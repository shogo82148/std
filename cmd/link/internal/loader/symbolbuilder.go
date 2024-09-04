// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loader

import (
	"github.com/shogo82148/std/cmd/internal/objabi"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/sym"
)

// SymbolBuilder is a helper designed to help with the construction
// of new symbol contents.
type SymbolBuilder struct {
	*extSymPayload
	symIdx Sym
	l      *Loader
}

// MakeSymbolBuilder creates a symbol builder for use in constructing
// an entirely new symbol.
func (l *Loader) MakeSymbolBuilder(name string) *SymbolBuilder

// MakeSymbolUpdater creates a symbol builder helper for an existing
// symbol 'symIdx'. If 'symIdx' is not an external symbol, then create
// a clone of it (copy name, properties, etc) fix things up so that
// the lookup tables and caches point to the new version, not the old
// version.
func (l *Loader) MakeSymbolUpdater(symIdx Sym) *SymbolBuilder

// CreateSymForUpdate creates a symbol with given name and version,
// returns a CreateSymForUpdate for update. If the symbol already
// exists, it will update in-place.
func (l *Loader) CreateSymForUpdate(name string, version int) *SymbolBuilder

func (sb *SymbolBuilder) Sym() Sym
func (sb *SymbolBuilder) Name() string
func (sb *SymbolBuilder) Version() int
func (sb *SymbolBuilder) Type() sym.SymKind
func (sb *SymbolBuilder) Size() int64
func (sb *SymbolBuilder) Data() []byte
func (sb *SymbolBuilder) Value() int64
func (sb *SymbolBuilder) Align() int32
func (sb *SymbolBuilder) Localentry() uint8
func (sb *SymbolBuilder) OnList() bool
func (sb *SymbolBuilder) External() bool
func (sb *SymbolBuilder) Extname() string
func (sb *SymbolBuilder) CgoExportDynamic() bool
func (sb *SymbolBuilder) Dynimplib() string
func (sb *SymbolBuilder) Dynimpvers() string
func (sb *SymbolBuilder) SubSym() Sym
func (sb *SymbolBuilder) GoType() Sym
func (sb *SymbolBuilder) VisibilityHidden() bool
func (sb *SymbolBuilder) Sect() *sym.Section

func (sb *SymbolBuilder) SetType(kind sym.SymKind)
func (sb *SymbolBuilder) SetSize(size int64)
func (sb *SymbolBuilder) SetData(data []byte)
func (sb *SymbolBuilder) SetOnList(v bool)
func (sb *SymbolBuilder) SetExternal(v bool)
func (sb *SymbolBuilder) SetValue(v int64)
func (sb *SymbolBuilder) SetAlign(align int32)
func (sb *SymbolBuilder) SetLocalentry(value uint8)
func (sb *SymbolBuilder) SetExtname(value string)
func (sb *SymbolBuilder) SetDynimplib(value string)
func (sb *SymbolBuilder) SetDynimpvers(value string)
func (sb *SymbolBuilder) SetPlt(value int32)
func (sb *SymbolBuilder) SetGot(value int32)
func (sb *SymbolBuilder) SetSpecial(value bool)
func (sb *SymbolBuilder) SetLocal(value bool)
func (sb *SymbolBuilder) SetVisibilityHidden(value bool)

func (sb *SymbolBuilder) SetNotInSymbolTable(value bool)

func (sb *SymbolBuilder) SetSect(sect *sym.Section)

func (sb *SymbolBuilder) AddBytes(data []byte)

func (sb *SymbolBuilder) Relocs() Relocs

// ResetRelocs removes all relocations on this symbol.
func (sb *SymbolBuilder) ResetRelocs()

// SetRelocType sets the type of the 'i'-th relocation on this sym to 't'
func (sb *SymbolBuilder) SetRelocType(i int, t objabi.RelocType)

// SetRelocSym sets the target sym of the 'i'-th relocation on this sym to 's'
func (sb *SymbolBuilder) SetRelocSym(i int, tgt Sym)

// SetRelocAdd sets the addend of the 'i'-th relocation on this sym to 'a'
func (sb *SymbolBuilder) SetRelocAdd(i int, a int64)

// Add n relocations, return a handle to the relocations.
func (sb *SymbolBuilder) AddRelocs(n int) Relocs

