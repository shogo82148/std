// Derived from Inferno utils/6l/l.h and related files.
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/6l/l.h
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

package obj

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/cmd/internal/dwarf"
	"github.com/shogo82148/std/cmd/internal/goobj"
	"github.com/shogo82148/std/cmd/internal/objabi"
	"github.com/shogo82148/std/cmd/internal/src"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/internal/abi"
	"github.com/shogo82148/std/sync"
)

type Addr struct {
	Reg    int16
	Index  int16
	Scale  int16
	Type   AddrType
	Name   AddrName
	Class  int8
	Offset int64
	Sym    *LSym

	// argument value:
	//	for TYPE_SCONST, a string
	//	for TYPE_FCONST, a float64
	//	for TYPE_BRANCH, a *Prog (optional)
	//	for TYPE_TEXTSIZE, an int32 (optional)
	Val interface{}
}

type AddrName int8

const (
	NAME_NONE AddrName = iota
	NAME_EXTERN
	NAME_STATIC
	NAME_AUTO
	NAME_PARAM
	// A reference to name@GOT(SB) is a reference to the entry in the global offset
	// table for 'name'.
	NAME_GOTREF
	// Indicates that this is a reference to a TOC anchor.
	NAME_TOCREF
)

type AddrType uint8

const (
	TYPE_NONE AddrType = iota
	TYPE_BRANCH
	TYPE_TEXTSIZE
	TYPE_MEM
	TYPE_CONST
	TYPE_FCONST
	TYPE_SCONST
	TYPE_REG
	TYPE_ADDR
	TYPE_SHIFT
	TYPE_REGREG
	TYPE_REGREG2
	TYPE_INDIR
	TYPE_REGLIST
	TYPE_SPECIAL
)

func (a *Addr) Target() *Prog

func (a *Addr) SetTarget(t *Prog)

func (a *Addr) SetConst(v int64)

// Prog describes a single machine instruction.
//
// The general instruction form is:
//
//	(1) As.Scond From [, ...RestArgs], To
//	(2) As.Scond From, Reg [, ...RestArgs], To, RegTo2
//
// where As is an opcode and the others are arguments:
// From, Reg are sources, and To, RegTo2 are destinations.
// RestArgs can hold additional sources and destinations.
// Usually, not all arguments are present.
// For example, MOVL R1, R2 encodes using only As=MOVL, From=R1, To=R2.
// The Scond field holds additional condition bits for systems (like arm)
// that have generalized conditional execution.
// (2) form is present for compatibility with older code,
// to avoid too much changes in a single swing.
// (1) scheme is enough to express any kind of operand combination.
//
// Jump instructions use the To.Val field to point to the target *Prog,
// which must be in the same linked list as the jump instruction.
//
// The Progs for a given function are arranged in a list linked through the Link field.
//
// Each Prog is charged to a specific source line in the debug information,
// specified by Pos.Line().
// Every Prog has a Ctxt field that defines its context.
// For performance reasons, Progs are usually bulk allocated, cached, and reused;
// those bulk allocators should always be used, rather than new(Prog).
//
// The other fields not yet mentioned are for use by the back ends and should
// be left zeroed by creators of Prog lists.
type Prog struct {
	Ctxt     *Link
	Link     *Prog
	From     Addr
	RestArgs []AddrPos
	To       Addr
	Pool     *Prog
	Forwd    *Prog
	Rel      *Prog
	Pc       int64
	Pos      src.XPos
	Spadj    int32
	As       As
	Reg      int16
	RegTo2   int16
	Mark     uint16
	Optab    uint16
	Scond    uint8
	Back     uint8
	Ft       uint8
	Tt       uint8
	Isize    uint8
}

// AddrPos indicates whether the operand is the source or the destination.
type AddrPos struct {
	Addr
	Pos OperandPos
}

type OperandPos int8

const (
	Source OperandPos = iota
	Destination
)

// From3Type returns p.GetFrom3().Type, or TYPE_NONE when
// p.GetFrom3() returns nil.
func (p *Prog) From3Type() AddrType

// GetFrom3 returns second source operand (the first is Prog.From).
// The same kinds of operands are saved in order so GetFrom3 actually
// return the first source operand in p.RestArgs.
// In combination with Prog.From and Prog.To it makes common 3 operand
// case easier to use.
func (p *Prog) GetFrom3() *Addr

