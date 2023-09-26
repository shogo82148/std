// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dsa implements the Digital Signature Algorithm, as defined in FIPS 186-3.
package dsa

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// Parameters represents the domain parameters for a key. These parameters can
// be shared across many keys. The bit length of Q must be a multiple of 8.
type Parameters struct {
	P, Q, G *big.Int
}

// PublicKey represents a DSA public key.
type PublicKey struct {
	Parameters
	Y *big.Int
}

// PrivateKey represents a DSA private key.
type PrivateKey struct {
	PublicKey
	X *big.Int
}

// ErrInvalidPublicKey results when a public key is not usable by this code.
// FIPS is quite strict about the format of DSA keys, but other code may be
// less so. Thus, when using keys which may have been generated by other code,
// this error must be handled.
var ErrInvalidPublicKey = errors.New("crypto/dsa: invalid public key")

// ParameterSizes is a enumeration of the acceptable bit lengths of the primes
// in a set of DSA parameters. See FIPS 186-3, section 4.2.
type ParameterSizes int

const (
	L1024N160 ParameterSizes = iota
	L2048N224
	L2048N256
	L3072N256
)

// numMRTests is the number of Miller-Rabin primality tests that we perform. We
// pick the largest recommended number from table C.1 of FIPS 186-3.

// GenerateParameters puts a random, valid set of DSA parameters into params.
// This function can take many seconds, even on fast machines.
func GenerateParameters(params *Parameters, rand io.Reader, sizes ParameterSizes) error

// GenerateKey generates a public&private key pair. The Parameters of the
// PrivateKey must already be valid (see GenerateParameters).
func GenerateKey(priv *PrivateKey, rand io.Reader) error

// Sign signs an arbitrary length hash (which should be the result of hashing a
// larger message) using the private key, priv. It returns the signature as a
// pair of integers. The security of the private key depends on the entropy of
// rand.
//
// Note that FIPS 186-3 section 4.6 specifies that the hash should be truncated
// to the byte-length of the subgroup. This function does not perform that
// truncation itself.
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

// Verify verifies the signature in r, s of hash using the public key, pub. It
// reports whether the signature is valid.
//
// Note that FIPS 186-3 section 4.6 specifies that the hash should be truncated
// to the byte-length of the subgroup. This function does not perform that
// truncation itself.
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool
