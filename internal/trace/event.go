// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/iter"
	"github.com/shogo82148/std/time"
)

// EventKind indicates the kind of event this is.
//
// Use this information to obtain a more specific event that
// allows access to more detailed information.
type EventKind uint16

const (
	EventBad EventKind = iota

	// EventKindSync is an event that indicates a global synchronization
	// point in the trace. At the point of a sync event, the
	// trace reader can be certain that all resources (e.g. threads,
	// goroutines) that have existed until that point have been enumerated.
	EventSync

	// EventMetric is an event that represents the value of a metric at
	// a particular point in time.
	EventMetric

	// EventLabel attaches a label to a resource.
	EventLabel

	// EventStackSample represents an execution sample, indicating what a
	// thread/proc/goroutine was doing at a particular point in time via
	// its backtrace.
	//
	// Note: Samples should be considered a close approximation of
	// what a thread/proc/goroutine was executing at a given point in time.
	// These events may slightly contradict the situation StateTransitions
	// describe, so they should only be treated as a best-effort annotation.
	EventStackSample

	// EventRangeBegin and EventRangeEnd are a pair of generic events representing
	// a special range of time. Ranges are named and scoped to some resource
	// (identified via ResourceKind). A range that has begun but has not ended
	// is considered active.
	//
	// EvRangeBegin and EvRangeEnd will share the same name, and an End will always
	// follow a Begin on the same instance of the resource. The associated
	// resource ID can be obtained from the Event. ResourceNone indicates the
	// range is globally scoped. That is, any goroutine/proc/thread can start or
	// stop, but only one such range may be active at any given time.
	//
	// EventRangeActive is like EventRangeBegin, but indicates that the range was
	// already active. In this case, the resource referenced may not be in the current
	// context.
	EventRangeBegin
	EventRangeActive
	EventRangeEnd

	// EvTaskBegin and EvTaskEnd are a pair of events representing a runtime/trace.Task.
	EventTaskBegin
	EventTaskEnd

	// EventRegionBegin and EventRegionEnd are a pair of events represent a runtime/trace.Region.
	EventRegionBegin
	EventRegionEnd

	// EventLog represents a runtime/trace.Log call.
	EventLog

	// EventStateTransition represents a state change for some resource.
	EventStateTransition

	// EventExperimental is an experimental event that is unvalidated and exposed in a raw form.
	// Users are expected to understand the format and perform their own validation. These events
	// may always be safely ignored.
	EventExperimental
)

// String returns a string form of the EventKind.
func (e EventKind) String() string

// Time is a timestamp in nanoseconds.
//
// It corresponds to the monotonic clock on the platform that the
// trace was taken, and so is possible to correlate with timestamps
// for other traces taken on the same machine using the same clock
// (i.e. no reboots in between).
//
// The actual absolute value of the timestamp is only meaningful in
// relation to other timestamps from the same clock.
//
// BUG: Timestamps coming from traces on Windows platforms are
// only comparable with timestamps from the same trace. Timestamps
// across traces cannot be compared, because the system clock is
// not used as of Go 1.22.
//
// BUG: Traces produced by Go versions 1.21 and earlier cannot be
// compared with timestamps from other traces taken on the same
// machine. This is because the system clock was not used at all
// to collect those timestamps.
type Time int64

// Sub subtracts t0 from t, returning the duration in nanoseconds.
func (t Time) Sub(t0 Time) time.Duration

// Metric provides details about a Metric event.
type Metric struct {
	// Name is the name of the sampled metric.
	//
	// Names follow the same convention as metric names in the
	// runtime/metrics package, meaning they include the unit.
	// Names that match with the runtime/metrics package represent
	// the same quantity. Note that this corresponds to the
	// runtime/metrics package for the Go version this trace was
	// collected for.
	Name string

	// Value is the sampled value of the metric.
	//
	// The Value's Kind is tied to the name of the metric, and so is
	// guaranteed to be the same for metric samples for the same metric.
	Value Value
}

// Label provides details about a Label event.
type Label struct {
	// Label is the label applied to some resource.
	Label string

	// Resource is the resource to which this label should be applied.
	Resource ResourceID
}

