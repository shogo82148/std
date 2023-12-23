// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// Phaseはxの位相（引数とも呼ばれる）を返します。
// 返される値の範囲は[-Pi, Pi]です。
func Phase(x complex128) float64
