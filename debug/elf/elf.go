/*
 * ELF constants and data structures
 *
 * Derived from:
 * $FreeBSD: src/sys/sys/elf32.h,v 1.8.14.1 2005/12/30 22:13:58 marcel Exp $
 * $FreeBSD: src/sys/sys/elf64.h,v 1.10.14.1 2005/12/30 22:13:58 marcel Exp $
 * $FreeBSD: src/sys/sys/elf_common.h,v 1.15.8.1 2005/12/30 22:13:58 marcel Exp $
 * $FreeBSD: src/sys/alpha/include/elf.h,v 1.14 2003/09/25 01:10:22 peter Exp $
 * $FreeBSD: src/sys/amd64/include/elf.h,v 1.18 2004/08/03 08:21:48 dfr Exp $
 * $FreeBSD: src/sys/arm/include/elf.h,v 1.5.2.1 2006/06/30 21:42:52 cognet Exp $
 * $FreeBSD: src/sys/i386/include/elf.h,v 1.16 2004/08/02 19:12:17 dfr Exp $
 * $FreeBSD: src/sys/powerpc/include/elf.h,v 1.7 2004/11/02 09:47:01 ssouhlal Exp $
 * $FreeBSD: src/sys/sparc64/include/elf.h,v 1.12 2003/09/25 01:10:26 peter Exp $
 * "ELF for the ARM® 64-bit Architecture (AArch64)" (ARM IHI 0056B)
 *
 * Copyright (c) 1996-1998 John D. Polstra.  All rights reserved.
 * Copyright (c) 2001 David E. O'Brien
 * Portions Copyright 2009 The Go Authors.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

package elf

// Indexes into the Header.Ident array.
const (
	EI_CLASS      = 4
	EI_DATA       = 5
	EI_VERSION    = 6
	EI_OSABI      = 7
	EI_ABIVERSION = 8
	EI_PAD        = 9
	EI_NIDENT     = 16
)

// Initial magic number for ELF files.
const ELFMAG = "\177ELF"

// Version is found in Header.Ident[EI_VERSION] and Header.Version.
type Version byte

const (
	EV_NONE    Version = 0
	EV_CURRENT Version = 1
)

func (i Version) String() string
func (i Version) GoString() string

// Class is found in Header.Ident[EI_CLASS] and Header.Class.
type Class byte

const (
	ELFCLASSNONE Class = 0
	ELFCLASS32   Class = 1
	ELFCLASS64   Class = 2
)

func (i Class) String() string
func (i Class) GoString() string

// Data is found in Header.Ident[EI_DATA] and Header.Data.
type Data byte

const (
	ELFDATANONE Data = 0
	ELFDATA2LSB Data = 1
	ELFDATA2MSB Data = 2
)

func (i Data) String() string
func (i Data) GoString() string

// OSABI is found in Header.Ident[EI_OSABI] and Header.OSABI.
type OSABI byte

const (
	ELFOSABI_NONE       OSABI = 0
	ELFOSABI_HPUX       OSABI = 1
	ELFOSABI_NETBSD     OSABI = 2
	ELFOSABI_LINUX      OSABI = 3
	ELFOSABI_HURD       OSABI = 4
	ELFOSABI_86OPEN     OSABI = 5
	ELFOSABI_SOLARIS    OSABI = 6
	ELFOSABI_AIX        OSABI = 7
	ELFOSABI_IRIX       OSABI = 8
	ELFOSABI_FREEBSD    OSABI = 9
	ELFOSABI_TRU64      OSABI = 10
	ELFOSABI_MODESTO    OSABI = 11
	ELFOSABI_OPENBSD    OSABI = 12
	ELFOSABI_OPENVMS    OSABI = 13
	ELFOSABI_NSK        OSABI = 14
	ELFOSABI_ARM        OSABI = 97
	ELFOSABI_STANDALONE OSABI = 255
)

func (i OSABI) String() string
func (i OSABI) GoString() string

// Type is found in Header.Type.
type Type uint16

const (
	ET_NONE   Type = 0
	ET_REL    Type = 1
	ET_EXEC   Type = 2
	ET_DYN    Type = 3
	ET_CORE   Type = 4
	ET_LOOS   Type = 0xfe00
	ET_HIOS   Type = 0xfeff
	ET_LOPROC Type = 0xff00
	ET_HIPROC Type = 0xffff
)

func (i Type) String() string
func (i Type) GoString() string

// Machine is found in Header.Machine.
type Machine uint16

const (
	EM_NONE        Machine = 0
	EM_M32         Machine = 1
	EM_SPARC       Machine = 2
	EM_386         Machine = 3
	EM_68K         Machine = 4
	EM_88K         Machine = 5
	EM_860         Machine = 7
	EM_MIPS        Machine = 8
	EM_S370        Machine = 9
	EM_MIPS_RS3_LE Machine = 10
	EM_PARISC      Machine = 15
	EM_VPP500      Machine = 17
	EM_SPARC32PLUS Machine = 18
	EM_960         Machine = 19
	EM_PPC         Machine = 20
	EM_PPC64       Machine = 21
	EM_S390        Machine = 22
	EM_V800        Machine = 36
	EM_FR20        Machine = 37
	EM_RH32        Machine = 38
	EM_RCE         Machine = 39
	EM_ARM         Machine = 40
	EM_SH          Machine = 42
	EM_SPARCV9     Machine = 43
	EM_TRICORE     Machine = 44
	EM_ARC         Machine = 45
	EM_H8_300      Machine = 46
	EM_H8_300H     Machine = 47
	EM_H8S         Machine = 48
	EM_H8_500      Machine = 49
	EM_IA_64       Machine = 50
	EM_MIPS_X      Machine = 51
	EM_COLDFIRE    Machine = 52
	EM_68HC12      Machine = 53
	EM_MMA         Machine = 54
	EM_PCP         Machine = 55
	EM_NCPU        Machine = 56
	EM_NDR1        Machine = 57
	EM_STARCORE    Machine = 58
	EM_ME16        Machine = 59
	EM_ST100       Machine = 60
	EM_TINYJ       Machine = 61
	EM_X86_64      Machine = 62
	EM_AARCH64     Machine = 183

	/* Non-standard or deprecated. */
	EM_486         Machine = 6
	EM_MIPS_RS4_BE Machine = 10
	EM_ALPHA_STD   Machine = 41
	EM_ALPHA       Machine = 0x9026
)

func (i Machine) String() string
func (i Machine) GoString() string

// Special section indices.
type SectionIndex int

const (
	SHN_UNDEF     SectionIndex = 0
	SHN_LORESERVE SectionIndex = 0xff00
	SHN_LOPROC    SectionIndex = 0xff00
	SHN_HIPROC    SectionIndex = 0xff1f
	SHN_LOOS      SectionIndex = 0xff20
	SHN_HIOS      SectionIndex = 0xff3f
	SHN_ABS       SectionIndex = 0xfff1
	SHN_COMMON    SectionIndex = 0xfff2
	SHN_XINDEX    SectionIndex = 0xffff
	SHN_HIRESERVE SectionIndex = 0xffff
)

