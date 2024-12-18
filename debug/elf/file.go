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
	Sections    []*Section
	Progs       []*Prog
	closer      io.Closer
	dynVers     []DynamicVersion
	dynVerNeeds []DynamicVersionNeed
	gnuVersym   []byte
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

	// FileSize is the size of this section in the file in bytes.
	// If a section is compressed, FileSize is the size of the
	// compressed data, while Size (above) is the size of the
	// uncompressed data.
	FileSize uint64
}

// A Section represents a single section in an ELF file.
type Section struct {
	SectionHeader

	// Embed ReaderAt for ReadAt method.
	// Do not embed SectionReader directly
	// to avoid having Read and Seek.
	// If a client wants Read and Seek it must use
	// Open() to avoid fighting over the seek offset
	// with other clients.
	//
	// ReaderAt may be nil if the section is not easily available
	// in a random-access form. For example, a compressed section
	// may have a nil ReaderAt.
	io.ReaderAt
	sr *io.SectionReader

	compressionType   CompressionType
	compressionOffset int64
}

// Data reads and returns the contents of the ELF section.
// Even if the section is stored compressed in the ELF file,
// Data returns uncompressed data.
//
// For an [SHT_NOBITS] section, Data always returns a non-nil error.
func (s *Section) Data() ([]byte, error)

// Open returns a new ReadSeeker reading the ELF section.
// Even if the section is stored compressed in the ELF file,
// the ReadSeeker reads uncompressed data.
//
// For an [SHT_NOBITS] section, all calls to the opened reader
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

	// Embed ReaderAt for ReadAt method.
	// Do not embed SectionReader directly
	// to avoid having Read and Seek.
	// If a client wants Read and Seek it must use
	// Open() to avoid fighting over the seek offset
	// with other clients.
	io.ReaderAt
	sr *io.SectionReader
}

// Open returns a new ReadSeeker reading the ELF program body.
func (p *Prog) Open() io.ReadSeeker

// A Symbol represents an entry in an ELF symbol table section.
type Symbol struct {
	Name        string
	Info, Other byte

	// VersionScope describes the version in which the symbol is defined.
	// This is only set for the dynamic symbol table.
	// When no symbol versioning information is available,
	// this is VersionScopeNone.
	VersionScope SymbolVersionScope
	// VersionIndex is the version index.
	// This is only set if VersionScope is VersionScopeSpecific or
	// VersionScopeHidden. This is only set for the dynamic symbol table.
	// This index will match either [DynamicVersion.Index]
	// in the slice returned by [File.DynamicVersions],
	// or [DynamicVersiondep.Index] in the Needs field
	// of the elements of the slice returned by [File.DynamicVersionNeeds].
	// In general, a defined symbol will have an index referring
	// to DynamicVersions, and an undefined symbol will have an index
	// referring to some version in DynamicVersionNeeds.
	VersionIndex int16

	Section     SectionIndex
	Value, Size uint64

	// These fields are present only for the dynamic symbol table.
	Version string
	Library string
}

type FormatError struct {
	off int64
	msg string
	val any
}

func (e *FormatError) Error() string

// Open opens the named file using [os.Open] and prepares it for use as an ELF binary.
func Open(name string) (*File, error)

// Close closes the [File].
// If the [File] was created using [NewFile] directly instead of [Open],
// Close has no effect.
func (f *File) Close() error

// SectionByType returns the first section in f with the
// given type, or nil if there is no such section.
func (f *File) SectionByType(typ SectionType) *Section

// NewFile creates a new [File] for accessing an ELF binary in an underlying reader.
// The ELF binary is expected to start at position 0 in the ReaderAt.
func NewFile(r io.ReaderAt) (*File, error)

// ErrNoSymbols is returned by [File.Symbols] and [File.DynamicSymbols]
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
// If f has a symbol version table, the returned [File.Symbols] will have
// initialized Version and Library fields.
//
// For compatibility with [File.Symbols], [File.DynamicSymbols] omits the null symbol at index 0.
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

// SymbolVersionScope describes the version in which a [Symbol] is defined.
// This is only used for the dynamic symbol table.
type SymbolVersionScope byte

const (
	VersionScopeNone SymbolVersionScope = iota
	VersionScopeLocal
	VersionScopeGlobal
	VersionScopeSpecific
	VersionScopeHidden
)

// DynamicVersion is a version defined by a dynamic object.
// This describes entries in the ELF SHT_GNU_verdef section.
// We assume that the vd_version field is 1.
// Note that the name of the version appears here;
// it is not in the first Deps entry as it is in the ELF file.
type DynamicVersion struct {
	Name  string
	Index uint16
	Flags DynamicVersionFlag
	Deps  []string
}

// DynamicVersionNeed describes a shared library needed by a dynamic object,
// with a list of the versions needed from that shared library.
// This describes entries in the ELF SHT_GNU_verneed section.
// We assume that the vn_version field is 1.
type DynamicVersionNeed struct {
	Name  string
	Needs []DynamicVersionDep
}

// DynamicVersionDep is a version needed from some shared library.
type DynamicVersionDep struct {
	Flags DynamicVersionFlag
	Index uint16
	Dep   string
}

// DynamicVersions returns version information for a dynamic object.
func (f *File) DynamicVersions() ([]DynamicVersion, error)

// DynamicVersionNeeds returns version dependencies for a dynamic object.
func (f *File) DynamicVersionNeeds() ([]DynamicVersionNeed, error)

// ImportedLibraries returns the names of all libraries
// referred to by the binary f that are expected to be
// linked with the binary at dynamic link time.
func (f *File) ImportedLibraries() ([]string, error)

// DynString returns the strings listed for the given tag in the file's dynamic
// section.
//
// The tag must be one that takes string values: [DT_NEEDED], [DT_SONAME], [DT_RPATH], or
// [DT_RUNPATH].
func (f *File) DynString(tag DynTag) ([]string, error)

// DynValue returns the values listed for the given tag in the file's dynamic
// section.
func (f *File) DynValue(tag DynTag) ([]uint64, error)
