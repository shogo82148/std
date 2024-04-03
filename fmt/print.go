// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt

import (
	"github.com/shogo82148/std/io"
)

// Stateはカスタムフォーマッタに渡されるプリンタの状態を表します。
// [io.Writer] インターフェースへのアクセスと、オペランドのフォーマット指定子に関するフラグとオプションの情報を提供します。
type State interface {
	// Writeは、出力をフォーマットして印刷するために呼び出す関数です。
	Write(b []byte) (n int, err error)
	// Widthは、幅オプションの値とその設定状態を返します。
	Width() (wid int, ok bool)
	// Precisionは、精度オプションの値とその設定状態を返します。
	Precision() (prec int, ok bool)

	// Flagは、文字cが設定されているかどうかを報告します。
	Flag(c int) bool
}

// Formatterは、Formatメソッドを持つ任意の値で実装されます。
// 実装は、[State] とruneの解釈方法を制御し、[Sprint] や [Fprint](f)などを呼び出して出力を生成することができます。
type Formatter interface {
	Format(f State, verb rune)
}

// Stringerは、Stringメソッドを持つ任意の値によって実装されます。
// このメソッドは、その値の「ネイティブ」フォーマットを定義します。
// Stringメソッドは、文字列を受け入れる任意のフォーマットや、
// [Print] のような書式のないプリンターにオペランドとして渡される値を表示するために使用されます。
type Stringer interface {
	String() string
}

// GoStringerは、GoStringメソッドを持つ任意の値で実装されており、
// その値のGo構文を定義します。
// GoStringメソッドは、%#vフォーマットに渡された値を出力するために使用されます。
type GoStringer interface {
	GoString() string
}

// FormatStringは、[State] によってキャプチャされた完全修飾のフォーマットディレクティブを表す文字列を返し、その後に引数の動詞が続きます。（[State] 自体は動詞を含みません。）結果には、先頭にパーセント記号があり、その後にフラグ、幅、および精度が続きます。フラグ、幅、および精度のない場合は省略されます。この関数により、[Formatter] はFormatへの呼び出しをトリガーした元のディレクティブを再構築することができます。
func FormatString(state State, verb rune) string

// Fprintfはフォーマット指定子に従って書式を整え、wに書き込みます。
// 書き込んだバイト数とエラーの有無を返します。
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)

// Printfはフォーマット指定子に従ってフォーマットを行い、標準出力に書き込みます。
// 書き込まれたバイト数とエラーがあれば、それらを返します。
func Printf(format string, a ...any) (n int, err error)

// Sprintfはフォーマット指定子に従ってフォーマットされた文字列を返します。
func Sprintf(format string, a ...any) string

// Appendfはフォーマット指定子に従ってフォーマットを行い、結果をバイトスライスに追加し、更新されたスライスを返します。
func Appendf(b []byte, format string, a ...any) []byte

// Fprint はオペランドのデフォルトのフォーマットを使用してwに書き込みます。
// オペランドが文字列でない場合、間にスペースが追加されます。
// 書き込まれたバイト数と発生した書き込みエラーが返されます。
func Fprint(w io.Writer, a ...any) (n int, err error)

// デフォルトの書式を使用して、オペランドの内容を表示し、標準出力に書き込みます。
// オペランドがどちらも文字列でない場合は、オペランド間にスペースが追加されます。
// 書き込まれたバイト数と発生した書き込みエラーが返されます。
func Print(a ...any) (n int, err error)

// Sprintは、オペランドのデフォルトのフォーマットを使用して結果の文字列を返します。
// オペランドがどちらも文字列でない場合、オペランド間にスペースが追加されます。
func Sprint(a ...any) string

// Appendはオペランドのデフォルトフォーマットを使用して、フォーマットを行い、結果をバイトスライスに追加し、更新されたスライスを返します。
func Append(b []byte, a ...any) []byte

// Fprintln はオペランドのデフォルト形式を使用してフォーマットし、w に書き込みます。
// オペランドの間には常にスペースが追加され、改行が追加されます。
// 書き込まれたバイト数と発生した書き込みエラーを返します。
func Fprintln(w io.Writer, a ...any) (n int, err error)

// Printlnは、オペランドのデフォルトのフォーマットを使用して整形し、標準出力に書き込みます。
// オペランドとオペランドの間には常にスペースが追加され、改行が追加されます。
// 書き込まれたバイト数とエラーの有無を返します。
func Println(a ...any) (n int, err error)

// Sprintln はオペランドのデフォルトの書式を使用してフォーマットし、結果の文字列を返します。
// オペランド間には常にスペースが追加され、改行が追加されます。
func Sprintln(a ...any) string

// Appendln はオペランドのデフォルトの書式を使用してフォーマットし、結果をバイトスライスに追加し、更新されたスライスを返します。オペランド間には常にスペースが追加され、改行が追加されます。
func Appendln(b []byte, a ...any) []byte