func (i SectionIndex) String() string
func (i SectionIndex) GoString() string

// Section type.
type SectionType uint32

const (
	SHT_NULL           SectionType = 0
	SHT_PROGBITS       SectionType = 1
	SHT_SYMTAB         SectionType = 2
	SHT_STRTAB         SectionType = 3
	SHT_RELA           SectionType = 4
	SHT_HASH           SectionType = 5
	SHT_DYNAMIC        SectionType = 6
	SHT_NOTE           SectionType = 7
	SHT_NOBITS         SectionType = 8
	SHT_REL            SectionType = 9
	SHT_SHLIB          SectionType = 10
	SHT_DYNSYM         SectionType = 11
	SHT_INIT_ARRAY     SectionType = 14
	SHT_FINI_ARRAY     SectionType = 15
	SHT_PREINIT_ARRAY  SectionType = 16
	SHT_GROUP          SectionType = 17
	SHT_SYMTAB_SHNDX   SectionType = 18
	SHT_LOOS           SectionType = 0x60000000
	SHT_GNU_ATTRIBUTES SectionType = 0x6ffffff5
	SHT_GNU_HASH       SectionType = 0x6ffffff6
	SHT_GNU_LIBLIST    SectionType = 0x6ffffff7
	SHT_GNU_VERDEF     SectionType = 0x6ffffffd
	SHT_GNU_VERNEED    SectionType = 0x6ffffffe
	SHT_GNU_VERSYM     SectionType = 0x6fffffff
	SHT_HIOS           SectionType = 0x6fffffff
	SHT_LOPROC         SectionType = 0x70000000
	SHT_HIPROC         SectionType = 0x7fffffff
	SHT_LOUSER         SectionType = 0x80000000
	SHT_HIUSER         SectionType = 0xffffffff
)

func (i SectionType) String() string
func (i SectionType) GoString() string

// Section flags.
type SectionFlag uint32

const (
	SHF_WRITE            SectionFlag = 0x1
	SHF_ALLOC            SectionFlag = 0x2
	SHF_EXECINSTR        SectionFlag = 0x4
	SHF_MERGE            SectionFlag = 0x10
	SHF_STRINGS          SectionFlag = 0x20
	SHF_INFO_LINK        SectionFlag = 0x40
	SHF_LINK_ORDER       SectionFlag = 0x80
	SHF_OS_NONCONFORMING SectionFlag = 0x100
	SHF_GROUP            SectionFlag = 0x200
	SHF_TLS              SectionFlag = 0x400
	SHF_COMPRESSED       SectionFlag = 0x800
	SHF_MASKOS           SectionFlag = 0x0ff00000
	SHF_MASKPROC         SectionFlag = 0xf0000000
)

func (i SectionFlag) String() string
func (i SectionFlag) GoString() string

// Section compression type.
type CompressionType int

const (
	COMPRESS_ZLIB   CompressionType = 1
	COMPRESS_LOOS   CompressionType = 0x60000000
	COMPRESS_HIOS   CompressionType = 0x6fffffff
	COMPRESS_LOPROC CompressionType = 0x70000000
	COMPRESS_HIPROC CompressionType = 0x7fffffff
)

func (i CompressionType) String() string
func (i CompressionType) GoString() string

// Prog.Type
type ProgType int

const (
	PT_NULL    ProgType = 0
	PT_LOAD    ProgType = 1
	PT_DYNAMIC ProgType = 2
	PT_INTERP  ProgType = 3
	PT_NOTE    ProgType = 4
	PT_SHLIB   ProgType = 5
	PT_PHDR    ProgType = 6
	PT_TLS     ProgType = 7
	PT_LOOS    ProgType = 0x60000000
	PT_HIOS    ProgType = 0x6fffffff
	PT_LOPROC  ProgType = 0x70000000
	PT_HIPROC  ProgType = 0x7fffffff
)

func (i ProgType) String() string
func (i ProgType) GoString() string

// Prog.Flag
type ProgFlag uint32

const (
	PF_X        ProgFlag = 0x1
	PF_W        ProgFlag = 0x2
	PF_R        ProgFlag = 0x4
	PF_MASKOS   ProgFlag = 0x0ff00000
	PF_MASKPROC ProgFlag = 0xf0000000
)

func (i ProgFlag) String() string
func (i ProgFlag) GoString() string

// Dyn.Tag
type DynTag int

const (
	DT_NULL         DynTag = 0
	DT_NEEDED       DynTag = 1
	DT_PLTRELSZ     DynTag = 2
	DT_PLTGOT       DynTag = 3
	DT_HASH         DynTag = 4
	DT_STRTAB       DynTag = 5
	DT_SYMTAB       DynTag = 6
	DT_RELA         DynTag = 7
	DT_RELASZ       DynTag = 8
	DT_RELAENT      DynTag = 9
	DT_STRSZ        DynTag = 10
	DT_SYMENT       DynTag = 11
	DT_INIT         DynTag = 12
	DT_FINI         DynTag = 13
	DT_SONAME       DynTag = 14
	DT_RPATH        DynTag = 15
	DT_SYMBOLIC     DynTag = 16
	DT_REL          DynTag = 17
	DT_RELSZ        DynTag = 18
	DT_RELENT       DynTag = 19
	DT_PLTREL       DynTag = 20
	DT_DEBUG        DynTag = 21
	DT_TEXTREL      DynTag = 22
	DT_JMPREL       DynTag = 23
	DT_BIND_NOW     DynTag = 24
	DT_INIT_ARRAY   DynTag = 25
	DT_FINI_ARRAY   DynTag = 26
	DT_INIT_ARRAYSZ DynTag = 27
	DT_FINI_ARRAYSZ DynTag = 28
	DT_RUNPATH      DynTag = 29
	DT_FLAGS        DynTag = 30
	DT_ENCODING     DynTag = 32

	DT_PREINIT_ARRAY   DynTag = 32
	DT_PREINIT_ARRAYSZ DynTag = 33
	DT_LOOS            DynTag = 0x6000000d
	DT_HIOS            DynTag = 0x6ffff000
	DT_VERSYM          DynTag = 0x6ffffff0
	DT_VERNEED         DynTag = 0x6ffffffe
	DT_VERNEEDNUM      DynTag = 0x6fffffff
	DT_LOPROC          DynTag = 0x70000000
	DT_HIPROC          DynTag = 0x7fffffff
)

func (i DynTag) String() string
func (i DynTag) GoString() string

// DT_FLAGS values.
type DynFlag int

