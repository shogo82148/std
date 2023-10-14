// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Tanはラジアンの引数xの正接（タンジェント）を返します。
//
// 特殊なケースは以下の通りです:
//
//	Tan(±0) = ±0
//	Tan(±Inf) = NaN
//	Tan(NaN) = NaN
func Tan(x float64) float64
