// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !typeparams
// +build !typeparams

package types

import "github.com/shogo82148/std/go/ast"

// Info holds result type information for a type-checked package.
// Only the information for which a map is provided is collected.
// If the package has type errors, the collected information may
// be incomplete.
type Info struct {
	Types map[ast.Expr]TypeAndValue

	Defs map[*ast.Ident]Object

	Uses map[*ast.Ident]Object

	Implicits map[ast.Node]Object

	Selections map[*ast.SelectorExpr]*Selection

	Scopes map[ast.Node]*Scope

	InitOrder []*Initializer
}