const (
	DF_ORIGIN DynFlag = 0x0001

	DF_SYMBOLIC DynFlag = 0x0002
	DF_TEXTREL  DynFlag = 0x0004
	DF_BIND_NOW DynFlag = 0x0008

	DF_STATIC_TLS DynFlag = 0x0010
)

func (i DynFlag) String() string
func (i DynFlag) GoString() string

// NType values; used in core files.
type NType int

const (
	NT_PRSTATUS NType = 1
	NT_FPREGSET NType = 2
	NT_PRPSINFO NType = 3
)

func (i NType) String() string
func (i NType) GoString() string

/* Symbol Binding - ELFNN_ST_BIND - st_info */
type SymBind int

const (
	STB_LOCAL  SymBind = 0
	STB_GLOBAL SymBind = 1
	STB_WEAK   SymBind = 2
	STB_LOOS   SymBind = 10
	STB_HIOS   SymBind = 12
	STB_LOPROC SymBind = 13
	STB_HIPROC SymBind = 15
)

func (i SymBind) String() string
func (i SymBind) GoString() string

/* Symbol type - ELFNN_ST_TYPE - st_info */
type SymType int

const (
	STT_NOTYPE  SymType = 0
	STT_OBJECT  SymType = 1
	STT_FUNC    SymType = 2
	STT_SECTION SymType = 3
	STT_FILE    SymType = 4
	STT_COMMON  SymType = 5
	STT_TLS     SymType = 6
	STT_LOOS    SymType = 10
	STT_HIOS    SymType = 12
	STT_LOPROC  SymType = 13
	STT_HIPROC  SymType = 15
)

func (i SymType) String() string
func (i SymType) GoString() string

/* Symbol visibility - ELFNN_ST_VISIBILITY - st_other */
type SymVis int

const (
	STV_DEFAULT   SymVis = 0x0
	STV_INTERNAL  SymVis = 0x1
	STV_HIDDEN    SymVis = 0x2
	STV_PROTECTED SymVis = 0x3
)

func (i SymVis) String() string
func (i SymVis) GoString() string

// Relocation types for x86-64.
type R_X86_64 int

const (
	R_X86_64_NONE     R_X86_64 = 0
	R_X86_64_64       R_X86_64 = 1
	R_X86_64_PC32     R_X86_64 = 2
	R_X86_64_GOT32    R_X86_64 = 3
	R_X86_64_PLT32    R_X86_64 = 4
	R_X86_64_COPY     R_X86_64 = 5
	R_X86_64_GLOB_DAT R_X86_64 = 6
	R_X86_64_JMP_SLOT R_X86_64 = 7
	R_X86_64_RELATIVE R_X86_64 = 8
	R_X86_64_GOTPCREL R_X86_64 = 9
	R_X86_64_32       R_X86_64 = 10
	R_X86_64_32S      R_X86_64 = 11
	R_X86_64_16       R_X86_64 = 12
	R_X86_64_PC16     R_X86_64 = 13
	R_X86_64_8        R_X86_64 = 14
	R_X86_64_PC8      R_X86_64 = 15
	R_X86_64_DTPMOD64 R_X86_64 = 16
	R_X86_64_DTPOFF64 R_X86_64 = 17
	R_X86_64_TPOFF64  R_X86_64 = 18
	R_X86_64_TLSGD    R_X86_64 = 19
	R_X86_64_TLSLD    R_X86_64 = 20
	R_X86_64_DTPOFF32 R_X86_64 = 21
	R_X86_64_GOTTPOFF R_X86_64 = 22
	R_X86_64_TPOFF32  R_X86_64 = 23
)

func (i R_X86_64) String() string
func (i R_X86_64) GoString() string

// Relocation types for AArch64 (aka arm64)
type R_AARCH64 int

