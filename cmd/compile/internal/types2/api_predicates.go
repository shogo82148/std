// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements exported type predicates.

package types2

// AssertableTo reports whether a value of type V can be asserted to have type T.
//
// The behavior of AssertableTo is unspecified in three cases:
//   - if T is Typ[Invalid]
//   - if V is a generalized interface; i.e., an interface that may only be used
//     as a type constraint in Go code
//   - if T is an uninstantiated generic type
func AssertableTo(V *Interface, T Type) bool

// AssignableTo reports whether a value of type V is assignable to a variable
// of type T.
//
// The behavior of AssignableTo is unspecified if V or T is Typ[Invalid] or an
// uninstantiated generic type.
func AssignableTo(V, T Type) bool

// ConvertibleTo reports whether a value of type V is convertible to a value of
// type T.
//
// The behavior of ConvertibleTo is unspecified if V or T is Typ[Invalid] or an
// uninstantiated generic type.
func ConvertibleTo(V, T Type) bool

// Implements reports whether type V implements interface T.
//
// The behavior of Implements is unspecified if V is Typ[Invalid] or an uninstantiated
// generic type.
func Implements(V Type, T *Interface) bool

// Satisfies reports whether type V satisfies the constraint T.
//
// The behavior of Satisfies is unspecified if V is Typ[Invalid] or an uninstantiated
// generic type.
func Satisfies(V Type, T *Interface) bool

// Identical reports whether x and y are identical types.
// Receivers of [Signature] types are ignored.
//
// Predicates such as [Identical], [Implements], and
// [Satisfies] assume that both operands belong to a
// consistent collection of symbols ([Object] values).
// For example, two [Named] types can be identical only if their
// [Named.Obj] methods return the same [TypeName] symbol.
// A collection of symbols is consistent if, for each logical
// package whose path is P, the creation of those symbols
// involved at most one call to [NewPackage](P, ...).
// To ensure consistency, use a single [Importer] for
// all loaded packages and their dependencies.
// For more information, see https://github.com/golang/go/issues/57497.
func Identical(x, y Type) bool

// IdenticalIgnoreTags reports whether x and y are identical types if tags are ignored.
// Receivers of [Signature] types are ignored.
func IdenticalIgnoreTags(x, y Type) bool
