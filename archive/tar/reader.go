// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

var (
	ErrHeader = errors.New("archive/tar: invalid tar header")
)

// A Reader provides sequential access to the contents of a tar archive.
// A tar archive consists of a sequence of files.
// The Next method advances to the next file in the archive (including the first),
// and then it can be treated as an io.Reader to access the file's data.
type Reader struct {
	r       io.Reader
	err     error
	pad     int64
	curr    numBytesReader
	hdrBuff [blockSize]byte
}

// A numBytesReader is an io.Reader with a numBytes method, returning the number
// of bytes remaining in the underlying encoded data.

// A regFileReader is a numBytesReader for reading file data from a tar archive.

// A sparseFileReader is a numBytesReader for reading sparse file data from a tar archive.

// Keywords for GNU sparse files in a PAX extended header

// Keywords for old GNU sparse headers

// NewReader creates a new Reader reading from r.
func NewReader(r io.Reader) *Reader

// Next advances to the next entry in the tar archive.
func (tr *Reader) Next() (*Header, error)

// A sparseEntry holds a single entry in a sparse file's sparse map.
// A sparse entry indicates the offset and size in a sparse file of a
// block of data.

// Read reads from the current entry in the tar archive.
// It returns 0, io.EOF when it reaches the end of that entry,
// until Next is called to advance to the next entry.
func (tr *Reader) Read(b []byte) (n int, err error)
