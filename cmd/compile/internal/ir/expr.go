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

// An Expr is a Node that can appear as an expression.
type Expr interface {
	Node
	isExpr()
}

// An AddStringExpr is a string concatenation List[0] + List[1] + ... + List[len(List)-1].
type AddStringExpr struct {
	miniExpr
	List     Nodes
	Prealloc *Name
}

func NewAddStringExpr(pos src.XPos, list []Node) *AddStringExpr

// An AddrExpr is an address-of expression &X.
// It may end up being a normal address-of or an allocation of a composite literal.
type AddrExpr struct {
	miniExpr
	X        Node
	Prealloc *Name
}

func NewAddrExpr(pos src.XPos, x Node) *AddrExpr

func (n *AddrExpr) Implicit() bool
func (n *AddrExpr) SetImplicit(b bool)

func (n *AddrExpr) SetOp(op Op)

// A BasicLit is a literal of basic type.
type BasicLit struct {
	miniExpr
	val constant.Value
}

// NewBasicLit returns an OLITERAL representing val with the given type.
func NewBasicLit(pos src.XPos, typ *types.Type, val constant.Value) Node

func (n *BasicLit) Val() constant.Value
func (n *BasicLit) SetVal(val constant.Value)

// NewConstExpr returns an OLITERAL representing val, copying the
// position and type from orig.
func NewConstExpr(val constant.Value, orig Node) Node

// A BinaryExpr is a binary expression X Op Y,
// or Op(X, Y) for builtin functions that do not become calls.
type BinaryExpr struct {
	miniExpr
	X     Node
	Y     Node
	RType Node `mknode:"-"`
}

func NewBinaryExpr(pos src.XPos, op Op, x, y Node) *BinaryExpr

func (n *BinaryExpr) SetOp(op Op)

// A CallExpr is a function call Fun(Args).
type CallExpr struct {
	miniExpr
	Fun       Node
	Args      Nodes
	DeferAt   Node
	RType     Node `mknode:"-"`
	KeepAlive []*Name
	IsDDD     bool
	NoInline  bool
}

func NewCallExpr(pos src.XPos, op Op, fun Node, args []Node) *CallExpr

func (n *CallExpr) SetOp(op Op)

// A ClosureExpr is a function literal expression.
type ClosureExpr struct {
	miniExpr
	Func     *Func `mknode:"-"`
	Prealloc *Name
	IsGoWrap bool
}

// A CompLitExpr is a composite literal Type{Vals}.
// Before type-checking, the type is Ntype.
type CompLitExpr struct {
	miniExpr
	List     Nodes
	RType    Node `mknode:"-"`
	Prealloc *Name
	// For OSLICELIT, Len is the backing array length.
	// For OMAPLIT, Len is the number of entries that we've removed from List and
	// generated explicit mapassign calls for. This is used to inform the map alloc hint.
	Len int64
}

func NewCompLitExpr(pos src.XPos, op Op, typ *types.Type, list []Node) *CompLitExpr

func (n *CompLitExpr) Implicit() bool
func (n *CompLitExpr) SetImplicit(b bool)

func (n *CompLitExpr) SetOp(op Op)

// A ConvExpr is a conversion Type(X).
// It may end up being a value or a type.
type ConvExpr struct {
	miniExpr
	X Node

	// For implementing OCONVIFACE expressions.
	//
	// TypeWord is an expression yielding a *runtime._type or
	// *runtime.itab value to go in the type word of the iface/eface
	// result. See reflectdata.ConvIfaceTypeWord for further details.
	//
	// SrcRType is an expression yielding a *runtime._type value for X,
	// if it's not pointer-shaped and needs to be heap allocated.
	TypeWord Node `mknode:"-"`
	SrcRType Node `mknode:"-"`

	// For -d=checkptr instrumentation of conversions from
	// unsafe.Pointer to *Elem or *[Len]Elem.
	//
	// TODO(mdempsky): We only ever need one of these, but currently we
	// don't decide which one until walk. Longer term, it probably makes
	// sense to have a dedicated IR op for `(*[Len]Elem)(ptr)[:n:m]`
	// expressions.
	ElemRType     Node `mknode:"-"`
	ElemElemRType Node `mknode:"-"`
}

