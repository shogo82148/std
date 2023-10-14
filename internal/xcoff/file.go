// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package xcoff implements access to XCOFF (Extended Common Object File Format) files.
package xcoff

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/io"
)

// SectionHeader holds information about an XCOFF section header.
type SectionHeader struct {
	Name           string
	VirtualAddress uint64
	Size           uint64
	Type           uint32
	Relptr         uint64
	Nreloc         uint32
}

type Section struct {
	SectionHeader
	Relocs []Reloc
	io.ReaderAt
	sr *io.SectionReader
}

// AuxiliaryCSect holds information about an XCOFF symbol in an AUX_CSECT entry.
type AuxiliaryCSect struct {
	Length              int64
	StorageMappingClass int
	SymbolType          int
}

// AuxiliaryFcn holds information about an XCOFF symbol in an AUX_FCN entry.
type AuxiliaryFcn struct {
	Size int64
}

type Symbol struct {
	Name          string
	Value         uint64
	SectionNumber int
	StorageClass  int
	AuxFcn        AuxiliaryFcn
	AuxCSect      AuxiliaryCSect
}

type Reloc struct {
	VirtualAddress   uint64
	Symbol           *Symbol
	Signed           bool
	InstructionFixed bool
	Length           uint8
	Type             uint8
}

// ImportedSymbol holds information about an imported XCOFF symbol.
type ImportedSymbol struct {
	Name    string
	Library string
}

// FileHeader holds information about an XCOFF file header.
type FileHeader struct {
	TargetMachine uint16
}

// A File represents an open XCOFF file.
type File struct {
	FileHeader
	Sections     []*Section
	Symbols      []*Symbol
	StringTable  []byte
	LibraryPaths []string

	closer io.Closer
}

// Open opens the named file using os.Open and prepares it for use as an XCOFF binary.
func Open(name string) (*File, error)

// Close closes the File.
// If the File was created using NewFile directly instead of Open,
// Close has no effect.
func (f *File) Close() error

// Section returns the first section with the given name, or nil if no such
// section exists.
// Xcoff have section's name limited to 8 bytes. Some sections like .gosymtab
// can be trunked but this method will still find them.
func (f *File) Section(name string) *Section

// SectionByType returns the first section in f with the
// given type, or nil if there is no such section.
func (f *File) SectionByType(typ uint32) *Section

// NewFile creates a new File for accessing an XCOFF binary in an underlying reader.
func NewFile(r io.ReaderAt) (*File, error)

// Data reads and returns the contents of the XCOFF section s.
func (s *Section) Data() ([]byte, error)

// CSect reads and returns the contents of a csect.
func (f *File) CSect(name string) []byte

func (f *File) DWARF() (*dwarf.Data, error)

// ImportedSymbols returns the names of all symbols
// referred to by the binary f that are expected to be
// satisfied by other libraries at dynamic load time.
// It does not return weak symbols.
func (f *File) ImportedSymbols() ([]ImportedSymbol, error)

// ImportedLibraries returns the names of all libraries
// referred to by the binary f that are expected to be
// linked with the binary at dynamic link time.
func (f *File) ImportedLibraries() ([]string, error)
