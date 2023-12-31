// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements rat-to-string conversion functions.

package big

import (
	"github.com/shogo82148/std/fmt"
)

var _ fmt.Scanner = &ratZero

// Scanは、fmt.Scannerのサポートルーチンです。これはフォーマット
// 'e', 'E', 'f', 'F', 'g', 'G', 'v'を受け入れます。すべてのフォーマットは同等です。
func (z *Rat) Scan(s fmt.ScanState, ch rune) error

// SetStringは、zをsの値に設定し、zと成功を示すブール値を返します。
// sは（符号付きの可能性のある）分数 "a/b"、または指数をオプションで追加した浮動小数点数として与えることができます。
// 分数が提供された場合、被除数と除数はともに10進整数であるか、またはそれぞれが独立して“0b”、“0”、“0o”、
// “0x”（またはそれらの大文字のバリエーション）のプレフィックスを使用して2進数、8進数、または16進数の整数を示すことができます。
// 除数は符号を持つことはできません。
// 浮動小数点数が提供された場合、それは10進数の形式であるか、または上記と同じプレフィックスのいずれかを使用して
// 10進数以外の仮数を示すことができます。先頭の“0”は10進数の先頭0と見なされます。この場合、8進数表現を示すものではありません。
// オプションの10進数の“e”または2進数の“p”（またはそれらの大文字のバリエーション）の指数も提供できます。
// ただし、16進数の浮動小数点数は（オプションの）“p”指数のみを受け入れます（“e”または“E”は仮数の桁と区別できないため）。
// 指数の絶対値が大きすぎる場合、操作は失敗する可能性があります。
// 文字列全体、つまりプレフィックスだけでなく、有効でなければならない。操作が失敗した場合、zの値は未定義ですが、返される値はnilです。
func (z *Rat) SetString(s string) (*Rat, bool)

// Stringは、xの文字列表現を形式 "a/b"（b == 1であっても）で返します。
func (x *Rat) String() string

// RatStringは、b != 1の場合は形式 "a/b"、b == 1の場合は形式 "a"でxの文字列表現を返します。
func (x *Rat) RatString() string

// FloatStringは、基数点の後にprec桁の精度で10進形式でxの文字列表現を返します。
// 最後の桁は最も近いものに丸められ、半分はゼロから丸められます。
func (x *Rat) FloatString(prec int) string

// FloatPrecは、xの10進表現の小数点直後の非繰り返し数字の数nを返します。
// ブール結果は、その多くの小数部分を持つxの10進表現が正確か丸められたかを示します。
//
// 例：
//
//	x      n    exact    n桁の小数部分を持つ10進表現
//	0      0    true     0
//	1      0    true     1
//	1/2    1    true     0.5
//	1/3    0    false    0       (0.333... は丸められます)
//	1/4    2    true     0.25
//	1/6    1    false    0.2     (0.166... は丸められます)
func (x *Rat) FloatPrec() (n int, exact bool)
