// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package elf implements access to ELF object files.
package elf

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/encoding/binary"
	"github.com/shogo82148/std/io"
)

// A FileHeader represents an ELF file header.
type FileHeader struct {
	Class      Class
	Data       Data
	Version    Version
	OSABI      OSABI
	ABIVersion uint8
	ByteOrder  binary.ByteOrder
	Type       Type
	Machine    Machine
}

// A File represents an open ELF file.
type File struct {
	FileHeader
	Sections  []*Section
	Progs     []*Prog
	closer    io.Closer
	gnuNeed   []verneed
	gnuVersym []byte
}

// A SectionHeader represents a single ELF section header.
type SectionHeader struct {
	Name      string
	Type      SectionType
	Flags     SectionFlag
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64
}

// A Section represents a single section in an ELF file.
type Section struct {
	SectionHeader

	io.ReaderAt
	sr *io.SectionReader
}

// Data reads and returns the contents of the ELF section.
func (s *Section) Data() ([]byte, error)

// Open returns a new ReadSeeker reading the ELF section.
func (s *Section) Open() io.ReadSeeker

// A ProgHeader represents a single ELF program header.
type ProgHeader struct {
	Type   ProgType
	Flags  ProgFlag
	Off    uint64
	Vaddr  uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}

// A Prog represents a single ELF program header in an ELF binary.
type Prog struct {
	ProgHeader

	io.ReaderAt
	sr *io.SectionReader
}

// Open returns a new ReadSeeker reading the ELF program body.
func (p *Prog) Open() io.ReadSeeker

// A Symbol represents an entry in an ELF symbol table section.
type Symbol struct {
	Name        string
	Info, Other byte
	Section     SectionIndex
	Value, Size uint64
}

type FormatError struct {
	off int64
	msg string
	val interface{}
}

func (e *FormatError) Error() string

// Open opens the named file using os.Open and prepares it for use as an ELF binary.
func Open(name string) (*File, error)

// Close closes the File.
// If the File was created using NewFile directly instead of Open,
// Close has no effect.
func (f *File) Close() error

// SectionByType returns the first section in f with the
// given type, or nil if there is no such section.
func (f *File) SectionByType(typ SectionType) *Section

// NewFile creates a new File for accessing an ELF binary in an underlying reader.
// The ELF binary is expected to start at position 0 in the ReaderAt.
func NewFile(r io.ReaderAt) (*File, error)

// Section returns a section with the given name, or nil if no such
// section exists.
func (f *File) Section(name string) *Section

func (f *File) DWARF() (*dwarf.Data, error)

// Symbols returns the symbol table for f.
func (f *File) Symbols() ([]Symbol, error)

type ImportedSymbol struct {
	Name    string
	Version string
	Library string
}

// ImportedSymbols returns the names of all symbols
// referred to by the binary f that are expected to be
// satisfied by other libraries at dynamic load time.
// It does not return weak symbols.
func (f *File) ImportedSymbols() ([]ImportedSymbol, error)

// ImportedLibraries returns the names of all libraries
// referred to by the binary f that are expected to be
// linked with the binary at dynamic link time.
func (f *File) ImportedLibraries() ([]string, error)
