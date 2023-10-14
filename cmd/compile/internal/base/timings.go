// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"github.com/shogo82148/std/io"
)

var Timer Timings

// Timings collects the execution times of labeled phases
// which are added through a sequence of Start/Stop calls.
// Events may be associated with each phase via AddEvent.
type Timings struct {
	list   []timestamp
	events map[int][]*event
}

// Start marks the beginning of a new phase and implicitly stops the previous phase.
// The phase name is the colon-separated concatenation of the labels.
func (t *Timings) Start(labels ...string)

// Stop marks the end of a phase and implicitly starts a new phase.
// The labels are added to the labels of the ended phase.
func (t *Timings) Stop(labels ...string)

// AddEvent associates an event, i.e., a count, or an amount of data,
// with the most recently started or stopped phase; or the very first
// phase if Start or Stop hasn't been called yet. The unit specifies
// the unit of measurement (e.g., MB, lines, no. of funcs, etc.).
func (t *Timings) AddEvent(size int64, unit string)

// Write prints the phase times to w.
// The prefix is printed at the start of each line.
func (t *Timings) Write(w io.Writer, prefix string)
