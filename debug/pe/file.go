// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package pe implements access to PE (Microsoft Windows Portable Executable) files.

# Security

This package is not designed to be hardened against adversarial inputs, and is
outside the scope of https://go.dev/security/policy. In particular, only basic
validation is done when parsing object files. As such, care should be taken when
parsing untrusted inputs, as parsing malformed files may consume significant
resources, or cause panics.
*/
package pe

import (
	"github.com/shogo82148/std/debug/dwarf"
	"github.com/shogo82148/std/io"
)

// A File represents an open PE file.
type File struct {
	FileHeader
	OptionalHeader any
	Sections       []*Section
	Symbols        []*Symbol
	COFFSymbols    []COFFSymbol
	StringTable    StringTable

	closer io.Closer
}

// Open opens the named file using [os.Open] and prepares it for use as a PE binary.
func Open(name string) (*File, error)

// Close closes the [File].
// If the [File] was created using [NewFile] directly instead of [Open],
// Close has no effect.
func (f *File) Close() error

// NewFile creates a new [File] for accessing a PE binary in an underlying reader.
func NewFile(r io.ReaderAt) (*File, error)

// Section returns the first section with the given name, or nil if no such
// section exists.
func (f *File) Section(name string) *Section

func (f *File) DWARF() (*dwarf.Data, error)

type ImportDirectory struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32

	dll string
}

// ImportedSymbols returns the names of all symbols
// referred to by the binary f that are expected to be
// satisfied by other libraries at dynamic load time.
// It does not return weak symbols.
func (f *File) ImportedSymbols() ([]string, error)

// ImportedLibraries returns the names of all libraries
// referred to by the binary f that are expected to be
// linked with the binary at dynamic link time.
func (f *File) ImportedLibraries() ([]string, error)

// FormatError is unused.
// The type is retained for compatibility.
type FormatError struct {
}

func (e *FormatError) Error() string