const (
	R_AARCH64_NONE                            R_AARCH64 = 0
	R_AARCH64_P32_ABS32                       R_AARCH64 = 1
	R_AARCH64_P32_ABS16                       R_AARCH64 = 2
	R_AARCH64_P32_PREL32                      R_AARCH64 = 3
	R_AARCH64_P32_PREL16                      R_AARCH64 = 4
	R_AARCH64_P32_MOVW_UABS_G0                R_AARCH64 = 5
	R_AARCH64_P32_MOVW_UABS_G0_NC             R_AARCH64 = 6
	R_AARCH64_P32_MOVW_UABS_G1                R_AARCH64 = 7
	R_AARCH64_P32_MOVW_SABS_G0                R_AARCH64 = 8
	R_AARCH64_P32_LD_PREL_LO19                R_AARCH64 = 9
	R_AARCH64_P32_ADR_PREL_LO21               R_AARCH64 = 10
	R_AARCH64_P32_ADR_PREL_PG_HI21            R_AARCH64 = 11
	R_AARCH64_P32_ADD_ABS_LO12_NC             R_AARCH64 = 12
	R_AARCH64_P32_LDST8_ABS_LO12_NC           R_AARCH64 = 13
	R_AARCH64_P32_LDST16_ABS_LO12_NC          R_AARCH64 = 14
	R_AARCH64_P32_LDST32_ABS_LO12_NC          R_AARCH64 = 15
	R_AARCH64_P32_LDST64_ABS_LO12_NC          R_AARCH64 = 16
	R_AARCH64_P32_LDST128_ABS_LO12_NC         R_AARCH64 = 17
	R_AARCH64_P32_TSTBR14                     R_AARCH64 = 18
	R_AARCH64_P32_CONDBR19                    R_AARCH64 = 19
	R_AARCH64_P32_JUMP26                      R_AARCH64 = 20
	R_AARCH64_P32_CALL26                      R_AARCH64 = 21
	R_AARCH64_P32_GOT_LD_PREL19               R_AARCH64 = 25
	R_AARCH64_P32_ADR_GOT_PAGE                R_AARCH64 = 26
	R_AARCH64_P32_LD32_GOT_LO12_NC            R_AARCH64 = 27
	R_AARCH64_P32_TLSGD_ADR_PAGE21            R_AARCH64 = 81
	R_AARCH64_P32_TLSGD_ADD_LO12_NC           R_AARCH64 = 82
	R_AARCH64_P32_TLSIE_ADR_GOTTPREL_PAGE21   R_AARCH64 = 103
	R_AARCH64_P32_TLSIE_LD32_GOTTPREL_LO12_NC R_AARCH64 = 104
	R_AARCH64_P32_TLSIE_LD_GOTTPREL_PREL19    R_AARCH64 = 105
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G1         R_AARCH64 = 106
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G0         R_AARCH64 = 107
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G0_NC      R_AARCH64 = 108
	R_AARCH64_P32_TLSLE_ADD_TPREL_HI12        R_AARCH64 = 109
	R_AARCH64_P32_TLSLE_ADD_TPREL_LO12        R_AARCH64 = 110
	R_AARCH64_P32_TLSLE_ADD_TPREL_LO12_NC     R_AARCH64 = 111
	R_AARCH64_P32_TLSDESC_LD_PREL19           R_AARCH64 = 122
	R_AARCH64_P32_TLSDESC_ADR_PREL21          R_AARCH64 = 123
	R_AARCH64_P32_TLSDESC_ADR_PAGE21          R_AARCH64 = 124
	R_AARCH64_P32_TLSDESC_LD32_LO12_NC        R_AARCH64 = 125
	R_AARCH64_P32_TLSDESC_ADD_LO12_NC         R_AARCH64 = 126
	R_AARCH64_P32_TLSDESC_CALL                R_AARCH64 = 127
	R_AARCH64_P32_COPY                        R_AARCH64 = 180
	R_AARCH64_P32_GLOB_DAT                    R_AARCH64 = 181
	R_AARCH64_P32_JUMP_SLOT                   R_AARCH64 = 182
	R_AARCH64_P32_RELATIVE                    R_AARCH64 = 183
	R_AARCH64_P32_TLS_DTPMOD                  R_AARCH64 = 184
	R_AARCH64_P32_TLS_DTPREL                  R_AARCH64 = 185
	R_AARCH64_P32_TLS_TPREL                   R_AARCH64 = 186
	R_AARCH64_P32_TLSDESC                     R_AARCH64 = 187
	R_AARCH64_P32_IRELATIVE                   R_AARCH64 = 188
	R_AARCH64_NULL                            R_AARCH64 = 256
	R_AARCH64_ABS64                           R_AARCH64 = 257
	R_AARCH64_ABS32                           R_AARCH64 = 258
	R_AARCH64_ABS16                           R_AARCH64 = 259
	R_AARCH64_PREL64                          R_AARCH64 = 260
	R_AARCH64_PREL32                          R_AARCH64 = 261
	R_AARCH64_PREL16                          R_AARCH64 = 262
	R_AARCH64_MOVW_UABS_G0                    R_AARCH64 = 263
	R_AARCH64_MOVW_UABS_G0_NC                 R_AARCH64 = 264
	R_AARCH64_MOVW_UABS_G1                    R_AARCH64 = 265
	R_AARCH64_MOVW_UABS_G1_NC                 R_AARCH64 = 266
	R_AARCH64_MOVW_UABS_G2                    R_AARCH64 = 267
	R_AARCH64_MOVW_UABS_G2_NC                 R_AARCH64 = 268
	R_AARCH64_MOVW_UABS_G3                    R_AARCH64 = 269
	R_AARCH64_MOVW_SABS_G0                    R_AARCH64 = 270
	R_AARCH64_MOVW_SABS_G1                    R_AARCH64 = 271
	R_AARCH64_MOVW_SABS_G2                    R_AARCH64 = 272
	R_AARCH64_LD_PREL_LO19                    R_AARCH64 = 273
	R_AARCH64_ADR_PREL_LO21                   R_AARCH64 = 274
	R_AARCH64_ADR_PREL_PG_HI21                R_AARCH64 = 275
	R_AARCH64_ADR_PREL_PG_HI21_NC             R_AARCH64 = 276
	R_AARCH64_ADD_ABS_LO12_NC                 R_AARCH64 = 277
	R_AARCH64_LDST8_ABS_LO12_NC               R_AARCH64 = 278
	R_AARCH64_TSTBR14                         R_AARCH64 = 279
	R_AARCH64_CONDBR19                        R_AARCH64 = 280
	R_AARCH64_JUMP26                          R_AARCH64 = 282
	R_AARCH64_CALL26                          R_AARCH64 = 283
	R_AARCH64_LDST16_ABS_LO12_NC              R_AARCH64 = 284
	R_AARCH64_LDST32_ABS_LO12_NC              R_AARCH64 = 285
	R_AARCH64_LDST64_ABS_LO12_NC              R_AARCH64 = 286
	R_AARCH64_LDST128_ABS_LO12_NC             R_AARCH64 = 299
	R_AARCH64_GOT_LD_PREL19                   R_AARCH64 = 309
	R_AARCH64_ADR_GOT_PAGE                    R_AARCH64 = 311
	R_AARCH64_LD64_GOT_LO12_NC                R_AARCH64 = 312
	R_AARCH64_TLSGD_ADR_PAGE21                R_AARCH64 = 513
	R_AARCH64_TLSGD_ADD_LO12_NC               R_AARCH64 = 514
	R_AARCH64_TLSIE_MOVW_GOTTPREL_G1          R_AARCH64 = 539
	R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC       R_AARCH64 = 540
	R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21       R_AARCH64 = 541
	R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC     R_AARCH64 = 542
	R_AARCH64_TLSIE_LD_GOTTPREL_PREL19        R_AARCH64 = 543
	R_AARCH64_TLSLE_MOVW_TPREL_G2             R_AARCH64 = 544
	R_AARCH64_TLSLE_MOVW_TPREL_G1             R_AARCH64 = 545
	R_AARCH64_TLSLE_MOVW_TPREL_G1_NC          R_AARCH64 = 546
	R_AARCH64_TLSLE_MOVW_TPREL_G0             R_AARCH64 = 547
	R_AARCH64_TLSLE_MOVW_TPREL_G0_NC          R_AARCH64 = 548
	R_AARCH64_TLSLE_ADD_TPREL_HI12            R_AARCH64 = 549
	R_AARCH64_TLSLE_ADD_TPREL_LO12            R_AARCH64 = 550
	R_AARCH64_TLSLE_ADD_TPREL_LO12_NC         R_AARCH64 = 551
	R_AARCH64_TLSDESC_LD_PREL19               R_AARCH64 = 560
	R_AARCH64_TLSDESC_ADR_PREL21              R_AARCH64 = 561
	R_AARCH64_TLSDESC_ADR_PAGE21              R_AARCH64 = 562
	R_AARCH64_TLSDESC_LD64_LO12_NC            R_AARCH64 = 563
	R_AARCH64_TLSDESC_ADD_LO12_NC             R_AARCH64 = 564
	R_AARCH64_TLSDESC_OFF_G1                  R_AARCH64 = 565
	R_AARCH64_TLSDESC_OFF_G0_NC               R_AARCH64 = 566
	R_AARCH64_TLSDESC_LDR                     R_AARCH64 = 567
	R_AARCH64_TLSDESC_ADD                     R_AARCH64 = 568
	R_AARCH64_TLSDESC_CALL                    R_AARCH64 = 569
	R_AARCH64_COPY                            R_AARCH64 = 1024
	R_AARCH64_GLOB_DAT                        R_AARCH64 = 1025
	R_AARCH64_JUMP_SLOT                       R_AARCH64 = 1026
	R_AARCH64_RELATIVE                        R_AARCH64 = 1027
	R_AARCH64_TLS_DTPMOD64                    R_AARCH64 = 1028
	R_AARCH64_TLS_DTPREL64                    R_AARCH64 = 1029
	R_AARCH64_TLS_TPREL64                     R_AARCH64 = 1030
	R_AARCH64_TLSDESC                         R_AARCH64 = 1031
	R_AARCH64_IRELATIVE                       R_AARCH64 = 1032
)

