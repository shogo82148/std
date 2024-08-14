// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// cmplxパッケージは、複素数の基本的な定数と数学関数を提供します。
// 特殊なケースの処理は、C99標準の付録G IEC 60559互換の複素数演算に準拠しています。
package cmplx

// Absはxの絶対値（またはモジュラスとも呼ばれる）を返します。
func Abs(x complex128) float64
