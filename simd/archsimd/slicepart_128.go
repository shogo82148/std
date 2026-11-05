// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && (amd64 || arm64 || wasm)

package archsimd

// LoadUint8x16Part loads n=min(len(s), 16) elements from slice s as a
// vector and returns the vector and n. If len(s) < 16, the
// remaining vector elements will be zero.
func LoadUint8x16Part(s []uint8) (z Uint8x16, n int)

// StorePart stores n=min(len(s), 16) elements of x into s and returns n.
func (x Uint8x16) StorePart(s []uint8) (n int)

// LoadUint16x8Part loads n=min(len(s), 8) elements from slice s as a
// vector and returns the vector and n. If len(s) < 8, the
// remaining vector elements will be zero.
func LoadUint16x8Part(s []uint16) (z Uint16x8, n int)

// StorePart stores n=min(len(s), 8) elements of x into s and returns n.
func (x Uint16x8) StorePart(s []uint16) (n int)

// LoadInt8x16Part loads n=min(len(s), 16) elements from slice s as a
// vector and returns the vector and n. If len(s) < 16, the
// remaining vector elements will be zero.
func LoadInt8x16Part(s []int8) (z Int8x16, n int)

// StorePart stores n=min(len(s), 16) elements of x into s and returns n.
func (x Int8x16) StorePart(s []int8) (n int)

// LoadInt16x8Part loads n=min(len(s), 8) elements from slice s as a
// vector and returns the vector and n. If len(s) < 8, the
// remaining vector elements will be zero.
func LoadInt16x8Part(s []int16) (z Int16x8, n int)

// StorePart stores n=min(len(s), 8) elements of x into s and returns n.
func (x Int16x8) StorePart(s []int16) (n int)
