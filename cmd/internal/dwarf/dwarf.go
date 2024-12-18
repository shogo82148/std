// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dwarf generates DWARF debugging information.
// DWARF generation is split between the compiler and the linker,
// this package contains the shared code.
package dwarf

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// InfoPrefix is the prefix for all the symbols containing DWARF info entries.
const InfoPrefix = "go:info."

// ConstInfoPrefix is the prefix for all symbols containing DWARF info
// entries that contain constants.
const ConstInfoPrefix = "go:constinfo."

// CUInfoPrefix is the prefix for symbols containing information to
// populate the DWARF compilation unit info entries.
const CUInfoPrefix = "go:cuinfo."

// Used to form the symbol name assigned to the DWARF 'abstract subprogram"
// info entry for a function
const AbstractFuncSuffix = "$abstract"

// Sym represents a symbol.
type Sym interface {
}

// A Var represents a local variable or a function parameter.
type Var struct {
	Name          string
	Tag           int
	WithLoclist   bool
	IsReturnValue bool
	IsInlFormal   bool
	DictIndex     uint16
	StackOffset   int32
	// This package can't use the ssa package, so it can't mention ssa.FuncDebug,
	// so indirect through a closure.
	PutLocationList func(listSym, startPC Sym)
	Scope           int32
	Type            Sym
	DeclFile        string
	DeclLine        uint
	DeclCol         uint
	InlIndex        int32
	ChildIndex      int32
	IsInAbstract    bool
	ClosureOffset   int64
}

// A Scope represents a lexical scope. All variables declared within a
// scope will only be visible to instructions covered by the scope.
// Lexical scopes are contiguous in source files but can end up being
// compiled to discontiguous blocks of instructions in the executable.
// The Ranges field lists all the blocks of instructions that belong
// in this scope.
type Scope struct {
	Parent int32
	Ranges []Range
	Vars   []*Var
}

// A Range represents a half-open interval [Start, End).
type Range struct {
	Start, End int64
}

// This container is used by the PutFunc* variants below when
// creating the DWARF subprogram DIE(s) for a function.
type FnState struct {
	Name          string
	Info          Sym
	Loc           Sym
	Ranges        Sym
	Absfn         Sym
	StartPC       Sym
	StartPos      src.Pos
	Size          int64
	External      bool
	Scopes        []Scope
	InlCalls      InlCalls
	UseBASEntries bool

	dictIndexToOffset []int64
}

func EnableLogging(doit bool)

// MergeRanges creates a new range list by merging the ranges from
// its two arguments, then returns the new list.
func MergeRanges(in1, in2 []Range) []Range

// UnifyRanges merges the ranges from 'c' into the list of ranges for 's'.
func (s *Scope) UnifyRanges(c *Scope)

// AppendRange adds r to s, if r is non-empty.
// If possible, it extends the last Range in s.Ranges; if not, it creates a new one.
func (s *Scope) AppendRange(r Range)

type InlCalls struct {
	Calls []InlCall
}

type InlCall struct {
	// index into ctx.InlTree describing the call inlined here
	InlIndex int

	// Position of the inlined call site.
	CallPos src.Pos

	// Dwarf abstract subroutine symbol (really *obj.LSym).
	AbsFunSym Sym

	// Indices of child inlines within Calls array above.
	Children []int

	// entries in this list are PAUTO's created by the inliner to
	// capture the promoted formals and locals of the inlined callee.
	InlVars []*Var

	// PC ranges for this inlined call.
	Ranges []Range

	// Root call (not a child of some other call).
	Root bool
}

// A Context specifies how to add data to a Sym.
type Context interface {
	PtrSize() int
	Size(s Sym) int64
	AddInt(s Sym, size int, i int64)
	AddBytes(s Sym, b []byte)
	AddAddress(s Sym, t interface{}, ofs int64)
	AddCURelativeAddress(s Sym, t interface{}, ofs int64)
	AddSectionOffset(s Sym, size int, t interface{}, ofs int64)
	AddDWARFAddrSectionOffset(s Sym, t interface{}, ofs int64)
	CurrentOffset(s Sym) int64
	RecordDclReference(from Sym, to Sym, dclIdx int, inlIndex int)
	RecordChildDieOffsets(s Sym, vars []*Var, offsets []int32)
	AddString(s Sym, v string)
	Logf(format string, args ...interface{})
}

// AppendUleb128 appends v to b using DWARF's unsigned LEB128 encoding.
func AppendUleb128(b []byte, v uint64) []byte

// AppendSleb128 appends v to b using DWARF's signed LEB128 encoding.
func AppendSleb128(b []byte, v int64) []byte

// Uleb128put appends v to s using DWARF's unsigned LEB128 encoding.
func Uleb128put(ctxt Context, s Sym, v int64)

// Sleb128put appends v to s using DWARF's signed LEB128 encoding.
func Sleb128put(ctxt Context, s Sym, v int64)

// Go-specific type attributes.
const (
	DW_AT_go_kind = 0x2900
	DW_AT_go_key  = 0x2901
	DW_AT_go_elem = 0x2902
	// Attribute for DW_TAG_member of a struct type.
	// Nonzero value indicates the struct field is an embedded field.
	DW_AT_go_embedded_field = 0x2903
	DW_AT_go_runtime_type   = 0x2904

	DW_AT_go_package_name   = 0x2905
	DW_AT_go_dict_index     = 0x2906
	DW_AT_go_closure_offset = 0x2907

	DW_AT_internal_location = 253
)

