// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Acoshはxの逆双曲線余弦を返します。
//
// 特殊な場合は以下の通りです:
//
//	Acosh(+Inf) = +Inf
//	Acosh(x) = NaN if x < 1
//	Acosh(NaN) = NaN
func Acosh(x float64) float64
