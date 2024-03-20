// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/io"
)

// FuncMapは、名前から関数へのマッピングを定義するマップの型です。
// 各関数は、単一の戻り値を持つか、または二つの戻り値を持つ必要があり、
// その二つ目の型はエラーです。その場合、もし二つ目の（エラー）戻り値が実行中に非nilに評価された場合、
// 実行は終了し、Executeはそのエラーを返します。
//
// Executeによって返されるエラーは、基礎となるエラーをラップします。それらをアンラップするには、[errors.As] を呼び出します。
//
// テンプレートの実行が引数リストを持つ関数を呼び出すとき、そのリストは
// 関数のパラメータタイプに割り当て可能でなければなりません。任意のタイプの引数に適用することを意図した
// 関数は、interface{}型または [reflect.Value] 型のパラメータを使用できます。同様に、任意の
// タイプの結果を返すことを意図した関数は、interface{}または [reflect.Value] を返すことができます。
type FuncMap map[string]any

// HTMLEscapeは、プレーンテキストデータbのエスケープされたHTML相当をwに書き込みます。
func HTMLEscape(w io.Writer, b []byte)

// HTMLEscapeStringは、プレーンテキストデータsのエスケープされたHTML相当を返します。
func HTMLEscapeString(s string) string

// HTMLEscaperは、その引数のテキスト表現のエスケープされたHTML相当を返します。
func HTMLEscaper(args ...any) string

// JSEscapeは、プレーンテキストデータbのエスケープされたJavaScript相当をwに書き込みます。
func JSEscape(w io.Writer, b []byte)

// JSEscapeStringは、プレーンテキストデータsのエスケープされたJavaScript相当を返します。
func JSEscapeString(s string) string

// JSEscaperは、その引数のテキスト表現のエスケープされたJavaScript相当を返します。
func JSEscaper(args ...any) string

// URLQueryEscaperは、URLクエリに埋め込む形式に適した、その引数のテキスト表現のエスケープされた値を返します。
func URLQueryEscaper(args ...any) string
