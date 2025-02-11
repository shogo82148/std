// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/internal/trace/version"
)

// Reader reads a byte stream, validates it, and produces trace events.
//
// Provided the trace is non-empty the Reader always produces a Sync
// event as the first event, and a Sync event as the last event.
// (There may also be any number of Sync events in the middle, too.)
type Reader struct {
	version    version.Version
	r          *bufio.Reader
	lastTs     Time
	gen        *generation
	spill      *spilledBatch
	spillErr   error
	frontier   []*batchCursor
	cpuSamples []cpuSample
	order      ordering
	syncs      int
	done       bool

	v1Events *traceV1Converter
}

// NewReader creates a new trace reader.
func NewReader(r io.Reader) (*Reader, error)

// ReadEvent reads a single event from the stream.
//
// If the stream has been exhausted, it returns an invalid event and io.EOF.
func (r *Reader) ReadEvent() (e Event, err error)
