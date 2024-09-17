// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hash implements hash functions used in the compiler toolchain.
package hash

import (
	"github.com/shogo82148/std/crypto/md5"
	"github.com/shogo82148/std/crypto/sha1"
	"github.com/shogo82148/std/crypto/sha256"
	"github.com/shogo82148/std/hash"
)

const (
	// Size32 is the size of 32 bytes hash checksum.
	Size32 = sha256.Size
	// Size20 is the size of 20 bytes hash checksum.
	Size20 = sha1.Size
	// Size16 is the size of 16 bytes hash checksum.
	Size16 = md5.Size
)

// New32 returns a new [hash.Hash] computing the 32 bytes hash checksum.
func New32() hash.Hash

// New20 returns a new [hash.Hash] computing the 20 bytes hash checksum.
func New20() hash.Hash

// New16 returns a new [hash.Hash] computing the 16 bytes hash checksum.
func New16() hash.Hash

// Sum32 returns the 32 bytes checksum of the data.
func Sum32(data []byte) [Size32]byte

// Sum20 returns the 20 bytes checksum of the data.
func Sum20(data []byte) [Size20]byte

// Sum16 returns the 16 bytes checksum of the data.
func Sum16(data []byte) [Size16]byte