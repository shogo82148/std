// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package elf implements access to ELF object files.

# Security

This package is not designed to be hardened against adversarial inputs, and is
outside the scope of https://go.dev/security/policy. In particular, only basic
validation is done when parsing object files. As such, care should be taken when
parsing untrusted inputs, as parsing malformed files may consume significant
resources, or cause panics.
*/
package elf

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/encoding/binary"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// seekStart, seekCurrent, seekEnd are copies of
// io.SeekStart, io.SeekCurrent, and io.SeekEnd.
// We can't use the ones from package io because
// we want this code to build with Go 1.4 during
// cmd/dist bootstrap.

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
	Entry      uint64
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

	FileSize uint64
}

// A Section represents a single section in an ELF file.
type Section struct {
	SectionHeader

	io.ReaderAt
	sr *io.SectionReader

	compressionType   CompressionType
	compressionOffset int64
}

// Data reads and returns the contents of the ELF section.
// Even if the section is stored compressed in the ELF file,
// Data returns uncompressed data.
//
// For an SHT_NOBITS section, Data always returns a non-nil error.
func (s *Section) Data() ([]byte, error)

// Open returns a new ReadSeeker reading the ELF section.
// Even if the section is stored compressed in the ELF file,
// the ReadSeeker reads uncompressed data.
//
// For an SHT_NOBITS section, all calls to the opened reader
// will return a non-nil error.
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

	Version string
	Library string
}

type FormatError struct {
	off int64
	msg string
	val any
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

// ErrNoSymbols is returned by File.Symbols and File.DynamicSymbols
// if there is no such section in the File.
var ErrNoSymbols = errors.New("no symbol section")

// Section returns a section with the given name, or nil if no such
// section exists.
func (f *File) Section(name string) *Section

func (f *File) DWARF() (*dwarf.Data, error)

// Symbols returns the symbol table for f. The symbols will be listed in the order
// they appear in f.
//
// For compatibility with Go 1.0, Symbols omits the null symbol at index 0.
// After retrieving the symbols as symtab, an externally supplied index x
// corresponds to symtab[x-1], not symtab[x].
func (f *File) Symbols() ([]Symbol, error)

// DynamicSymbols returns the dynamic symbol table for f. The symbols
// will be listed in the order they appear in f.
//
// If f has a symbol version table, the returned Symbols will have
// initialized Version and Library fields.
//
// For compatibility with Symbols, DynamicSymbols omits the null symbol at index 0.
// After retrieving the symbols as symtab, an externally supplied index x
// corresponds to symtab[x-1], not symtab[x].
func (f *File) DynamicSymbols() ([]Symbol, error)

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

// DynString returns the strings listed for the given tag in the file's dynamic
// section.
//
// The tag must be one that takes string values: DT_NEEDED, DT_SONAME, DT_RPATH, or
// DT_RUNPATH.
func (f *File) DynString(tag DynTag) ([]string, error)
