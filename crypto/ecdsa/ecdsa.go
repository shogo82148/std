// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ecdsa implements the Elliptic Curve Digital Signature Algorithm, as
// defined in FIPS 186-4 and SEC 1, Version 2.0.
//
// Signatures generated by this package are not deterministic, but entropy is
// mixed with the private key and the message, achieving the same level of
// security in case of randomness source failure.
package ecdsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/ecdh"
	"github.com/shogo82148/std/crypto/elliptic"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// PublicKey represents an ECDSA public key.
type PublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

// ECDH returns k as a [ecdh.PublicKey]. It returns an error if the key is
// invalid according to the definition of [ecdh.Curve.NewPublicKey], or if the
// Curve is not supported by crypto/ecdh.
func (k *PublicKey) ECDH() (*ecdh.PublicKey, error)

// Equal reports whether pub and x have the same value.
//
// Two keys are only considered to have the same value if they have the same Curve value.
// Note that for example [elliptic.P256] and elliptic.P256().Params() are different
// values, as the latter is a generic not constant time implementation.
func (pub *PublicKey) Equal(x crypto.PublicKey) bool

// PrivateKey represents an ECDSA private key.
type PrivateKey struct {
	PublicKey
	D *big.Int
}

// ECDH returns k as a [ecdh.PrivateKey]. It returns an error if the key is
// invalid according to the definition of [ecdh.Curve.NewPrivateKey], or if the
// Curve is not supported by [crypto/ecdh].
func (k *PrivateKey) ECDH() (*ecdh.PrivateKey, error)

// Public returns the public key corresponding to priv.
func (priv *PrivateKey) Public() crypto.PublicKey

// Equal reports whether priv and x have the same value.
//
// See [PublicKey.Equal] for details on how Curve is compared.
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool

// Sign signs digest with priv, reading randomness from rand. The opts argument
// is not currently used but, in keeping with the crypto.Signer interface,
// should be the hash function used to digest the message.
//
// This method implements crypto.Signer, which is an interface to support keys
// where the private part is kept in, for example, a hardware module. Common
// uses can use the [SignASN1] function in this package directly.
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)

// GenerateKey generates a new ECDSA private key for the specified curve.
//
// Most applications should use [crypto/rand.Reader] as rand. Note that the
// returned key does not depend deterministically on the bytes read from rand,
// and may change between calls and/or between versions.
func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error)

// SignASN1 signs a hash (which should be the result of hashing a larger message)
// using the private key, priv. If the hash is longer than the bit-length of the
// private key's curve order, the hash will be truncated to that length. It
// returns the ASN.1 encoded signature.
//
// The signature is randomized. Most applications should use [crypto/rand.Reader]
// as rand. Note that the returned signature does not depend deterministically on
// the bytes read from rand, and may change between calls and/or between versions.
func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error)

// VerifyASN1 verifies the ASN.1 encoded signature, sig, of hash using the
// public key, pub. Its return value records whether the signature is valid.
func VerifyASN1(pub *PublicKey, hash, sig []byte) bool
