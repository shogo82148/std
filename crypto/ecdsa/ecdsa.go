// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ecdsa implements the Elliptic Curve Digital Signature Algorithm, as
// defined in FIPS 186-3.
//
// This implementation derives the nonce from an AES-CTR CSPRNG keyed by:
//
// SHA2-512(priv.D || entropy || hash)[:32]
//
// The CSPRNG key is indifferentiable from a random oracle as shown in
// [Coron], the AES-CTR stream is indifferentiable from a random oracle
// under standard cryptographic assumptions (see [Larsson] for examples).
//
// References:
//
//	[Coron]
//	  https://cs.nyu.edu/~dodis/ps/merkle.pdf
//	[Larsson]
//	  https://www.nada.kth.se/kurser/kth/2D1441/semteo03/lecturenotes/assump.pdf
package ecdsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/elliptic"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// A invertible implements fast inverse mod Curve.Params().N

// combinedMult implements fast multiplication S1*g + S2*p (g - generator, p - arbitrary point)

// PublicKey represents an ECDSA public key.
type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

// Equal reports whether pub and x have the same value.
//
// Two keys are only considered to have the same value if they have the same Curve value.
// Note that for example elliptic.P256() and elliptic.P256().Params() are different
// values, as the latter is a generic not constant time implementation.
func (pub *PublicKey) Equal(x crypto.PublicKey) bool

// PrivateKey represents an ECDSA private key.
type PrivateKey struct {
	PublicKey
	D *big.Int
}

// Public returns the public key corresponding to priv.
func (priv *PrivateKey) Public() crypto.PublicKey

// Equal reports whether priv and x have the same value.
//
// See PublicKey.Equal for details on how Curve is compared.
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool

// Sign signs digest with priv, reading randomness from rand. The opts argument
// is not currently used but, in keeping with the crypto.Signer interface,
// should be the hash function used to digest the message.
//
// This method implements crypto.Signer, which is an interface to support keys
// where the private part is kept in, for example, a hardware module. Common
// uses should use the Sign function in this package directly.
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)

// GenerateKey generates a public and private key pair.
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)

// Sign signs a hash (which should be the result of hashing a larger message)
// using the private key, priv. If the hash is longer than the bit-length of the
// private key's curve order, the hash will be truncated to that length. It
// returns the signature as a pair of integers. The security of the private key
// depends on the entropy of rand.
func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error)

// SignASN1 signs a hash (which should be the result of hashing a larger message)
// using the private key, priv. If the hash is longer than the bit-length of the
// private key's curve order, the hash will be truncated to that length. It
// returns the ASN.1 encoded signature. The security of the private key
// depends on the entropy of rand.
func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error)

// Verify verifies the signature in r, s of hash using the public key, pub. Its
// return value records whether the signature is valid.
func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool

// VerifyASN1 verifies the ASN.1 encoded signature, sig, of hash using the
// public key, pub. Its return value records whether the signature is valid.
func VerifyASN1(pub *PublicKey, hash, sig []byte) bool
