// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto/internal/fips140/bigmod"
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
)

// TestingOnlyLargeExponentPublicKey is a variant of [PublicKey] that supports
// large public exponents. It is only meant for supporting the full ACVP test
// suite, which unfortunately forces us to choose between a fixed exponent and
// the full (2¹⁶, 2²⁵⁶) range. This type must not be used in production code,
// nor for e values < 2³¹, which instead must use [PublicKey].
type TestingOnlyLargeExponentPublicKey struct {
	N *bigmod.Modulus
	// E is the public exponent, represented as a big-endian byte slice.
	E []byte
}

func (pub *TestingOnlyLargeExponentPublicKey) Size() int

func TestingOnlyLargeExponentVerifyPKCS1v15(pub *TestingOnlyLargeExponentPublicKey, hash string, hashed []byte, sig []byte) error

func TestingOnlyLargeExponentVerifyPSS(pub *TestingOnlyLargeExponentPublicKey, hash hash.Hash, digest []byte, sig []byte) error

// TestingOnlyLargeExponentPrivateKey is a variant of [PrivateKey] that supports
// large public exponents. It is only meant for supporting the full ACVP test
// suite. This type must not be used in production code.
type TestingOnlyLargeExponentPrivateKey struct {
	n    *bigmod.Modulus
	e    []byte
	d    *bigmod.Nat
	p, q *bigmod.Modulus
	dP   []byte
	dQ   []byte
	qInv *bigmod.Nat
}

func (priv *TestingOnlyLargeExponentPrivateKey) Size() int

// TestingOnlyNewLargeExponentPrivateKeyWithPrecomputation creates a new RSA private key
// with a large public exponent from the given parameters. It is only meant for ACVP testing.
func TestingOnlyNewLargeExponentPrivateKeyWithPrecomputation(N []byte, e []byte, d, P, Q, dP, dQ, qInv []byte) (*TestingOnlyLargeExponentPrivateKey, error)

// TestingOnlyLargeExponentDecryptOAEP decrypts ciphertext using RSAES-OAEP with
// a private key that has a large public exponent. It is only meant for ACVP testing.
func TestingOnlyLargeExponentDecryptOAEP(hash, mgfHash hash.Hash, priv *TestingOnlyLargeExponentPrivateKey, ciphertext []byte, label []byte) ([]byte, error)

// TestingOnlyLargeExponentEncryptOAEP encrypts the given message with RSAES-OAEP
// using a public key with a large exponent. It is only meant for ACVP testing.
func TestingOnlyLargeExponentEncryptOAEP(hash, mgfHash hash.Hash, random io.Reader, pub *TestingOnlyLargeExponentPublicKey, msg []byte, label []byte) ([]byte, error)
