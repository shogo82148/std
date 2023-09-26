// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maphash provides hash functions on byte sequences.
// These hash functions are intended to be used to implement hash tables or
// other data structures that need to map arbitrary strings or byte
// sequences to a uniform distribution on unsigned 64-bit integers.
// Each different instance of a hash table or data structure should use its own Seed.
//
// The hash functions are not cryptographically secure.
// (See crypto/sha256 and crypto/sha512 for cryptographic use.)
package maphash

// A Seed is a random value that selects the specific hash function
// computed by a Hash. If two Hashes use the same Seeds, they
// will compute the same hash values for any given input.
// If two Hashes use different Seeds, they are very likely to compute
// distinct hash values for any given input.
//
// A Seed must be initialized by calling MakeSeed.
// The zero seed is uninitialized and not valid for use with Hash's SetSeed method.
//
// Each Seed value is local to a single process and cannot be serialized
// or otherwise recreated in a different process.
type Seed struct {
	s uint64
}

// A Hash computes a seeded hash of a byte sequence.
//
// The zero Hash is a valid Hash ready to use.
// A zero Hash chooses a random seed for itself during
// the first call to a Reset, Write, Seed, or Sum64 method.
// For control over the seed, use SetSeed.
//
// The computed hash values depend only on the initial seed and
// the sequence of bytes provided to the Hash object, not on the way
// in which the bytes are provided. For example, the three sequences
//
//	h.Write([]byte{'f','o','o'})
//	h.WriteByte('f'); h.WriteByte('o'); h.WriteByte('o')
//	h.WriteString("foo")
//
// all have the same effect.
//
// Hashes are intended to be collision-resistant, even for situations
// where an adversary controls the byte sequences being hashed.
//
// A Hash is not safe for concurrent use by multiple goroutines, but a Seed is.
// If multiple goroutines must compute the same seeded hash,
// each can declare its own Hash and call SetSeed with a common Seed.
type Hash struct {
	_     [0]func()
	seed  Seed
	state Seed
	buf   [bufSize]byte
	n     int
}

// bufSize is the size of the Hash write buffer.
// The buffer ensures that writes depend only on the sequence of bytes,
// not the sequence of WriteByte/Write/WriteString calls,
// by always calling rthash with a full buffer (except for the tail).

// WriteByte adds b to the sequence of bytes hashed by h.
// It never fails; the error result is for implementing io.ByteWriter.
func (h *Hash) WriteByte(b byte) error

// Write adds b to the sequence of bytes hashed by h.
// It always writes all of b and never fails; the count and error result are for implementing io.Writer.
func (h *Hash) Write(b []byte) (int, error)

// WriteString adds the bytes of s to the sequence of bytes hashed by h.
// It always writes all of s and never fails; the count and error result are for implementing io.StringWriter.
func (h *Hash) WriteString(s string) (int, error)

// Seed returns h's seed value.
func (h *Hash) Seed() Seed

// SetSeed sets h to use seed, which must have been returned by MakeSeed
// or by another Hash's Seed method.
// Two Hash objects with the same seed behave identically.
// Two Hash objects with different seeds will very likely behave differently.
// Any bytes added to h before this call will be discarded.
func (h *Hash) SetSeed(seed Seed)

// Reset discards all bytes added to h.
// (The seed remains the same.)
func (h *Hash) Reset()

// Sum64 returns h's current 64-bit value, which depends on
// h's seed and the sequence of bytes added to h since the
// last call to Reset or SetSeed.
//
// All bits of the Sum64 result are close to uniformly and
// independently distributed, so it can be safely reduced
// by using bit masking, shifting, or modular arithmetic.
func (h *Hash) Sum64() uint64

// MakeSeed returns a new random seed.
func MakeSeed() Seed

// Sum appends the hash's current 64-bit value to b.
// It exists for implementing hash.Hash.
// For direct calls, it is more efficient to use Sum64.
func (h *Hash) Sum(b []byte) []byte

// Size returns h's hash value size, 8 bytes.
func (h *Hash) Size() int

// BlockSize returns h's block size.
func (h *Hash) BlockSize() int