func (i R_AARCH64) String() string
func (i R_AARCH64) GoString() string

// Relocation types for Alpha.
type R_ALPHA int

const (
	R_ALPHA_NONE           R_ALPHA = 0
	R_ALPHA_REFLONG        R_ALPHA = 1
	R_ALPHA_REFQUAD        R_ALPHA = 2
	R_ALPHA_GPREL32        R_ALPHA = 3
	R_ALPHA_LITERAL        R_ALPHA = 4
	R_ALPHA_LITUSE         R_ALPHA = 5
	R_ALPHA_GPDISP         R_ALPHA = 6
	R_ALPHA_BRADDR         R_ALPHA = 7
	R_ALPHA_HINT           R_ALPHA = 8
	R_ALPHA_SREL16         R_ALPHA = 9
	R_ALPHA_SREL32         R_ALPHA = 10
	R_ALPHA_SREL64         R_ALPHA = 11
	R_ALPHA_OP_PUSH        R_ALPHA = 12
	R_ALPHA_OP_STORE       R_ALPHA = 13
	R_ALPHA_OP_PSUB        R_ALPHA = 14
	R_ALPHA_OP_PRSHIFT     R_ALPHA = 15
	R_ALPHA_GPVALUE        R_ALPHA = 16
	R_ALPHA_GPRELHIGH      R_ALPHA = 17
	R_ALPHA_GPRELLOW       R_ALPHA = 18
	R_ALPHA_IMMED_GP_16    R_ALPHA = 19
	R_ALPHA_IMMED_GP_HI32  R_ALPHA = 20
	R_ALPHA_IMMED_SCN_HI32 R_ALPHA = 21
	R_ALPHA_IMMED_BR_HI32  R_ALPHA = 22
	R_ALPHA_IMMED_LO32     R_ALPHA = 23
	R_ALPHA_COPY           R_ALPHA = 24
	R_ALPHA_GLOB_DAT       R_ALPHA = 25
	R_ALPHA_JMP_SLOT       R_ALPHA = 26
	R_ALPHA_RELATIVE       R_ALPHA = 27
)

func (i R_ALPHA) String() string
func (i R_ALPHA) GoString() string

// Relocation types for ARM.
type R_ARM int

const (
	R_ARM_NONE          R_ARM = 0
	R_ARM_PC24          R_ARM = 1
	R_ARM_ABS32         R_ARM = 2
	R_ARM_REL32         R_ARM = 3
	R_ARM_PC13          R_ARM = 4
	R_ARM_ABS16         R_ARM = 5
	R_ARM_ABS12         R_ARM = 6
	R_ARM_THM_ABS5      R_ARM = 7
	R_ARM_ABS8          R_ARM = 8
	R_ARM_SBREL32       R_ARM = 9
	R_ARM_THM_PC22      R_ARM = 10
	R_ARM_THM_PC8       R_ARM = 11
	R_ARM_AMP_VCALL9    R_ARM = 12
	R_ARM_SWI24         R_ARM = 13
	R_ARM_THM_SWI8      R_ARM = 14
	R_ARM_XPC25         R_ARM = 15
	R_ARM_THM_XPC22     R_ARM = 16
	R_ARM_COPY          R_ARM = 20
	R_ARM_GLOB_DAT      R_ARM = 21
	R_ARM_JUMP_SLOT     R_ARM = 22
	R_ARM_RELATIVE      R_ARM = 23
	R_ARM_GOTOFF        R_ARM = 24
	R_ARM_GOTPC         R_ARM = 25
	R_ARM_GOT32         R_ARM = 26
	R_ARM_PLT32         R_ARM = 27
	R_ARM_GNU_VTENTRY   R_ARM = 100
	R_ARM_GNU_VTINHERIT R_ARM = 101
	R_ARM_RSBREL32      R_ARM = 250
	R_ARM_THM_RPC22     R_ARM = 251
	R_ARM_RREL32        R_ARM = 252
	R_ARM_RABS32        R_ARM = 253
	R_ARM_RPC24         R_ARM = 254
	R_ARM_RBASE         R_ARM = 255
)

func (i R_ARM) String() string
func (i R_ARM) GoString() string

// Relocation types for 386.
type R_386 int

const (
	R_386_NONE         R_386 = 0
	R_386_32           R_386 = 1
	R_386_PC32         R_386 = 2
	R_386_GOT32        R_386 = 3
	R_386_PLT32        R_386 = 4
	R_386_COPY         R_386 = 5
	R_386_GLOB_DAT     R_386 = 6
	R_386_JMP_SLOT     R_386 = 7
	R_386_RELATIVE     R_386 = 8
	R_386_GOTOFF       R_386 = 9
	R_386_GOTPC        R_386 = 10
	R_386_TLS_TPOFF    R_386 = 14
	R_386_TLS_IE       R_386 = 15
	R_386_TLS_GOTIE    R_386 = 16
	R_386_TLS_LE       R_386 = 17
	R_386_TLS_GD       R_386 = 18
	R_386_TLS_LDM      R_386 = 19
	R_386_TLS_GD_32    R_386 = 24
	R_386_TLS_GD_PUSH  R_386 = 25
	R_386_TLS_GD_CALL  R_386 = 26
	R_386_TLS_GD_POP   R_386 = 27
	R_386_TLS_LDM_32   R_386 = 28
	R_386_TLS_LDM_PUSH R_386 = 29
	R_386_TLS_LDM_CALL R_386 = 30
	R_386_TLS_LDM_POP  R_386 = 31
	R_386_TLS_LDO_32   R_386 = 32
	R_386_TLS_IE_32    R_386 = 33
	R_386_TLS_LE_32    R_386 = 34
	R_386_TLS_DTPMOD32 R_386 = 35
	R_386_TLS_DTPOFF32 R_386 = 36
	R_386_TLS_TPOFF32  R_386 = 37
)

func (i R_386) String() string
func (i R_386) GoString() string

// Relocation types for MIPS.
type R_MIPS int

