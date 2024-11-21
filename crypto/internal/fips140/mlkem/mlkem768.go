// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlkem implements the quantum-resistant key encapsulation method
// ML-KEM (formerly known as Kyber), as specified in [NIST FIPS 203].
//
// [NIST FIPS 203]: https://doi.org/10.6028/NIST.FIPS.203
package mlkem

const (
	SharedKeySize = 32
	SeedSize      = 32 + 32
)

// ML-KEM-768 parameters.
const (
	CiphertextSize768       = k*encodingSize10 + encodingSize4
	EncapsulationKeySize768 = k*encodingSize12 + 32
)

// ML-KEM-1024 parameters.
const (
	CiphertextSize1024       = k1024*encodingSize11 + encodingSize5
	EncapsulationKeySize1024 = k1024*encodingSize12 + 32
)

// A DecapsulationKey768 is the secret key used to decapsulate a shared key from a
// ciphertext. It includes various precomputed values.
type DecapsulationKey768 struct {
	d [32]byte
	z [32]byte

	ρ [32]byte
	h [32]byte

	encryptionKey
	decryptionKey
}

// Bytes returns the decapsulation key as a 64-byte seed in the "d || z" form.
//
// The decapsulation key must be kept secret.
func (dk *DecapsulationKey768) Bytes() []byte

// EncapsulationKey returns the public encapsulation key necessary to produce
// ciphertexts.
func (dk *DecapsulationKey768) EncapsulationKey() *EncapsulationKey768

// An EncapsulationKey768 is the public key used to produce ciphertexts to be
// decapsulated by the corresponding [DecapsulationKey768].
type EncapsulationKey768 struct {
	ρ [32]byte
	h [32]byte
	encryptionKey
}

// Bytes returns the encapsulation key as a byte slice.
func (ek *EncapsulationKey768) Bytes() []byte

// GenerateKey768 generates a new decapsulation key, drawing random bytes from
// a DRBG. The decapsulation key must be kept secret.
func GenerateKey768() (*DecapsulationKey768, error)

// GenerateKeyInternal768 is a derandomized version of GenerateKey768,
// exclusively for use in tests.
func GenerateKeyInternal768(d, z *[32]byte) *DecapsulationKey768

// NewDecapsulationKey768 parses a decapsulation key from a 64-byte
// seed in the "d || z" form. The seed must be uniformly random.
func NewDecapsulationKey768(seed []byte) (*DecapsulationKey768, error)

// Encapsulate generates a shared key and an associated ciphertext from an
// encapsulation key, drawing random bytes from a DRBG.
//
// The shared key must be kept secret.
func (ek *EncapsulationKey768) Encapsulate() (ciphertext, sharedKey []byte)

// EncapsulateInternal is a derandomized version of Encapsulate, exclusively for
// use in tests.
func (ek *EncapsulationKey768) EncapsulateInternal(m *[32]byte) (ciphertext, sharedKey []byte)

// NewEncapsulationKey768 parses an encapsulation key from its encoded form.
// If the encapsulation key is not valid, NewEncapsulationKey768 returns an error.
func NewEncapsulationKey768(encapsulationKey []byte) (*EncapsulationKey768, error)

// Decapsulate generates a shared key from a ciphertext and a decapsulation key.
// If the ciphertext is not valid, Decapsulate returns an error.
//
// The shared key must be kept secret.
func (dk *DecapsulationKey768) Decapsulate(ciphertext []byte) (sharedKey []byte, err error)
