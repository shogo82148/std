// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha1 implements the SHA1 hash algorithm as defined in RFC 3174.
package sha1

import (
	"github.com/shogo82148/std/hash"
)

// The size of a SHA1 checksum in bytes.
const Size = 20

// The blocksize of SHA1 in bytes.
const BlockSize = 64

// digest represents the partial evaluation of a checksum.

// New returns a new hash.Hash computing the SHA1 checksum.
func New() hash.Hash
