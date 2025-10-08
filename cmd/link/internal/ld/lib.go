// Inferno utils/8l/asm.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/8l/asm.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ld

import (
	"github.com/shogo82148/std/cmd/internal/bio"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/loader"
	"github.com/shogo82148/std/cmd/link/internal/sym"
)

// ArchSyms holds a number of architecture specific symbols used during
// relocation.  Rather than allowing them universal access to all symbols,
// we keep a subset for relocation application.
type ArchSyms struct {
	Rel     loader.Sym
	Rela    loader.Sym
	RelPLT  loader.Sym
	RelaPLT loader.Sym

	LinkEditGOT loader.Sym
	LinkEditPLT loader.Sym

	TOC    loader.Sym
	DotTOC []loader.Sym

	GOT    loader.Sym
	PLT    loader.Sym
	GOTPLT loader.Sym

	Tlsg      loader.Sym
	Tlsoffset int

	Dynamic loader.Sym
	DynSym  loader.Sym
	DynStr  loader.Sym

	unreachableMethod loader.Sym

	// Symbol containing a list of all the inittasks that need
	// to be run at startup.
	mainInittasks loader.Sym
}

type Arch struct {
	Funcalign  int
	Maxalign   int
	Minalign   int
	Dwarfregsp int
	Dwarfreglr int

	// Threshold of total text size, used for trampoline insertion. If the total
	// text size is smaller than TrampLimit, we won't need to insert trampolines.
	// It is pretty close to the offset range of a direct CALL machine instruction.
	// We leave some room for extra stuff like PLT stubs.
	TrampLimit uint64

	// Empty spaces between codeblocks will be padded with this value.
	// For example an architecture might want to pad with a trap instruction to
	// catch wayward programs. Architectures that do not define a padding value
	// are padded with zeros.
	CodePad []byte

	// Plan 9 variables.
	Plan9Magic  uint32
	Plan9_64Bit bool

	Adddynrel func(*Target, *loader.Loader, *ArchSyms, loader.Sym, loader.Reloc, int) bool
	Archinit  func(*Link)
	// Archreloc is an arch-specific hook that assists in relocation processing
	// (invoked by 'relocsym'); it handles target-specific relocation tasks.
	// Here "rel" is the current relocation being examined, "sym" is the symbol
	// containing the chunk of data to which the relocation applies, and "off"
	// is the contents of the to-be-relocated data item (from sym.P). Return
	// value is the appropriately relocated value (to be written back to the
	// same spot in sym.P), number of external _host_ relocations needed (i.e.
	// ELF/Mach-O/etc. relocations, not Go relocations, this must match ELF.Reloc1,
	// etc.), and a boolean indicating success/failure (a failing value indicates
	// a fatal error).
	Archreloc func(*Target, *loader.Loader, *ArchSyms, loader.Reloc, loader.Sym,
		int64) (relocatedOffset int64, nExtReloc int, ok bool)
	// Archrelocvariant is a second arch-specific hook used for
	// relocation processing; it handles relocations where r.Type is
	// insufficient to describe the relocation (r.Variant !=
	// sym.RV_NONE). Here "rel" is the relocation being applied, "sym"
	// is the symbol containing the chunk of data to which the
	// relocation applies, and "off" is the contents of the
	// to-be-relocated data item (from sym.P). Return is an updated
	// offset value.
	Archrelocvariant func(target *Target, ldr *loader.Loader, rel loader.Reloc,
		rv sym.RelocVariant, sym loader.Sym, offset int64, data []byte) (relocatedOffset int64)

	// Generate a trampoline for a call from s to rs if necessary. ri is
	// index of the relocation.
	Trampoline func(ctxt *Link, ldr *loader.Loader, ri int, rs, s loader.Sym)

	// Assembling the binary breaks into two phases, writing the code/data/
	// dwarf information (which is rather generic), and some more architecture
	// specific work like setting up the elf headers/dynamic relocations, etc.
	// The phases are called "Asmb" and "Asmb2". Asmb2 needs to be defined for
	// every architecture, but only if architecture has an Asmb function will
	// it be used for assembly.  Otherwise a generic assembly Asmb function is
	// used.
	Asmb  func(*Link, *loader.Loader)
	Asmb2 func(*Link, *loader.Loader)

	// Extreloc is an arch-specific hook that converts a Go relocation to an
	// external relocation. Return the external relocation and whether it is
	// needed.
	Extreloc func(*Target, *loader.Loader, loader.Reloc, loader.Sym) (loader.ExtReloc, bool)

	Gentext        func(*Link, *loader.Loader)
	Machoreloc1    func(*sys.Arch, *OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int64) bool
	MachorelocSize uint32
	PEreloc1       func(*sys.Arch, *OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int64) bool
	Xcoffreloc1    func(*sys.Arch, *OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int64) bool

	// Generate additional symbols for the native symbol table just prior to
	// code generation.
	GenSymsLate func(*Link, *loader.Loader)

	// TLSIEtoLE converts a TLS Initial Executable relocation to
	// a TLS Local Executable relocation.
	//
	// This is possible when a TLS IE relocation refers to a local
	// symbol in an executable, which is typical when internally
	// linking PIE binaries.
	TLSIEtoLE func(P []byte, off, size int)

	// optional override for assignAddress
	AssignAddress func(ldr *loader.Loader, sect *sym.Section, n int, s loader.Sym, va uint64, isTramp bool) (*sym.Section, int, uint64)

	// ELF specific information.
	ELF ELFArch
}

// DynlinkingGo reports whether we are producing Go code that can live
// in separate shared libraries linked together at runtime.
func (ctxt *Link) DynlinkingGo() bool

// CanUsePlugins reports whether a plugins can be used
func (ctxt *Link) CanUsePlugins() bool

// NeedCodeSign reports whether we need to code-sign the output binary.
func (ctxt *Link) NeedCodeSign() bool

var (
	Funcalign int

	HEADR int32
)

var (
	Segtext      sym.Segment
	Segrodata    sym.Segment
	Segrelrodata sym.Segment
	Segdata      sym.Segment
	Segdwarf     sym.Segment
	Segpdata     sym.Segment
	Segxdata     sym.Segment

	Segments = []*sym.Segment{&Segtext, &Segrodata, &Segrelrodata, &Segdata, &Segdwarf, &Segpdata, &Segxdata}
)

func Lflag(ctxt *Link, arg string)

type Hostobj struct {
	ld     func(*Link, *bio.Reader, string, int64, string)
	pkg    string
	pn     string
	file   string
	off    int64
	length int64
}

type SymbolType int8

const (
	// see also https://9p.io/magic/man2html/1/nm
	TextSym      SymbolType = 'T'
	DataSym      SymbolType = 'D'
	BSSSym       SymbolType = 'B'
	UndefinedSym SymbolType = 'U'
	TLSSym       SymbolType = 't'
	FrameSym     SymbolType = 'm'
	ParamSym     SymbolType = 'p'
	AutoSym      SymbolType = 'a'

	// Deleted auto (not a real sym, just placeholder for type)
	DeletedAutoSym = 'x'
)

func Entryvalue(ctxt *Link) int64

func Rnd(v int64, r int64) int64

const (
	_ markKind = iota
)

func ElfSymForReloc(ctxt *Link, s loader.Sym) int32

func AddGotSym(target *Target, ldr *loader.Loader, syms *ArchSyms, s loader.Sym, elfRelocTyp uint32)
