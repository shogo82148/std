// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/time"
)

// MutatorUtil is a change in mutator utilization at a particular
// time. Mutator utilization functions are represented as a
// time-ordered []MutatorUtil.
type MutatorUtil struct {
	Time int64
	// Util is the mean mutator utilization starting at Time. This
	// is in the range [0, 1].
	Util float64
}

// UtilFlags controls the behavior of MutatorUtilization.
type UtilFlags int

const (
	// UtilSTW means utilization should account for STW events.
	// This includes non-GC STW events, which are typically user-requested.
	UtilSTW UtilFlags = 1 << iota
	// UtilBackground means utilization should account for
	// background mark workers.
	UtilBackground
	// UtilAssist means utilization should account for mark
	// assists.
	UtilAssist
	// UtilSweep means utilization should account for sweeping.
	UtilSweep

	// UtilPerProc means each P should be given a separate
	// utilization function. Otherwise, there is a single function
	// and each P is given a fraction of the utilization.
	UtilPerProc
)

// MutatorUtilization returns a set of mutator utilization functions
// for the given trace. Each function will always end with 0
// utilization. The bounds of each function are implicit in the first
// and last event; outside of these bounds each function is undefined.
//
// If the UtilPerProc flag is not given, this always returns a single
// utilization function. Otherwise, it returns one function per P.
func MutatorUtilization(events []*Event, flags UtilFlags) [][]MutatorUtil

// An MMUCurve is the minimum mutator utilization curve across
// multiple window sizes.
type MMUCurve struct {
	series []mmuSeries
}

// NewMMUCurve returns an MMU curve for the given mutator utilization
// function.
func NewMMUCurve(utils [][]MutatorUtil) *MMUCurve

// UtilWindow is a specific window at Time.
type UtilWindow struct {
	Time int64
	// MutatorUtil is the mean mutator utilization in this window.
	MutatorUtil float64
}

// MMU returns the minimum mutator utilization for the given time
// window. This is the minimum utilization for all windows of this
// duration across the execution. The returned value is in the range
// [0, 1].
func (c *MMUCurve) MMU(window time.Duration) (mmu float64)

// Examples returns n specific examples of the lowest mutator
// utilization for the given window size. The returned windows will be
// disjoint (otherwise there would be a huge number of
// mostly-overlapping windows at the single lowest point). There are
// no guarantees on which set of disjoint windows this returns.
func (c *MMUCurve) Examples(window time.Duration, n int) (worst []UtilWindow)

// MUD returns mutator utilization distribution quantiles for the
// given window size.
//
// The mutator utilization distribution is the distribution of mean
// mutator utilization across all windows of the given window size in
// the trace.
//
// The minimum mutator utilization is the minimum (0th percentile) of
// this distribution. (However, if only the minimum is desired, it's
// more efficient to use the MMU method.)
func (c *MMUCurve) MUD(window time.Duration, quantiles []float64) []float64
