// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Expm1はxの自然対数e**x - 1を返します。
// xがほぼゼロの場合、Exp(x) - 1よりも精度が高くなります。
//
// 特殊なケースは以下の通りです：
//
//	Expm1(+Inf) = +Inf
//	Expm1(-Inf) = -1
//	Expm1(NaN) = NaN
//
// 非常に大きな値は-1または+Infにオーバーフローします。
func Expm1(x float64) float64
