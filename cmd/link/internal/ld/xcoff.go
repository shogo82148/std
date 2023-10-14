// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

const (
	// Total amount of space to reserve at the start of the file
	// for File Header, Auxiliary Header, and Section Headers.
	// May waste some.
	XCOFFHDRRESERVE = FILHSZ_64 + AOUTHSZ_EXEC64 + SCNHSZ_64*23

	// base on dump -o, then rounded from 32B to 64B to
	// match worst case elf text section alignment on ppc64.
	XCOFFSECTALIGN int64 = 64

	// XCOFF binaries should normally have all its sections position-independent.
	// However, this is not yet possible for .text because of some R_ADDR relocations
	// inside RODATA symbols.
	// .data and .bss are position-independent so their address start inside an unreachable
	// segment during execution to force segfault if something is wrong.
	XCOFFTEXTBASE = 0x100000000
	XCOFFDATABASE = 0x200000000
)

// File Header
type XcoffFileHdr64 struct {
	Fmagic   uint16
	Fnscns   uint16
	Ftimedat int32
	Fsymptr  uint64
	Fopthdr  uint16
	Fflags   uint16
	Fnsyms   int32
}

const (
	U64_TOCMAGIC = 0767
)

// Flags that describe the type of the object file.
const (
	F_RELFLG    = 0x0001
	F_EXEC      = 0x0002
	F_LNNO      = 0x0004
	F_FDPR_PROF = 0x0010
	F_FDPR_OPTI = 0x0020
	F_DSA       = 0x0040
	F_VARPG     = 0x0100
	F_DYNLOAD   = 0x1000
	F_SHROBJ    = 0x2000
	F_LOADONLY  = 0x4000
)

// Auxiliary Header
type XcoffAoutHdr64 struct {
	Omagic      int16
	Ovstamp     int16
	Odebugger   uint32
	Otextstart  uint64
	Odatastart  uint64
	Otoc        uint64
	Osnentry    int16
	Osntext     int16
	Osndata     int16
	Osntoc      int16
	Osnloader   int16
	Osnbss      int16
	Oalgntext   int16
	Oalgndata   int16
	Omodtype    [2]byte
	Ocpuflag    uint8
	Ocputype    uint8
	Otextpsize  uint8
	Odatapsize  uint8
	Ostackpsize uint8
	Oflags      uint8
	Otsize      uint64
	Odsize      uint64
	Obsize      uint64
	Oentry      uint64
	Omaxstack   uint64
	Omaxdata    uint64
	Osntdata    int16
	Osntbss     int16
	Ox64flags   uint16
	Oresv3a     int16
	Oresv3      [2]int32
}

// Section Header
type XcoffScnHdr64 struct {
	Sname    [8]byte
	Spaddr   uint64
	Svaddr   uint64
	Ssize    uint64
	Sscnptr  uint64
	Srelptr  uint64
	Slnnoptr uint64
	Snreloc  uint32
	Snlnno   uint32
	Sflags   uint32
}

// Flags defining the section type.
const (
	STYP_DWARF  = 0x0010
	STYP_TEXT   = 0x0020
	STYP_DATA   = 0x0040
	STYP_BSS    = 0x0080
	STYP_EXCEPT = 0x0100
	STYP_INFO   = 0x0200
	STYP_TDATA  = 0x0400
	STYP_TBSS   = 0x0800
	STYP_LOADER = 0x1000
	STYP_DEBUG  = 0x2000
	STYP_TYPCHK = 0x4000
	STYP_OVRFLO = 0x8000
)
const (
	SSUBTYP_DWINFO  = 0x10000
	SSUBTYP_DWLINE  = 0x20000
	SSUBTYP_DWPBNMS = 0x30000
	SSUBTYP_DWPBTYP = 0x40000
	SSUBTYP_DWARNGE = 0x50000
	SSUBTYP_DWABREV = 0x60000
	SSUBTYP_DWSTR   = 0x70000
	SSUBTYP_DWRNGES = 0x80000
	SSUBTYP_DWLOC   = 0x90000
	SSUBTYP_DWFRAME = 0xA0000
	SSUBTYP_DWMAC   = 0xB0000
)

// Headers size
const (
	FILHSZ_32      = 20
	FILHSZ_64      = 24
	AOUTHSZ_EXEC32 = 72
	AOUTHSZ_EXEC64 = 120
	SCNHSZ_32      = 40
	SCNHSZ_64      = 72
	LDHDRSZ_32     = 32
	LDHDRSZ_64     = 56
	LDSYMSZ_64     = 24
	RELSZ_64       = 14
)

// Symbol Table Entry
type XcoffSymEnt64 struct {
	Nvalue  uint64
	Noffset uint32
	Nscnum  int16
	Ntype   uint16
	Nsclass uint8
	Nnumaux int8
}

const SYMESZ = 18

const (
	// Nscnum
	N_DEBUG = -2
	N_ABS   = -1
	N_UNDEF = 0

	//Ntype
	SYM_V_INTERNAL  = 0x1000
	SYM_V_HIDDEN    = 0x2000
	SYM_V_PROTECTED = 0x3000
	SYM_V_EXPORTED  = 0x4000
	SYM_TYPE_FUNC   = 0x0020
)

