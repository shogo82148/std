// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import (
	"github.com/shogo82148/std/internal/chacha8rand"
)

// A ChaCha8 is a ChaCha8-based cryptographically strong
// random number generator.
type ChaCha8 struct {
	state chacha8rand.State

	// The last readLen bytes of readBuf are still to be consumed by Read.
	readBuf [8]byte
	readLen int
}

// NewChaCha8 returns a new ChaCha8 seeded with the given seed.
func NewChaCha8(seed [32]byte) *ChaCha8

// Seed resets the ChaCha8 to behave the same way as NewChaCha8(seed).
func (c *ChaCha8) Seed(seed [32]byte)

// Uint64 returns a uniformly distributed random uint64 value.
func (c *ChaCha8) Uint64() uint64

// Read reads exactly len(p) bytes into p.
// It always returns len(p) and a nil error.
//
// If calls to Read and Uint64 are interleaved, the order in which bits are
// returned by the two is undefined, and Read may return bits generated before
// the last call to Uint64.
func (c *ChaCha8) Read(p []byte) (n int, err error)

// UnmarshalBinary implements the [encoding.BinaryUnmarshaler] interface.
func (c *ChaCha8) UnmarshalBinary(data []byte) error

// AppendBinary implements the [encoding.BinaryAppender] interface.
func (c *ChaCha8) AppendBinary(b []byte) ([]byte, error)

// MarshalBinary implements the [encoding.BinaryMarshaler] interface.
func (c *ChaCha8) MarshalBinary() ([]byte, error)