func NewConvExpr(pos src.XPos, op Op, typ *types.Type, x Node) *ConvExpr

func (n *ConvExpr) Implicit() bool
func (n *ConvExpr) SetImplicit(b bool)
func (n *ConvExpr) CheckPtr() bool
func (n *ConvExpr) SetCheckPtr(b bool)

func (n *ConvExpr) SetOp(op Op)

// An IndexExpr is an index expression X[Index].
type IndexExpr struct {
	miniExpr
	X        Node
	Index    Node
	RType    Node `mknode:"-"`
	Assigned bool
}

func NewIndexExpr(pos src.XPos, x, index Node) *IndexExpr

func (n *IndexExpr) SetOp(op Op)

// A KeyExpr is a Key: Value composite literal key.
type KeyExpr struct {
	miniExpr
	Key   Node
	Value Node
}

func NewKeyExpr(pos src.XPos, key, value Node) *KeyExpr

// A StructKeyExpr is an Field: Value composite literal key.
type StructKeyExpr struct {
	miniExpr
	Field *types.Field
	Value Node
}

func NewStructKeyExpr(pos src.XPos, field *types.Field, value Node) *StructKeyExpr

func (n *StructKeyExpr) Sym() *types.Sym

// An InlinedCallExpr is an inlined function call.
type InlinedCallExpr struct {
	miniExpr
	Body       Nodes
	ReturnVars Nodes
}

func NewInlinedCallExpr(pos src.XPos, body, retvars []Node) *InlinedCallExpr

func (n *InlinedCallExpr) SingleResult() Node

// A LogicalExpr is an expression X Op Y where Op is && or ||.
// It is separate from BinaryExpr to make room for statements
// that must be executed before Y but after X.
type LogicalExpr struct {
	miniExpr
	X Node
	Y Node
}

func NewLogicalExpr(pos src.XPos, op Op, x, y Node) *LogicalExpr

func (n *LogicalExpr) SetOp(op Op)

// A MakeExpr is a make expression: make(Type[, Len[, Cap]]).
// Op is OMAKECHAN, OMAKEMAP, OMAKESLICE, or OMAKESLICECOPY,
// but *not* OMAKE (that's a pre-typechecking CallExpr).
type MakeExpr struct {
	miniExpr
	RType Node `mknode:"-"`
	Len   Node
	Cap   Node
}

func NewMakeExpr(pos src.XPos, op Op, len, cap Node) *MakeExpr

func (n *MakeExpr) SetOp(op Op)

// A NilExpr represents the predefined untyped constant nil.
type NilExpr struct {
	miniExpr
}

func NewNilExpr(pos src.XPos, typ *types.Type) *NilExpr

// A ParenExpr is a parenthesized expression (X).
// It may end up being a value or a type.
type ParenExpr struct {
	miniExpr
	X Node
}

func NewParenExpr(pos src.XPos, x Node) *ParenExpr

func (n *ParenExpr) Implicit() bool
func (n *ParenExpr) SetImplicit(b bool)

// A ResultExpr represents a direct access to a result.
type ResultExpr struct {
	miniExpr
	Index int64
}

func NewResultExpr(pos src.XPos, typ *types.Type, index int64) *ResultExpr

// A LinksymOffsetExpr refers to an offset within a global variable.
// It is like a SelectorExpr but without the field name.
type LinksymOffsetExpr struct {
	miniExpr
	Linksym *obj.LSym
	Offset_ int64
}

func NewLinksymOffsetExpr(pos src.XPos, lsym *obj.LSym, offset int64, typ *types.Type) *LinksymOffsetExpr

// NewLinksymExpr is NewLinksymOffsetExpr, but with offset fixed at 0.
func NewLinksymExpr(pos src.XPos, lsym *obj.LSym, typ *types.Type) *LinksymOffsetExpr