const (
	R_MIPS_NONE          R_MIPS = 0
	R_MIPS_16            R_MIPS = 1
	R_MIPS_32            R_MIPS = 2
	R_MIPS_REL32         R_MIPS = 3
	R_MIPS_26            R_MIPS = 4
	R_MIPS_HI16          R_MIPS = 5
	R_MIPS_LO16          R_MIPS = 6
	R_MIPS_GPREL16       R_MIPS = 7
	R_MIPS_LITERAL       R_MIPS = 8
	R_MIPS_GOT16         R_MIPS = 9
	R_MIPS_PC16          R_MIPS = 10
	R_MIPS_CALL16        R_MIPS = 11
	R_MIPS_GPREL32       R_MIPS = 12
	R_MIPS_SHIFT5        R_MIPS = 16
	R_MIPS_SHIFT6        R_MIPS = 17
	R_MIPS_64            R_MIPS = 18
	R_MIPS_GOT_DISP      R_MIPS = 19
	R_MIPS_GOT_PAGE      R_MIPS = 20
	R_MIPS_GOT_OFST      R_MIPS = 21
	R_MIPS_GOT_HI16      R_MIPS = 22
	R_MIPS_GOT_LO16      R_MIPS = 23
	R_MIPS_SUB           R_MIPS = 24
	R_MIPS_INSERT_A      R_MIPS = 25
	R_MIPS_INSERT_B      R_MIPS = 26
	R_MIPS_DELETE        R_MIPS = 27
	R_MIPS_HIGHER        R_MIPS = 28
	R_MIPS_HIGHEST       R_MIPS = 29
	R_MIPS_CALL_HI16     R_MIPS = 30
	R_MIPS_CALL_LO16     R_MIPS = 31
	R_MIPS_SCN_DISP      R_MIPS = 32
	R_MIPS_REL16         R_MIPS = 33
	R_MIPS_ADD_IMMEDIATE R_MIPS = 34
	R_MIPS_PJUMP         R_MIPS = 35
	R_MIPS_RELGOT        R_MIPS = 36
	R_MIPS_JALR          R_MIPS = 37

	R_MIPS_TLS_DTPMOD32    R_MIPS = 38
	R_MIPS_TLS_DTPREL32    R_MIPS = 39
	R_MIPS_TLS_DTPMOD64    R_MIPS = 40
	R_MIPS_TLS_DTPREL64    R_MIPS = 41
	R_MIPS_TLS_GD          R_MIPS = 42
	R_MIPS_TLS_LDM         R_MIPS = 43
	R_MIPS_TLS_DTPREL_HI16 R_MIPS = 44
	R_MIPS_TLS_DTPREL_LO16 R_MIPS = 45
	R_MIPS_TLS_GOTTPREL    R_MIPS = 46
	R_MIPS_TLS_TPREL32     R_MIPS = 47
	R_MIPS_TLS_TPREL64     R_MIPS = 48
	R_MIPS_TLS_TPREL_HI16  R_MIPS = 49
	R_MIPS_TLS_TPREL_LO16  R_MIPS = 50
)

func (i R_MIPS) String() string
func (i R_MIPS) GoString() string

// Relocation types for PowerPC.
type R_PPC int

const (
	R_PPC_NONE            R_PPC = 0
	R_PPC_ADDR32          R_PPC = 1
	R_PPC_ADDR24          R_PPC = 2
	R_PPC_ADDR16          R_PPC = 3
	R_PPC_ADDR16_LO       R_PPC = 4
	R_PPC_ADDR16_HI       R_PPC = 5
	R_PPC_ADDR16_HA       R_PPC = 6
	R_PPC_ADDR14          R_PPC = 7
	R_PPC_ADDR14_BRTAKEN  R_PPC = 8
	R_PPC_ADDR14_BRNTAKEN R_PPC = 9
	R_PPC_REL24           R_PPC = 10
	R_PPC_REL14           R_PPC = 11
	R_PPC_REL14_BRTAKEN   R_PPC = 12
	R_PPC_REL14_BRNTAKEN  R_PPC = 13
	R_PPC_GOT16           R_PPC = 14
	R_PPC_GOT16_LO        R_PPC = 15
	R_PPC_GOT16_HI        R_PPC = 16
	R_PPC_GOT16_HA        R_PPC = 17
	R_PPC_PLTREL24        R_PPC = 18
	R_PPC_COPY            R_PPC = 19
	R_PPC_GLOB_DAT        R_PPC = 20
	R_PPC_JMP_SLOT        R_PPC = 21
	R_PPC_RELATIVE        R_PPC = 22
	R_PPC_LOCAL24PC       R_PPC = 23
	R_PPC_UADDR32         R_PPC = 24
	R_PPC_UADDR16         R_PPC = 25
	R_PPC_REL32           R_PPC = 26
	R_PPC_PLT32           R_PPC = 27
	R_PPC_PLTREL32        R_PPC = 28
	R_PPC_PLT16_LO        R_PPC = 29
	R_PPC_PLT16_HI        R_PPC = 30
	R_PPC_PLT16_HA        R_PPC = 31
	R_PPC_SDAREL16        R_PPC = 32
	R_PPC_SECTOFF         R_PPC = 33
	R_PPC_SECTOFF_LO      R_PPC = 34
	R_PPC_SECTOFF_HI      R_PPC = 35
	R_PPC_SECTOFF_HA      R_PPC = 36
	R_PPC_TLS             R_PPC = 67
	R_PPC_DTPMOD32        R_PPC = 68
	R_PPC_TPREL16         R_PPC = 69
	R_PPC_TPREL16_LO      R_PPC = 70
	R_PPC_TPREL16_HI      R_PPC = 71
	R_PPC_TPREL16_HA      R_PPC = 72
	R_PPC_TPREL32         R_PPC = 73
	R_PPC_DTPREL16        R_PPC = 74
	R_PPC_DTPREL16_LO     R_PPC = 75
	R_PPC_DTPREL16_HI     R_PPC = 76
	R_PPC_DTPREL16_HA     R_PPC = 77
	R_PPC_DTPREL32        R_PPC = 78
	R_PPC_GOT_TLSGD16     R_PPC = 79
	R_PPC_GOT_TLSGD16_LO  R_PPC = 80
	R_PPC_GOT_TLSGD16_HI  R_PPC = 81
	R_PPC_GOT_TLSGD16_HA  R_PPC = 82
	R_PPC_GOT_TLSLD16     R_PPC = 83
	R_PPC_GOT_TLSLD16_LO  R_PPC = 84
	R_PPC_GOT_TLSLD16_HI  R_PPC = 85
	R_PPC_GOT_TLSLD16_HA  R_PPC = 86
	R_PPC_GOT_TPREL16     R_PPC = 87
	R_PPC_GOT_TPREL16_LO  R_PPC = 88
	R_PPC_GOT_TPREL16_HI  R_PPC = 89
	R_PPC_GOT_TPREL16_HA  R_PPC = 90
	R_PPC_EMB_NADDR32     R_PPC = 101
	R_PPC_EMB_NADDR16     R_PPC = 102
	R_PPC_EMB_NADDR16_LO  R_PPC = 103
	R_PPC_EMB_NADDR16_HI  R_PPC = 104
	R_PPC_EMB_NADDR16_HA  R_PPC = 105
	R_PPC_EMB_SDAI16      R_PPC = 106
	R_PPC_EMB_SDA2I16     R_PPC = 107
	R_PPC_EMB_SDA2REL     R_PPC = 108
	R_PPC_EMB_SDA21       R_PPC = 109
	R_PPC_EMB_MRKREF      R_PPC = 110
	R_PPC_EMB_RELSEC16    R_PPC = 111
	R_PPC_EMB_RELST_LO    R_PPC = 112
	R_PPC_EMB_RELST_HI    R_PPC = 113
	R_PPC_EMB_RELST_HA    R_PPC = 114
	R_PPC_EMB_BIT_FLD     R_PPC = 115
	R_PPC_EMB_RELSDA      R_PPC = 116
)

