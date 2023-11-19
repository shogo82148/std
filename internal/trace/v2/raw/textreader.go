// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/internal/trace/v2/event"
	"github.com/shogo82148/std/internal/trace/v2/version"
)

// TextReader parses a text format trace with only very basic validation
// into an event stream.
type TextReader struct {
	v     version.Version
	specs []event.Spec
	names map[string]event.Type
	s     *bufio.Scanner
}

// NewTextReader creates a new reader for the trace text format.
func NewTextReader(r io.Reader) (*TextReader, error)

// Version returns the version of the trace that we're reading.
func (r *TextReader) Version() version.Version

// ReadEvent reads and returns the next trace event in the text stream.
func (r *TextReader) ReadEvent() (Event, error)
