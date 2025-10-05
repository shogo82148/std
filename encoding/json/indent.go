// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !goexperiment.jsonv2

package json

import "github.com/shogo82148/std/bytes"

// HTMLEscapeは、JSONエンコードされたsrcをdstに追加します。ただし、文字列リテラル内の<, >, &, U+2028およびU+2029
// 文字は\u003c, \u003e, \u0026, \u2028, \u2029に変更されます。
// これにより、JSONはHTMLの<script>タグ内に埋め込むのに安全になります。
// 歴史的な理由から、ウェブブラウザは<script>タグ内で標準的なHTMLエスケープを尊重しないため、
// 代替のJSONエンコーディングを使用する必要があります。
func HTMLEscape(dst *bytes.Buffer, src []byte)

// Compactは、無意味なスペース文字が省略された、JSONエンコードされたsrcをdstに追加します。
func Compact(dst *bytes.Buffer, src []byte) error

// Indentは、JSONエンコードされたsrcのインデント形式をdstに追加します。
// JSONオブジェクトまたは配列の各要素は、新しい、インデントされた行で始まり、
// その行はprefixで始まり、インデントのネストに応じてindentの1つ以上のコピーが続きます。
// dstに追加されるデータは、prefixやインデントで始まらないようになっています。
// これにより、他のフォーマットされたJSONデータ内に埋め込むのが容易になります。
// srcの先頭の空白文字（スペース、タブ、キャリッジリターン、改行）は削除されますが、
// srcの末尾の空白文字は保持され、dstにコピーされます。
// 例えば、srcが末尾のスペースを持っていない場合、dstも持ちません。
// srcが末尾の改行で終わる場合、dstも同様になります。
func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
