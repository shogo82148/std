// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/io"
)

// A Reader implements the [io.Reader], [io.ReaderAt], [io.WriterTo], [io.Seeker],
// [io.ByteScanner], and [io.RuneScanner] interfaces by reading from
// a byte slice.
// Unlike a [Buffer], a Reader is read-only and supports seeking.
// The zero value for Reader operates like a Reader of an empty slice.
type Reader struct {
	s        []byte
	i        int64
	prevRune int
}

// Len returns the number of bytes of the unread portion of the
// slice.
func (r *Reader) Len() int

// Size returns the original length of the underlying byte slice.
// Size is the number of bytes available for reading via [Reader.ReadAt].
// The result is unaffected by any method calls except [Reader.Reset].
func (r *Reader) Size() int64

// Read implements the [io.Reader] interface.
func (r *Reader) Read(b []byte) (n int, err error)

// ReadAt implements the [io.ReaderAt] interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// ReadByte implements the [io.ByteReader] interface.
func (r *Reader) ReadByte() (byte, error)

// UnreadByte complements [Reader.ReadByte] in implementing the [io.ByteScanner] interface.
func (r *Reader) UnreadByte() error

// ReadRune implements the [io.RuneReader] interface.
func (r *Reader) ReadRune() (ch rune, size int, err error)

// UnreadRune complements [Reader.ReadRune] in implementing the [io.RuneScanner] interface.
func (r *Reader) UnreadRune() error

// Seek implements the [io.Seeker] interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteTo implements the [io.WriterTo] interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// Reset resets the [Reader] to be reading from b.
func (r *Reader) Reset(b []byte)

// NewReader returns a new [Reader] reading from b.
func NewReader(b []byte) *Reader
