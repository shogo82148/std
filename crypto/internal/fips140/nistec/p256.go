// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (!amd64 && !arm64 && !ppc64le && !s390x) || purego

package nistec

import (
	"github.com/shogo82148/std/crypto/internal/fips140/nistec/fiat"
)

// P256Point is a P-256 point. The zero value is NOT valid.
type P256Point struct {
	// The point is represented in projective coordinates (X:Y:Z), where x = X/Z
	// and y = Y/Z. Infinity is (0:1:0).
	//
	// fiat.P256Element is a base field element in [0, P-1] in the Montgomery
	// domain (with R 2²⁵⁶ and P 2²⁵⁶ - 2²²⁴ + 2¹⁹² + 2⁹⁶ - 1) as four limbs in
	// little-endian order value.
	x, y, z fiat.P256Element
}

// NewP256Point returns a new P256Point representing the point at infinity point.
func NewP256Point() *P256Point

// SetGenerator sets p to the canonical generator and returns p.
func (p *P256Point) SetGenerator() *P256Point

// Set sets p = q and returns p.
func (p *P256Point) Set(q *P256Point) *P256Point

// SetBytes sets p to the compressed, uncompressed, or infinity value encoded in
// b, as specified in SEC 1, Version 2.0, Section 2.3.4. If the point is not on
// the curve, it returns nil and an error, and the receiver is unchanged.
// Otherwise, it returns p.
func (p *P256Point) SetBytes(b []byte) (*P256Point, error)

// Bytes returns the uncompressed or infinity encoding of p, as specified in
// SEC 1, Version 2.0, Section 2.3.3. Note that the encoding of the point at
// infinity is shorter than all other encodings.
func (p *P256Point) Bytes() []byte

// BytesX returns the encoding of the x-coordinate of p, as specified in SEC 1,
// Version 2.0, Section 2.3.5, or an error if p is the point at infinity.
func (p *P256Point) BytesX() ([]byte, error)

// BytesCompressed returns the compressed or infinity encoding of p, as
// specified in SEC 1, Version 2.0, Section 2.3.3. Note that the encoding of the
// point at infinity is shorter than all other encodings.
func (p *P256Point) BytesCompressed() []byte

// Add sets q = p1 + p2, and returns q. The points may overlap.
func (q *P256Point) Add(p1, p2 *P256Point) *P256Point

// Double sets q = p + p, and returns q. The points may overlap.
func (q *P256Point) Double(p *P256Point) *P256Point

// AddAffine sets q = p1 + p2, if infinity == 0, and to p1 if infinity == 1.
// p2 can't be the point at infinity as it can't be represented in affine
// coordinates, instead callers can set p2 to an arbitrary point and set
// infinity to 1.
func (q *P256Point) AddAffine(p1 *P256Point, p2 *p256AffinePoint, infinity int) *P256Point

// Select sets q to p1 if cond == 1, and to p2 if cond == 0.
func (q *P256Point) Select(p1, p2 *P256Point, cond int) *P256Point

// ScalarMult sets r = scalar * q, where scalar is a 32-byte big endian value,
// and returns r. If scalar is not 32 bytes long, ScalarMult returns an error
// and the receiver is unchanged.
func (p *P256Point) ScalarMult(q *P256Point, scalar []byte) (*P256Point, error)

// Negate sets p to -p, if cond == 1, and to p if cond == 0.
func (p *P256Point) Negate(cond int) *P256Point

// ScalarBaseMult sets p = scalar * generator, where scalar is a 32-byte big
// endian value, and returns r. If scalar is not 32 bytes long, ScalarBaseMult
// returns an error and the receiver is unchanged.
func (p *P256Point) ScalarBaseMult(scalar []byte) (*P256Point, error)