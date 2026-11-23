// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//specgen:category Math (concatenated)

package spec

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

// ConcatMaxPairs horizontally takes the maximum of adjacent pairs of elements
// in x and y and returns the concatenated result.
//
// {{if eq .xL 2}}
//
//	z = {max(x[0], x[1]), max(y[0], y[1])}
//
// {{else if eq .xL 4}}
//
//	z = {max(x[0], x[1]), max(x[2], x[3]), max(y[0], y[1]), max(y[2], y[3])}
//
// {{else}}
//
//	z = {max(x[0], x[1]), max(x[2], x[3]), ..., max(y[0], y[1]), max(y[2], y[3]), ...}
//
// {{end}}
//
// For floats, max is Go's max built-in: it returns NaN if either element is
// NaN, and orders -0 before +0.
func ConcatMaxPairs[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

// ConcatMinPairs horizontally takes the minimum of adjacent pairs of elements
// in x and y and returns the concatenated result.
//
// {{if eq .xL 2}}
//
//	z = {min(x[0], x[1]), min(y[0], y[1])}
//
// {{else if eq .xL 4}}
//
//	z = {min(x[0], x[1]), min(x[2], x[3]), min(y[0], y[1]), min(y[2], y[3])}
//
// {{else}}
//
//	z = {min(x[0], x[1]), min(x[2], x[3]), ..., min(y[0], y[1]), min(y[2], y[3]), ...}
//
// {{end}}
//
// For floats, min is Go's min built-in: it returns NaN if either element is
// NaN, and orders -0 before +0.
func ConcatMinPairs[E Nums, W Width](x, y Vec[E, W]) (z Vec[E, W])

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
//	z = {sat(x[0]-x[1]), sat(y[0]-y[1])}
//
// {{else if eq .xL 4}}
//
//	z = {sat(x[0]-x[1]), sat(x[2]-x[3]), sat(y[0]-y[1]), sat(y[2]-y[3])}
//
// {{else}}
//
//	z = {sat(x[0]-x[1]), sat(x[2]-x[3]), ..., sat(y[0]-y[1]), sat(y[2]-y[3]), ...}
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
