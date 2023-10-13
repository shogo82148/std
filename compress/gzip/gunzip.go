// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gzip implements reading and writing of gzip format compressed files,
// as specified in RFC 1952.
package gzip

import (
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

var (
	// ErrChecksum is returned when reading GZIP data that has an invalid checksum.
	ErrChecksum = errors.New("gzip: invalid checksum")
	// ErrHeader is returned when reading GZIP data that has an invalid header.
	ErrHeader = errors.New("gzip: invalid header")
)

// The gzip file stores a header giving metadata about the compressed file.
// That header is exposed as the fields of the [Writer] and [Reader] structs.
//
// Strings must be UTF-8 encoded and may only contain Unicode code points
// U+0001 through U+00FF, due to limitations of the GZIP file format.
type Header struct {
	Comment string
	Extra   []byte
	ModTime time.Time
	Name    string
	OS      byte
}

// A Reader is an [io.Reader] that can be read to retrieve
// uncompressed data from a gzip-format compressed file.
//
// In general, a gzip file can be a concatenation of gzip files,
// each with its own header. Reads from the Reader
// return the concatenation of the uncompressed data of each.
// Only the first header is recorded in the Reader fields.
//
// Gzip files store a length and checksum of the uncompressed data.
// The Reader will return an [ErrChecksum] when [Reader.Read]
// reaches the end of the uncompressed data if it does not
// have the expected length or checksum. Clients should treat data
// returned by [Reader.Read] as tentative until they receive the [io.EOF]
// marking the end of the data.
type Reader struct {
	Header
	r            flate.Reader
	decompressor io.ReadCloser
	digest       uint32
	size         uint32
	buf          [512]byte
	err          error
	multistream  bool
}

// NewReader creates a new [Reader] reading the given reader.
// If r does not also implement [io.ByteReader],
// the decompressor may read more data than necessary from r.
//
// It is the caller's responsibility to call Close on the [Reader] when done.
//
// The Reader.Header fields will be valid in the [Reader] returned.
func NewReader(r io.Reader) (*Reader, error)

// Reset discards the [Reader] z's state and makes it equivalent to the
// result of its original state from [NewReader], but reading from r instead.
// This permits reusing a [Reader] rather than allocating a new one.
func (z *Reader) Reset(r io.Reader) error

// Multistream controls whether the reader supports multistream files.
//
// If enabled (the default), the [Reader] expects the input to be a sequence
// of individually gzipped data streams, each with its own header and
// trailer, ending at EOF. The effect is that the concatenation of a sequence
// of gzipped files is treated as equivalent to the gzip of the concatenation
// of the sequence. This is standard behavior for gzip readers.
//
// Calling Multistream(false) disables this behavior; disabling the behavior
// can be useful when reading file formats that distinguish individual gzip
// data streams or mix gzip data streams with other data streams.
// In this mode, when the [Reader] reaches the end of the data stream,
// [Reader.Read] returns [io.EOF]. The underlying reader must implement [io.ByteReader]
// in order to be left positioned just after the gzip stream.
// To start the next stream, call z.Reset(r) followed by z.Multistream(false).
// If there is no next stream, z.Reset(r) will return [io.EOF].
func (z *Reader) Multistream(ok bool)

// Read implements [io.Reader], reading uncompressed bytes from its underlying [Reader].
func (z *Reader) Read(p []byte) (n int, err error)

// Close closes the [Reader]. It does not close the underlying [io.Reader].
// In order for the GZIP checksum to be verified, the reader must be
// fully consumed until the [io.EOF].
func (z *Reader) Close() error
