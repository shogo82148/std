// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

<<<<<<< HEAD
// LdexpはFrexpの逆です。
// それはfrac × 2 ** expを返します。
=======
// Ldexp is the inverse of [Frexp].
// It returns frac × 2**exp.
>>>>>>> upstream/master
//
// 特別なケースは以下の通りです：
//
//	Ldexp(±0, exp) = ±0
//	Ldexp(±Inf, exp) = ±Inf
//	Ldexp(NaN, exp) = NaN
func Ldexp(frac float64, exp int) float64
