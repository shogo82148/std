// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Tanhはxの双曲線正接を返します。
//
// 特殊な場合は次のとおりです：
//
//	Tanh(±0) = ±0
//	Tanh(±Inf) = ±1
//	Tanh(NaN) = NaN
func Tanh(x float64) float64
