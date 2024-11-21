// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdsa

import (
	"github.com/shogo82148/std/crypto/internal/fips140"
	"github.com/shogo82148/std/crypto/internal/fips140/bigmod"
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
	curve      curveID
	newPoint   func() P
	ordInverse func([]byte) ([]byte, error)
	N          *bigmod.Modulus
	nMinus2    []byte
}

// Point is a generic constraint for the [nistec] Point types.
type Point[P any] interface {
	*nistec.P224Point | *nistec.P256Point | *nistec.P384Point | *nistec.P521Point
	Bytes() []byte
	BytesX() ([]byte, error)
	SetBytes([]byte) (P, error)
	ScalarMult(P, []byte) (P, error)
	ScalarBaseMult([]byte) (P, error)
	Add(p1, p2 P) P
}

func P224() *Curve[*nistec.P224Point]

func P256() *Curve[*nistec.P256Point]

func P384() *Curve[*nistec.P384Point]

func P521() *Curve[*nistec.P521Point]

func NewPrivateKey[P Point[P]](c *Curve[P], D, Q []byte) (*PrivateKey, error)

func NewPublicKey[P Point[P]](c *Curve[P], Q []byte) (*PublicKey, error)

// GenerateKey generates a new ECDSA private key pair for the specified curve.
//
// In FIPS mode, rand is ignored.
func GenerateKey[P Point[P]](c *Curve[P], rand io.Reader) (*PrivateKey, error)

// Signature is an ECDSA signature, where r and s are represented as big-endian
// fixed-length byte slices.
type Signature struct {
	R, S []byte
}

// Sign signs a hash (which shall be the result of hashing a larger message with
// the hash function H) using the private key, priv. If the hash is longer than
// the bit-length of the private key's curve order, the hash will be truncated
// to that length.
//
// The signature is randomized. If FIPS mode is enabled, rand is ignored.
func Sign[P Point[P], H fips140.Hash](c *Curve[P], h func() H, priv *PrivateKey, rand io.Reader, hash []byte) (*Signature, error)

// SignDeterministic signs a hash (which shall be the result of hashing a
// larger message with the hash function H) using the private key, priv. If the
// hash is longer than the bit-length of the private key's curve order, the hash
// will be truncated to that length. This applies Deterministic ECDSA as
// specified in FIPS 186-5 and RFC 6979.
func SignDeterministic[P Point[P], H fips140.Hash](c *Curve[P], h func() H, priv *PrivateKey, hash []byte) (*Signature, error)

// Verify verifies the signature, sig, of hash (which should be the result of
// hashing a larger message) using the public key, pub. If the hash is longer
// than the bit-length of the private key's curve order, the hash will be
// truncated to that length.
//
// The inputs are not considered confidential, and may leak through timing side
// channels, or if an attacker has control of part of the inputs.
func Verify[P Point[P]](c *Curve[P], pub *PublicKey, hash []byte, sig *Signature) error
