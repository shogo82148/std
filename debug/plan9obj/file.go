// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package plan9obj implements access to Plan 9 a.out object files.

# Security

This package is not designed to be hardened against adversarial inputs, and is
outside the scope of https://go.dev/security/policy. In particular, only basic
validation is done when parsing object files. As such, care should be taken when
parsing untrusted inputs, as parsing malformed files may consume significant
resources, or cause panics.
*/
package plan9obj

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// A FileHeader represents a Plan 9 a.out file header.
type FileHeader struct {
	Magic       uint32
	Bss         uint32
	Entry       uint64
	PtrSize     int
	LoadAddress uint64
	HdrSize     uint64
}

// A File represents an open Plan 9 a.out file.
type File struct {
	FileHeader
	Sections []*Section
	closer   io.Closer
}

// A SectionHeader represents a single Plan 9 a.out section header.
// This structure doesn't exist on-disk, but eases navigation
// through the object file.
type SectionHeader struct {
	Name   string
	Size   uint32
	Offset uint32
}

// A Section represents a single section in a Plan 9 a.out file.
type Section struct {
	SectionHeader

	// Embed ReaderAt for ReadAt method.
	// Do not embed SectionReader directly
	// to avoid having Read and Seek.
	// If a client wants Read and Seek it must use
	// Open() to avoid fighting over the seek offset
	// with other clients.
	io.ReaderAt
	sr *io.SectionReader
}

// Data reads and returns the contents of the Plan 9 a.out section.
func (s *Section) Data() ([]byte, error)

// Open returns a new ReadSeeker reading the Plan 9 a.out section.
func (s *Section) Open() io.ReadSeeker

// A Symbol represents an entry in a Plan 9 a.out symbol table section.
type Sym struct {
	Value uint64
	Type  rune
	Name  string
}

// Open opens the named file using [os.Open] and prepares it for use as a Plan 9 a.out binary.
func Open(name string) (*File, error)

// Close closes the [File].
// If the [File] was created using [NewFile] directly instead of [Open],
// Close has no effect.
func (f *File) Close() error

// NewFile creates a new [File] for accessing a Plan 9 binary in an underlying reader.
// The Plan 9 binary is expected to start at position 0 in the ReaderAt.
func NewFile(r io.ReaderAt) (*File, error)

// ErrNoSymbols is returned by [File.Symbols] if there is no such section
// in the File.
var ErrNoSymbols = errors.New("no symbol section")

// Symbols returns the symbol table for f.
func (f *File) Symbols() ([]Sym, error)

// Section returns a section with the given name, or nil if no such
// section exists.
func (f *File) Section(name string) *Section
