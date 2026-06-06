// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package archsimd

// LoadInt8x32Part loads a Int8x32 from the slice s.
// If s has fewer than 32 elements, the remaining elements of the vector are filled with zeroes.
// If s has 32 or more elements, the function is equivalent to LoadInt8x32Slice.
func LoadInt8x32Part(s []int8) (Int8x32, int)

// LoadInt16x16Part loads a Int16x16 from the slice s.
// If s has fewer than 16 elements, the remaining elements of the vector are filled with zeroes.
// If s has 16 or more elements, the function is equivalent to LoadInt16x16Slice.
func LoadInt16x16Part(s []int16) (Int16x16, int)

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 32 or more elements, the method is equivalent to x.StoreSlice.
func (x Int8x32) StorePart(s []int8) int

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 16 or more elements, the method is equivalent to x.StoreSlice.
func (x Int16x16) StorePart(s []int16) int
