// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

func AssignExpr(curfunc *ir.Func, n ir.Node) ir.Node
func Expr(curfunc *ir.Func, n ir.Node) ir.Node
func Stmt(curfunc *ir.Func, n ir.Node) ir.Node

func Exprs(curfunc *ir.Func, exprs []ir.Node)
func Stmts(curfunc *ir.Func, stmts []ir.Node)

func Call(curfunc *ir.Func, pos src.XPos, callee ir.Node, args []ir.Node, dots bool) ir.Node

func Callee(curfunc *ir.Func, n ir.Node) ir.Node

// RewriteNonNameCall replaces non-Name call expressions with temps,
// rewriting f()(...) to t0 := f(); t0(...).
func RewriteNonNameCall(curfunc *ir.Func, n *ir.CallExpr)

// RewriteMultiValueCall rewrites multi-valued f() to use temporaries,
// so the backend wouldn't need to worry about tuple-valued expressions.
func RewriteMultiValueCall(curfunc *ir.Func, n ir.InitNode, call ir.Node)

// Lookdot1 looks up the specified method s in the list fs of methods, returning
// the matching field or nil. If dostrcmp is 0, it matches the symbols. If
// dostrcmp is 1, it matches by name exactly. If dostrcmp is 2, it matches names
// with case folding.
func Lookdot1(errnode ir.Node, s *types.Sym, t *types.Type, fs []*types.Field, dostrcmp int) *types.Field

// NewMethodExpr returns an OMETHEXPR node representing method
// expression "recv.sym".
func NewMethodExpr(pos src.XPos, recv *types.Type, sym *types.Sym) *ir.SelectorExpr

// Lookdot looks up field or method n.Sel in the type t and returns the matching
// field. It transforms the op of node n to ODOTINTER or ODOTMETH, if appropriate.
// It also may add a StarExpr node to n.X as needed for access to non-pointer
// methods. If dostrcmp is 0, it matches the field/method with the exact symbol
// as n.Sel (appropriate for exported fields). If dostrcmp is 1, it matches by name
// exactly. If dostrcmp is 2, it matches names with case folding.
func Lookdot(curfunc *ir.Func, n *ir.SelectorExpr, t *types.Type, dostrcmp int) *types.Field

func Conv(curfunc *ir.Func, n ir.Node, t *types.Type) ir.Node

// ConvNop converts node n to type t using the OCONVNOP op
// and typechecks the result with ctxExpr.
func ConvNop(curfunc *ir.Func, n ir.Node, t *types.Type) ir.Node
