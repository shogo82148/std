// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fnv implements FNV-1 and FNV-1a, non-cryptographic hash functions
// created by Glenn Fowler, Landon Curt Noll, and Phong Vo.
// See
// http://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function.
package fnv

import (
	"github.com/shogo82148/std/hash"
)

// New32 returns a new 32-bit FNV-1 hash.Hash.
func New32() hash.Hash32

// New32a returns a new 32-bit FNV-1a hash.Hash.
func New32a() hash.Hash32

// New64 returns a new 64-bit FNV-1 hash.Hash.
func New64() hash.Hash64

// New64a returns a new 64-bit FNV-1a hash.Hash.
func New64a() hash.Hash64
