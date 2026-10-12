// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

// SVE mask (predicate) tests, in the same utility-function shape as
// compare_amd64_test.go.

package simd_test