// Index into the abbrevs table below.
const (
	DW_ABRV_NULL = iota
	DW_ABRV_COMPUNIT
	DW_ABRV_COMPUNIT_TEXTLESS
	DW_ABRV_FUNCTION
	DW_ABRV_WRAPPER
	DW_ABRV_FUNCTION_ABSTRACT
	DW_ABRV_FUNCTION_CONCRETE
	DW_ABRV_WRAPPER_CONCRETE
	DW_ABRV_INLINED_SUBROUTINE
	DW_ABRV_INLINED_SUBROUTINE_RANGES
	DW_ABRV_VARIABLE
	DW_ABRV_INT_CONSTANT
	DW_ABRV_LEXICAL_BLOCK_RANGES
	DW_ABRV_LEXICAL_BLOCK_SIMPLE
	DW_ABRV_STRUCTFIELD
	DW_ABRV_FUNCTYPEPARAM
	DW_ABRV_DOTDOTDOT
	DW_ABRV_ARRAYRANGE
	DW_ABRV_NULLTYPE
	DW_ABRV_BASETYPE
	DW_ABRV_ARRAYTYPE
	DW_ABRV_CHANTYPE
	DW_ABRV_FUNCTYPE
	DW_ABRV_IFACETYPE
	DW_ABRV_MAPTYPE
	DW_ABRV_PTRTYPE
	DW_ABRV_BARE_PTRTYPE
	DW_ABRV_SLICETYPE
	DW_ABRV_STRINGTYPE
	DW_ABRV_STRUCTTYPE
	DW_ABRV_TYPEDECL
	DW_ABRV_DICT_INDEX
	DW_ABRV_PUTVAR_START
)

// Abbrevs returns the finalized abbrev array for the platform,
// expanding any DW_FORM pseudo-ops to real values.
func Abbrevs() []dwAbbrev

// GetAbbrev returns the contents of the .debug_abbrev section.
func GetAbbrev() []byte

// DWAttr represents an attribute of a DWDie.
//
// For DW_CLS_string and _block, value should contain the length, and
// data the data, for _reference, value is 0 and data is a DWDie* to
// the referenced instance, for all others, value is the whole thing
// and data is null.
type DWAttr struct {
	Link  *DWAttr
	Atr   uint16
	Cls   uint8
	Value int64
	Data  interface{}
}

// DWDie represents a DWARF debug info entry.
type DWDie struct {
	Abbrev int
	Link   *DWDie
	Child  *DWDie
	Attr   *DWAttr
	Sym    Sym
}

// PutAttrs writes the attributes for a DIE to symbol 's'.
//
// Note that we can (and do) add arbitrary attributes to a DIE, but
// only the ones actually listed in the Abbrev will be written out.
func PutAttrs(ctxt Context, s Sym, abbrev int, attr *DWAttr)

// HasChildren reports whether 'die' uses an abbrev that supports children.
func HasChildren(die *DWDie) bool

// PutIntConst writes a DIE for an integer constant
func PutIntConst(ctxt Context, info, typ Sym, name string, val int64)

// PutGlobal writes a DIE for a global variable.
func PutGlobal(ctxt Context, info, typ, gvar Sym, name string)

// PutBasedRanges writes a range table to sym. All addresses in ranges are
// relative to some base address, which must be arranged by the caller
// (e.g., with a DW_AT_low_pc attribute, or in a BASE-prefixed range).
func PutBasedRanges(ctxt Context, sym Sym, ranges []Range)

// PutRanges writes a range table to s.Ranges.
// All addresses in ranges are relative to s.base.
func (s *FnState) PutRanges(ctxt Context, ranges []Range)

// Emit DWARF attributes and child DIEs for an 'abstract' subprogram.
// The abstract subprogram DIE for a function contains its
// location-independent attributes (name, type, etc). Other instances
// of the function (any inlined copy of it, or the single out-of-line
// 'concrete' instance) will contain a pointer back to this abstract
// DIE (as a space-saving measure, so that name/type etc doesn't have
// to be repeated for each inlined copy).
func PutAbstractFunc(ctxt Context, s *FnState) error

// Emit DWARF attributes and child DIEs for a 'concrete' subprogram,
// meaning the out-of-line copy of a function that was inlined at some
// point during the compilation of its containing package. The first
// attribute for a concrete DIE is a reference to the 'abstract' DIE
// for the function (which holds location-independent attributes such
// as name, type), then the remainder of the attributes are specific
// to this instance (location, frame base, etc).
func PutConcreteFunc(ctxt Context, s *FnState, isWrapper bool) error

// Emit DWARF attributes and child DIEs for a subprogram. Here
// 'default' implies that the function in question was not inlined
// when its containing package was compiled (hence there is no need to
// emit an abstract version for it to use as a base for inlined
// routine records).
func PutDefaultFunc(ctxt Context, s *FnState, isWrapper bool) error

// IsDWARFEnabledOnAIXLd returns true if DWARF is possible on the
// current extld.
// AIX ld doesn't support DWARF with -bnoobjreorder with version
// prior to 7.2.2.
func IsDWARFEnabledOnAIXLd(extld []string) (bool, error)
