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
	// ReadRuneは、入力から次のルーン（Unicodeコードポイント）を読み取ります。
	// Scanln、Fscanln、またはSscanln中に呼び出された場合、ReadRune()は、最初の'\n'を返した後、または指定された幅を超えて読み取りを行った後にEOFを返します。
	ReadRune() (r rune, size int, err error)
	// UnreadRuneは、次のReadRune呼び出しで同じルーンを返します。
	UnreadRune() error
	// SkipSpaceは、入力内のスペースをスキップします。
	// 操作に応じて、改行は適切に処理されます。
	// 詳細については、パッケージのドキュメントを参照してください。
	SkipSpace()
	// Tokenは、skipSpaceがtrueの場合、入力内のスペースをスキップして、f(c)を満たすUnicodeコードポイントcのランを返します。
	// fがnilの場合、!unicode.IsSpace(c)が使用されます。つまり、トークンにはスペース以外の文字が含まれます。
	// 操作に応じて、改行は適切に処理されます。詳細については、パッケージのドキュメントを参照してください。
	// 返されたスライスは、Tokenの次の呼び出し、ScanStateを入力として使用するScan関数の呼び出し、または呼び出し元のScanメソッドが返されたときに上書きされる可能性がある共有データを指します。
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	// Widthは、幅オプションの値とその設定状態を返します。
	// 単位はUnicodeコードポイントです。
	Width() (wid int, ok bool)
	// ReadRuneはインターフェースによって実装されているため、スキャンルーチンからReadが呼び出されることはありません。
	// また、ScanStateの有効な実装は、常にReadからエラーを返すことがあります。
	Read(buf []byte) (n int, err error)
}

<<<<<<< HEAD
// Scannerは、値のスキャンメソッドを持つ任意の値によって実装されており、
// 入力を値の表現形式でスキャンし、結果をレシーバに格納します。
// 有用にするために、レシーバはポインタでなければなりません。
// スキャンメソッドは、Scan、Scanf、またはScanlnの引数として実装するものです。
=======
// Scanner is implemented by any value that has a Scan method, which scans
// the input for the representation of a value and stores the result in the
// receiver, which must be a pointer to be useful. The Scan method is called
// for any argument to [Scan], [Scanf], or [Scanln] that implements it.
>>>>>>> upstream/master
type Scanner interface {
	Scan(state ScanState, verb rune) error
}

// Scanは標準入力から読み取ったテキストをスキャンし、連続するスペースで区切られた値を連続した引数に格納します。改行はスペースとして扱われます。スキャンに成功したアイテムの数を返します。引数の数よりも少ない場合、errにはエラーの理由が報告されます。
func Scan(a ...any) (n int, err error)

<<<<<<< HEAD
// ScanlnはScanに似ていますが、改行でスキャンを停止し、
// 最後のアイテムの後には改行またはEOFが必要です。
=======
// Scanln is similar to [Scan], but stops scanning at a newline and
// after the final item there must be a newline or EOF.
>>>>>>> upstream/master
func Scanln(a ...any) (n int, err error)

// Scanfは標準入力から読み取ったテキストをスキャンし、形式に応じて連続したスペースで区切られた値を連続した引数に保存します。成功したスキャンのアイテム数を返します。引数の数よりも少ない場合、エラーが発生した理由がerrに報告されます。入力にある改行は、形式にある改行と一致する必要があります。ただし例外として、動詞%cは常に入力の次のルーンをスキャンします。それがスペース（またはタブなど）や改行であってもです。
func Scanf(format string, a ...any) (n int, err error)

// Sscanは引数の文字列をスキャンし、連続するスペースで区切られた値を連続する引数に格納します。改行もスペースとして扱われます。正常にスキャンできたアイテムの数を返します。もし正常にスキャンされたアイテムの数が引数の数よりも少ない場合、errがその理由を報告します。
func Sscan(str string, a ...any) (n int, err error)

<<<<<<< HEAD
// SscanlnはSscanに似ていますが、改行でスキャンを停止し、最後のアイテムの後には改行またはEOFが必要です。
=======
// Sscanln is similar to [Sscan], but stops scanning at a newline and
// after the final item there must be a newline or EOF.
>>>>>>> upstream/master
func Sscanln(str string, a ...any) (n int, err error)

// Sscanfは引数文字列をスキャンし、フォーマットによって決まる連続するスペースで区切られた値を連続した引数に格納します。正常に解析されたアイテムの数を返します。
// 入力の改行は、フォーマットと一致する必要があります。
func Sscanf(str string, format string, a ...any) (n int, err error)

// Fscan は、r から読み取ったテキストをスキャンし、連続した空白で区切られた値を連続した引数に格納します。改行も空白としてカウントされます。成功したスキャンのアイテム数を返します。引数の数よりも少ない場合、err がなぜエラーが発生したのかを報告します。
func Fscan(r io.Reader, a ...any) (n int, err error)

<<<<<<< HEAD
// FscanlnはFscanに似ていますが、改行でスキャンを終了し、最後の項目の後には改行かEOFが必要です。
=======
// Fscanln is similar to [Fscan], but stops scanning at a newline and
// after the final item there must be a newline or EOF.
>>>>>>> upstream/master
func Fscanln(r io.Reader, a ...any) (n int, err error)

// Fscanfはrから読み取ったテキストをスキャンし、
// フォーマットに従って連続したスペースで区切られた値を連続した引数に格納します。
// 成功した解析のアイテム数を返します。
// 入力の改行はフォーマットの改行と一致する必要があります。
func Fscanf(r io.Reader, format string, a ...any) (n int, err error)
