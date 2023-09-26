// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package big implements multi-precision arithmetic (big numbers).
// The following numeric types are supported:
//
//   - Int	signed integers
//   - Rat	rational numbers
//
// Methods are typically of the form:
//
//	func (z *Int) Op(x, y *Int) *Int	(similar for *Rat)
//
// and implement operations z = x Op y with the result as receiver; if it
// is one of the operands it may be overwritten (and its memory reused).
// To enable chaining of operations, the result is also returned. Methods
// returning a result other than *Int or *Rat take one of the operands as
// the receiver.
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

// MaxBase is the largest number base accepted for string conversions.
const MaxBase = 'z' - 'a' + 10 + 1

// Character sets for string conversion.

// Split blocks greater than leafSize Words (or set to 0 to disable recursive conversion)
// Benchmark and configure leafSize using: go test -bench="Leaf"
//   8 and 16 effective on 3.0 GHz Xeon "Clovertown" CPU (128 byte cache lines)
//   8 and 16 effective on 2.66 GHz Core 2 Duo "Penryn" CPU
