// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// An Alias represents an alias type.
// Whether or not Alias types are created is controlled by the
// gotypesalias setting with the GODEBUG environment variable.
// For gotypesalias=1, alias declarations produce an Alias type.
// Otherwise, the alias information is only in the type name,
// which points directly to the actual (aliased) type.
type Alias struct {
	obj     *TypeName
	orig    *Alias
	tparams *TypeParamList
	targs   *TypeList
	fromRHS Type
	actual  Type
}

// NewAlias creates a new Alias type with the given type name and rhs.
// rhs must not be nil.
func NewAlias(obj *TypeName, rhs Type) *Alias

// Obj returns the type name for the declaration defining the alias type a.
// For instantiated types, this is same as the type name of the origin type.
func (a *Alias) Obj() *TypeName

func (a *Alias) String() string

// Underlying returns the [underlying type] of the alias type a, which is the
// underlying type of the aliased type. Underlying types are never Named,
// TypeParam, or Alias types.
//
// [underlying type]: https://go.dev/ref/spec#Underlying_types.
func (a *Alias) Underlying() Type

// Origin returns the generic Alias type of which a is an instance.
// If a is not an instance of a generic alias, Origin returns a.
func (a *Alias) Origin() *Alias

// TypeParams returns the type parameters of the alias type a, or nil.
// A generic Alias and its instances have the same type parameters.
func (a *Alias) TypeParams() *TypeParamList

// SetTypeParams sets the type parameters of the alias type a.
// The alias a must not have type arguments.
func (a *Alias) SetTypeParams(tparams []*TypeParam)

// TypeArgs returns the type arguments used to instantiate the Alias type.
// If a is not an instance of a generic alias, the result is nil.
func (a *Alias) TypeArgs() *TypeList

// Rhs returns the type R on the right-hand side of an alias
// declaration "type A = R", which may be another alias.
func (a *Alias) Rhs() Type

// Unalias returns t if it is not an alias type;
// otherwise it follows t's alias chain until it
// reaches a non-alias type which is then returned.
// Consequently, the result is never an alias type.
func Unalias(t Type) Type
