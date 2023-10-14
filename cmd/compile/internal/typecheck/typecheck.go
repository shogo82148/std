// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/base"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// Function collecting autotmps generated during typechecking,
// to be included in the package-level init function.
var InitTodoFunc = ir.NewFunc(base.Pos)

var TypecheckAllowed bool

var (
	NeedRuntimeType = func(*types.Type) {}
)

func AssignExpr(n ir.Node) ir.Node
func Expr(n ir.Node) ir.Node
func Stmt(n ir.Node) ir.Node

func Exprs(exprs []ir.Node)
func Stmts(stmts []ir.Node)

func Call(pos src.XPos, callee ir.Node, args []ir.Node, dots bool) ir.Node

func Callee(n ir.Node) ir.Node

// Resolve resolves an ONONAME node to a definition, if any. If n is not an ONONAME node,
// Resolve returns n unchanged. If n is an ONONAME node and not in the same package,
// then n.Sym() is resolved using import data. Otherwise, Resolve returns
// n.Sym().Def. An ONONAME node can be created using ir.NewIdent(), so an imported
// symbol can be resolved via Resolve(ir.NewIdent(src.NoXPos, sym)).
func Resolve(n ir.Node) (res ir.Node)

func Func(fn *ir.Func)

// RewriteNonNameCall replaces non-Name call expressions with temps,
// rewriting f()(...) to t0 := f(); t0(...).
func RewriteNonNameCall(n *ir.CallExpr)

// RewriteMultiValueCall rewrites multi-valued f() to use temporaries,
// so the backend wouldn't need to worry about tuple-valued expressions.
func RewriteMultiValueCall(n ir.InitNode, call ir.Node)

// Lookdot1 looks up the specified method s in the list fs of methods, returning
// the matching field or nil. If dostrcmp is 0, it matches the symbols. If
// dostrcmp is 1, it matches by name exactly. If dostrcmp is 2, it matches names
// with case folding.
func Lookdot1(errnode ir.Node, s *types.Sym, t *types.Type, fs *types.Fields, dostrcmp int) *types.Field

// Lookdot looks up field or method n.Sel in the type t and returns the matching
// field. It transforms the op of node n to ODOTINTER or ODOTMETH, if appropriate.
// It also may add a StarExpr node to n.X as needed for access to non-pointer
// methods. If dostrcmp is 0, it matches the field/method with the exact symbol
// as n.Sel (appropriate for exported fields). If dostrcmp is 1, it matches by name
// exactly. If dostrcmp is 2, it matches names with case folding.
func Lookdot(n *ir.SelectorExpr, t *types.Type, dostrcmp int) *types.Field

func Conv(n ir.Node, t *types.Type) ir.Node

// ConvNop converts node n to type t using the OCONVNOP op
// and typechecks the result with ctxExpr.
func ConvNop(n ir.Node, t *types.Type) ir.Node