// AddRestSource assigns []Args{{a, Source}} to p.RestArgs.
func (p *Prog) AddRestSource(a Addr)

// AddRestSourceReg calls p.AddRestSource with a register Addr containing reg.
func (p *Prog) AddRestSourceReg(reg int16)

// AddRestSourceConst calls p.AddRestSource with a const Addr containing off.
func (p *Prog) AddRestSourceConst(off int64)

// AddRestDest assigns []Args{{a, Destination}} to p.RestArgs when the second destination
// operand does not fit into prog.RegTo2.
func (p *Prog) AddRestDest(a Addr)

// GetTo2 returns the second destination operand.
// The same kinds of operands are saved in order so GetTo2 actually
// return the first destination operand in Prog.RestArgs[]
func (p *Prog) GetTo2() *Addr

// AddRestSourceArgs assigns more than one source operands to p.RestArgs.
func (p *Prog) AddRestSourceArgs(args []Addr)

// An As denotes an assembler opcode.
// There are some portable opcodes, declared here in package obj,
// that are common to all architectures.
// However, the majority of opcodes are arch-specific
// and are declared in their respective architecture's subpackage.
type As int16

// These are the portable opcodes.
const (
	AXXX As = iota
	ACALL
	ADUFFCOPY
	ADUFFZERO
	AEND
	AFUNCDATA
	AJMP
	ANOP
	APCALIGN
	APCALIGNMAX
	APCDATA
	ARET
	AGETCALLERPC
	ATEXT
	AUNDEF
	A_ARCHSPECIFIC
)

// Each architecture is allotted a distinct subspace of opcode values
// for declaring its arch-specific opcodes.
// Within this subspace, the first arch-specific opcode should be
// at offset A_ARCHSPECIFIC.
//
// Subspaces are aligned to a power of two so opcodes can be masked
// with AMask and used as compact array indices.
const (
	ABase386 = (1 + iota) << 11
	ABaseARM
	ABaseAMD64
	ABasePPC64
	ABaseARM64
	ABaseMIPS
	ABaseLoong64
	ABaseRISCV
	ABaseS390X
	ABaseWasm

	AllowedOpCodes = 1 << 11
	AMask          = AllowedOpCodes - 1
)

// An LSym is the sort of symbol that is written to an object file.
// It represents Go symbols in a flat pkg+"."+name namespace.
type LSym struct {
	Name string
	Type objabi.SymKind
	Attribute

	Size   int64
	Gotype *LSym
	P      []byte
	R      []Reloc

	Extra *interface{}

	Pkg    string
	PkgIdx int32
	SymIdx int32
}

// A FuncInfo contains extra fields for STEXT symbols.
type FuncInfo struct {
	Args      int32
	Locals    int32
	Align     int32
	FuncID    abi.FuncID
	FuncFlag  abi.FuncFlag
	StartLine int32
	Text      *Prog
	Autot     map[*LSym]struct{}
	Pcln      Pcln
	InlMarks  []InlMark
	spills    []RegSpill

	dwarfInfoSym       *LSym
	dwarfLocSym        *LSym
	dwarfRangesSym     *LSym
	dwarfAbsFnSym      *LSym
	dwarfDebugLinesSym *LSym

	GCArgs             *LSym
	GCLocals           *LSym
	StackObjects       *LSym
	OpenCodedDeferInfo *LSym
	ArgInfo            *LSym
	ArgLiveInfo        *LSym
	WrapInfo           *LSym
	JumpTables         []JumpTable

	FuncInfoSym *LSym

	WasmImport *WasmImport
	WasmExport *WasmExport

	sehUnwindInfoSym *LSym
}

// JumpTable represents a table used for implementing multi-way
// computed branching, used typically for implementing switches.
// Sym is the table itself, and Targets is a list of target
// instructions to go to for the computed branch index.
type JumpTable struct {
	Sym     *LSym
	Targets []*Prog
}

// NewFuncInfo allocates and returns a FuncInfo for LSym.
func (s *LSym) NewFuncInfo() *FuncInfo

// Func returns the *FuncInfo associated with s, or else nil.
func (s *LSym) Func() *FuncInfo

type VarInfo struct {
	dwarfInfoSym *LSym
}