// Range provides details about a Range event.
type Range struct {
	// Name is a human-readable name for the range.
	//
	// This name can be used to identify the end of the range for the resource
	// its scoped to, because only one of each type of range may be active on
	// a particular resource. The relevant resource should be obtained from the
	// Event that produced these details. The corresponding RangeEnd will have
	// an identical name.
	Name string

	// Scope is the resource that the range is scoped to.
	//
	// For example, a ResourceGoroutine scope means that the same goroutine
	// must have a start and end for the range, and that goroutine can only
	// have one range of a particular name active at any given time. The
	// ID that this range is scoped to may be obtained via Event.Goroutine.
	//
	// The ResourceNone scope means that the range is globally scoped. As a
	// result, any goroutine/proc/thread may start or end the range, and only
	// one such named range may be active globally at any given time.
	//
	// For RangeBegin and RangeEnd events, this will always reference some
	// resource ID in the current execution context. For RangeActive events,
	// this may reference a resource not in the current context. Prefer Scope
	// over the current execution context.
	Scope ResourceID
}

// RangeAttributes provides attributes about a completed Range.
type RangeAttribute struct {
	// Name is the human-readable name for the range.
	Name string

	// Value is the value of the attribute.
	Value Value
}

// TaskID is the internal ID of a task used to disambiguate tasks (even if they
// are of the same type).
type TaskID uint64

const (
	// NoTask indicates the lack of a task.
	NoTask = TaskID(^uint64(0))

	// BackgroundTask is the global task that events are attached to if there was
	// no other task in the context at the point the event was emitted.
	BackgroundTask = TaskID(0)
)

// Task provides details about a Task event.
type Task struct {
	// ID is a unique identifier for the task.
	//
	// This can be used to associate the beginning of a task with its end.
	ID TaskID

	// ParentID is the ID of the parent task.
	Parent TaskID

	// Type is the taskType that was passed to runtime/trace.NewTask.
	//
	// May be "" if a task's TaskBegin event isn't present in the trace.
	Type string
}

// Region provides details about a Region event.
type Region struct {
	// Task is the ID of the task this region is associated with.
	Task TaskID

	// Type is the regionType that was passed to runtime/trace.StartRegion or runtime/trace.WithRegion.
	Type string
}

// Log provides details about a Log event.
type Log struct {
	// Task is the ID of the task this region is associated with.
	Task TaskID

	// Category is the category that was passed to runtime/trace.Log or runtime/trace.Logf.
	Category string

	// Message is the message that was passed to runtime/trace.Log or runtime/trace.Logf.
	Message string
}

// Stack represents a stack. It's really a handle to a stack and it's trivially comparable.
//
// If two Stacks are equal then their Frames are guaranteed to be identical. If they are not
// equal, however, their Frames may still be equal.
type Stack struct {
	table *evTable
	id    stackID
}

// Frames is an iterator over the frames in a Stack.
func (s Stack) Frames() iter.Seq[StackFrame]

// NoStack is a sentinel value that can be compared against any Stack value, indicating
// a lack of a stack trace.
var NoStack = Stack{}

// StackFrame represents a single frame of a stack.
type StackFrame struct {
	// PC is the program counter of the function call if this
	// is not a leaf frame. If it's a leaf frame, it's the point
	// at which the stack trace was taken.
	PC uint64

	// Func is the name of the function this frame maps to.
	Func string

	// File is the file which contains the source code of Func.
	File string

	// Line is the line number within File which maps to PC.
	Line uint64
}

// ExperimentalEvent presents a raw view of an experimental event's arguments and their names.
type ExperimentalEvent struct {
	// Name is the name of the event.
	Name string

	// Experiment is the name of the experiment this event is a part of.
	Experiment string

	// Args lists the names of the event's arguments in order.
	Args []string

	// argValues contains the raw integer arguments which are interpreted
	// by ArgValue using table.
	table     *evTable
	argValues []uint64
}

// ArgValue returns a typed Value for the i'th argument in the experimental event.
func (e ExperimentalEvent) ArgValue(i int) Value

// ExperimentalBatch represents a packet of unparsed data along with metadata about that packet.
type ExperimentalBatch struct {
	// Thread is the ID of the thread that produced a packet of data.
	Thread ThreadID

	// Data is a packet of unparsed data all produced by one thread.
	Data []byte
}

