// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements Scopes.

package types2

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/io"
)

// A Scope maintains a set of objects and links to its containing
// (parent) and contained (children) scopes. Objects may be inserted
// and looked up by name. The zero value for Scope is a ready-to-use
// empty scope.
type Scope struct {
	parent   *Scope
	children []*Scope
	number   int
	elems    map[string]Object
	pos, end syntax.Pos
	comment  string
	isFunc   bool
}

// NewScope returns a new, empty scope contained in the given parent
// scope, if any. The comment is for debugging only.
func NewScope(parent *Scope, pos, end syntax.Pos, comment string) *Scope

// Parent returns the scope's containing (parent) scope.
func (s *Scope) Parent() *Scope

// Len returns the number of scope elements.
func (s *Scope) Len() int

// Names returns the scope's element names in sorted order.
func (s *Scope) Names() []string

// NumChildren returns the number of scopes nested in s.
func (s *Scope) NumChildren() int

// Child returns the i'th child scope for 0 <= i < NumChildren().
func (s *Scope) Child(i int) *Scope

// Lookup returns the object in scope s with the given name if such an
// object exists; otherwise the result is nil.
func (s *Scope) Lookup(name string) Object

// Insert attempts to insert an object obj into scope s.
// If s already contains an alternative object alt with
// the same name, Insert leaves s unchanged and returns alt.
// Otherwise it inserts obj, sets the object's parent scope
// if not already set, and returns nil.
func (s *Scope) Insert(obj Object) Object

// InsertLazy is like Insert, but allows deferring construction of the
// inserted object until it's accessed with Lookup. The Object
// returned by resolve must have the same name as given to InsertLazy.
// If s already contains an alternative object with the same name,
// InsertLazy leaves s unchanged and returns false. Otherwise it
// records the binding and returns true. The object's parent scope
// will be set to s after resolve is called.
func (s *Scope) InsertLazy(name string, resolve func() Object) bool

// WriteTo writes a string representation of the scope to w,
// with the scope elements sorted by name.
// The level of indentation is controlled by n >= 0, with
// n == 0 for no indentation.
// If recurse is set, it also writes nested (children) scopes.
func (s *Scope) WriteTo(w io.Writer, n int, recurse bool)

// String returns a string representation of the scope, for debugging.
func (s *Scope) String() string