// NewNameOffsetExpr is NewLinksymOffsetExpr, but taking a *Name
// representing a global variable instead of an *obj.LSym directly.
func NewNameOffsetExpr(pos src.XPos, name *Name, offset int64, typ *types.Type) *LinksymOffsetExpr

// A SelectorExpr is a selector expression X.Sel.
type SelectorExpr struct {
	miniExpr
	X Node
	// Sel is the name of the field or method being selected, without (in the
	// case of methods) any preceding type specifier. If the field/method is
	// exported, than the Sym uses the local package regardless of the package
	// of the containing type.
	Sel *types.Sym
	// The actual selected field - may not be filled in until typechecking.
	Selection *types.Field
	Prealloc  *Name
}

func NewSelectorExpr(pos src.XPos, op Op, x Node, sel *types.Sym) *SelectorExpr

func (n *SelectorExpr) SetOp(op Op)

func (n *SelectorExpr) Sym() *types.Sym
func (n *SelectorExpr) Implicit() bool
func (n *SelectorExpr) SetImplicit(b bool)
func (n *SelectorExpr) Offset() int64

func (n *SelectorExpr) FuncName() *Name

// A SliceExpr is a slice expression X[Low:High] or X[Low:High:Max].
type SliceExpr struct {
	miniExpr
	X    Node
	Low  Node
	High Node
	Max  Node
}

func NewSliceExpr(pos src.XPos, op Op, x, low, high, max Node) *SliceExpr

func (n *SliceExpr) SetOp(op Op)

// IsSlice3 reports whether o is a slice3 op (OSLICE3, OSLICE3ARR).
// o must be a slicing op.
func (o Op) IsSlice3() bool

// A SliceHeader expression constructs a slice header from its parts.
type SliceHeaderExpr struct {
	miniExpr
	Ptr Node
	Len Node
	Cap Node
}

func NewSliceHeaderExpr(pos src.XPos, typ *types.Type, ptr, len, cap Node) *SliceHeaderExpr

// A StringHeaderExpr expression constructs a string header from its parts.
type StringHeaderExpr struct {
	miniExpr
	Ptr Node
	Len Node
}

func NewStringHeaderExpr(pos src.XPos, ptr, len Node) *StringHeaderExpr

// A StarExpr is a dereference expression *X.
// It may end up being a value or a type.
type StarExpr struct {
	miniExpr
	X Node
}

func NewStarExpr(pos src.XPos, x Node) *StarExpr

func (n *StarExpr) Implicit() bool
func (n *StarExpr) SetImplicit(b bool)

// A TypeAssertionExpr is a selector expression X.(Type).
// Before type-checking, the type is Ntype.
type TypeAssertExpr struct {
	miniExpr
	X Node

	// Runtime type information provided by walkDotType for
	// assertions from non-empty interface to concrete type.
	ITab Node `mknode:"-"`

	// An internal/abi.TypeAssert descriptor to pass to the runtime.
	Descriptor *obj.LSym
}

func NewTypeAssertExpr(pos src.XPos, x Node, typ *types.Type) *TypeAssertExpr

func (n *TypeAssertExpr) SetOp(op Op)

// A DynamicTypeAssertExpr asserts that X is of dynamic type RType.
type DynamicTypeAssertExpr struct {
	miniExpr
	X Node

	// SrcRType is an expression that yields a *runtime._type value
	// representing X's type. It's used in failed assertion panic
	// messages.
	SrcRType Node

	// RType is an expression that yields a *runtime._type value
	// representing the asserted type.
	//
	// BUG(mdempsky): If ITab is non-nil, RType may be nil.
	RType Node

	// ITab is an expression that yields a *runtime.itab value
	// representing the asserted type within the assertee expression's
	// original interface type.
	//
	// ITab is only used for assertions from non-empty interface type to
	// a concrete (i.e., non-interface) type. For all other assertions,
	// ITab is nil.
	ITab Node
}

func NewDynamicTypeAssertExpr(pos src.XPos, op Op, x, rtype Node) *DynamicTypeAssertExpr

