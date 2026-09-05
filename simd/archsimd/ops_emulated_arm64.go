// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

package archsimd

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: NEON
func (x Float32x4) ReduceSum() float32

// ReduceSum returns the sum of all elements in x.
//
// Emulated, CPU Feature: NEON
func (x Float64x2) ReduceSum() float64
