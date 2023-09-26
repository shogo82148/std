// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha1 implements the SHA-1 hash algorithm as defined in RFC 3174.
//
// SHA-1 is cryptographically broken and should not be used for secure
// applications.
package sha1

import (
	"github.com/shogo82148/std/hash"
)

// The size of a SHA-1 checksum in bytes.
const Size = 20

// The blocksize of SHA-1 in bytes.
const BlockSize = 64

// digest represents the partial evaluation of a checksum.

// New returns a new hash.Hash computing the SHA1 checksum.
func New() hash.Hash

// Sum returns the SHA-1 checksum of the data.
func Sum(data []byte) [Size]byte
