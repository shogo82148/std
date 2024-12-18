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
//
// Deprecated: use the type checker [go/types] instead; see [Object].
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
//	Pkg     *Scope            package scope
//	Con     int               iota for the respective declaration
//
// Deprecated: The relationship between Idents and Objects cannot be
// correctly computed without type information. For example, the
// expression T{K: 0} may denote a struct, map, slice, or array
// literal, depending on the type of T. If T is a struct, then K
// refers to a field of T, whereas for the other types it refers to a
// value in the environment.
//
// New programs should set the [parser.SkipObjectResolution] parser
// flag to disable syntactic object resolution (which also saves CPU
// and memory), and instead use the type checker [go/types] if object
// resolution is desired. See the Defs, Uses, and Implicits fields of
// the [types.Info] struct for details.
type Object struct {
	Kind ObjKind
	Name string
	Decl any
	Data any
	Type any
}

// NewObj creates a new object of a given kind and name.
func NewObj(kind ObjKind, name string) *Object

// Pos computes the source position of the declaration of an object name.
// The result may be an invalid position if it cannot be computed
// (obj.Decl may be nil or not correct).
func (obj *Object) Pos() token.Pos

// ObjKind describes what an [Object] represents.
type ObjKind int

// The list of possible [Object] kinds.
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
