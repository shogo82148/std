// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha512 implements the SHA-384, SHA-512, SHA-512/224, and SHA-512/256
// hash algorithms as defined in FIPS 180-4.
package sha512

// Digest is a SHA-384, SHA-512, SHA-512/224, or SHA-512/256 [hash.Hash]
// implementation.
type Digest struct {
	h    [8]uint64
	x    [chunk]byte
	nx   int
	len  uint64
	size int
}

func (d *Digest) Reset()

func (d *Digest) MarshalBinary() ([]byte, error)

func (d *Digest) AppendBinary(b []byte) ([]byte, error)

func (d *Digest) UnmarshalBinary(b []byte) error

// New returns a new Digest computing the SHA-512 hash.
func New() *Digest

// New512_224 returns a new Digest computing the SHA-512/224 hash.
func New512_224() *Digest

// New512_256 returns a new Digest computing the SHA-512/256 hash.
func New512_256() *Digest

// New384 returns a new Digest computing the SHA-384 hash.
func New384() *Digest

func (d *Digest) Size() int

func (d *Digest) BlockSize() int

func (d *Digest) Write(p []byte) (nn int, err error)

func (d *Digest) Sum(in []byte) []byte