// NewVarInfo allocates and returns a VarInfo for LSym.
func (s *LSym) NewVarInfo() *VarInfo

// VarInfo returns the *VarInfo associated with s, or else nil.
func (s *LSym) VarInfo() *VarInfo

// A FileInfo contains extra fields for SDATA symbols backed by files.
// (If LSym.Extra is a *FileInfo, LSym.P == nil.)
type FileInfo struct {
	Name string
	Size int64
}

// NewFileInfo allocates and returns a FileInfo for LSym.
func (s *LSym) NewFileInfo() *FileInfo

// File returns the *FileInfo associated with s, or else nil.
func (s *LSym) File() *FileInfo

// A TypeInfo contains information for a symbol
// that contains a runtime._type.
type TypeInfo struct {
	Type interface{}
}

func (s *LSym) NewTypeInfo() *TypeInfo

// An ItabInfo contains information for a symbol
// that contains a runtime.itab.
type ItabInfo struct {
	Type interface{}
}

func (s *LSym) NewItabInfo() *ItabInfo

// WasmImport represents a WebAssembly (WASM) imported function with
// parameters and results translated into WASM types based on the Go function
// declaration.
type WasmImport struct {
	// Module holds the WASM module name specified by the //go:wasmimport
	// directive.
	Module string
	// Name holds the WASM imported function name specified by the
	// //go:wasmimport directive.
	Name string

	WasmFuncType

	// aux symbol to pass metadata to the linker, serialization of
	// the fields above.
	AuxSym *LSym
}

func (wi *WasmImport) CreateAuxSym()

func (wi *WasmImport) Write(w *bytes.Buffer)

func (wi *WasmImport) Read(b []byte)

// WasmFuncType represents a WebAssembly (WASM) function type with
// parameters and results translated into WASM types based on the Go function
// declaration.
type WasmFuncType struct {
	// Params holds the function parameter fields.
	Params []WasmField
	// Results holds the function result fields.
	Results []WasmField
}

func (ft *WasmFuncType) Write(w *bytes.Buffer)

func (ft *WasmFuncType) Read(b []byte)

// WasmExport represents a WebAssembly (WASM) exported function with
// parameters and results translated into WASM types based on the Go function
// declaration.
type WasmExport struct {
	WasmFuncType

	WrappedSym *LSym
	AuxSym     *LSym
}

func (we *WasmExport) CreateAuxSym()

type WasmField struct {
	Type WasmFieldType
	// Offset holds the frame-pointer-relative locations for Go's stack-based
	// ABI. This is used by the src/cmd/internal/wasm package to map WASM
	// import parameters to the Go stack in a wrapper function.
	Offset int64
}

type WasmFieldType byte

const (
	WasmI32 WasmFieldType = iota
	WasmI64
	WasmF32
	WasmF64
	WasmPtr

	// bool is not really a wasm type, but we allow it on wasmimport/wasmexport
	// function parameters/results. 32-bit on Wasm side, 8-bit on Go side.
	WasmBool
)

type InlMark struct {
	// When unwinding from an instruction in an inlined body, mark
	// where we should unwind to.
	// id records the global inlining id of the inlined body.
	// p records the location of an instruction in the parent (inliner) frame.
	p  *Prog
	id int32
}

// Mark p as the instruction to set as the pc when
// "unwinding" the inlining global frame id. Usually it should be
// instruction with a file:line at the callsite, and occur
// just before the body of the inlined function.
func (fi *FuncInfo) AddInlMark(p *Prog, id int32)

// AddSpill appends a spill record to the list for FuncInfo fi
func (fi *FuncInfo) AddSpill(s RegSpill)

// Record the type symbol for an auto variable so that the linker
// an emit DWARF type information for the type.
func (fi *FuncInfo) RecordAutoType(gotype *LSym)

// ABI is the calling convention of a text symbol.
type ABI uint8

const (
	// ABI0 is the stable stack-based ABI. It's important that the
	// value of this is "0": we can't distinguish between
	// references to data and ABI0 text symbols in assembly code,
	// and hence this doesn't distinguish between symbols without
	// an ABI and text symbols with ABI0.
	ABI0 ABI = iota

	// ABIInternal is the internal ABI that may change between Go
	// versions. All Go functions use the internal ABI and the
	// compiler generates wrappers for calls to and from other
	// ABIs.
	ABIInternal

	ABICount
)

