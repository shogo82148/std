// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build typeparams
// +build typeparams

package ast

import "github.com/shogo82148/std/go/token"

type (
	// A FuncType node represents a function type.
	FuncType struct {
		Func    token.Pos
		TParams *FieldList
		Params  *FieldList
		Results *FieldList
	}

	// A TypeSpec node represents a type declaration (TypeSpec production).
	TypeSpec struct {
		Doc     *CommentGroup
		Name    *Ident
		TParams *FieldList
		Assign  token.Pos
		Type    Expr
		Comment *CommentGroup
	}

	// A ListExpr node represents a list of expressions separated by commas.
	// ListExpr nodes are used as index in IndexExpr nodes representing type
	// or function instantiations with more than one type argument.
	ListExpr struct {
		ElemList []Expr
	}
)

func (x *ListExpr) Pos() token.Pos

func (x *ListExpr) End() token.Pos
