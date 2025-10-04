// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

// New224 returns a new Digest computing the SHA3-224 hash.
func New224() *Digest

// New256 returns a new Digest computing the SHA3-256 hash.
func New256() *Digest

// New384 returns a new Digest computing the SHA3-384 hash.
func New384() *Digest

// New512 returns a new Digest computing the SHA3-512 hash.
func New512() *Digest

// NewLegacyKeccak256 returns a new Digest computing the legacy, non-standard
// Keccak-256 hash.
func NewLegacyKeccak256() *Digest

// NewLegacyKeccak512 returns a new Digest computing the legacy, non-standard
// Keccak-512 hash.
func NewLegacyKeccak512() *Digest
