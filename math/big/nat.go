// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package big implements multi-precision arithmetic (big numbers).
// The following numeric types are supported:
//
//	Int    signed integers
//	Rat    rational numbers
//	Float  floating-point numbers
//
// Methods are typically of the form:
//
//	func (z *T) Unary(x *T) *T        // z = op x
//	func (z *T) Binary(x, y *T) *T    // z = x op y
//	func (x *T) M() T1                // v = x.M()
//
// with T one of Int, Rat, or Float. For unary and binary operations, the
// result is the receiver (usually named z in that case); if it is one of
// the operands x or y it may be overwritten (and its memory reused).
// To enable chaining of operations, the result is also returned. Methods
// returning a result other than *Int, *Rat, or *Float take an operand as
// the receiver (usually named x in that case).
package big

// An unsigned integer x of the form
//
//   x = x[n-1]*_B^(n-1) + x[n-2]*_B^(n-2) + ... + x[1]*_B + x[0]
//
// with 0 <= x[i] < _B and 0 <= i < n is stored in a slice of length n,
// with the digits x[i] as the slice elements.
//
// A number is normalized if the slice contains no leading 0 digits.
// During arithmetic operations, denormalized values may occur but are
// always normalized before returning the final result. The normalized
// representation of 0 is the empty or nil slice (length = 0).
//

// Operands that are shorter than karatsubaThreshold are multiplied using
// "grade school" multiplication; for longer operands the Karatsuba algorithm
// is used.
