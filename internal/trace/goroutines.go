// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

// GDesc contains statistics and execution details of a single goroutine.
type GDesc struct {
	ID           uint64
	Name         string
	PC           uint64
	CreationTime int64
	StartTime    int64
	EndTime      int64

	// List of regions in the goroutine, sorted based on the start time.
	Regions []*UserRegionDesc

	// Statistics of execution time during the goroutine execution.
	GExecutionStat

	*gdesc
}

// UserRegionDesc represents a region and goroutine execution stats
// while the region was active.
type UserRegionDesc struct {
	TaskID uint64
	Name   string

	// Region start event. Normally EvUserRegion start event or nil,
	// but can be EvGoCreate event if the region is a synthetic
	// region representing task inheritance from the parent goroutine.
	Start *Event

	// Region end event. Normally EvUserRegion end event or nil,
	// but can be EvGoStop or EvGoEnd event if the goroutine
	// terminated without explicitly ending the region.
	End *Event

	GExecutionStat
}

// GExecutionStat contains statistics about a goroutine's execution
// during a period of time.
type GExecutionStat struct {
	ExecTime      int64
	SchedWaitTime int64
	IOTime        int64
	BlockTime     int64
	SyscallTime   int64
	GCTime        int64
	SweepTime     int64
	TotalTime     int64
}

// GoroutineStats generates statistics for all goroutines in the trace.
func GoroutineStats(events []*Event) map[uint64]*GDesc

// RelatedGoroutines finds a set of goroutines related to goroutine goid.
func RelatedGoroutines(events []*Event, goid uint64) map[uint64]bool

func IsSystemGoroutine(entryFn string) bool
