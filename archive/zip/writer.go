// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/io"
)

// Writer implements a zip file writer.
type Writer struct {
	cw     *countWriter
	dir    []*header
	last   *fileWriter
	closed bool
}

// NewWriter returns a new Writer writing a zip file to w.
func NewWriter(w io.Writer) *Writer

// Close finishes writing the zip file by writing the central directory.
// It does not (and can not) close the underlying writer.
func (w *Writer) Close() error

// Create adds a file to the zip file using the provided name.
// It returns a Writer to which the file contents should be written.
// The file's contents must be written to the io.Writer before the next
// call to Create, CreateHeader, or Close.
func (w *Writer) Create(name string) (io.Writer, error)

// CreateHeader adds a file to the zip file using the provided FileHeader
// for the file metadata.
// It returns a Writer to which the file contents should be written.
// The file's contents must be written to the io.Writer before the next
// call to Create, CreateHeader, or Close.
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
