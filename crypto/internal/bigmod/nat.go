// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bigmod

import (
	"github.com/shogo82148/std/math/big"
)

// Nat represents an arbitrary natural number
//
// Each Nat has an announced length, which is the number of limbs it has stored.
// Operations on this number are allowed to leak this length, but will not leak
// any information about the values contained in those limbs.
type Nat struct {
	// limbs is little-endian in base 2^W with W = bits.UintSize.
	limbs []uint
}

// NewNat returns a new nat with a size of zero, just like new(Nat), but with
// the preallocated capacity to hold a number of up to preallocTarget bits.
// NewNat inlines, so the allocation can live on the stack.
func NewNat() *Nat

// Bytes returns x as a zero-extended big-endian byte slice. The size of the
// slice will match the size of m.
//
// x must have the same size as m and it must be reduced modulo m.
func (x *Nat) Bytes(m *Modulus) []byte

// SetBytes assigns x = b, where b is a slice of big-endian bytes.
// SetBytes returns an error if b >= m.
//
// The output will be resized to the size of m and overwritten.
func (x *Nat) SetBytes(b []byte, m *Modulus) (*Nat, error)

// SetOverflowingBytes assigns x = b, where b is a slice of big-endian bytes.
// SetOverflowingBytes returns an error if b has a longer bit length than m, but
// reduces overflowing values up to 2^⌈log2(m)⌉ - 1.
//
// The output will be resized to the size of m and overwritten.
func (x *Nat) SetOverflowingBytes(b []byte, m *Modulus) (*Nat, error)

// Equal returns 1 if x == y, and 0 otherwise.
//
// Both operands must have the same announced length.
func (x *Nat) Equal(y *Nat) choice

// IsZero returns 1 if x == 0, and 0 otherwise.
func (x *Nat) IsZero() choice

// Modulus is used for modular arithmetic, precomputing relevant constants.
//
// Moduli are assumed to be odd numbers. Moduli can also leak the exact
// number of bits needed to store their value, and are stored without padding.
//
// Their actual value is still kept secret.
type Modulus struct {
	// The underlying natural number for this modulus.
	//
	// This will be stored without any padding, and shouldn't alias with any
	// other natural number being used.
	nat     *Nat
	leading int
	m0inv   uint
	rr      *Nat
}

// NewModulusFromBig creates a new Modulus from a [big.Int].
//
// The Int must be odd. The number of significant bits (and nothing else) is
// leaked through timing side-channels.
func NewModulusFromBig(n *big.Int) (*Modulus, error)

// Size returns the size of m in bytes.
func (m *Modulus) Size() int

// BitLen returns the size of m in bits.
func (m *Modulus) BitLen() int

// Nat returns m as a Nat. The return value must not be written to.
func (m *Modulus) Nat() *Nat

// Mod calculates out = x mod m.
//
// This works regardless how large the value of x is.
//
// The output will be resized to the size of m and overwritten.
func (out *Nat) Mod(x *Nat, m *Modulus) *Nat

// ExpandFor ensures x has the right size to work with operations modulo m.
//
// The announced size of x must be smaller than or equal to that of m.
func (x *Nat) ExpandFor(m *Modulus) *Nat

// Sub computes x = x - y mod m.
//
// The length of both operands must be the same as the modulus. Both operands
// must already be reduced modulo m.
func (x *Nat) Sub(y *Nat, m *Modulus) *Nat

// Add computes x = x + y mod m.
//
// The length of both operands must be the same as the modulus. Both operands
// must already be reduced modulo m.
func (x *Nat) Add(y *Nat, m *Modulus) *Nat

// Mul calculates x = x * y mod m.
//
// The length of both operands must be the same as the modulus. Both operands
// must already be reduced modulo m.
func (x *Nat) Mul(y *Nat, m *Modulus) *Nat

// Exp calculates out = x^e mod m.
//
// The exponent e is represented in big-endian order. The output will be resized
// to the size of m and overwritten. x must already be reduced modulo m.
func (out *Nat) Exp(x *Nat, e []byte, m *Modulus) *Nat

// ExpShortVarTime calculates out = x^e mod m.
//
// The output will be resized to the size of m and overwritten. x must already
// be reduced modulo m. This leaks the exponent through timing side-channels.
func (out *Nat) ExpShortVarTime(x *Nat, e uint, m *Modulus) *Nat
