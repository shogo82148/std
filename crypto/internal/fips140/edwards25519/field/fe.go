// Copyright (c) 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package field implements fast arithmetic modulo 2^255-19.
package field

// Element represents an element of the field GF(2^255-19). Note that this
// is not a cryptographically secure group, and should only be used to interact
// with edwards25519.Point coordinates.
//
// This type works similarly to math/big.Int, and all arguments and receivers
// are allowed to alias.
//
// The zero value is a valid zero element.
type Element struct {
	// An element t represents the integer
	//     t.l0 + t.l1*2^51 + t.l2*2^102 + t.l3*2^153 + t.l4*2^204
	//
	// Between operations, all limbs are expected to be lower than 2^52.
	l0 uint64
	l1 uint64
	l2 uint64
	l3 uint64
	l4 uint64
}

// Zero sets v = 0, and returns v.
func (v *Element) Zero() *Element

// One sets v = 1, and returns v.
func (v *Element) One() *Element

// Add sets v = a + b, and returns v.
func (v *Element) Add(a, b *Element) *Element

// Subtract sets v = a - b, and returns v.
func (v *Element) Subtract(a, b *Element) *Element

// Negate sets v = -a, and returns v.
func (v *Element) Negate(a *Element) *Element

// Invert sets v = 1/z mod p, and returns v.
//
// If z == 0, Invert returns v = 0.
func (v *Element) Invert(z *Element) *Element

// Set sets v = a, and returns v.
func (v *Element) Set(a *Element) *Element

// SetBytes sets v to x, where x is a 32-byte little-endian encoding. If x is
// not of the right length, SetBytes returns nil and an error, and the
// receiver is unchanged.
//
// Consistent with RFC 7748, the most significant bit (the high bit of the
// last byte) is ignored, and non-canonical values (2^255-19 through 2^255-1)
// are accepted. Note that this is laxer than specified by RFC 8032, but
// consistent with most Ed25519 implementations.
func (v *Element) SetBytes(x []byte) (*Element, error)

// Bytes returns the canonical 32-byte little-endian encoding of v.
func (v *Element) Bytes() []byte

// Equal returns 1 if v and u are equal, and 0 otherwise.
func (v *Element) Equal(u *Element) int

// Select sets v to a if cond == 1, and to b if cond == 0.
func (v *Element) Select(a, b *Element, cond int) *Element

// Swap swaps v and u if cond == 1 or leaves them unchanged if cond == 0, and returns v.
func (v *Element) Swap(u *Element, cond int)

// IsNegative returns 1 if v is negative, and 0 otherwise.
func (v *Element) IsNegative() int

// Absolute sets v to |u|, and returns v.
func (v *Element) Absolute(u *Element) *Element

// Multiply sets v = x * y, and returns v.
func (v *Element) Multiply(x, y *Element) *Element

// Square sets v = x * x, and returns v.
func (v *Element) Square(x *Element) *Element

// Mult32 sets v = x * y, and returns v.
func (v *Element) Mult32(x *Element, y uint32) *Element

// Pow22523 set v = x^((p-5)/8), and returns v. (p-5)/8 is 2^252-3.
func (v *Element) Pow22523(x *Element) *Element

// SqrtRatio sets r to the non-negative square root of the ratio of u and v.
//
// If u/v is square, SqrtRatio returns r and 1. If u/v is not square, SqrtRatio
// sets r according to Section 4.3 of draft-irtf-cfrg-ristretto255-decaf448-00,
// and returns r and 0.
func (r *Element) SqrtRatio(u, v *Element) (R *Element, wasSquare int)
