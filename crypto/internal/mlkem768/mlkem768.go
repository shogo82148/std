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

// GenerateKey generates an encapsulation key and a corresponding decapsulation
// key, drawing random bytes from crypto/rand.
//
// The decapsulation key must be kept secret.
func GenerateKey() (encapsulationKey, decapsulationKey []byte, err error)

// NewKeyFromSeed deterministically generates an encapsulation key and a
// corresponding decapsulation key from a 64-byte seed. The seed must be
// uniformly random.
func NewKeyFromSeed(seed []byte) (encapsulationKey, decapsulationKey []byte, err error)

// Encapsulate generates a shared key and an associated ciphertext from an
// encapsulation key, drawing random bytes from crypto/rand.
// If the encapsulation key is not valid, Encapsulate returns an error.
//
// The shared key must be kept secret.
func Encapsulate(encapsulationKey []byte) (ciphertext, sharedKey []byte, err error)

// Decapsulate generates a shared key from a ciphertext and a decapsulation key.
// If the decapsulation key or the ciphertext are not valid, Decapsulate returns
// an error.
//
// The shared key must be kept secret.
func Decapsulate(decapsulationKey, ciphertext []byte) (sharedKey []byte, err error)
