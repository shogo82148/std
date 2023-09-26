// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

import (
	"github.com/shogo82148/std/time"
)

// GCStats collect information about recent garbage collections.
type GCStats struct {
	LastGC         time.Time
	NumGC          int64
	PauseTotal     time.Duration
	Pause          []time.Duration
	PauseQuantiles []time.Duration
}

// ReadGCStats reads statistics about garbage collection into stats.
// The number of entries in the pause history is system-dependent;
// stats.Pause slice will be reused if large enough, reallocated otherwise.
// ReadGCStats may use the full capacity of the stats.Pause slice.
// If stats.PauseQuantiles is non-empty, ReadGCStats fills it with quantiles
// summarizing the distribution of pause time. For example, if
// len(stats.PauseQuantiles) is 5, it will be filled with the minimum,
// 25%, 50%, 75%, and maximum pause times.
func ReadGCStats(stats *GCStats)

// SetGCPercent sets the garbage collection target percentage:
// a collection is triggered when the ratio of freshly allocated data
// to live data remaining after the previous collection reaches this percentage.
// SetGCPercent returns the previous setting.
// The initial setting is the value of the GOGC environment variable
// at startup, or 100 if the variable is not set.
// A negative percentage disables garbage collection.
func SetGCPercent(percent int) int

// FreeOSMemory forces a garbage collection followed by an
// attempt to return as much memory to the operating system
// as possible. (Even if this is not called, the runtime gradually
// returns memory to the operating system in a background task.)
func FreeOSMemory()
