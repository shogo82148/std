// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate.go. DO NOT EDIT.

package fiat

// P256Element is an integer modulo 2^256 - 2^224 + 2^192 + 2^96 - 1.
//
// The zero value is a valid zero element.
type P256Element struct {
	// Values are represented internally always in the Montgomery domain, and
	// converted in Bytes and SetBytes.
	x p256MontgomeryDomainFieldElement
}

// One sets e = 1, and returns e.
func (e *P256Element) One() *P256Element

// Equal returns 1 if e == t, and zero otherwise.
func (e *P256Element) Equal(t *P256Element) int

// IsZero returns 1 if e == 0, and zero otherwise.
func (e *P256Element) IsZero() int

// Set sets e = t, and returns e.
func (e *P256Element) Set(t *P256Element) *P256Element

// Bytes returns the 32-byte big-endian encoding of e.
func (e *P256Element) Bytes() []byte

// SetBytes sets e = v, where v is a big-endian 32-byte encoding, and returns e.
// If v is not 32 bytes or it encodes a value higher than 2^256 - 2^224 + 2^192 + 2^96 - 1,
// SetBytes returns nil and an error, and e is unchanged.
func (e *P256Element) SetBytes(v []byte) (*P256Element, error)

// Add sets e = t1 + t2, and returns e.
func (e *P256Element) Add(t1, t2 *P256Element) *P256Element

// Sub sets e = t1 - t2, and returns e.
func (e *P256Element) Sub(t1, t2 *P256Element) *P256Element

// Mul sets e = t1 * t2, and returns e.
func (e *P256Element) Mul(t1, t2 *P256Element) *P256Element

// Square sets e = t * t, and returns e.
func (e *P256Element) Square(t *P256Element) *P256Element

// Select sets v to a if cond == 1, and to b if cond == 0.
func (v *P256Element) Select(a, b *P256Element, cond int) *P256Element
