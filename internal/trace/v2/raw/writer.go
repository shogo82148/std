// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/internal/trace/v2/event"
	"github.com/shogo82148/std/internal/trace/v2/version"
)

// Writer emits the wire format of a trace.
//
// It may not produce a byte-for-byte compatible trace from what is
// produced by the runtime, because it may be missing extra padding
// in the LEB128 encoding that the runtime adds but isn't necessary
// when you know the data up-front.
type Writer struct {
	w     io.Writer
	buf   []byte
	v     version.Version
	specs []event.Spec
}

// NewWriter creates a new byte format writer.
func NewWriter(w io.Writer, v version.Version) (*Writer, error)

// WriteEvent writes a single event to the trace wire format stream.
func (w *Writer) WriteEvent(e Event) error
