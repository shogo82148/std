// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loader

import (
	"github.com/shogo82148/std/cmd/internal/bio"
	"github.com/shogo82148/std/cmd/internal/goobj"
	"github.com/shogo82148/std/cmd/internal/objabi"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/sym"
	"github.com/shogo82148/std/debug/elf"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/internal/abi"
)

var _ = fmt.Print

// Sym encapsulates a global symbol index, used to identify a specific
// Go symbol. The 0-valued Sym is corresponds to an invalid symbol.
type Sym = sym.LoaderSym

// Relocs encapsulates the set of relocations on a given symbol; an
// instance of this type is returned by the Loader Relocs() method.
type Relocs struct {
	rs []goobj.Reloc

	li uint32
	r  *oReader
	l  *Loader
}

// ExtReloc contains the payload for an external relocation.
type ExtReloc struct {
	Xsym Sym
	Xadd int64
	Type objabi.RelocType
	Size uint8
}

// Reloc holds a "handle" to access a relocation record from an
// object file.
type Reloc struct {
	*goobj.Reloc
	r *oReader
	l *Loader
}

func (rel Reloc) Type() objabi.RelocType
func (rel Reloc) Weak() bool
func (rel Reloc) SetType(t objabi.RelocType)
func (rel Reloc) Sym() Sym
func (rel Reloc) SetSym(s Sym)
func (rel Reloc) IsMarker() bool

// Aux holds a "handle" to access an aux symbol record from an
// object file.
type Aux struct {
	*goobj.Aux
	r *oReader
	l *Loader
}

func (a Aux) Sym() Sym

type Bitmap []uint32

// set the i-th bit.
func (bm Bitmap) Set(i Sym)

// unset the i-th bit.
func (bm Bitmap) Unset(i Sym)

// whether the i-th bit is set.
func (bm Bitmap) Has(i Sym) bool

// return current length of bitmap in bits.
func (bm Bitmap) Len() int

// return the number of bits set.
func (bm Bitmap) Count() int

func MakeBitmap(n int) Bitmap

// A Loader loads new object files and resolves indexed symbol references.
//
// Notes on the layout of global symbol index space:
//
//   - Go object files are read before host object files; each Go object
//     read adds its defined package symbols to the global index space.
//     Nonpackage symbols are not yet added.
//
//   - In loader.LoadNonpkgSyms, add non-package defined symbols and
//     references in all object files to the global index space.
//
//   - Host object file loading happens; the host object loader does a
//     name/version lookup for each symbol it finds; this can wind up
//     extending the external symbol index space range. The host object
//     loader stores symbol payloads in loader.payloads using SymbolBuilder.
//
//   - Each symbol gets a unique global index. For duplicated and
//     overwriting/overwritten symbols, the second (or later) appearance
//     of the symbol gets the same global index as the first appearance.
type Loader struct {
	start       map[*oReader]Sym
	objs        []objIdx
	extStart    Sym
	builtinSyms []Sym

	objSyms []objSym

	symsByName    [2]map[string]Sym
	extStaticSyms map[nameVer]Sym

	extReader    *oReader
	payloadBatch []extSymPayload
	payloads     []*extSymPayload
	values       []int64

	sects    []*sym.Section
	symSects []uint16

	align []uint8

	deferReturnTramp map[Sym]bool

	objByPkg map[string]uint32

	anonVersion int

	// Bitmaps and other side structures used to store data used to store
	// symbol flags/attributes; these are to be accessed via the
	// corresponding loader "AttrXXX" and "SetAttrXXX" methods. Please
	// visit the comments on these methods for more details on the
	// semantics / interpretation of the specific flags or attribute.
	attrReachable        Bitmap
	attrOnList           Bitmap
	attrLocal            Bitmap
	attrNotInSymbolTable Bitmap
	attrUsedInIface      Bitmap
	attrSpecial          Bitmap
	attrVisibilityHidden Bitmap
	attrDuplicateOK      Bitmap
	attrShared           Bitmap
	attrExternal         Bitmap
	generatedSyms        Bitmap

	attrReadOnly         map[Sym]bool
	attrCgoExportDynamic map[Sym]struct{}
	attrCgoExportStatic  map[Sym]struct{}

	// Outer and Sub relations for symbols.
	outer []Sym
	sub   map[Sym]Sym

	dynimplib   map[Sym]string
	dynimpvers  map[Sym]string
	localentry  map[Sym]uint8
	extname     map[Sym]string
	elfType     map[Sym]elf.SymType
	elfSym      map[Sym]int32
	localElfSym map[Sym]int32
	symPkg      map[Sym]string
	plt         map[Sym]int32
	got         map[Sym]int32
	dynid       map[Sym]int32

	relocVariant map[relocId]sym.RelocVariant

	// Used to implement field tracking; created during deadcode if
	// field tracking is enabled. Reachparent[K] contains the index of
	// the symbol that triggered the marking of symbol K as live.
	Reachparent []Sym

	// CgoExports records cgo-exported symbols by SymName.
	CgoExports map[string]Sym

	flags uint32

	strictDupMsgs int

	errorReporter *ErrorReporter

	npkgsyms    int
	nhashedsyms int
}

