// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func Map[T, U any](f func(T) U, in []T) []U

// File paths
const (
	SIMD = "simd/archsimd/"
	TD   = "simd/archsimd/internal/simd_test/"
	SSA  = "cmd/compile/internal/ssacompile/"
)
