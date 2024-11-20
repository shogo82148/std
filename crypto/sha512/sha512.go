// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256
// hash algorithms as defined in FIPS 180-4.
//
// All the hash.Hash implementations returned by this package also
// implement encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to
// marshal and unmarshal the internal state of the hash.
package sha512

import (
	"github.com/shogo82148/std/hash"
)

const (
	// Size is the size, in bytes, of a SHA-512 checksum.
	Size = 64

	// Size224 is the size, in bytes, of a SHA-512/224 checksum.
	Size224 = 28

	// Size256 is the size, in bytes, of a SHA-512/256 checksum.
	Size256 = 32

	// Size384 is the size, in bytes, of a SHA-384 checksum.
	Size384 = 48

	// BlockSize is the block size, in bytes, of the SHA-512/224,
	// SHA-512/256, SHA-384 and SHA-512 hash functions.
	BlockSize = 128
)

// New returns a new [hash.Hash] computing the SHA-512 checksum. The Hash
// also implements [encoding.BinaryMarshaler], [encoding.BinaryAppender] and
// [encoding.BinaryUnmarshaler] to marshal and unmarshal the internal
// state of the hash.
func New() hash.Hash

// New512_224 returns a new [hash.Hash] computing the SHA-512/224 checksum. The Hash
// also implements [encoding.BinaryMarshaler], [encoding.BinaryAppender] and
// [encoding.BinaryUnmarshaler] to marshal and unmarshal the internal
// state of the hash.
func New512_224() hash.Hash

// New512_256 returns a new [hash.Hash] computing the SHA-512/256 checksum. The Hash
// also implements [encoding.BinaryMarshaler], [encoding.BinaryAppender] and
// [encoding.BinaryUnmarshaler] to marshal and unmarshal the internal
// state of the hash.
func New512_256() hash.Hash

// New384 returns a new [hash.Hash] computing the SHA-384 checksum. The Hash
// also implements [encoding.BinaryMarshaler], [encoding.BinaryAppender] and
// [encoding.BinaryUnmarshaler] to marshal and unmarshal the internal
// state of the hash.
func New384() hash.Hash

// Sum512 returns the SHA512 checksum of the data.
func Sum512(data []byte) [Size]byte

// Sum384 returns the SHA384 checksum of the data.
func Sum384(data []byte) [Size384]byte

// Sum512_224 returns the Sum512/224 checksum of the data.
func Sum512_224(data []byte) [Size224]byte

// Sum512_256 returns the Sum512/256 checksum of the data.
func Sum512_256(data []byte) [Size256]byte
