// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/iter"
)

// Linesは、バイトスライスs内の改行で終わる行に対するイテレータを返します。
// イテレータが生成する行には、終端の改行が含まれます。
// sが空の場合、イテレータは行をまったく生成しません。
// sが改行で終わらない場合、最後に生成される行は改行で終わりません。
// これは使い捨てのイテレータを返します。
func Lines(s []byte) iter.Seq[[]byte]

// SplitSeqは、sepで区切られたsのすべてのサブスライスに対するイテレータを返します。
// イテレータは [Split](s, sep)が返すのと同じサブスライスを生成しますが、
// サブスライスを含む新しいスライスを構築しません。
// これは使い捨てのイテレータを返します。
func SplitSeq(s, sep []byte) iter.Seq[[]byte]

// SplitAfterSeqは、sepの各出現の後で分割されたsのサブスライスに対するイテレータを返します。
// イテレータは [SplitAfter](s, sep)が返すのと同じサブスライスを生成しますが、
// サブスライスを含む新しいスライスを構築しません。
// これは使い捨てのイテレータを返します。
func SplitAfterSeq(s, sep []byte) iter.Seq[[]byte]

// FieldsSeqは、[unicode.IsSpace] で定義される連続する空白文字の周りで分割された
// sのサブスライスに対するイテレータを返します。
// イテレータは [Fields](s)が返すのと同じサブスライスを生成しますが、
// サブスライスを含む新しいスライスを構築しません。
func FieldsSeq(s []byte) iter.Seq[[]byte]

// FieldsFuncSeqは、f(c)を満たす連続するUnicodeコードポイントの周りで分割された
// sのサブスライスに対するイテレータを返します。
// イテレータは [FieldsFunc](s)が返すのと同じサブスライスを生成しますが、
// サブスライスを含む新しいスライスを構築しません。
func FieldsFuncSeq(s []byte, f func(rune) bool) iter.Seq[[]byte]
