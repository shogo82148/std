// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && (arm64 || wasm)

package archsimd

// All returns true when all positions in mask x are true.
//
// Emulated
func (x Mask8x16) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated
func (x Mask8x16) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated
func (x Mask8x16) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated
func (x Mask16x8) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated
func (x Mask16x8) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated
func (x Mask16x8) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated
func (x Mask32x4) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated
func (x Mask32x4) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated
func (x Mask32x4) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated
func (x Mask64x2) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated
func (x Mask64x2) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated
func (x Mask64x2) None() bool

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated
func (m Mask8x16) TrailingZeros() int

// TrailingZeros returns the number of trailing (low-order) zeroes in mask m
//
// Emulated
func (m Mask16x8) TrailingZeros() int

// TrailingZeros returns the number of trailing (low-order) zeroes in mask m
//
// Emulated
func (m Mask32x4) TrailingZeros() int

// TrailingZeros returns the number of trailing (low-order) zeroes in mask m
//
// Emulated
func (m Mask64x2) TrailingZeros() int
