// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !typeparams
// +build !typeparams

package ast

import "github.com/shogo82148/std/go/token"

type (
	// A FuncType node represents a function type.
	FuncType struct {
		Func    token.Pos
		Params  *FieldList
		Results *FieldList
	}

	// A TypeSpec node represents a type declaration (TypeSpec production).
	TypeSpec struct {
		Doc     *CommentGroup
		Name    *Ident
		Assign  token.Pos
		Type    Expr
		Comment *CommentGroup
	}
)
