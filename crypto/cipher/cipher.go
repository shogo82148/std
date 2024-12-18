// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cipher implements standard block cipher modes that can be wrapped
// around low-level block cipher implementations.
// See https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html
// and NIST Special Publication 800-38A.
package cipher

// A Block represents an implementation of block cipher
// using a given key. It provides the capability to encrypt
// or decrypt individual blocks. The mode implementations
// extend that capability to streams of blocks.
type Block interface {
	BlockSize() int

	Encrypt(dst, src []byte)

	Decrypt(dst, src []byte)
}

// A Stream represents a stream cipher.
type Stream interface {
	XORKeyStream(dst, src []byte)
}

// A BlockMode represents a block cipher running in a block-based mode (CBC,
// ECB etc).
type BlockMode interface {
	BlockSize() int

	CryptBlocks(dst, src []byte)
}

// AEAD is a cipher mode providing authenticated encryption with associated
// data. For a description of the methodology, see
// https://en.wikipedia.org/wiki/Authenticated_encryption.
type AEAD interface {
	NonceSize() int

	Overhead() int

	Seal(dst, nonce, plaintext, additionalData []byte) []byte

	Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}
