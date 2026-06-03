// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func Map[T, U any](f func(T) U, in []T) []U

const SIMD = "../../"
const TD = "../../internal/simd_test/"
const SSA = "../../../../cmd/compile/internal/ssa/"
