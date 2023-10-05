// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

// Writer provides sequential writing of a tar archive.
// Write.WriteHeader begins a new file with the provided Header,
// and then Writer can be treated as an io.Writer to supply that file's data.
type Writer struct {
	w    io.Writer
	pad  int64
	curr fileWriter
	hdr  Header
	blk  block

	// err is a persistent error.
	// It is only the responsibility of every exported method of Writer to
	// ensure that this error is sticky.
	err error
}

// NewWriter creates a new Writer writing to w.
func NewWriter(w io.Writer) *Writer

// Flush finishes writing the current file's block padding.
// The current file must be fully written before Flush can be called.
//
// This is unnecessary as the next call to WriteHeader or Close
// will implicitly flush out the file's padding.
func (tw *Writer) Flush() error

// WriteHeader writes hdr and prepares to accept the file's contents.
// The Header.Size determines how many bytes can be written for the next file.
// If the current file is not fully written, then this returns an error.
// This implicitly flushes any padding necessary before writing the header.
func (tw *Writer) WriteHeader(hdr *Header) error

// AddFS adds the files from fs.FS to the archive.
// It walks the directory tree starting at the root of the filesystem
// adding each file to the tar archive while maintaining the directory structure.
func (tw *Writer) AddFS(fsys fs.FS) error

// Write writes to the current file in the tar archive.
// Write returns the error ErrWriteTooLong if more than
// Header.Size bytes are written after WriteHeader.
//
// Calling Write on special types like TypeLink, TypeSymlink, TypeChar,
// TypeBlock, TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless
// of what the Header.Size claims.
func (tw *Writer) Write(b []byte) (int, error)

// Close closes the tar archive by flushing the padding, and writing the footer.
// If the current file (from a prior call to WriteHeader) is not fully written,
// then this returns an error.
func (tw *Writer) Close() error
