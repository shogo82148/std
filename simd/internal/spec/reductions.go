// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//specgen:category Reductions

package spec

// ReduceSumInts returns the scalar sum of the elements of x.
//
//	z = x[0] + x[1] + ...
//
//specgen:name ReduceSum
func ReduceSumInts[E Ints | Uints, W Width](x Vec[E, W]) (z E)

// ReduceSumFloats returns the scalar sum of the elements of x.
//
// {{if eq .xL 2}}
//
//	z = x[0] + x[1]
//
// {{else}}
//
// This uses repeated pairwise summation, which may affect the final precision
// of the result.
//
// {{if eq .xL 4}}
//
//	z = (x[0] + x[1]) + (x[2] + x[3])
//
// {{else}}
//
//	z = ((x[0] + x[1]) + (x[2] + x[3])) + ...
//
// {{end}}
//
// {{end}}
//
//specgen:name ReduceSum
func ReduceSumFloats[E Floats, W Width](x Vec[E, W]) (z E)

// ReduceMax returns the largest element of x.
//
//	z = max(x[0], x[1], ...)
//
// {{if eq .xB "float"}}
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's max built-in and the IEEE 754-2019 maximum operation. Note
// that on AMD64 (and only AMD64), it differs from [Max].
// {{end}}
func ReduceMax[E Elt, W Width](x Vec[E, W]) (z E)

// ReduceMin returns the smallest element of x.
//
//	z = min(x[0], x[1], ...)
//
// {{if eq .xB "float"}}
// If any element is NaN, the result is NaN, and -0 is considered less than +0.
// This matches Go's min built-in and the IEEE 754-2019 minimum operation. Note
// that on AMD64 (and only AMD64), it differs from [Min].
// {{end}}
func ReduceMin[E Elt, W Width](x Vec[E, W]) (z E)
