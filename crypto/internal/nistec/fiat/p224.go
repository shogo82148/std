// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate.go. DO NOT EDIT.

package fiat

// P224Element is an integer modulo 2^224 - 2^96 + 1.
//
// The zero value is a valid zero element.
type P224Element struct {
	// Values are represented internally always in the Montgomery domain, and
	// converted in Bytes and SetBytes.
	x p224MontgomeryDomainFieldElement
}

// One sets e = 1, and returns e.
func (e *P224Element) One() *P224Element

// Equal returns 1 if e == t, and zero otherwise.
func (e *P224Element) Equal(t *P224Element) int

// IsZero returns 1 if e == 0, and zero otherwise.
func (e *P224Element) IsZero() int

// Set sets e = t, and returns e.
func (e *P224Element) Set(t *P224Element) *P224Element

// Bytes returns the 28-byte big-endian encoding of e.
func (e *P224Element) Bytes() []byte

// SetBytes sets e = v, where v is a big-endian 28-byte encoding, and returns e.
// If v is not 28 bytes or it encodes a value higher than 2^224 - 2^96 + 1,
// SetBytes returns nil and an error, and e is unchanged.
func (e *P224Element) SetBytes(v []byte) (*P224Element, error)

// Add sets e = t1 + t2, and returns e.
func (e *P224Element) Add(t1, t2 *P224Element) *P224Element

// Sub sets e = t1 - t2, and returns e.
func (e *P224Element) Sub(t1, t2 *P224Element) *P224Element

// Mul sets e = t1 * t2, and returns e.
func (e *P224Element) Mul(t1, t2 *P224Element) *P224Element

// Square sets e = t * t, and returns e.
func (e *P224Element) Square(t *P224Element) *P224Element

// Select sets v to a if cond == 1, and to b if cond == 0.
func (v *P224Element) Select(a, b *P224Element, cond int) *P224Element