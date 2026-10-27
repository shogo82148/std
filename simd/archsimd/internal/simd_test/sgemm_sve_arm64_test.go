// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

// SGEMM (C += A*B on float32) written with the scalable SVE API, checked
// against a scalar reference and benchmarked against the hand-written
// assembly kernel in sgemm_sve_arm64.s.

package simd
