// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// # Go execution tracer
//
// The tracer captures a wide range of execution events like goroutine
// creation/blocking/unblocking, syscall enter/exit/block, GC-related events,
// changes of heap size, processor start/stop, etc and writes them to a buffer
// in a compact form. A precise nanosecond-precision timestamp and a stack
// trace is captured for most events.
//
// ## Design
//
// The basic idea behind the execution tracer is to have per-M buffers that
// trace data may be written into. Each M maintains a write flag indicating whether
// its trace buffer is currently in use.
//
// Tracing is initiated by StartTrace, and proceeds in "generations," with each
// generation being marked by a call to traceAdvance, to advance to the next
// generation. Generations are a global synchronization point for trace data,
// and we proceed to a new generation by moving forward trace.gen. Each M reads
// trace.gen under its own write flag to determine which generation it is writing
// trace data for. To this end, each M has 2 slots for buffers: one slot for the
// previous generation, one slot for the current one. It uses tl.gen to select
// which buffer slot to write to. Simultaneously, traceAdvance uses the write flag
// to determine whether every thread is guaranteed to observe an updated
// trace.gen. Once it is sure, it may then flush any buffers that are left over
// from the previous generation safely, since it knows the Ms will not mutate
// it.
//
// Flushed buffers are processed by the ReadTrace function, which is called by
// the trace reader goroutine. The first goroutine to call ReadTrace is designated
// as the trace reader goroutine until tracing completes. (There may only be one at
// a time.)
//
// Once all buffers are flushed, any extra post-processing complete, and flushed
// buffers are processed by the trace reader goroutine, the trace emits an
// EndOfGeneration event to mark the global synchronization point in the trace.
//
// All other trace features, including CPU profile samples, stack information,
// string tables, etc. all revolve around this generation system, and typically
// appear in pairs: one for the previous generation, and one for the current one.
// Like the per-M buffers, which of the two is written to is selected using trace.gen,
// and anything managed this way must similarly be mutated only in traceAdvance or
// under the M's write flag.
//
// Trace events themselves are simple. They consist of a single byte for the event type,
// followed by zero or more LEB128-encoded unsigned varints. They are decoded using
// a pre-determined table for each trace version: internal/trace/tracev2.specs.
//
// To avoid relying on timestamps for correctness and validation, each G and P have
// sequence counters that are written into trace events to encode a partial order.
// The sequence counters reset on each generation. Ms do not need sequence counters
// because they are the source of truth for execution: trace events, and even whole
// buffers, are guaranteed to appear in order in the trace data stream, simply because
// that's the order the thread emitted them in.
//
// See traceruntime.go for the API the tracer exposes to the runtime for emitting events.
//
// In each generation, we ensure that we enumerate all goroutines, such that each
// generation's data is fully self-contained. This makes features like the flight
// recorder easy to implement. To this end, we guarantee that every live goroutine is
// listed at least once by emitting a status event for the goroutine, indicating its
// starting state. These status events are emitted based on context, generally based
// on the event that's about to be emitted.
//
// The traceEventWriter type encapsulates these details, and is the backbone of
// the API exposed in traceruntime.go, though there are deviations where necessary.
//
// This is the overall design, but as always, there are many details. Beyond this,
// look to the invariants and select corner cases below and the code itself for the
// source of truth.
//
// See https://go.dev/issue/60773 for a link to a more complete design with rationale,
// though parts of it are out-of-date.
//
// ## Invariants
//
// 1. An m that has a trace buffer MUST be on either the allm or sched.freem lists.
//
// Otherwise, traceAdvance might miss an M with a buffer that needs to be flushed.
//
// 2. Trace buffers MUST only be mutated in traceAdvance or under a traceAcquire/traceRelease.
//
// Otherwise, traceAdvance may race with Ms writing trace data when trying to flush buffers.
//
// 3. traceAdvance MUST NOT return until all of the current generation's buffers are flushed.
//
// Otherwise, callers cannot rely on all the data they need being available (for example, for
// the flight recorder).
//
// 4. P and goroutine state transition events MUST be emitted by an M that owns its ability
//    to transition.
//
// What this means is that the M must either be the owner of the P, the owner of the goroutine,
// or owner of a non-running goroutine's _Gscan bit. There are a lot of bad things that can
// happen if this invariant isn't maintained, mostly around generating inconsistencies in the
// trace due to racy emission of events.
//
// 5. Acquisition of a P (pidleget or takeP/gcstopP) MUST NOT be performed under a traceAcquire/traceRelease pair.
//
// Notably, it's important that traceAcquire/traceRelease not cover a state in which the
// goroutine or P is not yet owned. For example, if traceAcquire is held across both wirep and
// pidleget, then we could end up emitting an event in the wrong generation. Suppose T1
// traceAcquires in generation 1, a generation transition happens, T2 emits a ProcStop and
// executes pidleput in generation 2, and finally T1 calls pidleget and emits ProcStart.
// The ProcStart must follow the ProcStop in the trace to make any sense, but ProcStop was
// emitted in a latter generation.
//
// 6. Goroutine state transitions, with the exception of transitions into _Grunning, MUST be
//    performed under the traceAcquire/traceRelease pair where the event is emitted.
//
// Otherwise, traceAdvance may observe a goroutine state that is inconsistent with the
// events being emitted. traceAdvance inspects all goroutines' states in order to emit
// a status event for any goroutine that did not have an event emitted for it already.
// If the generation then advances in between that observation and the event being emitted,
// then the trace will contain a status that doesn't line up with the event. For example,
// if the event is emitted after the state transition _Gwaiting -> _Grunnable, then
// traceAdvance may observe the goroutine in _Grunnable, emit a status event, advance the
// generation, and the following generation contains a GoUnblock event. The trace parser
// will get confused because it sees that goroutine in _Grunnable in the previous generation
// trying to be transitioned from _Gwaiting into _Grunnable in the following one. Something
// similar happens if the trace event is emitted before the state transition, so that does
// not help either.
//
// Transitions to _Grunning do not have the same problem because traceAdvance is unable to
// observe running goroutines directly. It must stop them, or wait for them to emit an event.
// Note that it cannot even stop them with asynchronous preemption in any "bad" window between
// the state transition to _Grunning and the event emission because async preemption cannot
// stop goroutines in the runtime.
//
// 7. Goroutine state transitions into _Grunning MUST emit an event for the transition after
//    the state transition.
//
// This follows from invariants (4), (5), and the explanation of (6).
// The relevant part of the previous invariant is that in order for the tracer to be unable to
// stop a goroutine, it must be in _Grunning and in the runtime. So to close any windows between
// event emission and the state transition, the event emission must happen *after* the transition
// to _Grunning.
//
// ## Select corner cases
//
// ### CGO calls / system calls
//
// CGO calls and system calls are mostly straightforward, except for P stealing. For historical
// reasons, this introduces a new trace-level P state called ProcSyscall which used to model
// _Psyscall (now _Psyscall_unused). This state is used to indicate in the trace that a P
// is eligible for stealing as part of the parser's ordering logic.
//
// Another quirk of this corner case is the ProcSyscallAbandoned trace-level P state, which
// is used only in status events to indicate a relaxation of verification requirements. It
// means that if the execution trace parser can't find the corresponding thread that the P
// was stolen from in the state it expects it to be, to accept the trace anyway. This is also
// historical. When _Psyscall still existed, one would steal and then ProcSteal, and there
// was no ordering between the ProcSteal and the subsequent GoSyscallEndBlocked. One clearly
// happened before the other, but since P stealing was a single atomic, there was no way
// to enforce the order. The GoSyscallEndBlocked thread could move on and end up in any
// state, and the GoSyscallEndBlocked could be in a completely different generation to the
// ProcSteal. Today this is no longer possible as the ProcSteal is always ordered before
// the GoSyscallEndBlocked event in the runtime.
//
// Both ProcSyscall and ProcSyscallAbandoned are likely no longer be necessary.
//
// ### CGO callbacks
//
// When a C thread calls into Go, the execution tracer models that as the creation of a new
// goroutine. When the thread exits back into C, that is modeled as the destruction of that
// goroutine. These are the GoCreateSyscall and GoDestroySyscall events, which represent the
// creation and destruction of a goroutine with its starting and ending states being _Gsyscall.
//
// This model is simple to reason about but contradicts the runtime implementation, which
// doesn't do this directly for performance reasons. The runtime implementation instead caches
// a G on the M created for the C thread. On Linux this M is then cached in the thread's TLS,
// and on other systems, the M is put on a global list on exit from Go. We need to do some
// extra work to make sure that this is modeled correctly in the tracer. For example,
// a C thread exiting Go may leave a P hanging off of its M (whether that M is kept in TLS
// or placed back on a list). In order to correctly model goroutine creation and destruction,
// we must behave as if the P was at some point stolen by the runtime, if the C thread
// reenters Go with the same M (and thus, same P) once more.

package runtime

// StartTrace enables tracing for the current process.
// While tracing, the data will be buffered and available via [ReadTrace].
// StartTrace returns an error if tracing is already enabled.
// Most clients should use the [runtime/trace] package or the [testing] package's
// -test.trace flag instead of calling StartTrace directly.
func StartTrace() error

// StopTrace stops tracing, if it was previously enabled.
// StopTrace only returns after all the reads for the trace have completed.
func StopTrace()

// ReadTrace returns the next chunk of binary tracing data, blocking until data
// is available. If tracing is turned off and all the data accumulated while it
// was on has been returned, ReadTrace returns nil. The caller must copy the
// returned data before calling ReadTrace again.
// ReadTrace must be called from one goroutine at a time.
func ReadTrace() (buf []byte)
