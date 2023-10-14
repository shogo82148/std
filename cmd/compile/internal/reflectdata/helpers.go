// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectdata

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/internal/src"
)

// AppendElemRType asserts that n is an "append" operation, and
// returns an expression that yields the *runtime._type value
// representing the result slice type's element type.
func AppendElemRType(pos src.XPos, n *ir.CallExpr) ir.Node

// CompareRType asserts that n is a comparison (== or !=) operation
// between expressions of interface and non-interface type, and
// returns an expression that yields the *runtime._type value
// representing the non-interface type.
func CompareRType(pos src.XPos, n *ir.BinaryExpr) ir.Node

// ConvIfaceTypeWord asserts that n is conversion to interface type,
// and returns an expression that yields the *runtime._type or
// *runtime.itab value necessary for implementing the conversion.
//
//   - *runtime._type for the destination type, for I2I conversions
//   - *runtime.itab, for T2I conversions
//   - *runtime._type for the source type, for T2E conversions
func ConvIfaceTypeWord(pos src.XPos, n *ir.ConvExpr) ir.Node

// ConvIfaceSrcRType asserts that n is a conversion from
// non-interface type to interface type, and
// returns an expression that yields the *runtime._type for copying
// the convertee value to the heap.
func ConvIfaceSrcRType(pos src.XPos, n *ir.ConvExpr) ir.Node

// CopyElemRType asserts that n is a "copy" operation, and returns an
// expression that yields the *runtime._type value representing the
// destination slice type's element type.
func CopyElemRType(pos src.XPos, n *ir.BinaryExpr) ir.Node

// DeleteMapRType asserts that n is a "delete" operation, and returns
// an expression that yields the *runtime._type value representing the
// map type.
func DeleteMapRType(pos src.XPos, n *ir.CallExpr) ir.Node

// IndexMapRType asserts that n is a map index operation, and returns
// an expression that yields the *runtime._type value representing the
// map type.
func IndexMapRType(pos src.XPos, n *ir.IndexExpr) ir.Node

// MakeChanRType asserts that n is a "make" operation for a channel
// type, and returns an expression that yields the *runtime._type
// value representing that channel type.
func MakeChanRType(pos src.XPos, n *ir.MakeExpr) ir.Node

// MakeMapRType asserts that n is a "make" operation for a map type,
// and returns an expression that yields the *runtime._type value
// representing that map type.
func MakeMapRType(pos src.XPos, n *ir.MakeExpr) ir.Node

// MakeSliceElemRType asserts that n is a "make" operation for a slice
// type, and returns an expression that yields the *runtime._type
// value representing that slice type's element type.
func MakeSliceElemRType(pos src.XPos, n *ir.MakeExpr) ir.Node

// RangeMapRType asserts that n is a "range" loop over a map value,
// and returns an expression that yields the *runtime._type value
// representing that map type.
func RangeMapRType(pos src.XPos, n *ir.RangeStmt) ir.Node

// UnsafeSliceElemRType asserts that n is an "unsafe.Slice" operation,
// and returns an expression that yields the *runtime._type value
// representing the result slice type's element type.
func UnsafeSliceElemRType(pos src.XPos, n *ir.BinaryExpr) ir.Node
