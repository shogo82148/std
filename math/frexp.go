// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Frexpは、fを正規化された分数と2の整数冪に分解します。
// それは、f == frac × 2**expを満たすfracとexpを返します。
// fracの絶対値は[½, 1)の範囲にあります。
//
// 特殊な場合は以下の通りです:
//
//	Frexp(±0) = ±0, 0
//	Frexp(±Inf) = ±Inf, 0
//	Frexp(NaN) = NaN, 0
func Frexp(f float64) (frac float64, exp int)
