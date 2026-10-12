// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/shogo82148/std/iter"
)

// Linesは、文字列s内の改行で終わる行に対するイテレーターを返します。
// イテレーターが生成する行には、終端の改行が含まれます。
// sが空の場合、イテレーターは行をまったく生成しません。
// sが改行で終わらない場合、最後に生成される行は改行で終わりません。
// これは使い捨てのイテレーターを返します。
func Lines(s string) iter.Seq[string]

// SplitSeqは、sepで区切られたsのすべての部分文字列に対するイテレーターを返します。
// イテレーターは [Split](s, sep)が返すのと同じ文字列を生成しますが、スライスを構築しません。
// これは使い捨てのイテレーターを返します。
func SplitSeq(s, sep string) iter.Seq[string]

// SplitAfterSeqは、sepの各出現の後で分割されたsの部分文字列に対するイテレーターを返します。
// イテレーターは [SplitAfter](s, sep)が返すのと同じ文字列を生成しますが、スライスを構築しません。
// これは使い捨てのイテレーターを返します。
func SplitAfterSeq(s, sep string) iter.Seq[string]

// FieldsSeqは、[unicode.IsSpace] で定義される連続する空白文字の周りで分割された
// sの部分文字列に対するイテレーターを返します。
// イテレーターは [Fields](s)が返すのと同じ文字列を生成しますが、スライスを構築しません。
func FieldsSeq(s string) iter.Seq[string]

// FieldsFuncSeqは、f(c)を満たす連続するUnicodeコードポイントの周りで分割された
// sの部分文字列に対するイテレーターを返します。
// イテレーターは [FieldsFunc](s)が返すのと同じ文字列を生成しますが、スライスを構築しません。
func FieldsFuncSeq(s string, f func(rune) bool) iter.Seq[string]
