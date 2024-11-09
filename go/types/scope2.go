// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements go/types-specific scope methods.
// These methods do not exist in types2.

package types

import "github.com/shogo82148/std/go/token"

// LookupParent follows the parent chain of scopes starting with s until
// it finds a scope where Lookup(name) returns a non-nil object, and then
// returns that scope and object. If a valid position pos is provided,
// only objects that were declared at or before pos are considered.
// If no such scope and object exists, the result is (nil, nil).
// The results are guaranteed to be valid only if the type-checked
// AST has complete position information.
//
// Note that obj.Parent() may be different from the returned scope if the
// object was inserted into the scope and already had a parent at that
// time (see Insert). This can only happen for dot-imported objects
// whose parent is the scope of the package that exported them.
func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)

// Pos and End describe the scope's source code extent [pos, end).
// The results are guaranteed to be valid only if the type-checked
// AST has complete position information. The extent is undefined
// for Universe and package scopes.
func (s *Scope) Pos() token.Pos
func (s *Scope) End() token.Pos

// Contains reports whether pos is within the scope's extent.
// The result is guaranteed to be valid only if the type-checked
// AST has complete position information.
func (s *Scope) Contains(pos token.Pos) bool

// Innermost returns the innermost (child) scope containing
// pos. If pos is not within any scope, the result is nil.
// The result is also nil for the Universe scope.
// The result is guaranteed to be valid only if the type-checked
// AST has complete position information.
func (s *Scope) Innermost(pos token.Pos) *Scope