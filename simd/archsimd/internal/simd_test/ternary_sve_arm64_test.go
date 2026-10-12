// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

// SVE ternary-op tests, the three-input counterpart of the drivers in
// binary_sve_arm64_test.go. flakiness absorbs the float32 double rounding of
// the math.FMA-based emulation, as in the NEON MulAdd tests.

package simd_test
