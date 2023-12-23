// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// Polarはxの絶対値rと位相θを返します。
// そのため、x = r * e**θiとなります。
// 位相は範囲[-Pi, Pi]内にあります。
func Polar(x complex128) (r, θ float64)