// ParseABI converts from a string representation in 'abistr' to the
// corresponding ABI value. Second return value is TRUE if the
// abi string is recognized, FALSE otherwise.
func ParseABI(abistr string) (ABI, bool)

// ABISet is a bit set of ABI values.
type ABISet uint8

const (
	// ABISetCallable is the set of all ABIs any function could
	// potentially be called using.
	ABISetCallable ABISet = (1 << ABI0) | (1 << ABIInternal)
)

// Ensure ABISet is big enough to hold all ABIs.
var _ ABISet = 1 << (ABICount - 1)

func ABISetOf(abi ABI) ABISet

func (a *ABISet) Set(abi ABI, value bool)

func (a *ABISet) Get(abi ABI) bool

func (a ABISet) String() string

// Attribute is a set of symbol attributes.
type Attribute uint32

const (
	AttrDuplicateOK Attribute = 1 << iota
	AttrCFunc
	AttrNoSplit
	AttrLeaf
	AttrWrapper
	AttrNeedCtxt
	AttrNoFrame
	AttrOnList
	AttrStatic

	// MakeTypelink means that the type should have an entry in the typelink table.
	AttrMakeTypelink

	// ReflectMethod means the function may call reflect.Type.Method or
	// reflect.Type.MethodByName. Matching is imprecise (as reflect.Type
	// can be used through a custom interface), so ReflectMethod may be
	// set in some cases when the reflect package is not called.
	//
	// Used by the linker to determine what methods can be pruned.
	AttrReflectMethod

	// Local means make the symbol local even when compiling Go code to reference Go
	// symbols in other shared libraries, as in this mode symbols are global by
	// default. "local" here means in the sense of the dynamic linker, i.e. not
	// visible outside of the module (shared library or executable) that contains its
	// definition. (When not compiling to support Go shared libraries, all symbols are
	// local in this sense unless there is a cgo_export_* directive).
	AttrLocal

	// For function symbols; indicates that the specified function was the
	// target of an inline during compilation
	AttrWasInlined

	// Indexed indicates this symbol has been assigned with an index (when using the
	// new object file format).
	AttrIndexed

	// Only applied on type descriptor symbols, UsedInIface indicates this type is
	// converted to an interface.
	//
	// Used by the linker to determine what methods can be pruned.
	AttrUsedInIface

	// ContentAddressable indicates this is a content-addressable symbol.
	AttrContentAddressable

	// ABI wrapper is set for compiler-generated text symbols that
	// convert between ABI0 and ABIInternal calling conventions.
	AttrABIWrapper

	// IsPcdata indicates this is a pcdata symbol.
	AttrPcdata

	// PkgInit indicates this is a compiler-generated package init func.
	AttrPkgInit

	// Linkname indicates this is a go:linkname'd symbol.
	AttrLinkname
)

func (a *Attribute) DuplicateOK() bool
func (a *Attribute) MakeTypelink() bool
func (a *Attribute) CFunc() bool
func (a *Attribute) NoSplit() bool
func (a *Attribute) Leaf() bool
func (a *Attribute) OnList() bool
func (a *Attribute) ReflectMethod() bool
func (a *Attribute) Local() bool
func (a *Attribute) Wrapper() bool
func (a *Attribute) NeedCtxt() bool
func (a *Attribute) NoFrame() bool
func (a *Attribute) Static() bool
func (a *Attribute) WasInlined() bool
func (a *Attribute) Indexed() bool
func (a *Attribute) UsedInIface() bool
func (a *Attribute) ContentAddressable() bool
func (a *Attribute) ABIWrapper() bool
func (a *Attribute) IsPcdata() bool
func (a *Attribute) IsPkgInit() bool
func (a *Attribute) IsLinkname() bool

func (a *Attribute) Set(flag Attribute, value bool)

func (a *Attribute) ABI() ABI
func (a *Attribute) SetABI(abi ABI)

// String formats a for printing in as part of a TEXT prog.
func (a Attribute) String() string

// TextAttrString formats the symbol attributes for printing in as part of a TEXT prog.
func (s *LSym) TextAttrString() string

func (s *LSym) String() string

// The compiler needs *LSym to be assignable to cmd/compile/internal/ssa.Sym.
func (*LSym) CanBeAnSSASym()
func (*LSym) CanBeAnSSAAux()

