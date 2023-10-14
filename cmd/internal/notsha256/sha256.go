// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package notsha256 implements the NOTSHA256 algorithm,
// a hash defined as bitwise NOT of SHA256.
// It is used in situations where exact fidelity to SHA256 is unnecessary.
// In particular, it is used in the compiler toolchain,
// which cannot depend directly on cgo when GOEXPERIMENT=boringcrypto
// (and in that mode the real sha256 uses cgo).
package notsha256

import (
	"github.com/shogo82148/std/hash"
)

// The size of a checksum in bytes.
const Size = 32

// The blocksize in bytes.
const BlockSize = 64

// New returns a new hash.Hash computing the NOTSHA256 checksum.
// state of the hash.
func New() hash.Hash

// Sum256 returns the SHA256 checksum of the data.
func Sum256(data []byte) [Size]byte
