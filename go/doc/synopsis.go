// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

// Synopsisはtextの最初の文のクリーニングされたバージョンを返します。
//
// Deprecated: 新しいプログラムは代わりに[Package.Synopsis]を使用するべきです。
// これはテキスト内のリンクを正しく扱います。
func Synopsis(text string) string

// IllegalPrefixesは、ドキュメントコメントではないコメントを識別するための小文字の接頭辞のリストです。
// これにより、パッケージ文の直前にある著作権表示のよくある間違いをドキュメントコメントと誤解しないようにします。
var IllegalPrefixes = []string{
	"copyright",
	"all rights",
	"author",
}

// Synopsisは、テキスト内の最初の文のクリーニングされたバージョンを返します。
// その文は、最初のピリオドの後に空白が続き、正確に1つの大文字で始まっていない、または最初の段落区切りの直前に終わります。
// 結果の文字列には、\n、\r、または\tの文字がなく、単語間には単一のスペースのみが使用されます。
// テキストが [IllegalPrefixes] のいずれかで始まる場合、結果は空の文字列です。
func (p *Package) Synopsis(text string) string
