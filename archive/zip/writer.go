// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/io"
)

// Writer implements a zip file writer.
type Writer struct {
	cw          *countWriter
	dir         []*header
	last        *fileWriter
	closed      bool
	compressors map[uint16]Compressor
	comment     string

	testHookCloseSizeOffset func(size, offset uint64)
}

// NewWriter returns a new Writer writing a zip file to w.
func NewWriter(w io.Writer) *Writer

// SetOffset sets the offset of the beginning of the zip data within the
// underlying writer. It should be used when the zip data is appended to an
// existing file, such as a binary executable.
// It must be called before any data is written.
func (w *Writer) SetOffset(n int64)

// Flush flushes any buffered data to the underlying writer.
// Calling Flush is not normally necessary; calling Close is sufficient.
func (w *Writer) Flush() error

// SetComment sets the end-of-central-directory comment field.
// It can only be called before Close.
func (w *Writer) SetComment(comment string) error

// Close finishes writing the zip file by writing the central directory.
// It does not (and cannot) close the underlying writer.
func (w *Writer) Close() error

// Create adds a file to the zip file using the provided name.
// It returns a Writer to which the file contents should be written.
// The file contents will be compressed using the Deflate method.
// The name must be a relative path: it must not start with a drive
// letter (e.g. C:) or leading slash, and only forward slashes are
// allowed.
// The file's contents must be written to the io.Writer before the next
// call to Create, CreateHeader, or Close.
func (w *Writer) Create(name string) (io.Writer, error)

// CreateHeader adds a file to the zip archive using the provided FileHeader
// for the file metadata. Writer takes ownership of fh and may mutate
// its fields. The caller must not modify fh after calling CreateHeader.
//
// This returns a Writer to which the file contents should be written.
// The file's contents must be written to the io.Writer before the next
// call to Create, CreateHeader, or Close.
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)

// RegisterCompressor registers or overrides a custom compressor for a specific
// method ID. If a compressor for a given method is not found, Writer will
// default to looking up the compressor at the package level.
func (w *Writer) RegisterCompressor(method uint16, comp Compressor)
