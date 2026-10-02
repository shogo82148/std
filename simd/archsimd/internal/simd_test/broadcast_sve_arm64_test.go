// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

// SVE broadcast tests, in the shape of the binary/unary drivers in
// binary_sve_arm64_test.go: broadcast each scalar drawn from the pool and
// check that every lane the hardware actually populated equals the scalar.

package simd_test
