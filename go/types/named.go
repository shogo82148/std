// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/sync"
)

// A Named represents a named (defined) type.
type Named struct {
	check      *Checker
	obj        *TypeName
	orig       *Named
	fromRHS    Type
	underlying Type
	tparams    *TypeParamList
	targs      *TypeList

	methods *methodList

	resolver func(*Context, *Named) (tparams *TypeParamList, underlying Type, methods *methodList)
	once     sync.Once
}

// NewNamed returns a new named type for the given type name, underlying type, and associated methods.
// If the given type name obj doesn't have a type yet, its type is set to the returned named type.
// The underlying type must not be a *Named.
func NewNamed(obj *TypeName, underlying Type, methods []*Func) *Named

// Obj returns the type name for the declaration defining the named type t. For
// instantiated types, this is same as the type name of the origin type.
func (t *Named) Obj() *TypeName

// Origin returns the generic type from which the named type t is
// instantiated. If t is not an instantiated type, the result is t.
func (t *Named) Origin() *Named

// TypeParams returns the type parameters of the named type t, or nil.
// The result is non-nil for an (originally) generic type even if it is instantiated.
func (t *Named) TypeParams() *TypeParamList

// SetTypeParams sets the type parameters of the named type t.
// t must not have type arguments.
func (t *Named) SetTypeParams(tparams []*TypeParam)

// TypeArgs returns the type arguments used to instantiate the named type t.
func (t *Named) TypeArgs() *TypeList

// NumMethods returns the number of explicit methods defined for t.
//
// For an ordinary or instantiated type t, the receiver base type of these
// methods will be the named type t. For an uninstantiated generic type t, each
// method receiver will be instantiated with its receiver type parameters.
func (t *Named) NumMethods() int

// Method returns the i'th method of named type t for 0 <= i < t.NumMethods().
func (t *Named) Method(i int) *Func

// SetUnderlying sets the underlying type and marks t as complete.
// t must not have type arguments.
func (t *Named) SetUnderlying(underlying Type)

// AddMethod adds method m unless it is already in the method list.
// t must not have type arguments.
func (t *Named) AddMethod(m *Func)

func (t *Named) Underlying() Type
func (t *Named) String() string
