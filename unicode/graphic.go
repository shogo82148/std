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

<<<<<<< HEAD
// IsGraphicはUnicodeによってグラフィックとして定義されたルーンかどうかを報告します。
// これには、カテゴリL、M、N、P、S、Zsの文字が含まれます。
func IsGraphic(r rune) bool

// IsPrintは、Goによって印字可能として定義されているルーンかどうかを報告します。これには文字、マーク、数字、句読点、記号、およびASCIIスペース文字が含まれます。これはカテゴリL、M、N、P、S、およびASCIIスペース文字と同じ分類です（ただし、唯一のスペース文字はASCIIスペース、U+0020です）。IsGraphicとは異なり、この区分にはASCIIスペース文字のみが含まれています。
=======
// IsGraphic reports whether the rune is defined as a Graphic by Unicode.
// Such characters include letters, marks, numbers, punctuation, symbols, and
// spaces, from categories [L], [M], [N], [P], [S], [Zs].
func IsGraphic(r rune) bool

// IsPrint reports whether the rune is defined as printable by Go. Such
// characters include letters, marks, numbers, punctuation, symbols, and the
// ASCII space character, from categories [L], [M], [N], [P], [S] and the ASCII space
// character. This categorization is the same as [IsGraphic] except that the
// only spacing character is ASCII space, U+0020.
>>>>>>> upstream/master
func IsPrint(r rune) bool

// IsOneOfは、ルーンがいずれかの範囲のメンバーであるかどうかを報告します。
// 関数"In"はより良いシグネチャを提供し、IsOneOfよりも使われるべきです。
func IsOneOf(ranges []*RangeTable, r rune) bool

// ランジのいずれかのメンバーかどうかを報告する。
func In(r rune, ranges ...*RangeTable) bool

<<<<<<< HEAD
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
=======
// IsControl reports whether the rune is a control character.
// The [C] ([Other]) Unicode category includes more code points
// such as surrogates; use [Is](C, r) to test for them.
func IsControl(r rune) bool

// IsLetter reports whether the rune is a letter (category [L]).
func IsLetter(r rune) bool

// IsMark reports whether the rune is a mark character (category [M]).
func IsMark(r rune) bool

// IsNumber reports whether the rune is a number (category [N]).
func IsNumber(r rune) bool

// IsPunct reports whether the rune is a Unicode punctuation character
// (category [P]).
>>>>>>> upstream/master
func IsPunct(r rune) bool

// IsSpaceは、UnicodeのWhite Spaceプロパティによって定義された空白文字であるかどうかを報告します。
// これには、Latin-1スペースに次の文字が含まれます。
//
// '\t'、'\n'、'\v'、'\f'、'\r'、' '、U+0085（NEL）、U+00A0（NBSP）。
//
<<<<<<< HEAD
// スペーシング文字の他の定義は、カテゴリZおよびプロパティPattern_White_Spaceによって設定されています。
=======
// Other definitions of spacing characters are set by category
// Z and property [Pattern_White_Space].
>>>>>>> upstream/master
func IsSpace(r rune) bool

// IsSymbolはルーンが記号の文字であるかどうかを報告します。
func IsSymbol(r rune) bool
