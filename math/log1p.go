// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Log1pは引数xの1を加えたものの自然対数を返します。
// xがゼロに近い場合、Log(1 + x)よりも正確です。
//
// 特殊な場合は次のとおりです：
//
//	Log1p(+Inf) = +Inf
//	Log1p(±0) = ±0
//	Log1p(-1) = -Inf
//	Log1p(x < -1) = NaN
//	Log1p(NaN) = NaN
func Log1p(x float64) float64
