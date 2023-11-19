// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// unicodeパッケージは、Unicodeコードポイントのいくつかのプロパティをテストするためのデータと関数を提供します。
package unicode

const (
	MaxRune         = '\U0010FFFF'
	ReplacementChar = '\uFFFD'
	MaxASCII        = '\u007F'
	MaxLatin1       = '\u00FF'
)

// RangeTableは、セット内のUnicodeコードポイントを範囲ごとにリストアップして定義します。
// 範囲は2つのスライスにリストされます。16ビット範囲のスライスと32ビット範囲のスライスです。
// 2つのスライスはソートされた順序で且つ重複しないようにする必要があります。
// また、R32には値が0x10000（1<<16）以上のもののみ含まれるべきです。
type RangeTable struct {
	R16         []Range16
	R32         []Range32
	LatinOffset int
}

// Range16は16ビットのUnicodeコードポイントの範囲を表します。範囲はLoからHiまでの値を含み、指定されたストライドを持ちます。
type Range16 struct {
	Lo     uint16
	Hi     uint16
	Stride uint16
}

// Range32はUnicodeコードポイントの範囲を表し、16ビットに収まらない値が1つ以上含まれる場合に使用されます。この範囲はLoからHiまでの間（Hiも含む）で、指定されたストライドを持ちます。LoとHiは常に1<<16以上である必要があります。
type Range32 struct {
	Lo     uint32
	Hi     uint32
	Stride uint32
}

// CaseRangeは、単純な（1つのコードポイントから別のコードポイントへの）大文字小文字変換のためのUnicodeコードポイントの範囲を表します。
// 範囲は、LoからHiまでの範囲で、固定のストライド1で実行されます。デルタは、その文字の異なるケースに到達するためにコードポイントに追加する数値です。デルタは負数になることもあります。ゼロの場合、対応するケースに文字があることを意味します。特別な場合として、交互に対応する大文字と小文字のペアのシーケンスを表すものがあります。これは、固定デルタの
//
// {UpperLower、UpperLower、UpperLower}
//
// と表示されます。
//
// 定数UpperLowerには、通常ありえないデルタ値があります。
type CaseRange struct {
	Lo    uint32
	Hi    uint32
	Delta d
}

// SpecialCaseはトルコ語などの言語固有の大文字小文字マッピングを表します。
// SpecialCaseのメソッドは、標準のマッピングをカスタマイズするために（オーバーライドすることによって）使用されます。
type SpecialCase []CaseRange

// CaseMapping 内の CaseRanges のデルタ配列への索引。
const (
	UpperCase = iota
	LowerCase
	TitleCase
	MaxCase
)

// [CaseRange] のDeltaフィールドがUpperLowerである場合、
// これはCaseRangeが（例えば） [Upper] [Lower] [Upper] [Lower] の形式のシーケンスを表していることを意味します。
const (
	UpperLower = MaxRune + 1
)

// 指定された範囲テーブル内にルーンが含まれているかどうかを報告します。
func Is(rangeTab *RangeTable, r rune) bool

// IsUpperはルーンが大文字のアルファベットかどうかを報告します。
func IsUpper(r rune) bool

// IsLowerは、ルーンが小文字の文字かどうかを報告します。
func IsLower(r rune) bool

// IsTitleは、与えられたルーンがタイトルケースの文字であるかどうかを報告します。
func IsTitle(r rune) bool

// Toは、指定されたケース（[UpperCase]、[LowerCase]、または[TitleCase]）にルーンをマッピングします。
func To(_case int, r rune) rune

// ToUpperはルーンを大文字にマッピングします。
func ToUpper(r rune) rune

// ToLowerはルーン文字を小文字にマッピングします。
func ToLower(r rune) rune

// ToTitleはルーンをタイトルケースにマッピングします。
func ToTitle(r rune) rune

// ToUpperはルーンを大文字にマッピングしますが、特別なマッピングに優先します。
func (special SpecialCase) ToUpper(r rune) rune

// ToTitleはルーンをタイトルケースにマッピングし、特別なマッピングを優先します。
func (special SpecialCase) ToTitle(r rune) rune

// ToLowerはルーンを小文字にマッピングしますが、特別なマッピングに優先します。
func (special SpecialCase) ToLower(r rune) rune

// SimpleFoldは、Unicodeが定義するシンプルな大文字小文字変換に基づいて、Unicodeコードポイントに相当するコードポイントを反復します。rune自体を含むruneに相当するコードポイントの中で、SimpleFoldは、存在する場合はrよりも大きい最小のruneを返し、存在しない場合は0以上の最小のruneを返します。rが有効なUnicodeコードポイントでない場合、SimpleFold(r)はrを返します。
// 例えば：
//
// SimpleFold('A') = 'a'
// SimpleFold('a') = 'A'
//
// SimpleFold('K') = 'k'
// SimpleFold('k') = '\u212A' (ケルビン記号、K)
// SimpleFold('\u212A') = 'K'
//
// SimpleFold('1') = '1'
//
// SimpleFold(-2) = -2
func SimpleFold(r rune) rune
