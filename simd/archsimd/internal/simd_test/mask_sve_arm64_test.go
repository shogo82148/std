// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && arm64

// SVE mask accessor tests. A mask is built from a uint64 of predicate bits
// through its memory representation (see maskBits), and each accessor is
// checked against a lane-by-lane reference over a set of lane patterns.

package simd_test
