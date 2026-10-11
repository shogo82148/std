// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/bytes"
)

// HTMLEscapeは、文字列リテラル内の<、>、&、U+2028、U+2029文字を
// \u003c、\u003e、\u0026、\u2028、\u2029に変更したJSONエンコードされたsrcをdstに追加します。
// これにより、JSONをHTMLの<script>タグ内に安全に埋め込むことができます。
// 歴史的な理由により、Webブラウザは<script>タグ内で標準のHTMLエスケープを
// 尊重しないため、代替のJSONエンコーディングを使用する必要があります。
func HTMLEscape(dst *bytes.Buffer, src []byte)

// Compactは、重要でない空白文字を省いたJSONエンコードされたsrcをdstに追加します。
func Compact(dst *bytes.Buffer, src []byte) error

// Indentは、JSONエンコードされたsrcのインデント形式をdstに追加します。
// JSONオブジェクトまたは配列の各要素は、新しいインデントされた行で始まり、
// その行はprefixで始まり、その後にインデントのネストに応じて
// indentの1つ以上のコピーが続きます。
// 他のフォーマットされたJSONデータに埋め込みやすくするために、dstに追加されるデータは
// prefixやインデントで始まりません。
// src先頭の空白文字（スペース、タブ、キャリッジリターン、改行）は削除されますが、
// src末尾の空白文字は保持され、dstにコピーされます。
// 例えば、srcに末尾の空白がない場合、dstにもありません。
// srcが末尾の改行で終わる場合、dstも同様です。
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
