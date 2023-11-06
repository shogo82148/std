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
 * "System V ABI" (http://www.sco.com/developers/gabi/latest/ch4.eheader.html)
 * "ELF for the ARM® 64-bit Architecture (AArch64)" (ARM IHI 0056B)
 * "RISC-V ELF psABI specification" (https://github.com/riscv/riscv-elf-psabi-doc/blob/master/riscv-elf.adoc)
 * llvm/BinaryFormat/ELF.h - ELF constants and structures
 *
 * Copyright (c) 1996-1998 John D. Polstra.  All rights reserved.
 * Copyright (c) 2001 David E. O'Brien
 * Portions Copyright 2009 The Go Authors. All rights reserved.
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
	ELFOSABI_AROS       OSABI = 15
	ELFOSABI_FENIXOS    OSABI = 16
	ELFOSABI_CLOUDABI   OSABI = 17
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
	EM_NONE          Machine = 0
	EM_M32           Machine = 1
	EM_SPARC         Machine = 2
	EM_386           Machine = 3
	EM_68K           Machine = 4
	EM_88K           Machine = 5
	EM_860           Machine = 7
	EM_MIPS          Machine = 8
	EM_S370          Machine = 9
	EM_MIPS_RS3_LE   Machine = 10
	EM_PARISC        Machine = 15
	EM_VPP500        Machine = 17
	EM_SPARC32PLUS   Machine = 18
	EM_960           Machine = 19
	EM_PPC           Machine = 20
	EM_PPC64         Machine = 21
	EM_S390          Machine = 22
	EM_V800          Machine = 36
	EM_FR20          Machine = 37
	EM_RH32          Machine = 38
	EM_RCE           Machine = 39
	EM_ARM           Machine = 40
	EM_SH            Machine = 42
	EM_SPARCV9       Machine = 43
	EM_TRICORE       Machine = 44
	EM_ARC           Machine = 45
	EM_H8_300        Machine = 46
	EM_H8_300H       Machine = 47
	EM_H8S           Machine = 48
	EM_H8_500        Machine = 49
	EM_IA_64         Machine = 50
	EM_MIPS_X        Machine = 51
	EM_COLDFIRE      Machine = 52
	EM_68HC12        Machine = 53
	EM_MMA           Machine = 54
	EM_PCP           Machine = 55
	EM_NCPU          Machine = 56
	EM_NDR1          Machine = 57
	EM_STARCORE      Machine = 58
	EM_ME16          Machine = 59
	EM_ST100         Machine = 60
	EM_TINYJ         Machine = 61
	EM_X86_64        Machine = 62
	EM_PDSP          Machine = 63
	EM_PDP10         Machine = 64
	EM_PDP11         Machine = 65
	EM_FX66          Machine = 66
	EM_ST9PLUS       Machine = 67
	EM_ST7           Machine = 68
	EM_68HC16        Machine = 69
	EM_68HC11        Machine = 70
	EM_68HC08        Machine = 71
	EM_68HC05        Machine = 72
	EM_SVX           Machine = 73
	EM_ST19          Machine = 74
	EM_VAX           Machine = 75
	EM_CRIS          Machine = 76
	EM_JAVELIN       Machine = 77
	EM_FIREPATH      Machine = 78
	EM_ZSP           Machine = 79
	EM_MMIX          Machine = 80
	EM_HUANY         Machine = 81
	EM_PRISM         Machine = 82
	EM_AVR           Machine = 83
	EM_FR30          Machine = 84
	EM_D10V          Machine = 85
	EM_D30V          Machine = 86
	EM_V850          Machine = 87
	EM_M32R          Machine = 88
	EM_MN10300       Machine = 89
	EM_MN10200       Machine = 90
	EM_PJ            Machine = 91
	EM_OPENRISC      Machine = 92
	EM_ARC_COMPACT   Machine = 93
	EM_XTENSA        Machine = 94
	EM_VIDEOCORE     Machine = 95
	EM_TMM_GPP       Machine = 96
	EM_NS32K         Machine = 97
	EM_TPC           Machine = 98
	EM_SNP1K         Machine = 99
	EM_ST200         Machine = 100
	EM_IP2K          Machine = 101
	EM_MAX           Machine = 102
	EM_CR            Machine = 103
	EM_F2MC16        Machine = 104
	EM_MSP430        Machine = 105
	EM_BLACKFIN      Machine = 106
	EM_SE_C33        Machine = 107
	EM_SEP           Machine = 108
	EM_ARCA          Machine = 109
	EM_UNICORE       Machine = 110
	EM_EXCESS        Machine = 111
	EM_DXP           Machine = 112
	EM_ALTERA_NIOS2  Machine = 113
	EM_CRX           Machine = 114
	EM_XGATE         Machine = 115
	EM_C166          Machine = 116
	EM_M16C          Machine = 117
	EM_DSPIC30F      Machine = 118
	EM_CE            Machine = 119
	EM_M32C          Machine = 120
	EM_TSK3000       Machine = 131
	EM_RS08          Machine = 132
	EM_SHARC         Machine = 133
	EM_ECOG2         Machine = 134
	EM_SCORE7        Machine = 135
	EM_DSP24         Machine = 136
	EM_VIDEOCORE3    Machine = 137
	EM_LATTICEMICO32 Machine = 138
	EM_SE_C17        Machine = 139
	EM_TI_C6000      Machine = 140
	EM_TI_C2000      Machine = 141
	EM_TI_C5500      Machine = 142
	EM_TI_ARP32      Machine = 143
	EM_TI_PRU        Machine = 144
	EM_MMDSP_PLUS    Machine = 160
	EM_CYPRESS_M8C   Machine = 161
	EM_R32C          Machine = 162
	EM_TRIMEDIA      Machine = 163
	EM_QDSP6         Machine = 164
	EM_8051          Machine = 165
	EM_STXP7X        Machine = 166
	EM_NDS32         Machine = 167
	EM_ECOG1         Machine = 168
	EM_ECOG1X        Machine = 168
	EM_MAXQ30        Machine = 169
	EM_XIMO16        Machine = 170
	EM_MANIK         Machine = 171
	EM_CRAYNV2       Machine = 172
	EM_RX            Machine = 173
	EM_METAG         Machine = 174
	EM_MCST_ELBRUS   Machine = 175
	EM_ECOG16        Machine = 176
	EM_CR16          Machine = 177
	EM_ETPU          Machine = 178
	EM_SLE9X         Machine = 179
	EM_L10M          Machine = 180
	EM_K10M          Machine = 181
	EM_AARCH64       Machine = 183
	EM_AVR32         Machine = 185
	EM_STM8          Machine = 186
	EM_TILE64        Machine = 187
	EM_TILEPRO       Machine = 188
	EM_MICROBLAZE    Machine = 189
	EM_CUDA          Machine = 190
	EM_TILEGX        Machine = 191
	EM_CLOUDSHIELD   Machine = 192
	EM_COREA_1ST     Machine = 193
	EM_COREA_2ND     Machine = 194
	EM_ARC_COMPACT2  Machine = 195
	EM_OPEN8         Machine = 196
	EM_RL78          Machine = 197
	EM_VIDEOCORE5    Machine = 198
	EM_78KOR         Machine = 199
	EM_56800EX       Machine = 200
	EM_BA1           Machine = 201
	EM_BA2           Machine = 202
	EM_XCORE         Machine = 203
	EM_MCHP_PIC      Machine = 204
	EM_INTEL205      Machine = 205
	EM_INTEL206      Machine = 206
	EM_INTEL207      Machine = 207
	EM_INTEL208      Machine = 208
	EM_INTEL209      Machine = 209
	EM_KM32          Machine = 210
	EM_KMX32         Machine = 211
	EM_KMX16         Machine = 212
	EM_KMX8          Machine = 213
	EM_KVARC         Machine = 214
	EM_CDP           Machine = 215
	EM_COGE          Machine = 216
	EM_COOL          Machine = 217
	EM_NORC          Machine = 218
	EM_CSR_KALIMBA   Machine = 219
	EM_Z80           Machine = 220
	EM_VISIUM        Machine = 221
	EM_FT32          Machine = 222
	EM_MOXIE         Machine = 223
	EM_AMDGPU        Machine = 224
	EM_RISCV         Machine = 243
	EM_LANAI         Machine = 244
	EM_BPF           Machine = 247
	EM_LOONGARCH     Machine = 258

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
	SHT_MIPS_ABIFLAGS  SectionType = 0x7000002a
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
	COMPRESS_ZSTD   CompressionType = 2
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

	PT_LOOS ProgType = 0x60000000

	PT_GNU_EH_FRAME ProgType = 0x6474e550
	PT_GNU_STACK    ProgType = 0x6474e551
	PT_GNU_RELRO    ProgType = 0x6474e552
	PT_GNU_PROPERTY ProgType = 0x6474e553
	PT_GNU_MBIND_LO ProgType = 0x6474e555
	PT_GNU_MBIND_HI ProgType = 0x6474f554

	PT_PAX_FLAGS ProgType = 0x65041580

	PT_OPENBSD_RANDOMIZE ProgType = 0x65a3dbe6
	PT_OPENBSD_WXNEEDED  ProgType = 0x65a3dbe7
	PT_OPENBSD_BOOTDATA  ProgType = 0x65a41be6

	PT_SUNW_EH_FRAME ProgType = 0x6474e550
	PT_SUNWSTACK     ProgType = 0x6ffffffb

	PT_HIOS ProgType = 0x6fffffff

	PT_LOPROC ProgType = 0x70000000

	PT_ARM_ARCHEXT ProgType = 0x70000000
	PT_ARM_EXIDX   ProgType = 0x70000001

	PT_AARCH64_ARCHEXT ProgType = 0x70000000
	PT_AARCH64_UNWIND  ProgType = 0x70000001

	PT_MIPS_REGINFO  ProgType = 0x70000000
	PT_MIPS_RTPROC   ProgType = 0x70000001
	PT_MIPS_OPTIONS  ProgType = 0x70000002
	PT_MIPS_ABIFLAGS ProgType = 0x70000003

	PT_S390_PGSTE ProgType = 0x70000000

	PT_HIPROC ProgType = 0x7fffffff
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
	DT_SYMTAB_SHNDX    DynTag = 34

	DT_LOOS DynTag = 0x6000000d
	DT_HIOS DynTag = 0x6ffff000

	DT_VALRNGLO       DynTag = 0x6ffffd00
	DT_GNU_PRELINKED  DynTag = 0x6ffffdf5
	DT_GNU_CONFLICTSZ DynTag = 0x6ffffdf6
	DT_GNU_LIBLISTSZ  DynTag = 0x6ffffdf7
	DT_CHECKSUM       DynTag = 0x6ffffdf8
	DT_PLTPADSZ       DynTag = 0x6ffffdf9
	DT_MOVEENT        DynTag = 0x6ffffdfa
	DT_MOVESZ         DynTag = 0x6ffffdfb
	DT_FEATURE        DynTag = 0x6ffffdfc
	DT_POSFLAG_1      DynTag = 0x6ffffdfd
	DT_SYMINSZ        DynTag = 0x6ffffdfe
	DT_SYMINENT       DynTag = 0x6ffffdff
	DT_VALRNGHI       DynTag = 0x6ffffdff

	DT_ADDRRNGLO    DynTag = 0x6ffffe00
	DT_GNU_HASH     DynTag = 0x6ffffef5
	DT_TLSDESC_PLT  DynTag = 0x6ffffef6
	DT_TLSDESC_GOT  DynTag = 0x6ffffef7
	DT_GNU_CONFLICT DynTag = 0x6ffffef8
	DT_GNU_LIBLIST  DynTag = 0x6ffffef9
	DT_CONFIG       DynTag = 0x6ffffefa
	DT_DEPAUDIT     DynTag = 0x6ffffefb
	DT_AUDIT        DynTag = 0x6ffffefc
	DT_PLTPAD       DynTag = 0x6ffffefd
	DT_MOVETAB      DynTag = 0x6ffffefe
	DT_SYMINFO      DynTag = 0x6ffffeff
	DT_ADDRRNGHI    DynTag = 0x6ffffeff

	DT_VERSYM     DynTag = 0x6ffffff0
	DT_RELACOUNT  DynTag = 0x6ffffff9
	DT_RELCOUNT   DynTag = 0x6ffffffa
	DT_FLAGS_1    DynTag = 0x6ffffffb
	DT_VERDEF     DynTag = 0x6ffffffc
	DT_VERDEFNUM  DynTag = 0x6ffffffd
	DT_VERNEED    DynTag = 0x6ffffffe
	DT_VERNEEDNUM DynTag = 0x6fffffff

	DT_LOPROC DynTag = 0x70000000

	DT_MIPS_RLD_VERSION           DynTag = 0x70000001
	DT_MIPS_TIME_STAMP            DynTag = 0x70000002
	DT_MIPS_ICHECKSUM             DynTag = 0x70000003
	DT_MIPS_IVERSION              DynTag = 0x70000004
	DT_MIPS_FLAGS                 DynTag = 0x70000005
	DT_MIPS_BASE_ADDRESS          DynTag = 0x70000006
	DT_MIPS_MSYM                  DynTag = 0x70000007
	DT_MIPS_CONFLICT              DynTag = 0x70000008
	DT_MIPS_LIBLIST               DynTag = 0x70000009
	DT_MIPS_LOCAL_GOTNO           DynTag = 0x7000000a
	DT_MIPS_CONFLICTNO            DynTag = 0x7000000b
	DT_MIPS_LIBLISTNO             DynTag = 0x70000010
	DT_MIPS_SYMTABNO              DynTag = 0x70000011
	DT_MIPS_UNREFEXTNO            DynTag = 0x70000012
	DT_MIPS_GOTSYM                DynTag = 0x70000013
	DT_MIPS_HIPAGENO              DynTag = 0x70000014
	DT_MIPS_RLD_MAP               DynTag = 0x70000016
	DT_MIPS_DELTA_CLASS           DynTag = 0x70000017
	DT_MIPS_DELTA_CLASS_NO        DynTag = 0x70000018
	DT_MIPS_DELTA_INSTANCE        DynTag = 0x70000019
	DT_MIPS_DELTA_INSTANCE_NO     DynTag = 0x7000001a
	DT_MIPS_DELTA_RELOC           DynTag = 0x7000001b
	DT_MIPS_DELTA_RELOC_NO        DynTag = 0x7000001c
	DT_MIPS_DELTA_SYM             DynTag = 0x7000001d
	DT_MIPS_DELTA_SYM_NO          DynTag = 0x7000001e
	DT_MIPS_DELTA_CLASSSYM        DynTag = 0x70000020
	DT_MIPS_DELTA_CLASSSYM_NO     DynTag = 0x70000021
	DT_MIPS_CXX_FLAGS             DynTag = 0x70000022
	DT_MIPS_PIXIE_INIT            DynTag = 0x70000023
	DT_MIPS_SYMBOL_LIB            DynTag = 0x70000024
	DT_MIPS_LOCALPAGE_GOTIDX      DynTag = 0x70000025
	DT_MIPS_LOCAL_GOTIDX          DynTag = 0x70000026
	DT_MIPS_HIDDEN_GOTIDX         DynTag = 0x70000027
	DT_MIPS_PROTECTED_GOTIDX      DynTag = 0x70000028
	DT_MIPS_OPTIONS               DynTag = 0x70000029
	DT_MIPS_INTERFACE             DynTag = 0x7000002a
	DT_MIPS_DYNSTR_ALIGN          DynTag = 0x7000002b
	DT_MIPS_INTERFACE_SIZE        DynTag = 0x7000002c
	DT_MIPS_RLD_TEXT_RESOLVE_ADDR DynTag = 0x7000002d
	DT_MIPS_PERF_SUFFIX           DynTag = 0x7000002e
	DT_MIPS_COMPACT_SIZE          DynTag = 0x7000002f
	DT_MIPS_GP_VALUE              DynTag = 0x70000030
	DT_MIPS_AUX_DYNAMIC           DynTag = 0x70000031
	DT_MIPS_PLTGOT                DynTag = 0x70000032
	DT_MIPS_RWPLT                 DynTag = 0x70000034
	DT_MIPS_RLD_MAP_REL           DynTag = 0x70000035

	DT_PPC_GOT DynTag = 0x70000000
	DT_PPC_OPT DynTag = 0x70000001

	DT_PPC64_GLINK DynTag = 0x70000000
	DT_PPC64_OPD   DynTag = 0x70000001
	DT_PPC64_OPDSZ DynTag = 0x70000002
	DT_PPC64_OPT   DynTag = 0x70000003

	DT_SPARC_REGISTER DynTag = 0x70000001

	DT_AUXILIARY DynTag = 0x7ffffffd
	DT_USED      DynTag = 0x7ffffffe
	DT_FILTER    DynTag = 0x7fffffff

	DT_HIPROC DynTag = 0x7fffffff
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

// DT_FLAGS_1 values.
type DynFlag1 uint32

const (
	// Indicates that all relocations for this object must be processed before
	// returning control to the program.
	DF_1_NOW DynFlag1 = 0x00000001
	// Unused.
	DF_1_GLOBAL DynFlag1 = 0x00000002
	// Indicates that the object is a member of a group.
	DF_1_GROUP DynFlag1 = 0x00000004
	// Indicates that the object cannot be deleted from a process.
	DF_1_NODELETE DynFlag1 = 0x00000008
	// Meaningful only for filters. Indicates that all associated filtees be
	// processed immediately.
	DF_1_LOADFLTR DynFlag1 = 0x00000010
	// Indicates that this object's initialization section be run before any other
	// objects loaded.
	DF_1_INITFIRST DynFlag1 = 0x00000020
	// Indicates that the object cannot be added to a running process with dlopen.
	DF_1_NOOPEN DynFlag1 = 0x00000040
	// Indicates the object requires $ORIGIN processing.
	DF_1_ORIGIN DynFlag1 = 0x00000080
	// Indicates that the object should use direct binding information.
	DF_1_DIRECT DynFlag1 = 0x00000100
	// Unused.
	DF_1_TRANS DynFlag1 = 0x00000200
	// Indicates that the objects symbol table is to interpose before all symbols
	// except the primary load object, which is typically the executable.
	DF_1_INTERPOSE DynFlag1 = 0x00000400
	// Indicates that the search for dependencies of this object ignores any
	// default library search paths.
	DF_1_NODEFLIB DynFlag1 = 0x00000800
	// Indicates that this object is not dumped by dldump. Candidates are objects
	// with no relocations that might get included when generating alternative
	// objects using.
	DF_1_NODUMP DynFlag1 = 0x00001000
	// Identifies this object as a configuration alternative object generated by
	// crle. Triggers the runtime linker to search for a configuration file $ORIGIN/ld.config.app-name.
	DF_1_CONFALT DynFlag1 = 0x00002000
	// Meaningful only for filtees. Terminates a filters search for any
	// further filtees.
	DF_1_ENDFILTEE DynFlag1 = 0x00004000
	// Indicates that this object has displacement relocations applied.
	DF_1_DISPRELDNE DynFlag1 = 0x00008000
	// Indicates that this object has displacement relocations pending.
	DF_1_DISPRELPND DynFlag1 = 0x00010000
	// Indicates that this object contains symbols that cannot be directly
	// bound to.
	DF_1_NODIRECT DynFlag1 = 0x00020000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_IGNMULDEF DynFlag1 = 0x00040000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_NOKSYMS DynFlag1 = 0x00080000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_NOHDR DynFlag1 = 0x00100000
	// Indicates that this object has been edited or has been modified since the
	// objects original construction by the link-editor.
	DF_1_EDITED DynFlag1 = 0x00200000
	// Reserved for internal use by the kernel runtime-linker.
	DF_1_NORELOC DynFlag1 = 0x00400000
	// Indicates that the object contains individual symbols that should interpose
	// before all symbols except the primary load object, which is typically the
	// executable.
	DF_1_SYMINTPOSE DynFlag1 = 0x00800000
	// Indicates that the executable requires global auditing.
	DF_1_GLOBAUDIT DynFlag1 = 0x01000000
	// Indicates that the object defines, or makes reference to singleton symbols.
	DF_1_SINGLETON DynFlag1 = 0x02000000
	// Indicates that the object is a stub.
	DF_1_STUB DynFlag1 = 0x04000000
	// Indicates that the object is a position-independent executable.
	DF_1_PIE DynFlag1 = 0x08000000
	// Indicates that the object is a kernel module.
	DF_1_KMOD DynFlag1 = 0x10000000
	// Indicates that the object is a weak standard filter.
	DF_1_WEAKFILTER DynFlag1 = 0x20000000
	// Unused.
	DF_1_NOCOMMON DynFlag1 = 0x40000000
)

func (i DynFlag1) String() string
func (i DynFlag1) GoString() string

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
	R_X86_64_NONE            R_X86_64 = 0
	R_X86_64_64              R_X86_64 = 1
	R_X86_64_PC32            R_X86_64 = 2
	R_X86_64_GOT32           R_X86_64 = 3
	R_X86_64_PLT32           R_X86_64 = 4
	R_X86_64_COPY            R_X86_64 = 5
	R_X86_64_GLOB_DAT        R_X86_64 = 6
	R_X86_64_JMP_SLOT        R_X86_64 = 7
	R_X86_64_RELATIVE        R_X86_64 = 8
	R_X86_64_GOTPCREL        R_X86_64 = 9
	R_X86_64_32              R_X86_64 = 10
	R_X86_64_32S             R_X86_64 = 11
	R_X86_64_16              R_X86_64 = 12
	R_X86_64_PC16            R_X86_64 = 13
	R_X86_64_8               R_X86_64 = 14
	R_X86_64_PC8             R_X86_64 = 15
	R_X86_64_DTPMOD64        R_X86_64 = 16
	R_X86_64_DTPOFF64        R_X86_64 = 17
	R_X86_64_TPOFF64         R_X86_64 = 18
	R_X86_64_TLSGD           R_X86_64 = 19
	R_X86_64_TLSLD           R_X86_64 = 20
	R_X86_64_DTPOFF32        R_X86_64 = 21
	R_X86_64_GOTTPOFF        R_X86_64 = 22
	R_X86_64_TPOFF32         R_X86_64 = 23
	R_X86_64_PC64            R_X86_64 = 24
	R_X86_64_GOTOFF64        R_X86_64 = 25
	R_X86_64_GOTPC32         R_X86_64 = 26
	R_X86_64_GOT64           R_X86_64 = 27
	R_X86_64_GOTPCREL64      R_X86_64 = 28
	R_X86_64_GOTPC64         R_X86_64 = 29
	R_X86_64_GOTPLT64        R_X86_64 = 30
	R_X86_64_PLTOFF64        R_X86_64 = 31
	R_X86_64_SIZE32          R_X86_64 = 32
	R_X86_64_SIZE64          R_X86_64 = 33
	R_X86_64_GOTPC32_TLSDESC R_X86_64 = 34
	R_X86_64_TLSDESC_CALL    R_X86_64 = 35
	R_X86_64_TLSDESC         R_X86_64 = 36
	R_X86_64_IRELATIVE       R_X86_64 = 37
	R_X86_64_RELATIVE64      R_X86_64 = 38
	R_X86_64_PC32_BND        R_X86_64 = 39
	R_X86_64_PLT32_BND       R_X86_64 = 40
	R_X86_64_GOTPCRELX       R_X86_64 = 41
	R_X86_64_REX_GOTPCRELX   R_X86_64 = 42
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
	R_AARCH64_LD64_GOTOFF_LO15                R_AARCH64 = 310
	R_AARCH64_ADR_GOT_PAGE                    R_AARCH64 = 311
	R_AARCH64_LD64_GOT_LO12_NC                R_AARCH64 = 312
	R_AARCH64_LD64_GOTPAGE_LO15               R_AARCH64 = 313
	R_AARCH64_TLSGD_ADR_PREL21                R_AARCH64 = 512
	R_AARCH64_TLSGD_ADR_PAGE21                R_AARCH64 = 513
	R_AARCH64_TLSGD_ADD_LO12_NC               R_AARCH64 = 514
	R_AARCH64_TLSGD_MOVW_G1                   R_AARCH64 = 515
	R_AARCH64_TLSGD_MOVW_G0_NC                R_AARCH64 = 516
	R_AARCH64_TLSLD_ADR_PREL21                R_AARCH64 = 517
	R_AARCH64_TLSLD_ADR_PAGE21                R_AARCH64 = 518
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
	R_AARCH64_TLSLE_LDST128_TPREL_LO12        R_AARCH64 = 570
	R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC     R_AARCH64 = 571
	R_AARCH64_TLSLD_LDST128_DTPREL_LO12       R_AARCH64 = 572
	R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC    R_AARCH64 = 573
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
	R_ARM_NONE               R_ARM = 0
	R_ARM_PC24               R_ARM = 1
	R_ARM_ABS32              R_ARM = 2
	R_ARM_REL32              R_ARM = 3
	R_ARM_PC13               R_ARM = 4
	R_ARM_ABS16              R_ARM = 5
	R_ARM_ABS12              R_ARM = 6
	R_ARM_THM_ABS5           R_ARM = 7
	R_ARM_ABS8               R_ARM = 8
	R_ARM_SBREL32            R_ARM = 9
	R_ARM_THM_PC22           R_ARM = 10
	R_ARM_THM_PC8            R_ARM = 11
	R_ARM_AMP_VCALL9         R_ARM = 12
	R_ARM_SWI24              R_ARM = 13
	R_ARM_THM_SWI8           R_ARM = 14
	R_ARM_XPC25              R_ARM = 15
	R_ARM_THM_XPC22          R_ARM = 16
	R_ARM_TLS_DTPMOD32       R_ARM = 17
	R_ARM_TLS_DTPOFF32       R_ARM = 18
	R_ARM_TLS_TPOFF32        R_ARM = 19
	R_ARM_COPY               R_ARM = 20
	R_ARM_GLOB_DAT           R_ARM = 21
	R_ARM_JUMP_SLOT          R_ARM = 22
	R_ARM_RELATIVE           R_ARM = 23
	R_ARM_GOTOFF             R_ARM = 24
	R_ARM_GOTPC              R_ARM = 25
	R_ARM_GOT32              R_ARM = 26
	R_ARM_PLT32              R_ARM = 27
	R_ARM_CALL               R_ARM = 28
	R_ARM_JUMP24             R_ARM = 29
	R_ARM_THM_JUMP24         R_ARM = 30
	R_ARM_BASE_ABS           R_ARM = 31
	R_ARM_ALU_PCREL_7_0      R_ARM = 32
	R_ARM_ALU_PCREL_15_8     R_ARM = 33
	R_ARM_ALU_PCREL_23_15    R_ARM = 34
	R_ARM_LDR_SBREL_11_10_NC R_ARM = 35
	R_ARM_ALU_SBREL_19_12_NC R_ARM = 36
	R_ARM_ALU_SBREL_27_20_CK R_ARM = 37
	R_ARM_TARGET1            R_ARM = 38
	R_ARM_SBREL31            R_ARM = 39
	R_ARM_V4BX               R_ARM = 40
	R_ARM_TARGET2            R_ARM = 41
	R_ARM_PREL31             R_ARM = 42
	R_ARM_MOVW_ABS_NC        R_ARM = 43
	R_ARM_MOVT_ABS           R_ARM = 44
	R_ARM_MOVW_PREL_NC       R_ARM = 45
	R_ARM_MOVT_PREL          R_ARM = 46
	R_ARM_THM_MOVW_ABS_NC    R_ARM = 47
	R_ARM_THM_MOVT_ABS       R_ARM = 48
	R_ARM_THM_MOVW_PREL_NC   R_ARM = 49
	R_ARM_THM_MOVT_PREL      R_ARM = 50
	R_ARM_THM_JUMP19         R_ARM = 51
	R_ARM_THM_JUMP6          R_ARM = 52
	R_ARM_THM_ALU_PREL_11_0  R_ARM = 53
	R_ARM_THM_PC12           R_ARM = 54
	R_ARM_ABS32_NOI          R_ARM = 55
	R_ARM_REL32_NOI          R_ARM = 56
	R_ARM_ALU_PC_G0_NC       R_ARM = 57
	R_ARM_ALU_PC_G0          R_ARM = 58
	R_ARM_ALU_PC_G1_NC       R_ARM = 59
	R_ARM_ALU_PC_G1          R_ARM = 60
	R_ARM_ALU_PC_G2          R_ARM = 61
	R_ARM_LDR_PC_G1          R_ARM = 62
	R_ARM_LDR_PC_G2          R_ARM = 63
	R_ARM_LDRS_PC_G0         R_ARM = 64
	R_ARM_LDRS_PC_G1         R_ARM = 65
	R_ARM_LDRS_PC_G2         R_ARM = 66
	R_ARM_LDC_PC_G0          R_ARM = 67
	R_ARM_LDC_PC_G1          R_ARM = 68
	R_ARM_LDC_PC_G2          R_ARM = 69
	R_ARM_ALU_SB_G0_NC       R_ARM = 70
	R_ARM_ALU_SB_G0          R_ARM = 71
	R_ARM_ALU_SB_G1_NC       R_ARM = 72
	R_ARM_ALU_SB_G1          R_ARM = 73
	R_ARM_ALU_SB_G2          R_ARM = 74
	R_ARM_LDR_SB_G0          R_ARM = 75
	R_ARM_LDR_SB_G1          R_ARM = 76
	R_ARM_LDR_SB_G2          R_ARM = 77
	R_ARM_LDRS_SB_G0         R_ARM = 78
	R_ARM_LDRS_SB_G1         R_ARM = 79
	R_ARM_LDRS_SB_G2         R_ARM = 80
	R_ARM_LDC_SB_G0          R_ARM = 81
	R_ARM_LDC_SB_G1          R_ARM = 82
	R_ARM_LDC_SB_G2          R_ARM = 83
	R_ARM_MOVW_BREL_NC       R_ARM = 84
	R_ARM_MOVT_BREL          R_ARM = 85
	R_ARM_MOVW_BREL          R_ARM = 86
	R_ARM_THM_MOVW_BREL_NC   R_ARM = 87
	R_ARM_THM_MOVT_BREL      R_ARM = 88
	R_ARM_THM_MOVW_BREL      R_ARM = 89
	R_ARM_TLS_GOTDESC        R_ARM = 90
	R_ARM_TLS_CALL           R_ARM = 91
	R_ARM_TLS_DESCSEQ        R_ARM = 92
	R_ARM_THM_TLS_CALL       R_ARM = 93
	R_ARM_PLT32_ABS          R_ARM = 94
	R_ARM_GOT_ABS            R_ARM = 95
	R_ARM_GOT_PREL           R_ARM = 96
	R_ARM_GOT_BREL12         R_ARM = 97
	R_ARM_GOTOFF12           R_ARM = 98
	R_ARM_GOTRELAX           R_ARM = 99
	R_ARM_GNU_VTENTRY        R_ARM = 100
	R_ARM_GNU_VTINHERIT      R_ARM = 101
	R_ARM_THM_JUMP11         R_ARM = 102
	R_ARM_THM_JUMP8          R_ARM = 103
	R_ARM_TLS_GD32           R_ARM = 104
	R_ARM_TLS_LDM32          R_ARM = 105
	R_ARM_TLS_LDO32          R_ARM = 106
	R_ARM_TLS_IE32           R_ARM = 107
	R_ARM_TLS_LE32           R_ARM = 108
	R_ARM_TLS_LDO12          R_ARM = 109
	R_ARM_TLS_LE12           R_ARM = 110
	R_ARM_TLS_IE12GP         R_ARM = 111
	R_ARM_PRIVATE_0          R_ARM = 112
	R_ARM_PRIVATE_1          R_ARM = 113
	R_ARM_PRIVATE_2          R_ARM = 114
	R_ARM_PRIVATE_3          R_ARM = 115
	R_ARM_PRIVATE_4          R_ARM = 116
	R_ARM_PRIVATE_5          R_ARM = 117
	R_ARM_PRIVATE_6          R_ARM = 118
	R_ARM_PRIVATE_7          R_ARM = 119
	R_ARM_PRIVATE_8          R_ARM = 120
	R_ARM_PRIVATE_9          R_ARM = 121
	R_ARM_PRIVATE_10         R_ARM = 122
	R_ARM_PRIVATE_11         R_ARM = 123
	R_ARM_PRIVATE_12         R_ARM = 124
	R_ARM_PRIVATE_13         R_ARM = 125
	R_ARM_PRIVATE_14         R_ARM = 126
	R_ARM_PRIVATE_15         R_ARM = 127
	R_ARM_ME_TOO             R_ARM = 128
	R_ARM_THM_TLS_DESCSEQ16  R_ARM = 129
	R_ARM_THM_TLS_DESCSEQ32  R_ARM = 130
	R_ARM_THM_GOT_BREL12     R_ARM = 131
	R_ARM_THM_ALU_ABS_G0_NC  R_ARM = 132
	R_ARM_THM_ALU_ABS_G1_NC  R_ARM = 133
	R_ARM_THM_ALU_ABS_G2_NC  R_ARM = 134
	R_ARM_THM_ALU_ABS_G3     R_ARM = 135
	R_ARM_IRELATIVE          R_ARM = 160
	R_ARM_RXPC25             R_ARM = 249
	R_ARM_RSBREL32           R_ARM = 250
	R_ARM_THM_RPC22          R_ARM = 251
	R_ARM_RREL32             R_ARM = 252
	R_ARM_RABS32             R_ARM = 253
	R_ARM_RPC24              R_ARM = 254
	R_ARM_RBASE              R_ARM = 255
)

func (i R_ARM) String() string
func (i R_ARM) GoString() string

// Relocation types for 386.
type R_386 int

const (
	R_386_NONE          R_386 = 0
	R_386_32            R_386 = 1
	R_386_PC32          R_386 = 2
	R_386_GOT32         R_386 = 3
	R_386_PLT32         R_386 = 4
	R_386_COPY          R_386 = 5
	R_386_GLOB_DAT      R_386 = 6
	R_386_JMP_SLOT      R_386 = 7
	R_386_RELATIVE      R_386 = 8
	R_386_GOTOFF        R_386 = 9
	R_386_GOTPC         R_386 = 10
	R_386_32PLT         R_386 = 11
	R_386_TLS_TPOFF     R_386 = 14
	R_386_TLS_IE        R_386 = 15
	R_386_TLS_GOTIE     R_386 = 16
	R_386_TLS_LE        R_386 = 17
	R_386_TLS_GD        R_386 = 18
	R_386_TLS_LDM       R_386 = 19
	R_386_16            R_386 = 20
	R_386_PC16          R_386 = 21
	R_386_8             R_386 = 22
	R_386_PC8           R_386 = 23
	R_386_TLS_GD_32     R_386 = 24
	R_386_TLS_GD_PUSH   R_386 = 25
	R_386_TLS_GD_CALL   R_386 = 26
	R_386_TLS_GD_POP    R_386 = 27
	R_386_TLS_LDM_32    R_386 = 28
	R_386_TLS_LDM_PUSH  R_386 = 29
	R_386_TLS_LDM_CALL  R_386 = 30
	R_386_TLS_LDM_POP   R_386 = 31
	R_386_TLS_LDO_32    R_386 = 32
	R_386_TLS_IE_32     R_386 = 33
	R_386_TLS_LE_32     R_386 = 34
	R_386_TLS_DTPMOD32  R_386 = 35
	R_386_TLS_DTPOFF32  R_386 = 36
	R_386_TLS_TPOFF32   R_386 = 37
	R_386_SIZE32        R_386 = 38
	R_386_TLS_GOTDESC   R_386 = 39
	R_386_TLS_DESC_CALL R_386 = 40
	R_386_TLS_DESC      R_386 = 41
	R_386_IRELATIVE     R_386 = 42
	R_386_GOT32X        R_386 = 43
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

	R_MIPS_PC32 R_MIPS = 248
)

func (i R_MIPS) String() string
func (i R_MIPS) GoString() string

// Relocation types for LoongArch.
type R_LARCH int

const (
	R_LARCH_NONE                       R_LARCH = 0
	R_LARCH_32                         R_LARCH = 1
	R_LARCH_64                         R_LARCH = 2
	R_LARCH_RELATIVE                   R_LARCH = 3
	R_LARCH_COPY                       R_LARCH = 4
	R_LARCH_JUMP_SLOT                  R_LARCH = 5
	R_LARCH_TLS_DTPMOD32               R_LARCH = 6
	R_LARCH_TLS_DTPMOD64               R_LARCH = 7
	R_LARCH_TLS_DTPREL32               R_LARCH = 8
	R_LARCH_TLS_DTPREL64               R_LARCH = 9
	R_LARCH_TLS_TPREL32                R_LARCH = 10
	R_LARCH_TLS_TPREL64                R_LARCH = 11
	R_LARCH_IRELATIVE                  R_LARCH = 12
	R_LARCH_MARK_LA                    R_LARCH = 20
	R_LARCH_MARK_PCREL                 R_LARCH = 21
	R_LARCH_SOP_PUSH_PCREL             R_LARCH = 22
	R_LARCH_SOP_PUSH_ABSOLUTE          R_LARCH = 23
	R_LARCH_SOP_PUSH_DUP               R_LARCH = 24
	R_LARCH_SOP_PUSH_GPREL             R_LARCH = 25
	R_LARCH_SOP_PUSH_TLS_TPREL         R_LARCH = 26
	R_LARCH_SOP_PUSH_TLS_GOT           R_LARCH = 27
	R_LARCH_SOP_PUSH_TLS_GD            R_LARCH = 28
	R_LARCH_SOP_PUSH_PLT_PCREL         R_LARCH = 29
	R_LARCH_SOP_ASSERT                 R_LARCH = 30
	R_LARCH_SOP_NOT                    R_LARCH = 31
	R_LARCH_SOP_SUB                    R_LARCH = 32
	R_LARCH_SOP_SL                     R_LARCH = 33
	R_LARCH_SOP_SR                     R_LARCH = 34
	R_LARCH_SOP_ADD                    R_LARCH = 35
	R_LARCH_SOP_AND                    R_LARCH = 36
	R_LARCH_SOP_IF_ELSE                R_LARCH = 37
	R_LARCH_SOP_POP_32_S_10_5          R_LARCH = 38
	R_LARCH_SOP_POP_32_U_10_12         R_LARCH = 39
	R_LARCH_SOP_POP_32_S_10_12         R_LARCH = 40
	R_LARCH_SOP_POP_32_S_10_16         R_LARCH = 41
	R_LARCH_SOP_POP_32_S_10_16_S2      R_LARCH = 42
	R_LARCH_SOP_POP_32_S_5_20          R_LARCH = 43
	R_LARCH_SOP_POP_32_S_0_5_10_16_S2  R_LARCH = 44
	R_LARCH_SOP_POP_32_S_0_10_10_16_S2 R_LARCH = 45
	R_LARCH_SOP_POP_32_U               R_LARCH = 46
	R_LARCH_ADD8                       R_LARCH = 47
	R_LARCH_ADD16                      R_LARCH = 48
	R_LARCH_ADD24                      R_LARCH = 49
	R_LARCH_ADD32                      R_LARCH = 50
	R_LARCH_ADD64                      R_LARCH = 51
	R_LARCH_SUB8                       R_LARCH = 52
	R_LARCH_SUB16                      R_LARCH = 53
	R_LARCH_SUB24                      R_LARCH = 54
	R_LARCH_SUB32                      R_LARCH = 55
	R_LARCH_SUB64                      R_LARCH = 56
	R_LARCH_GNU_VTINHERIT              R_LARCH = 57
	R_LARCH_GNU_VTENTRY                R_LARCH = 58
	R_LARCH_B16                        R_LARCH = 64
	R_LARCH_B21                        R_LARCH = 65
	R_LARCH_B26                        R_LARCH = 66
	R_LARCH_ABS_HI20                   R_LARCH = 67
	R_LARCH_ABS_LO12                   R_LARCH = 68
	R_LARCH_ABS64_LO20                 R_LARCH = 69
	R_LARCH_ABS64_HI12                 R_LARCH = 70
	R_LARCH_PCALA_HI20                 R_LARCH = 71
	R_LARCH_PCALA_LO12                 R_LARCH = 72
	R_LARCH_PCALA64_LO20               R_LARCH = 73
	R_LARCH_PCALA64_HI12               R_LARCH = 74
	R_LARCH_GOT_PC_HI20                R_LARCH = 75
	R_LARCH_GOT_PC_LO12                R_LARCH = 76
	R_LARCH_GOT64_PC_LO20              R_LARCH = 77
	R_LARCH_GOT64_PC_HI12              R_LARCH = 78
	R_LARCH_GOT_HI20                   R_LARCH = 79
	R_LARCH_GOT_LO12                   R_LARCH = 80
	R_LARCH_GOT64_LO20                 R_LARCH = 81
	R_LARCH_GOT64_HI12                 R_LARCH = 82
	R_LARCH_TLS_LE_HI20                R_LARCH = 83
	R_LARCH_TLS_LE_LO12                R_LARCH = 84
	R_LARCH_TLS_LE64_LO20              R_LARCH = 85
	R_LARCH_TLS_LE64_HI12              R_LARCH = 86
	R_LARCH_TLS_IE_PC_HI20             R_LARCH = 87
	R_LARCH_TLS_IE_PC_LO12             R_LARCH = 88
	R_LARCH_TLS_IE64_PC_LO20           R_LARCH = 89
	R_LARCH_TLS_IE64_PC_HI12           R_LARCH = 90
	R_LARCH_TLS_IE_HI20                R_LARCH = 91
	R_LARCH_TLS_IE_LO12                R_LARCH = 92
	R_LARCH_TLS_IE64_LO20              R_LARCH = 93
	R_LARCH_TLS_IE64_HI12              R_LARCH = 94
	R_LARCH_TLS_LD_PC_HI20             R_LARCH = 95
	R_LARCH_TLS_LD_HI20                R_LARCH = 96
	R_LARCH_TLS_GD_PC_HI20             R_LARCH = 97
	R_LARCH_TLS_GD_HI20                R_LARCH = 98
	R_LARCH_32_PCREL                   R_LARCH = 99
	R_LARCH_RELAX                      R_LARCH = 100
	R_LARCH_DELETE                     R_LARCH = 101
	R_LARCH_ALIGN                      R_LARCH = 102
	R_LARCH_PCREL20_S2                 R_LARCH = 103
	R_LARCH_CFA                        R_LARCH = 104
	R_LARCH_ADD6                       R_LARCH = 105
	R_LARCH_SUB6                       R_LARCH = 106
	R_LARCH_ADD_ULEB128                R_LARCH = 107
	R_LARCH_SUB_ULEB128                R_LARCH = 108
	R_LARCH_64_PCREL                   R_LARCH = 109
)

func (i R_LARCH) String() string
func (i R_LARCH) GoString() string

// Relocation types for PowerPC.
//
// Values that are shared by both R_PPC and R_PPC64 are prefixed with
// R_POWERPC_ in the ELF standard. For the R_PPC type, the relevant
// shared relocations have been renamed with the prefix R_PPC_.
// The original name follows the value in a comment.
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
//
// Values that are shared by both R_PPC and R_PPC64 are prefixed with
// R_POWERPC_ in the ELF standard. For the R_PPC64 type, the relevant
// shared relocations have been renamed with the prefix R_PPC64_.
// The original name follows the value in a comment.
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
	R_PPC64_COPY               R_PPC64 = 19
	R_PPC64_GLOB_DAT           R_PPC64 = 20
	R_PPC64_JMP_SLOT           R_PPC64 = 21
	R_PPC64_RELATIVE           R_PPC64 = 22
	R_PPC64_UADDR32            R_PPC64 = 24
	R_PPC64_UADDR16            R_PPC64 = 25
	R_PPC64_REL32              R_PPC64 = 26
	R_PPC64_PLT32              R_PPC64 = 27
	R_PPC64_PLTREL32           R_PPC64 = 28
	R_PPC64_PLT16_LO           R_PPC64 = 29
	R_PPC64_PLT16_HI           R_PPC64 = 30
	R_PPC64_PLT16_HA           R_PPC64 = 31
	R_PPC64_SECTOFF            R_PPC64 = 33
	R_PPC64_SECTOFF_LO         R_PPC64 = 34
	R_PPC64_SECTOFF_HI         R_PPC64 = 35
	R_PPC64_SECTOFF_HA         R_PPC64 = 36
	R_PPC64_REL30              R_PPC64 = 37
	R_PPC64_ADDR64             R_PPC64 = 38
	R_PPC64_ADDR16_HIGHER      R_PPC64 = 39
	R_PPC64_ADDR16_HIGHERA     R_PPC64 = 40
	R_PPC64_ADDR16_HIGHEST     R_PPC64 = 41
	R_PPC64_ADDR16_HIGHESTA    R_PPC64 = 42
	R_PPC64_UADDR64            R_PPC64 = 43
	R_PPC64_REL64              R_PPC64 = 44
	R_PPC64_PLT64              R_PPC64 = 45
	R_PPC64_PLTREL64           R_PPC64 = 46
	R_PPC64_TOC16              R_PPC64 = 47
	R_PPC64_TOC16_LO           R_PPC64 = 48
	R_PPC64_TOC16_HI           R_PPC64 = 49
	R_PPC64_TOC16_HA           R_PPC64 = 50
	R_PPC64_TOC                R_PPC64 = 51
	R_PPC64_PLTGOT16           R_PPC64 = 52
	R_PPC64_PLTGOT16_LO        R_PPC64 = 53
	R_PPC64_PLTGOT16_HI        R_PPC64 = 54
	R_PPC64_PLTGOT16_HA        R_PPC64 = 55
	R_PPC64_ADDR16_DS          R_PPC64 = 56
	R_PPC64_ADDR16_LO_DS       R_PPC64 = 57
	R_PPC64_GOT16_DS           R_PPC64 = 58
	R_PPC64_GOT16_LO_DS        R_PPC64 = 59
	R_PPC64_PLT16_LO_DS        R_PPC64 = 60
	R_PPC64_SECTOFF_DS         R_PPC64 = 61
	R_PPC64_SECTOFF_LO_DS      R_PPC64 = 62
	R_PPC64_TOC16_DS           R_PPC64 = 63
	R_PPC64_TOC16_LO_DS        R_PPC64 = 64
	R_PPC64_PLTGOT16_DS        R_PPC64 = 65
	R_PPC64_PLTGOT_LO_DS       R_PPC64 = 66
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
	R_PPC64_TOCSAVE            R_PPC64 = 109
	R_PPC64_ADDR16_HIGH        R_PPC64 = 110
	R_PPC64_ADDR16_HIGHA       R_PPC64 = 111
	R_PPC64_TPREL16_HIGH       R_PPC64 = 112
	R_PPC64_TPREL16_HIGHA      R_PPC64 = 113
	R_PPC64_DTPREL16_HIGH      R_PPC64 = 114
	R_PPC64_DTPREL16_HIGHA     R_PPC64 = 115
	R_PPC64_REL24_NOTOC        R_PPC64 = 116
	R_PPC64_ADDR64_LOCAL       R_PPC64 = 117
	R_PPC64_ENTRY              R_PPC64 = 118
	R_PPC64_PLTSEQ             R_PPC64 = 119
	R_PPC64_PLTCALL            R_PPC64 = 120
	R_PPC64_PLTSEQ_NOTOC       R_PPC64 = 121
	R_PPC64_PLTCALL_NOTOC      R_PPC64 = 122
	R_PPC64_PCREL_OPT          R_PPC64 = 123
	R_PPC64_REL24_P9NOTOC      R_PPC64 = 124
	R_PPC64_D34                R_PPC64 = 128
	R_PPC64_D34_LO             R_PPC64 = 129
	R_PPC64_D34_HI30           R_PPC64 = 130
	R_PPC64_D34_HA30           R_PPC64 = 131
	R_PPC64_PCREL34            R_PPC64 = 132
	R_PPC64_GOT_PCREL34        R_PPC64 = 133
	R_PPC64_PLT_PCREL34        R_PPC64 = 134
	R_PPC64_PLT_PCREL34_NOTOC  R_PPC64 = 135
	R_PPC64_ADDR16_HIGHER34    R_PPC64 = 136
	R_PPC64_ADDR16_HIGHERA34   R_PPC64 = 137
	R_PPC64_ADDR16_HIGHEST34   R_PPC64 = 138
	R_PPC64_ADDR16_HIGHESTA34  R_PPC64 = 139
	R_PPC64_REL16_HIGHER34     R_PPC64 = 140
	R_PPC64_REL16_HIGHERA34    R_PPC64 = 141
	R_PPC64_REL16_HIGHEST34    R_PPC64 = 142
	R_PPC64_REL16_HIGHESTA34   R_PPC64 = 143
	R_PPC64_D28                R_PPC64 = 144
	R_PPC64_PCREL28            R_PPC64 = 145
	R_PPC64_TPREL34            R_PPC64 = 146
	R_PPC64_DTPREL34           R_PPC64 = 147
	R_PPC64_GOT_TLSGD_PCREL34  R_PPC64 = 148
	R_PPC64_GOT_TLSLD_PCREL34  R_PPC64 = 149
	R_PPC64_GOT_TPREL_PCREL34  R_PPC64 = 150
	R_PPC64_GOT_DTPREL_PCREL34 R_PPC64 = 151
	R_PPC64_REL16_HIGH         R_PPC64 = 240
	R_PPC64_REL16_HIGHA        R_PPC64 = 241
	R_PPC64_REL16_HIGHER       R_PPC64 = 242
	R_PPC64_REL16_HIGHERA      R_PPC64 = 243
	R_PPC64_REL16_HIGHEST      R_PPC64 = 244
	R_PPC64_REL16_HIGHESTA     R_PPC64 = 245
	R_PPC64_REL16DX_HA         R_PPC64 = 246
	R_PPC64_JMP_IREL           R_PPC64 = 247
	R_PPC64_IRELATIVE          R_PPC64 = 248
	R_PPC64_REL16              R_PPC64 = 249
	R_PPC64_REL16_LO           R_PPC64 = 250
	R_PPC64_REL16_HI           R_PPC64 = 251
	R_PPC64_REL16_HA           R_PPC64 = 252
	R_PPC64_GNU_VTINHERIT      R_PPC64 = 253
	R_PPC64_GNU_VTENTRY        R_PPC64 = 254
)

