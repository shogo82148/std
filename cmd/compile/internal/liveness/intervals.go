// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package liveness

// Interval hols the range [st,en).
type Interval struct {
	st, en int
}

// Intervals is a sequence of sorted, disjoint intervals.
type Intervals []Interval

func (i Interval) String() string

// Overlaps returns true if here is any overlap between i and i2.
func (i Interval) Overlaps(i2 Interval) bool

// MergeInto merges interval i2 into i1. This version happens to
// require that the two intervals either overlap or are adjacent.
func (i1 *Interval) MergeInto(i2 Interval) error

// IntervalsBuilder is a helper for constructing intervals based on
// live dataflow sets for a series of BBs where we're making a
// backwards pass over each BB looking for uses and kills. The
// expected use case is:
//
//   - invoke MakeIntervalsBuilder to create a new object "b"
//   - series of calls to b.Live/b.Kill based on a backwards reverse layout
//     order scan over instructions
//   - invoke b.Finish() to produce final set
//
// See the Live method comment for an IR example.
type IntervalsBuilder struct {
	s Intervals
	// index of last instruction visited plus 1
	lidx int
}

func (c *IntervalsBuilder) Finish() (Intervals, error)

// Live method should be invoked on instruction at position p if instr
// contains an upwards-exposed use of a resource. See the example in
// the comment at the beginning of this file for an example.
func (c *IntervalsBuilder) Live(pos int) error

// Kill method should be invoked on instruction at position p if instr
// should be treated as having a kill (lifetime end) for the
// resource. See the example in the comment at the beginning of this
// file for an example. Note that if we see a kill at position K for a
// resource currently live since J, this will result in a lifetime
// segment of [K+1,J+1), the assumption being that the first live
// instruction will be the one after the kill position, not the kill
// position itself.
func (c *IntervalsBuilder) Kill(pos int) error

func (is *Intervals) String() string

// Overlaps returns whether any of the component ranges in is overlaps
// with some range in is2.
func (is Intervals) Overlaps(is2 Intervals) bool

// Merge combines the intervals from "is" and "is2" and returns
// a new Intervals object containing all combined ranges from the
// two inputs.
func (is Intervals) Merge(is2 Intervals) Intervals