const (
	// Loader.flags
	FlagStrictDups = 1 << iota
)

func NewLoader(flags uint32, reporter *ErrorReporter) *Loader

// LookupOrCreateSym looks up the symbol with the specified name/version,
// returning its Sym index if found. If the lookup fails, a new external
// Sym will be created, entered into the lookup tables, and returned.
func (l *Loader) LookupOrCreateSym(name string, ver int) Sym

// AddCgoExport records a cgo-exported symbol in l.CgoExports.
// This table is used to identify the correct Go symbol ABI to use
// to resolve references from host objects (which don't have ABIs).
func (l *Loader) AddCgoExport(s Sym)

// LookupOrCreateCgoExport is like LookupOrCreateSym, but if ver
// indicates a global symbol, it uses the CgoExport table to determine
// the appropriate symbol version (ABI) to use. ver must be either 0
// or a static symbol version.
func (l *Loader) LookupOrCreateCgoExport(name string, ver int) Sym

func (l *Loader) IsExternal(i Sym) bool

// Look up a symbol by name, return global index, or 0 if not found.
// This is more like Syms.ROLookup than Lookup -- it doesn't create
// new symbol.
func (l *Loader) Lookup(name string, ver int) Sym

func (l *Loader) NStrictDupMsgs() int

// Number of total symbols.
func (l *Loader) NSym() int

// Number of defined Go symbols.
func (l *Loader) NDef() int

// Number of reachable symbols.
func (l *Loader) NReachableSym() int

// Returns the name of the i-th symbol.
func (l *Loader) SymName(i Sym) string

// Returns the version of the i-th symbol.
func (l *Loader) SymVersion(i Sym) int

func (l *Loader) IsFileLocal(i Sym) bool

// IsFromAssembly returns true if this symbol is derived from an
// object file generated by the Go assembler.
func (l *Loader) IsFromAssembly(i Sym) bool

// Returns the type of the i-th symbol.
func (l *Loader) SymType(i Sym) sym.SymKind

// Returns the attributes of the i-th symbol.
func (l *Loader) SymAttr(i Sym) uint8

// Returns the size of the i-th symbol.
func (l *Loader) SymSize(i Sym) int64

// AttrReachable returns true for symbols that are transitively
// referenced from the entry points. Unreachable symbols are not
// written to the output.
func (l *Loader) AttrReachable(i Sym) bool

// SetAttrReachable sets the reachability property for a symbol (see
// AttrReachable).
func (l *Loader) SetAttrReachable(i Sym, v bool)

// AttrOnList returns true for symbols that are on some list (such as
// the list of all text symbols, or one of the lists of data symbols)
// and is consulted to avoid bugs where a symbol is put on a list
// twice.
func (l *Loader) AttrOnList(i Sym) bool

// SetAttrOnList sets the "on list" property for a symbol (see
// AttrOnList).
func (l *Loader) SetAttrOnList(i Sym, v bool)

// AttrLocal returns true for symbols that are only visible within the
// module (executable or shared library) being linked. This attribute
// is applied to thunks and certain other linker-generated symbols.
func (l *Loader) AttrLocal(i Sym) bool

