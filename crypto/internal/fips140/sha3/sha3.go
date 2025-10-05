// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sha3 implements the SHA-3 fixed-output-length hash functions and
// the SHAKE variable-output-length functions defined by [FIPS 202], as well as
// the cSHAKE extendable-output-length functions defined by [SP 800-185].
//
// [FIPS 202]: https://doi.org/10.6028/NIST.FIPS.202
// [SP 800-185]: https://doi.org/10.6028/NIST.SP.800-185
package sha3

type Digest struct {
	a [1600 / 8]byte

	// a[n:rate] is the buffer. If absorbing, it's the remaining space to XOR
	// into before running the permutation. If squeezing, it's the remaining
	// output to produce before running the permutation.
	n, rate int

	// dsbyte contains the "domain separation" bits and the first bit of
	// the padding. Sections 6.1 and 6.2 of [1] separate the outputs of the
	// SHA-3 and SHAKE functions by appending bitstrings to the message.
	// Using a little-endian bit-ordering convention, these are "01" for SHA-3
	// and "1111" for SHAKE, or 00000010b and 00001111b, respectively. Then the
	// padding rule from section 5.1 is applied to pad the message to a multiple
	// of the rate, which involves adding a "1" bit, zero or more "0" bits, and
	// a final "1" bit. We merge the first "1" bit from the padding into dsbyte,
	// giving 00000110b (0x06) and 00011111b (0x1f).
	// [1] http://csrc.nist.gov/publications/drafts/fips-202/fips_202_draft.pdf
	//     "Draft FIPS 202: SHA-3 Standard: Permutation-Based Hash and
	//      Extendable-Output Functions (May 2014)"
	dsbyte byte

	outputLen int
	state     spongeDirection
}

// BlockSize returns the rate of sponge underlying this hash function.
func (d *Digest) BlockSize() int

// Size returns the output size of the hash function in bytes.
func (d *Digest) Size() int

// Reset resets the Digest to its initial state.
func (d *Digest) Reset()

func (d *Digest) Clone() *Digest

// Write absorbs more data into the hash's state.
func (d *Digest) Write(p []byte) (n int, err error)

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (d *Digest) Sum(b []byte) []byte

func (d *Digest) MarshalBinary() ([]byte, error)

func (d *Digest) AppendBinary(b []byte) ([]byte, error)

func (d *Digest) UnmarshalBinary(b []byte) error
