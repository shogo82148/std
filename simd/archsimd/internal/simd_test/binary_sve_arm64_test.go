// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

// SVE binary-op tests. Unlike amd64, SVE has only a handful of (scalable)
// vector types, so there is nothing to generate — these drivers are hand-written
// in the same shape as the generated testXxxBinary helpers. Each loads two input
// windows via the fixed-array API, runs the op, stores the result, and compares
// the lanes the hardware actually populated: the vector's runtime Len() (VL is
// <= the 32-byte backing, enforced at package init).

package simd_test
