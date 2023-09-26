// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package subtle implements functions that are often useful in cryptographic
// code but require careful thought to use correctly.
package subtle

// ConstantTimeCompare returns 1 iff the two equal length slices, x
// and y, have equal contents. The time taken is a function of the length of
// the slices and is independent of the contents.
func ConstantTimeCompare(x, y []byte) int

// ConstantTimeSelect returns x if v is 1 and y if v is 0.
// Its behavior is undefined if v takes any other value.
func ConstantTimeSelect(v, x, y int) int

// ConstantTimeByteEq returns 1 if x == y and 0 otherwise.
func ConstantTimeByteEq(x, y uint8) int

// ConstantTimeEq returns 1 if x == y and 0 otherwise.
func ConstantTimeEq(x, y int32) int

// ConstantTimeCopy copies the contents of y into x iff v == 1. If v == 0, x is left unchanged.
// Its behavior is undefined if v takes any other value.
func ConstantTimeCopy(v int, x, y []byte)
