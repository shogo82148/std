// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bio implements common I/O abstractions used within the Go toolchain.
package bio

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/os"
)

// Reader implements a seekable buffered io.Reader.
type Reader struct {
	f *os.File
	*bufio.Reader
}

// Writer implements a seekable buffered io.Writer.
type Writer struct {
	f *os.File
	*bufio.Writer
}

// Create creates the file named name and returns a Writer
// for that file.
func Create(name string) (*Writer, error)

// Open returns a Reader for the file named name.
func Open(name string) (*Reader, error)

// NewReader returns a Reader from an open file.
func NewReader(f *os.File) *Reader

func (r *Reader) MustSeek(offset int64, whence int) int64

func (w *Writer) MustSeek(offset int64, whence int) int64

func (r *Reader) Offset() int64

func (w *Writer) Offset() int64

func (r *Reader) Close() error

func (w *Writer) Close() error

func (r *Reader) File() *os.File

func (w *Writer) File() *os.File

// Slice reads the next length bytes of r into a slice.
//
// This slice may be backed by mmap'ed memory. Currently, this memory
// will never be unmapped. The second result reports whether the
// backing memory is read-only.
func (r *Reader) Slice(length uint64) ([]byte, bool, error)

// SliceRO returns a slice containing the next length bytes of r
// backed by a read-only mmap'd data. If the mmap cannot be
// established (limit exceeded, region too small, etc) a nil slice
// will be returned. If mmap succeeds, it will never be unmapped.
func (r *Reader) SliceRO(length uint64) []byte