// SetAttrLocal the "local" property for a symbol (see AttrLocal above).
func (l *Loader) SetAttrLocal(i Sym, v bool)

// AttrUsedInIface returns true for a type symbol that is used in
// an interface.
func (l *Loader) AttrUsedInIface(i Sym) bool

func (l *Loader) SetAttrUsedInIface(i Sym, v bool)

// SymAddr checks that a symbol is reachable, and returns its value.
func (l *Loader) SymAddr(i Sym) int64

// AttrNotInSymbolTable returns true for symbols that should not be
// added to the symbol table of the final generated load module.
func (l *Loader) AttrNotInSymbolTable(i Sym) bool

// SetAttrNotInSymbolTable the "not in symtab" property for a symbol
// (see AttrNotInSymbolTable above).
func (l *Loader) SetAttrNotInSymbolTable(i Sym, v bool)

// AttrVisibilityHidden symbols returns true for ELF symbols with
// visibility set to STV_HIDDEN. They become local symbols in
// the final executable. Only relevant when internally linking
// on an ELF platform.
func (l *Loader) AttrVisibilityHidden(i Sym) bool

// SetAttrVisibilityHidden sets the "hidden visibility" property for a
// symbol (see AttrVisibilityHidden).
func (l *Loader) SetAttrVisibilityHidden(i Sym, v bool)

// AttrDuplicateOK returns true for a symbol that can be present in
// multiple object files.
func (l *Loader) AttrDuplicateOK(i Sym) bool

// SetAttrDuplicateOK sets the "duplicate OK" property for an external
// symbol (see AttrDuplicateOK).
func (l *Loader) SetAttrDuplicateOK(i Sym, v bool)

// AttrShared returns true for symbols compiled with the -shared option.
func (l *Loader) AttrShared(i Sym) bool

// SetAttrShared sets the "shared" property for an external
// symbol (see AttrShared).
func (l *Loader) SetAttrShared(i Sym, v bool)

// AttrExternal returns true for function symbols loaded from host
// object files.
func (l *Loader) AttrExternal(i Sym) bool

// SetAttrExternal sets the "external" property for a host object
// symbol (see AttrExternal).
func (l *Loader) SetAttrExternal(i Sym, v bool)

// AttrSpecial returns true for a symbols that do not have their
// address (i.e. Value) computed by the usual mechanism of
// data.go:dodata() & data.go:address().
func (l *Loader) AttrSpecial(i Sym) bool

// SetAttrSpecial sets the "special" property for a symbol (see
// AttrSpecial).
func (l *Loader) SetAttrSpecial(i Sym, v bool)

// AttrCgoExportDynamic returns true for a symbol that has been
// specially marked via the "cgo_export_dynamic" compiler directive
// written by cgo (in response to //export directives in the source).
func (l *Loader) AttrCgoExportDynamic(i Sym) bool

// SetAttrCgoExportDynamic sets the "cgo_export_dynamic" for a symbol
// (see AttrCgoExportDynamic).
func (l *Loader) SetAttrCgoExportDynamic(i Sym, v bool)

// ForAllCgoExportDynamic calls f for every symbol that has been
// marked with the "cgo_export_dynamic" compiler directive.
func (l *Loader) ForAllCgoExportDynamic(f func(Sym))

// AttrCgoExportStatic returns true for a symbol that has been
// specially marked via the "cgo_export_static" directive
// written by cgo.
func (l *Loader) AttrCgoExportStatic(i Sym) bool

// SetAttrCgoExportStatic sets the "cgo_export_static" for a symbol
// (see AttrCgoExportStatic).
func (l *Loader) SetAttrCgoExportStatic(i Sym, v bool)

// IsGeneratedSym returns true if a symbol's been previously marked as a
// generator symbol through the SetIsGeneratedSym. The functions for generator
// symbols are kept in the Link context.
func (l *Loader) IsGeneratedSym(i Sym) bool

// SetIsGeneratedSym marks symbols as generated symbols. Data shouldn't be
// stored in generated symbols, and a function is registered and called for
// each of these symbols.
func (l *Loader) SetIsGeneratedSym(i Sym, v bool)

func (l *Loader) AttrCgoExport(i Sym) bool

