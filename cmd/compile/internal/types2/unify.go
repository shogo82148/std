// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements type unification.
//
// Type unification attempts to make two types x and y structurally
// equivalent by determining the types for a given list of (bound)
// type parameters which may occur within x and y. If x and y are
// structurally different (say []T vs chan T), or conflicting
// types are determined for type parameters, unification fails.
// If unification succeeds, as a side-effect, the types of the
// bound type parameters may be determined.
//
// Unification typically requires multiple calls u.unify(x, y) to
// a given unifier u, with various combinations of types x and y.
// In each call, additional type parameter types may be determined
// as a side effect and recorded in u.
// If a call fails (returns false), unification fails.
//
// In the unification context, structural equivalence of two types
// ignores the difference between a defined type and its underlying
// type if one type is a defined type and the other one is not.
// It also ignores the difference between an (external, unbound)
// type parameter and its core type.
// If two types are not structurally equivalent, they cannot be Go
// identical types. On the other hand, if they are structurally
// equivalent, they may be Go identical or at least assignable, or
// they may be in the type set of a constraint.
// Whether they indeed are identical or assignable is determined
// upon instantiation and function argument passing.

package types2