// Add a relocation with given type, return its handle and index
// (to set other fields).
func (sb *SymbolBuilder) AddRel(typ objabi.RelocType) (Reloc, int)

// SortRelocs Sort relocations by offset.
func (sb *SymbolBuilder) SortRelocs()

func (sb *SymbolBuilder) Reachable() bool

func (sb *SymbolBuilder) SetReachable(v bool)

func (sb *SymbolBuilder) ReadOnly() bool

func (sb *SymbolBuilder) SetReadOnly(v bool)

func (sb *SymbolBuilder) DuplicateOK() bool

func (sb *SymbolBuilder) SetDuplicateOK(v bool)

func (sb *SymbolBuilder) Outer() Sym

func (sb *SymbolBuilder) Sub() Sym

func (sb *SymbolBuilder) SortSub()

func (sb *SymbolBuilder) AddInteriorSym(sub Sym)

func (sb *SymbolBuilder) AddUint8(v uint8) int64

func (sb *SymbolBuilder) AddUintXX(arch *sys.Arch, v uint64, wid int) int64

func (sb *SymbolBuilder) AddUint16(arch *sys.Arch, v uint16) int64

func (sb *SymbolBuilder) AddUint32(arch *sys.Arch, v uint32) int64

func (sb *SymbolBuilder) AddUint64(arch *sys.Arch, v uint64) int64

func (sb *SymbolBuilder) AddUint(arch *sys.Arch, v uint64) int64

func (sb *SymbolBuilder) SetUint8(arch *sys.Arch, r int64, v uint8) int64

func (sb *SymbolBuilder) SetUint16(arch *sys.Arch, r int64, v uint16) int64

func (sb *SymbolBuilder) SetUint32(arch *sys.Arch, r int64, v uint32) int64

func (sb *SymbolBuilder) SetUint(arch *sys.Arch, r int64, v uint64) int64

func (sb *SymbolBuilder) SetUintptr(arch *sys.Arch, r int64, v uintptr) int64

func (sb *SymbolBuilder) SetAddrPlus(arch *sys.Arch, off int64, tgt Sym, add int64) int64

func (sb *SymbolBuilder) SetAddr(arch *sys.Arch, off int64, tgt Sym) int64

func (sb *SymbolBuilder) AddStringAt(off int64, str string) int64

// AddCStringAt adds str plus a null terminating byte.
func (sb *SymbolBuilder) AddCStringAt(off int64, str string) int64

func (sb *SymbolBuilder) Addstring(str string) int64

func (sb *SymbolBuilder) SetBytesAt(off int64, b []byte) int64

// Add a symbol reference (relocation) with given type, addend, and size
// (the most generic form).
func (sb *SymbolBuilder) AddSymRef(arch *sys.Arch, tgt Sym, add int64, typ objabi.RelocType, rsize int) int64

func (sb *SymbolBuilder) AddAddrPlus(arch *sys.Arch, tgt Sym, add int64) int64

func (sb *SymbolBuilder) AddAddrPlus4(arch *sys.Arch, tgt Sym, add int64) int64

func (sb *SymbolBuilder) AddAddr(arch *sys.Arch, tgt Sym) int64

func (sb *SymbolBuilder) AddPEImageRelativeAddrPlus(arch *sys.Arch, tgt Sym, add int64) int64

func (sb *SymbolBuilder) AddPCRelPlus(arch *sys.Arch, tgt Sym, add int64) int64

func (sb *SymbolBuilder) AddCURelativeAddrPlus(arch *sys.Arch, tgt Sym, add int64) int64

func (sb *SymbolBuilder) AddSize(arch *sys.Arch, tgt Sym) int64

// GenAddAddrPlusFunc returns a function to be called when capturing
// a function symbol's address. In later stages of the link (when
// address assignment is done) when doing internal linking and
// targeting an executable, we can just emit the address of a function
// directly instead of generating a relocation. Clients can call
// this function (setting 'internalExec' based on build mode and target)
// and then invoke the returned function in roughly the same way that
// loader.*SymbolBuilder.AddAddrPlus would be used.
func GenAddAddrPlusFunc(internalExec bool) func(s *SymbolBuilder, arch *sys.Arch, tgt Sym, add int64) int64

func (sb *SymbolBuilder) MakeWritable()

func (sb *SymbolBuilder) AddUleb(v uint64)