func (i R_PPC64) String() string
func (i R_PPC64) GoString() string

// Relocation types for RISC-V processors.
type R_RISCV int

const (
	R_RISCV_NONE          R_RISCV = 0
	R_RISCV_32            R_RISCV = 1
	R_RISCV_64            R_RISCV = 2
	R_RISCV_RELATIVE      R_RISCV = 3
	R_RISCV_COPY          R_RISCV = 4
	R_RISCV_JUMP_SLOT     R_RISCV = 5
	R_RISCV_TLS_DTPMOD32  R_RISCV = 6
	R_RISCV_TLS_DTPMOD64  R_RISCV = 7
	R_RISCV_TLS_DTPREL32  R_RISCV = 8
	R_RISCV_TLS_DTPREL64  R_RISCV = 9
	R_RISCV_TLS_TPREL32   R_RISCV = 10
	R_RISCV_TLS_TPREL64   R_RISCV = 11
	R_RISCV_BRANCH        R_RISCV = 16
	R_RISCV_JAL           R_RISCV = 17
	R_RISCV_CALL          R_RISCV = 18
	R_RISCV_CALL_PLT      R_RISCV = 19
	R_RISCV_GOT_HI20      R_RISCV = 20
	R_RISCV_TLS_GOT_HI20  R_RISCV = 21
	R_RISCV_TLS_GD_HI20   R_RISCV = 22
	R_RISCV_PCREL_HI20    R_RISCV = 23
	R_RISCV_PCREL_LO12_I  R_RISCV = 24
	R_RISCV_PCREL_LO12_S  R_RISCV = 25
	R_RISCV_HI20          R_RISCV = 26
	R_RISCV_LO12_I        R_RISCV = 27
	R_RISCV_LO12_S        R_RISCV = 28
	R_RISCV_TPREL_HI20    R_RISCV = 29
	R_RISCV_TPREL_LO12_I  R_RISCV = 30
	R_RISCV_TPREL_LO12_S  R_RISCV = 31
	R_RISCV_TPREL_ADD     R_RISCV = 32
	R_RISCV_ADD8          R_RISCV = 33
	R_RISCV_ADD16         R_RISCV = 34
	R_RISCV_ADD32         R_RISCV = 35
	R_RISCV_ADD64         R_RISCV = 36
	R_RISCV_SUB8          R_RISCV = 37
	R_RISCV_SUB16         R_RISCV = 38
	R_RISCV_SUB32         R_RISCV = 39
	R_RISCV_SUB64         R_RISCV = 40
	R_RISCV_GNU_VTINHERIT R_RISCV = 41
	R_RISCV_GNU_VTENTRY   R_RISCV = 42
	R_RISCV_ALIGN         R_RISCV = 43
	R_RISCV_RVC_BRANCH    R_RISCV = 44
	R_RISCV_RVC_JUMP      R_RISCV = 45
	R_RISCV_RVC_LUI       R_RISCV = 46
	R_RISCV_GPREL_I       R_RISCV = 47
	R_RISCV_GPREL_S       R_RISCV = 48
	R_RISCV_TPREL_I       R_RISCV = 49
	R_RISCV_TPREL_S       R_RISCV = 50
	R_RISCV_RELAX         R_RISCV = 51
	R_RISCV_SUB6          R_RISCV = 52
	R_RISCV_SET6          R_RISCV = 53
	R_RISCV_SET8          R_RISCV = 54
	R_RISCV_SET16         R_RISCV = 55
	R_RISCV_SET32         R_RISCV = 56
	R_RISCV_32_PCREL      R_RISCV = 57
)

