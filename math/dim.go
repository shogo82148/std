// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Dim関数は、x-yまたは0のうちの大きい方を返します。
//
// 特別なケースは以下の通りです:
//
//	Dim(+Inf, +Inf) = NaN
//	Dim(-Inf, -Inf) = NaN
//	Dim(x, NaN) = Dim(NaN, x) = NaN
func Dim(x, y float64) float64

// Maxはxとyのうち大きい方を返します。
//
// 特殊なケースは以下の通りです:
//
//	Max(x, +Inf) = Max(+Inf, x) = +Inf
//	Max(x, NaN) = Max(NaN, x) = NaN
//	Max(+0, ±0) = Max(±0, +0) = +0
//	Max(-0, -0) = -0
//
// なお、これはNaNと+Infの場合に組み込み関数maxとは異なります。
func Max(x, y float64) float64

// Min関数は、xとyのうち小さい方を返します。
//
// 特殊な場合は以下の通りです：
//
//	Min(x, -Inf) = Min(-Inf, x) = -Inf
//	Min(x, NaN) = Min(NaN, x) = NaN
//	Min(-0, ±0) = Min(±0, -0) = -0
//
// 注意：これは、NaNと-Infを引数にして呼び出すと、組み込みのmin関数とは異なる結果になります。
func Min(x, y float64) float64