// Storage Class.
const (
	C_NULL    = 0
	C_EXT     = 2
	C_STAT    = 3
	C_BLOCK   = 100
	C_FCN     = 101
	C_FILE    = 103
	C_HIDEXT  = 107
	C_BINCL   = 108
	C_EINCL   = 109
	C_WEAKEXT = 111
	C_DWARF   = 112
	C_GSYM    = 128
	C_LSYM    = 129
	C_PSYM    = 130
	C_RSYM    = 131
	C_RPSYM   = 132
	C_STSYM   = 133
	C_BCOMM   = 135
	C_ECOML   = 136
	C_ECOMM   = 137
	C_DECL    = 140
	C_ENTRY   = 141
	C_FUN     = 142
	C_BSTAT   = 143
	C_ESTAT   = 144
	C_GTLS    = 145
	C_STTLS   = 146
)

// File Auxiliary Entry
type XcoffAuxFile64 struct {
	Xzeroes  uint32
	Xoffset  uint32
	X_pad1   [6]byte
	Xftype   uint8
	X_pad2   [2]byte
	Xauxtype uint8
}

// Function Auxiliary Entry
type XcoffAuxFcn64 struct {
	Xlnnoptr uint64
	Xfsize   uint32
	Xendndx  uint32
	Xpad     uint8
	Xauxtype uint8
}

// csect Auxiliary Entry.
type XcoffAuxCSect64 struct {
	Xscnlenlo uint32
	Xparmhash uint32
	Xsnhash   uint16
	Xsmtyp    uint8
	Xsmclas   uint8
	Xscnlenhi uint32
	Xpad      uint8
	Xauxtype  uint8
}

// DWARF Auxiliary Entry
type XcoffAuxDWARF64 struct {
	Xscnlen  uint64
	X_pad    [9]byte
	Xauxtype uint8
}

// Xftype field
const (
	XFT_FN = 0
	XFT_CT = 1
	XFT_CV = 2
	XFT_CD = 128
)

// Symbol type field.
const (
	XTY_ER  = 0
	XTY_SD  = 1
	XTY_LD  = 2
	XTY_CM  = 3
	XTY_WK  = 0x8
	XTY_EXP = 0x10
	XTY_ENT = 0x20
	XTY_IMP = 0x40
)

// Storage-mapping class.
const (
	XMC_PR     = 0
	XMC_RO     = 1
	XMC_DB     = 2
	XMC_TC     = 3
	XMC_UA     = 4
	XMC_RW     = 5
	XMC_GL     = 6
	XMC_XO     = 7
	XMC_SV     = 8
	XMC_BS     = 9
	XMC_DS     = 10
	XMC_UC     = 11
	XMC_TC0    = 15
	XMC_TD     = 16
	XMC_SV64   = 17
	XMC_SV3264 = 18
	XMC_TL     = 20
	XMC_UL     = 21
	XMC_TE     = 22
)

// Loader Header
type XcoffLdHdr64 struct {
	Lversion int32
	Lnsyms   int32
	Lnreloc  int32
	Listlen  uint32
	Lnimpid  int32
	Lstlen   uint32
	Limpoff  uint64
	Lstoff   uint64
	Lsymoff  uint64
	Lrldoff  uint64
}

// Loader Symbol
type XcoffLdSym64 struct {
	Lvalue  uint64
	Loffset uint32
	Lscnum  int16
	Lsmtype int8
	Lsmclas int8
	Lifile  int32
	Lparm   uint32
}

type XcoffLdImportFile64 struct {
	Limpidpath string
	Limpidbase string
	Limpidmem  string
}

type XcoffLdRel64 struct {
	Lvaddr  uint64
	Lrtype  uint16
	Lrsecnm int16
	Lsymndx int32
}

const (
	XCOFF_R_POS = 0x00
	XCOFF_R_NEG = 0x01
	XCOFF_R_REL = 0x02
	XCOFF_R_TOC = 0x03
	XCOFF_R_TRL = 0x12

	XCOFF_R_TRLA = 0x13
	XCOFF_R_GL   = 0x05
	XCOFF_R_TCL  = 0x06
	XCOFF_R_RL   = 0x0C
	XCOFF_R_RLA  = 0x0D
	XCOFF_R_REF  = 0x0F
	XCOFF_R_BA   = 0x08
	XCOFF_R_RBA  = 0x18
	XCOFF_R_BR   = 0x0A
	XCOFF_R_RBR  = 0x1A

	XCOFF_R_TLS    = 0x20
	XCOFF_R_TLS_IE = 0x21
	XCOFF_R_TLS_LD = 0x22
	XCOFF_R_TLS_LE = 0x23
	XCOFF_R_TLSM   = 0x24
	XCOFF_R_TLSML  = 0x25

	XCOFF_R_TOCU = 0x30
	XCOFF_R_TOCL = 0x31
)

type XcoffLdStr64 struct {
	size uint16
	name string
}

// Xcoffinit initialised some internal value and setups
// already known header information.
func Xcoffinit(ctxt *Link)

// Xcoffadddynrel adds a dynamic relocation in a XCOFF file.
// This relocation will be made by the loader.
func Xcoffadddynrel(target *Target, ldr *loader.Loader, syms *ArchSyms, s loader.Sym, r loader.Reloc, rIdx int) bool

// Create loader section and returns its size.
func Loaderblk(ctxt *Link, off uint64)
