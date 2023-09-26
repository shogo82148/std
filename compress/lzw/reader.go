// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lzw implements the Lempel-Ziv-Welch compressed data format,
// described in T. A. Welch, “A Technique for High-Performance Data
// Compression”, Computer, 17(6) (June 1984), pp 8-19.
//
// In particular, it implements LZW as used by the GIF, TIFF and PDF file
// formats, which means variable-width codes up to 12 bits and the first
// two non-literal codes are a clear code and an EOF code.
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

// decoder is the state from which the readXxx method converts a byte
// stream into a code stream.

// NewReader creates a new io.ReadCloser that satisfies reads by decompressing
// the data read from r.
// It is the caller's responsibility to call Close on the ReadCloser when
// finished reading.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8.
func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
