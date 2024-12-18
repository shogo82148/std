// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// A Writer writes records using CSV encoding.
//
// As returned by [NewWriter], a Writer writes records terminated by a
// newline and uses ',' as the field delimiter. The exported fields can be
// changed to customize the details before
// the first call to [Writer.Write] or [Writer.WriteAll].
//
// [Writer.Comma] is the field delimiter.
//
// If [Writer.UseCRLF] is true,
// the Writer ends each output line with \r\n instead of \n.
//
// The writes of individual records are buffered.
// After all data has been written, the client should call the
// [Writer.Flush] method to guarantee all data has been forwarded to
// the underlying [io.Writer].  Any errors that occurred should
// be checked by calling the [Writer.Error] method.
type Writer struct {
	Comma   rune
	UseCRLF bool
	w       *bufio.Writer
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer

// Write writes a single CSV record to w along with any necessary quoting.
// A record is a slice of strings with each string being one field.
// Writes are buffered, so [Writer.Flush] must eventually be called to ensure
// that the record is written to the underlying [io.Writer].
func (w *Writer) Write(record []string) error

// Flush writes any buffered data to the underlying [io.Writer].
// To check if an error occurred during Flush, call [Writer.Error].
func (w *Writer) Flush()

// Error reports any error that has occurred during
// a previous [Writer.Write] or [Writer.Flush].
func (w *Writer) Error() error

// WriteAll writes multiple CSV records to w using [Writer.Write] and
// then calls [Writer.Flush], returning any error from the Flush.
func (w *Writer) WriteAll(records [][]string) error
