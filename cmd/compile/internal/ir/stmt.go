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

// A Decl is a declaration of a const, type, or var. (A declared func is a Func.)
type Decl struct {
	miniNode
	X *Name
}

func NewDecl(pos src.XPos, op Op, x *Name) *Decl

// A Stmt is a Node that can appear as a statement.
// This includes statement-like expressions such as f().
//
// (It's possible it should include <-c, but that would require
// splitting ORECV out of UnaryExpr, which hasn't yet been
// necessary. Maybe instead we will introduce ExprStmt at
// some point.)
type Stmt interface {
	Node
	isStmt()
	PtrInit() *Nodes
}

// An AssignListStmt is an assignment statement with
// more than one item on at least one side: Lhs = Rhs.
// If Def is true, the assignment is a :=.
type AssignListStmt struct {
	miniStmt
	Lhs Nodes
	Def bool
	Rhs Nodes
}

func NewAssignListStmt(pos src.XPos, op Op, lhs, rhs []Node) *AssignListStmt

func (n *AssignListStmt) SetOp(op Op)

// An AssignStmt is a simple assignment statement: X = Y.
// If Def is true, the assignment is a :=.
type AssignStmt struct {
	miniStmt
	X   Node
	Def bool
	Y   Node
}

func NewAssignStmt(pos src.XPos, x, y Node) *AssignStmt

func (n *AssignStmt) SetOp(op Op)

// An AssignOpStmt is an AsOp= assignment statement: X AsOp= Y.
type AssignOpStmt struct {
	miniStmt
	X      Node
	AsOp   Op
	Y      Node
	IncDec bool
}

func NewAssignOpStmt(pos src.XPos, asOp Op, x, y Node) *AssignOpStmt

// A BlockStmt is a block: { List }.
type BlockStmt struct {
	miniStmt
	List Nodes
}

func NewBlockStmt(pos src.XPos, list []Node) *BlockStmt

// A BranchStmt is a break, continue, fallthrough, or goto statement.
type BranchStmt struct {
	miniStmt
	Label *types.Sym
}

func NewBranchStmt(pos src.XPos, op Op, label *types.Sym) *BranchStmt

func (n *BranchStmt) SetOp(op Op)

func (n *BranchStmt) Sym() *types.Sym

// A CaseClause is a case statement in a switch or select: case List: Body.
type CaseClause struct {
	miniStmt
	Var  *Name
	List Nodes

	// RTypes is a list of RType expressions, which are copied to the
	// corresponding OEQ nodes that are emitted when switch statements
	// are desugared. RTypes[i] must be non-nil if the emitted
	// comparison for List[i] will be a mixed interface/concrete
	// comparison; see reflectdata.CompareRType for details.
	//
	// Because mixed interface/concrete switch cases are rare, we allow
	// len(RTypes) < len(List). Missing entries are implicitly nil.
	RTypes Nodes

	Body Nodes
}

func NewCaseStmt(pos src.XPos, list, body []Node) *CaseClause

type CommClause struct {
	miniStmt
	Comm Node
	Body Nodes
}

func NewCommStmt(pos src.XPos, comm Node, body []Node) *CommClause

// A ForStmt is a non-range for loop: for Init; Cond; Post { Body }
type ForStmt struct {
	miniStmt
	Label        *types.Sym
	Cond         Node
	Post         Node
	Body         Nodes
	DistinctVars bool
}

func NewForStmt(pos src.XPos, init Node, cond, post Node, body []Node, distinctVars bool) *ForStmt

// A GoDeferStmt is a go or defer statement: go Call / defer Call.
//
// The two opcodes use a single syntax because the implementations
// are very similar: both are concerned with saving Call and running it
// in a different context (a separate goroutine or a later time).
type GoDeferStmt struct {
	miniStmt
	Call    Node
	DeferAt Expr
}

func NewGoDeferStmt(pos src.XPos, op Op, call Node) *GoDeferStmt

// An IfStmt is a return statement: if Init; Cond { Body } else { Else }.
type IfStmt struct {
	miniStmt
	Cond   Node
	Body   Nodes
	Else   Nodes
	Likely bool
}

func NewIfStmt(pos src.XPos, cond Node, body, els []Node) *IfStmt

// A JumpTableStmt is used to implement switches. Its semantics are:
//
//	tmp := jt.Idx
//	if tmp == Cases[0] goto Targets[0]
//	if tmp == Cases[1] goto Targets[1]
//	...
//	if tmp == Cases[n] goto Targets[n]
//
// Note that a JumpTableStmt is more like a multiway-goto than
// a multiway-if. In particular, the case bodies are just
// labels to jump to, not full Nodes lists.
type JumpTableStmt struct {
	miniStmt

	// Value used to index the jump table.
	// We support only integer types that
	// are at most the size of a uintptr.
	Idx Node

	// If Idx is equal to Cases[i], jump to Targets[i].
	// Cases entries must be distinct and in increasing order.
	// The length of Cases and Targets must be equal.
	Cases   []constant.Value
	Targets []*types.Sym
}

