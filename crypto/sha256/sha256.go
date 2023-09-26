// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha256 implements the SHA224 and SHA256 hash algorithms as defined
// in FIPS 180-4.
package sha256

import (
	"github.com/shogo82148/std/hash"
)

// The size of a SHA256 checksum in bytes.
const Size = 32

// The size of a SHA224 checksum in bytes.
const Size224 = 28

// The blocksize of SHA256 and SHA224 in bytes.
const BlockSize = 64

// digest represents the partial evaluation of a checksum.

// New returns a new hash.Hash computing the SHA256 checksum. The Hash
// also implements encoding.BinaryMarshaler and
// encoding.BinaryUnmarshaler to marshal and unmarshal the internal
// state of the hash.
func New() hash.Hash

// New224 returns a new hash.Hash computing the SHA224 checksum.
func New224() hash.Hash

// Sum256 returns the SHA256 checksum of the data.
func Sum256(data []byte) [Size]byte

// Sum224 returns the SHA224 checksum of the data.
func Sum224(data []byte) (sum224 [Size224]byte)
