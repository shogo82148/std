// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/io"
)

// A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
// io.ByteScanner, and io.RuneScanner interfaces by reading from
// a byte slice.
// Unlike a Buffer, a Reader is read-only and supports seeking.
type Reader struct {
	s        []byte
	i        int64
	prevRune int
}

// Len returns the number of bytes of the unread portion of the
// slice.
func (r *Reader) Len() int

// Size returns the original length of the underlying byte slice.
// Size is the number of bytes available for reading via ReadAt.
// The returned value is always the same and is not affected by calls
// to any other method.
func (r *Reader) Size() int64

func (r *Reader) Read(b []byte) (n int, err error)

func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

func (r *Reader) ReadByte() (byte, error)

func (r *Reader) UnreadByte() error

func (r *Reader) ReadRune() (ch rune, size int, err error)

func (r *Reader) UnreadRune() error

// Seek implements the io.Seeker interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteTo implements the io.WriterTo interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// Reset resets the Reader to be reading from b.
func (r *Reader) Reset(b []byte)

// NewReader returns a new Reader reading from b.
func NewReader(b []byte) *Reader
