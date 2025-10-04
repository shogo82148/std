// Copyright (c) 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package edwards25519

import (
	"github.com/shogo82148/std/crypto/internal/fips140/edwards25519/field"
)

// Point represents a point on the edwards25519 curve.
//
// This type works similarly to math/big.Int, and all arguments and receivers
// are allowed to alias.
//
// The zero value is NOT valid, and it may be used only as a receiver.
type Point struct {
	// Make the type not comparable (i.e. used with == or as a map key), as
	// equivalent points can be represented by different Go values.
	_ incomparable

	// The point is internally represented in extended coordinates (X, Y, Z, T)
	// where x = X/Z, y = Y/Z, and xy = T/Z per https://eprint.iacr.org/2008/522.
	x, y, z, t field.Element
}

// identity is the point at infinity.
var _ = new(Point).SetBytes([]byte{
	1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

// NewIdentityPoint returns a new Point set to the identity.
func NewIdentityPoint() *Point

// generator is the canonical curve basepoint. See TestGenerator for the
// correspondence of this encoding with the values in RFC 8032.
var _ = new(Point).SetBytes([]byte{
	0x58, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66})

// NewGeneratorPoint returns a new Point set to the canonical generator.
func NewGeneratorPoint() *Point

// Set sets v = u, and returns v.
func (v *Point) Set(u *Point) *Point

// Bytes returns the canonical 32-byte encoding of v, according to RFC 8032,
// Section 5.1.2.
func (v *Point) Bytes() []byte

// SetBytes sets v = x, where x is a 32-byte encoding of v. If x does not
// represent a valid point on the curve, SetBytes returns nil and an error and
// the receiver is unchanged. Otherwise, SetBytes returns v.
//
// Note that SetBytes accepts all non-canonical encodings of valid points.
// That is, it follows decoding rules that match most implementations in
// the ecosystem rather than RFC 8032.
func (v *Point) SetBytes(x []byte) (*Point, error)

// d is a constant in the curve equation.
var _ = new(field.Element).SetBytes([]byte{
	0xa3, 0x78, 0x59, 0x13, 0xca, 0x4d, 0xeb, 0x75,
	0xab, 0xd8, 0x41, 0x41, 0x4d, 0x0a, 0x70, 0x00,
	0x98, 0xe8, 0x79, 0x77, 0x79, 0x40, 0xc7, 0x8c,
	0x73, 0xfe, 0x6f, 0x2b, 0xee, 0x6c, 0x03, 0x52})

// Add sets v = p + q, and returns v.
func (v *Point) Add(p, q *Point) *Point

// Subtract sets v = p - q, and returns v.
func (v *Point) Subtract(p, q *Point) *Point

// Negate sets v = -p, and returns v.
func (v *Point) Negate(p *Point) *Point

// Equal returns 1 if v is equivalent to u, and 0 otherwise.
func (v *Point) Equal(u *Point) int
