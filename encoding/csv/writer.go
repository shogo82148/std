// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// A Writer writes records to a CSV encoded file.
//
// As returned by NewWriter, a Writer writes records terminated by a
// newline and uses ',' as the field delimiter.  The exported fields can be
// changed to customize the details before the first call to Write or WriteAll.
//
// Comma is the field delimiter.
//
// If UseCRLF is true, the Writer ends each record with \r\n instead of \n.
type Writer struct {
	Comma   rune
	UseCRLF bool
	w       *bufio.Writer
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer

// Writer writes a single CSV record to w along with any necessary quoting.
// A record is a slice of strings with each string being one field.
func (w *Writer) Write(record []string) (err error)

// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush()

// WriteAll writes multiple CSV records to w using Write and then calls Flush.
func (w *Writer) WriteAll(records [][]string) (err error)