func (n *DynamicTypeAssertExpr) SetOp(op Op)

// A UnaryExpr is a unary expression Op X,
// or Op(X) for a builtin function that does not end up being a call.
type UnaryExpr struct {
	miniExpr
	X Node
}

func NewUnaryExpr(pos src.XPos, op Op, x Node) *UnaryExpr

func (n *UnaryExpr) SetOp(op Op)

func IsZero(n Node) bool

// lvalue etc
func IsAddressable(n Node) bool

// StaticValue analyzes n to find the earliest expression that always
// evaluates to the same value as n, which might be from an enclosing
// function.
//
// For example, given:
//
//	var x int = g()
//	func() {
//		y := x
//		*p = int(y)
//	}
//
// calling StaticValue on the "int(y)" expression returns the outer
// "g()" expression.
func StaticValue(n Node) Node

// Reassigned takes an ONAME node, walks the function in which it is
// defined, and returns a boolean indicating whether the name has any
// assignments other than its declaration.
// NB: global variables are always considered to be re-assigned.
// TODO: handle initial declaration not including an assignment and
// followed by a single assignment?
func Reassigned(name *Name) bool

// StaticCalleeName returns the ONAME/PFUNC for n, if known.
func StaticCalleeName(n Node) *Name

// IsIntrinsicCall reports whether the compiler back end will treat the call as an intrinsic operation.
var IsIntrinsicCall = func(*CallExpr) bool { return false }

// SameSafeExpr checks whether it is safe to reuse one of l and r
// instead of computing both. SameSafeExpr assumes that l and r are
// used in the same statement or expression. In order for it to be
// safe to reuse l or r, they must:
//   - be the same expression
//   - not have side-effects (no function calls, no channel ops);
//     however, panics are ok
//   - not cause inappropriate aliasing; e.g. two string to []byte
//     conversions, must result in two distinct slices
//
// The handling of OINDEXMAP is subtle. OINDEXMAP can occur both
// as an lvalue (map assignment) and an rvalue (map access). This is
// currently OK, since the only place SameSafeExpr gets used on an
// lvalue expression is for OSLICE and OAPPEND optimizations, and it
// is correct in those settings.
func SameSafeExpr(l Node, r Node) bool

// ShouldCheckPtr reports whether pointer checking should be enabled for
// function fn at a given level. See debugHelpFooter for defined
// levels.
func ShouldCheckPtr(fn *Func, level int) bool

// ShouldAsanCheckPtr reports whether pointer checking should be enabled for
// function fn when -asan is enabled.
func ShouldAsanCheckPtr(fn *Func) bool

// IsReflectHeaderDataField reports whether l is an expression p.Data
// where p has type reflect.SliceHeader or reflect.StringHeader.
func IsReflectHeaderDataField(l Node) bool

func ParamNames(ft *types.Type) []Node

// MethodSym returns the method symbol representing a method name
// associated with a specific receiver type.
//
// Method symbols can be used to distinguish the same method appearing
// in different method sets. For example, T.M and (*T).M have distinct
// method symbols.
//
// The returned symbol will be marked as a function.
func MethodSym(recv *types.Type, msym *types.Sym) *types.Sym

// MethodSymSuffix is like MethodSym, but allows attaching a
// distinguisher suffix. To avoid collisions, the suffix must not
// start with a letter, number, or period.
func MethodSymSuffix(recv *types.Type, msym *types.Sym, suffix string) *types.Sym

// LookupMethodSelector returns the types.Sym of the selector for a method
// named in local symbol name, as well as the types.Sym of the receiver.
//
// TODO(prattmic): this does not attempt to handle method suffixes (wrappers).
func LookupMethodSelector(pkg *types.Pkg, name string) (typ, meth *types.Sym, err error)

// MethodExprName returns the ONAME representing the method
// referenced by expression n, which must be a method selector,
// method expression, or method value.
func MethodExprName(n Node) *Name

// MethodExprFunc is like MethodExprName, but returns the types.Field instead.
func MethodExprFunc(n Node) *types.Field
