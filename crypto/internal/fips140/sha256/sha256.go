// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha256 implements the SHA-224 and SHA-256 hash algorithms as defined
// in FIPS 180-4.
package sha256

import (
	"github.com/shogo82148/std/hash"
)

// Digest is a SHA-224 or SHA-256 [hash.Hash] implementation.
type Digest struct {
	h     [8]uint32
	x     [chunk]byte
	nx    int
	len   uint64
	is224 bool
}

func (d *Digest) MarshalBinary() ([]byte, error)

func (d *Digest) AppendBinary(b []byte) ([]byte, error)

func (d *Digest) UnmarshalBinary(b []byte) error

func (d *Digest) Clone() (hash.Cloner, error)

func (d *Digest) Reset()

// New returns a new Digest computing the SHA-256 hash.
func New() *Digest

// New224 returns a new Digest computing the SHA-224 hash.
func New224() *Digest

func (d *Digest) Size() int

func (d *Digest) BlockSize() int

func (d *Digest) Write(p []byte) (nn int, err error)

func (d *Digest) Sum(in []byte) []byte
