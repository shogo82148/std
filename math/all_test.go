// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math_test

// arguments and expected results for boundary cases
const (
	SmallestNormalFloat64   = 2.2250738585072014e-308
	LargestSubnormalFloat64 = SmallestNormalFloat64 - SmallestNonzeroFloat64
)

var PortableFMA = FMA

// Global exported variables are used to store the
// return values of functions measured in the benchmarks.
// Storing the results in these variables prevents the compiler
// from completely optimizing the benchmarked functions away.
var (
	GlobalI int
	GlobalB bool
	GlobalF float64
)
