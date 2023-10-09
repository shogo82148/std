// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

import (
	"github.com/shogo82148/std/io"
)

// ToHTMLはコメントテキストをフォーマットされたHTMLに変換します。
//
// Deprecated: ToHTMLはドキュメントリンクを識別できません
// ドキュメントコメント内のリンクは、テキストがパッケージから取得された内容を知っている必要があるため、このAPIには含まれていません。
//
// *[doc.Package] pがテキストが見つかった場所で見つかる場合、
// ToHTML(w, text, nil)は以下のように置き換えられます：
//
//	w.Write(p.HTML(text))
//
// これは次の省略形です：
//
//	w.Write(p.Printer().HTML(p.Parser().Parse(text)))
//
// wordsがnilでない場合、より長い置換は次のとおりです：
//
//	parser := p.Parser()
//	parser.Words = words
//	w.Write(p.Printer().HTML(parser.Parse(d)))
func ToHTML(w io.Writer, text string, words map[string]string)

// ToTextはコメントテキストを整形されたテキストに変換します。
//
// Deprecated: ToTextはドキュメントリンクを識別できません。
// ドキュメントリンクはテキストが含まれるパッケージを知る必要があるため、
// このAPIには含まれていません。
//
// *[doc.Package] pでテキストが見つかった場合、
// ToText(w, text, "", "\t", 80)は次のように置き換えられます:
//
//	w.Write(p.Text(text))
//
// 一般的な場合、ToText(w, text, prefix, codePrefix, width)は次のように置き換えられます:
//
//	d := p.Parser().Parse(text)
//	pr := p.Printer()
//	pr.TextPrefix = prefix
//	pr.TextCodePrefix = codePrefix
//	pr.TextWidth = width
//	w.Write(pr.Text(d))
//
// 詳細については、[Package.Text] と [comment.Printer.Text] のドキュメントを参照してください。
func ToText(w io.Writer, text string, prefix, codePrefix string, width int)
