// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package archsimd

// LoadUint32x4Part loads a Int32x4 from the slice s.
// If s has fewer than 4 elements, the remaining elements of the vector are filled with zeroes.
// If s has 4 or more elements, the function is equivalent to LoadUint32x4.
func LoadUint32x4Part(s []uint32) (z Uint32x4, n int)

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 4 or more elements, the method is equivalent to x.Store.
func (x Uint32x4) StorePart(s []uint32) (n int)

// LoadUint64x2Part loads a Int64x2 from the slice s.
// If s has fewer than 2 elements, the remaining elements of the vector are filled with zeroes.
// If s has 2 or more elements, the function is equivalent to LoadUint64x2.
func LoadUint64x2Part(s []uint64) (z Uint64x2, n int)

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 2 or more elements, the method is equivalent to x.Store.
func (x Uint64x2) StorePart(s []uint64) (n int)

// LoadInt32x4Part loads a Int32x4 from the slice s.
// If s has fewer than 4 elements, the remaining elements of the vector are filled with zeroes.
// If s has 4 or more elements, the function is equivalent to LoadInt32x4.
func LoadInt32x4Part(s []int32) (z Int32x4, n int)

// StorePart stores the 4 elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 4 or more elements, the method is equivalent to x.Store.
func (x Int32x4) StorePart(s []int32) (n int)

// LoadInt64x2Part loads a Int64x2 from the slice s.
// If s has fewer than 2 elements, the remaining elements of the vector are filled with zeroes.
// If s has 2 or more elements, the function is equivalent to LoadInt64x2.
func LoadInt64x2Part(s []int64) (z Int64x2, n int)

// StorePart stores the 2 elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 2 or more elements, the method is equivalent to x.Store.
func (x Int64x2) StorePart(s []int64) (n int)

// LoadFloat32x4Part loads a Float32x4 from the slice s.
// If s has fewer than 4 elements, the remaining elements of the vector are filled with zeroes.
// If s has 4 or more elements, the function is equivalent to LoadFloat32x4.
func LoadFloat32x4Part(s []float32) (z Float32x4, n int)

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 4 or more elements, the method is equivalent to x.Store.
func (x Float32x4) StorePart(s []float32) (n int)

// LoadFloat64x2Part loads a Float64x2 from the slice s.
// If s has fewer than 2 elements, the remaining elements of the vector are filled with zeroes.
// If s has 2 or more elements, the function is equivalent to LoadFloat64x2.
func LoadFloat64x2Part(s []float64) (z Float64x2, n int)

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 2 or more elements, the method is equivalent to x.Store.
func (x Float64x2) StorePart(s []float64) (n int)
