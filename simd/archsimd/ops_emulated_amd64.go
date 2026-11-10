// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package archsimd

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX
func (x Float32x4) Abs() (z Float32x4)

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX2
func (x Float32x8) Abs() (z Float32x8)

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX512
func (x Float32x16) Abs() (z Float32x16)

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX
func (x Float64x2) Abs() (z Float64x2)

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX2
func (x Float64x4) Abs() (z Float64x4)

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX512
func (x Float64x8) Abs() (z Float64x8)

// Neg returns the elementwise negation of x.
//
//	z[i] = -x[i]
//
// Emulated, CPU Feature: AVX
func (x Float32x4) Neg() (z Float32x4)

// Neg returns the elementwise negation of x.
//
//	z[i] = -x[i]
//
// Emulated, CPU Feature: AVX2
func (x Float32x8) Neg() (z Float32x8)

// Neg returns the elementwise negation of x.
//
//	z[i] = -x[i]
//
// Emulated, CPU Feature: AVX512
func (x Float32x16) Neg() (z Float32x16)

// Neg returns the elementwise negation of x.
//
//	z[i] = -x[i]
//
// Emulated, CPU Feature: AVX
func (x Float64x2) Neg() (z Float64x2)

// Neg returns the elementwise negation of x.
//
//	z[i] = -x[i]
//
// Emulated, CPU Feature: AVX2
func (x Float64x4) Neg() (z Float64x4)

// Neg returns the elementwise negation of x.
//
//	z[i] = -x[i]
//
// Emulated, CPU Feature: AVX512
func (x Float64x8) Neg() (z Float64x8)

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX
func (x Int64x2) Abs() (z Uint64x2)

// Abs returns the elementwise absolute value of x.
//
// Emulated, CPU Feature: AVX2
func (x Int64x4) Abs() (z Uint64x4)

// ShiftAllRight arithmetically shifts each element of x right by shift bits.
// If shift is greater than the element width, the result is 0 or -1.
//
//	z[i] = x[i] >> shift
//
// Emulated, CPU Feature: AVX
func (x Int64x2) ShiftAllRight(shift uint64) (z Int64x2)

// ShiftAllRight arithmetically shifts each element of x right by shift bits.
// If shift is greater than the element width, the result is 0 or -1.
//
//	z[i] = x[i] >> shift
//
// Emulated, CPU Feature: AVX2
func (x Int64x4) ShiftAllRight(shift uint64) (z Int64x4)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated, CPU Feature: AVX
func (x Int8x16) Mul(y Int8x16) (z Int8x16)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated, CPU Feature: AVX
func (x Uint8x16) Mul(y Uint8x16) (z Uint8x16)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) Mul(y Int8x32) (z Int8x32)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated, CPU Feature: AVX512
func (x Int8x64) Mul(y Int8x64) (z Int8x64)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) Mul(y Uint8x32) (z Uint8x32)

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
// Emulated, CPU Feature: AVX512
func (x Uint8x64) Mul(y Uint8x64) (z Uint8x64)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated, CPU Feature: AVX
func (x Int8x16) OnesCount() (z Int8x16)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated, CPU Feature: AVX
func (x Uint8x16) OnesCount() (z Uint8x16)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) OnesCount() (z Int8x32)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) OnesCount() (z Uint8x32)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Asm: VPOPCNTB, CPU Feature: AVX512BITALG
func (x Int8x64) OnesCount() (z Int8x64)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Asm: VPOPCNTB, CPU Feature: AVX512BITALG
func (x Uint8x64) OnesCount() (z Uint8x64)

// ReduceSum returns the scalar sum of the elements of x.
//
// This uses repeated pairwise summation, which may affect the final precision
// of the result.
//
//	z = (x[0] + x[1]) + (x[2] + x[3])
//
// Emulated, CPU Feature: AVX
func (x Float32x4) ReduceSum() (z float32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1]
//
// Emulated, CPU Feature: AVX
func (x Float64x2) ReduceSum() (z float64)

// ReduceSum returns the scalar sum of the elements of x.
//
// This uses repeated pairwise summation, which may affect the final precision
// of the result.
//
//	z = ((x[0] + x[1]) + (x[2] + x[3])) + ...
//
// Emulated, CPU Feature: AVX
func (x Float32x8) ReduceSum() (z float32)

// ReduceSum returns the scalar sum of the elements of x.
//
// This uses repeated pairwise summation, which may affect the final precision
// of the result.
//
//	z = (x[0] + x[1]) + (x[2] + x[3])
//
// Emulated, CPU Feature: AVX
func (x Float64x4) ReduceSum() (z float64)

// ReduceSum returns the scalar sum of the elements of x.
//
// This uses repeated pairwise summation, which may affect the final precision
// of the result.
//
//	z = ((x[0] + x[1]) + (x[2] + x[3])) + ...
//
// Emulated, CPU Feature: AVX512
func (x Float32x16) ReduceSum() (z float32)

// ReduceSum returns the scalar sum of the elements of x.
//
// This uses repeated pairwise summation, which may affect the final precision
// of the result.
//
//	z = ((x[0] + x[1]) + (x[2] + x[3])) + ...
//
// Emulated, CPU Feature: AVX512
func (x Float64x8) ReduceSum() (z float64)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX
func (x Int16x8) ReduceSum() (z int16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX
func (x Uint16x8) ReduceSum() (z uint16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX
func (x Int32x4) ReduceSum() (z int32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX
func (x Uint32x4) ReduceSum() (z uint32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX2
func (x Int16x16) ReduceSum() (z int16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX2
func (x Uint16x16) ReduceSum() (z uint16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX2
func (x Int32x8) ReduceSum() (z int32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX2
func (x Uint32x8) ReduceSum() (z uint32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX512
func (x Int16x32) ReduceSum() (z int16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX512
func (x Uint16x32) ReduceSum() (z uint16)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX512
func (x Int32x16) ReduceSum() (z int32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX512
func (x Uint32x16) ReduceSum() (z uint32)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX
func (x Int8x16) ReduceSum() (z int8)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX
func (x Uint8x16) ReduceSum() (z uint8)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) ReduceSum() (z int8)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) ReduceSum() (z uint8)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX512
func (x Int8x64) ReduceSum() (z int8)

// ReduceSum returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
// Emulated, CPU Feature: AVX512
func (x Uint8x64) ReduceSum() (z uint8)

// LookupOrZero permutes x, returning 0 for out-of-range indices.
//
//	if indices[i] < 16 {
//	    z[i] = x[indices[i]]
//	} else {
//	    z[i] = 0
//	}
//
// Emulated, CPU Feature: AVX
func (x Int8x16) LookupOrZero(indices Uint8x16) (z Int8x16)

// LookupOrZero permutes x, returning 0 for out-of-range indices.
//
//	if indices[i] < 32 {
//	    z[i] = x[indices[i]]
//	} else {
//	    z[i] = 0
//	}
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) LookupOrZero(indices Uint8x32) (z Int8x32)

// LookupOrZero permutes x, returning 0 for out-of-range indices.
//
//	if indices[i] < 64 {
//	    z[i] = x[indices[i]]
//	} else {
//	    z[i] = 0
//	}
//
// Emulated, CPU Feature: AVX512
func (x Int8x64) LookupOrZero(indices Uint8x64) (z Int8x64)
