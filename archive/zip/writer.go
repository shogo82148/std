// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

// Writer implements a zip file writer.
type Writer struct {
	cw          *countWriter
	dir         []*header
	last        *fileWriter
	closed      bool
	compressors map[uint16]Compressor
	comment     string

	// testHookCloseSizeOffset if non-nil is called with the size
	// of offset of the central directory at Close.
	testHookCloseSizeOffset func(size, offset uint64)
}

// NewWriter returns a new [Writer] writing a zip file to w.
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
// It can only be called before [Writer.Close].
func (w *Writer) SetComment(comment string) error

// Close finishes writing the zip file by writing the central directory.
// It does not close the underlying writer.
func (w *Writer) Close() error

// Create adds a file to the zip file using the provided name.
// It returns a [Writer] to which the file contents should be written.
// The file contents will be compressed using the [Deflate] method.
// The name must be a relative path: it must not start with a drive
// letter (e.g. C:) or leading slash, and only forward slashes are
// allowed. To create a directory instead of a file, add a trailing
// slash to the name.
// The file's contents must be written to the [io.Writer] before the next
// call to [Writer.Create], [Writer.CreateHeader], or [Writer.Close].
func (w *Writer) Create(name string) (io.Writer, error)

// CreateHeader adds a file to the zip archive using the provided [FileHeader]
// for the file metadata. [Writer] takes ownership of fh and may mutate
// its fields. The caller must not modify fh after calling [Writer.CreateHeader].
//
// This returns a [Writer] to which the file contents should be written.
// The file's contents must be written to the io.Writer before the next
// call to [Writer.Create], [Writer.CreateHeader], [Writer.CreateRaw], or [Writer.Close].
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)

// CreateRaw adds a file to the zip archive using the provided [FileHeader] and
// returns a [Writer] to which the file contents should be written. The file's
// contents must be written to the io.Writer before the next call to [Writer.Create],
// [Writer.CreateHeader], [Writer.CreateRaw], or [Writer.Close].
//
// In contrast to [Writer.CreateHeader], the bytes passed to Writer are not compressed.
//
// CreateRaw's argument is stored in w. If the argument is a pointer to the embedded
// [FileHeader] in a [File] obtained from a [Reader] created from in-memory data,
// then w will refer to all of that memory.
func (w *Writer) CreateRaw(fh *FileHeader) (io.Writer, error)

// Copy copies the file f (obtained from a [Reader]) into w. It copies the raw
// form directly bypassing decompression, compression, and validation.
func (w *Writer) Copy(f *File) error

// RegisterCompressor registers or overrides a custom compressor for a specific
// method ID. If a compressor for a given method is not found, [Writer] will
// default to looking up the compressor at the package level.
func (w *Writer) RegisterCompressor(method uint16, comp Compressor)

// AddFS adds the files from fs.FS to the archive.
// It walks the directory tree starting at the root of the filesystem
// adding each file to the zip using deflate while maintaining the directory structure.
func (w *Writer) AddFS(fsys fs.FS) error
