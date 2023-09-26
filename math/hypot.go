// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Hypot computes Sqrt(p*p + q*q), taking care to avoid
// unnecessary overflow and underflow.
//
// Special cases are:
//
//	Hypot(p, q) = +Inf if p or q is infinite
//	Hypot(p, q) = NaN if p or q is NaN
func Hypot(p, q float64) float64
