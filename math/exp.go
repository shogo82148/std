// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Expはe**x、つまりxの自然対数の底eによる指数関数です。
//
// 特殊な場合は次の通りです:
//
//	Exp(+Inf) = +Inf
//	Exp(NaN) = NaN
//
// 非常に大きな値は0または+Infにオーバーフローします。
// 非常に小さい値は1にアンダーフローします。
func Exp(x float64) float64

// Exp2はxの2の累乗（2**x）を返します。
//
// 特殊なケースはExpと同じです。
func Exp2(x float64) float64
