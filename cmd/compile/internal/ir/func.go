// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Func corresponds to a single function in a Go program
// (and vice versa: each function is denoted by exactly one *Func).
//
// There are multiple nodes that represent a Func in the IR.
//
// The ONAME node (Func.Nname) is used for plain references to it.
// The ODCLFUNC node (the Func itself) is used for its declaration code.
// The OCLOSURE node (Func.OClosure) is used for a reference to a
// function literal.
//
// An imported function will have an ONAME node which points to a Func
// with an empty body.
// A declared function or method has an ODCLFUNC (the Func itself) and an ONAME.
// A function literal is represented directly by an OCLOSURE, but it also
// has an ODCLFUNC (and a matching ONAME) representing the compiled
// underlying form of the closure, which accesses the captured variables
// using a special data structure passed in a register.
//
// A method declaration is represented like functions, except f.Sym
// will be the qualified method name (e.g., "T.m").
//
// A method expression (T.M) is represented as an OMETHEXPR node,
// in which n.Left and n.Right point to the type and method, respectively.
// Each distinct mention of a method expression in the source code
// constructs a fresh node.
//
// A method value (t.M) is represented by ODOTMETH/ODOTINTER
// when it is called directly and by OMETHVALUE otherwise.
// These are like method expressions, except that for ODOTMETH/ODOTINTER,
// the method name is stored in Sym instead of Right.
// Each OMETHVALUE ends up being implemented as a new
// function, a bit like a closure, with its own ODCLFUNC.
// The OMETHVALUE uses n.Func to record the linkage to
// the generated ODCLFUNC, but there is no
// pointer from the Func back to the OMETHVALUE.
type Func struct {
	miniNode
	Body Nodes

	Nname    *Name
	OClosure *ClosureExpr

	// Extra entry code for the function. For example, allocate and initialize
	// memory for escaping parameters.
	Enter Nodes
	Exit  Nodes

	// ONAME nodes for all params/locals for this func/closure, does NOT
	// include closurevars until transforming closures during walk.
	// Names must be listed PPARAMs, PPARAMOUTs, then PAUTOs,
	// with PPARAMs and PPARAMOUTs in order corresponding to the function signature.
	// However, as anonymous or blank PPARAMs are not actually declared,
	// they are omitted from Dcl.
	// Anonymous and blank PPARAMOUTs are declared as ~rNN and ~bNN Names, respectively.
	Dcl []*Name

	// ClosureVars lists the free variables that are used within a
	// function literal, but formally declared in an enclosing
	// function. The variables in this slice are the closure function's
	// own copy of the variables, which are used within its function
	// body. They will also each have IsClosureVar set, and will have
	// Byval set if they're captured by value.
	ClosureVars []*Name

	// Enclosed functions that need to be compiled.
	// Populated during walk.
	Closures []*Func

	// Parents records the parent scope of each scope within a
	// function. The root scope (0) has no parent, so the i'th
	// scope's parent is stored at Parents[i-1].
	Parents []ScopeID

	// Marks records scope boundary changes.
	Marks []Mark

	FieldTrack map[*obj.LSym]struct{}
	DebugInfo  interface{}
	LSym       *obj.LSym

	Inl *Inline

	// Closgen tracks how many closures have been generated within
	// this function. Used by closurename for creating unique
	// function names.
	Closgen int32

	Label int32

	Endlineno src.XPos
	WBPos     src.XPos

	Pragma PragmaFlag

	flags bitset16

	// ABI is a function's "definition" ABI. This is the ABI that
	// this function's generated code is expecting to be called by.
	//
	// For most functions, this will be obj.ABIInternal. It may be
	// a different ABI for functions defined in assembly or ABI wrappers.
	//
	// This is included in the export data and tracked across packages.
	ABI obj.ABI
	// ABIRefs is the set of ABIs by which this function is referenced.
	// For ABIs other than this function's definition ABI, the
	// compiler generates ABI wrapper functions. This is only tracked
	// within a package.
	ABIRefs obj.ABISet

	NumDefers  int32
	NumReturns int32

	// nwbrCalls records the LSyms of functions called by this
	// function for go:nowritebarrierrec analysis. Only filled in
	// if nowritebarrierrecCheck != nil.
	NWBRCalls *[]SymAndPos

	// For wrapper functions, WrappedFunc point to the original Func.
	// Currently only used for go/defer wrappers.
	WrappedFunc *Func

	// WasmImport is used by the //go:wasmimport directive to store info about
	// a WebAssembly function import.
	WasmImport *WasmImport
}

// WasmImport stores metadata associated with the //go:wasmimport pragma.
type WasmImport struct {
	Module string
	Name   string
}

