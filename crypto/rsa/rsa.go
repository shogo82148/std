// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rsa implements RSA encryption as specified in PKCS#1.
package rsa

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
)

// A PublicKey represents the public part of an RSA key.
type PublicKey struct {
	N *big.Int
	E int
}

// A PrivateKey represents an RSA key
type PrivateKey struct {
	PublicKey
	D      *big.Int
	Primes []*big.Int

	Precomputed PrecomputedValues
}

type PrecomputedValues struct {
	Dp, Dq *big.Int
	Qinv   *big.Int

	CRTValues []CRTValue
}

// CRTValue contains the precomputed chinese remainder theorem values.
type CRTValue struct {
	Exp   *big.Int
	Coeff *big.Int
	R     *big.Int
}

// Validate performs basic sanity checks on the key.
// It returns nil if the key is valid, or else an error describing a problem.
func (priv *PrivateKey) Validate() error

// GenerateKey generates an RSA keypair of the given bit size.
func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error)

// GenerateMultiPrimeKey generates a multi-prime RSA keypair of the given bit
// size, as suggested in [1]. Although the public keys are compatible
// (actually, indistinguishable) from the 2-prime case, the private keys are
// not. Thus it may not be possible to export multi-prime private keys in
// certain formats or to subsequently import them into other code.
//
// Table 1 in [2] suggests maximum numbers of primes for a given size.
//
// [1] US patent 4405829 (1972, expired)
// [2] http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf
func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (priv *PrivateKey, err error)

// ErrMessageTooLong is returned when attempting to encrypt a message which is
// too large for the size of the public key.
var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA public key size")

// EncryptOAEP encrypts the given message with RSA-OAEP.
// The message must be no longer than the length of the public modulus less
// twice the hash length plus 2.
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) (out []byte, err error)

// ErrDecryption represents a failure to decrypt a message.
// It is deliberately vague to avoid adaptive attacks.
var ErrDecryption = errors.New("crypto/rsa: decryption error")

// ErrVerification represents a failure to verify a signature.
// It is deliberately vague to avoid adaptive attacks.
var ErrVerification = errors.New("crypto/rsa: verification error")

// Precompute performs some calculations that speed up private key operations
// in the future.
func (priv *PrivateKey) Precompute()

// DecryptOAEP decrypts ciphertext using RSA-OAEP.
// If random != nil, DecryptOAEP uses RSA blinding to avoid timing side-channel attacks.
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) (msg []byte, err error)
