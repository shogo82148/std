// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/go/constant"
	"github.com/shogo82148/std/go/token"
)

// An Object describes a named language entity such as a package,
// constant, type, variable, function (incl. methods), or label.
// All objects implement the Object interface.
type Object interface {
	Parent() *Scope
	Pos() token.Pos
	Pkg() *Package
	Name() string
	Type() Type
	Exported() bool
	Id() string

	String() string

	order() uint32

	setOrder(uint32)

	setParent(*Scope)

	sameId(pkg *Package, name string) bool

	scopePos() token.Pos

	setScopePos(pos token.Pos)
}

// Id returns name if it is exported, otherwise it
// returns the name qualified with the package path.
func Id(pkg *Package, name string) string

// An object implements the common parts of an Object.

// A PkgName represents an imported Go package.
type PkgName struct {
	object
	imported *Package
	used     bool
}

func NewPkgName(pos token.Pos, pkg *Package, name string, imported *Package) *PkgName

// Imported returns the package that was imported.
// It is distinct from Pkg(), which is the package containing the import statement.
func (obj *PkgName) Imported() *Package

// A Const represents a declared constant.
type Const struct {
	object
	val     constant.Value
	visited bool
}

func NewConst(pos token.Pos, pkg *Package, name string, typ Type, val constant.Value) *Const

func (obj *Const) Val() constant.Value

// A TypeName represents a declared type.
type TypeName struct {
	object
}

func NewTypeName(pos token.Pos, pkg *Package, name string, typ Type) *TypeName

// A Variable represents a declared variable (including function parameters and results, and struct fields).
type Var struct {
	object
	anonymous bool
	visited   bool
	isField   bool
	used      bool
}

func NewVar(pos token.Pos, pkg *Package, name string, typ Type) *Var

func NewParam(pos token.Pos, pkg *Package, name string, typ Type) *Var

func NewField(pos token.Pos, pkg *Package, name string, typ Type, anonymous bool) *Var

func (obj *Var) Anonymous() bool
func (obj *Var) IsField() bool

// A Func represents a declared function, concrete method, or abstract
// (interface) method. Its Type() is always a *Signature.
// An abstract method may belong to many interfaces due to embedding.
type Func struct {
	object
}

func NewFunc(pos token.Pos, pkg *Package, name string, sig *Signature) *Func

// FullName returns the package- or receiver-type-qualified name of
// function or method obj.
func (obj *Func) FullName() string

func (obj *Func) Scope() *Scope

// An Alias represents a declared alias.

// A Label represents a declared label.
type Label struct {
	object
	used bool
}

func NewLabel(pos token.Pos, pkg *Package, name string) *Label

// A Builtin represents a built-in function.
// Builtins don't have a valid type.
type Builtin struct {
	object
	id builtinId
}

// Nil represents the predeclared value nil.
type Nil struct {
	object
}

// ObjectString returns the string form of obj.
// The Qualifier controls the printing of
// package-level objects, and may be nil.
func ObjectString(obj Object, qf Qualifier) string

func (obj *PkgName) String() string
func (obj *Const) String() string
func (obj *TypeName) String() string
func (obj *Var) String() string
func (obj *Func) String() string

func (obj *Label) String() string
func (obj *Builtin) String() string
func (obj *Nil) String() string
