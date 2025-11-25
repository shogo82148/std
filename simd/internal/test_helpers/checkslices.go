// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package test_helpers

import (
	"github.com/shogo82148/std/testing"
)

func CheckSlices[T number](t *testing.T, got, want []T) bool

// CheckSlices compares two slices for equality,
// reporting a test error if there is a problem,
// and also consumes the two slices so that a
// test/benchmark won't be dead-code eliminated.
func CheckSlicesLogInput[T number](t *testing.T, got, want []T, flakiness float64, logInput func()) bool
