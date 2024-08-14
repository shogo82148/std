// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import "github.com/shogo82148/std/iter"

// Methods returns a go1.23 iterator over all the methods of an
// interface, ordered by Id.
//
// Example: for m := range t.Methods() { ... }
func (t *Interface) Methods() iter.Seq[*Func]

// ExplicitMethods returns a go1.23 iterator over the explicit methods of
// an interface, ordered by Id.
//
// Example: for m := range t.ExplicitMethods() { ... }
func (t *Interface) ExplicitMethods() iter.Seq[*Func]

// EmbeddedTypes returns a go1.23 iterator over the types embedded within an interface.
//
// Example: for e := range t.EmbeddedTypes() { ... }
func (t *Interface) EmbeddedTypes() iter.Seq[Type]

// Methods returns a go1.23 iterator over the declared methods of a named type.
//
// Example: for m := range t.Methods() { ... }
func (t *Named) Methods() iter.Seq[*Func]

// Children returns a go1.23 iterator over the child scopes nested within scope s.
//
// Example: for child := range scope.Children() { ... }
func (s *Scope) Children() iter.Seq[*Scope]

// Fields returns a go1.23 iterator over the fields of a struct type.
//
// Example: for field := range s.Fields() { ... }
func (s *Struct) Fields() iter.Seq[*Var]

// Variables returns a go1.23 iterator over the variables of a tuple type.
//
// Example: for v := range tuple.Variables() { ... }
func (t *Tuple) Variables() iter.Seq[*Var]

// Methods returns a go1.23 iterator over the methods of a method set.
//
// Example: for method := range s.Methods() { ... }
func (s *MethodSet) Methods() iter.Seq[*Selection]

// Terms returns a go1.23 iterator over the terms of a union.
//
// Example: for term := range union.Terms() { ... }
func (u *Union) Terms() iter.Seq[*Term]

// TypeParams returns a go1.23 iterator over a list of type parameters.
//
// Example: for tparam := range l.TypeParams() { ... }
func (l *TypeParamList) TypeParams() iter.Seq[*TypeParam]

// Types returns a go1.23 iterator over the elements of a list of types.
//
// Example: for t := range l.Types() { ... }
func (l *TypeList) Types() iter.Seq[Type]
