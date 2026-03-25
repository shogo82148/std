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

// IMAGE_LOAD_CONFIG_DIRECTORY64.GuardFlags and IMAGE_LOAD_CONFIG_DIRECTORY32.GuardFlags
// values. These can be combined together.
const (
	IMAGE_GUARD_CF_INSTRUMENTED                    = 0x00000100
	IMAGE_GUARD_CFW_INSTRUMENTED                   = 0x00000200
	IMAGE_GUARD_CF_FUNCTION_TABLE_PRESENT          = 0x00000400
	IMAGE_GUARD_SECURITY_COOKIE_UNUSED             = 0x00000800
	IMAGE_GUARD_PROTECT_DELAYLOAD_IAT              = 0x00001000
	IMAGE_GUARD_DELAYLOAD_IAT_IN_ITS_OWN_SECTION   = 0x00002000
	IMAGE_GUARD_CF_EXPORT_SUPPRESSION_INFO_PRESENT = 0x00004000
	IMAGE_GUARD_CF_ENABLE_EXPORT_SUPPRESSION       = 0x00008000
	IMAGE_GUARD_CF_LONGJUMP_TABLE_PRESENT          = 0x00010000
	IMAGE_GUARD_RF_INSTRUMENTED                    = 0x00020000
	IMAGE_GUARD_RF_ENABLE                          = 0x00040000
	IMAGE_GUARD_RF_STRICT                          = 0x00080000
	IMAGE_GUARD_CF_FUNCTION_TABLE_SIZE_MASK        = 0xF0000000
	IMAGE_GUARD_CF_FUNCTION_TABLE_SIZE_SHIFT       = 28
)

type IMAGE_LOAD_CONFIG_CODE_INTEGRITY struct {
	Flags         uint16
	Catalog       uint16
	CatalogOffset uint32
	Reserved      uint32
}

type IMAGE_LOAD_CONFIG_DIRECTORY32 struct {
	Size                                     uint32
	TimeDateStamp                            uint32
	MajorVersion                             uint16
	MinorVersion                             uint16
	GlobalFlagsClear                         uint32
	GlobalFlagsSet                           uint32
	CriticalSectionDefaultTimeout            uint32
	DeCommitFreeBlockThreshold               uint32
	DeCommitTotalFreeThreshold               uint32
	LockPrefixTable                          uint32
	MaximumAllocationSize                    uint32
	VirtualMemoryThreshold                   uint32
	ProcessHeapFlags                         uint32
	ProcessAffinityMask                      uint32
	CSDVersion                               uint16
	DependentLoadFlags                       uint16
	EditList                                 uint32
	SecurityCookie                           uint32
	SEHandlerTable                           uint32
	SEHandlerCount                           uint32
	GuardCFCheckFunctionPointer              uint32
	GuardCFDispatchFunctionPointer           uint32
	GuardCFFunctionTable                     uint32
	GuardCFFunctionCount                     uint32
	GuardFlags                               uint32
	CodeIntegrity                            IMAGE_LOAD_CONFIG_CODE_INTEGRITY
	GuardAddressTakenIatEntryTable           uint32
	GuardAddressTakenIatEntryCount           uint32
	GuardLongJumpTargetTable                 uint32
	GuardLongJumpTargetCount                 uint32
	DynamicValueRelocTable                   uint32
	CHPEMetadataPointer                      uint32
	GuardRFFailureRoutine                    uint32
	GuardRFFailureRoutineFunctionPointer     uint32
	DynamicValueRelocTableOffset             uint32
	DynamicValueRelocTableSection            uint16
	Reserved2                                uint16
	GuardRFVerifyStackPointerFunctionPointer uint32
	HotPatchTableOffset                      uint32
	Reserved3                                uint32
	EnclaveConfigurationPointer              uint32
	VolatileMetadataPointer                  uint32
	GuardEHContinuationTable                 uint32
	GuardEHContinuationCount                 uint32
	GuardXFGCheckFunctionPointer             uint32
	GuardXFGDispatchFunctionPointer          uint32
	GuardXFGTableDispatchFunctionPointer     uint32
	CastGuardOsDeterminedFailureMode         uint32
	GuardMemcpyFunctionPointer               uint32
}

type IMAGE_LOAD_CONFIG_DIRECTORY64 struct {
	Size                                     uint32
	TimeDateStamp                            uint32
	MajorVersion                             uint16
	MinorVersion                             uint16
	GlobalFlagsClear                         uint32
	GlobalFlagsSet                           uint32
	CriticalSectionDefaultTimeout            uint32
	DeCommitFreeBlockThreshold               uint64
	DeCommitTotalFreeThreshold               uint64
	LockPrefixTable                          uint64
	MaximumAllocationSize                    uint64
	VirtualMemoryThreshold                   uint64
	ProcessAffinityMask                      uint64
	ProcessHeapFlags                         uint32
	CSDVersion                               uint16
	DependentLoadFlags                       uint16
	EditList                                 uint64
	SecurityCookie                           uint64
	SEHandlerTable                           uint64
	SEHandlerCount                           uint64
	GuardCFCheckFunctionPointer              uint64
	GuardCFDispatchFunctionPointer           uint64
	GuardCFFunctionTable                     uint64
	GuardCFFunctionCount                     uint64
	GuardFlags                               uint32
	CodeIntegrity                            IMAGE_LOAD_CONFIG_CODE_INTEGRITY
	GuardAddressTakenIatEntryTable           uint64
	GuardAddressTakenIatEntryCount           uint64
	GuardLongJumpTargetTable                 uint64
	GuardLongJumpTargetCount                 uint64
	DynamicValueRelocTable                   uint64
	CHPEMetadataPointer                      uint64
	GuardRFFailureRoutine                    uint64
	GuardRFFailureRoutineFunctionPointer     uint64
	DynamicValueRelocTableOffset             uint32
	DynamicValueRelocTableSection            uint16
	Reserved2                                uint16
	GuardRFVerifyStackPointerFunctionPointer uint64
	HotPatchTableOffset                      uint32
	Reserved3                                uint32
	EnclaveConfigurationPointer              uint64
	VolatileMetadataPointer                  uint64
	GuardEHContinuationTable                 uint64
	GuardEHContinuationCount                 uint64
	GuardXFGCheckFunctionPointer             uint64
	GuardXFGDispatchFunctionPointer          uint64
	GuardXFGTableDispatchFunctionPointer     uint64
	CastGuardOsDeterminedFailureMode         uint64
	GuardMemcpyFunctionPointer               uint64
}

const (
	PeMinimumTargetMajorVersion = 10
	PeMinimumTargetMinorVersion = 0
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
