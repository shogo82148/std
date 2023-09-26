// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Go execution tracer.
// The tracer captures a wide range of execution events like goroutine
// creation/blocking/unblocking, syscall enter/exit/block, GC-related events,
// changes of heap size, processor start/stop, etc and writes them to a buffer
// in a compact form. A precise nanosecond-precision timestamp and a stack
// trace is captured for most events.
// See https://golang.org/s/go15trace for more info.

package runtime

// Event types in the trace, args are given in square brackets.

// traceBlockReason is an enumeration of reasons a goroutine might block.
// This is the interface the rest of the runtime uses to tell the
// tracer why a goroutine blocked. The tracer then propagates this information
// into the trace however it sees fit.
//
// Note that traceBlockReasons should not be compared, since reasons that are
// distinct by name may *not* be distinct by value.

// For maximal efficiency, just map the trace block reason directly to a trace
// event.

// trace is global tracing context.

// gTraceState is per-G state for the tracer.

// mTraceState is per-M state for the tracer.

// pTraceState is per-P state for the tracer.

// traceBufHeader is per-P tracing buffer.

// traceBuf is per-P tracing buffer.

// traceBufPtr is a *traceBuf that is not traced by the garbage
// collector and doesn't have write barriers. traceBufs are not
// allocated from the GC'd heap, so this is safe, and are often
// manipulated in contexts where write barriers are not allowed, so
// this is necessary.
//
// TODO: Since traceBuf is now embedded runtime/internal/sys.NotInHeap, this isn't necessary.

// StartTrace enables tracing for the current process.
// While tracing, the data will be buffered and available via ReadTrace.
// StartTrace returns an error if tracing is already enabled.
// Most clients should use the runtime/trace package or the testing package's
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
func ReadTrace() []byte

// logicalStackSentinel is a sentinel value at pcBuf[0] signifying that
// pcBuf[1:] holds a logical stack requiring no further processing. Any other
// value at pcBuf[0] represents a skip value to apply to the physical stack in
// pcBuf[1:] after inline expansion.

// traceStackTable maps stack traces (arrays of PC's) to unique uint32 ids.
// It is lock-free for reading.

// traceStack is a single stack in traceStackTable.

// traceAlloc is a non-thread-safe region allocator.
// It holds a linked list of traceAllocBlock.

// traceAllocBlock is a block in traceAlloc.
//
// traceAllocBlock is allocated from non-GC'd memory, so it must not
// contain heap pointers. Writes to pointers to traceAllocBlocks do
// not need write barriers.

// TODO: Since traceAllocBlock is now embedded runtime/internal/sys.NotInHeap, this isn't necessary.

// traceTime represents a timestamp for the trace.
