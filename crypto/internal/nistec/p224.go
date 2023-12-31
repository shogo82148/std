// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate.go. DO NOT EDIT.

package nistec

import (
	"github.com/shogo82148/std/crypto/internal/nistec/fiat"
)

// P224Point is a P224 point. The zero value is NOT valid.
type P224Point struct {
	// The point is represented in projective coordinates (X:Y:Z),
	// where x = X/Z and y = Y/Z.
	x, y, z *fiat.P224Element
}

// NewP224Point returns a new P224Point representing the point at infinity point.
func NewP224Point() *P224Point

// SetGenerator sets p to the canonical generator and returns p.
func (p *P224Point) SetGenerator() *P224Point

// Set sets p = q and returns p.
func (p *P224Point) Set(q *P224Point) *P224Point

// SetBytes sets p to the compressed, uncompressed, or infinity value encoded in
// b, as specified in SEC 1, Version 2.0, Section 2.3.4. If the point is not on
// the curve, it returns nil and an error, and the receiver is unchanged.
// Otherwise, it returns p.
func (p *P224Point) SetBytes(b []byte) (*P224Point, error)

// Bytes returns the uncompressed or infinity encoding of p, as specified in
// SEC 1, Version 2.0, Section 2.3.3. Note that the encoding of the point at
// infinity is shorter than all other encodings.
func (p *P224Point) Bytes() []byte

// BytesX returns the encoding of the x-coordinate of p, as specified in SEC 1,
// Version 2.0, Section 2.3.5, or an error if p is the point at infinity.
func (p *P224Point) BytesX() ([]byte, error)

// BytesCompressed returns the compressed or infinity encoding of p, as
// specified in SEC 1, Version 2.0, Section 2.3.3. Note that the encoding of the
// point at infinity is shorter than all other encodings.
func (p *P224Point) BytesCompressed() []byte

// Add sets q = p1 + p2, and returns q. The points may overlap.
func (q *P224Point) Add(p1, p2 *P224Point) *P224Point

// Double sets q = p + p, and returns q. The points may overlap.
func (q *P224Point) Double(p *P224Point) *P224Point

// Select sets q to p1 if cond == 1, and to p2 if cond == 0.
func (q *P224Point) Select(p1, p2 *P224Point, cond int) *P224Point

// ScalarMult sets p = scalar * q, and returns p.
func (p *P224Point) ScalarMult(q *P224Point, scalar []byte) (*P224Point, error)

// ScalarBaseMult sets p = scalar * B, where B is the canonical generator, and
// returns p.
func (p *P224Point) ScalarBaseMult(scalar []byte) (*P224Point, error)
