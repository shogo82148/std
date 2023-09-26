// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go -full -output md5block.go

// Package md5 implements the MD5 hash algorithm as defined in RFC 1321.
package md5

import (
	"github.com/shogo82148/std/hash"
)

// The size of an MD5 checksum in bytes.
const Size = 16

// The blocksize of MD5 in bytes.
const BlockSize = 64

// digest represents the partial evaluation of a checksum.

// New returns a new hash.Hash computing the MD5 checksum.
func New() hash.Hash

// Sum returns the MD5 checksum of the data.
func Sum(data []byte) [Size]byte
