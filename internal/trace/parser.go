// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

// Frame is a frame in stack traces.
type Frame struct {
	PC   uint64
	Fn   string
	File string
	Line int
}

const (
	// Special P identifiers:
	FakeP    = 1000000 + iota
	TimerP
	NetpollP
	SyscallP
	GCP
	ProfileP
)

// Event types in the trace.
// Verbatim copy from src/runtime/trace.go with the "trace" prefix removed.
const (
	EvNone              = 0
	EvBatch             = 1
	EvFrequency         = 2
	EvStack             = 3
	EvGomaxprocs        = 4
	EvProcStart         = 5
	EvProcStop          = 6
	EvGCStart           = 7
	EvGCDone            = 8
	EvSTWStart          = 9
	EvSTWDone           = 10
	EvGCSweepStart      = 11
	EvGCSweepDone       = 12
	EvGoCreate          = 13
	EvGoStart           = 14
	EvGoEnd             = 15
	EvGoStop            = 16
	EvGoSched           = 17
	EvGoPreempt         = 18
	EvGoSleep           = 19
	EvGoBlock           = 20
	EvGoUnblock         = 21
	EvGoBlockSend       = 22
	EvGoBlockRecv       = 23
	EvGoBlockSelect     = 24
	EvGoBlockSync       = 25
	EvGoBlockCond       = 26
	EvGoBlockNet        = 27
	EvGoSysCall         = 28
	EvGoSysExit         = 29
	EvGoSysBlock        = 30
	EvGoWaiting         = 31
	EvGoInSyscall       = 32
	EvHeapAlloc         = 33
	EvHeapGoal          = 34
	EvTimerGoroutine    = 35
	EvFutileWakeup      = 36
	EvString            = 37
	EvGoStartLocal      = 38
	EvGoUnblockLocal    = 39
	EvGoSysExitLocal    = 40
	EvGoStartLabel      = 41
	EvGoBlockGC         = 42
	EvGCMarkAssistStart = 43
	EvGCMarkAssistDone  = 44
	EvUserTaskCreate    = 45
	EvUserTaskEnd       = 46
	EvUserRegion        = 47
	EvUserLog           = 48
	EvCPUSample         = 49
	EvCount             = 50
)
