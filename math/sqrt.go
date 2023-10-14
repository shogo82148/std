// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Sqrtはxの平方根を返します。
//
// 特殊な場合は次の通りです：
//
//	Sqrt(+Inf) = +Inf
//	Sqrt(±0) = ±0
//	Sqrt(x < 0) = NaN
//	Sqrt(NaN) = NaN
func Sqrt(x float64) float64
