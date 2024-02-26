// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeparams

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

func PackIndexExpr(x ast.Expr, lbrack token.Pos, exprs []ast.Expr, rbrack token.Pos) ast.Expr

// IndexExpr wraps an ast.IndexExpr or ast.IndexListExpr.
//
// Orig holds the original ast.Expr from which this IndexExpr was derived.
//
// Note: IndexExpr (intentionally) does not wrap ast.Expr, as that leads to
// accidental misuse such as encountered in golang/go#63933.
//
// TODO(rfindley): remove this helper, in favor of just having a helper
// function that returns indices.
type IndexExpr struct {
	Orig    ast.Expr
	X       ast.Expr
	Lbrack  token.Pos
	Indices []ast.Expr
	Rbrack  token.Pos
}

func (x *IndexExpr) Pos() token.Pos

func UnpackIndexExpr(n ast.Node) *IndexExpr
