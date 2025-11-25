// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package simd

// LoadInt8x32SlicePart loads a Int8x32 from the slice s.
// If s has fewer than 32 elements, the remaining elements of the vector are filled with zeroes.
// If s has 32 or more elements, the function is equivalent to LoadInt8x32Slice.
func LoadInt8x32SlicePart(s []int8) Int8x32

// LoadInt16x16SlicePart loads a Int16x16 from the slice s.
// If s has fewer than 16 elements, the remaining elements of the vector are filled with zeroes.
// If s has 16 or more elements, the function is equivalent to LoadInt16x16Slice.
func LoadInt16x16SlicePart(s []int16) Int16x16

// StoreSlicePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 32 or more elements, the method is equivalent to x.StoreSlice.
func (x Int8x32) StoreSlicePart(s []int8)

// StoreSlicePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 16 or more elements, the method is equivalent to x.StoreSlice.
func (x Int16x16) StoreSlicePart(s []int16)

// LoadInt8x16SlicePart loads a Int8x16 from the slice s.
// If s has fewer than 16 elements, the remaining elements of the vector are filled with zeroes.
// If s has 16 or more elements, the function is equivalent to LoadInt8x16Slice.
func LoadInt8x16SlicePart(s []int8) Int8x16

// StoreSlicePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 16 or more elements, the method is equivalent to x.StoreSlice.
func (x Int8x16) StoreSlicePart(s []int8)

// LoadInt16x8SlicePart loads a Int16x8 from the slice s.
// If s has fewer than 8 elements, the remaining elements of the vector are filled with zeroes.
// If s has 8 or more elements, the function is equivalent to LoadInt16x8Slice.
func LoadInt16x8SlicePart(s []int16) Int16x8

// StoreSlicePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 8 or more elements, the method is equivalent to x.StoreSlice.
func (x Int16x8) StoreSlicePart(s []int16)
