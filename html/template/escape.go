// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/io"
)

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

// URLQueryEscaperは、その引数のテキスト表現のエスケープされた値を、
// URLクエリに埋め込むのに適した形式で返します。
func URLQueryEscaper(args ...any) string
