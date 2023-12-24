// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements Float-to-string conversion functions.
// It is closely following the corresponding implementation
// in strconv/ftoa.go, but modified and simplified for Float.

package big

import (
	"github.com/shogo82148/std/fmt"
)

// Textは、与えられたフォーマットと精度precに従って、浮動小数点数xを文字列に変換します。
// フォーマットは次のいずれかです:
//
//	'e'	-d.dddde±dd, 10進数の指数、少なくとも2つ（可能性のある0）の指数の桁
//	'E'	-d.ddddE±dd, 10進数の指数、少なくとも2つ（可能性のある0）の指数の桁
//	'f'	-ddddd.dddd, 指数なし
//	'g'	大きな指数の場合は'e'のように、それ以外の場合は'f'のように
//	'G'	大きな指数の場合は'E'のように、それ以外の場合は'f'のように
//	'x'	-0xd.dddddp±dd, 16進数の仮数、2の力の10進数の指数
//	'p'	-0x.dddp±dd, 16進数の仮数、2の力の10進数の指数（非標準）
//	'b'	-ddddddp±dd, 10進数の仮数、2の力の10進数の指数（非標準）
//
// 2の力の指数形式の場合、仮数は正規化された形式で印刷されます:
//
//	'x'	16進数の仮数は[1, 2)、または0
//	'p'	16進数の仮数は[½, 1)、または0
//	'b'	x.Prec()ビットを使用した10進数の整数仮数、または0
//
// 'x'形式は、他のほとんどの言語やライブラリで使用されている形式であることに注意してください。
//
// フォーマットが異なる文字の場合、Textは"%"と認識されないフォーマット文字を続けて返します。
//
// 精度precは、'e'、'E'、'f'、'g'、'G'、および'x'の形式で印刷される桁数（指数を除く）を制御します。
// 'e'、'E'、'f'、および'x'の場合、それは小数点の後の桁数です。
// 'g'と'G'の場合、それは全体の桁数です。負の精度は、x.Prec()マンティッサビットを使用して
// 値xを一意に識別するために必要な最小の10進数の桁数を選択します。
// 'b'と'p'の形式では、prec値は無視されます。
func (x *Float) Text(format byte, prec int) string

<<<<<<< HEAD
// String formats x like x.Text('g', 10).
// (String must be called explicitly, [Float.Format] does not support %s verb.)
=======
// Stringはxをx.Text('g', 10)のようにフォーマットします。
// (Stringは明示的に呼び出す必要があります、Float.Formatは%s動詞をサポートしていません。)
>>>>>>> release-branch.go1.21
func (x *Float) String() string

// Appendは、x.Textによって生成された浮動小数点数xの文字列形式をbufに追加し、
// 拡張されたバッファを返します。
func (x *Float) Append(buf []byte, fmt byte, prec int) []byte

var _ fmt.Formatter = &floatZero

<<<<<<< HEAD
// Format implements [fmt.Formatter]. It accepts all the regular
// formats for floating-point numbers ('b', 'e', 'E', 'f', 'F',
// 'g', 'G', 'x') as well as 'p' and 'v'. See (*Float).Text for the
// interpretation of 'p'. The 'v' format is handled like 'g'.
// Format also supports specification of the minimum precision
// in digits, the output field width, as well as the format flags
// '+' and ' ' for sign control, '0' for space or zero padding,
// and '-' for left or right justification. See the fmt package
// for details.
=======
// Formatはfmt.Formatterを実装します。通常の浮動小数点数のフォーマット('b', 'e', 'E', 'f', 'F',
// 'g', 'G', 'x')と同様に'p'と'v'も受け入れます。'p'の解釈については(*Float).Textを参照してください。
// 'v'フォーマットは'g'のように扱われます。
// Formatはまた、最小精度の指定、出力フィールドの幅、および符号制御のための'+'と' 'フラグ、
// スペースまたはゼロパディングのための'0'、左または右の正当化のための'-'フラグもサポートしています。
// 詳細はfmtパッケージを参照してください。
>>>>>>> release-branch.go1.21
func (x *Float) Format(s fmt.State, format rune)
