// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package defines the Go object file format, and provide "low-level" functions
// for reading and writing object files.

// The object file is understood by the compiler, assembler, linker, and tools. They
// have "high level" code that operates on object files, handling application-specific
// logics, and use this package for the actual reading and writing. Specifically, the
// code below:
//
// - cmd/internal/obj/objfile.go (used by cmd/asm and cmd/compile)
// - cmd/internal/objfile/goobj.go (used cmd/nm, cmd/objdump)
// - cmd/link/internal/loader package (used by cmd/link)
//
// If the object file format changes, they may (or may not) need to change.

package goobj

import (
	"github.com/shogo82148/std/cmd/internal/bio"
)

type FingerprintType [8]byte

func (fp FingerprintType) IsZero() bool

// Package Index.
const (
	PkgIdxNone = (1<<31 - 1) - iota
	PkgIdxHashed64
	PkgIdxHashed
	PkgIdxBuiltin
	PkgIdxSelf
	PkgIdxSpecial = PkgIdxSelf
	PkgIdxInvalid = 0
)

// Blocks
const (
	BlkAutolib = iota
	BlkPkgIdx
	BlkFile
	BlkSymdef
	BlkHashed64def
	BlkHasheddef
	BlkNonpkgdef
	BlkNonpkgref
	BlkRefFlags
	BlkHash64
	BlkHash
	BlkRelocIdx
	BlkAuxIdx
	BlkDataIdx
	BlkReloc
	BlkAux
	BlkData
	BlkRefName
	BlkEnd
	NBlk
)

// File header.
// TODO: probably no need to export this.
type Header struct {
	Magic       string
	Fingerprint FingerprintType
	Flags       uint32
	Offsets     [NBlk]uint32
}

const Magic = "\x00go120ld"

func (h *Header) Write(w *Writer)

func (h *Header) Read(r *Reader) error

func (h *Header) Size() int

// Autolib
type ImportedPkg struct {
	Pkg         string
	Fingerprint FingerprintType
}

func (p *ImportedPkg) Write(w *Writer)

// Symbol definition.
//
// Serialized format:
//
//	Sym struct {
//	   Name  string
//	   ABI   uint16
//	   Type  uint8
//	   Flag  uint8
//	   Flag2 uint8
//	   Siz   uint32
//	   Align uint32
//	}
type Sym [SymSize]byte

const SymSize = stringRefSize + 2 + 1 + 1 + 1 + 4 + 4

const SymABIstatic = ^uint16(0)

const (
	ObjFlagShared = 1 << iota
	_
	ObjFlagFromAssembly
	ObjFlagUnlinkable
	ObjFlagStd
)

// Sym.Flag
const (
	SymFlagDupok = 1 << iota
	SymFlagLocal
	SymFlagTypelink
	SymFlagLeaf
	SymFlagNoSplit
	SymFlagReflectMethod
	SymFlagGoType
)

// Sym.Flag2
const (
	SymFlagUsedInIface = 1 << iota
	SymFlagItab
	SymFlagDict
	SymFlagPkgInit
	SymFlagLinkname
	SymFlagABIWrapper
	SymFlagWasmExport
)

// Returns the length of the name of the symbol.
func (s *Sym) NameLen(r *Reader) int

func (s *Sym) Name(r *Reader) string

func (s *Sym) ABI() uint16
func (s *Sym) Type() uint8
func (s *Sym) Flag() uint8
func (s *Sym) Flag2() uint8
func (s *Sym) Siz() uint32
func (s *Sym) Align() uint32

func (s *Sym) Dupok() bool
func (s *Sym) Local() bool
func (s *Sym) Typelink() bool
func (s *Sym) Leaf() bool
func (s *Sym) NoSplit() bool
func (s *Sym) ReflectMethod() bool
func (s *Sym) IsGoType() bool
func (s *Sym) UsedInIface() bool
func (s *Sym) IsItab() bool
func (s *Sym) IsDict() bool
func (s *Sym) IsPkgInit() bool
func (s *Sym) IsLinkname() bool
func (s *Sym) ABIWrapper() bool
func (s *Sym) WasmExport() bool