// Event represents a single event in the trace.
type Event struct {
	table *evTable
	ctx   schedCtx
	base  baseEvent
}

// Kind returns the kind of event that this is.
func (e Event) Kind() EventKind

// Time returns the timestamp of the event.
func (e Event) Time() Time

// Goroutine returns the ID of the goroutine that was executing when
// this event happened. It describes part of the execution context
// for this event.
//
// Note that for goroutine state transitions this always refers to the
// state before the transition. For example, if a goroutine is just
// starting to run on this thread and/or proc, then this will return
// NoGoroutine. In this case, the goroutine starting to run will be
// can be found at Event.StateTransition().Resource.
func (e Event) Goroutine() GoID

// Proc returns the ID of the proc this event event pertains to.
//
// Note that for proc state transitions this always refers to the
// state before the transition. For example, if a proc is just
// starting to run on this thread, then this will return NoProc.
func (e Event) Proc() ProcID

// Thread returns the ID of the thread this event pertains to.
//
// Note that for thread state transitions this always refers to the
// state before the transition. For example, if a thread is just
// starting to run, then this will return NoThread.
//
// Note: tracking thread state is not currently supported, so this
// will always return a valid thread ID. However thread state transitions
// may be tracked in the future, and callers must be robust to this
// possibility.
func (e Event) Thread() ThreadID

// Stack returns a handle to a stack associated with the event.
//
// This represents a stack trace at the current moment in time for
// the current execution context.
func (e Event) Stack() Stack

// Metric returns details about a Metric event.
//
// Panics if Kind != EventMetric.
func (e Event) Metric() Metric

// Label returns details about a Label event.
//
// Panics if Kind != EventLabel.
func (e Event) Label() Label

// Range returns details about an EventRangeBegin, EventRangeActive, or EventRangeEnd event.
//
// Panics if Kind != EventRangeBegin, Kind != EventRangeActive, and Kind != EventRangeEnd.
func (e Event) Range() Range

// RangeAttributes returns attributes for a completed range.
//
// Panics if Kind != EventRangeEnd.
func (e Event) RangeAttributes() []RangeAttribute

// Task returns details about a TaskBegin or TaskEnd event.
//
// Panics if Kind != EventTaskBegin and Kind != EventTaskEnd.
func (e Event) Task() Task

// Region returns details about a RegionBegin or RegionEnd event.
//
// Panics if Kind != EventRegionBegin and Kind != EventRegionEnd.
func (e Event) Region() Region

// Log returns details about a Log event.
//
// Panics if Kind != EventLog.
func (e Event) Log() Log

// StateTransition returns details about a StateTransition event.
//
// Panics if Kind != EventStateTransition.
func (e Event) StateTransition() StateTransition

// Sync returns details that are relevant for the following events, up to but excluding the
// next EventSync event.
func (e Event) Sync() Sync

// Sync contains details potentially relevant to all the following events, up to but excluding
// the next EventSync event.
type Sync struct {
	// N indicates that this is the Nth sync event in the trace.
	N int

	// ClockSnapshot is a snapshot of different clocks taken in close in time
	// that can be used to correlate trace events with data captured by other
	// tools. May be nil for older trace versions.
	ClockSnapshot *ClockSnapshot

	// ExperimentalBatches contain all the unparsed batches of data for a given experiment.
	ExperimentalBatches map[string][]ExperimentalBatch
}

// ClockSnapshot represents a near-simultaneous clock reading of several
// different system clocks. The snapshot can be used as a reference to convert
// timestamps to different clocks, which is helpful for correlating timestamps
// with data captured by other tools.
type ClockSnapshot struct {
	// Trace is a snapshot of the trace clock.
	Trace Time

	// Wall is a snapshot of the system's wall clock.
	Wall time.Time

	// Mono is a snapshot of the system's monotonic clock.
	Mono uint64
}

// Experimental returns a view of the raw event for an experimental event.
//
// Panics if Kind != EventExperimental.
func (e Event) Experimental() ExperimentalEvent

// String returns the event as a human-readable string.
//
// The format of the string is intended for debugging and is subject to change.
func (e Event) String() string
