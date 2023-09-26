// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// J1 returns the order-one Bessel function of the first kind.
//
// Special cases are:
//
//	J1(±Inf) = 0
//	J1(NaN) = NaN
func J1(x float64) float64

// Y1 returns the order-one Bessel function of the second kind.
//
// Special cases are:
//
//	Y1(+Inf) = 0
//	Y1(0) = -Inf
//	Y1(x < 0) = NaN
//	Y1(NaN) = NaN
func Y1(x float64) float64

// for x in [inf, 8]=1/[0,0.125]

// for x in [8,4.5454] = 1/[0.125,0.22001]

// for x in[4.5453,2.8571] = 1/[0.2199,0.35001]

// for x in [2.8570,2] = 1/[0.3499,0.5]

// for x in [inf, 8] = 1/[0,0.125]

// for x in [8,4.5454] = 1/[0.125,0.22001]

// for x in [4.5454,2.8571] = 1/[0.2199,0.35001] ???

// for x in [2.8570,2] = 1/[0.3499,0.5]
