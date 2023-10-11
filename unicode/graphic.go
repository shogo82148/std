// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unicode

// GraphicRangesはUnicodeに基づいてグラフィック文字のセットを定義します。
var GraphicRanges = []*RangeTable{
	L, M, N, P, S, Zs,
}

// PrintRangesはGoによる印刷可能な文字のセットを定義します。
// ASCIIスペース、U+0020は別途扱われます。
var PrintRanges = []*RangeTable{
	L, M, N, P, S,
}

// IsGraphicはUnicodeによってグラフィックとして定義されたルーンかどうかを報告します。
// これには、カテゴリL、M、N、P、S、Zsの文字が含まれます。
func IsGraphic(r rune) bool

// IsPrintは、Goによって印字可能として定義されているルーンかどうかを報告します。これには文字、マーク、数字、句読点、記号、およびASCIIスペース文字が含まれます。これはカテゴリL、M、N、P、S、およびASCIIスペース文字と同じ分類です（ただし、唯一のスペース文字はASCIIスペース、U+0020です）。IsGraphicとは異なり、この区分にはASCIIスペース文字のみが含まれています。
func IsPrint(r rune) bool

// IsOneOfは、ルーンがいずれかの範囲のメンバーであるかどうかを報告します。
// 関数"In"はより良いシグネチャを提供し、IsOneOfよりも使われるべきです。
func IsOneOf(ranges []*RangeTable, r rune) bool

// ランジのいずれかのメンバーかどうかを報告する。
func In(r rune, ranges ...*RangeTable) bool

// IsControlはルーンが制御文字であるかどうかを報告します。
// C (その他)のUnicodeカテゴリにはサロゲートなど、より多くのコードポイントが含まれています。
// それらをテストするにはIs(C, r)を使用してください。
func IsControl(r rune) bool

// IsLetterはルーンが文字（カテゴリーL）であるかどうかを報告します。
func IsLetter(r rune) bool

// IsMarkは、ルーンがマーク文字（カテゴリM）であるかどうかを報告します。
func IsMark(r rune) bool

// IsNumberはルーンが数字（カテゴリーN）であるかどうかを報告します。
func IsNumber(r rune) bool

// IsPunctはruneがUnicodeの句読点（カテゴリP）であるかどうかを報告します。
func IsPunct(r rune) bool

// IsSpaceは、UnicodeのWhite Spaceプロパティによって定義された空白文字であるかどうかを報告します。
// これには、Latin-1スペースに次の文字が含まれます。
//
// '\t'、'\n'、'\v'、'\f'、'\r'、' '、U+0085（NEL）、U+00A0（NBSP）。
//
// スペーシング文字の他の定義は、カテゴリZおよびプロパティPattern_White_Spaceによって設定されています。
func IsSpace(r rune) bool

// IsSymbolはルーンが記号の文字であるかどうかを報告します。
func IsSymbol(r rune) bool
