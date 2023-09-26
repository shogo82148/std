// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lzw

import (
	"github.com/shogo82148/std/io"
)

// A writer is a buffered, flushable writer.

// Writer is an LZW compressor. It writes the compressed form of the data
// to an underlying writer (see NewWriter).
type Writer struct {
	w writer

	order Order
	write func(*Writer, uint32) error
	bits  uint32
	nBits uint
	width uint

	litWidth uint

	hi, overflow uint32

	savedCode uint32

	err error

	table [tableSize]uint32
}

// errOutOfCodes is an internal error that means that the writer has run out
// of unused codes and a clear code needs to be sent next.

// Write writes a compressed representation of p to w's underlying writer.
func (w *Writer) Write(p []byte) (n int, err error)

// Close closes the Writer, flushing any pending output. It does not close
// w's underlying writer.
func (w *Writer) Close() error

// Reset clears the Writer's state and allows it to be reused again
// as a new Writer.
func (w *Writer) Reset(dst io.Writer, order Order, litWidth int)

// NewWriter creates a new io.WriteCloser.
// Writes to the returned io.WriteCloser are compressed and written to w.
// It is the caller's responsibility to call Close on the WriteCloser when
// finished writing.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8. Input bytes must be less than 1<<litWidth.
//
// It is guaranteed that the underlying type of the returned io.WriteCloser
// is a *Writer.
func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
