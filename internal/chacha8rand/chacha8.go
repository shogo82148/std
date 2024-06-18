// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package chacha8rand implements a pseudorandom generator
// based on ChaCha8. It is used by both runtime and math/rand/v2
// and must have minimal dependencies.
package chacha8rand

// A State holds the state for a single random generator.
// It must be used from one goroutine at a time.
// If used by multiple goroutines at a time, the goroutines
// may see the same random values, but the code will not
// crash or cause out-of-bounds memory accesses.
type State struct {
	buf  [32]uint64
	seed [4]uint64
	i    uint32
	n    uint32
	c    uint32
}

// Next returns the next random value, along with a boolean
// indicating whether one was available.
// If one is not available, the caller should call Refill
// and then repeat the call to Next.
//
// Next is //go:nosplit to allow its use in the runtime
// with per-m data without holding the per-m lock.
//
//go:nosplit
func (s *State) Next() (uint64, bool)

// Init seeds the State with the given seed value.
func (s *State) Init(seed [32]byte)

// Init64 seeds the state with the given seed value.
func (s *State) Init64(seed [4]uint64)

// Refill refills the state with more random values.
// After a call to Refill, an immediate call to Next will succeed
// (unless multiple goroutines are incorrectly sharing a state).
func (s *State) Refill()

// Reseed reseeds the state with new random values.
// After a call to Reseed, any previously returned random values
// have been erased from the memory of the state and cannot be
// recovered.
func (s *State) Reseed()

// Marshal marshals the state into a byte slice.
// Marshal and Unmarshal are functions, not methods,
// so that they will not be linked into the runtime
// when it uses the State struct, since the runtime
// does not need these.
func Marshal(s *State) []byte

// Unmarshal unmarshals the state from a byte slice.
func Unmarshal(s *State, data []byte) error
