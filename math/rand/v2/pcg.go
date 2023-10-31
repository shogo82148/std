// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand

// A PCG is a PCG generator with 128 bits of internal state.
// A zero PCG is equivalent to NewPCG(0, 0).
type PCG struct {
	hi uint64
	lo uint64
}

// NewPCG returns a new PCG seeded with the given values.
func NewPCG(seed1, seed2 uint64) *PCG

// Seed resets the PCG to behave the same way as NewPCG(seed1, seed2).
func (p *PCG) Seed(seed1, seed2 uint64)

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (p *PCG) MarshalBinary() ([]byte, error)

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (p *PCG) UnmarshalBinary(data []byte) error

// Uint64 return a uniformly-distributed random uint64 value.
func (p *PCG) Uint64() uint64
