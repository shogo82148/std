// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha3

type SHAKE struct {
	d Digest

	// initBlock is the cSHAKE specific initialization set of bytes. It is initialized
	// by newCShake function and stores concatenation of N followed by S, encoded
	// by the method specified in 3.3 of [1].
	// It is stored here in order for Reset() to be able to put context into
	// initial state.
	initBlock []byte
}

func (s *SHAKE) BlockSize() int
func (s *SHAKE) Size() int

// Sum appends a portion of output to b and returns the resulting slice. The
// output length is selected to provide full-strength generic security: 32 bytes
// for SHAKE128 and 64 bytes for SHAKE256. It does not change the underlying
// state. It panics if any output has already been read.
func (s *SHAKE) Sum(in []byte) []byte

// Write absorbs more data into the hash's state.
// It panics if any output has already been read.
func (s *SHAKE) Write(p []byte) (n int, err error)

func (s *SHAKE) Read(out []byte) (n int, err error)

// Reset resets the hash to initial state.
func (s *SHAKE) Reset()

// Clone returns a copy of the SHAKE context in its current state.
func (s *SHAKE) Clone() *SHAKE

func (s *SHAKE) MarshalBinary() ([]byte, error)

func (s *SHAKE) AppendBinary(b []byte) ([]byte, error)

func (s *SHAKE) UnmarshalBinary(b []byte) error

// NewShake128 creates a new SHAKE128 XOF.
func NewShake128() *SHAKE

// NewShake256 creates a new SHAKE256 XOF.
func NewShake256() *SHAKE

// NewCShake128 creates a new cSHAKE128 XOF.
//
// N is used to define functions based on cSHAKE, it can be empty when plain
// cSHAKE is desired. S is a customization byte string used for domain
// separation. When N and S are both empty, this is equivalent to NewShake128.
func NewCShake128(N, S []byte) *SHAKE

// NewCShake256 creates a new cSHAKE256 XOF.
//
// N is used to define functions based on cSHAKE, it can be empty when plain
// cSHAKE is desired. S is a customization byte string used for domain
// separation. When N and S are both empty, this is equivalent to NewShake256.
func NewCShake256(N, S []byte) *SHAKE
