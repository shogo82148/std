// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//simdgen:category Logical

package spec

// And returns the bitwise AND of x and y.
//
//	z[i] = x[i] & y[i]
//
//specgen:commutative
func And[E Ints | Uints | MaskElt, W Width](x, y Vec[E, W]) (z Vec[E, W])

// AndNot returns the bitwise AND NOT of x and y.
//
//	z[i] = x[i] &^ y[i]
func AndNot[E Ints | Uints | MaskElt, W Width](x, y Vec[E, W]) (z Vec[E, W])

// Or returns the bitwise OR of x and y.
//
//	z[i] = x[i] | y[i]
//
//specgen:commutative
func Or[E Ints | Uints | MaskElt, W Width](x, y Vec[E, W]) (z Vec[E, W])

// Xor returns the bitwise XOR of x and y.
//
//	z[i] = x[i] ^ y[i]
//
//specgen:commutative
func Xor[E Ints | Uints | MaskElt, W Width](x, y Vec[E, W]) (z Vec[E, W])

// Not returns the bitwise negation of x.
//
//	z[i] = ^x[i]
func Not[E Ints | Uints | MaskElt, W Width](x Vec[E, W]) (z Vec[E, W])