// AttrReadOnly returns true for a symbol whose underlying data
// is stored via a read-only mmap.
func (l *Loader) AttrReadOnly(i Sym) bool

// SetAttrReadOnly sets the "data is read only" property for a symbol
// (see AttrReadOnly).
func (l *Loader) SetAttrReadOnly(i Sym, v bool)

func (l *Loader) AttrSubSymbol(i Sym) bool

// Returns whether the i-th symbol has ReflectMethod attribute set.
func (l *Loader) IsReflectMethod(i Sym) bool

// Returns whether the i-th symbol is nosplit.
func (l *Loader) IsNoSplit(i Sym) bool

// Returns whether this is a Go type symbol.
func (l *Loader) IsGoType(i Sym) bool

// Returns whether this symbol should be included in typelink.
func (l *Loader) IsTypelink(i Sym) bool

// Returns whether this symbol is an itab symbol.
func (l *Loader) IsItab(i Sym) bool

// Returns whether this symbol is a dictionary symbol.
func (l *Loader) IsDict(i Sym) bool

// Returns whether this symbol is a compiler-generated package init func.
func (l *Loader) IsPkgInit(i Sym) bool

// Return whether this is a trampoline of a deferreturn call.
func (l *Loader) IsDeferReturnTramp(i Sym) bool

// Set that i is a trampoline of a deferreturn call.
func (l *Loader) SetIsDeferReturnTramp(i Sym, v bool)

// SymValue returns the value of the i-th symbol. i is global index.
func (l *Loader) SymValue(i Sym) int64

// SetSymValue sets the value of the i-th symbol. i is global index.
func (l *Loader) SetSymValue(i Sym, val int64)

// AddToSymValue adds to the value of the i-th symbol. i is the global index.
func (l *Loader) AddToSymValue(i Sym, val int64)

// Returns the symbol content of the i-th symbol. i is global index.
func (l *Loader) Data(i Sym) []byte

// Returns the symbol content of the i-th symbol as a string. i is global index.
func (l *Loader) DataString(i Sym) string

// FreeData clears the symbol data of an external symbol, allowing the memory
// to be freed earlier. No-op for non-external symbols.
// i is global index.
func (l *Loader) FreeData(i Sym)

// SymAlign returns the alignment for a symbol.
func (l *Loader) SymAlign(i Sym) int32

// SetSymAlign sets the alignment for a symbol.
func (l *Loader) SetSymAlign(i Sym, align int32)

// SymSect returns the section of the i-th symbol. i is global index.
func (l *Loader) SymSect(i Sym) *sym.Section

// SetSymSect sets the section of the i-th symbol. i is global index.
func (l *Loader) SetSymSect(i Sym, sect *sym.Section)

// NewSection creates a new (output) section.
func (l *Loader) NewSection() *sym.Section

// SymDynimplib returns the "dynimplib" attribute for the specified
// symbol, making up a portion of the info for a symbol specified
// on a "cgo_import_dynamic" compiler directive.
func (l *Loader) SymDynimplib(i Sym) string

// SetSymDynimplib sets the "dynimplib" attribute for a symbol.
func (l *Loader) SetSymDynimplib(i Sym, value string)

// SymDynimpvers returns the "dynimpvers" attribute for the specified
// symbol, making up a portion of the info for a symbol specified
// on a "cgo_import_dynamic" compiler directive.
func (l *Loader) SymDynimpvers(i Sym) string

// SetSymDynimpvers sets the "dynimpvers" attribute for a symbol.
func (l *Loader) SetSymDynimpvers(i Sym, value string)

// SymExtname returns the "extname" value for the specified
// symbol.
func (l *Loader) SymExtname(i Sym) string

// SetSymExtname sets the  "extname" attribute for a symbol.
func (l *Loader) SetSymExtname(i Sym, value string)

// SymElfType returns the previously recorded ELF type for a symbol
// (used only for symbols read from shared libraries by ldshlibsyms).
// It is not set for symbols defined by the packages being linked or
// by symbols read by ldelf (and so is left as elf.STT_NOTYPE).
func (l *Loader) SymElfType(i Sym) elf.SymType

