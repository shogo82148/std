// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Logはxの自然対数を返します。
//
// 特別なケースは以下の通りです:
//
//	Log(+Inf) = +Inf
//	Log(0) = -Inf
//	Log(x < 0) = NaN
//	Log(NaN) = NaN
func Log(x float64) float64
