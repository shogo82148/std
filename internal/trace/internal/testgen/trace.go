// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testgen

import (
	"github.com/shogo82148/std/regexp"
	"github.com/shogo82148/std/time"

	"github.com/shogo82148/std/internal/trace"
	"github.com/shogo82148/std/internal/trace/raw"
	"github.com/shogo82148/std/internal/trace/tracev2"
	"github.com/shogo82148/std/internal/trace/version"
)

func Main(ver version.Version, f func(*Trace))

// Trace represents an execution trace for testing.
//
// It does a little bit of work to ensure that the produced trace is valid,
// just for convenience. It mainly tracks batches and batch sizes (so they're
// trivially correct), tracks strings and stacks, and makes sure emitted string
// and stack batches are valid. That last part can be controlled by a few options.
//
// Otherwise, it performs no validation on the trace at all.
type Trace struct {
	// Trace data state.
	ver             version.Version
	names           map[string]tracev2.EventType
	specs           []tracev2.EventSpec
	events          []raw.Event
	gens            []*Generation
	validTimestamps bool
	lastTs          Time

	// Expectation state.
	bad      bool
	badMatch *regexp.Regexp
}

// NewTrace creates a new trace.
func NewTrace(ver version.Version) *Trace

// ExpectFailure writes down that the trace should be broken. The caller
// must provide a pattern matching the expected error produced by the parser.
func (t *Trace) ExpectFailure(pattern string)

// ExpectSuccess writes down that the trace should successfully parse.
func (t *Trace) ExpectSuccess()

// RawEvent emits an event into the trace. name must correspond to one
// of the names in Specs() result for the version that was passed to
// this trace.
func (t *Trace) RawEvent(typ tracev2.EventType, data []byte, args ...uint64)

// DisableTimestamps makes the timestamps for all events generated after
// this call zero. Raw events are exempted from this because the caller
// has to pass their own timestamp into those events anyway.
func (t *Trace) DisableTimestamps()

// Generation creates a new trace generation.
//
// This provides more structure than Event to allow for more easily
// creating complex traces that are mostly or completely correct.
func (t *Trace) Generation(gen uint64) *Generation

// Generate creates a test file for the trace.
func (t *Trace) Generate() []byte

var (
	NoString = ""
	NoStack  = []trace.StackFrame{}
)

// Generation represents a single generation in the trace.
type Generation struct {
	trace   *Trace
	gen     uint64
	batches []*Batch
	strings map[string]uint64
	stacks  map[stack]uint64
	sync    sync

	// Options applied when Trace.Generate is called.
	ignoreStringBatchSizeLimit bool
	ignoreStackBatchSizeLimit  bool
}

// Batch starts a new event batch in the trace data.
//
// This is convenience function for generating correct batches.
func (g *Generation) Batch(thread trace.ThreadID, time Time) *Batch

// String registers a string with the trace.
//
// This is a convenience function for easily adding correct
// strings to traces.
func (g *Generation) String(s string) uint64

// Stack registers a stack with the trace.
//
// This is a convenience function for easily adding correct
// stacks to traces.
func (g *Generation) Stack(stk []trace.StackFrame) uint64

// Sync configures the sync batch for the generation. For go1.25 and later,
// the time value is the timestamp of the EvClockSnapshot event. For earlier
// version, the time value is the timestamp of the batch containing a lone
// EvFrequency event.
func (g *Generation) Sync(freq uint64, time Time, mono uint64, wall time.Time)

// Batch represents an event batch.
type Batch struct {
	gen       *Generation
	thread    trace.ThreadID
	timestamp Time
	size      uint64
	events    []raw.Event
}

// Event emits an event into a batch. name must correspond to one
// of the names in Specs() result for the version that was passed to
// this trace. Callers must omit the timestamp delta.
func (b *Batch) Event(name string, args ...any)

// RawEvent emits an event into a batch. name must correspond to one
// of the names in Specs() result for the version that was passed to
// this trace.
func (b *Batch) RawEvent(typ tracev2.EventType, data []byte, args ...uint64)

// Seq represents a sequence counter.
type Seq uint64

// Time represents a low-level trace timestamp (which does not necessarily
// correspond to nanoseconds, like trace.Time does).
type Time uint64
