// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

type MachoHdr struct {
	cpu    uint32
	subcpu uint32
}

type MachoSect struct {
	name    string
	segname string
	addr    uint64
	size    uint64
	off     uint32
	align   uint32
	reloc   uint32
	nreloc  uint32
	flag    uint32
	res1    uint32
	res2    uint32
}

type MachoSeg struct {
	name       string
	vsize      uint64
	vaddr      uint64
	fileoffset uint64
	filesize   uint64
	prot1      uint32
	prot2      uint32
	nsect      uint32
	msect      uint32
	sect       []MachoSect
	flag       uint32
}

// MachoPlatformLoad represents a LC_VERSION_MIN_* or
// LC_BUILD_VERSION load command.
type MachoPlatformLoad struct {
	platform MachoPlatform
	cmd      MachoLoad
}

type MachoLoad struct {
	type_ uint32
	data  []uint32
}

type MachoPlatform int

/*
 * Total amount of space to reserve at the start of the file
 * for Header, PHeaders, and SHeaders.
 * May waste some.
 */
const (
	INITIAL_MACHO_HEADR = 4 * 1024
)

const (
	MACHO_CPU_AMD64                      = 1<<24 | 7
	MACHO_CPU_386                        = 7
	MACHO_SUBCPU_X86                     = 3
	MACHO_CPU_ARM                        = 12
	MACHO_SUBCPU_ARM                     = 0
	MACHO_SUBCPU_ARMV7                   = 9
	MACHO_CPU_ARM64                      = 1<<24 | 12
	MACHO_SUBCPU_ARM64_ALL               = 0
	MACHO_SUBCPU_ARM64_V8                = 1
	MACHO_SUBCPU_ARM64E                  = 2
	MACHO32SYMSIZE                       = 12
	MACHO64SYMSIZE                       = 16
	MACHO_X86_64_RELOC_UNSIGNED          = 0
	MACHO_X86_64_RELOC_SIGNED            = 1
	MACHO_X86_64_RELOC_BRANCH            = 2
	MACHO_X86_64_RELOC_GOT_LOAD          = 3
	MACHO_X86_64_RELOC_GOT               = 4
	MACHO_X86_64_RELOC_SUBTRACTOR        = 5
	MACHO_X86_64_RELOC_SIGNED_1          = 6
	MACHO_X86_64_RELOC_SIGNED_2          = 7
	MACHO_X86_64_RELOC_SIGNED_4          = 8
	MACHO_ARM_RELOC_VANILLA              = 0
	MACHO_ARM_RELOC_PAIR                 = 1
	MACHO_ARM_RELOC_SECTDIFF             = 2
	MACHO_ARM_RELOC_BR24                 = 5
	MACHO_ARM64_RELOC_UNSIGNED           = 0
	MACHO_ARM64_RELOC_SUBTRACTOR         = 1
	MACHO_ARM64_RELOC_BRANCH26           = 2
	MACHO_ARM64_RELOC_PAGE21             = 3
	MACHO_ARM64_RELOC_PAGEOFF12          = 4
	MACHO_ARM64_RELOC_GOT_LOAD_PAGE21    = 5
	MACHO_ARM64_RELOC_GOT_LOAD_PAGEOFF12 = 6
	MACHO_ARM64_RELOC_POINTER_TO_GOT     = 7
	MACHO_ARM64_RELOC_ADDEND             = 10
	MACHO_GENERIC_RELOC_VANILLA          = 0
	MACHO_FAKE_GOTPCREL                  = 100
)

const (
	MH_MAGIC    = 0xfeedface
	MH_MAGIC_64 = 0xfeedfacf

	MH_OBJECT  = 0x1
	MH_EXECUTE = 0x2

	MH_NOUNDEFS = 0x1
	MH_DYLDLINK = 0x4
	MH_PIE      = 0x200000
)

const (
	S_REGULAR                  = 0x0
	S_ZEROFILL                 = 0x1
	S_NON_LAZY_SYMBOL_POINTERS = 0x6
	S_SYMBOL_STUBS             = 0x8
	S_MOD_INIT_FUNC_POINTERS   = 0x9
	S_ATTR_PURE_INSTRUCTIONS   = 0x80000000
	S_ATTR_DEBUG               = 0x02000000
	S_ATTR_SOME_INSTRUCTIONS   = 0x00000400
)

