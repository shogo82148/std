// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xcoff

// File Header.
type FileHeader32 struct {
	Fmagic   uint16
	Fnscns   uint16
	Ftimedat uint32
	Fsymptr  uint32
	Fnsyms   uint32
	Fopthdr  uint16
	Fflags   uint16
}

type FileHeader64 struct {
	Fmagic   uint16
	Fnscns   uint16
	Ftimedat uint32
	Fsymptr  uint64
	Fopthdr  uint16
	Fflags   uint16
	Fnsyms   uint32
}

const (
	FILHSZ_32 = 20
	FILHSZ_64 = 24
)
const (
	U802TOCMAGIC = 0737
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

// Section Header.
type SectionHeader32 struct {
	Sname    [8]byte
	Spaddr   uint32
	Svaddr   uint32
	Ssize    uint32
	Sscnptr  uint32
	Srelptr  uint32
	Slnnoptr uint32
	Snreloc  uint16
	Snlnno   uint16
	Sflags   uint32
}

type SectionHeader64 struct {
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
	Spad     uint32
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

// Symbol Table Entry.
type SymEnt32 struct {
	Nname   [8]byte
	Nvalue  uint32
	Nscnum  uint16
	Ntype   uint16
	Nsclass uint8
	Nnumaux uint8
}

type SymEnt64 struct {
	Nvalue  uint64
	Noffset uint32
	Nscnum  uint16
	Ntype   uint16
	Nsclass uint8
	Nnumaux uint8
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
type AuxFile64 struct {
	Xfname   [8]byte
	Xftype   uint8
	Xauxtype uint8
}

// Function Auxiliary Entry
type AuxFcn32 struct {
	Xexptr   uint32
	Xfsize   uint32
	Xlnnoptr uint32
	Xendndx  uint32
	Xpad     uint16
}
type AuxFcn64 struct {
	Xlnnoptr uint64
	Xfsize   uint32
	Xendndx  uint32
	Xpad     uint8
	Xauxtype uint8
}

type AuxSect64 struct {
	Xscnlen  uint64
	Xnreloc  uint64
	pad      uint8
	Xauxtype uint8
}

// csect Auxiliary Entry.
type AuxCSect32 struct {
	Xscnlen   uint32
	Xparmhash uint32
	Xsnhash   uint16
	Xsmtyp    uint8
	Xsmclas   uint8
	Xstab     uint32
	Xsnstab   uint16
}

type AuxCSect64 struct {
	Xscnlenlo uint32
	Xparmhash uint32
	Xsnhash   uint16
	Xsmtyp    uint8
	Xsmclas   uint8
	Xscnlenhi uint32
	Xpad      uint8
	Xauxtype  uint8
}

// Symbol type field.
const (
	XTY_ER = 0
	XTY_SD = 1
	XTY_LD = 2
	XTY_CM = 3
)

// Defines for File auxiliary definitions: x_ftype field of x_file
const (
	XFT_FN = 0
	XFT_CT = 1
	XFT_CV = 2
	XFT_CD = 128
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

// Loader Header.
type LoaderHeader32 struct {
	Lversion uint32
	Lnsyms   uint32
	Lnreloc  uint32
	Listlen  uint32
	Lnimpid  uint32
	Limpoff  uint32
	Lstlen   uint32
	Lstoff   uint32
}

type LoaderHeader64 struct {
	Lversion uint32
	Lnsyms   uint32
	Lnreloc  uint32
	Listlen  uint32
	Lnimpid  uint32
	Lstlen   uint32
	Limpoff  uint64
	Lstoff   uint64
	Lsymoff  uint64
	Lrldoff  uint64
}

const (
	LDHDRSZ_32 = 32
	LDHDRSZ_64 = 56
)

// Loader Symbol.
type LoaderSymbol32 struct {
	Lname   [8]byte
	Lvalue  uint32
	Lscnum  uint16
	Lsmtype uint8
	Lsmclas uint8
	Lifile  uint32
	Lparm   uint32
}

type LoaderSymbol64 struct {
	Lvalue  uint64
	Loffset uint32
	Lscnum  uint16
	Lsmtype uint8
	Lsmclas uint8
	Lifile  uint32
	Lparm   uint32
}

type Reloc32 struct {
	Rvaddr  uint32
	Rsymndx uint32
	Rsize   uint8
	Rtype   uint8
}

type Reloc64 struct {
	Rvaddr  uint64
	Rsymndx uint32
	Rsize   uint8
	Rtype   uint8
}

const (
	R_POS = 0x00
	R_NEG = 0x01
	R_REL = 0x02
	R_TOC = 0x03
	R_TRL = 0x12

	R_TRLA = 0x13
	R_GL   = 0x05
	R_TCL  = 0x06
	R_RL   = 0x0C
	R_RLA  = 0x0D
	R_REF  = 0x0F
	R_BA   = 0x08
	R_RBA  = 0x18
	R_BR   = 0x0A
	R_RBR  = 0x1A

	R_TLS    = 0x20
	R_TLS_IE = 0x21
	R_TLS_LD = 0x22
	R_TLS_LE = 0x23
	R_TLSM   = 0x24
	R_TLSML  = 0x25

	R_TOCU = 0x30
	R_TOCL = 0x31
)