// SetSymElfType sets the elf type attribute for a symbol.
func (l *Loader) SetSymElfType(i Sym, et elf.SymType)

// SymElfSym returns the ELF symbol index for a given loader
// symbol, assigned during ELF symtab generation.
func (l *Loader) SymElfSym(i Sym) int32

// SetSymElfSym sets the elf symbol index for a symbol.
func (l *Loader) SetSymElfSym(i Sym, es int32)

// SymLocalElfSym returns the "local" ELF symbol index for a given loader
// symbol, assigned during ELF symtab generation.
func (l *Loader) SymLocalElfSym(i Sym) int32

// SetSymLocalElfSym sets the "local" elf symbol index for a symbol.
func (l *Loader) SetSymLocalElfSym(i Sym, es int32)

// SymPlt returns the PLT offset of symbol s.
func (l *Loader) SymPlt(s Sym) int32

// SetPlt sets the PLT offset of symbol i.
func (l *Loader) SetPlt(i Sym, v int32)

// SymGot returns the GOT offset of symbol s.
func (l *Loader) SymGot(s Sym) int32

// SetGot sets the GOT offset of symbol i.
func (l *Loader) SetGot(i Sym, v int32)

// SymDynid returns the "dynid" property for the specified symbol.
func (l *Loader) SymDynid(i Sym) int32

// SetSymDynid sets the "dynid" property for a symbol.
func (l *Loader) SetSymDynid(i Sym, val int32)

// DynidSyms returns the set of symbols for which dynID is set to an
// interesting (non-default) value. This is expected to be a fairly
// small set.
func (l *Loader) DynidSyms() []Sym

// SymGoType returns the 'Gotype' property for a given symbol (set by
// the Go compiler for variable symbols). This version relies on
// reading aux symbols for the target sym -- it could be that a faster
// approach would be to check for gotype during preload and copy the
// results in to a map (might want to try this at some point and see
// if it helps speed things up).
func (l *Loader) SymGoType(i Sym) Sym

// SymUnit returns the compilation unit for a given symbol (which will
// typically be nil for external or linker-manufactured symbols).
func (l *Loader) SymUnit(i Sym) *sym.CompilationUnit

// SymPkg returns the package where the symbol came from (for
// regular compiler-generated Go symbols), but in the case of
// building with "-linkshared" (when a symbol is read from a
// shared library), will hold the library name.
// NOTE: this corresponds to sym.Symbol.File field.
func (l *Loader) SymPkg(i Sym) string

// SetSymPkg sets the package/library for a symbol. This is
// needed mainly for external symbols, specifically those imported
// from shared libraries.
func (l *Loader) SetSymPkg(i Sym, pkg string)

// SymLocalentry returns an offset in bytes of the "local entry" of a symbol.
//
// On PPC64, a value of 1 indicates the symbol does not use or preserve a TOC
// pointer in R2, nor does it have a distinct local entry.
func (l *Loader) SymLocalentry(i Sym) uint8

// SetSymLocalentry sets the "local entry" offset attribute for a symbol.
func (l *Loader) SetSymLocalentry(i Sym, value uint8)

// Returns the number of aux symbols given a global index.
func (l *Loader) NAux(i Sym) int

// Returns the "handle" to the j-th aux symbol of the i-th symbol.
func (l *Loader) Aux(i Sym, j int) Aux

// WasmImportSym returns the auxiliary WebAssembly import symbol associated with
// a given function symbol. The aux sym only exists for Go function stubs that
// have been annotated with the //go:wasmimport directive.  The aux sym
// contains the information necessary for the linker to add a WebAssembly
// import statement.
// (https://webassembly.github.io/spec/core/syntax/modules.html#imports)
func (l *Loader) WasmImportSym(fnSymIdx Sym) (Sym, bool)

// SEHUnwindSym returns the auxiliary SEH unwind symbol associated with
// a given function symbol.
func (l *Loader) SEHUnwindSym(fnSymIdx Sym) Sym

// GetFuncDwarfAuxSyms collects and returns the auxiliary DWARF
// symbols associated with a given function symbol.  Prior to the
// introduction of the loader, this was done purely using name
// lookups, e.f. for function with name XYZ we would then look up
// go.info.XYZ, etc.
func (l *Loader) GetFuncDwarfAuxSyms(fnSymIdx Sym) (auxDwarfInfo, auxDwarfLoc, auxDwarfRanges, auxDwarfLines Sym)

