// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
)

// Event describes one event in the trace.
type Event struct {
	Off   int
	Type  byte
	seq   int64
	Ts    int64
	P     int
	G     uint64
	StkID uint64
	Stk   []*Frame
	Args  [3]uint64
	SArgs []string
	// linked event (can be nil), depends on event type:
	// for GCStart: the GCStop
	// for GCSTWStart: the GCSTWDone
	// for GCSweepStart: the GCSweepDone
	// for GoCreate: first GoStart of the created goroutine
	// for GoStart/GoStartLabel: the associated GoEnd, GoBlock or other blocking event
	// for GoSched/GoPreempt: the next GoStart
	// for GoBlock and other blocking events: the unblock event
	// for GoUnblock: the associated GoStart
	// for blocking GoSysCall: the associated GoSysExit
	// for GoSysExit: the next GoStart
	// for GCMarkAssistStart: the associated GCMarkAssistDone
	// for UserTaskCreate: the UserTaskEnd
	// for UserRegion: if the start region, the corresponding UserRegion end event
	Link *Event
}

// Frame is a frame in stack traces.
type Frame struct {
	PC   uint64
	Fn   string
	File string
	Line int
}

const (
	// Special P identifiers:
	FakeP = 1000000 + iota
	TimerP
	NetpollP
	SyscallP
	GCP
	ProfileP
)

// ParseResult is the result of Parse.
type ParseResult struct {
	// Events is the sorted list of Events in the trace.
	Events []*Event
	// Stacks is the stack traces keyed by stack IDs from the trace.
	Stacks map[uint64][]*Frame
}

// Parse parses, post-processes and verifies the trace.
func Parse(r io.Reader, bin string) (ParseResult, error)

// ErrTimeOrder is returned by Parse when the trace contains
// time stamps that do not respect actual event ordering.
var ErrTimeOrder = fmt.Errorf("time stamps out of order")

// Print dumps events to stdout. For debugging.
func Print(events []*Event)

// PrintEvent dumps the event to stdout. For debugging.
func PrintEvent(ev *Event)

func (ev *Event) String() string

// BreakTimestampsForTesting causes the parser to randomly alter timestamps (for testing of broken cputicks).
var BreakTimestampsForTesting bool

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

var EventDescriptions = [EvCount]struct {
	Name       string
	minVersion int
	Stack      bool
	Args       []string
	SArgs      []string
}{
	EvNone:              {"None", 1005, false, []string{}, nil},
	EvBatch:             {"Batch", 1005, false, []string{"p", "ticks"}, nil},
	EvFrequency:         {"Frequency", 1005, false, []string{"freq"}, nil},
	EvStack:             {"Stack", 1005, false, []string{"id", "siz"}, nil},
	EvGomaxprocs:        {"Gomaxprocs", 1005, true, []string{"procs"}, nil},
	EvProcStart:         {"ProcStart", 1005, false, []string{"thread"}, nil},
	EvProcStop:          {"ProcStop", 1005, false, []string{}, nil},
	EvGCStart:           {"GCStart", 1005, true, []string{"seq"}, nil},
	EvGCDone:            {"GCDone", 1005, false, []string{}, nil},
	EvSTWStart:          {"STWStart", 1005, false, []string{"kindid"}, []string{"kind"}},
	EvSTWDone:           {"STWDone", 1005, false, []string{}, nil},
	EvGCSweepStart:      {"GCSweepStart", 1005, true, []string{}, nil},
	EvGCSweepDone:       {"GCSweepDone", 1005, false, []string{"swept", "reclaimed"}, nil},
	EvGoCreate:          {"GoCreate", 1005, true, []string{"g", "stack"}, nil},
	EvGoStart:           {"GoStart", 1005, false, []string{"g", "seq"}, nil},
	EvGoEnd:             {"GoEnd", 1005, false, []string{}, nil},
	EvGoStop:            {"GoStop", 1005, true, []string{}, nil},
	EvGoSched:           {"GoSched", 1005, true, []string{}, nil},
	EvGoPreempt:         {"GoPreempt", 1005, true, []string{}, nil},
	EvGoSleep:           {"GoSleep", 1005, true, []string{}, nil},
	EvGoBlock:           {"GoBlock", 1005, true, []string{}, nil},
	EvGoUnblock:         {"GoUnblock", 1005, true, []string{"g", "seq"}, nil},
	EvGoBlockSend:       {"GoBlockSend", 1005, true, []string{}, nil},
	EvGoBlockRecv:       {"GoBlockRecv", 1005, true, []string{}, nil},
	EvGoBlockSelect:     {"GoBlockSelect", 1005, true, []string{}, nil},
	EvGoBlockSync:       {"GoBlockSync", 1005, true, []string{}, nil},
	EvGoBlockCond:       {"GoBlockCond", 1005, true, []string{}, nil},
	EvGoBlockNet:        {"GoBlockNet", 1005, true, []string{}, nil},
	EvGoSysCall:         {"GoSysCall", 1005, true, []string{}, nil},
	EvGoSysExit:         {"GoSysExit", 1005, false, []string{"g", "seq", "ts"}, nil},
	EvGoSysBlock:        {"GoSysBlock", 1005, false, []string{}, nil},
	EvGoWaiting:         {"GoWaiting", 1005, false, []string{"g"}, nil},
	EvGoInSyscall:       {"GoInSyscall", 1005, false, []string{"g"}, nil},
	EvHeapAlloc:         {"HeapAlloc", 1005, false, []string{"mem"}, nil},
	EvHeapGoal:          {"HeapGoal", 1005, false, []string{"mem"}, nil},
	EvTimerGoroutine:    {"TimerGoroutine", 1005, false, []string{"g"}, nil},
	EvFutileWakeup:      {"FutileWakeup", 1005, false, []string{}, nil},
	EvString:            {"String", 1007, false, []string{}, nil},
	EvGoStartLocal:      {"GoStartLocal", 1007, false, []string{"g"}, nil},
	EvGoUnblockLocal:    {"GoUnblockLocal", 1007, true, []string{"g"}, nil},
	EvGoSysExitLocal:    {"GoSysExitLocal", 1007, false, []string{"g", "ts"}, nil},
	EvGoStartLabel:      {"GoStartLabel", 1008, false, []string{"g", "seq", "labelid"}, []string{"label"}},
	EvGoBlockGC:         {"GoBlockGC", 1008, true, []string{}, nil},
	EvGCMarkAssistStart: {"GCMarkAssistStart", 1009, true, []string{}, nil},
	EvGCMarkAssistDone:  {"GCMarkAssistDone", 1009, false, []string{}, nil},
	EvUserTaskCreate:    {"UserTaskCreate", 1011, true, []string{"taskid", "pid", "typeid"}, []string{"name"}},
	EvUserTaskEnd:       {"UserTaskEnd", 1011, true, []string{"taskid"}, nil},
	EvUserRegion:        {"UserRegion", 1011, true, []string{"taskid", "mode", "typeid"}, []string{"name"}},
	EvUserLog:           {"UserLog", 1011, true, []string{"id", "keyid"}, []string{"category", "message"}},
	EvCPUSample:         {"CPUSample", 1019, true, []string{"ts", "p", "g"}, nil},
}
