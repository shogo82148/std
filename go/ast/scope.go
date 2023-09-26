// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements scopes and the objects they contain.

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// A Scope maintains the set of named language entities declared
// in the scope and a link to the immediately surrounding (outer)
// scope.
type Scope struct {
	Outer   *Scope
	Objects map[string]*Object
}

// NewScope creates a new scope nested in the outer scope.
func NewScope(outer *Scope) *Scope

// Lookup returns the object with the given name if it is
// found in scope s, otherwise it returns nil. Outer scopes
// are ignored.
func (s *Scope) Lookup(name string) *Object

// Insert attempts to insert a named object obj into the scope s.
// If the scope already contains an object alt with the same name,
// Insert leaves the scope unchanged and returns alt. Otherwise
// it inserts obj and returns nil.
func (s *Scope) Insert(obj *Object) (alt *Object)

// Debugging support
func (s *Scope) String() string

// An Object describes a named language entity such as a package,
// constant, type, variable, function (incl. methods), or label.
//
// The Data fields contains object-specific data:
//
//	Kind    Data type         Data value
//	Pkg	*types.Package    package scope
//	Con     int               iota for the respective declaration
//	Con     != nil            constant value
//	Typ     *Scope            (used as method scope during type checking - transient)
type Object struct {
	Kind ObjKind
	Name string
	Decl interface{}
	Data interface{}
	Type interface{}
}

// NewObj creates a new object of a given kind and name.
func NewObj(kind ObjKind, name string) *Object

// Pos computes the source position of the declaration of an object name.
// The result may be an invalid position if it cannot be computed
// (obj.Decl may be nil or not correct).
func (obj *Object) Pos() token.Pos

// ObjKind describes what an object represents.
type ObjKind int

// The list of possible Object kinds.
const (
	Bad ObjKind = iota
	Pkg
	Con
	Typ
	Var
	Fun
	Lbl
)

func (kind ObjKind) String() string
