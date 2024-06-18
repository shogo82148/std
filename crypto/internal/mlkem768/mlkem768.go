// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mlkem768 implements the quantum-resistant key encapsulation method
// ML-KEM (formerly known as Kyber).
//
// Only the recommended ML-KEM-768 parameter set is provided.
//
// The version currently implemented is the one specified by [NIST FIPS 203 ipd],
// with the unintentional transposition of the matrix A reverted to match the
// behavior of [Kyber version 3.0]. Future versions of this package might
// introduce backwards incompatible changes to implement changes to FIPS 203.
//
// [Kyber version 3.0]: https://pq-crystals.org/kyber/data/kyber-specification-round3-20210804.pdf
// [NIST FIPS 203 ipd]: https://doi.org/10.6028/NIST.FIPS.203.ipd
package mlkem768

const (
	CiphertextSize       = k*encodingSize10 + encodingSize4
	EncapsulationKeySize = encryptionKeySize
	DecapsulationKeySize = decryptionKeySize + encryptionKeySize + 32 + 32
	SharedKeySize        = 32
	SeedSize             = 32 + 32
)

// A DecapsulationKey is the secret key used to decapsulate a shared key from a
// ciphertext. It includes various precomputed values.
type DecapsulationKey struct {
	dk [DecapsulationKeySize]byte
	encryptionKey
	decryptionKey
}

// Bytes returns the extended encoding of the decapsulation key, according to
// FIPS 203 (DRAFT).
func (dk *DecapsulationKey) Bytes() []byte

// EncapsulationKey returns the public encapsulation key necessary to produce
// ciphertexts.
func (dk *DecapsulationKey) EncapsulationKey() []byte

// GenerateKey generates a new decapsulation key, drawing random bytes from
// crypto/rand. The decapsulation key must be kept secret.
func GenerateKey() (*DecapsulationKey, error)

// NewKeyFromSeed deterministically generates a decapsulation key from a 64-byte
// seed in the "d || z" form. The seed must be uniformly random.
func NewKeyFromSeed(seed []byte) (*DecapsulationKey, error)

// NewKeyFromExtendedEncoding parses a decapsulation key from its FIPS 203
// (DRAFT) extended encoding.
func NewKeyFromExtendedEncoding(decapsulationKey []byte) (*DecapsulationKey, error)

// Encapsulate generates a shared key and an associated ciphertext from an
// encapsulation key, drawing random bytes from crypto/rand.
// If the encapsulation key is not valid, Encapsulate returns an error.
//
// The shared key must be kept secret.
func Encapsulate(encapsulationKey []byte) (ciphertext, sharedKey []byte, err error)

// Decapsulate generates a shared key from a ciphertext and a decapsulation key.
// If the ciphertext is not valid, Decapsulate returns an error.
//
// The shared key must be kept secret.
func Decapsulate(dk *DecapsulationKey, ciphertext []byte) (sharedKey []byte, err error)
