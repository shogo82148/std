// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// J0 returns the order-zero Bessel function of the first kind.
//
// Special cases are:
//
//	J0(±Inf) = 0
//	J0(0) = 1
//	J0(NaN) = NaN
func J0(x float64) float64

// Y0 returns the order-zero Bessel function of the second kind.
//
// Special cases are:
//
//	Y0(+Inf) = 0
//	Y0(0) = -Inf
//	Y0(x < 0) = NaN
//	Y0(NaN) = NaN
func Y0(x float64) float64

// for x in [inf, 8]=1/[0,0.125]

// for x in [8,4.5454]=1/[0.125,0.22001]

// for x in [4.547,2.8571]=1/[0.2199,0.35001]

// for x in [2.8570,2]=1/[0.3499,0.5]

// for x in [inf, 8]=1/[0,0.125]

// for x in [8,4.5454]=1/[0.125,0.22001]

// for x in [4.547,2.8571]=1/[0.2199,0.35001]

// for x in [2.8570,2]=1/[0.3499,0.5]
