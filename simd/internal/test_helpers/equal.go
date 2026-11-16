// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package test_helpers

import (
	"github.com/shogo82148/std/testing"
)

// EqVector is a vector type T that can be compared element-wise with another T
// to produce a mask of type M.
type EqVector[T any, M EqMask] interface {
	comparable
	Equal(T) M
	Len() int
}

// EqMask is a mask type, returned by Vec.Equal(Vec).
type EqMask interface {
	comparable
	All() bool
}

// CheckVectorEqual checks == and != on vectors of type T against comparing
// their elements one at a time, for many pairs of vectors loaded from values.
//
// eq and ne must be x == y and x != y, written for the concrete type T, so that
// they test the compiler's direct lowering of ==. CheckVectorEqual also tests
// == in generic code, on structs and arrays containing vectors, on interfaces,
// through reflect, and on map keys.
func CheckVectorEqual[T EqVector[T, M], M EqMask, E number](t *testing.T, load func([]E) T, eq, ne func(x, y T) bool, values []E)

// CheckMaskEqual checks == and != on masks of type M against comparing their
// elements one at a time. It makes each mask by using load to make a vector of
// type T whose elements are zero where the mask should be true, and comparing
// that vector to zero.
//
// eq and ne must be x == y and x != y, written for the concrete type M, so that
// they test the compiler's direct lowering of ==. CheckMaskEqual also tests ==
// in generic code, on structs and arrays containing masks, on interfaces,
// through reflect, and on map keys.
func CheckMaskEqual[T EqVector[T, M], M EqMask, E number](t *testing.T, load func([]E) T, eq, ne func(x, y M) bool)
