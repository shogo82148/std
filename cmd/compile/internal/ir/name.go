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
	Opt       interface{}
	Embed     *[]Embed

	// For a local variable (not param) or extern, the initializing assignment (OAS or OAS2).
	// For a closure var, the ONAME node of the outer captured variable.
	// For the case-local variables of a type switch, the type switch guard (OTYPESW).
	// For a range variable, the range statement (ORANGE)
	// For a recv variable in a case of a select statement, the receive assignment (OSELRECV2)
	// For the name of a function, points to corresponding Func node.
	Defn Node

	// The function, method, or closure in which local variable or param is declared.
	Curfn *Func

	Heapaddr *Name

	// ONAME closure linkage
	// Consider:
	//
	//	func f() {
	//		x := 1 // x1
	//		func() {
	//			use(x) // x2
	//			func() {
	//				use(x) // x3
	//				--- parser is here ---
	//			}()
	//		}()
	//	}
	//
	// There is an original declaration of x and then a chain of mentions of x
	// leading into the current function. Each time x is mentioned in a new closure,
	// we create a variable representing x for use in that specific closure,
	// since the way you get to x is different in each closure.
	//
	// Let's number the specific variables as shown in the code:
	// x1 is the original x, x2 is when mentioned in the closure,
	// and x3 is when mentioned in the closure in the closure.
	//
	// We keep these linked (assume N > 1):
	//
	//   - x1.Defn = original declaration statement for x (like most variables)
	//   - x1.Innermost = current innermost closure x (in this case x3), or nil for none
	//   - x1.IsClosureVar() = false
	//
	//   - xN.Defn = x1, N > 1
	//   - xN.IsClosureVar() = true, N > 1
	//   - x2.Outer = nil
	//   - xN.Outer = x(N-1), N > 2
	//
	//
	// When we look up x in the symbol table, we always get x1.
	// Then we can use x1.Innermost (if not nil) to get the x
	// for the innermost known closure function,
	// but the first reference in a closure will find either no x1.Innermost
	// or an x1.Innermost with .Funcdepth < Funcdepth.
	// In that case, a new xN must be created, linked in with:
	//
	//     xN.Defn = x1
	//     xN.Outer = x1.Innermost
	//     x1.Innermost = xN
	//
	// When we finish the function, we'll process its closure variables
	// and find xN and pop it off the list using:
	//
	//     x1 := xN.Defn
	//     x1.Innermost = xN.Outer
	//
	// We leave x1.Innermost set so that we can still get to the original
	// variable quickly. Not shown here, but once we're
	// done parsing a function and no longer need xN.Outer for the
	// lexical x reference links as described above, funcLit
	// recomputes xN.Outer as the semantic x reference link tree,
	// even filling in x in intermediate closures that might not
	// have mentioned it along the way to inner closures that did.
	// See funcLit for details.
	//
	// During the eventual compilation, then, for closure variables we have:
	//
	//     xN.Defn = original variable
	//     xN.Outer = variable captured in next outward scope
	//                to make closure where xN appears
	//
	// Because of the sharding of pieces of the node, x.Defn means x.Name.Defn
	// and x.Innermost/Outer means x.Name.Param.Innermost/Outer.
	Innermost *Name
	Outer     *Name
}

// RecordFrameOffset records the frame offset for the name.
// It is used by package types when laying out function arguments.
func (n *Name) RecordFrameOffset(offset int64)

// NewNameAt returns a new ONAME Node associated with symbol s at position pos.
// The caller is responsible for setting Curfn.
func NewNameAt(pos src.XPos, sym *types.Sym) *Name

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
func (n *Name) Offset() int64
func (n *Name) SetOffset(x int64)

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
func (n *Name) CoverageCounter() bool
func (n *Name) CoverageAuxVar() bool

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
func (n *Name) SetCoverageCounter(b bool)
func (n *Name) SetCoverageAuxVar(b bool)

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

// CaptureName returns a Name suitable for referring to n from within function
// fn or from the package block if fn is nil. If n is a free variable declared
// within a function that encloses fn, then CaptureName returns the closure
// variable that refers to n within fn, creating it if necessary.
// Otherwise, it simply returns n.
func CaptureName(pos src.XPos, fn *Func, n *Name) *Name

// FinishCaptureNames handles any work leftover from calling CaptureName
// earlier. outerfn should be the function that immediately encloses fn.
func FinishCaptureNames(pos src.XPos, outerfn, fn *Func)

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
