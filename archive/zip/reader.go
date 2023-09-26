// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
)

var (
	ErrFormat    = errors.New("zip: not a valid zip file")
	ErrAlgorithm = errors.New("zip: unsupported compression algorithm")
	ErrChecksum  = errors.New("zip: checksum error")
)

type Reader struct {
	r       io.ReaderAt
	File    []*File
	Comment string
}

type ReadCloser struct {
	f *os.File
	Reader
}

type File struct {
	FileHeader
	zipr         io.ReaderAt
	zipsize      int64
	headerOffset int64
}

// OpenReader will open the Zip file specified by name and return a ReadCloser.
func OpenReader(name string) (*ReadCloser, error)

// NewReader returns a new Reader reading from r, which is assumed to
// have the given size in bytes.
func NewReader(r io.ReaderAt, size int64) (*Reader, error)

// Close closes the Zip file, rendering it unusable for I/O.
func (rc *ReadCloser) Close() error

// Open returns a ReadCloser that provides access to the File's contents.
// Multiple files may be read concurrently.
func (f *File) Open() (rc io.ReadCloser, err error)
