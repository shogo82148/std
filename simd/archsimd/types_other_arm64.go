// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package archsimd

// Int8s is an SVE vector of int8s.
type Int8s struct {
	int8s vsve
	vals  [32]int8
}

// Mask8s is an SVE predicate for 8-bit elements.
type Mask8s struct {
	mask8s psve
	vals   uint32
}

func (v *Int8s) Len() int

func (m *Mask8s) Len() int
