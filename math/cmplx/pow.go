// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

<<<<<<< HEAD
// Pow returns x**y, the base-x exponential of y.
// For generalized compatibility with [math.Pow]:
=======
// Powはx**y、すなわちyの底xの指数を返します。
// math.Powとの一般的な互換性のために：
>>>>>>> release-branch.go1.21
//
//	Pow(0, ±0)は1+0iを返します
//	real(c)<0の場合のPow(0, c)は、imag(c)がゼロの場合はInf+0iを返し、それ以外の場合はInf+Inf iを返します。
func Pow(x, y complex128) complex128
