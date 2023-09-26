// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

var (
	ErrWriteTooLong    = errors.New("archive/tar: write too long")
	ErrFieldTooLong    = errors.New("archive/tar: header field too long")
	ErrWriteAfterClose = errors.New("archive/tar: write after close")
)

// A Writer provides sequential writing of a tar archive in POSIX.1 format.
// A tar archive consists of a sequence of files.
// Call WriteHeader to begin a new file, and then call Write to supply that file's data,
// writing at most hdr.Size bytes in total.
//
// Example:
//
//	tw := tar.NewWriter(w)
//	hdr := new(Header)
//	hdr.Size = length of data in bytes
//	// populate other hdr fields as desired
//	if err := tw.WriteHeader(hdr); err != nil {
//		// handle error
//	}
//	io.Copy(tw, data)
//	tw.Close()
type Writer struct {
	w          io.Writer
	err        error
	nb         int64
	pad        int64
	closed     bool
	usedBinary bool
}

// NewWriter creates a new Writer writing to w.
func NewWriter(w io.Writer) *Writer

// Flush finishes writing the current file (optional).
func (tw *Writer) Flush() error

// WriteHeader writes hdr and prepares to accept the file's contents.
// WriteHeader calls Flush if it is not the first header.
// Calling after a Close will return ErrWriteAfterClose.
func (tw *Writer) WriteHeader(hdr *Header) error

// Write writes to the current entry in the tar archive.
// Write returns the error ErrWriteTooLong if more than
// hdr.Size bytes are written after WriteHeader.
func (tw *Writer) Write(b []byte) (n int, err error)

// Close closes the tar archive, flushing any unwritten
// data to the underlying writer.
func (tw *Writer) Close() error
