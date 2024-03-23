// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package go122

import (
	"github.com/shogo82148/std/internal/trace/v2/event"
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
