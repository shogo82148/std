// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

// SVE shuffle tests. The expected result of a shuffle depends on the runtime
// vector length, so unlike testSVEBinary, these drivers pass only the active
// lanes to want.

package simd_test
