// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Erfはxの誤差関数を返します。
//
// 特殊なケースは以下の通りです:
//
//	Erf(+Inf) = 1
//	Erf(-Inf) = -1
//	Erf(NaN) = NaN
func Erf(x float64) float64

// Erfcはxの補完誤差関数を返します。
//
// 特別なケースは以下の通りです:
//
//	Erfc(+Inf) = 0
//	Erfc(-Inf) = 2
//	Erfc(NaN) = NaN
func Erfc(x float64) float64
