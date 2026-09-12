// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package archsimd

// Abs returns the absolute values of the elements of x
//
// Emulated, CPU Feature AVX
func (x Float32x4) Abs() Float32x4

// Abs returns the absolute values of the elements of x
//
// Emulated, CPU Feature AVX2
func (x Float32x8) Abs() Float32x8

// Abs returns the absolute values of the elements of x
//
// Emulated, CPU Feature AVX512
func (x Float32x16) Abs() Float32x16

// Abs returns the absolute values of the elements of x
//
// Emulated, CPU Feature AVX
func (x Float64x2) Abs() Float64x2

// Abs returns the absolute values of the elements of x
//
// Emulated, CPU Feature AVX2
func (x Float64x4) Abs() Float64x4

// Abs returns the absolute values of the elements of x
//
// Emulated, CPU Feature AVX512
func (x Float64x8) Abs() Float64x8

// Neg returns the negation of the elements of x
//
// Emulated, CPU Feature AVX
func (x Float32x4) Neg() Float32x4

// Neg returns the negation of the elements of x
//
// Emulated, CPU Feature AVX2
func (x Float32x8) Neg() Float32x8

// Neg returns the negation of the elements of x
//
// Emulated, CPU Feature AVX512
func (x Float32x16) Neg() Float32x16

// Neg returns the negation of the elements of x
//
// Emulated, CPU Feature AVX
func (x Float64x2) Neg() Float64x2

// Neg returns the negation of the elements of x
//
// Emulated, CPU Feature AVX2
func (x Float64x4) Neg() Float64x4

// Neg returns the negation of the elements of x
//
// Emulated, CPU Feature AVX512
func (x Float64x8) Neg() Float64x8

// Mul multiplies corresponding elements of two vectors, modulo 2ⁿ.
//
// Emulated, CPU Feature: AVX
func (x Int8x16) Mul(y Int8x16) Int8x16

// Mul multiplies corresponding elements of two vectors, modulo 2ⁿ.
//
// Emulated, CPU Feature: AVX
func (x Uint8x16) Mul(y Uint8x16) Uint8x16

// Mul multiplies corresponding elements of two vectors, modulo 2ⁿ.
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) Mul(y Int8x32) Int8x32

// Mul multiplies corresponding elements of two vectors, modulo 2ⁿ.
//
// Emulated, CPU Feature: AVX512
func (x Int8x64) Mul(y Int8x64) Int8x64

// Mul multiplies corresponding elements of two vectors, modulo 2ⁿ.
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) Mul(y Uint8x32) Uint8x32

// Mul multiplies corresponding elements of two vectors, modulo 2ⁿ.
//
// Emulated, CPU Feature: AVX512
func (x Uint8x64) Mul(y Uint8x64) Uint8x64

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Float32x4) ReduceSum() float32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Float64x2) ReduceSum() float64

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Float32x8) ReduceSum() float32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Float64x4) ReduceSum() float64

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Float32x16) ReduceSum() float32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Float64x8) ReduceSum() float64

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Int16x8) ReduceSum() int16

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Uint16x8) ReduceSum() uint16

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Int32x4) ReduceSum() int32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Uint32x4) ReduceSum() uint32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX2
func (x Int16x16) ReduceSum() int16

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX2
func (x Uint16x16) ReduceSum() uint16

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX2
func (x Int32x8) ReduceSum() int32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX2
func (x Uint32x8) ReduceSum() uint32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Int16x32) ReduceSum() int16

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Uint16x32) ReduceSum() uint16

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Int32x16) ReduceSum() int32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Uint32x16) ReduceSum() uint32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Int8x16) ReduceSum() int8

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX
func (x Uint8x16) ReduceSum() uint8

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX2
func (x Int8x32) ReduceSum() int8

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX2
func (x Uint8x32) ReduceSum() uint8

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Int8x64) ReduceSum() int8

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: AVX512
func (x Uint8x64) ReduceSum() uint8
