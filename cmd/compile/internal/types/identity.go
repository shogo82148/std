// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Identical reports whether t1 and t2 are identical types, following the spec rules.
// Receiver parameter types are ignored. Named (defined) types are only equal if they
// are pointer-equal - i.e. there must be a unique types.Type for each specific named
// type. Also, a type containing a shape type is considered identical to another type
// (shape or not) if their underlying types are the same, or they are both pointers.
func Identical(t1, t2 *Type) bool

// IdenticalIgnoreTags is like Identical, but it ignores struct tags
// for struct identity.
func IdenticalIgnoreTags(t1, t2 *Type) bool

// IdenticalStrict is like Identical, but matches types exactly, without the
// exception for shapes.
func IdenticalStrict(t1, t2 *Type) bool
