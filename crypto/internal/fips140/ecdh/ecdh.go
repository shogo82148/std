// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdh

import (
	"github.com/shogo82148/std/crypto/internal/fips140/nistec"
	"github.com/shogo82148/std/io"
)

type PrivateKey struct {
	pub PublicKey
	d   []byte
}

func (priv *PrivateKey) Bytes() []byte

func (priv *PrivateKey) PublicKey() *PublicKey

type PublicKey struct {
	curve curveID
	q     []byte
}

func (pub *PublicKey) Bytes() []byte

type Curve[P Point[P]] struct {
	curve    curveID
	newPoint func() P
	N        []byte
}

// Point is a generic constraint for the [nistec] Point types.
type Point[P any] interface {
	*nistec.P224Point | *nistec.P256Point | *nistec.P384Point | *nistec.P521Point
	Bytes() []byte
	BytesX() ([]byte, error)
	SetBytes([]byte) (P, error)
	ScalarMult(P, []byte) (P, error)
	ScalarBaseMult([]byte) (P, error)
}

func P224() *Curve[*nistec.P224Point]

func P256() *Curve[*nistec.P256Point]

func P384() *Curve[*nistec.P384Point]

func P521() *Curve[*nistec.P521Point]

// GenerateKey generates a new ECDSA private key pair for the specified curve.
func GenerateKey[P Point[P]](c *Curve[P], rand io.Reader) (*PrivateKey, error)

func NewPrivateKey[P Point[P]](c *Curve[P], key []byte) (*PrivateKey, error)

func NewPublicKey[P Point[P]](c *Curve[P], key []byte) (*PublicKey, error)

func ECDH[P Point[P]](c *Curve[P], k *PrivateKey, peer *PublicKey) ([]byte, error)
