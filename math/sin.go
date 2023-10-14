// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Cosは、ラジアンの引数xの余弦を返します。
//
// 特殊な場合は:
//
//	Cos(±Inf) = NaN
//	Cos(NaN) = NaN
func Cos(x float64) float64

// Sinは、ラジアンの引数xの正弦を返します。
//
// 特殊なケースは次の通りです：
//
//	Sin(±0) = ±0
//	Sin(±Inf) = NaN
//	Sin(NaN) = NaN
func Sin(x float64) float64
