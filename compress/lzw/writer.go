// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lzw

import (
	"github.com/shogo82148/std/io"
)

// Writer is an LZW compressor. It writes the compressed form of the data
// to an underlying writer (see [NewWriter]).
type Writer struct {
	// w is the writer that compressed bytes are written to.
	w writer
	// litWidth is the width in bits of literal codes.
	litWidth uint
	// order, write, bits, nBits and width are the state for
	// converting a code stream into a byte stream.
	order Order
	write func(*Writer, uint32) error
	nBits uint
	width uint
	bits  uint32
	// hi is the code implied by the next code emission.
	// overflow is the code at which hi overflows the code width.
	hi, overflow uint32
	// savedCode is the accumulated code at the end of the most recent Write
	// call. It is equal to invalidCode if there was no such call.
	savedCode uint32
	// err is the first error encountered during writing. Closing the writer
	// will make any future Write calls return errClosed
	err error
	// table is the hash table from 20-bit keys to 12-bit values. Each table
	// entry contains key<<12|val and collisions resolve by linear probing.
	// The keys consist of a 12-bit code prefix and an 8-bit byte suffix.
	// The values are a 12-bit code.
	table [tableSize]uint32
}

// Write writes a compressed representation of p to w's underlying writer.
func (w *Writer) Write(p []byte) (n int, err error)

// Close closes the [Writer], flushing any pending output. It does not close
// w's underlying writer.
func (w *Writer) Close() error

// Reset clears the [Writer]'s state and allows it to be reused again
// as a new [Writer].
func (w *Writer) Reset(dst io.Writer, order Order, litWidth int)

// NewWriter creates a new [io.WriteCloser].
// Writes to the returned [io.WriteCloser] are compressed and written to w.
// It is the caller's responsibility to call Close on the WriteCloser when
// finished writing.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8. Input bytes must be less than 1<<litWidth.
//
// It is guaranteed that the underlying type of the returned [io.WriteCloser]
// is a *[Writer].
func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
