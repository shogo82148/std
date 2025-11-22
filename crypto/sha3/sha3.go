// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha3 implements the SHA-3 hash algorithms and the SHAKE extendable
// output functions defined in FIPS 202.
package sha3

import (
	"github.com/shogo82148/std/crypto/internal/fips140/sha3"
	"github.com/shogo82148/std/hash"
)

// Sum224 returns the SHA3-224 hash of data.
func Sum224(data []byte) [28]byte

// Sum256 returns the SHA3-256 hash of data.
func Sum256(data []byte) [32]byte

// Sum384 returns the SHA3-384 hash of data.
func Sum384(data []byte) [48]byte

// Sum512 returns the SHA3-512 hash of data.
func Sum512(data []byte) [64]byte

// SumSHAKE128 applies the SHAKE128 extendable output function to data and
// returns an output of the given length in bytes.
func SumSHAKE128(data []byte, length int) []byte

// SumSHAKE256 applies the SHAKE256 extendable output function to data and
// returns an output of the given length in bytes.
func SumSHAKE256(data []byte, length int) []byte

// SHA3 is an instance of a SHA-3 hash. It implements [hash.Hash].
// The zero value is a usable SHA3-256 hash.
type SHA3 struct {
	s sha3.Digest
}

// New224 creates a new SHA3-224 hash.
func New224() *SHA3

// New256 creates a new SHA3-256 hash.
func New256() *SHA3

// New384 creates a new SHA3-384 hash.
func New384() *SHA3

// New512 creates a new SHA3-512 hash.
func New512() *SHA3

// Write absorbs more data into the hash's state.
func (s *SHA3) Write(p []byte) (n int, err error)

// Sum appends the current hash to b and returns the resulting slice.
func (s *SHA3) Sum(b []byte) []byte

// Reset resets the hash to its initial state.
func (s *SHA3) Reset()

// Size returns the number of bytes Sum will produce.
func (s *SHA3) Size() int

// BlockSize returns the hash's rate.
func (s *SHA3) BlockSize() int

// MarshalBinary implements [encoding.BinaryMarshaler].
func (s *SHA3) MarshalBinary() ([]byte, error)

// AppendBinary implements [encoding.BinaryAppender].
func (s *SHA3) AppendBinary(p []byte) ([]byte, error)

// UnmarshalBinary implements [encoding.BinaryUnmarshaler].
func (s *SHA3) UnmarshalBinary(data []byte) error

// Clone implements [hash.Cloner].
func (d *SHA3) Clone() (hash.Cloner, error)

// SHAKE is an instance of a SHAKE extendable output function.
// The zero value is a usable SHAKE256 hash.
type SHAKE struct {
	s sha3.SHAKE
}

// NewSHAKE128 creates a new SHAKE128 XOF.
func NewSHAKE128() *SHAKE

// NewSHAKE256 creates a new SHAKE256 XOF.
func NewSHAKE256() *SHAKE

// NewCSHAKE128 creates a new cSHAKE128 XOF.
//
// N is used to define functions based on cSHAKE, it can be empty when plain
// cSHAKE is desired. S is a customization byte string used for domain
// separation. When N and S are both empty, this is equivalent to NewSHAKE128.
func NewCSHAKE128(N, S []byte) *SHAKE

// NewCSHAKE256 creates a new cSHAKE256 XOF.
//
// N is used to define functions based on cSHAKE, it can be empty when plain
// cSHAKE is desired. S is a customization byte string used for domain
// separation. When N and S are both empty, this is equivalent to NewSHAKE256.
func NewCSHAKE256(N, S []byte) *SHAKE

// Write absorbs more data into the XOF's state.
//
// It panics if any output has already been read.
func (s *SHAKE) Write(p []byte) (n int, err error)

// Read squeezes more output from the XOF.
//
// Any call to Write after a call to Read will panic.
func (s *SHAKE) Read(p []byte) (n int, err error)

// Reset resets the XOF to its initial state.
func (s *SHAKE) Reset()

// BlockSize returns the rate of the XOF.
func (s *SHAKE) BlockSize() int

// MarshalBinary implements [encoding.BinaryMarshaler].
func (s *SHAKE) MarshalBinary() ([]byte, error)

// AppendBinary implements [encoding.BinaryAppender].
func (s *SHAKE) AppendBinary(p []byte) ([]byte, error)

// UnmarshalBinary implements [encoding.BinaryUnmarshaler].
func (s *SHAKE) UnmarshalBinary(data []byte) error
