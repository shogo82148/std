// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Remainderはx/yのIEEE 754浮動小数点余りを返します。
//
// 特殊なケースは以下の通りです：
//
//	Remainder(±Inf, y) = NaN
//	Remainder(NaN, y) = NaN
//	Remainder(x, 0) = NaN
//	Remainder(x, ±Inf) = x
//	Remainder(x, NaN) = NaN
func Remainder(x, y float64) float64
