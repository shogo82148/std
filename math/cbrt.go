// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Cbrtはxの立方根を返します。
//
// 特殊なケースは次の通りです:
//
//	Cbrt(±0) = ±0
//	Cbrt(±Inf) = ±Inf
//	Cbrt(NaN) = NaN
func Cbrt(x float64) float64
