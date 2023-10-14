// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reflectlite implements lightweight version of reflect, not using
// any package except for "runtime", "unsafe", and "internal/abi"
package reflectlite

import (
	"github.com/shogo82148/std/internal/abi"
)

// Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types. Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of type before
// calling kind-specific methods. Calling a method
// inappropriate to the kind of type causes a run-time panic.
//
// Type values are comparable, such as with the == operator,
// so they can be used as map keys.
// Two Type values are equal if they represent identical types.
type Type interface {
	Name() string

	PkgPath() string

	Size() uintptr

	Kind() Kind

	Implements(u Type) bool

	AssignableTo(u Type) bool

	Comparable() bool

	String() string

	Elem() Type

	common() *abi.Type
	uncommon() *uncommonType
}

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind = abi.Kind

const Ptr = abi.Pointer

const (
	// Import-and-export these constants as necessary
	Interface = abi.Interface
	Slice     = abi.Slice
	String    = abi.String
	Struct    = abi.Struct
)

// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i any) Type
