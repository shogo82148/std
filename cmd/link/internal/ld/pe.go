// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// PE (Portable Executable) file writing
// https://docs.microsoft.com/en-us/windows/win32/debug/pe-format

package ld

import (
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

type IMAGE_IMPORT_DESCRIPTOR struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32
}

type IMAGE_EXPORT_DIRECTORY struct {
	Characteristics       uint32
	TimeDateStamp         uint32
	MajorVersion          uint16
	MinorVersion          uint16
	Name                  uint32
	Base                  uint32
	NumberOfFunctions     uint32
	NumberOfNames         uint32
	AddressOfFunctions    uint32
	AddressOfNames        uint32
	AddressOfNameOrdinals uint32
}

var (
	// PEBASE is the base address for the executable.
	// It is small for 32-bit and large for 64-bit.
	PEBASE int64

	// SectionAlignment must be greater than or equal to FileAlignment.
	// The default is the page size for the architecture.
	PESECTALIGN int64 = 0x1000

	// FileAlignment should be a power of 2 between 512 and 64 K, inclusive.
	// The default is 512. If the SectionAlignment is less than
	// the architecture's page size, then FileAlignment must match SectionAlignment.
	PEFILEALIGN int64 = 2 << 8
)

const (
	IMAGE_SCN_CNT_CODE               = 0x00000020
	IMAGE_SCN_CNT_INITIALIZED_DATA   = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA = 0x00000080
	IMAGE_SCN_LNK_OTHER              = 0x00000100
	IMAGE_SCN_LNK_INFO               = 0x00000200
	IMAGE_SCN_LNK_REMOVE             = 0x00000800
	IMAGE_SCN_LNK_COMDAT             = 0x00001000
	IMAGE_SCN_GPREL                  = 0x00008000
	IMAGE_SCN_MEM_PURGEABLE          = 0x00020000
	IMAGE_SCN_MEM_16BIT              = 0x00020000
	IMAGE_SCN_MEM_LOCKED             = 0x00040000
	IMAGE_SCN_MEM_PRELOAD            = 0x00080000
	IMAGE_SCN_ALIGN_1BYTES           = 0x00100000
	IMAGE_SCN_ALIGN_2BYTES           = 0x00200000
	IMAGE_SCN_ALIGN_4BYTES           = 0x00300000
	IMAGE_SCN_ALIGN_8BYTES           = 0x00400000
	IMAGE_SCN_ALIGN_16BYTES          = 0x00500000
	IMAGE_SCN_ALIGN_32BYTES          = 0x00600000
	IMAGE_SCN_ALIGN_64BYTES          = 0x00700000
	IMAGE_SCN_ALIGN_128BYTES         = 0x00800000
	IMAGE_SCN_ALIGN_256BYTES         = 0x00900000
	IMAGE_SCN_ALIGN_512BYTES         = 0x00A00000
	IMAGE_SCN_ALIGN_1024BYTES        = 0x00B00000
	IMAGE_SCN_ALIGN_2048BYTES        = 0x00C00000
	IMAGE_SCN_ALIGN_4096BYTES        = 0x00D00000
	IMAGE_SCN_ALIGN_8192BYTES        = 0x00E00000
	IMAGE_SCN_LNK_NRELOC_OVFL        = 0x01000000
	IMAGE_SCN_MEM_DISCARDABLE        = 0x02000000
	IMAGE_SCN_MEM_NOT_CACHED         = 0x04000000
	IMAGE_SCN_MEM_NOT_PAGED          = 0x08000000
	IMAGE_SCN_MEM_SHARED             = 0x10000000
	IMAGE_SCN_MEM_EXECUTE            = 0x20000000
	IMAGE_SCN_MEM_READ               = 0x40000000
	IMAGE_SCN_MEM_WRITE              = 0x80000000
)

// See https://docs.microsoft.com/en-us/windows/win32/debug/pe-format.
// TODO(crawshaw): add these constants to debug/pe.
const (
	IMAGE_SYM_TYPE_NULL      = 0
	IMAGE_SYM_TYPE_STRUCT    = 8
	IMAGE_SYM_DTYPE_FUNCTION = 2
	IMAGE_SYM_DTYPE_ARRAY    = 3
	IMAGE_SYM_CLASS_EXTERNAL = 2
	IMAGE_SYM_CLASS_STATIC   = 3

	IMAGE_REL_I386_DIR32   = 0x0006
	IMAGE_REL_I386_DIR32NB = 0x0007
	IMAGE_REL_I386_SECREL  = 0x000B
	IMAGE_REL_I386_REL32   = 0x0014

	IMAGE_REL_AMD64_ADDR64   = 0x0001
	IMAGE_REL_AMD64_ADDR32   = 0x0002
	IMAGE_REL_AMD64_ADDR32NB = 0x0003
	IMAGE_REL_AMD64_REL32    = 0x0004
	IMAGE_REL_AMD64_SECREL   = 0x000B

	IMAGE_REL_ARM_ABSOLUTE = 0x0000
	IMAGE_REL_ARM_ADDR32   = 0x0001
	IMAGE_REL_ARM_ADDR32NB = 0x0002
	IMAGE_REL_ARM_BRANCH24 = 0x0003
	IMAGE_REL_ARM_BRANCH11 = 0x0004
	IMAGE_REL_ARM_SECREL   = 0x000F

	IMAGE_REL_ARM64_ABSOLUTE       = 0x0000
	IMAGE_REL_ARM64_ADDR32         = 0x0001
	IMAGE_REL_ARM64_ADDR32NB       = 0x0002
	IMAGE_REL_ARM64_BRANCH26       = 0x0003
	IMAGE_REL_ARM64_PAGEBASE_REL21 = 0x0004
	IMAGE_REL_ARM64_REL21          = 0x0005
	IMAGE_REL_ARM64_PAGEOFFSET_12A = 0x0006
	IMAGE_REL_ARM64_PAGEOFFSET_12L = 0x0007
	IMAGE_REL_ARM64_SECREL         = 0x0008
	IMAGE_REL_ARM64_SECREL_LOW12A  = 0x0009
	IMAGE_REL_ARM64_SECREL_HIGH12A = 0x000A
	IMAGE_REL_ARM64_SECREL_LOW12L  = 0x000B
	IMAGE_REL_ARM64_TOKEN          = 0x000C
	IMAGE_REL_ARM64_SECTION        = 0x000D
	IMAGE_REL_ARM64_ADDR64         = 0x000E
	IMAGE_REL_ARM64_BRANCH19       = 0x000F
	IMAGE_REL_ARM64_BRANCH14       = 0x0010
	IMAGE_REL_ARM64_REL32          = 0x0011

	IMAGE_REL_BASED_HIGHLOW = 3
	IMAGE_REL_BASED_DIR64   = 10
)

const (
	PeMinimumTargetMajorVersion = 6
	PeMinimumTargetMinorVersion = 1
)

type Imp struct {
	s       loader.Sym
	off     uint64
	next    *Imp
	argsize int
}

type Dll struct {
	name     string
	nameoff  uint64
	thunkoff uint64
	ms       *Imp
	next     *Dll
}

var (
	PESECTHEADR int32
	PEFILEHEADR int32
)

func AddPELabelSym(ldr *loader.Loader, s loader.Sym)

func Peinit(ctxt *Link)
