// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && wasm

package archsimd

// Less return a mask vector of x[i] < y[i]
func (x Uint64x2) Less(y Uint64x2) Mask64x2

// LessEqual return a mask vector of x[i] <= y[i]
func (x Uint64x2) LessEqual(y Uint64x2) Mask64x2

// Greater return a mask vector of x[i] > y[i]
func (x Uint64x2) Greater(y Uint64x2) Mask64x2

// GreaterEqual return a mask vector of x[i] >= y[i]
func (x Uint64x2) GreaterEqual(y Uint64x2) Mask64x2

// Max returns the elementswise maximum of elements in x and y
func (x Int64x2) Max(y Int64x2) Int64x2

// Min returns the elementswise minimum of elements in x and y
func (x Int64x2) Min(y Int64x2) Int64x2

// Max returns the elementswise maximum of elements in x and y
func (x Uint64x2) Max(y Uint64x2) Uint64x2

// Min returns the elementswise minimum of elements in x and y
func (x Uint64x2) Min(y Uint64x2) Uint64x2

// Mul returns the elementswise product of elements in x and y
func (x Int8x16) Mul(y Int8x16) Int8x16

// Mul returns the elementswise product of elements in x and y
func (x Uint8x16) Mul(y Uint8x16) Uint8x16

// OnesCount returns the number of set bits in each vector element
func (x Int16x8) OnesCount() Int16x8

// OnesCount returns the number of set bits in each vector element
func (x Int32x4) OnesCount() Int32x4

// OnesCount returns the number of set bits in each vector element
func (x Int64x2) OnesCount() Int64x2

// OnesCount returns the number of set bits in each vector element
func (x Uint8x16) OnesCount() Uint8x16

// OnesCount returns the number of set bits in each vector element
func (x Uint16x8) OnesCount() Uint16x8

// OnesCount returns the number of set bits in each vector element
func (x Uint32x4) OnesCount() Uint32x4

// OnesCount returns the number of set bits in each vector element
func (x Uint64x2) OnesCount() Uint64x2

// CarrylessMultiplyEven computes the carryless
// multiplications of selected even halves of the elements of x and y.
//
// A carryless multiplication uses bitwise XOR instead of
// add-with-carry, for example (in base two):
//
//	11 * 11 = 11 * (10 ^ 1) = (11 * 10) ^ (11 * 1) = 110 ^ 11 = 101
//
// This also models multiplication of polynomials with coefficients
// from GF(2) -- 11 * 11 models (x+1)*(x+1) = x**2 + (1^1)x + 1 =
// x**2 + 0x + 1 = x**2 + 1 modeled by 101.  (Note that "+" adds
// polynomial terms, but coefficients "add" with XOR.)
//
// Emulated
func (x Uint64x2) CarrylessMultiplyEven(y Uint64x2) Uint64x2

// CarrylessMultiplyOdd computes the carryless
// multiplications of selected odd halves of the elements of x and y.
//
// A carryless multiplication uses bitwise XOR instead of
// add-with-carry, for example (in base two):
//
//	11 * 11 = 11 * (10 ^ 1) = (11 * 10) ^ (11 * 1) = 110 ^ 11 = 101
//
// This also models multiplication of polynomials with coefficients
// from GF(2) -- 11 * 11 models (x+1)*(x+1) = x**2 + (1^1)x + 1 =
// x**2 + 0x + 1 = x**2 + 1 modeled by 101.  (Note that "+" adds
// polynomial terms, but coefficients "add" with XOR.)
//
// Emulated
func (x Uint64x2) CarrylessMultiplyOdd(y Uint64x2) Uint64x2
