// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Dim returns the maximum of x-y or 0.
//
// Special cases are:
//
//	Dim(+Inf, +Inf) = NaN
//	Dim(-Inf, -Inf) = NaN
//	Dim(x, NaN) = Dim(NaN, x) = NaN
func Dim(x, y float64) float64

// Max returns the larger of x or y.
//
// Special cases are:
//
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
//
// Note that this differs from the built-in function max when called
// with NaN and +Inf.
func Max(x, y float64) float64

// Min returns the smaller of x or y.
//
// Special cases are:
//
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
//
// Note that this differs from the built-in function min when called
// with NaN and -Inf.
func Min(x, y float64) float64
