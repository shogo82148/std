// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// SincosはSin(x), Cos(x)を返します。
//
// 特殊なケースは以下の通りです:
//
//	Sincos(±0) = ±0, 1
//	Sincos(±Inf) = NaN, NaN
//	Sincos(NaN) = NaN, NaN
func Sincos(x float64) (sin, cos float64)
