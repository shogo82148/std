// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzip

import (
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/io"
)

// These constants are copied from the flate package, so that code that imports
// "compress/gzip" does not also have to import "compress/flate".
const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)

// A Writer is an io.WriteCloser.
// Writes to a Writer are compressed and written to w.
type Writer struct {
	Header
	w           io.Writer
	level       int
	wroteHeader bool
	compressor  *flate.Writer
	digest      uint32
	size        uint32
	closed      bool
	buf         [10]byte
	err         error
}

// NewWriter returns a new [Writer].
// Writes to the returned writer are compressed and written to w.
//
// It is the caller's responsibility to call Close on the [Writer] when done.
// Writes may be buffered and not flushed until Close.
//
// Callers that wish to set the fields in Writer.Header must do so before
// the first call to Write, Flush, or Close.
func NewWriter(w io.Writer) *Writer

// NewWriterLevel is like [NewWriter] but specifies the compression level instead
// of assuming [DefaultCompression].
//
// The compression level can be [DefaultCompression], [NoCompression], [HuffmanOnly]
// or any integer value between [BestSpeed] and [BestCompression] inclusive.
// The error returned will be nil if the level is valid.
func NewWriterLevel(w io.Writer, level int) (*Writer, error)

// Reset discards the [Writer] z's state and makes it equivalent to the
// result of its original state from [NewWriter] or [NewWriterLevel], but
// writing to w instead. This permits reusing a [Writer] rather than
// allocating a new one.
func (z *Writer) Reset(w io.Writer)

// Write writes a compressed form of p to the underlying [io.Writer]. The
// compressed bytes are not necessarily flushed until the [Writer] is closed.
func (z *Writer) Write(p []byte) (int, error)

// Flush flushes any pending compressed data to the underlying writer.
//
// It is useful mainly in compressed network protocols, to ensure that
// a remote reader has enough data to reconstruct a packet. Flush does
// not return until the data has been written. If the underlying
// writer returns an error, Flush returns that error.
//
// In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.
func (z *Writer) Flush() error

// Close closes the [Writer] by flushing any unwritten data to the underlying
// [io.Writer] and writing the GZIP footer.
// It does not close the underlying [io.Writer].
func (z *Writer) Close() error
