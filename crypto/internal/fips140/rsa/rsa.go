// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto/internal/fips140/bigmod"
	"github.com/shogo82148/std/errors"
)

type PublicKey struct {
	N *bigmod.Modulus
	E int
}

// Size returns the modulus size in bytes. Raw signatures and ciphertexts
// for or by this public key will have the same size.
func (pub *PublicKey) Size() int

type PrivateKey struct {
	// pub has already been checked with checkPublicKey.
	pub PublicKey
	d   *bigmod.Nat
	// The following values are not set for deprecated multi-prime keys.
	//
	// Since they are always set for keys in FIPS mode, for SP 800-56B Rev. 2
	// purposes we always use the Chinese Remainder Theorem (CRT) format.
	p, q *bigmod.Modulus
	// dP and dQ are used as exponents, so we store them as big-endian byte
	// slices to be passed to [bigmod.Nat.Exp].
	dP   []byte
	dQ   []byte
	qInv *bigmod.Nat
}

func (priv *PrivateKey) PublicKey() *PublicKey

// NewPrivateKey creates a new RSA private key from the given parameters.
//
// All values are in big-endian byte slice format, and may have leading zeros
// or be shorter if leading zeroes were trimmed.
func NewPrivateKey(N []byte, e int, d, P, Q []byte) (*PrivateKey, error)

// NewPrivateKeyWithPrecomputation creates a new RSA private key from the given
// parameters, which include precomputed CRT values.
func NewPrivateKeyWithPrecomputation(N []byte, e int, d, P, Q, dP, dQ, qInv []byte) (*PrivateKey, error)

// NewPrivateKeyWithoutCRT creates a new RSA private key from the given parameters.
//
// This is meant for deprecated multi-prime keys, and is not FIPS 140 compliant.
func NewPrivateKeyWithoutCRT(N []byte, e int, d []byte) (*PrivateKey, error)

// Export returns the key parameters in big-endian byte slice format.
//
// P, Q, dP, dQ, and qInv may be nil if the key was created with
// NewPrivateKeyWithoutCRT.
func (priv *PrivateKey) Export() (N []byte, e int, d, P, Q, dP, dQ, qInv []byte)

// Encrypt performs the RSA public key operation.
func Encrypt(pub *PublicKey, plaintext []byte) ([]byte, error)

var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA key size")
var ErrDecryption = errors.New("crypto/rsa: decryption error")
var ErrVerification = errors.New("crypto/rsa: verification error")

// DecryptWithoutCheck performs the RSA private key operation.
func DecryptWithoutCheck(priv *PrivateKey, ciphertext []byte) ([]byte, error)

// DecryptWithCheck performs the RSA private key operation and checks the
// result to defend against errors in the CRT computation.
func DecryptWithCheck(priv *PrivateKey, ciphertext []byte) ([]byte, error)
