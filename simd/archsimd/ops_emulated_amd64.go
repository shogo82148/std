// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simdstable && amd64

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
// As in Go, if shift is greater than or equal to the element width, the
// result is 0 or -1, depending on the sign of the element.
//
//	z[i] = x[i] >> shift
//
// Emulated, CPU Feature: AVX
func (x Int64x2) ShiftAllRight(shift uint64) (z Int64x2)

// ShiftAllRight arithmetically shifts each element of x right by shift bits.
// As in Go, if shift is greater than or equal to the element width, the
// result is 0 or -1, depending on the sign of the element.
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
// Asm: VPOPCNTB if CPU feature AVX512BITALG, Emulated otherwise, CPU Feature: AVX
func (x Int8x16) OnesCount() (z Int8x16)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Asm: VPOPCNTB if CPU feature AVX512BITALG, Emulated otherwise, CPU Feature: AVX
func (x Uint8x16) OnesCount() (z Uint8x16)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Asm: VPOPCNTB if CPU feature AVX512BITALG, Emulated otherwise, CPU Feature: AVX2
func (x Int8x32) OnesCount() (z Int8x32)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Asm: VPOPCNTB if CPU feature AVX512BITALG, Emulated otherwise, CPU Feature: AVX2
func (x Uint8x32) OnesCount() (z Uint8x32)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Asm: VPOPCNTB if CPU feature AVX512BITALG, Emulated otherwise, CPU Feature: AVX512
func (x Int8x64) OnesCount() (z Int8x64)

// OnesCount counts the number of one bits ("population count") in
// each element.
//
//	z[i] = bits.OnesCount(x[i])
//
// Asm: VPOPCNTB if CPU feature AVX512BITALG, Emulated otherwise, CPU Feature: AVX512
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

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint16x8) ReduceMin() (z uint16)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint16x8) ReduceMax() (z uint16)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int16x8) ReduceMin() (z int16)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int16x8) ReduceMax() (z int16)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint16x16) ReduceMin() (z uint16)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint16x16) ReduceMax() (z uint16)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int16x16) ReduceMin() (z int16)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int16x16) ReduceMax() (z int16)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint16x32) ReduceMin() (z uint16)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint16x32) ReduceMax() (z uint16)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int16x32) ReduceMin() (z int16)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int16x32) ReduceMax() (z int16)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint8x16) ReduceMin() (z uint8)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint8x16) ReduceMax() (z uint8)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int8x16) ReduceMin() (z int8)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int8x16) ReduceMax() (z int8)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) ReduceMin() (z uint8)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) ReduceMax() (z uint8)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) ReduceMin() (z int8)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) ReduceMax() (z int8)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint8x64) ReduceMin() (z uint8)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint8x64) ReduceMax() (z uint8)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int8x64) ReduceMin() (z int8)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int8x64) ReduceMax() (z int8)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int32x4) ReduceMax() (z int32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int32x4) ReduceMin() (z int32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint32x4) ReduceMax() (z uint32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint32x4) ReduceMin() (z uint32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int32x8) ReduceMax() (z int32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int32x8) ReduceMin() (z int32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint32x8) ReduceMax() (z uint32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint32x8) ReduceMin() (z uint32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int32x16) ReduceMax() (z int32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int32x16) ReduceMin() (z int32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint32x16) ReduceMax() (z uint32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint32x16) ReduceMin() (z uint32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int64x2) ReduceMax() (z int64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Int64x2) ReduceMin() (z int64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint64x2) ReduceMax() (z uint64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX
func (x Uint64x2) ReduceMin() (z uint64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int64x4) ReduceMax() (z int64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Int64x4) ReduceMin() (z int64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint64x4) ReduceMax() (z uint64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX2
func (x Uint64x4) ReduceMin() (z uint64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int64x8) ReduceMax() (z int64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Int64x8) ReduceMin() (z int64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint64x8) ReduceMax() (z uint64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// Emulated, CPU Feature: AVX512
func (x Uint64x8) ReduceMin() (z uint64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's max built-in and the IEEE 754-2019 maximum operation. Note
// that on AMD64 (and only AMD64), it differs from [Max].
//
// Emulated, CPU Feature: AVX
func (x Float32x4) ReduceMax() (z float32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's min built-in and the IEEE 754-2019 minimum operation. Note
// that on AMD64 (and only AMD64), it differs from [Min].
//
// Emulated, CPU Feature: AVX
func (x Float32x4) ReduceMin() (z float32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's max built-in and the IEEE 754-2019 maximum operation. Note
// that on AMD64 (and only AMD64), it differs from [Max].
//
// Emulated, CPU Feature: AVX
func (x Float32x8) ReduceMax() (z float32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's min built-in and the IEEE 754-2019 minimum operation. Note
// that on AMD64 (and only AMD64), it differs from [Min].
//
// Emulated, CPU Feature: AVX
func (x Float32x8) ReduceMin() (z float32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's max built-in and the IEEE 754-2019 maximum operation. Note
// that on AMD64 (and only AMD64), it differs from [Max].
//
// Emulated, CPU Feature: AVX512
func (x Float32x16) ReduceMax() (z float32)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's min built-in and the IEEE 754-2019 minimum operation. Note
// that on AMD64 (and only AMD64), it differs from [Min].
//
// Emulated, CPU Feature: AVX512
func (x Float32x16) ReduceMin() (z float32)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's max built-in and the IEEE 754-2019 maximum operation. Note
// that on AMD64 (and only AMD64), it differs from [Max].
//
// Emulated, CPU Feature: AVX
func (x Float64x2) ReduceMax() (z float64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's min built-in and the IEEE 754-2019 minimum operation. Note
// that on AMD64 (and only AMD64), it differs from [Min].
//
// Emulated, CPU Feature: AVX
func (x Float64x2) ReduceMin() (z float64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's max built-in and the IEEE 754-2019 maximum operation. Note
// that on AMD64 (and only AMD64), it differs from [Max].
//
// Emulated, CPU Feature: AVX
func (x Float64x4) ReduceMax() (z float64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's min built-in and the IEEE 754-2019 minimum operation. Note
// that on AMD64 (and only AMD64), it differs from [Min].
//
// Emulated, CPU Feature: AVX
func (x Float64x4) ReduceMin() (z float64)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's max built-in and the IEEE 754-2019 maximum operation. Note
// that on AMD64 (and only AMD64), it differs from [Max].
//
// Emulated, CPU Feature: AVX512
func (x Float64x8) ReduceMax() (z float64)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's min built-in and the IEEE 754-2019 minimum operation. Note
// that on AMD64 (and only AMD64), it differs from [Min].
//
// Emulated, CPU Feature: AVX512
func (x Float64x8) ReduceMin() (z float64)

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

// LookupOrZero permutes x, returning 0 for out-of-range indices.
//
//	if indices[i] < 16 {
//	    z[i] = x[indices[i]]
//	} else {
//	    z[i] = 0
//	}
//
// Emulated
func (x Uint8x16) LookupOrZero(indices Uint8x16) (z Uint8x16)

// LookupOrZero permutes x, returning 0 for out-of-range indices.
//
//	if indices[i] < 32 {
//	    z[i] = x[indices[i]]
//	} else {
//	    z[i] = 0
//	}
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) LookupOrZero(indices Uint8x32) (z Uint8x32)

// LookupOrZero permutes x, returning 0 for out-of-range indices.
//
//	if indices[i] < 64 {
//	    z[i] = x[indices[i]]
//	} else {
//	    z[i] = 0
//	}
//
// Emulated
func (x Uint8x64) LookupOrZero(indices Uint8x64) (z Uint8x64)
