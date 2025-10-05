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
// fmtの形式は以下のいずれかです
//   - 'b' (-ddddp±ddd、2進指数),
//   - 'e' (-d.dddde±dd、10進指数),
//   - 'E' (-d.ddddE±dd、10進指数),
//   - 'f' (-ddd.dddd、指数なし),
//   - 'g' (大きな指数の場合は'e'、それ以外は'f'),
//   - 'G' (大きな指数の場合は'E'、それ以外は'f'),
//   - 'x' (-0xd.ddddp±ddd、16進小数と2進指数)、または
//   - 'X' (-0Xd.ddddP±ddd、16進小数と2進指数)。
//
// 精度precは、'e', 'E', 'f', 'g', 'G', 'x', 'X'形式で出力される桁数（指数部を除く）を制御します。
// 'e', 'E', 'f', 'x', 'X'の場合は小数点以下の桁数です。
// 'g', 'G'の場合は有効数字の最大桁数です（末尾のゼロは削除されます）。
// 特別な精度-1は、ParseFloatが正確にfを返すのに必要な最小限の桁数を使用します。
// 指数は10進整数で書かれます。
// 'b'以外のすべての形式では、指数部は少なくとも2桁になります。
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

// AppendFloatは、[FormatFloat] によって生成された浮動小数点数fの文字列形式をdstに追加し、拡張されたバッファを返します。
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
