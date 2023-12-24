// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements int-to-string conversion functions.

package big

import (
	"github.com/shogo82148/std/fmt"
)

// Textは、指定された基数でのxの文字列表現を返します。
// 基数は2から62までの間でなければなりません。結果は、
// 数字の値10から35に対して小文字の'a'から'z'を、
// 数字の値36から61に対して大文字の'A'から'Z'を使用します。
// 文字列にはプレフィックス（例えば"0x"）は追加されません。xがnilポインタの場合、
// "<nil>"を返します。
func (x *Int) Text(base int) string

// Appendは、x.Text(base)によって生成されたxの文字列表現をbufに追加し、
// 拡張されたバッファを返します。
func (x *Int) Append(buf []byte, base int) []byte

// Stringは、x.Text(10)によって生成されるxの10進表現を返します。
func (x *Int) String() string

var _ fmt.Formatter = intOne

<<<<<<< HEAD
// Format implements [fmt.Formatter]. It accepts the formats
// 'b' (binary), 'o' (octal with 0 prefix), 'O' (octal with 0o prefix),
// 'd' (decimal), 'x' (lowercase hexadecimal), and
// 'X' (uppercase hexadecimal).
// Also supported are the full suite of package fmt's format
// flags for integral types, including '+' and ' ' for sign
// control, '#' for leading zero in octal and for hexadecimal,
// a leading "0x" or "0X" for "%#x" and "%#X" respectively,
// specification of minimum digits precision, output field
// width, space or zero padding, and '-' for left or right
// justification.
=======
// Formatは、fmt.Formatterを実装します。次の形式を受け入れます
// 'b'（二進数）、'o'（0接頭辞付きの8進数）、'O'（0o接頭辞付きの8進数）、
// 'd'（10進数）、'x'（小文字の16進数）、そして
// 'X'（大文字の16進数）。
// また、符号制御のための'+'と' '、8進数の先頭ゼロと16進数のための'#'、
// "%#x"と"%#X"に対する先頭の"0x"または"0X"、最小桁数の精度の指定、出力フィールド
// 幅、スペースまたはゼロパディング、そして左または右
// 寄せのための'-'を含む、パッケージfmtの整数型のための完全な形式
// フラグもサポートされています。
>>>>>>> release-branch.go1.21
func (x *Int) Format(s fmt.State, ch rune)

var _ fmt.Scanner = intOne

<<<<<<< HEAD
// Scan is a support routine for [fmt.Scanner]; it sets z to the value of
// the scanned number. It accepts the formats 'b' (binary), 'o' (octal),
// 'd' (decimal), 'x' (lowercase hexadecimal), and 'X' (uppercase hexadecimal).
=======
// Scanは、fmt.Scannerのサポートルーチンであり、zをスキャンされた数値に設定します。
// 形式 'b'（二進数）、'o'（8進数）、'd'（10進数）、'x'（小文字の16進数）、'X'（大文字の16進数）を受け入れます。
>>>>>>> release-branch.go1.21
func (z *Int) Scan(s fmt.ScanState, ch rune) error
