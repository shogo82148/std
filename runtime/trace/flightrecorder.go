// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

// FlightRecorder represents a single consumer of a Go execution
// trace.
// It tracks a moving window over the execution trace produced by
// the runtime, always containing the most recent trace data.
//
// At most one flight recorder may be active at any given time,
// though flight recording is allowed to be concurrently active
// with a trace consumer using trace.Start.
// This restriction of only a single flight recorder may be removed
// in the future.
type FlightRecorder struct {
	err error

	// State specific to the recorder.
	header [16]byte
	active rawGeneration
	ringMu sync.Mutex
	ring   []rawGeneration
	freq   frequency

	// Externally-set options.
	targetSize   uint64
	targetPeriod time.Duration

	enabled bool
	writing sync.Mutex

	// The values of targetSize and targetPeriod we've committed to since the last Start.
	wantSize uint64
	wantDur  time.Duration
}

// NewFlightRecorder creates a new flight recorder from the provided configuration.
func NewFlightRecorder(cfg FlightRecorderConfig) *FlightRecorder

// Start activates the flight recorder and begins recording trace data.
// Only one call to trace.Start may be active at any given time.
// In addition, currently only one flight recorder may be active in the program.
// Returns an error if the flight recorder cannot be started or is already started.
func (fr *FlightRecorder) Start() error

// Stop ends recording of trace data. It blocks until any concurrent WriteTo calls complete.
func (fr *FlightRecorder) Stop()

// Enabled returns true if the flight recorder is active.
// Specifically, it will return true if Start did not return an error, and Stop has not yet been called.
// It is safe to call from multiple goroutines simultaneously.
func (fr *FlightRecorder) Enabled() bool

// WriteTo snapshots the moving window tracked by the flight recorder.
// The snapshot is expected to contain data that is up-to-date as of when WriteTo is called,
// though this is not a hard guarantee.
// Only one goroutine may execute WriteTo at a time.
// An error is returned upon failure to write to w, if another WriteTo call is already in-progress,
// or if the flight recorder is inactive.
func (fr *FlightRecorder) WriteTo(w io.Writer) (n int64, err error)

type FlightRecorderConfig struct {
	// MinAge is a lower bound on the age of an event in the flight recorder's window.
	//
	// The flight recorder will strive to promptly discard events older than the minimum age,
	// but older events may appear in the window snapshot. The age setting will always be
	// overridden by MaxSize.
	//
	// If this is 0, the minimum age is implementation defined, but can be assumed to be on the order
	// of seconds.
	MinAge time.Duration

	// MaxBytes is an upper bound on the size of the window in bytes.
	//
	// This setting takes precedence over MinAge.
	// However, it does not make any guarantees on the size of the data WriteTo will write,
	// nor does it guarantee memory overheads will always stay below MaxBytes. Treat it
	// as a hint.
	//
	// If this is 0, the maximum size is implementation defined.
	MaxBytes uint64
}
