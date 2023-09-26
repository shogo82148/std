// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/context"
	_ "github.com/shogo82148/std/unsafe"
)

// NewTask creates a task instance with the type taskType and returns
// it along with a Context that carries the task.
// If the input context contains a task, the new task is its subtask.
//
// The taskType is used to classify task instances. Analysis tools
// like the Go execution tracer may assume there are only a bounded
// number of unique task types in the system.
//
// The returned end function is used to mark the task's end.
// The trace tool measures task latency as the time between task creation
// and when the end function is called, and provides the latency
// distribution per task type.
// If the end function is called multiple times, only the first
// call is used in the latency measurement.
//
//	ctx, task := trace.NewTask(ctx, "awesomeTask")
//	trace.WithRegion(ctx, "preparation", prepWork)
//	// preparation of the task
//	go func() {  // continue processing the task in a separate goroutine.
//	    defer task.End()
//	    trace.WithRegion(ctx, "remainingWork", remainingWork)
//	}()
func NewTask(pctx context.Context, taskType string) (ctx context.Context, task *Task)

// Task is a data type for tracing a user-defined, logical operation.
type Task struct {
	id uint64
}

// End marks the end of the operation represented by the Task.
func (t *Task) End()

// Log emits a one-off event with the given category and message.
// Category can be empty and the API assumes there are only a handful of
// unique categories in the system.
func Log(ctx context.Context, category, message string)

// Logf is like Log, but the value is formatted using the specified format spec.
func Logf(ctx context.Context, category, format string, args ...any)

// WithRegion starts a region associated with its calling goroutine, runs fn,
// and then ends the region. If the context carries a task, the region is
// associated with the task. Otherwise, the region is attached to the background
// task.
//
// The regionType is used to classify regions, so there should be only a
// handful of unique region types.
func WithRegion(ctx context.Context, regionType string, fn func())

// StartRegion starts a region and returns a function for marking the
// end of the region. The returned Region's End function must be called
// from the same goroutine where the region was started.
// Within each goroutine, regions must nest. That is, regions started
// after this region must be ended before this region can be ended.
// Recommended usage is
//
//	defer trace.StartRegion(ctx, "myTracedRegion").End()
func StartRegion(ctx context.Context, regionType string) *Region

// Region is a region of code whose execution time interval is traced.
type Region struct {
	id         uint64
	regionType string
}

// End marks the end of the traced code region.
func (r *Region) End()

// IsEnabled reports whether tracing is enabled.
// The information is advisory only. The tracing status
// may have changed by the time this function returns.
func IsEnabled() bool
