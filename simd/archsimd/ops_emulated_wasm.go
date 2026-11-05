// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && wasm

package archsimd

// Less returns a mask indicating which elements of x are less than y.
//
//	z[i] = x[i] < y[i]
//
// Emulated
func (x Uint64x2) Less(y Uint64x2) (z Mask64x2)

// LessEqual returns a mask indicating which elements of x are less than or
// equal to y.
//
//	z[i] = x[i] <= y[i]
//
// Emulated
func (x Uint64x2) LessEqual(y Uint64x2) (z Mask64x2)

// Greater returns a mask indicating which elements of x are greater than y.
//
//	z[i] = x[i] > y[i]
//
// Emulated
func (x Uint64x2) Greater(y Uint64x2) (z Mask64x2)

// GreaterEqual returns a mask indicating which elements of x are greater than
// or equal to y.
//
//	z[i] = x[i] >= y[i]
//
// Emulated
func (x Uint64x2) GreaterEqual(y Uint64x2) (z Mask64x2)

// Max returns the elementwise maximum of x and y.
//
//	z[i] = max(x[i], y[i])
//
// Emulated
func (x Int64x2) Max(y Int64x2) (z Int64x2)

// Min returns the elementwise minimum of x and y.
//
//	z[i] = min(x[i], y[i])
//
// Emulated
func (x Int64x2) Min(y Int64x2) (z Int64x2)

// Max returns the elementwise maximum of x and y.
//
//	z[i] = max(x[i], y[i])
//
// Emulated
func (x Uint64x2) Max(y Uint64x2) (z Uint64x2)

// Min returns the elementwise minimum of x and y.
//
//	z[i] = min(x[i], y[i])
//
// Emulated
func (x Uint64x2) Min(y Uint64x2) (z Uint64x2)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated
func (x Int8x16) Mul(y Int8x16) (z Int8x16)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated
func (x Uint8x16) Mul(y Uint8x16) (z Uint8x16)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated
func (x Int16x8) OnesCount() (z Int16x8)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated
func (x Int32x4) OnesCount() (z Int32x4)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated
func (x Int64x2) OnesCount() (z Int64x2)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated
func (x Uint8x16) OnesCount() (z Uint8x16)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated
func (x Uint16x8) OnesCount() (z Uint16x8)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated
func (x Uint32x4) OnesCount() (z Uint32x4)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated
func (x Uint64x2) OnesCount() (z Uint64x2)

// MulAdd returns x * y + z elementwise.
//
//	w[i] = x[i] * y[i] + z[i]
//
// This performs a fused multiply-add, computing x[i]*y[i] + z[i] with a single
// rounding rather than two and no intermediate overflow, except on wasm, where
// it rounds x[i]*y[i] before adding z[i].
//
// Emulated
func (x Float32x4) MulAdd(y Float32x4, z Float32x4) (w Float32x4)

// MulAdd returns x * y + z elementwise.
//
//	w[i] = x[i] * y[i] + z[i]
//
// This performs a fused multiply-add, computing x[i]*y[i] + z[i] with a single
// rounding rather than two and no intermediate overflow, except on wasm, where
// it rounds x[i]*y[i] before adding z[i].
//
// Emulated
func (x Float64x2) MulAdd(y Float64x2, z Float64x2) (w Float64x2)

// CarrylessMultiplyEven computes the elementwise carryless multiplication of
// even-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i], y[2*i])
//
// A carryless multiplication uses bitwise XOR instead of add-with-carry.
// For example, to compute the carryless multiply of 0b1110 and 0b1011:
//
//	     1110
//	   ⊗ 1011
//	  ───────
//	     1110
//	    1110
//	   0000
//	⊕ 1110
//	  ───────
//	  1100010
//
// Carryless multiply can also be viewed as multiplying polynomials with
// coefficients from GF(2). For example, the above example can be represented as
//
//	  (x^3 + x^2 + x^1) * (x^3 + x^1 + x^0)
//	= (x^6 + x^5 + (1^1)x^4 + (1^1)x^3 + (1^1)x^2 + x^1)
//	= (x^6 + x^5 + x^1)
//
// Emulated
func (x Uint64x2) CarrylessMultiplyEven(y Uint64x2) (z Uint64x2)

// CarrylessMultiplyOdd computes the elementwise carryless multiplication of
// odd-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i+1], y[2*i+1])
//
// See [CarrylessMultiplyEven] for details about carryless multiply.
//
// Emulated
func (x Uint64x2) CarrylessMultiplyOdd(y Uint64x2) (z Uint64x2)

// ReduceSum returns the scalar sum of the elements of x.
//
// This uses repeated pairwise summation, which may affect the final precision
// of the result.
//
//	z = (x[0] + x[1]) + (x[2] + x[3])
//
// Emulated
func (x Float32x4) ReduceSum() (z float32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1]
//
// Emulated
func (x Float64x2) ReduceSum() (z float64)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated
func (x Int16x8) ReduceSum() (z int16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated
func (x Uint16x8) ReduceSum() (z uint16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated
func (x Int32x4) ReduceSum() (z int32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated
func (x Uint32x4) ReduceSum() (z uint32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated
func (x Int8x16) ReduceSum() (z int8)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated
func (x Uint8x16) ReduceSum() (z uint8)
