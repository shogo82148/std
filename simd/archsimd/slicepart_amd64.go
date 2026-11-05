// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package archsimd

// LoadInt8x32Part loads n=min(len(s), 32) elements from slice s as a
// vector and returns the vector and n. If len(s) < 32, the
// remaining vector elements will be zero.
func LoadInt8x32Part(s []int8) (z Int8x32, n int)

// LoadInt16x16Part loads n=min(len(s), 16) elements from slice s as a
// vector and returns the vector and n. If len(s) < 16, the
// remaining vector elements will be zero.
func LoadInt16x16Part(s []int16) (z Int16x16, n int)

// StorePart stores n=min(len(s), 32) elements of x into s and returns n.
func (x Int8x32) StorePart(s []int8) (n int)

// StorePart stores n=min(len(s), 16) elements of x into s and returns n.
func (x Int16x16) StorePart(s []int16) (n int)
