// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// J1は第1種の1次ベッセル関数を返します。
//
// 特殊なケースは以下の通りです：
//
//	J1(±Inf) = 0
//	J1(NaN) = NaN
func J1(x float64) float64

// Y1は第2種ベッセル関数の1次の値を返します。
//
// 特殊な場合は以下の通りです:
//
//	Y1(+Inf) = 0
//	Y1(0) = -Inf
//	Y1(x < 0) = NaN
//	Y1(NaN) = NaN
func Y1(x float64) float64
