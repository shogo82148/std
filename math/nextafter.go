// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Nextafter32はxからyの方向に次に表現可能なfloat32値を返します。
//
// 特殊なケースは以下の通りです:
//
//	Nextafter32(x, x)   = x
//	Nextafter32(NaN, y) = NaN
//	Nextafter32(x, NaN) = NaN
func Nextafter32(x, y float32) (r float32)

// Nextafterはxからyに向かって次の表現可能なfloat64値を返します。
//
// 特別な場合は以下の通りです:
//
//	Nextafter(x, x)   = x
//	Nextafter(NaN, y) = NaN
//	Nextafter(x, NaN) = NaN
func Nextafter(x, y float64) (r float64)
