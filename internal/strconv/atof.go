// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

// ParseFloatは、文字列sを、bitSizeで指定された精度の浮動小数点数に変換します。
// bitSize=32の場合、結果はfloat64型のままですが、値を変更せずにfloat32型に変換できます。
//
// ParseFloatは、[浮動小数点リテラル] のGo構文で定義された10進数および16進数の浮動小数点数を受け入れます。
// sが正しく形成され、有効な浮動小数点数に近い場合、ParseFloatはIEEE754のバイアスのない丸めを使用して最も近い浮動小数点数を返します。
// （16進数の浮動小数点値を解析する場合、16進数表現にビットが多すぎてマンティッサに収まらない場合にのみ丸めが行われます。）
//
// ParseFloatが返すエラーは、*NumErrorの具体的な型であり、err.Num = sを含みます。
//
// sが構文的に正しくない場合、ParseFloatはerr.Err = ErrSyntaxを返します。
//
// sが構文的に正しく、与えられたサイズの最大浮動小数点数から1/2 ULP以上離れている場合、
// ParseFloatはf = ±Inf、err.Err = ErrRangeを返します。
//
// ParseFloatは、文字列 "NaN" および (可能な場合は符号付きの) 文字列 "Inf" および "Infinity" を、
// それぞれ特別な浮動小数点値として認識します。大文字小文字は区別されません。
//
// [浮動小数点リテラル]: https://go.dev/ref/spec#Floating-point_literals
func ParseFloat(s string, bitSize int) (float64, error)
