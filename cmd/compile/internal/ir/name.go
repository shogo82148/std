// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"

	"github.com/shogo82148/std/go/constant"
)

// An Ident is an identifier, possibly qualified.
type Ident struct {
	miniExpr
	sym *types.Sym
}

func NewIdent(pos src.XPos, sym *types.Sym) *Ident

func (n *Ident) Sym() *types.Sym

// Name holds Node fields used only by named nodes (ONAME, OTYPE, some OLITERAL).
type Name struct {
	miniExpr
	BuiltinOp Op
	Class     Class
	pragma    PragmaFlag
	flags     bitset16
	DictIndex uint16
	sym       *types.Sym
	Func      *Func
	Offset_   int64
	val       constant.Value
	Opt       any
	Embed     *[]Embed

	// For a local variable (not param) or extern, the initializing assignment (OAS or OAS2).
	// For a closure var, the ONAME node of the original (outermost) captured variable.
	// For the case-local variables of a type switch, the type switch guard (OTYPESW).
	// For a range variable, the range statement (ORANGE)
	// For a recv variable in a case of a select statement, the receive assignment (OSELRECV2)
	// For the name of a function, points to corresponding Func node.
	Defn Node

	// The function, method, or closure in which local variable or param is declared.
	Curfn *Func

	Heapaddr *Name

	// Outer points to the immediately enclosing function's copy of this
	// closure variable. If not a closure variable, then Outer is nil.
	Outer *Name
}

// RecordFrameOffset records the frame offset for the name.
// It is used by package types when laying out function arguments.
func (n *Name) RecordFrameOffset(offset int64)

// NewNameAt returns a new ONAME Node associated with symbol s at position pos.
// The caller is responsible for setting Curfn.
func NewNameAt(pos src.XPos, sym *types.Sym, typ *types.Type) *Name

// NewBuiltin returns a new Name representing a builtin function,
// either predeclared or from package unsafe.
func NewBuiltin(sym *types.Sym, op Op) *Name

// NewLocal returns a new function-local variable with the given name and type.
func (fn *Func) NewLocal(pos src.XPos, sym *types.Sym, typ *types.Type) *Name

// NewDeclNameAt returns a new Name associated with symbol s at position pos.
// The caller is responsible for setting Curfn.
func NewDeclNameAt(pos src.XPos, op Op, sym *types.Sym) *Name

// NewConstAt returns a new OLITERAL Node associated with symbol s at position pos.
func NewConstAt(pos src.XPos, sym *types.Sym, typ *types.Type, val constant.Value) *Name

func (n *Name) Name() *Name
func (n *Name) Sym() *types.Sym
func (n *Name) SetSym(x *types.Sym)
func (n *Name) SubOp() Op
func (n *Name) SetSubOp(x Op)
func (n *Name) SetFunc(x *Func)
func (n *Name) FrameOffset() int64
func (n *Name) SetFrameOffset(x int64)

func (n *Name) Linksym() *obj.LSym
func (n *Name) LinksymABI(abi obj.ABI) *obj.LSym

func (*Name) CanBeNtype()
func (*Name) CanBeAnSSASym()
func (*Name) CanBeAnSSAAux()

// Pragma returns the PragmaFlag for p, which must be for an OTYPE.
func (n *Name) Pragma() PragmaFlag

// SetPragma sets the PragmaFlag for p, which must be for an OTYPE.
func (n *Name) SetPragma(flag PragmaFlag)

// Alias reports whether p, which must be for an OTYPE, is a type alias.
func (n *Name) Alias() bool

// SetAlias sets whether p, which must be for an OTYPE, is a type alias.
func (n *Name) SetAlias(alias bool)

func (n *Name) Readonly() bool
func (n *Name) Needzero() bool
func (n *Name) AutoTemp() bool
func (n *Name) Used() bool
func (n *Name) IsClosureVar() bool
func (n *Name) IsOutputParamHeapAddr() bool
func (n *Name) IsOutputParamInRegisters() bool
func (n *Name) Addrtaken() bool
func (n *Name) InlFormal() bool
func (n *Name) InlLocal() bool
func (n *Name) OpenDeferSlot() bool
func (n *Name) Libfuzzer8BitCounter() bool
func (n *Name) CoverageAuxVar() bool
func (n *Name) NonMergeable() bool

func (n *Name) SetNeedzero(b bool)
func (n *Name) SetAutoTemp(b bool)
func (n *Name) SetUsed(b bool)
func (n *Name) SetIsClosureVar(b bool)
func (n *Name) SetIsOutputParamHeapAddr(b bool)
func (n *Name) SetIsOutputParamInRegisters(b bool)
func (n *Name) SetAddrtaken(b bool)
func (n *Name) SetInlFormal(b bool)
func (n *Name) SetInlLocal(b bool)
func (n *Name) SetOpenDeferSlot(b bool)
func (n *Name) SetLibfuzzer8BitCounter(b bool)
func (n *Name) SetCoverageAuxVar(b bool)
func (n *Name) SetNonMergeable(b bool)

// OnStack reports whether variable n may reside on the stack.
func (n *Name) OnStack() bool

// MarkReadonly indicates that n is an ONAME with readonly contents.
func (n *Name) MarkReadonly()

// Val returns the constant.Value for the node.
func (n *Name) Val() constant.Value

// SetVal sets the constant.Value for the node.
func (n *Name) SetVal(v constant.Value)

// Canonical returns the logical declaration that n represents. If n
// is a closure variable, then Canonical returns the original Name as
// it appears in the function that immediately contains the
// declaration. Otherwise, Canonical simply returns n itself.
func (n *Name) Canonical() *Name

func (n *Name) SetByval(b bool)

func (n *Name) Byval() bool

// NewClosureVar returns a new closure variable for fn to refer to
// outer variable n.
func NewClosureVar(pos src.XPos, fn *Func, n *Name) *Name

// NewHiddenParam returns a new hidden parameter for fn with the given
// name and type.
func NewHiddenParam(pos src.XPos, fn *Func, sym *types.Sym, typ *types.Type) *Name

// SameSource reports whether two nodes refer to the same source
// element.
//
// It exists to help incrementally migrate the compiler towards
// allowing the introduction of IdentExpr (#42990). Once we have
// IdentExpr, it will no longer be safe to directly compare Node
// values to tell if they refer to the same Name. Instead, code will
// need to explicitly get references to the underlying Name object(s),
// and compare those instead.
//
// It will still be safe to compare Nodes directly for checking if two
// nodes are syntactically the same. The SameSource function exists to
// indicate code that intentionally compares Nodes for syntactic
// equality as opposed to code that has yet to be updated in
// preparation for IdentExpr.
func SameSource(n1, n2 Node) bool

// Uses reports whether expression x is a (direct) use of the given
// variable.
func Uses(x Node, v *Name) bool

// DeclaredBy reports whether expression x refers (directly) to a
// variable that was declared by the given statement.
func DeclaredBy(x, stmt Node) bool

// The Class of a variable/function describes the "storage class"
// of a variable or function. During parsing, storage classes are
// called declaration contexts.
type Class uint8

//go:generate stringer -type=Class name.go
const (
	Pxxx Class = iota
	PEXTERN
	PAUTO
	PAUTOHEAP
	PPARAM
	PPARAMOUT
	PTYPEPARAM
	PFUNC

	// Careful: Class is stored in three bits in Node.flags.
	_ = uint((1 << 3) - iota)
)

type Embed struct {
	Pos      src.XPos
	Patterns []string
}
