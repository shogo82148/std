// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

import (
	"github.com/shogo82148/std/io"
)

// SectionHeader32 represents real PE COFF section header.
type SectionHeader32 struct {
	Name                 [8]uint8
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}

// _Reloc represents a PE COFF relocation.
// Each section contains its own relocation list.

// SectionHeader is similar to SectionHeader32 with Name
// field replaced by Go string.
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

// Section provides access to PE COFF section.
type Section struct {
	SectionHeader
	_Relocs []_Reloc

	io.ReaderAt
	sr *io.SectionReader
}

// Data reads and returns the contents of the PE section s.
func (s *Section) Data() ([]byte, error)

// Open returns a new ReadSeeker reading the PE section s.
func (s *Section) Open() io.ReadSeeker
