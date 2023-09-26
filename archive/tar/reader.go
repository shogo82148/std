// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"github.com/shogo82148/std/io"
)

// Reader provides sequential access to the contents of a tar archive.
// Reader.Next advances to the next file in the archive (including the first),
// and then Reader can be treated as an io.Reader to access the file's data.
type Reader struct {
	r    io.Reader
	pad  int64
	curr fileReader
	blk  block

	err error
}

// NewReader creates a new Reader reading from r.
func NewReader(r io.Reader) *Reader

// Next advances to the next entry in the tar archive.
// The Header.Size determines how many bytes can be read for the next file.
// Any remaining data in the current file is automatically discarded.
//
// io.EOF is returned at the end of the input.
func (tr *Reader) Next() (*Header, error)

// Read reads from the current file in the tar archive.
// It returns (0, io.EOF) when it reaches the end of that file,
// until Next is called to advance to the next file.
//
// If the current file is sparse, then the regions marked as a hole
// are read back as NUL-bytes.
//
// Calling Read on special types like TypeLink, TypeSymlink, TypeChar,
// TypeBlock, TypeDir, and TypeFifo returns (0, io.EOF) regardless of what
// the Header.Size claims.
func (tr *Reader) Read(b []byte) (int, error)

// regFileReader is a fileReader for reading data from a regular file entry.

// sparseFileReader is a fileReader for reading data from a sparse file entry.
