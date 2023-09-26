// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math_test

import (
	. "math"
)

// The expected results below were computed by the high precision calculators
// at http://keisan.casio.com/.  More exact input values (array vf[], above)
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
