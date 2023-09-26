// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Note: This is a uint32 rather than a uint64 because the
// respective 64 bit atomic instructions are not available
// on all platforms.

// A TypeParam represents a type parameter type.
type TypeParam struct {
	check *Checker
	id    uint64
	obj   *TypeName
	index int
	bound Type
}

// NewTypeParam returns a new TypeParam. Type parameters may be set on a Named
// or Signature type by calling SetTypeParams. Setting a type parameter on more
// than one type will result in a panic.
//
// The constraint argument can be nil, and set later via SetConstraint. If the
// constraint is non-nil, it must be fully defined.
func NewTypeParam(obj *TypeName, constraint Type) *TypeParam

// Index returns the index of the type param within its param list, or -1 if
// the type parameter has not yet been bound to a type.
func (t *TypeParam) Index() int

// Obj returns the type name for t.
func (t *TypeParam) Obj() *TypeName

// Constraint returns the type constraint specified for t.
func (t *TypeParam) Constraint() Type

// SetConstraint sets the type constraint for t.
//
// It must be called by users of NewTypeParam after the bound's underlying is
// fully defined, and before using the type parameter in any way other than to
// form other types. Once SetConstraint returns the receiver, t is safe for
// concurrent use.
func (t *TypeParam) SetConstraint(bound Type)

func (t *TypeParam) Underlying() Type

func (t *TypeParam) String() string
