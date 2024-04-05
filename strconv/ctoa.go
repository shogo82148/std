// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// FormatComplexは複素数cを，(a+bi)という形式の文字列に変換します．ここでaとbは実部と虚部を表し，フォーマットfmtと精度precに従って整形されます．
//
<<<<<<< HEAD
// フォーマットfmtと精度precは，FormatFloatの意味と同じです．元の値がcomplex64の場合はbitSizeが64で，complex128の場合は128であることを前提にして結果を丸めます．
=======
// The format fmt and precision prec have the same meaning as in [FormatFloat].
// It rounds the result assuming that the original was obtained from a complex
// value of bitSize bits, which must be 64 for complex64 and 128 for complex128.
>>>>>>> upstream/master
func FormatComplex(c complex128, fmt byte, prec, bitSize int) string
