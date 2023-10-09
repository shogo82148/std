// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

// Printerはドキュメントコメントのプリンターです。
// 構造体のフィールドは、印刷の詳細をカスタマイズするために、
// 印刷メソッドのいずれかを呼び出す前に埋めることができます。
type Printer struct {

	// HeadingLevelはHTMLとMarkdownの見出しに使用されるネストレベルです。
	// HeadingLevelがゼロであれば、デフォルトでレベル3に設定され、<h3>と###が使用されます。
	HeadingLevel int

	// HeadingIDは、HTMLとMarkdownを生成する際に使用する見出しhの見出しID（アンカータグ）を計算する関数です。HeadingIDが空の文字列を返す場合、見出しIDは省略されます。HeadingIDがnilの場合、h.DefaultIDが使用されます。
	HeadingID func(h *Heading) string

	// DocLinkURLは、与えられたDocLinkのURLを計算する関数です。
	// DocLinkURLがnilの場合、link.DefaultURL（p.DocLinkBaseURL）が使用されます。
	DocLinkURL func(link *DocLink) string

	// DocLinkBaseURLは、DocLinkURLがnilの場合に使用され、DocLinkのURLを構築するために[DocLink.DefaultURL]に渡されます。
	// 詳細については、そのメソッドのドキュメントを参照してください。
	DocLinkBaseURL string

	// TextPrefixは、Textメソッドを使用してテキスト出力を生成する際に、
	// 各行の先頭に表示するプレフィックスです。
	TextPrefix string

	// TextCodePrefixは、テキスト出力を生成する際に各事前に書式設定された（コードブロック）行の先頭に印刷する接頭辞です。
	// （TextPrefixに追加ではなく）。
	// TextCodePrefixが空の文字列の場合、TextPrefix +"\t"がデフォルト値となります。
	TextCodePrefix string

	// TextWidthは生成するテキスト行の最大幅であり、Unicodeのコードポイントで測定されます。
	// TextPrefixと改行文字を除いたものです。
	// TextWidthがゼロの場合、TextPrefixのコードポイント数を除いた80から始まります。
	// TextWidthが負の場合、制限はありません。
	TextWidth int
}

// DefaultURLは、baseURLをプレフィックスとして他のパッケージへのリンクのためにlのドキュメンテーションURLを構築して返します。
//
// DefaultURLが返す可能性のある形式は以下の通りです：
//   - baseURL/ImportPath、他のパッケージへのリンク
//   - baseURL/ImportPath#Name、他のパッケージのconst、func、type、またはvarへのリンク
//   - baseURL/ImportPath#Recv.Name、他のパッケージのメソッドへのリンク
//   - #Name、このパッケージのconst、func、type、またはvarへのリンク
//   - #Recv.Name、このパッケージのメソッドへのリンク
//
// baseURLが末尾にスラッシュで終わる場合、アンカー形式のImportPathと#の間にスラッシュが挿入されます。
// 例えば、以下はいくつかのbaseURL値とそれらが生成するURLの例です：
//
// "/pkg/" → "/pkg/math/#Sqrt"
// "/pkg"  → "/pkg/math#Sqrt"
// "/"     → "/math/#Sqrt"
// ""      → "/math#Sqrt"
func (l *DocLink) DefaultURL(baseURL string) string

// DefaultIDは見出しhのデフォルトのアンカーIDを返します。
//
// デフォルトのアンカーIDは、全てのアルファベットや数字以外のASCIIのランを
// アンダースコアに変換し、その前に「hdr-」の接頭辞を付けることで構築されます。
// 例えば、見出しのテキストが「Go Doc Comments」の場合、デフォルトのIDは「hdr-Go_Doc_Comments」となります。
func (h *Heading) DefaultID() string

// Commentは、コメントマーカーなしでのDocの標準的なGoのフォーマットを返します。
func (p *Printer) Comment(d *Doc) []byte
