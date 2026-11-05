// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package archsimd

// LoadUint32x4Part loads n=min(len(s), 4) elements from slice s as a
// vector and returns the vector and n. If len(s) < 4, the
// remaining vector elements will be zero.
func LoadUint32x4Part(s []uint32) (z Uint32x4, n int)

// StorePart stores n=min(len(s), 4) elements of x into s and returns n.
func (x Uint32x4) StorePart(s []uint32) (n int)

// LoadUint64x2Part loads n=min(len(s), 2) elements from slice s as a
// vector and returns the vector and n. If len(s) < 2, the
// remaining vector elements will be zero.
func LoadUint64x2Part(s []uint64) (z Uint64x2, n int)

// StorePart stores n=min(len(s), 2) elements of x into s and returns n.
func (x Uint64x2) StorePart(s []uint64) (n int)

// LoadInt32x4Part loads n=min(len(s), 4) elements from slice s as a
// vector and returns the vector and n. If len(s) < 4, the
// remaining vector elements will be zero.
func LoadInt32x4Part(s []int32) (z Int32x4, n int)

// StorePart stores n=min(len(s), 4) elements of x into s and returns n.
func (x Int32x4) StorePart(s []int32) (n int)

// LoadInt64x2Part loads n=min(len(s), 2) elements from slice s as a
// vector and returns the vector and n. If len(s) < 2, the
// remaining vector elements will be zero.
func LoadInt64x2Part(s []int64) (z Int64x2, n int)

// StorePart stores n=min(len(s), 2) elements of x into s and returns n.
func (x Int64x2) StorePart(s []int64) (n int)

// LoadFloat32x4Part loads n=min(len(s), 4) elements from slice s as a
// vector and returns the vector and n. If len(s) < 4, the
// remaining vector elements will be zero.
func LoadFloat32x4Part(s []float32) (z Float32x4, n int)

// StorePart stores n=min(len(s), 4) elements of x into s and returns n.
func (x Float32x4) StorePart(s []float32) (n int)

// LoadFloat64x2Part loads n=min(len(s), 2) elements from slice s as a
// vector and returns the vector and n. If len(s) < 2, the
// remaining vector elements will be zero.
func LoadFloat64x2Part(s []float64) (z Float64x2, n int)

// StorePart stores n=min(len(s), 2) elements of x into s and returns n.
func (x Float64x2) StorePart(s []float64) (n int)
