// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/shogo82148/std/io"
)

// A Reader implements the [io.Reader], [io.ReaderAt], [io.ByteReader], [io.ByteScanner],
// [io.RuneReader], [io.RuneScanner], [io.Seeker], and [io.WriterTo] interfaces by reading
// from a string.
// The zero value for Reader operates like a Reader of an empty string.
type Reader struct {
	s        string
	i        int64
	prevRune int
}

// Len returns the number of bytes of the unread portion of the
// string.
func (r *Reader) Len() int

// Size returns the original length of the underlying string.
// Size is the number of bytes available for reading via [Reader.ReadAt].
// The returned value is always the same and is not affected by calls
// to any other method.
func (r *Reader) Size() int64

// Read implements the [io.Reader] interface.
func (r *Reader) Read(b []byte) (n int, err error)

// ReadAt implements the [io.ReaderAt] interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// ReadByte implements the [io.ByteReader] interface.
func (r *Reader) ReadByte() (byte, error)

// UnreadByte implements the [io.ByteScanner] interface.
func (r *Reader) UnreadByte() error

// ReadRune implements the [io.RuneReader] interface.
func (r *Reader) ReadRune() (ch rune, size int, err error)

// UnreadRune implements the [io.RuneScanner] interface.
func (r *Reader) UnreadRune() error

// Seek implements the [io.Seeker] interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteTo implements the [io.WriterTo] interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// Reset resets the [Reader] to be reading from s.
func (r *Reader) Reset(s string)

// NewReader returns a new [Reader] reading from s.
// It is similar to [bytes.NewBufferString] but more efficient and non-writable.
func NewReader(s string) *Reader