func NewJumpTableStmt(pos src.XPos, idx Node) *JumpTableStmt

// An InterfaceSwitchStmt is used to implement type switches.
// Its semantics are:
//
//	if RuntimeType implements Descriptor.Cases[0] {
//	    Case, Itab = 0, itab<RuntimeType, Descriptor.Cases[0]>
//	} else if RuntimeType implements Descriptor.Cases[1] {
//	    Case, Itab = 1, itab<RuntimeType, Descriptor.Cases[1]>
//	...
//	} else if RuntimeType implements Descriptor.Cases[N-1] {
//	    Case, Itab = N-1, itab<RuntimeType, Descriptor.Cases[N-1]>
//	} else {
//	    Case, Itab = len(cases), nil
//	}
//
// RuntimeType must be a non-nil *runtime._type.
// Hash must be the hash field of RuntimeType (or its copy loaded from an itab).
// Descriptor must represent an abi.InterfaceSwitch global variable.
type InterfaceSwitchStmt struct {
	miniStmt

	Case        Node
	Itab        Node
	RuntimeType Node
	Hash        Node
	Descriptor  *obj.LSym
}

func NewInterfaceSwitchStmt(pos src.XPos, case_, itab, runtimeType, hash Node, descriptor *obj.LSym) *InterfaceSwitchStmt

// An InlineMarkStmt is a marker placed just before an inlined body.
type InlineMarkStmt struct {
	miniStmt
	Index int64
}

func NewInlineMarkStmt(pos src.XPos, index int64) *InlineMarkStmt

func (n *InlineMarkStmt) Offset() int64
func (n *InlineMarkStmt) SetOffset(x int64)

// A LabelStmt is a label statement (just the label, not including the statement it labels).
type LabelStmt struct {
	miniStmt
	Label *types.Sym
}

func NewLabelStmt(pos src.XPos, label *types.Sym) *LabelStmt

func (n *LabelStmt) Sym() *types.Sym

// A RangeStmt is a range loop: for Key, Value = range X { Body }
type RangeStmt struct {
	miniStmt
	Label        *types.Sym
	Def          bool
	X            Node
	RType        Node `mknode:"-"`
	Key          Node
	Value        Node
	Body         Nodes
	DistinctVars bool
	Prealloc     *Name

	// When desugaring the RangeStmt during walk, the assignments to Key
	// and Value may require OCONVIFACE operations. If so, these fields
	// will be copied to their respective ConvExpr fields.
	KeyTypeWord   Node `mknode:"-"`
	KeySrcRType   Node `mknode:"-"`
	ValueTypeWord Node `mknode:"-"`
	ValueSrcRType Node `mknode:"-"`
}

func NewRangeStmt(pos src.XPos, key, value, x Node, body []Node, distinctVars bool) *RangeStmt

// A ReturnStmt is a return statement.
type ReturnStmt struct {
	miniStmt
	Results Nodes
}

func NewReturnStmt(pos src.XPos, results []Node) *ReturnStmt

// A SelectStmt is a block: { Cases }.
type SelectStmt struct {
	miniStmt
	Label *types.Sym
	Cases []*CommClause

	// TODO(rsc): Instead of recording here, replace with a block?
	Compiled Nodes
}

func NewSelectStmt(pos src.XPos, cases []*CommClause) *SelectStmt

// A SendStmt is a send statement: X <- Y.
type SendStmt struct {
	miniStmt
	Chan  Node
	Value Node
}

func NewSendStmt(pos src.XPos, ch, value Node) *SendStmt

// A SwitchStmt is a switch statement: switch Init; Tag { Cases }.
type SwitchStmt struct {
	miniStmt
	Tag   Node
	Cases []*CaseClause
	Label *types.Sym

	// TODO(rsc): Instead of recording here, replace with a block?
	Compiled Nodes
}

func NewSwitchStmt(pos src.XPos, tag Node, cases []*CaseClause) *SwitchStmt

// A TailCallStmt is a tail call statement, which is used for back-end
// code generation to jump directly to another function entirely.
type TailCallStmt struct {
	miniStmt
	Call *CallExpr
}

func NewTailCallStmt(pos src.XPos, call *CallExpr) *TailCallStmt

// A TypeSwitchGuard is the [Name :=] X.(type) in a type switch.
type TypeSwitchGuard struct {
	miniNode
	Tag  *Ident
	X    Node
	Used bool
}

func NewTypeSwitchGuard(pos src.XPos, tag *Ident, x Node) *TypeSwitchGuard
