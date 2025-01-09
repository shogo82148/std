// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlkem implements the quantum-resistant key encapsulation method
// ML-KEM (formerly known as Kyber), as specified in [NIST FIPS 203].
//
// [NIST FIPS 203]: https://doi.org/10.6028/NIST.FIPS.203
package mlkem

import "github.com/shogo82148/std/crypto/internal/fips140/mlkem"

const (
	// SharedKeySize is the size of a shared key produced by ML-KEM.
	SharedKeySize = 32

	// SeedSize is the size of a seed used to generate a decapsulation key.
	SeedSize = 64

	// CiphertextSize768 is the size of a ciphertext produced by the 768-bit
	// variant of ML-KEM.
	CiphertextSize768 = 1088

	// EncapsulationKeySize768 is the size of an encapsulation key for the
	// 768-bit variant of ML-KEM.
	EncapsulationKeySize768 = 1184
)

// DecapsulationKey768 is the secret key used to decapsulate a shared key
// from a ciphertext. It includes various precomputed values.
type DecapsulationKey768 struct {
	key *mlkem.DecapsulationKey768
}

// GenerateKey768 generates a new decapsulation key, drawing random bytes from
// crypto/rand. The decapsulation key must be kept secret.
func GenerateKey768() (*DecapsulationKey768, error)

// NewDecapsulationKey768 parses a decapsulation key from a 64-byte seed in the
// "d || z" form. The seed must be uniformly random.
func NewDecapsulationKey768(seed []byte) (*DecapsulationKey768, error)

// Bytes returns the decapsulation key as a 64-byte seed in the "d || z" form.
//
// The decapsulation key must be kept secret.
func (dk *DecapsulationKey768) Bytes() []byte

// Decapsulate generates a shared key from a ciphertext and a decapsulation
// key. If the ciphertext is not valid, Decapsulate returns an error.
//
// The shared key must be kept secret.
func (dk *DecapsulationKey768) Decapsulate(ciphertext []byte) (sharedKey []byte, err error)

// EncapsulationKey returns the public encapsulation key necessary to produce
// ciphertexts.
func (dk *DecapsulationKey768) EncapsulationKey() *EncapsulationKey768

// An EncapsulationKey768 is the public key used to produce ciphertexts to be
// decapsulated by the corresponding DecapsulationKey768.
type EncapsulationKey768 struct {
	key *mlkem.EncapsulationKey768
}

// NewEncapsulationKey768 parses an encapsulation key from its encoded form. If
// the encapsulation key is not valid, NewEncapsulationKey768 returns an error.
func NewEncapsulationKey768(encapsulationKey []byte) (*EncapsulationKey768, error)

// Bytes returns the encapsulation key as a byte slice.
func (ek *EncapsulationKey768) Bytes() []byte

// Encapsulate generates a shared key and an associated ciphertext from an
// encapsulation key, drawing random bytes from crypto/rand.
//
// The shared key must be kept secret.
func (ek *EncapsulationKey768) Encapsulate() (sharedKey, ciphertext []byte)
