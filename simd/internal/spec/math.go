// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//specgen:category Math

package spec

// AbsFloat returns the elementwise absolute value of x.
//
//specgen:name Abs
func AbsFloat[E Floats, W Width](x Vec[E, W]) (z Vec[E, W])

// AbsInt returns the elementwise absolute value of x.
//
//specgen:name Abs
//specgen:require z=Uint{xN}x{xL}
func AbsInt[E Ints, W Width, zE Uints](x Vec[E, W]) (z Vec[zE, W])

// Add adds x and y elementwise.
//
//	z[i] = x[i] + y[i]
//
//specgen:commutative
func Add[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// Sub subtracts y from x elementwise.
//
//	z[i] = x[i] - y[i]
func Sub[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// AddSaturated adds x and y elementwise with saturation.
//
//	z[i] = sat(x[i] + y[i])
//
//specgen:commutative
func AddSaturated[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// SubSaturated subtracts y from x elementwise with saturation.
//
//	z[i] = sat(x[i] - y[i])
func SubSaturated[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatAddPairs horizontally adds adjacent pairs of elements in x and y and
// returns the concatenated result.
//
// {{if eq .xL 2}}
//
//	z = {x[0]+x[1], y[0]+y[1]}
//
// {{else if eq .xL 4}}
//
//	z = {x[0]+x[1], x[2]+x[3], y[0]+y[1], y[2]+y[3]}
//
// {{else}}
//
//	z = {x[0]+x[1], x[2]+x[3], ..., y[0]+y[1], y[2]+y[3], ...}
//
// {{end}}
func ConcatAddPairs[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatSubPairs horizontally subtracts adjacent pairs of elements in x and y
// and returns the concatenated result.
//
// {{if eq .xL 2}}
//
//	z = {x[0]-x[1], y[0]-y[1]}
//
// {{else if eq .xL 4}}
//
//	z = {x[0]-x[1], x[2]-x[3], y[0]-y[1], y[2]-y[3]}
//
// {{else}}
//
//	z = {x[0]-x[1], x[2]-x[3], ..., y[0]-y[1], y[2]-y[3], ...}
//
// {{end}}
func ConcatSubPairs[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatAddPairsSaturated horizontally adds adjacent pairs of elements in x and
// y with saturation and returns the concatenated result.
//
// {{if eq .xL 2}}
//
//	z = {sat(x[0]+x[1]), sat(y[0]+y[1])}
//
// {{else if eq .xL 4}}
//
//	z = {sat(x[0]+x[1]), sat(x[2]+x[3]), sat(y[0]+y[1]), sat(y[2]+y[3])}
//
// {{else}}
//
//	z = {sat(x[0]+x[1]), sat(x[2]+x[3]), ..., sat(y[0]+y[1]), sat(y[2]+y[3]), ...}
//
// {{end}}
func ConcatAddPairsSaturated[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatSubPairsSaturated horizontally subtracts adjacent pairs of elements in
// x and y with saturation and returns the concatenated result.
//
// {{if eq .xL 2}}
//
//	z = {x[0]-x[1], y[0]-y[1]}
//
// {{else if eq .xL 4}}
//
//	z = {sat(x[0]-x[1]), sat(x[2]-x[3]), sat(y[0]-y[1]), sat(y[2]-y[3])}
//
// {{else}}
//
//	z = {x[0]-x[1], x[2]-x[3], ..., y[0]-y[1], y[2]-y[3], ...}
//
// {{end}}
func ConcatSubPairsSaturated[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatAddPairsGrouped divides x, y, and z into groups of 128 bits and
// performs [ConcatAddPairs] on each group.
func ConcatAddPairsGrouped[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatSubPairsGrouped divides x, y, and z into groups of 128 bits and
// performs [ConcatSubPairs] on each group.
func ConcatSubPairsGrouped[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatAddPairsSaturatedGrouped divides x, y, and z into groups of 128 bits
// and performs [ConcatAddPairsSaturated] on each group.
func ConcatAddPairsSaturatedGrouped[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatSubPairsSaturatedGrouped divides x, y, and z into groups of 128 bits
// and performs [ConcatSubPairsSaturated] on each group.
func ConcatSubPairsSaturatedGrouped[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// Average returns the elementwise average of x and y, rounded toward +∞.
//
//	z[i] = (x[i] + y[i] + 1) / 2
func Average[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// DivInts divides x by y elementwise.
//
//	z[i] = x[i] / y[i]
//
// The result is rounded toward zero (truncated division), just like Go's /
// operator.{{if eq .xB "int"}} Also like Go's / operator, dividing Min{{title .xE}}
// by -1 results in Min{{title .xE}}, since the true result is unrepresentable.{{end}}
// Unlike Go's / operator, division by zero does not panic and instead produces
// a zero result for that element.
//
//specgen:name Div
func DivInts[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// DivFloats divides x by y elementwise.
//
//	z[i] = x[i] / y[i]
//
// Division by zero does not panic, and the result follows IEEE 754. That is,
// dividing a non-zero value by zero results in +/- infinity, and dividing zero
// by zero results in NaN.
//
//specgen:name Div
func DivFloats[E Floats, W Width](x, y Vec[E, W]) (z Vec[E, W])

// DotProductPairs multiplies corresponding elements of x and y, and sums
// adjacent pairs, returning a vector of half as many elements, each with twice
// the input element size.
//
//	w[i] = x[i] * y[i]        // Double width
//	z[j] = w[2*j] + w[2*j+1]
//
// While the products can never overflow, the sum wraps on overflow.
// {{if eq .xB "int"}}
// This happens in exactly one situation: if all four input elements are Min{{title .xE}}.
// {{end}}
//
//specgen:commutative
//specgen:require z={xB}{xN*2}x{xL/2}
func DotProductPairs[E Ints | Uints, W Width, zE Nums](x, y Vec[E, W]) (z Vec[zE, W])

// DotProductPairsSaturated multiplies corresponding elements of x and y, and
// sums adjacent pairs, all with saturation. It returns a vector of half as many
// elements, each with twice the input element size.
//
//	w[i] = sat(x[i] * y[i])        // Double width
//	z[j] = sat(w[2*j] + w[2*j+1])
//
//specgen:commutative
//specgen:require y=Int{xN}x{xL} z=Int{xN*2}x{xL/2}
func DotProductPairsSaturated[xE Uints, xW Width, yE Ints, zE Ints](x Vec[xE, xW], y Vec[yE, xW]) (z Vec[zE, xW])

// MaxInts returns the elementwise maximum of x and y.
//
//	z[i] = max(x[i], y[i])
//
//specgen:name Max
//specgen:commutative
func MaxInts[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// MaxFloats returns the elementwise maximum of x and y.
//
//	z[i] = max(x[i], y[i])
//
// The handling of NaN and signed zeroes is architecture-dependent. ARM64 and
// WASM follow the IEEE 754-2019 maximum algorithm, which also agrees with Go's
// max built-in. AMD64 handles these differently. Notably, AMD64 is
// non-commutative.
//
//	x     y     IEEE  AMD64
//	----- ----- ----- -----
//	NaN   y     NaN   y
//	x     NaN   NaN   NaN
//	-Inf  y     y     y
//	x     -Inf  x     x
//	+Inf  y     +Inf  +Inf
//	x     +Inf  +Inf  +Inf
//	+0    -0    +0    -0
//	-0    +0    +0    +0
//
//specgen:name Max
func MaxFloats[E Floats, W Width](x, y Vec[E, W]) (z Vec[Variation[E], W])

// MinInts returns the elementwise minimum of x and y.
//
//	z[i] = min(x[i], y[i])
//
//specgen:name Min
//specgen:commutative
func MinInts[E Ints | Uints, W Width](x, y Vec[E, W]) (z Vec[E, W])

// MinFloats returns the elementwise minimum of x and y.
//
//	z[i] = min(x[i], y[i])
//
// The handling of NaN and signed zeroes is architecture-dependent. ARM64 and
// WASM follow the IEEE 754-2019 minimum algorithm, which also agrees with Go's
// min built-in. AMD64 handles these differently. Notably, AMD64 is
// non-commutative.
//
//	x     y     IEEE  AMD64
//	----- ----- ----- -----
//	NaN   y     NaN   y
//	x     NaN   NaN   NaN
//	-Inf  y     -Inf  -Inf
//	x     -Inf  -Inf  -Inf
//	+Inf  y     y     y
//	x     +Inf  x     x
//	+0    -0    -0    -0
//	-0    +0    -0    +0
//
//specgen:name Min
func MinFloats[E Floats, W Width](x, y Vec[E, W]) (z Vec[Variation[E], W])

// Mul returns the elementwise product of x and y.
//
//	z[i] = x[i] * y[i]
//
//specgen:commutative
func Mul[E Elt, W Width](x, y Vec[E, W]) (z Vec[E, W])

// MulAddInts returns x * y + z elementwise.
//
//	w[i] = x[i] * y[i] + z[i]
//
//specgen:name MulAdd
func MulAddInts[E Ints | Uints, W Width](x, y, z Vec[E, W]) (w Vec[E, W])

// MulAddFloats returns x * y + z elementwise.
//
//	w[i] = x[i] * y[i] + z[i]
//
// This performs a fused multiply-add, computing x[i]*y[i] + z[i] with a single
// rounding rather than two and no intermediate overflow, except on wasm, where
// it rounds x[i]*y[i] before adding z[i].
//
//specgen:name MulAdd
func MulAddFloats[E Floats, W Width](x, y, z Vec[E, W]) (w Vec[Variation[E], W])

// Neg returns the elementwise negation of x.
//
//	z[i] = -x[i]
func Neg[E Ints | Floats, W Width](x Vec[E, W]) (z Vec[E, W])

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
func ReduceMax[E Elt, W Width](x Vec[E, W]) (z E)

// Sqrt returns the elementwise square root of x.
//
//	z[i] = sqrt(x[i])
//
// The result is correctly rounded. Sqrt(+Inf) = +Inf, Sqrt(±0) = ±0, and
// Sqrt(x) is NaN if x < 0 or x is NaN. The sign and payload of a NaN result
// are unspecified.
func Sqrt[E Floats, W Width](x Vec[E, W]) (z Vec[E, W])
