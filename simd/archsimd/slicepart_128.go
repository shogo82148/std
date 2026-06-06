// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && (amd64 || arm64 || wasm)

package archsimd

// LoadUint8x16Part loads a Uint8x16 from the slice s.
// If s has fewer than 16 elements, the remaining elements of the vector are filled with zeroes.
// If s has 16 or more elements, the function is equivalent to LoadInt8x16.
func LoadUint8x16Part(s []uint8) (Uint8x16, int)

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 16 or more elements, the method is equivalent to x.Store.
func (x Uint8x16) StorePart(s []uint8) int

// LoadUint16x8Part loads a Uint16x8 from the slice s.
// If s has fewer than 8 elements, the remaining elements of the vector are filled with zeroes.
// If s has 8 or more elements, the function is equivalent to LoadInt16x8.
func LoadUint16x8Part(s []uint16) (Uint16x8, int)

// StorePart stores the elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 8 or more elements, the method is equivalent to x.Store.
func (x Uint16x8) StorePart(s []uint16) int

// LoadInt8x16Part loads a Int8x16 from the slice s, it returns the loaded vector and the
// number of elements loaded.
// If s has fewer than 16 elements, the remaining elements of the vector are filled with zeroes.
// If s has 16 or more elements, the function is equivalent to LoadInt8x16.
func LoadInt8x16Part(s []int8) (Int8x16, int)

// StorePart stores the 16 elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 16 or more elements, the method is equivalent to x.Store.
func (x Int8x16) StorePart(s []int8) int

// LoadInt16x8Part loads a Int16x8 from the slice s, it returns the loaded vector and the
// number of elements loaded.
// If s has fewer than 8 elements, the remaining elements of the vector are filled with zeroes.
// If s has 8 or more elements, the function is equivalent to LoadInt16x8.
func LoadInt16x8Part(s []int16) (Int16x8, int)

// StorePart stores the 8 elements of x into the slice s.
// It stores as many elements as will fit in s.
// If s has 8 or more elements, the method is equivalent to x.Store.
func (x Int16x8) StorePart(s []int16) int
