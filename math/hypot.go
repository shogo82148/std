// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

<<<<<<< HEAD
// Hypotはp*p + q*qの平方根を返します。不必要なオーバーフローやアンダーフローを避けるように注意します。
=======
// Hypot returns [Sqrt](p*p + q*q), taking care to avoid
// unnecessary overflow and underflow.
>>>>>>> upstream/master
//
// 特殊な場合は以下の通りです:
//
//	Hypot(±Inf, q) = +Inf
//	Hypot(p, ±Inf) = +Inf
//	Hypot(NaN, q) = NaN
//	Hypot(p, NaN) = NaN
func Hypot(p, q float64) float64
