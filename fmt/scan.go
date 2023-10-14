// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt

import (
	"github.com/shogo82148/std/io"
)

// ScanStateはカスタムスキャナーに渡されるスキャナーの状態を表します。
// スキャナーは一文字ずつスキャンすることもでき、またScanStateに次のスペース区切りのトークンを見つけるように依頼することもできます。
type ScanState interface {
	// ReadRune reads the next rune (Unicode code point) from the input.
	// If invoked during Scanln, Fscanln, or Sscanln, ReadRune() will
	// return EOF after returning the first '\n' or when reading beyond
	// the specified width.
	ReadRune() (r rune, size int, err error)
	// UnreadRune causes the next call to ReadRune to return the same rune.
	UnreadRune() error
	// SkipSpace skips space in the input. Newlines are treated appropriately
	// for the operation being performed; see the package documentation
	// for more information.
	SkipSpace()
	// Token skips space in the input if skipSpace is true, then returns the
	// run of Unicode code points c satisfying f(c).  If f is nil,
	// !unicode.IsSpace(c) is used; that is, the token will hold non-space
	// characters. Newlines are treated appropriately for the operation being
	// performed; see the package documentation for more information.
	// The returned slice points to shared data that may be overwritten
	// by the next call to Token, a call to a Scan function using the ScanState
	// as input, or when the calling Scan method returns.
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	// Width returns the value of the width option and whether it has been set.
	// The unit is Unicode code points.
	Width() (wid int, ok bool)
	// Because ReadRune is implemented by the interface, Read should never be
	// called by the scanning routines and a valid implementation of
	// ScanState may choose always to return an error from Read.
	Read(buf []byte) (n int, err error)
}

// Scannerは、値のスキャンメソッドを持つ任意の値によって実装されており、
// 入力を値の表現形式でスキャンし、結果をレシーバに格納します。
// 有用にするために、レシーバはポインタでなければなりません。
// スキャンメソッドは、Scan、Scanf、またはScanlnの引数として実装するものです。
type Scanner interface {
	Scan(state ScanState, verb rune) error
}

// Scanは標準入力から読み取ったテキストをスキャンし、連続するスペースで区切られた値を連続した引数に格納します。改行はスペースとして扱われます。スキャンに成功したアイテムの数を返します。引数の数よりも少ない場合、errにはエラーの理由が報告されます。
func Scan(a ...any) (n int, err error)

// ScanlnはScanに似ていますが、改行でスキャンを停止し、
// 最後のアイテムの後には改行またはEOFが必要です。
func Scanln(a ...any) (n int, err error)

// Scanfは標準入力から読み取ったテキストをスキャンし、形式に応じて連続したスペースで区切られた値を連続した引数に保存します。成功したスキャンのアイテム数を返します。引数の数よりも少ない場合、エラーが発生した理由がerrに報告されます。入力にある改行は、形式にある改行と一致する必要があります。ただし例外として、動詞%cは常に入力の次のルーンをスキャンします。それがスペース（またはタブなど）や改行であってもです。
func Scanf(format string, a ...any) (n int, err error)

// Sscanは引数の文字列をスキャンし、連続するスペースで区切られた値を連続する引数に格納します。改行もスペースとして扱われます。正常にスキャンできたアイテムの数を返します。もし正常にスキャンされたアイテムの数が引数の数よりも少ない場合、errがその理由を報告します。
func Sscan(str string, a ...any) (n int, err error)

// SscanlnはSscanに似ていますが、改行でスキャンを停止し、最後のアイテムの後には改行またはEOFが必要です。
func Sscanln(str string, a ...any) (n int, err error)

// Sscanfは引数文字列をスキャンし、フォーマットによって決まる連続するスペースで区切られた値を連続した引数に格納します。正常に解析されたアイテムの数を返します。
// 入力の改行は、フォーマットと一致する必要があります。
func Sscanf(str string, format string, a ...any) (n int, err error)

// Fscan は、r から読み取ったテキストをスキャンし、連続した空白で区切られた値を連続した引数に格納します。改行も空白としてカウントされます。成功したスキャンのアイテム数を返します。引数の数よりも少ない場合、err がなぜエラーが発生したのかを報告します。
func Fscan(r io.Reader, a ...any) (n int, err error)

// FscanlnはFscanに似ていますが、改行でスキャンを終了し、最後の項目の後には改行かEOFが必要です。
func Fscanln(r io.Reader, a ...any) (n int, err error)

// Fscanfはrから読み取ったテキストをスキャンし、
// フォーマットに従って連続したスペースで区切られた値を連続した引数に格納します。
// 成功した解析のアイテム数を返します。
// 入力の改行はフォーマットの改行と一致する必要があります。
func Fscanf(r io.Reader, format string, a ...any) (n int, err error)
