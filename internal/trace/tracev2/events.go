// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tracev2

// Event types in the trace, args are given in square brackets.
//
// Naming scheme:
//   - Time range event pairs have suffixes "Begin" and "End".
//   - "Start", "Stop", "Create", "Destroy", "Block", "Unblock"
//     are suffixes reserved for scheduling resources.
//
// NOTE: If you add an event type, make sure you also update all
// tables in this file!
const (
	EvNone EventType = iota

	// Structural events.
	EvEventBatch
	EvStacks
	EvStack
	EvStrings
	EvString
	EvCPUSamples
	EvCPUSample
	EvFrequency

	// Procs.
	EvProcsChange
	EvProcStart
	EvProcStop
	EvProcSteal
	EvProcStatus

	// Goroutines.
	EvGoCreate
	EvGoCreateSyscall
	EvGoStart
	EvGoDestroy
	EvGoDestroySyscall
	EvGoStop
	EvGoBlock
	EvGoUnblock
	EvGoSyscallBegin
	EvGoSyscallEnd
	EvGoSyscallEndBlocked
	EvGoStatus

	// STW.
	EvSTWBegin
	EvSTWEnd

	// GC events.
	EvGCActive
	EvGCBegin
	EvGCEnd
	EvGCSweepActive
	EvGCSweepBegin
	EvGCSweepEnd
	EvGCMarkAssistActive
	EvGCMarkAssistBegin
	EvGCMarkAssistEnd
	EvHeapAlloc
	EvHeapGoal

	// Annotations.
	EvGoLabel
	EvUserTaskBegin
	EvUserTaskEnd
	EvUserRegionBegin
	EvUserRegionEnd
	EvUserLog

	// Coroutines. Added in Go 1.23.
	EvGoSwitch
	EvGoSwitchDestroy
	EvGoCreateBlocked

	// GoStatus with stack. Added in Go 1.23.
	EvGoStatusStack

	// Batch event for an experimental batch with a custom format. Added in Go 1.23.
	EvExperimentalBatch

	// Sync batch. Added in Go 1.25. Previously a lone EvFrequency event.
	EvSync
	EvClockSnapshot

	NumEvents
)

func (ev EventType) Experimental() bool

// Experiments.
const (
	// AllocFree is the alloc-free events experiment.
	AllocFree Experiment = 1 + iota

	NumExperiments
)

func Experiments() []string

// Experimental events.
const (
	MaxEvent EventType = 127 + iota

	// Experimental heap span events. Added in Go 1.23.
	EvSpan
	EvSpanAlloc
	EvSpanFree

	// Experimental heap object events. Added in Go 1.23.
	EvHeapObject
	EvHeapObjectAlloc
	EvHeapObjectFree

	// Experimental goroutine stack events. Added in Go 1.23.
	EvGoroutineStack
	EvGoroutineStackAlloc
	EvGoroutineStackFree

	MaxExperimentalEvent
)

const NumExperimentalEvents = MaxExperimentalEvent - MaxEvent

// MaxTimedEventArgs is the maximum number of arguments for timed events.
const MaxTimedEventArgs = 5

func Specs() []EventSpec

// GoStatus is the status of a goroutine.
//
// They correspond directly to the various goroutine states.
type GoStatus uint8

const (
	GoBad GoStatus = iota
	GoRunnable
	GoRunning
	GoSyscall
	GoWaiting
)

func (s GoStatus) String() string

// ProcStatus is the status of a P.
//
// They mostly correspond to the various P states.
type ProcStatus uint8

const (
	ProcBad ProcStatus = iota
	ProcRunning
	ProcIdle
	ProcSyscall

	// ProcSyscallAbandoned is a special case of
	// ProcSyscall. It's used in the very specific case
	// where the first a P is mentioned in a generation is
	// part of a ProcSteal event. If that's the first time
	// it's mentioned, then there's no GoSyscallBegin to
	// connect the P stealing back to at that point. This
	// special state indicates this to the parser, so it
	// doesn't try to find a GoSyscallEndBlocked that
	// corresponds with the ProcSteal.
	ProcSyscallAbandoned
)

func (s ProcStatus) String() string

const (
	// MaxBatchSize sets the maximum size that a batch can be.
	//
	// Directly controls the trace batch size in the runtime.
	//
	// NOTE: If this number decreases, the trace format version must change.
	MaxBatchSize = 64 << 10

	// Maximum number of PCs in a single stack trace.
	//
	// Since events contain only stack ID rather than whole stack trace,
	// we can allow quite large values here.
	//
	// Directly controls the maximum number of frames per stack
	// in the runtime.
	//
	// NOTE: If this number decreases, the trace format version must change.
	MaxFramesPerStack = 128

	// MaxEventTrailerDataSize controls the amount of trailer data that
	// an event can have in bytes. Must be smaller than MaxBatchSize.
	// Controls the maximum string size in the trace.
	//
	// Directly controls the maximum such value in the runtime.
	//
	// NOTE: If this number decreases, the trace format version must change.
	MaxEventTrailerDataSize = 1 << 10
)