func NewFunc(pos src.XPos) *Func

func (f *Func) Type() *types.Type
func (f *Func) Sym() *types.Sym
func (f *Func) Linksym() *obj.LSym
func (f *Func) LinksymABI(abi obj.ABI) *obj.LSym

// An Inline holds fields used for function bodies that can be inlined.
type Inline struct {
	Cost int32

	// Copies of Func.Dcl and Func.Body for use during inlining. Copies are
	// needed because the function's dcl/body may be changed by later compiler
	// transformations. These fields are also populated when a function from
	// another package is imported.
	Dcl  []*Name
	Body []Node

	// CanDelayResults reports whether it's safe for the inliner to delay
	// initializing the result parameters until immediately before the
	// "return" statement.
	CanDelayResults bool
}

// A Mark represents a scope boundary.
type Mark struct {
	// Pos is the position of the token that marks the scope
	// change.
	Pos src.XPos

	// Scope identifies the innermost scope to the right of Pos.
	Scope ScopeID
}

// A ScopeID represents a lexical scope within a function.
type ScopeID int32

type SymAndPos struct {
	Sym *obj.LSym
	Pos src.XPos
}

func (f *Func) Dupok() bool
func (f *Func) Wrapper() bool
func (f *Func) ABIWrapper() bool
func (f *Func) Needctxt() bool
func (f *Func) ReflectMethod() bool
func (f *Func) IsHiddenClosure() bool
func (f *Func) IsDeadcodeClosure() bool
func (f *Func) HasDefer() bool
func (f *Func) NilCheckDisabled() bool
func (f *Func) InlinabilityChecked() bool
func (f *Func) ExportInline() bool
func (f *Func) InstrumentBody() bool
func (f *Func) OpenCodedDeferDisallowed() bool
func (f *Func) ClosureCalled() bool
func (f *Func) IsPackageInit() bool

func (f *Func) SetDupok(b bool)
func (f *Func) SetWrapper(b bool)
func (f *Func) SetABIWrapper(b bool)
func (f *Func) SetNeedctxt(b bool)
func (f *Func) SetReflectMethod(b bool)
func (f *Func) SetIsHiddenClosure(b bool)
func (f *Func) SetIsDeadcodeClosure(b bool)
func (f *Func) SetHasDefer(b bool)
func (f *Func) SetNilCheckDisabled(b bool)
func (f *Func) SetInlinabilityChecked(b bool)
func (f *Func) SetExportInline(b bool)
func (f *Func) SetInstrumentBody(b bool)
func (f *Func) SetOpenCodedDeferDisallowed(b bool)
func (f *Func) SetClosureCalled(b bool)
func (f *Func) SetIsPackageInit(b bool)

func (f *Func) SetWBPos(pos src.XPos)

// FuncName returns the name (without the package) of the function f.
func FuncName(f *Func) string

// PkgFuncName returns the name of the function referenced by f, with package
// prepended.
//
// This differs from the compiler's internal convention where local functions
// lack a package. This is primarily useful when the ultimate consumer of this
// is a human looking at message.
func PkgFuncName(f *Func) string

// LinkFuncName returns the name of the function f, as it will appear in the
// symbol table of the final linked binary.
func LinkFuncName(f *Func) string

// IsEqOrHashFunc reports whether f is type eq/hash function.
func IsEqOrHashFunc(f *Func) bool

var CurFunc *Func

// WithFunc invokes do with CurFunc and base.Pos set to curfn and
// curfn.Pos(), respectively, and then restores their previous values
// before returning.
func WithFunc(curfn *Func, do func())

func FuncSymName(s *types.Sym) string

// MarkFunc marks a node as a function.
func MarkFunc(n *Name)

// ClosureDebugRuntimeCheck applies boilerplate checks for debug flags
// and compiling runtime.
func ClosureDebugRuntimeCheck(clo *ClosureExpr)

// IsTrivialClosure reports whether closure clo has an
// empty list of captured vars.
func IsTrivialClosure(clo *ClosureExpr) bool

// NewClosureFunc creates a new Func to represent a function literal.
// If hidden is true, then the closure is marked hidden (i.e., as a
// function literal contained within another function, rather than a
// package-scope variable initialization expression).
func NewClosureFunc(pos src.XPos, hidden bool) *Func

// NameClosure generates a unique for the given function literal,
// which must have appeared within outerfn.
func NameClosure(clo *ClosureExpr, outerfn *Func)

// UseClosure checks that the given function literal has been setup
// correctly, and then returns it as an expression.
// It must be called after clo.Func.ClosureVars has been set.
func UseClosure(clo *ClosureExpr, pkg *Package) Node