type Pcln struct {
	// Aux symbols for pcln
	Pcsp      *LSym
	Pcfile    *LSym
	Pcline    *LSym
	Pcinline  *LSym
	Pcdata    []*LSym
	Funcdata  []*LSym
	UsedFiles map[goobj.CUFileIndex]struct{}
	InlTree   InlTree
}

type Reloc struct {
	Off  int32
	Siz  uint8
	Type objabi.RelocType
	Add  int64
	Sym  *LSym
}

type Auto struct {
	Asym    *LSym
	Aoffset int32
	Name    AddrName
	Gotype  *LSym
}

// RegSpill provides spill/fill information for a register-resident argument
// to a function.  These need spilling/filling in the safepoint/stackgrowth case.
// At the time of fill/spill, the offset must be adjusted by the architecture-dependent
// adjustment to hardware SP that occurs in a call instruction.  E.g., for AMD64,
// at Offset+8 because the return address was pushed.
type RegSpill struct {
	Addr           Addr
	Reg            int16
	Reg2           int16
	Spill, Unspill As
}

// A Func represents a Go function. If non-nil, it must be a *ir.Func.
type Func interface {
	Pos() src.XPos
}

// Link holds the context for writing object code from a compiler
// to be linker input or for reading that input into the linker.
type Link struct {
	Headtype           objabi.HeadType
	Arch               *LinkArch
	Debugasm           int
	Debugvlog          bool
	Debugpcln          string
	Flag_shared        bool
	Flag_dynlink       bool
	Flag_linkshared    bool
	Flag_optimize      bool
	Flag_locationlists bool
	Flag_noRefName     bool
	Retpoline          bool
	Flag_maymorestack  string
	Bso                *bufio.Writer
	Pathname           string
	Pkgpath            string
	hashmu             sync.Mutex
	hash               map[string]*LSym
	funchash           map[string]*LSym
	statichash         map[string]*LSym
	PosTable           src.PosTable
	InlTree            InlTree
	DwFixups           *DwarfFixupTable
	DwTextCount        int
	Imports            []goobj.ImportedPkg
	DiagFunc           func(string, ...interface{})
	DiagFlush          func()
	DebugInfo          func(ctxt *Link, fn *LSym, info *LSym, curfn Func) ([]dwarf.Scope, dwarf.InlCalls)
	GenAbstractFunc    func(fn *LSym)
	Errors             int

	InParallel    bool
	UseBASEntries bool
	IsAsm         bool
	Std           bool

	// state for writing objects
	Text []*LSym
	Data []*LSym

	// Constant symbols (e.g. $i64.*) are data symbols created late
	// in the concurrent phase. To ensure a deterministic order, we
	// add them to a separate list, sort at the end, and append it
	// to Data.
	constSyms []*LSym

	// Windows SEH symbols are also data symbols that can be created
	// concurrently.
	SEHSyms []*LSym

	// pkgIdx maps package path to index. The index is used for
	// symbol reference in the object file.
	pkgIdx map[string]int32

	defs         []*LSym
	hashed64defs []*LSym
	hasheddefs   []*LSym
	nonpkgdefs   []*LSym
	nonpkgrefs   []*LSym

	Fingerprint goobj.FingerprintType
}

func (ctxt *Link) Diag(format string, args ...interface{})

func (ctxt *Link) Logf(format string, args ...interface{})

// SpillRegisterArgs emits the code to spill register args into whatever
// locations the spill records specify.
func (fi *FuncInfo) SpillRegisterArgs(last *Prog, pa ProgAlloc) *Prog

// UnspillRegisterArgs emits the code to restore register args from whatever
// locations the spill records specify.
func (fi *FuncInfo) UnspillRegisterArgs(last *Prog, pa ProgAlloc) *Prog

// LinkArch is the definition of a single architecture.
type LinkArch struct {
	*sys.Arch
	Init           func(*Link)
	ErrorCheck     func(*Link, *LSym)
	Preprocess     func(*Link, *LSym, ProgAlloc)
	Assemble       func(*Link, *LSym, ProgAlloc)
	Progedit       func(*Link, *Prog, ProgAlloc)
	SEH            func(*Link, *LSym) *LSym
	UnaryDst       map[As]bool
	DWARFRegisters map[int16]int16
}