func (i R_PPC) String() string
func (i R_PPC) GoString() string

// Relocation types for 64-bit PowerPC or Power Architecture processors.
type R_PPC64 int

const (
	R_PPC64_NONE               R_PPC64 = 0
	R_PPC64_ADDR32             R_PPC64 = 1
	R_PPC64_ADDR24             R_PPC64 = 2
	R_PPC64_ADDR16             R_PPC64 = 3
	R_PPC64_ADDR16_LO          R_PPC64 = 4
	R_PPC64_ADDR16_HI          R_PPC64 = 5
	R_PPC64_ADDR16_HA          R_PPC64 = 6
	R_PPC64_ADDR14             R_PPC64 = 7
	R_PPC64_ADDR14_BRTAKEN     R_PPC64 = 8
	R_PPC64_ADDR14_BRNTAKEN    R_PPC64 = 9
	R_PPC64_REL24              R_PPC64 = 10
	R_PPC64_REL14              R_PPC64 = 11
	R_PPC64_REL14_BRTAKEN      R_PPC64 = 12
	R_PPC64_REL14_BRNTAKEN     R_PPC64 = 13
	R_PPC64_GOT16              R_PPC64 = 14
	R_PPC64_GOT16_LO           R_PPC64 = 15
	R_PPC64_GOT16_HI           R_PPC64 = 16
	R_PPC64_GOT16_HA           R_PPC64 = 17
	R_PPC64_JMP_SLOT           R_PPC64 = 21
	R_PPC64_REL32              R_PPC64 = 26
	R_PPC64_ADDR64             R_PPC64 = 38
	R_PPC64_ADDR16_HIGHER      R_PPC64 = 39
	R_PPC64_ADDR16_HIGHERA     R_PPC64 = 40
	R_PPC64_ADDR16_HIGHEST     R_PPC64 = 41
	R_PPC64_ADDR16_HIGHESTA    R_PPC64 = 42
	R_PPC64_REL64              R_PPC64 = 44
	R_PPC64_TOC16              R_PPC64 = 47
	R_PPC64_TOC16_LO           R_PPC64 = 48
	R_PPC64_TOC16_HI           R_PPC64 = 49
	R_PPC64_TOC16_HA           R_PPC64 = 50
	R_PPC64_TOC                R_PPC64 = 51
	R_PPC64_ADDR16_DS          R_PPC64 = 56
	R_PPC64_ADDR16_LO_DS       R_PPC64 = 57
	R_PPC64_GOT16_DS           R_PPC64 = 58
	R_PPC64_GOT16_LO_DS        R_PPC64 = 59
	R_PPC64_TOC16_DS           R_PPC64 = 63
	R_PPC64_TOC16_LO_DS        R_PPC64 = 64
	R_PPC64_TLS                R_PPC64 = 67
	R_PPC64_DTPMOD64           R_PPC64 = 68
	R_PPC64_TPREL16            R_PPC64 = 69
	R_PPC64_TPREL16_LO         R_PPC64 = 70
	R_PPC64_TPREL16_HI         R_PPC64 = 71
	R_PPC64_TPREL16_HA         R_PPC64 = 72
	R_PPC64_TPREL64            R_PPC64 = 73
	R_PPC64_DTPREL16           R_PPC64 = 74
	R_PPC64_DTPREL16_LO        R_PPC64 = 75
	R_PPC64_DTPREL16_HI        R_PPC64 = 76
	R_PPC64_DTPREL16_HA        R_PPC64 = 77
	R_PPC64_DTPREL64           R_PPC64 = 78
	R_PPC64_GOT_TLSGD16        R_PPC64 = 79
	R_PPC64_GOT_TLSGD16_LO     R_PPC64 = 80
	R_PPC64_GOT_TLSGD16_HI     R_PPC64 = 81
	R_PPC64_GOT_TLSGD16_HA     R_PPC64 = 82
	R_PPC64_GOT_TLSLD16        R_PPC64 = 83
	R_PPC64_GOT_TLSLD16_LO     R_PPC64 = 84
	R_PPC64_GOT_TLSLD16_HI     R_PPC64 = 85
	R_PPC64_GOT_TLSLD16_HA     R_PPC64 = 86
	R_PPC64_GOT_TPREL16_DS     R_PPC64 = 87
	R_PPC64_GOT_TPREL16_LO_DS  R_PPC64 = 88
	R_PPC64_GOT_TPREL16_HI     R_PPC64 = 89
	R_PPC64_GOT_TPREL16_HA     R_PPC64 = 90
	R_PPC64_GOT_DTPREL16_DS    R_PPC64 = 91
	R_PPC64_GOT_DTPREL16_LO_DS R_PPC64 = 92
	R_PPC64_GOT_DTPREL16_HI    R_PPC64 = 93
	R_PPC64_GOT_DTPREL16_HA    R_PPC64 = 94
	R_PPC64_TPREL16_DS         R_PPC64 = 95
	R_PPC64_TPREL16_LO_DS      R_PPC64 = 96
	R_PPC64_TPREL16_HIGHER     R_PPC64 = 97
	R_PPC64_TPREL16_HIGHERA    R_PPC64 = 98
	R_PPC64_TPREL16_HIGHEST    R_PPC64 = 99
	R_PPC64_TPREL16_HIGHESTA   R_PPC64 = 100
	R_PPC64_DTPREL16_DS        R_PPC64 = 101
	R_PPC64_DTPREL16_LO_DS     R_PPC64 = 102
	R_PPC64_DTPREL16_HIGHER    R_PPC64 = 103
	R_PPC64_DTPREL16_HIGHERA   R_PPC64 = 104
	R_PPC64_DTPREL16_HIGHEST   R_PPC64 = 105
	R_PPC64_DTPREL16_HIGHESTA  R_PPC64 = 106
	R_PPC64_TLSGD              R_PPC64 = 107
	R_PPC64_TLSLD              R_PPC64 = 108
	R_PPC64_REL16              R_PPC64 = 249
	R_PPC64_REL16_LO           R_PPC64 = 250
	R_PPC64_REL16_HI           R_PPC64 = 251
	R_PPC64_REL16_HA           R_PPC64 = 252
)

func (i R_PPC64) String() string
func (i R_PPC64) GoString() string

// Relocation types for SPARC.
type R_SPARC int

