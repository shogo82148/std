// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package traceviewer

import (
	"github.com/shogo82148/std/internal/trace"
	"github.com/shogo82148/std/internal/trace/traceviewer/format"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

type TraceConsumer struct {
	ConsumeTimeUnit    func(unit string)
	ConsumeViewerEvent func(v *format.Event, required bool)
	ConsumeViewerFrame func(key string, f format.Frame)
	Flush              func()
}

// ViewerDataTraceConsumer returns a TraceConsumer that writes to w. The
// startIdx and endIdx are used for splitting large traces. They refer to
// indexes in the traceEvents output array, not the events in the trace input.
func ViewerDataTraceConsumer(w io.Writer, startIdx, endIdx int64) TraceConsumer

func SplittingTraceConsumer(max int) (*splitter, TraceConsumer)

// WalkStackFrames calls fn for id and all of its parent frames from allFrames.
func WalkStackFrames(allFrames map[string]format.Frame, id int, fn func(id int))

type Mode int

const (
	ModeGoroutineOriented Mode = 1 << iota
	ModeTaskOriented
	ModeThreadOriented
)

// NewEmitter returns a new Emitter that writes to c. The rangeStart and
// rangeEnd args are used for splitting large traces.
func NewEmitter(c TraceConsumer, rangeStart, rangeEnd time.Duration) *Emitter

type Emitter struct {
	c          TraceConsumer
	rangeStart time.Duration
	rangeEnd   time.Duration

	heapStats, prevHeapStats     heapStats
	gstates, prevGstates         [gStateCount]int64
	threadStats, prevThreadStats [threadStateCount]int64
	gomaxprocs                   uint64
	frameTree                    frameNode
	frameSeq                     int
	arrowSeq                     uint64
	filter                       func(uint64) bool
	resourceType                 string
	resources                    map[uint64]string
	focusResource                uint64
	tasks                        map[uint64]task
	asyncSliceSeq                uint64
}

func (e *Emitter) Gomaxprocs(v uint64)

func (e *Emitter) Resource(id uint64, name string)

func (e *Emitter) SetResourceType(name string)

func (e *Emitter) SetResourceFilter(filter func(uint64) bool)

func (e *Emitter) Task(id uint64, name string, sortIndex int)

func (e *Emitter) Slice(s SliceEvent)

func (e *Emitter) TaskSlice(s SliceEvent)

type SliceEvent struct {
	Name     string
	Ts       time.Duration
	Dur      time.Duration
	Resource uint64
	Stack    int
	EndStack int
	Arg      any
}

func (e *Emitter) AsyncSlice(s AsyncSliceEvent)

type AsyncSliceEvent struct {
	SliceEvent
	Category       string
	Scope          string
	TaskColorIndex uint64
}

func (e *Emitter) Instant(i InstantEvent)

type InstantEvent struct {
	Ts       time.Duration
	Name     string
	Category string
	Resource uint64
	Stack    int
	Arg      any
}

func (e *Emitter) Arrow(a ArrowEvent)

func (e *Emitter) TaskArrow(a ArrowEvent)

type ArrowEvent struct {
	Name         string
	Start        time.Duration
	End          time.Duration
	FromResource uint64
	FromStack    int
	ToResource   uint64
}

func (e *Emitter) Event(ev *format.Event)

func (e *Emitter) HeapAlloc(ts time.Duration, v uint64)

func (e *Emitter) Focus(id uint64)

func (e *Emitter) GoroutineTransition(ts time.Duration, from, to GState)

func (e *Emitter) IncThreadStateCount(ts time.Duration, state ThreadState, delta int64)

func (e *Emitter) HeapGoal(ts time.Duration, v uint64)

// Err returns an error if the emitter is in an invalid state.
func (e *Emitter) Err() error

// OptionalEvent emits ev if it's within the time range of the consumer, i.e.
// the selected trace split range.
func (e *Emitter) OptionalEvent(ev *format.Event)

func (e *Emitter) Flush()

// Stack emits the given frames and returns a unique id for the stack. No
// pointers to the given data are being retained beyond the call to Stack.
func (e *Emitter) Stack(stk []*trace.Frame) int

type GState int

const (
	GDead GState = iota
	GRunnable
	GRunning
	GWaiting
	GWaitingGC
)

type ThreadState int

const (
	ThreadStateInSyscall ThreadState = iota
	ThreadStateInSyscallRuntime
	ThreadStateRunning
)
