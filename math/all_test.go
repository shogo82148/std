// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math_test

import (
	. "math"
)

// The expected results below were computed by the high precision calculators
// at https://keisan.casio.com/.  More exact input values (array vf[], above)
// were obtained by printing them with "%.26f".  The answers were calculated
// to 26 digits (by using the "Digit number" drop-down control of each
// calculator).

// Results for 100000 * Pi + vf[i]

// Results for 100000 * Pi + vf[i]

// Results for 100000 * Pi + vf[i]

// arguments and expected results for special cases

// arguments and expected results for boundary cases
const (
	SmallestNormalFloat64   = 2.2250738585072014e-308
	LargestSubnormalFloat64 = SmallestNormalFloat64 - SmallestNonzeroFloat64
)

// Global exported variables are used to store the
// return values of functions measured in the benchmarks.
// Storing the results in these variables prevents the compiler
// from completely optimizing the benchmarked functions away.
var (
	GlobalI int
	GlobalB bool
	GlobalF float64
)
