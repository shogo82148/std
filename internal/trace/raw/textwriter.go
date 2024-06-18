// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/internal/trace/version"
)

// TextWriter emits the text format of a trace.
type TextWriter struct {
	w io.Writer
	v version.Version
}

// NewTextWriter creates a new write for the trace text format.
func NewTextWriter(w io.Writer, v version.Version) (*TextWriter, error)

// WriteEvent writes a single event to the stream.
func (w *TextWriter) WriteEvent(e Event) error
