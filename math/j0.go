// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// J0は第1種のゼロ次ベッセル関数を返します。
//
// 特殊な場合は次の通りです:
//
//	J0(±Inf) = 0
//	J0(0) = 1
//	J0(NaN) = NaN
func J0(x float64) float64

// Y0は第二種ベッセル関数のゼロ次の値を返します。
//
// 特殊なケースは以下の通りです：
//
//	Y0(+Inf) = 0
//	Y0(0) = -Inf
//	Y0(x < 0) = NaN
//	Y0(NaN) = NaN
func Y0(x float64) float64
