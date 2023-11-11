// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// Reader reads a byte stream, validates it, and produces trace events.
type Reader struct {
	r           *bufio.Reader
	lastTs      Time
	gen         *generation
	spill       *spilledBatch
	frontier    []*batchCursor
	cpuSamples  []cpuSample
	order       ordering
	emittedSync bool
}

// NewReader creates a new trace reader.
func NewReader(r io.Reader) (*Reader, error)

// ReadEvent reads a single event from the stream.
//
// If the stream has been exhausted, it returns an invalid
// event and io.EOF.
func (r *Reader) ReadEvent() (e Event, err error)
