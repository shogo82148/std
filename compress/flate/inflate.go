// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package flate implements the DEFLATE compressed data format, described in
// RFC 1951.  The gzip and zlib packages implement access to DEFLATE-based file
// formats.
package flate

import (
	"github.com/shogo82148/std/io"
	mathbits "math/bits"
)

// Initialize the fixedHuffmanDecoder only once upon first use.

// A CorruptInputError reports the presence of corrupt input at a given offset.
type CorruptInputError int64

func (e CorruptInputError) Error() string

// An InternalError reports an error in the flate code itself.
type InternalError string

func (e InternalError) Error() string

// A ReadError reports an error encountered while reading input.
//
// Deprecated: No longer returned.
type ReadError struct {
	Offset int64
	Err    error
}

func (e *ReadError) Error() string

// A WriteError reports an error encountered while writing output.
//
// Deprecated: No longer returned.
type WriteError struct {
	Offset int64
	Err    error
}

func (e *WriteError) Error() string

// Resetter resets a ReadCloser returned by NewReader or NewReaderDict to
// to switch to a new underlying Reader. This permits reusing a ReadCloser
// instead of allocating a new one.
type Resetter interface {
	Reset(r io.Reader, dict []byte) error
}

// The actual read interface needed by NewReader.
// If the passed in io.Reader does not also have ReadByte,
// the NewReader will introduce its own buffering.
type Reader interface {
	io.Reader
	io.ByteReader
}

// Decompress state.

// NewReader returns a new ReadCloser that can be used
// to read the uncompressed version of r.
// If r does not also implement io.ByteReader,
// the decompressor may read more data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser
// when finished reading.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReader(r io.Reader) io.ReadCloser

// NewReaderDict is like NewReader but initializes the reader
// with a preset dictionary. The returned Reader behaves as if
// the uncompressed data stream started with the given dictionary,
// which has already been read. NewReaderDict is typically used
// to read data compressed by NewWriterDict.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
