// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package loadelf implements an ELF file reader.
package loadelf

import (
	"github.com/shogo82148/std/cmd/internal/bio"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/loader"
	"github.com/shogo82148/std/debug/elf"
	"github.com/shogo82148/std/encoding/binary"
)

const (
	SHT_ARM_ATTRIBUTES = 0x70000003
)

type ElfSect struct {
	name        string
	nameoff     uint32
	type_       elf.SectionType
	flags       elf.SectionFlag
	addr        uint64
	off         uint64
	size        uint64
	link        uint32
	info        uint32
	align       uint64
	entsize     uint64
	base        []byte
	readOnlyMem bool
	sym         loader.Sym
}

type ElfObj struct {
	f         *bio.Reader
	base      int64
	length    int64
	is64      int
	name      string
	e         binary.ByteOrder
	sect      []ElfSect
	nsect     uint
	nsymtab   int
	symtab    *ElfSect
	symstr    *ElfSect
	type_     uint32
	machine   uint32
	version   uint32
	entry     uint64
	phoff     uint64
	shoff     uint64
	flags     uint32
	ehsize    uint32
	phentsize uint32
	phnum     uint32
	shentsize uint32
	shnum     uint32
	shstrndx  uint32
}

type ElfSym struct {
	name  string
	value uint64
	size  uint64
	bind  elf.SymBind
	type_ elf.SymType
	other uint8
	shndx elf.SectionIndex
	sym   loader.Sym
}

const (
	TagFile               = 1
	TagCPUName            = 4
	TagCPURawName         = 5
	TagCompatibility      = 32
	TagNoDefaults         = 64
	TagAlsoCompatibleWith = 65
	TagABIVFPArgs         = 28
)

// Load loads the ELF file pn from f.
// Symbols are installed into the loader, and a slice of the text symbols is returned.
//
// On ARM systems, Load will attempt to determine what ELF header flags to
// emit by scanning the attributes in the ELF file being loaded. The
// parameter initEhdrFlags contains the current header flags for the output
// object, and the returned ehdrFlags contains what this Load function computes.
// TODO: find a better place for this logic.
func Load(l *loader.Loader, arch *sys.Arch, localSymVersion int, f *bio.Reader, pkg string, length int64, pn string, initEhdrFlags uint32) (textp []loader.Sym, ehdrFlags uint32, err error)
