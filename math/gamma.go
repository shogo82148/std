// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Gamma(x) returns the Gamma function of x.
//
// Special cases are:
//
//	Gamma(±Inf) = ±Inf
//	Gamma(NaN) = NaN
//
// Large values overflow to +Inf.
// Zero and negative integer arguments return ±Inf.
func Gamma(x float64) float64