const (
	R_SPARC_NONE     R_SPARC = 0
	R_SPARC_8        R_SPARC = 1
	R_SPARC_16       R_SPARC = 2
	R_SPARC_32       R_SPARC = 3
	R_SPARC_DISP8    R_SPARC = 4
	R_SPARC_DISP16   R_SPARC = 5
	R_SPARC_DISP32   R_SPARC = 6
	R_SPARC_WDISP30  R_SPARC = 7
	R_SPARC_WDISP22  R_SPARC = 8
	R_SPARC_HI22     R_SPARC = 9
	R_SPARC_22       R_SPARC = 10
	R_SPARC_13       R_SPARC = 11
	R_SPARC_LO10     R_SPARC = 12
	R_SPARC_GOT10    R_SPARC = 13
	R_SPARC_GOT13    R_SPARC = 14
	R_SPARC_GOT22    R_SPARC = 15
	R_SPARC_PC10     R_SPARC = 16
	R_SPARC_PC22     R_SPARC = 17
	R_SPARC_WPLT30   R_SPARC = 18
	R_SPARC_COPY     R_SPARC = 19
	R_SPARC_GLOB_DAT R_SPARC = 20
	R_SPARC_JMP_SLOT R_SPARC = 21
	R_SPARC_RELATIVE R_SPARC = 22
	R_SPARC_UA32     R_SPARC = 23
	R_SPARC_PLT32    R_SPARC = 24
	R_SPARC_HIPLT22  R_SPARC = 25
	R_SPARC_LOPLT10  R_SPARC = 26
	R_SPARC_PCPLT32  R_SPARC = 27
	R_SPARC_PCPLT22  R_SPARC = 28
	R_SPARC_PCPLT10  R_SPARC = 29
	R_SPARC_10       R_SPARC = 30
	R_SPARC_11       R_SPARC = 31
	R_SPARC_64       R_SPARC = 32
	R_SPARC_OLO10    R_SPARC = 33
	R_SPARC_HH22     R_SPARC = 34
	R_SPARC_HM10     R_SPARC = 35
	R_SPARC_LM22     R_SPARC = 36
	R_SPARC_PC_HH22  R_SPARC = 37
	R_SPARC_PC_HM10  R_SPARC = 38
	R_SPARC_PC_LM22  R_SPARC = 39
	R_SPARC_WDISP16  R_SPARC = 40
	R_SPARC_WDISP19  R_SPARC = 41
	R_SPARC_GLOB_JMP R_SPARC = 42
	R_SPARC_7        R_SPARC = 43
	R_SPARC_5        R_SPARC = 44
	R_SPARC_6        R_SPARC = 45
	R_SPARC_DISP64   R_SPARC = 46
	R_SPARC_PLT64    R_SPARC = 47
	R_SPARC_HIX22    R_SPARC = 48
	R_SPARC_LOX10    R_SPARC = 49
	R_SPARC_H44      R_SPARC = 50
	R_SPARC_M44      R_SPARC = 51
	R_SPARC_L44      R_SPARC = 52
	R_SPARC_REGISTER R_SPARC = 53
	R_SPARC_UA64     R_SPARC = 54
	R_SPARC_UA16     R_SPARC = 55
)

func (i R_SPARC) String() string
func (i R_SPARC) GoString() string

// Magic number for the elf trampoline, chosen wisely to be an immediate value.
const ARM_MAGIC_TRAMP_NUMBER = 0x5c000003

// ELF32 File header.
type Header32 struct {
	Ident     [EI_NIDENT]byte
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint32
	Phoff     uint32
	Shoff     uint32
	Flags     uint32
	Ehsize    uint16
	Phentsize uint16
	Phnum     uint16
	Shentsize uint16
	Shnum     uint16
	Shstrndx  uint16
}

// ELF32 Section header.
type Section32 struct {
	Name      uint32
	Type      uint32
	Flags     uint32
	Addr      uint32
	Off       uint32
	Size      uint32
	Link      uint32
	Info      uint32
	Addralign uint32
	Entsize   uint32
}

// ELF32 Program header.
type Prog32 struct {
	Type   uint32
	Off    uint32
	Vaddr  uint32
	Paddr  uint32
	Filesz uint32
	Memsz  uint32
	Flags  uint32
	Align  uint32
}

// ELF32 Dynamic structure.  The ".dynamic" section contains an array of them.
type Dyn32 struct {
	Tag int32
	Val uint32
}

// ELF32 Compression header.
type Chdr32 struct {
	Type      uint32
	Size      uint32
	Addralign uint32
}

// ELF32 Relocations that don't need an addend field.
type Rel32 struct {
	Off  uint32
	Info uint32
}

// ELF32 Relocations that need an addend field.
type Rela32 struct {
	Off    uint32
	Info   uint32
	Addend int32
}

func R_SYM32(info uint32) uint32
func R_TYPE32(info uint32) uint32
func R_INFO32(sym, typ uint32) uint32

// ELF32 Symbol.
type Sym32 struct {
	Name  uint32
	Value uint32
	Size  uint32
	Info  uint8
	Other uint8
	Shndx uint16
}

const Sym32Size = 16

func ST_BIND(info uint8) SymBind
func ST_TYPE(info uint8) SymType
func ST_INFO(bind SymBind, typ SymType) uint8

func ST_VISIBILITY(other uint8) SymVis

// ELF64 file header.
type Header64 struct {
	Ident     [EI_NIDENT]byte
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint64
	Phoff     uint64
	Shoff     uint64
	Flags     uint32
	Ehsize    uint16
	Phentsize uint16
	Phnum     uint16
	Shentsize uint16
	Shnum     uint16
	Shstrndx  uint16
}

// ELF64 Section header.
type Section64 struct {
	Name      uint32
	Type      uint32
	Flags     uint64
	Addr      uint64
	Off       uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64
}

// ELF64 Program header.
type Prog64 struct {
	Type   uint32
	Flags  uint32
	Off    uint64
	Vaddr  uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}

// ELF64 Dynamic structure.  The ".dynamic" section contains an array of them.
type Dyn64 struct {
	Tag int64
	Val uint64
}

// ELF64 Compression header.
type Chdr64 struct {
	Type      uint32
	_         uint32
	Size      uint64
	Addralign uint64
}

/* ELF64 relocations that don't need an addend field. */
type Rel64 struct {
	Off  uint64
	Info uint64
}

/* ELF64 relocations that need an addend field. */
type Rela64 struct {
	Off    uint64
	Info   uint64
	Addend int64
}

func R_SYM64(info uint64) uint32
func R_TYPE64(info uint64) uint32
func R_INFO(sym, typ uint32) uint64

// ELF64 symbol table entries.
type Sym64 struct {
	Name  uint32
	Info  uint8
	Other uint8
	Shndx uint16
	Value uint64
	Size  uint64
}

const Sym64Size = 24
