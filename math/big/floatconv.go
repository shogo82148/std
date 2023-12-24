// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements string-to-Float conversion functions.

package big

import (
	"github.com/shogo82148/std/fmt"
)

<<<<<<< HEAD
// SetString sets z to the value of s and returns z and a boolean indicating
// success. s must be a floating-point number of the same format as accepted
// by [Float.Parse], with base argument 0. The entire string (not just a prefix) must
// be valid for success. If the operation failed, the value of z is undefined
// but the returned value is nil.
=======
// SetStringは、zをsの値に設定し、zと成功を示すブール値を返します。
// sは、基数引数0でParseによって受け入れられるのと同じ形式の浮動小数点数でなければなりません。
// 成功するためには、文字列全体（プレフィックスだけでなく）が有効でなければなりません。
// 操作が失敗した場合、zの値は未定義ですが、返される値はnilです。
>>>>>>> release-branch.go1.21
func (z *Float) SetString(s string) (*Float, bool)

// Parseは、指定された変換基数で仮数を持つ浮動小数点数のテキスト表現、
// または無限大を表す文字列を含むsを解析します（指数は常に10進数です）。
//
// 基数が0の場合、アンダースコア文字 "_" は基数のプレフィックスと隣接する数字の間、
// または連続する数字の間に現れることがあります。そのようなアンダースコアは、
// 数値の値や返される数字の数に変化を与えません。アンダースコアの配置が間違っている場合、
// 他にエラーがない場合に限り、エラーとして報告されます。基数が0でない場合、
// アンダースコアは認識されず、したがって、有効な基数点または数字でない他の任意の文字と同様に、
// スキャンを終了します。
//
// zを対応する浮動小数点値の（可能性のある丸められた）値に設定し、
// z、実際の基数b、およびエラーerr（ある場合）を返します。
// 成功するためには、文字列全体（プレフィックスだけでなく）が消費されなければなりません。
// zの精度が0の場合、丸めが効く前に64に変更されます。
// 数字は次の形式でなければなりません:
//
//	number    = [ sign ] ( float | "inf" | "Inf" ) .
//	sign      = "+" | "-" .
//	float     = ( mantissa | prefix pmantissa ) [ exponent ] .
//	prefix    = "0" [ "b" | "B" | "o" | "O" | "x" | "X" ] .
//	mantissa  = digits "." [ digits ] | digits | "." digits .
//	pmantissa = [ "_" ] digits "." [ digits ] | [ "_" ] digits | "." digits .
//	exponent  = ( "e" | "E" | "p" | "P" ) [ sign ] digits .
//	digits    = digit { [ "_" ] digit } .
//	digit     = "0" ... "9" | "a" ... "z" | "A" ... "Z" .
//
// 基数引数は0、2、8、10、または16でなければなりません。無効な基数引数を提供すると、
// 実行時にパニックが発生します。
//
// 基数が0の場合、数値のプレフィックスが実際の基数を決定します：プレフィックスが
// "0b"または"0B"は基数2を選択し、"0o"または"0O"は基数8を選択し、
// "0x"または"0X"は基数16を選択します。それ以外の場合、実際の基数は10であり、
// プレフィックスは受け入れられません。8進数のプレフィックス"0"はサポートされていません（先頭の
// "0"は単に"0"と見なされます）。
//
// "p"または"P"の指数は、基数10ではなく基数2の指数を示します。
// 例えば、"0x1.fffffffffffffp1023"（基数0を使用）は、最大のfloat64値を表します。
// 16進数の仮数については、指数文字が存在する場合、'p'または'P'のいずれかでなければなりません
// （"e"または"E"の指数指示子は、仮数の数字と区別できません）。
//
// エラーが報告された場合、返される*Float fはnilで、zの値は有効ですが定義されていません。
func (z *Float) Parse(s string, base int) (f *Float, b int, err error)

// ParseFloatは、指定された精度と丸めモードでfを設定した状態のf.Parse(s, base)と同じです。
func ParseFloat(s string, base int, prec uint, mode RoundingMode) (f *Float, b int, err error)

var _ fmt.Scanner = (*Float)(nil)

<<<<<<< HEAD
// Scan is a support routine for [fmt.Scanner]; it sets z to the value of
// the scanned number. It accepts formats whose verbs are supported by
// [fmt.Scan] for floating point values, which are:
// 'b' (binary), 'e', 'E', 'f', 'F', 'g' and 'G'.
// Scan doesn't handle ±Inf.
=======
// Scanは、fmt.Scannerのサポートルーチンで、zをスキャンされた数値に設定します。
// これは、浮動小数点値に対してfmt.Scanがサポートする動詞を持つ形式を受け入れます。それらは次のとおりです:
// 'b'（バイナリ）、'e'、'E'、'f'、'F'、'g'、'G'。
// Scanは±Infを処理しません。
>>>>>>> release-branch.go1.21
func (z *Float) Scan(s fmt.ScanState, ch rune) error
