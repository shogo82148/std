// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go122

import (
	"github.com/shogo82148/std/internal/trace/event"
)

const (
	EvNone event.Type = iota

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
)

// Experiments.
const (
	// AllocFree is the alloc-free events experiment.
	AllocFree event.Experiment = 1 + iota
)

// Experimental events.
const (
	_ event.Type = 127 + iota

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
)

// EventString returns the name of a Go 1.22 event.
func EventString(typ event.Type) string

func Specs() []event.Spec

type GoStatus uint8

const (
	GoBad GoStatus = iota
	GoRunnable
	GoRunning
	GoSyscall
	GoWaiting
)

func (s GoStatus) String() string

type ProcStatus uint8

const (
	ProcBad ProcStatus = iota
	ProcRunning
	ProcIdle
	ProcSyscall
	ProcSyscallAbandoned
)

func (s ProcStatus) String() string

const (
	// Various format-specific constants.
	MaxBatchSize      = 64 << 10
	MaxFramesPerStack = 128
	MaxStringSize     = 1 << 10
)
