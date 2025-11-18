// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlkem implements the quantum-resistant key encapsulation method
// ML-KEM (formerly known as Kyber), as specified in [NIST FIPS 203].
//
// Most applications should use the ML-KEM-768 parameter set, as implemented by
// [DecapsulationKey768] and [EncapsulationKey768].
//
// [NIST FIPS 203]: https://doi.org/10.6028/NIST.FIPS.203
package mlkem

import "github.com/shogo82148/std/crypto/internal/fips140/mlkem"

const (
	// SharedKeySize is the size of a shared key produced by ML-KEM.
	SharedKeySize = 32

	// SeedSize is the size of a seed used to generate a decapsulation key.
	SeedSize = 64

	// CiphertextSize768 is the size of a ciphertext produced by ML-KEM-768.
	CiphertextSize768 = 1088

	// EncapsulationKeySize768 is the size of an ML-KEM-768 encapsulation key.
	EncapsulationKeySize768 = 1184

	// CiphertextSize1024 is the size of a ciphertext produced by ML-KEM-1024.
	CiphertextSize1024 = 1568

	// EncapsulationKeySize1024 is the size of an ML-KEM-1024 encapsulation key.
	EncapsulationKeySize1024 = 1568
)

// DecapsulationKey768 is the secret key used to decapsulate a shared key
// from a ciphertext. It includes various precomputed values.
type DecapsulationKey768 struct {
	key *mlkem.DecapsulationKey768
}

// GenerateKey768 generates a new decapsulation key, drawing random bytes from
// the default crypto/rand source. The decapsulation key must be kept secret.
func GenerateKey768() (*DecapsulationKey768, error)

// NewDecapsulationKey768 expands a decapsulation key from a 64-byte seed in the
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
// encapsulation key, drawing random bytes from the default crypto/rand source.
//
// The shared key must be kept secret.
//
// For testing, derandomized encapsulation is provided by the
// [crypto/mlkem/mlkemtest] package.
func (ek *EncapsulationKey768) Encapsulate() (sharedKey, ciphertext []byte)

// DecapsulationKey1024 is the secret key used to decapsulate a shared key
// from a ciphertext. It includes various precomputed values.
type DecapsulationKey1024 struct {
	key *mlkem.DecapsulationKey1024
}

// GenerateKey1024 generates a new decapsulation key, drawing random bytes from
// the default crypto/rand source. The decapsulation key must be kept secret.
func GenerateKey1024() (*DecapsulationKey1024, error)

// NewDecapsulationKey1024 expands a decapsulation key from a 64-byte seed in the
// "d || z" form. The seed must be uniformly random.
func NewDecapsulationKey1024(seed []byte) (*DecapsulationKey1024, error)

// Bytes returns the decapsulation key as a 64-byte seed in the "d || z" form.
//
// The decapsulation key must be kept secret.
func (dk *DecapsulationKey1024) Bytes() []byte

// Decapsulate generates a shared key from a ciphertext and a decapsulation
// key. If the ciphertext is not valid, Decapsulate returns an error.
//
// The shared key must be kept secret.
func (dk *DecapsulationKey1024) Decapsulate(ciphertext []byte) (sharedKey []byte, err error)

// EncapsulationKey returns the public encapsulation key necessary to produce
// ciphertexts.
func (dk *DecapsulationKey1024) EncapsulationKey() *EncapsulationKey1024

// An EncapsulationKey1024 is the public key used to produce ciphertexts to be
// decapsulated by the corresponding DecapsulationKey1024.
type EncapsulationKey1024 struct {
	key *mlkem.EncapsulationKey1024
}

// NewEncapsulationKey1024 parses an encapsulation key from its encoded form. If
// the encapsulation key is not valid, NewEncapsulationKey1024 returns an error.
func NewEncapsulationKey1024(encapsulationKey []byte) (*EncapsulationKey1024, error)

// Bytes returns the encapsulation key as a byte slice.
func (ek *EncapsulationKey1024) Bytes() []byte

// Encapsulate generates a shared key and an associated ciphertext from an
// encapsulation key, drawing random bytes from the default crypto/rand source.
//
// The shared key must be kept secret.
//
// For testing, derandomized encapsulation is provided by the
// [crypto/mlkem/mlkemtest] package.
func (ek *EncapsulationKey1024) Encapsulate() (sharedKey, ciphertext []byte)