func (s *Sym) SetName(x string, w *Writer)

func (s *Sym) SetABI(x uint16)
func (s *Sym) SetType(x uint8)
func (s *Sym) SetFlag(x uint8)
func (s *Sym) SetFlag2(x uint8)
func (s *Sym) SetSiz(x uint32)
func (s *Sym) SetAlign(x uint32)

func (s *Sym) Write(w *Writer)

// Symbol reference.
type SymRef struct {
	PkgIdx uint32
	SymIdx uint32
}

func (s SymRef) IsZero() bool

// Hash64
type Hash64Type [Hash64Size]byte

const Hash64Size = 8

// Hash
type HashType [HashSize]byte

const HashSize = 16

// Relocation.
//
// Serialized format:
//
//	Reloc struct {
//	   Off  int32
//	   Siz  uint8
//	   Type uint16
//	   Add  int64
//	   Sym  SymRef
//	}
type Reloc [RelocSize]byte

const RelocSize = 4 + 1 + 2 + 8 + 8

func (r *Reloc) Off() int32
func (r *Reloc) Siz() uint8
func (r *Reloc) Type() uint16
func (r *Reloc) Add() int64
func (r *Reloc) Sym() SymRef

func (r *Reloc) SetOff(x int32)
func (r *Reloc) SetSiz(x uint8)
func (r *Reloc) SetType(x uint16)
func (r *Reloc) SetAdd(x int64)
func (r *Reloc) SetSym(x SymRef)

func (r *Reloc) Set(off int32, size uint8, typ uint16, add int64, sym SymRef)

func (r *Reloc) Write(w *Writer)

// Aux symbol info.
//
// Serialized format:
//
//	Aux struct {
//	   Type uint8
//	   Sym  SymRef
//	}
type Aux [AuxSize]byte

const AuxSize = 1 + 8

// Aux Type
const (
	AuxGotype = iota
	AuxFuncInfo
	AuxFuncdata
	AuxDwarfInfo
	AuxDwarfLoc
	AuxDwarfRanges
	AuxDwarfLines
	AuxPcsp
	AuxPcfile
	AuxPcline
	AuxPcinline
	AuxPcdata
	AuxWasmImport
	AuxWasmType
	AuxSehUnwindInfo
)

func (a *Aux) Type() uint8
func (a *Aux) Sym() SymRef

func (a *Aux) SetType(x uint8)
func (a *Aux) SetSym(x SymRef)

func (a *Aux) Write(w *Writer)

// Referenced symbol flags.
//
// Serialized format:
//
//	RefFlags struct {
//	   Sym   symRef
//	   Flag  uint8
//	   Flag2 uint8
//	}
type RefFlags [RefFlagsSize]byte

const RefFlagsSize = 8 + 1 + 1

func (r *RefFlags) Sym() SymRef

func (r *RefFlags) Flag() uint8
func (r *RefFlags) Flag2() uint8

func (r *RefFlags) SetSym(x SymRef)

func (r *RefFlags) SetFlag(x uint8)
func (r *RefFlags) SetFlag2(x uint8)

func (r *RefFlags) Write(w *Writer)

// Referenced symbol name.
//
// Serialized format:
//
//	RefName struct {
//	   Sym  symRef
//	   Name string
//	}
type RefName [RefNameSize]byte

const RefNameSize = 8 + stringRefSize

func (n *RefName) Sym() SymRef

func (n *RefName) Name(r *Reader) string

func (n *RefName) SetSym(x SymRef)

func (n *RefName) SetName(x string, w *Writer)

func (n *RefName) Write(w *Writer)

type Writer struct {
	wr        *bio.Writer
	stringMap map[string]uint32
	off       uint32

	b [8]byte
}

func NewWriter(wr *bio.Writer) *Writer

func (w *Writer) AddString(s string)

func (w *Writer) StringRef(s string)

func (w *Writer) RawString(s string)

func (w *Writer) Bytes(s []byte)

