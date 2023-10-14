// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import "github.com/shogo82148/std/bytes"

// Writer is a test trace writer.
type Writer struct {
	bytes.Buffer
}

func NewWriter() *Writer

// Emit writes an event record to the trace.
// See Event types for valid types and required arguments.
func (w *Writer) Emit(typ byte, args ...uint64)
