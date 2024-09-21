// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/go/constant"
)

// An Object is a named language entity.
// An Object may be a constant ([Const]), type name ([TypeName]),
// variable or struct field ([Var]), function or method ([Func]),
// imported package ([PkgName]), label ([Label]),
// built-in function ([Builtin]),
// or the predeclared identifier 'nil' ([Nil]).
//
// The environment, which is structured as a tree of Scopes,
// maps each name to the unique Object that it denotes.
type Object interface {
	Parent() *Scope
	Pos() syntax.Pos
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

	sameId(pkg *Package, name string, foldCase bool) bool

	scopePos() syntax.Pos

	setScopePos(pos syntax.Pos)
}

// Id returns name if it is exported, otherwise it
// returns the name qualified with the package path.
func Id(pkg *Package, name string) string

// A PkgName represents an imported Go package.
// PkgNames don't have a type.
type PkgName struct {
	object
	imported *Package
	used     bool
}

// NewPkgName returns a new PkgName object representing an imported package.
// The remaining arguments set the attributes found with all Objects.
func NewPkgName(pos syntax.Pos, pkg *Package, name string, imported *Package) *PkgName

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
func NewConst(pos syntax.Pos, pkg *Package, name string, typ Type, val constant.Value) *Const

// Val returns the constant's value.
func (obj *Const) Val() constant.Value

// A TypeName is an [Object] that represents a type with a name:
// a defined type ([Named]),
// an alias type ([Alias]),
// a type parameter ([TypeParam]),
// or a predeclared type such as int or error.
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
func NewTypeName(pos syntax.Pos, pkg *Package, name string, typ Type) *TypeName

// NewTypeNameLazy returns a new defined type like NewTypeName, but it
// lazily calls resolve to finish constructing the Named object.
func NewTypeNameLazy(pos syntax.Pos, pkg *Package, name string, load func(named *Named) (tparams []*TypeParam, underlying Type, methods []*Func)) *TypeName

// IsAlias reports whether obj is an alias name for a type.
func (obj *TypeName) IsAlias() bool

// A Variable represents a declared variable (including function parameters and results, and struct fields).
type Var struct {
	object
	embedded bool
	isField  bool
	used     bool
	origin   *Var
}

// NewVar returns a new variable.
// The arguments set the attributes found with all Objects.
func NewVar(pos syntax.Pos, pkg *Package, name string, typ Type) *Var

// NewParam returns a new variable representing a function parameter.
func NewParam(pos syntax.Pos, pkg *Package, name string, typ Type) *Var

// NewField returns a new variable representing a struct field.
// For embedded fields, the name is the unqualified type name
// under which the field is accessible.
func NewField(pos syntax.Pos, pkg *Package, name string, typ Type, embedded bool) *Var

// Anonymous reports whether the variable is an embedded field.
// Same as Embedded; only present for backward-compatibility.
func (obj *Var) Anonymous() bool

// Embedded reports whether the variable is an embedded field.
func (obj *Var) Embedded() bool

// IsField reports whether the variable is a struct field.
func (obj *Var) IsField() bool

// Origin returns the canonical Var for its receiver, i.e. the Var object
// recorded in Info.Defs.
//
// For synthetic Vars created during instantiation (such as struct fields or
// function parameters that depend on type arguments), this will be the
// corresponding Var on the generic (uninstantiated) type. For all other Vars
// Origin returns the receiver.
func (obj *Var) Origin() *Var

// A Func represents a declared function, concrete method, or abstract
// (interface) method. Its Type() is always a *Signature.
// An abstract method may belong to many interfaces due to embedding.
type Func struct {
	object
	hasPtrRecv_ bool
	origin      *Func
}

// NewFunc returns a new function with the given signature, representing
// the function's type.
func NewFunc(pos syntax.Pos, pkg *Package, name string, sig *Signature) *Func

// Signature returns the signature (type) of the function or method.
func (obj *Func) Signature() *Signature

// FullName returns the package- or receiver-type-qualified name of
// function or method obj.
func (obj *Func) FullName() string

// Scope returns the scope of the function's body block.
// The result is nil for imported or instantiated functions and methods
// (but there is also no mechanism to get to an instantiated function).
func (obj *Func) Scope() *Scope

// Origin returns the canonical Func for its receiver, i.e. the Func object
// recorded in Info.Defs.
//
// For synthetic functions created during instantiation (such as methods on an
// instantiated Named type or interface methods that depend on type arguments),
// this will be the corresponding Func on the generic (uninstantiated) type.
// For all other Funcs Origin returns the receiver.
func (obj *Func) Origin() *Func

// Pkg returns the package to which the function belongs.
//
// The result is nil for methods of types in the Universe scope,
// like method Error of the error built-in interface type.
func (obj *Func) Pkg() *Package

// A Label represents a declared label.
// Labels don't have a type.
type Label struct {
	object
	used bool
}

// NewLabel returns a new label.
func NewLabel(pos syntax.Pos, pkg *Package, name string) *Label

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
