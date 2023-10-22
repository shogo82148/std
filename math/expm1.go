// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

<<<<<<< HEAD
// Expm1はxの自然対数e**x - 1を返します。
// xがほぼゼロの場合、Exp(x) - 1よりも精度が高くなります。
=======
// Expm1 returns e**x - 1, the base-e exponential of x minus 1.
// It is more accurate than [Exp](x) - 1 when x is near zero.
>>>>>>> upstream/master
//
// 特殊なケースは以下の通りです：
//
//	Expm1(+Inf) = +Inf
//	Expm1(-Inf) = -1
//	Expm1(NaN) = NaN
//
// 非常に大きな値は-1または+Infにオーバーフローします。
func Expm1(x float64) float64
