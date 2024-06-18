// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/internal/trace/event"
	"github.com/shogo82148/std/internal/trace/version"
)

// Reader parses trace bytes with only very basic validation
// into an event stream.
type Reader struct {
	r     *bufio.Reader
	v     version.Version
	specs []event.Spec
}

// NewReader creates a new reader for the trace wire format.
func NewReader(r io.Reader) (*Reader, error)

// Version returns the version of the trace that we're reading.
func (r *Reader) Version() version.Version

// ReadEvent reads and returns the next trace event in the byte stream.
func (r *Reader) ReadEvent() (Event, error)
