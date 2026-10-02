// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package archsimd

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX
func (x Mask8x16) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX
func (x Mask8x16) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX
func (x Mask8x16) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX
func (x Mask16x8) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX
func (x Mask16x8) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX
func (x Mask16x8) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX
func (x Mask32x4) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX
func (x Mask32x4) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX
func (x Mask32x4) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX
func (x Mask64x2) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX
func (x Mask64x2) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX
func (x Mask64x2) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX2
func (x Mask8x32) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX2
func (x Mask8x32) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX2
func (x Mask8x32) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX2
func (x Mask16x16) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX2
func (x Mask16x16) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX2
func (x Mask16x16) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX2
func (x Mask32x8) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX2
func (x Mask32x8) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX2
func (x Mask32x8) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX2
func (x Mask64x4) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX2
func (x Mask64x4) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX2
func (x Mask64x4) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX512
func (x Mask8x64) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX512
func (x Mask8x64) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX512
func (x Mask8x64) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX512
func (x Mask16x32) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX512
func (x Mask16x32) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX512
func (x Mask16x32) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX512
func (x Mask32x16) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX512
func (x Mask32x16) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX512
func (x Mask32x16) None() bool

// All returns true when all positions in mask x are true.
//
// Emulated, CPU Feature AVX512
func (x Mask64x8) All() bool

// Any returns true when any position in mask x is true.
//
// Emulated, CPU Feature AVX512
func (x Mask64x8) Any() bool

// None returns true when no positions in mask x are set.
//
// Emulated, CPU Feature AVX512
func (x Mask64x8) None() bool

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX
func (m Mask8x16) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX
func (m Mask8x32) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX
func (m Mask8x64) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX
func (x Mask16x8) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX2
func (x Mask16x16) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX512
func (x Mask16x32) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX
func (x Mask32x4) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX2
func (x Mask32x8) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX512
func (x Mask32x16) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX
func (x Mask64x2) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX2
func (x Mask64x4) TrailingZeros() int

// TrailingZeros returns the number of low-order false (zero) elements in mask m.
//
// Emulated, CPU Feature AVX512
func (x Mask64x8) TrailingZeros() int
