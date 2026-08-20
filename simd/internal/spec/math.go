// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spec

// Add adds corresponding elements of two vectors.
//
//	z[i] = x[i] + y[i]
func Add[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// DotProductPairs multiplies corresponding elements of x and y, and sums
// adjacent pairs, yielding a vector of half as many elements with twice the
// input element size.
//
//	w[i] = x[i] * y[i]        // Double width
//	z[i] = w[2*i] + w[2*i+1]
//
//specgen:require z={xB}{xN*2}x{xL/2}
func DotProductPairs[E Nums, W Width, zE Nums](x, y Vec[E, W]) (z Vec[zE, W])

// DotProductPairsSaturated multiplies corresponding elements of x and y, and
// sums adjacent pairs, all with saturation. It yields a vector of half as many
// elements with twice the input element size.
//
//	w[i] = x[i] * y[i]        // Double width, saturated
//	z[i] = w[2*i] + w[2*i+1]  // Saturated
//
//specgen:require y=Int{xN}x{xL} z=Int{xN*2}x{xL/2}
func DotProductPairsSaturated[xE Uints, xW Width, yE Ints, zE Ints](x Vec[xE, xW], y Vec[yE, xW]) (z Vec[zE, xW])