func (l *Loader) GetVarDwarfAuxSym(i Sym) Sym

// AddInteriorSym sets up 'interior' as an interior symbol of
// container/payload symbol 'container'. An interior symbol does not
// itself have data, but gives a name to a subrange of the data in its
// container symbol. The container itself may or may not have a name.
// This method is intended primarily for use in the host object
// loaders, to capture the semantics of symbols and sections in an
// object file. When reading a host object file, we'll typically
// encounter a static section symbol (ex: ".text") containing content
// for a collection of functions, then a series of ELF (or macho, etc)
// symbol table entries each of which points into a sub-section
// (offset and length) of its corresponding container symbol. Within
// the go linker we create a loader.Sym for the container (which is
// expected to have the actual content/payload) and then a set of
// interior loader.Sym's that point into a portion of the container.
func (l *Loader) AddInteriorSym(container Sym, interior Sym)

// OuterSym gets the outer/container symbol.
func (l *Loader) OuterSym(i Sym) Sym

// SubSym gets the subsymbol for host object loaded symbols.
func (l *Loader) SubSym(i Sym) Sym

// SetCarrierSym declares that 'c' is the carrier or container symbol
// for 's'. Carrier symbols are used in the linker to as a container
// for a collection of sub-symbols where the content of the
// sub-symbols is effectively concatenated to form the content of the
// carrier. The carrier is given a name in the output symbol table
// while the sub-symbol names are not. For example, the Go compiler
// emits named string symbols (type SGOSTRING) when compiling a
// package; after being deduplicated, these symbols are collected into
// a single unit by assigning them a new carrier symbol named
// "go:string.*" (which appears in the final symbol table for the
// output load module).
func (l *Loader) SetCarrierSym(s Sym, c Sym)

// Initialize Reachable bitmap and its siblings for running deadcode pass.
func (l *Loader) InitReachable()

// SortSub walks through the sub-symbols for 's' and sorts them
// in place by increasing value. Return value is the new
// sub symbol for the specified outer symbol.
func (l *Loader) SortSub(s Sym) Sym

// SortSyms sorts a list of symbols by their value.
func (l *Loader) SortSyms(ss []Sym)

func (relocs *Relocs) Count() int

// At returns the j-th reloc for a global symbol.
func (relocs *Relocs) At(j int) Reloc

// Relocs returns a Relocs object for the given global sym.
func (l *Loader) Relocs(i Sym) Relocs

func (l *Loader) Pcsp(i Sym) Sym

// Returns all aux symbols of per-PC data for symbol i.
// tmp is a scratch space for the pcdata slice.
func (l *Loader) PcdataAuxs(i Sym, tmp []Sym) (pcsp, pcfile, pcline, pcinline Sym, pcdata []Sym)

// Returns the number of pcdata for symbol i.
func (l *Loader) NumPcdata(i Sym) int

// Returns all funcdata symbols of symbol i.
// tmp is a scratch space.
func (l *Loader) Funcdata(i Sym, tmp []Sym) []Sym

// Returns the number of funcdata for symbol i.
func (l *Loader) NumFuncdata(i Sym) int

// FuncInfo provides hooks to access goobj.FuncInfo in the objects.
type FuncInfo struct {
	l       *Loader
	r       *oReader
	data    []byte
	lengths goobj.FuncInfoLengths
}

func (fi *FuncInfo) Valid() bool

func (fi *FuncInfo) Args() int

func (fi *FuncInfo) Locals() int

func (fi *FuncInfo) FuncID() abi.FuncID

func (fi *FuncInfo) FuncFlag() abi.FuncFlag

func (fi *FuncInfo) StartLine() int32

// Preload has to be called prior to invoking the various methods
// below related to pcdata, funcdataoff, files, and inltree nodes.
func (fi *FuncInfo) Preload()

func (fi *FuncInfo) NumFile() uint32

func (fi *FuncInfo) File(k int) goobj.CUFileIndex

