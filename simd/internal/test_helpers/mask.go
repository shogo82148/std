// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package test_helpers

import "github.com/shogo82148/std/testing"

func TestMaskAllAny[E integer, V interface {
	Len() int
	Equal(V) M
	NotEqual(V) M
}, M interface {
	All() bool
	Any() bool
	None() bool
	TrailingZeros() int
}](t *testing.T, load func([]E) V)