func (w *Writer) Uint64(x uint64)

func (w *Writer) Uint32(x uint32)

func (w *Writer) Uint16(x uint16)

func (w *Writer) Uint8(x uint8)

func (w *Writer) Offset() uint32

type Reader struct {
	b        []byte
	readonly bool

	start uint32
	h     Header
}

func NewReaderFromBytes(b []byte, readonly bool) *Reader

func (r *Reader) BytesAt(off uint32, len int) []byte

func (r *Reader) StringAt(off uint32, len uint32) string

func (r *Reader) StringRef(off uint32) string

func (r *Reader) Fingerprint() FingerprintType

func (r *Reader) Autolib() []ImportedPkg

func (r *Reader) Pkglist() []string

func (r *Reader) NPkg() int

func (r *Reader) Pkg(i int) string

func (r *Reader) NFile() int

func (r *Reader) File(i int) string

func (r *Reader) NSym() int

func (r *Reader) NHashed64def() int

func (r *Reader) NHasheddef() int

func (r *Reader) NNonpkgdef() int

func (r *Reader) NNonpkgref() int

// SymOff returns the offset of the i-th symbol.
func (r *Reader) SymOff(i uint32) uint32

// Sym returns a pointer to the i-th symbol.
func (r *Reader) Sym(i uint32) *Sym

// NRefFlags returns the number of referenced symbol flags.
func (r *Reader) NRefFlags() int

// RefFlags returns a pointer to the i-th referenced symbol flags.
// Note: here i is not a local symbol index, just a counter.
func (r *Reader) RefFlags(i int) *RefFlags

// Hash64 returns the i-th short hashed symbol's hash.
// Note: here i is the index of short hashed symbols, not all symbols
// (unlike other accessors).
func (r *Reader) Hash64(i uint32) uint64

// Hash returns a pointer to the i-th hashed symbol's hash.
// Note: here i is the index of hashed symbols, not all symbols
// (unlike other accessors).
func (r *Reader) Hash(i uint32) *HashType

// NReloc returns the number of relocations of the i-th symbol.
func (r *Reader) NReloc(i uint32) int

// RelocOff returns the offset of the j-th relocation of the i-th symbol.
func (r *Reader) RelocOff(i uint32, j int) uint32

// Reloc returns a pointer to the j-th relocation of the i-th symbol.
func (r *Reader) Reloc(i uint32, j int) *Reloc

// Relocs returns a pointer to the relocations of the i-th symbol.
func (r *Reader) Relocs(i uint32) []Reloc

// NAux returns the number of aux symbols of the i-th symbol.
func (r *Reader) NAux(i uint32) int

// AuxOff returns the offset of the j-th aux symbol of the i-th symbol.
func (r *Reader) AuxOff(i uint32, j int) uint32

// Aux returns a pointer to the j-th aux symbol of the i-th symbol.
func (r *Reader) Aux(i uint32, j int) *Aux

// Auxs returns the aux symbols of the i-th symbol.
func (r *Reader) Auxs(i uint32) []Aux

// DataOff returns the offset of the i-th symbol's data.
func (r *Reader) DataOff(i uint32) uint32

// DataSize returns the size of the i-th symbol's data.
func (r *Reader) DataSize(i uint32) int

// Data returns the i-th symbol's data.
func (r *Reader) Data(i uint32) []byte

// DataString returns the i-th symbol's data as a string.
func (r *Reader) DataString(i uint32) string

// NRefName returns the number of referenced symbol names.
func (r *Reader) NRefName() int

// RefName returns a pointer to the i-th referenced symbol name.
// Note: here i is not a local symbol index, just a counter.
func (r *Reader) RefName(i int) *RefName

// ReadOnly returns whether r.BytesAt returns read-only bytes.
func (r *Reader) ReadOnly() bool

// Flags returns the flag bits read from the object file header.
func (r *Reader) Flags() uint32

func (r *Reader) Shared() bool
func (r *Reader) FromAssembly() bool
func (r *Reader) Unlinkable() bool
func (r *Reader) Std() bool