// TopFrame returns true if the function associated with this FuncInfo
// is an entry point, meaning that unwinders should stop when they hit
// this function.
func (fi *FuncInfo) TopFrame() bool

type InlTreeNode struct {
	Parent   int32
	File     goobj.CUFileIndex
	Line     int32
	Func     Sym
	ParentPC int32
}

func (fi *FuncInfo) NumInlTree() uint32

func (fi *FuncInfo) InlTree(k int) InlTreeNode

func (l *Loader) FuncInfo(i Sym) FuncInfo

// Preload a package: adds autolib.
// Does not add defined package or non-packaged symbols to the symbol table.
// These are done in LoadSyms.
// Does not read symbol data.
// Returns the fingerprint of the object.
func (l *Loader) Preload(localSymVersion int, f *bio.Reader, lib *sym.Library, unit *sym.CompilationUnit, length int64) goobj.FingerprintType

// Add syms, hashed (content-addressable) symbols, non-package symbols, and
// references to external symbols (which are always named).
func (l *Loader) LoadSyms(arch *sys.Arch)

// TopLevelSym tests a symbol (by name and kind) to determine whether
// the symbol first class sym (participating in the link) or is an
// anonymous aux or sub-symbol containing some sub-part or payload of
// another symbol.
func (l *Loader) TopLevelSym(s Sym) bool

// Copy the payload of symbol src to dst. Both src and dst must be external
// symbols.
// The intended use case is that when building/linking against a shared library,
// where we do symbol name mangling, the Go object file may have reference to
// the original symbol name whereas the shared library provides a symbol with
// the mangled name. When we do mangling, we copy payload of mangled to original.
func (l *Loader) CopySym(src, dst Sym)

// CreateExtSym creates a new external symbol with the specified name
// without adding it to any lookup tables, returning a Sym index for it.
func (l *Loader) CreateExtSym(name string, ver int) Sym

// CreateStaticSym creates a new static symbol with the specified name
// without adding it to any lookup tables, returning a Sym index for it.
func (l *Loader) CreateStaticSym(name string) Sym

func (l *Loader) FreeSym(i Sym)

// SetRelocVariant sets the 'variant' property of a relocation on
// some specific symbol.
func (l *Loader) SetRelocVariant(s Sym, ri int, v sym.RelocVariant)

// RelocVariant returns the 'variant' property of a relocation on
// some specific symbol.
func (l *Loader) RelocVariant(s Sym, ri int) sym.RelocVariant

// UndefinedRelocTargets iterates through the global symbol index
// space, looking for symbols with relocations targeting undefined
// references. The linker's loadlib method uses this to determine if
// there are unresolved references to functions in system libraries
// (for example, libgcc.a), presumably due to CGO code. Return value
// is a pair of lists of loader.Sym's. First list corresponds to the
// corresponding to the undefined symbols themselves, the second list
// is the symbol that is making a reference to the undef. The "limit"
// param controls the maximum number of results returned; if "limit"
// is -1, then all undefs are returned.
func (l *Loader) UndefinedRelocTargets(limit int) ([]Sym, []Sym)

// AssignTextSymbolOrder populates the Textp slices within each
// library and compilation unit, insuring that packages are laid down
// in dependency order (internal first, then everything else). Return value
// is a slice of all text syms.
func (l *Loader) AssignTextSymbolOrder(libs []*sym.Library, intlibs []bool, extsyms []Sym) []Sym

// ErrorReporter is a helper class for reporting errors.
type ErrorReporter struct {
	ldr              *Loader
	AfterErrorAction func()
}

// Errorf method logs an error message.
//
// After each error, the error actions function will be invoked; this
// will either terminate the link immediately (if -h option given)
// or it will keep a count and exit if more than 20 errors have been printed.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
func (reporter *ErrorReporter) Errorf(s Sym, format string, args ...interface{})

// GetErrorReporter returns the loader's associated error reporter.
func (l *Loader) GetErrorReporter() *ErrorReporter

// Errorf method logs an error message. See ErrorReporter.Errorf for details.
func (l *Loader) Errorf(s Sym, format string, args ...interface{})

// Symbol statistics.
func (l *Loader) Stat() string

// For debugging.
func (l *Loader) Dump()
