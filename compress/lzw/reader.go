// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lzw implements the Lempel-Ziv-Welch compressed data format,
// described in T. A. Welch, “A Technique for High-Performance Data
// Compression”, Computer, 17(6) (June 1984), pp 8-19.
//
// In particular, it implements LZW as used by the GIF and PDF file
// formats, which means variable-width codes up to 12 bits and the first
// two non-literal codes are a clear code and an EOF code.
//
// The TIFF file format uses a similar but incompatible version of the LZW
// algorithm. See the golang.org/x/image/tiff/lzw package for an
// implementation.
package lzw

import (
	"github.com/shogo82148/std/io"
)

// Order specifies the bit ordering in an LZW data stream.
type Order int

const (
	// LSB means Least Significant Bits first, as used in the GIF file format.
	LSB Order = iota
	// MSB means Most Significant Bits first, as used in the TIFF and PDF
	// file formats.
	MSB
)

// Reader is an io.Reader which can be used to read compressed data in the
// LZW format.
type Reader struct {
	r        io.ByteReader
	bits     uint32
	nBits    uint
	width    uint
	read     func(*Reader) (uint16, error)
	litWidth int
	err      error

	// The first 1<<litWidth codes are literal codes.
	// The next two codes mean clear and EOF.
	// Other valid codes are in the range [lo, hi] where lo := clear + 2,
	// with the upper bound incrementing on each code seen.
	//
	// overflow is the code at which hi overflows the code width. It always
	// equals 1 << width.
	//
	// last is the most recently seen code, or decoderInvalidCode.
	//
	// An invariant is that hi < overflow.
	clear, eof, hi, overflow, last uint16

	// Each code c in [lo, hi] expands to two or more bytes. For c != hi:
	//   suffix[c] is the last of these bytes.
	//   prefix[c] is the code for all but the last byte.
	//   This code can either be a literal code or another code in [lo, c).
	// The c == hi case is a special case.
	suffix [1 << maxWidth]uint8
	prefix [1 << maxWidth]uint16

	// output is the temporary output buffer.
	// Literal codes are accumulated from the start of the buffer.
	// Non-literal codes decode to a sequence of suffixes that are first
	// written right-to-left from the end of the buffer before being copied
	// to the start of the buffer.
	// It is flushed when it contains >= 1<<maxWidth bytes,
	// so that there is always room to decode an entire code.
	output [2 * 1 << maxWidth]byte
	o      int
	toRead []byte
}

// Read implements io.Reader, reading uncompressed bytes from its underlying Reader.
func (r *Reader) Read(b []byte) (int, error)

// Close closes the Reader and returns an error for any future read operation.
// It does not close the underlying io.Reader.
func (r *Reader) Close() error

// Reset clears the Reader's state and allows it to be reused again
// as a new Reader.
func (r *Reader) Reset(src io.Reader, order Order, litWidth int)

// NewReader creates a new io.ReadCloser.
// Reads from the returned io.ReadCloser read and decompress data from r.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser when
// finished reading.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8. It must equal the litWidth
// used during compression.
//
// It is guaranteed that the underlying type of the returned io.ReadCloser
// is a *Reader.
func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
