// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// ParseComplexは文字列sを、bitSizeで指定された精度で複素数に変換します。
// bitSizeが64の場合、結果は型complex128になりますが、値を変更せずにcomplex64に変換することができます。
//
// sに表される数値は、N、Ni、またはN±Niの形式でなければなりません。ここで、Nは [ParseFloat] で認識される浮動小数点数を表し、iは虚数成分です。
// 2つ目のNが非負である場合、±によって示される2つの成分の間に+記号が必要です。2つ目のNがNaNである場合、+記号のみが受け入れられます。
// 形式は括弧で囲んでも良く、スペースを含んではいけません。
// 変換される複素数は、ParseFloatによって変換された2つの成分から構成されます。
//
// ParseComplexが返すエラーは、[*NumError] という具体的な型であり、err.Num = sとなります。
//
// sが構文的に正しくない場合、ParseComplexはerr.Err = ErrSyntaxを返します。
//
// sが構文的に正しくなっているが、いずれかの成分が与えられた成分のサイズで最大の浮動小数点数から1/2 ULP以上離れている場合、
// ParseComplexはerr.Err = ErrRangeとc = ±Infを返します。
func ParseComplex(s string, bitSize int) (complex128, error)
