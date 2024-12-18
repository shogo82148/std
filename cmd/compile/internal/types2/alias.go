// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// An Alias represents an alias type.
//
// Alias types are created by alias declarations such as:
//
//	type A = int
//
// The type on the right-hand side of the declaration can be accessed
// using [Alias.Rhs]. This type may itself be an alias.
// Call [Unalias] to obtain the first non-alias type in a chain of
// alias type declarations.
//
// Like a defined ([Named]) type, an alias type has a name.
// Use the [Alias.Obj] method to access its [TypeName] object.
//
// Historically, Alias types were not materialized so that, in the example
// above, A's type was represented by a Basic (int), not an Alias
// whose [Alias.Rhs] is int. But Go 1.24 allows you to declare an
// alias type with type parameters or arguments:
//
//	type Set[K comparable] = map[K]bool
//	s := make(Set[String])
//
// and this requires that Alias types be materialized. Use the
// [Alias.TypeParams] and [Alias.TypeArgs] methods to access them.
//
// To ease the transition, the Alias type was introduced in go1.22,
// but the type-checker would not construct values of this type unless
// the GODEBUG=gotypesalias=1 environment variable was provided.
// Starting in go1.23, this variable is enabled by default.
// This setting also causes the predeclared type "any" to be
// represented as an Alias, not a bare [Interface].
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
