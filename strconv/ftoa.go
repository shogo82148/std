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
<<<<<<< HEAD
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
=======
// The format fmt is one of
//   - 'b' (-ddddp±ddd, a binary exponent),
//   - 'e' (-d.dddde±dd, a decimal exponent),
//   - 'E' (-d.ddddE±dd, a decimal exponent),
//   - 'f' (-ddd.dddd, no exponent),
//   - 'g' ('e' for large exponents, 'f' otherwise),
//   - 'G' ('E' for large exponents, 'f' otherwise),
//   - 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
//   - 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
//
// The precision prec controls the number of digits (excluding the exponent)
// printed by the 'e', 'E', 'f', 'g', 'G', 'x', and 'X' formats.
// For 'e', 'E', 'f', 'x', and 'X', it is the number of digits after the decimal point.
// For 'g' and 'G' it is the maximum number of significant digits (trailing
// zeros are removed).
// The special precision -1 uses the smallest number of digits
// necessary such that ParseFloat will return f exactly.
// The exponent is written as a decimal integer;
// for all formats other than 'b', it will be at least two digits.
>>>>>>> upstream/release-branch.go1.25
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

// AppendFloatは、[FormatFloat] によって生成された浮動小数点数fの文字列形式をdstに追加し、拡張されたバッファを返します。
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
