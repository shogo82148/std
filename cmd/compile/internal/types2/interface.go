// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
)

// An Interface represents an interface type.
type Interface struct {
	check     *Checker
	methods   []*Func
	embeddeds []Type
	embedPos  *[]syntax.Pos
	implicit  bool
	complete  bool

	tset *_TypeSet
}

// NewInterfaceType returns a new interface for the given methods and embedded types.
// NewInterfaceType takes ownership of the provided methods and may modify their types
// by setting missing receivers.
func NewInterfaceType(methods []*Func, embeddeds []Type) *Interface

// MarkImplicit marks the interface t as implicit, meaning this interface
// corresponds to a constraint literal such as ~T or A|B without explicit
// interface embedding. MarkImplicit should be called before any concurrent use
// of implicit interfaces.
func (t *Interface) MarkImplicit()

// NumExplicitMethods returns the number of explicitly declared methods of interface t.
func (t *Interface) NumExplicitMethods() int

// ExplicitMethod returns the i'th explicitly declared method of interface t for 0 <= i < t.NumExplicitMethods().
// The methods are ordered by their unique Id.
func (t *Interface) ExplicitMethod(i int) *Func

// NumEmbeddeds returns the number of embedded types in interface t.
func (t *Interface) NumEmbeddeds() int

// EmbeddedType returns the i'th embedded type of interface t for 0 <= i < t.NumEmbeddeds().
func (t *Interface) EmbeddedType(i int) Type

// NumMethods returns the total number of methods of interface t.
func (t *Interface) NumMethods() int

// Method returns the i'th method of interface t for 0 <= i < t.NumMethods().
// The methods are ordered by their unique Id.
func (t *Interface) Method(i int) *Func

// Empty reports whether t is the empty interface.
func (t *Interface) Empty() bool

// IsComparable reports whether each type in interface t's type set is comparable.
func (t *Interface) IsComparable() bool

// IsMethodSet reports whether the interface t is fully described by its method set.
func (t *Interface) IsMethodSet() bool

// IsImplicit reports whether the interface t is a wrapper for a type set literal.
func (t *Interface) IsImplicit() bool

func (t *Interface) Underlying() Type
func (t *Interface) String() string
