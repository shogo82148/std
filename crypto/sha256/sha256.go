// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha256 implements the SHA224 and SHA256 hash algorithms as defined
// in FIPS 180-2.
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

// New returns a new hash.Hash computing the SHA256 checksum.
func New() hash.Hash

// New224 returns a new hash.Hash computing the SHA224 checksum.
func New224() hash.Hash
