// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pe implements access to PE (Microsoft Windows Portable Executable) files.
package pe

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/io"
)

// A File represents an open PE file.
type File struct {
	FileHeader
	Sections []*Section

	closer io.Closer
}

type SectionHeader struct {
	Name                 string
	VirtualSize          uint32
	VirtualAddress       uint32
	Size                 uint32
	Offset               uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}

type Section struct {
	SectionHeader

	io.ReaderAt
	sr *io.SectionReader
}

type ImportDirectory struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32

	dll string
}

// Data reads and returns the contents of the PE section.
func (s *Section) Data() ([]byte, error)

// Open returns a new ReadSeeker reading the PE section.
func (s *Section) Open() io.ReadSeeker

type FormatError struct {
	off int64
	msg string
	val interface{}
}

func (e *FormatError) Error() string

// Open opens the named file using os.Open and prepares it for use as a PE binary.
func Open(name string) (*File, error)

// Close closes the File.
// If the File was created using NewFile directly instead of Open,
// Close has no effect.
func (f *File) Close() error

// NewFile creates a new File for accessing a PE binary in an underlying reader.
func NewFile(r io.ReaderAt) (*File, error)

// Section returns the first section with the given name, or nil if no such
// section exists.
func (f *File) Section(name string) *Section

func (f *File) DWARF() (*dwarf.Data, error)

// ImportedSymbols returns the names of all symbols
// referred to by the binary f that are expected to be
// satisfied by other libraries at dynamic load time.
// It does not return weak symbols.
func (f *File) ImportedSymbols() ([]string, error)

// ImportedLibraries returns the names of all libraries
// referred to by the binary f that are expected to be
// linked with the binary at dynamic link time.
func (f *File) ImportedLibraries() ([]string, error)
