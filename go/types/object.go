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

	color() color

	setType(Type)

	setOrder(uint32)

	setColor(color color)

	setParent(*Scope)

	sameId(pkg *Package, name string) bool

	scopePos() token.Pos

	setScopePos(pos token.Pos)
}

// Id returns name if it is exported, otherwise it
// returns the name qualified with the package path.
func Id(pkg *Package, name string) string

// An object implements the common parts of an Object.

// color encodes the color of an object (see Checker.objDecl for details).

// An object may be painted in one of three colors.
// Color values other than white or black are considered grey.

// A PkgName represents an imported Go package.
// PkgNames don't have a type.
type PkgName struct {
	object
	imported *Package
	used     bool
}

// NewPkgName returns a new PkgName object representing an imported package.
// The remaining arguments set the attributes found with all Objects.
func NewPkgName(pos token.Pos, pkg *Package, name string, imported *Package) *PkgName

// Imported returns the package that was imported.
// It is distinct from Pkg(), which is the package containing the import statement.
func (obj *PkgName) Imported() *Package

// A Const represents a declared constant.
type Const struct {
	object
	val constant.Value
}

// NewConst returns a new constant with value val.
// The remaining arguments set the attributes found with all Objects.
func NewConst(pos token.Pos, pkg *Package, name string, typ Type, val constant.Value) *Const

// Val returns the constant's value.
func (obj *Const) Val() constant.Value

// A TypeName represents a name for a (defined or alias) type.
type TypeName struct {
	object
}

// NewTypeName returns a new type name denoting the given typ.
// The remaining arguments set the attributes found with all Objects.
//
// The typ argument may be a defined (Named) type or an alias type.
// It may also be nil such that the returned TypeName can be used as
// argument for NewNamed, which will set the TypeName's type as a side-
// effect.
func NewTypeName(pos token.Pos, pkg *Package, name string, typ Type) *TypeName

// IsAlias reports whether obj is an alias name for a type.
func (obj *TypeName) IsAlias() bool

// A Variable represents a declared variable (including function parameters and results, and struct fields).
type Var struct {
	object
	embedded bool
	isField  bool
	used     bool
}

// NewVar returns a new variable.
// The arguments set the attributes found with all Objects.
func NewVar(pos token.Pos, pkg *Package, name string, typ Type) *Var

// NewParam returns a new variable representing a function parameter.
func NewParam(pos token.Pos, pkg *Package, name string, typ Type) *Var

// NewField returns a new variable representing a struct field.
// For embedded fields, the name is the unqualified type name
// / under which the field is accessible.
func NewField(pos token.Pos, pkg *Package, name string, typ Type, embedded bool) *Var

// Anonymous reports whether the variable is an embedded field.
// Same as Embedded; only present for backward-compatibility.
func (obj *Var) Anonymous() bool

// Embedded reports whether the variable is an embedded field.
func (obj *Var) Embedded() bool

// IsField reports whether the variable is a struct field.
func (obj *Var) IsField() bool

// A Func represents a declared function, concrete method, or abstract
// (interface) method. Its Type() is always a *Signature.
// An abstract method may belong to many interfaces due to embedding.
type Func struct {
	object
	hasPtrRecv bool
}

// NewFunc returns a new function with the given signature, representing
// the function's type.
func NewFunc(pos token.Pos, pkg *Package, name string, sig *Signature) *Func

// FullName returns the package- or receiver-type-qualified name of
// function or method obj.
func (obj *Func) FullName() string

// Scope returns the scope of the function's body block.
func (obj *Func) Scope() *Scope

// A Label represents a declared label.
// Labels don't have a type.
type Label struct {
	object
	used bool
}

// NewLabel returns a new label.
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
