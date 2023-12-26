// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// Sqrtはxの平方根を返します。
// 結果のrは、real(r) ≥ 0 かつ imag(r)がimag(x)と同じ符号になるように選ばれます。
func Sqrt(x complex128) complex128
