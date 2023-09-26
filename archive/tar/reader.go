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

// A sparseFileReader is a numBytesReader for reading sparse file data from a
// tar archive.

// A sparseEntry holds a single entry in a sparse file's sparse map.
//
// Sparse files are represented using a series of sparseEntrys.
// Despite the name, a sparseEntry represents an actual data fragment that
// references data found in the underlying archive stream. All regions not
// covered by a sparseEntry are logically filled with zeros.
//
// For example, if the underlying raw file contains the 10-byte data:
//	var compactData = "abcdefgh"
//
// And the sparse map has the following entries:
//	var sp = []sparseEntry{
//		{offset: 2,  numBytes: 5} // Data fragment for [2..7]
//		{offset: 18, numBytes: 3} // Data fragment for [18..21]
//	}
//
// Then the content of the resulting sparse file with a "real" size of 25 is:
//	var sparseData = "\x00"*2 + "abcde" + "\x00"*11 + "fgh" + "\x00"*4

// Keywords for GNU sparse files in a PAX extended header

// Keywords for old GNU sparse headers

// NewReader creates a new Reader reading from r.
func NewReader(r io.Reader) *Reader

// Next advances to the next entry in the tar archive.
//
// io.EOF is returned at the end of the input.
func (tr *Reader) Next() (*Header, error)

// Read reads from the current entry in the tar archive.
// It returns 0, io.EOF when it reaches the end of that entry,
// until Next is called to advance to the next entry.
//
// Calling Read on special types like TypeLink, TypeSymLink, TypeChar,
// TypeBlock, TypeDir, and TypeFifo returns 0, io.EOF regardless of what
// the Header.Size claims.
func (tr *Reader) Read(b []byte) (n int, err error)
