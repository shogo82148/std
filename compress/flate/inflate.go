// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package flate implements the DEFLATE compressed data format, described in
// RFC 1951.  The gzip and zlib packages implement access to DEFLATE-based file
// formats.
package flate

import (
	"github.com/shogo82148/std/io"
)

// A CorruptInputError reports the presence of corrupt input at a given offset.
type CorruptInputError int64

func (e CorruptInputError) Error() string

// An InternalError reports an error in the flate code itself.
type InternalError string

func (e InternalError) Error() string

// A ReadError reports an error encountered while reading input.
type ReadError struct {
	Offset int64
	Err    error
}

func (e *ReadError) Error() string

// A WriteError reports an error encountered while writing output.
type WriteError struct {
	Offset int64
	Err    error
}

func (e *WriteError) Error() string

// Huffman decoder is based on
// J. Brian Connell, ``A Huffman-Shannon-Fano Code,''
// Proceedings of the IEEE, 61(7) (July 1973), pp 1046-1047.

// Hard-coded Huffman tables for DEFLATE algorithm.
// See RFC 1951, section 3.2.6.

// The actual read interface needed by NewReader.
// If the passed in io.Reader does not also have ReadByte,
// the NewReader will introduce its own buffering.
type Reader interface {
	io.Reader
	ReadByte() (c byte, err error)
}

// Decompress state.

// NewReader returns a new ReadCloser that can be used
// to read the uncompressed version of r.  It is the caller's
// responsibility to call Close on the ReadCloser when
// finished reading.
func NewReader(r io.Reader) io.ReadCloser

// NewReaderDict is like NewReader but initializes the reader
// with a preset dictionary.  The returned Reader behaves as if
// the uncompressed data stream started with the given dictionary,
// which has already been read.  NewReaderDict is typically used
// to read data compressed by NewWriterDict.
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
