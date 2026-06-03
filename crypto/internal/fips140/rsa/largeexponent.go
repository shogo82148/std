// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto/internal/fips140/bigmod"
	"github.com/shogo82148/std/hash"
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
