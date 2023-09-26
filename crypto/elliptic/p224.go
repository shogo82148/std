// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elliptic

// P224 returns a Curve which implements P-224 (see FIPS 186-3, section D.2.2)
func P224() Curve

// Field element functions.
//
// The field that we're dealing with is ℤ/pℤ where p = 2**224 - 2**96 + 1.
//
// Field elements are represented by a FieldElement, which is a typedef to an
// array of 8 uint32's. The value of a FieldElement, a, is:
//   a[0] + 2**28·a[1] + 2**56·a[1] + ... + 2**196·a[7]
//
// Using 28-bit limbs means that there's only 4 bits of headroom, which is less
// than we would really like. But it has the useful feature that we hit 2**224
// exactly, making the reflections during a reduce much nicer.

// p224P is the order of the field, represented as a p224FieldElement.

// p224ZeroModP31 is 0 mod p where bit 31 is set in all limbs so that we can
// subtract smaller amounts without underflow. See the section "Subtraction" in
// [1] for reasoning.

// LargeFieldElement also represents an element of the field. The limbs are
// still spaced 28-bits apart and in little-endian order. So the limbs are at
// 0, 28, 56, ..., 392 bits, each 64-bits wide.

// p224ZeroModP63 is 0 mod p where bit 63 is set in all limbs. See the section
// "Subtraction" in [1] for why.
