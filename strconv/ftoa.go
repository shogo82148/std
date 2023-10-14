// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Binary to decimal floating point conversion.
// Algorithm:
//   1) store mantissa in multiprecision decimal
//   2) shift decimal by exponent
//   3) read digits out & format

package strconv

// FormatFloatは、浮動小数点数fを、フォーマットfmtと精度precに従って文字列に変換します。
// それは、元の値がbitSizeビット（float32の場合は32、float64の場合は64）の浮動小数点値から得られたものと仮定して、
// 結果を四捨五入します。
//
// フォーマットfmtは、次のいずれかです。
// 'b' (-ddddp±ddd、2進指数表記),
// 'e' (-d.dddde±dd、10進指数表記),
// 'E' (-d.ddddE±dd、10進指数表記),
// 'f' (-ddd.dddd、指数なし),
// 'g' (指数が大きい場合は'e'、そうでない場合は'f'),
// 'G' (指数が大きい場合は'E'、そうでない場合は'f'),
// 'x' (-0xd.ddddp±ddd、16進数の小数部と2進指数表記), または
// 'X' (-0Xd.ddddP±ddd、16進数の小数部と2進指数表記)。
//
// 精度precは、'e'、'E'、'f'、'g'、'G'、'x'、および'X'フォーマットによって出力される指数を除く桁数を制御します。
// 'e'、'E'、'f'、'x'、および'X'の場合、小数点以下の桁数です。
// 'g'および'G'の場合、最大有効桁数です（末尾のゼロは削除されます）。
// 特別な精度-1は、ParseFloatがfを正確に返すために必要な最小桁数を使用します。
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

// AppendFloatは、FormatFloatによって生成された浮動小数点数fの文字列形式をdstに追加し、拡張されたバッファを返します。
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
