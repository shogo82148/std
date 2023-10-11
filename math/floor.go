// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Floorはx以下の最大の整数値を返します。
//
// 特殊なケースは次の通りです：
//
//	Floor(±0) = ±0
//	Floor(±Inf) = ±Inf
//	Floor(NaN) = NaN
func Floor(x float64) float64

// Ceil（天井）は、x以上の最小の整数値を返します。
//
// 特別なケースは:
//
//	Ceil(±0) = ±0
//	Ceil(±Inf) = ±Inf
//	Ceil(NaN) = NaN
func Ceil(x float64) float64

// Truncはxの整数値を返します。
//
// 特殊なケースは以下の通りです:
//
//	Trunc(±0) = ±0
//	Trunc(±Inf) = ±Inf
//	Trunc(NaN) = NaN
func Trunc(x float64) float64

// Roundは、最も近い整数を返します。半の場合はゼロから離れます。
//
// 特殊なケースは次の通りです：
//
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func Round(x float64) float64

// RoundToEvenは、最も近い整数を返し、丸めの際に引き分けを偶数に丸めます。
//
// 特別なケースは以下の通りです：
//
//	RoundToEven(±0) = ±0
//	RoundToEven(±Inf) = ±Inf
//	RoundToEven(NaN) = NaN
func RoundToEven(x float64) float64
