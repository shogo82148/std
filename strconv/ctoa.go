// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// FormatComplexは複素数cを，(a+bi)という形式の文字列に変換します．ここでaとbは実部と虚部を表し，フォーマットfmtと精度precに従って整形されます．
//
// フォーマットfmtと精度precは，[FormatFloat] の意味と同じです．元の値がcomplex64の場合はbitSizeが64で，complex128の場合は128であることを前提にして結果を丸めます．
func FormatComplex(c complex128, fmt byte, prec, bitSize int) string