func (i R_RISCV) String() string
func (i R_RISCV) GoString() string

// Relocation types for s390x processors.
type R_390 int

const (
	R_390_NONE        R_390 = 0
	R_390_8           R_390 = 1
	R_390_12          R_390 = 2
	R_390_16          R_390 = 3
	R_390_32          R_390 = 4
	R_390_PC32        R_390 = 5
	R_390_GOT12       R_390 = 6
	R_390_GOT32       R_390 = 7
	R_390_PLT32       R_390 = 8
	R_390_COPY        R_390 = 9
	R_390_GLOB_DAT    R_390 = 10
	R_390_JMP_SLOT    R_390 = 11
	R_390_RELATIVE    R_390 = 12
	R_390_GOTOFF      R_390 = 13
	R_390_GOTPC       R_390 = 14
	R_390_GOT16       R_390 = 15
	R_390_PC16        R_390 = 16
	R_390_PC16DBL     R_390 = 17
	R_390_PLT16DBL    R_390 = 18
	R_390_PC32DBL     R_390 = 19
	R_390_PLT32DBL    R_390 = 20
	R_390_GOTPCDBL    R_390 = 21
	R_390_64          R_390 = 22
	R_390_PC64        R_390 = 23
	R_390_GOT64       R_390 = 24
	R_390_PLT64       R_390 = 25
	R_390_GOTENT      R_390 = 26
	R_390_GOTOFF16    R_390 = 27
	R_390_GOTOFF64    R_390 = 28
	R_390_GOTPLT12    R_390 = 29
	R_390_GOTPLT16    R_390 = 30
	R_390_GOTPLT32    R_390 = 31
	R_390_GOTPLT64    R_390 = 32
	R_390_GOTPLTENT   R_390 = 33
	R_390_GOTPLTOFF16 R_390 = 34
	R_390_GOTPLTOFF32 R_390 = 35
	R_390_GOTPLTOFF64 R_390 = 36
	R_390_TLS_LOAD    R_390 = 37
	R_390_TLS_GDCALL  R_390 = 38
	R_390_TLS_LDCALL  R_390 = 39
	R_390_TLS_GD32    R_390 = 40
	R_390_TLS_GD64    R_390 = 41
	R_390_TLS_GOTIE12 R_390 = 42
	R_390_TLS_GOTIE32 R_390 = 43
	R_390_TLS_GOTIE64 R_390 = 44
	R_390_TLS_LDM32   R_390 = 45
	R_390_TLS_LDM64   R_390 = 46
	R_390_TLS_IE32    R_390 = 47
	R_390_TLS_IE64    R_390 = 48
	R_390_TLS_IEENT   R_390 = 49
	R_390_TLS_LE32    R_390 = 50
	R_390_TLS_LE64    R_390 = 51
	R_390_TLS_LDO32   R_390 = 52
	R_390_TLS_LDO64   R_390 = 53
	R_390_TLS_DTPMOD  R_390 = 54
	R_390_TLS_DTPOFF  R_390 = 55
	R_390_TLS_TPOFF   R_390 = 56
	R_390_20          R_390 = 57
	R_390_GOT20       R_390 = 58
	R_390_GOTPLT20    R_390 = 59
	R_390_TLS_GOTIE20 R_390 = 60
)

func (i R_390) String() string
func (i R_390) GoString() string

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

// ELF32 Dynamic structure. The ".dynamic" section contains an array of them.
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

// ELF64 Dynamic structure. The ".dynamic" section contains an array of them.
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