const (
	PLATFORM_MACOS       MachoPlatform = 1
	PLATFORM_IOS         MachoPlatform = 2
	PLATFORM_TVOS        MachoPlatform = 3
	PLATFORM_WATCHOS     MachoPlatform = 4
	PLATFORM_BRIDGEOS    MachoPlatform = 5
	PLATFORM_MACCATALYST MachoPlatform = 6
)

// rebase table opcode
const (
	REBASE_TYPE_POINTER         = 1
	REBASE_TYPE_TEXT_ABSOLUTE32 = 2
	REBASE_TYPE_TEXT_PCREL32    = 3

	REBASE_OPCODE_MASK                               = 0xF0
	REBASE_IMMEDIATE_MASK                            = 0x0F
	REBASE_OPCODE_DONE                               = 0x00
	REBASE_OPCODE_SET_TYPE_IMM                       = 0x10
	REBASE_OPCODE_SET_SEGMENT_AND_OFFSET_ULEB        = 0x20
	REBASE_OPCODE_ADD_ADDR_ULEB                      = 0x30
	REBASE_OPCODE_ADD_ADDR_IMM_SCALED                = 0x40
	REBASE_OPCODE_DO_REBASE_IMM_TIMES                = 0x50
	REBASE_OPCODE_DO_REBASE_ULEB_TIMES               = 0x60
	REBASE_OPCODE_DO_REBASE_ADD_ADDR_ULEB            = 0x70
	REBASE_OPCODE_DO_REBASE_ULEB_TIMES_SKIPPING_ULEB = 0x80
)

// bind table opcode
const (
	BIND_TYPE_POINTER         = 1
	BIND_TYPE_TEXT_ABSOLUTE32 = 2
	BIND_TYPE_TEXT_PCREL32    = 3

	BIND_SPECIAL_DYLIB_SELF            = 0
	BIND_SPECIAL_DYLIB_MAIN_EXECUTABLE = -1
	BIND_SPECIAL_DYLIB_FLAT_LOOKUP     = -2
	BIND_SPECIAL_DYLIB_WEAK_LOOKUP     = -3

	BIND_SYMBOL_FLAGS_WEAK_IMPORT = 0x1

	BIND_OPCODE_MASK                                         = 0xF0
	BIND_IMMEDIATE_MASK                                      = 0x0F
	BIND_OPCODE_DONE                                         = 0x00
	BIND_OPCODE_SET_DYLIB_ORDINAL_IMM                        = 0x10
	BIND_OPCODE_SET_DYLIB_ORDINAL_ULEB                       = 0x20
	BIND_OPCODE_SET_DYLIB_SPECIAL_IMM                        = 0x30
	BIND_OPCODE_SET_SYMBOL_TRAILING_FLAGS_IMM                = 0x40
	BIND_OPCODE_SET_TYPE_IMM                                 = 0x50
	BIND_OPCODE_SET_ADDEND_SLEB                              = 0x60
	BIND_OPCODE_SET_SEGMENT_AND_OFFSET_ULEB                  = 0x70
	BIND_OPCODE_ADD_ADDR_ULEB                                = 0x80
	BIND_OPCODE_DO_BIND                                      = 0x90
	BIND_OPCODE_DO_BIND_ADD_ADDR_ULEB                        = 0xA0
	BIND_OPCODE_DO_BIND_ADD_ADDR_IMM_SCALED                  = 0xB0
	BIND_OPCODE_DO_BIND_ULEB_TIMES_SKIPPING_ULEB             = 0xC0
	BIND_OPCODE_THREADED                                     = 0xD0
	BIND_SUBOPCODE_THREADED_SET_BIND_ORDINAL_TABLE_SIZE_ULEB = 0x00
	BIND_SUBOPCODE_THREADED_APPLY                            = 0x01
)

const (
	SymKindLocal = 0 + iota
	SymKindExtdef
	SymKindUndef
	NumSymKind
)

// AddMachoSym adds s to Mach-O symbol table, used in GenSymLate.
// Currently only used on ARM64 when external linking.
func AddMachoSym(ldr *loader.Loader, s loader.Sym)

func MachoAddRebase(s loader.Sym, off int64)

func MachoAddBind(off int64, targ loader.Sym)
