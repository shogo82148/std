// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package test_helpers

import (
	"github.com/shogo82148/std/testing"
)

func CheckSlices[T number](t *testing.T, got, want []T) bool

// CheckSlicesLogInput compares two slices for equality,
// reporting a test error if there is a problem,
// and also consumes the two slices so that a
// test/benchmark won't be dead-code eliminated.
func CheckSlicesLogInput[T number](t *testing.T, got, want []T, flakiness float64, logInput func()) bool

// CheckScalarsLogInput compares two values for equality,
// reporting a test error if there is a problem,
// and also consumes the two values so that a
// test/benchmark won't be dead-code eliminated.
// The goal is to match the behavior and output for
// CheckSlicesLogInput.
func CheckScalarsLogInput[T number](t *testing.T, got, want T, flakiness float64, logInput func()) bool

// CheckStringsLogInput compares two strings for equality,
// reporting a test error if there is a problem,
// and also consumes the two values so that a
// test/benchmark won't be dead-code eliminated.
// The goal is to match the behavior and output for
// CheckSlicesLogInput.
func CheckStringsLogInput(t *testing.T, got, want string, flakiness float64, logInput func()) bool

const (
	PN22  = 1.0 / 1024 / 1024 / 4
	PN24  = 1.0 / 1024 / 1024 / 16
	PN53  = PN24 * PN24 / 32
	F0    = float32(1.0 + 513*PN22/2)
	F1    = float32(1.0 + 511*PN22*8)
	Aeasy = float32(2046 * PN53)
	Ahard = float32(2047 * PN53)
)

// N controls how large the test vectors are
const N = 144

func Float64s() []float64

func Int64s() []int64

func Uint64s() []uint64

func Float32s() []float32

func Int32s() []int32

func Uint32s() []uint32

func Int16s() []int16

func Uint16s() []uint16

func Int8s() []int8

func Uint8s() []uint8

func Bools() []bool

func Shift8s() []uint64

func Shift16s() []uint64

func Shift32s() []uint64

func Shift64s() []uint64

// ImpliedDo returns a slice T with length n and each element i
// initialized to val(i).
func ImpliedDo[T number](n int, val func(i int) T) []T

// ImpliedBools returns a string
func ImpliedBools(n int, val func(i int) bool) string
