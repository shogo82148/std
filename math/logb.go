// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Logbはxの二進指数を返します。
//
// 特殊なケースは次の通りです：
//
//	Logb(±Inf) = +Inf
//	Logb(0) = -Inf
//	Logb(NaN) = NaN
func Logb(x float64) float64

// Ilogbはxのバイナリ指数を整数として返します。
//
// 特殊なケースは以下の通りです:
//
//	Ilogb(±Inf) = MaxInt32
//	Ilogb(0) = MinInt32
//	Ilogb(NaN) = MaxInt32
func Ilogb(x float64) int
