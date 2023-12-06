// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

import "github.com/shogo82148/std/internal/chacha8rand"

// A ChaCha8 is a ChaCha8-based cryptographically strong
// random number generator.
type ChaCha8 struct {
	state chacha8rand.State
}

// NewChaCha8 returns a new ChaCha8 seeded with the given seed.
func NewChaCha8(seed [32]byte) *ChaCha8

// Seed resets the ChaCha8 to behave the same way as NewChaCha8(seed).
func (c *ChaCha8) Seed(seed [32]byte)

// Uint64 returns a uniformly distributed random uint64 value.
func (c *ChaCha8) Uint64() uint64

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (c *ChaCha8) UnmarshalBinary(data []byte) error

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (c *ChaCha8) MarshalBinary() ([]byte, error)
