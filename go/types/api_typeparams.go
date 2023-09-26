// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build typeparams
// +build typeparams

package types

import (
	"github.com/shogo82148/std/go/ast"
)

type (
	Inferred  = _Inferred
	Sum       = _Sum
	TypeParam = _TypeParam
)

func NewSum(types []Type) Type

func (s *Signature) TParams() []*TypeName
func (s *Signature) SetTParams(tparams []*TypeName)

func (t *Interface) HasTypeList() bool
func (t *Interface) IsComparable() bool
func (t *Interface) IsConstraint() bool

func (t *Named) TParams() []*TypeName
func (t *Named) TArgs() []Type
func (t *Named) SetTArgs(args []Type)

// Info is documented in api_notypeparams.go.
type Info struct {
	Types map[ast.Expr]TypeAndValue

	Inferred map[ast.Expr]_Inferred

	Defs       map[*ast.Ident]Object
	Uses       map[*ast.Ident]Object
	Implicits  map[ast.Node]Object
	Selections map[*ast.SelectorExpr]*Selection
	Scopes     map[ast.Node]*Scope
	InitOrder  []*Initializer
}
