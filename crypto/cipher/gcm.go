// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

// AEAD is a cipher mode providing authenticated encryption with associated
// data.
type AEAD interface {
	NonceSize() int

	Overhead() int

	Seal(dst, nonce, plaintext, data []byte) []byte

	Open(dst, nonce, ciphertext, data []byte) ([]byte, error)
}

// gcmFieldElement represents a value in GF(2¹²⁸). In order to reflect the GCM
// standard and make getUint64 suitable for marshaling these values, the bits
// are stored backwards. For example:
//   the coefficient of x⁰ can be obtained by v.low >> 63.
//   the coefficient of x⁶³ can be obtained by v.low & 1.
//   the coefficient of x⁶⁴ can be obtained by v.high >> 63.
//   the coefficient of x¹²⁷ can be obtained by v.high & 1.

// gcm represents a Galois Counter Mode with a specific key. See
// http://csrc.nist.gov/groups/ST/toolkit/BCM/documents/proposedmodes/gcm/gcm-revised-spec.pdf

// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode.
func NewGCM(cipher Block) (AEAD, error)
