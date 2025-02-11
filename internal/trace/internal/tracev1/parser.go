// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tracev1 implements a parser for Go execution traces from versions
// 1.11–1.21.
//
// The package started as a copy of Go 1.19's internal/trace, but has been
// optimized to be faster while using less memory and fewer allocations. It has
// been further modified for the specific purpose of converting traces to the
// new 1.22+ format.
package tracev1

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/internal/trace/version"
	"github.com/shogo82148/std/io"
)

// Timestamp represents a count of nanoseconds since the beginning of the trace.
// They can only be meaningfully compared with other timestamps from the same
// trace.
type Timestamp int64

// Event describes one event in the trace.
type Event struct {
	Ts    Timestamp
	G     uint64
	Args  [4]uint64
	StkID uint32
	P     int32
	Type  EventType
}

// Frame is a frame in stack traces.
type Frame struct {
	PC uint64
	// string ID of the function name
	Fn uint64
	// string ID of the file name
	File uint64
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

// Trace is the result of Parse.
type Trace struct {
	Version version.Version

	// Events is the sorted list of Events in the trace.
	Events Events
	// Stacks is the stack traces (stored as slices of PCs), keyed by stack IDs
	// from the trace.
	Stacks        map[uint32][]uint64
	PCs           map[uint64]Frame
	Strings       map[uint64]string
	InlineStrings []string
}

// Parse parses Go execution traces from versions 1.11–1.21. The provided reader
// will be read to completion and the entire trace will be materialized in
// memory. That is, this function does not allow incremental parsing.
//
// The reader has to be positioned just after the trace header and vers needs to
// be the version of the trace. This can be achieved by using
// version.ReadHeader.
func Parse(r io.Reader, vers version.Version) (Trace, error)

type Events struct {
	// Events is a slice of slices that grows one slice of size eventsBucketSize
	// at a time. This avoids the O(n) cost of slice growth in append, and
	// additionally allows consumers to drop references to parts of the data,
	// freeing memory piecewise.
	n       int
	buckets []*[eventsBucketSize]Event
	off     int
}

func (l *Events) Ptr(i int) *Event

func (l *Events) Len() int

func (l *Events) Less(i, j int) bool

func (l *Events) Swap(i, j int)

func (l *Events) Pop() (*Event, bool)

func (l *Events) Peek() (*Event, bool)

func (l *Events) All() func(yield func(ev *Event) bool)

// ErrTimeOrder is returned by Parse when the trace contains
// time stamps that do not respect actual event ordering.
var ErrTimeOrder = errors.New("time stamps out of order")

func (ev *Event) String() string

type EventType uint8

// Event types in the trace.
// Verbatim copy from src/runtime/trace.go with the "trace" prefix removed.
const (
	EvNone              EventType = 0
	EvBatch             EventType = 1
	EvFrequency         EventType = 2
	EvStack             EventType = 3
	EvGomaxprocs        EventType = 4
	EvProcStart         EventType = 5
	EvProcStop          EventType = 6
	EvGCStart           EventType = 7
	EvGCDone            EventType = 8
	EvSTWStart          EventType = 9
	EvSTWDone           EventType = 10
	EvGCSweepStart      EventType = 11
	EvGCSweepDone       EventType = 12
	EvGoCreate          EventType = 13
	EvGoStart           EventType = 14
	EvGoEnd             EventType = 15
	EvGoStop            EventType = 16
	EvGoSched           EventType = 17
	EvGoPreempt         EventType = 18
	EvGoSleep           EventType = 19
	EvGoBlock           EventType = 20
	EvGoUnblock         EventType = 21
	EvGoBlockSend       EventType = 22
	EvGoBlockRecv       EventType = 23
	EvGoBlockSelect     EventType = 24
	EvGoBlockSync       EventType = 25
	EvGoBlockCond       EventType = 26
	EvGoBlockNet        EventType = 27
	EvGoSysCall         EventType = 28
	EvGoSysExit         EventType = 29
	EvGoSysBlock        EventType = 30
	EvGoWaiting         EventType = 31
	EvGoInSyscall       EventType = 32
	EvHeapAlloc         EventType = 33
	EvHeapGoal          EventType = 34
	EvTimerGoroutine    EventType = 35
	EvFutileWakeup      EventType = 36
	EvString            EventType = 37
	EvGoStartLocal      EventType = 38
	EvGoUnblockLocal    EventType = 39
	EvGoSysExitLocal    EventType = 40
	EvGoStartLabel      EventType = 41
	EvGoBlockGC         EventType = 42
	EvGCMarkAssistStart EventType = 43
	EvGCMarkAssistDone  EventType = 44
	EvUserTaskCreate    EventType = 45
	EvUserTaskEnd       EventType = 46
	EvUserRegion        EventType = 47
	EvUserLog           EventType = 48
	EvCPUSample         EventType = 49
	EvCount             EventType = 50
)

var EventDescriptions = [256]struct {
	Name       string
	minVersion version.Version
	Stack      bool
	Args       []string
	SArgs      []string
}{
	EvNone:              {"None", 5, false, []string{}, nil},
	EvBatch:             {"Batch", 5, false, []string{"p", "ticks"}, nil},
	EvFrequency:         {"Frequency", 5, false, []string{"freq"}, nil},
	EvStack:             {"Stack", 5, false, []string{"id", "siz"}, nil},
	EvGomaxprocs:        {"Gomaxprocs", 5, true, []string{"procs"}, nil},
	EvProcStart:         {"ProcStart", 5, false, []string{"thread"}, nil},
	EvProcStop:          {"ProcStop", 5, false, []string{}, nil},
	EvGCStart:           {"GCStart", 5, true, []string{"seq"}, nil},
	EvGCDone:            {"GCDone", 5, false, []string{}, nil},
	EvSTWStart:          {"GCSTWStart", 5, false, []string{"kindid"}, []string{"kind"}},
	EvSTWDone:           {"GCSTWDone", 5, false, []string{}, nil},
	EvGCSweepStart:      {"GCSweepStart", 5, true, []string{}, nil},
	EvGCSweepDone:       {"GCSweepDone", 5, false, []string{"swept", "reclaimed"}, nil},
	EvGoCreate:          {"GoCreate", 5, true, []string{"g", "stack"}, nil},
	EvGoStart:           {"GoStart", 5, false, []string{"g", "seq"}, nil},
	EvGoEnd:             {"GoEnd", 5, false, []string{}, nil},
	EvGoStop:            {"GoStop", 5, true, []string{}, nil},
	EvGoSched:           {"GoSched", 5, true, []string{}, nil},
	EvGoPreempt:         {"GoPreempt", 5, true, []string{}, nil},
	EvGoSleep:           {"GoSleep", 5, true, []string{}, nil},
	EvGoBlock:           {"GoBlock", 5, true, []string{}, nil},
	EvGoUnblock:         {"GoUnblock", 5, true, []string{"g", "seq"}, nil},
	EvGoBlockSend:       {"GoBlockSend", 5, true, []string{}, nil},
	EvGoBlockRecv:       {"GoBlockRecv", 5, true, []string{}, nil},
	EvGoBlockSelect:     {"GoBlockSelect", 5, true, []string{}, nil},
	EvGoBlockSync:       {"GoBlockSync", 5, true, []string{}, nil},
	EvGoBlockCond:       {"GoBlockCond", 5, true, []string{}, nil},
	EvGoBlockNet:        {"GoBlockNet", 5, true, []string{}, nil},
	EvGoSysCall:         {"GoSysCall", 5, true, []string{}, nil},
	EvGoSysExit:         {"GoSysExit", 5, false, []string{"g", "seq", "ts"}, nil},
	EvGoSysBlock:        {"GoSysBlock", 5, false, []string{}, nil},
	EvGoWaiting:         {"GoWaiting", 5, false, []string{"g"}, nil},
	EvGoInSyscall:       {"GoInSyscall", 5, false, []string{"g"}, nil},
	EvHeapAlloc:         {"HeapAlloc", 5, false, []string{"mem"}, nil},
	EvHeapGoal:          {"HeapGoal", 5, false, []string{"mem"}, nil},
	EvTimerGoroutine:    {"TimerGoroutine", 5, false, []string{"g"}, nil},
	EvFutileWakeup:      {"FutileWakeup", 5, false, []string{}, nil},
	EvString:            {"String", 7, false, []string{}, nil},
	EvGoStartLocal:      {"GoStartLocal", 7, false, []string{"g"}, nil},
	EvGoUnblockLocal:    {"GoUnblockLocal", 7, true, []string{"g"}, nil},
	EvGoSysExitLocal:    {"GoSysExitLocal", 7, false, []string{"g", "ts"}, nil},
	EvGoStartLabel:      {"GoStartLabel", 8, false, []string{"g", "seq", "labelid"}, []string{"label"}},
	EvGoBlockGC:         {"GoBlockGC", 8, true, []string{}, nil},
	EvGCMarkAssistStart: {"GCMarkAssistStart", 9, true, []string{}, nil},
	EvGCMarkAssistDone:  {"GCMarkAssistDone", 9, false, []string{}, nil},
	EvUserTaskCreate:    {"UserTaskCreate", 11, true, []string{"taskid", "pid", "typeid"}, []string{"name"}},
	EvUserTaskEnd:       {"UserTaskEnd", 11, true, []string{"taskid"}, nil},
	EvUserRegion:        {"UserRegion", 11, true, []string{"taskid", "mode", "typeid"}, []string{"name"}},
	EvUserLog:           {"UserLog", 11, true, []string{"id", "keyid"}, []string{"category", "message"}},
	EvCPUSample:         {"CPUSample", 19, true, []string{"ts", "p", "g"}, nil},
}

func (tr *Trace) STWReason(kindID uint64) STWReason

type STWReason int

const (
	STWUnknown                 STWReason = 0
	STWGCMarkTermination       STWReason = 1
	STWGCSweepTermination      STWReason = 2
	STWWriteHeapDump           STWReason = 3
	STWGoroutineProfile        STWReason = 4
	STWGoroutineProfileCleanup STWReason = 5
	STWAllGoroutinesStackTrace STWReason = 6
	STWReadMemStats            STWReason = 7
	STWAllThreadsSyscall       STWReason = 8
	STWGOMAXPROCS              STWReason = 9
	STWStartTrace              STWReason = 10
	STWStopTrace               STWReason = 11
	STWCountPagesInUse         STWReason = 12
	STWReadMetricsSlow         STWReason = 13
	STWReadMemStatsSlow        STWReason = 14
	STWPageCachePagesLeaked    STWReason = 15
	STWResetDebugLog           STWReason = 16

	NumSTWReasons = 17
)
