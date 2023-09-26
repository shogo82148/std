// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package plan9obj implements access to Plan 9 a.out object files.
package plan9obj

import (
	"github.com/shogo82148/std/io"
)

// A FileHeader represents a Plan 9 a.out file header.
type FileHeader struct {
	Magic   uint32
	Bss     uint32
	Entry   uint64
	PtrSize int
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

// formatError is returned by some operations if the data does
// not have the correct format for an object file.

// Open opens the named file using os.Open and prepares it for use as a Plan 9 a.out binary.
func Open(name string) (*File, error)

// Close closes the File.
// If the File was created using NewFile directly instead of Open,
// Close has no effect.
func (f *File) Close() error

// NewFile creates a new File for accessing a Plan 9 binary in an underlying reader.
// The Plan 9 binary is expected to start at position 0 in the ReaderAt.
func NewFile(r io.ReaderAt) (*File, error)

// Symbols returns the symbol table for f.
func (f *File) Symbols() ([]Sym, error)

// Section returns a section with the given name, or nil if no such
// section exists.
func (f *File) Section(name string) *Section
