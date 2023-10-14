// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Modfは整数部分と小数部分の浮動小数点数を返します。両方の値はfと同じ符号を持ちます。
//
// 特殊なケースは以下の通りです：
//
//	Modf(±Inf) = ±Inf, NaN
//	Modf(NaN) = NaN, NaN
func Modf(f float64) (int float64, frac float64)
