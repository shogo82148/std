// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/time"
)

// Summary is the analysis result produced by the summarizer.
type Summary struct {
	Goroutines map[tracev2.GoID]*GoroutineSummary
	Tasks      map[tracev2.TaskID]*UserTaskSummary
}

// GoroutineSummary contains statistics and execution details of a single goroutine.
// (For v2 traces.)
type GoroutineSummary struct {
	ID           tracev2.GoID
	Name         string
	PC           uint64
	CreationTime tracev2.Time
	StartTime    tracev2.Time
	EndTime      tracev2.Time

	// List of regions in the goroutine, sorted based on the start time.
	Regions []*UserRegionSummary

	// Statistics of execution time during the goroutine execution.
	GoroutineExecStats

	// goroutineSummary is state used just for computing this structure.
	// It's dropped before being returned to the caller.
	//
	// More specifically, if it's nil, it indicates that this summary has
	// already been finalized.
	*goroutineSummary
}

// UserTaskSummary represents a task in the trace.
type UserTaskSummary struct {
	ID       tracev2.TaskID
	Name     string
	Parent   *UserTaskSummary
	Children []*UserTaskSummary

	// Task begin event. An EventTaskBegin event or nil.
	Start *tracev2.Event

	// End end event. Normally EventTaskEnd event or nil.
	End *tracev2.Event

	// Logs is a list of tracev2.EventLog events associated with the task.
	Logs []*tracev2.Event

	// List of regions in the task, sorted based on the start time.
	Regions []*UserRegionSummary

	// Goroutines is the set of goroutines associated with this task.
	Goroutines map[tracev2.GoID]*GoroutineSummary
}

// Complete returns true if we have complete information about the task
// from the trace: both a start and an end.
func (s *UserTaskSummary) Complete() bool

// Descendents returns a slice consisting of itself (always the first task returned),
// and the transitive closure of all of its children.
func (s *UserTaskSummary) Descendents() []*UserTaskSummary

// UserRegionSummary represents a region and goroutine execution stats
// while the region was active. (For v2 traces.)
type UserRegionSummary struct {
	TaskID tracev2.TaskID
	Name   string

	// Region start event. Normally EventRegionBegin event or nil,
	// but can be a state transition event from NotExist or Undetermined
	// if the region is a synthetic region representing task inheritance
	// from the parent goroutine.
	Start *tracev2.Event

	// Region end event. Normally EventRegionEnd event or nil,
	// but can be a state transition event to NotExist if the goroutine
	// terminated without explicitly ending the region.
	End *tracev2.Event

	GoroutineExecStats
}

// GoroutineExecStats contains statistics about a goroutine's execution
// during a period of time.
type GoroutineExecStats struct {
	// These stats are all non-overlapping.
	ExecTime          time.Duration
	SchedWaitTime     time.Duration
	BlockTimeByReason map[string]time.Duration
	SyscallTime       time.Duration
	SyscallBlockTime  time.Duration

	// TotalTime is the duration of the goroutine's presence in the trace.
	// Necessarily overlaps with other stats.
	TotalTime time.Duration

	// Total time the goroutine spent in certain ranges; may overlap
	// with other stats.
	RangeTime map[string]time.Duration
}

func (s GoroutineExecStats) NonOverlappingStats() map[string]time.Duration

// UnknownTime returns whatever isn't accounted for in TotalTime.
func (s GoroutineExecStats) UnknownTime() time.Duration

// Summarizer constructs per-goroutine time statistics for v2 traces.
type Summarizer struct {
	// gs contains the map of goroutine summaries we're building up to return to the caller.
	gs map[tracev2.GoID]*GoroutineSummary

	// tasks contains the map of task summaries we're building up to return to the caller.
	tasks map[tracev2.TaskID]*UserTaskSummary

	// syscallingP and syscallingG represent a binding between a P and G in a syscall.
	// Used to correctly identify and clean up after syscalls (blocking or otherwise).
	syscallingP map[tracev2.ProcID]tracev2.GoID
	syscallingG map[tracev2.GoID]tracev2.ProcID

	// rangesP is used for optimistic tracking of P-based ranges for goroutines.
	//
	// It's a best-effort mapping of an active range on a P to the goroutine we think
	// is associated with it.
	rangesP map[rangeP]tracev2.GoID

	lastTs tracev2.Time
	syncTs tracev2.Time
}

// NewSummarizer creates a new struct to build goroutine stats from a trace.
func NewSummarizer() *Summarizer

// Event feeds a single event into the stats summarizer.
func (s *Summarizer) Event(ev *tracev2.Event)

// Finalize indicates to the summarizer that we're done processing the trace.
// It cleans up any remaining state and returns the full summary.
func (s *Summarizer) Finalize() *Summary

// RelatedGoroutinesV2 finds a set of goroutines related to goroutine goid for v2 traces.
// The association is based on whether they have synchronized with each other in the Go
// scheduler (one has unblocked another).
func RelatedGoroutinesV2(events []tracev2.Event, goid tracev2.GoID) map[tracev2.GoID]struct{}
