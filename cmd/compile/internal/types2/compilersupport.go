// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Helper functions exported for the compiler.
// Do not use internally.

package types2

// If t is a pointer, AsPointer returns that type, otherwise it returns nil.
func AsPointer(t Type) *Pointer

// If t is a signature, AsSignature returns that type, otherwise it returns nil.
func AsSignature(t Type) *Signature

// If typ is a type parameter, CoreType returns the single underlying
// type of all types in the corresponding type constraint if it exists, or
// nil otherwise. If the type set contains only unrestricted and restricted
// channel types (with identical element types), the single underlying type
// is the restricted channel type if the restrictions are always the same.
// If typ is not a type parameter, CoreType returns the underlying type.
func CoreType(t Type) Type

// RangeKeyVal returns the key and value types for a range over typ.
// It panics if range over typ is invalid.
func RangeKeyVal(typ Type) (Type, Type)
