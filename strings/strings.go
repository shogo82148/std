// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package strings implements simple functions to manipulate UTF-8 encoded strings.
//
// For information about UTF-8 strings in Go, see https://blog.golang.org/strings.
package strings

import (
	"github.com/shogo82148/std/unicode"
)

// Countは、s内の重複しないsubstrのインスタンス数を数えます。
// substrが空の文字列の場合、Countはs内のUnicodeコードポイントの数に1を加えたものを返します。
func Count(s, substr string) int

// Containsは、substrがs内に含まれているかどうかを報告します。
func Contains(s, substr string) bool

// ContainsAnyは、chars内の任意のUnicodeコードポイントがs内に含まれているかどうかを報告します。
func ContainsAny(s, chars string) bool

// ContainsRuneは、Unicodeコードポイントrがs内に含まれているかどうかを報告します。
func ContainsRune(s string, r rune) bool

// ContainsFuncは、s内の任意のUnicodeコードポイントrがf(r)を満たすかどうかを報告します。
func ContainsFunc(s string, f func(rune) bool) bool

// LastIndexは、s内のsubstrの最後のインスタンスのインデックスを返します。
// substrがs内に存在しない場合は-1を返します。
func LastIndex(s, substr string) int

// IndexByteは、s内の最初のcのインスタンスのインデックスを返します。
// cがsに存在しない場合は-1を返します。
func IndexByte(s string, c byte) int

// IndexRuneは、Unicodeコードポイントrの最初のインスタンスのインデックスを返します。
// rがsに存在しない場合は-1を返します。
// rがutf8.RuneErrorの場合、無効なUTF-8バイトシーケンスの最初のインスタンスを返します。
func IndexRune(s string, r rune) int

// IndexAnyは、sにcharsの任意のUnicodeコードポイントの最初のインスタンスのインデックスを返します。
// charsのUnicodeコードポイントがsに存在しない場合は-1を返します。
func IndexAny(s, chars string) int

// LastIndexAnyは、sにcharsの任意のUnicodeコードポイントの最後のインスタンスのインデックスを返します。
// charsのUnicodeコードポイントがsに存在しない場合は-1を返します。
func LastIndexAny(s, chars string) int

// LastIndexByteは、sの最後のインスタンスのインデックスを返します。
// cがsに存在しない場合は-1を返します。
func LastIndexByte(s string, c byte) int

// SplitNは、sをsepで区切った部分文字列のスライスを返します。
//
// countは、返す部分文字列の数を決定します。
//
//	n > 0：最大n個の部分文字列。最後の部分文字列は区切り文字以降の残りの部分です。
//	n == 0：結果はnil（部分文字列がゼロ個）
//	n < 0：すべての部分文字列
//
<<<<<<< HEAD
// sとsepのエッジケース（空の文字列など）は、Splitのドキュメントで説明されているように処理されます。
=======
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for [Split].
>>>>>>> upstream/master
//
// 最初の区切り文字を基準に分割するには、Cutを参照してください。
func SplitN(s, sep string, n int) []string

// SplitAfterNは、sをsepの後にスライスし、それらの部分文字列のスライスを返します。
//
// countは、返す部分文字列の数を決定します。
//
//	n > 0：最大n個の部分文字列。最後の部分文字列は区切り文字以降の残りの部分です。
//	n == 0：結果はnil（部分文字列がゼロ個）
//	n < 0：すべての部分文字列
//
// sとsepのエッジケース（空の文字列など）は、SplitAfterのドキュメントで説明されているように処理されます。
func SplitAfterN(s, sep string, n int) []string

// Splitは、sをsepで区切り、それらの区切り文字の間の部分文字列のスライスを返します。
//
// sがsepを含まず、sepが空でない場合、Splitは長さ1のスライスを返します。その唯一の要素はsです。
//
// sepが空の場合、Splitは各UTF-8シーケンスの後に分割します。sとsepの両方が空の場合、Splitは空のスライスを返します。
//
<<<<<<< HEAD
// countが-1のSplitNと同等です。
=======
// It is equivalent to [SplitN] with a count of -1.
>>>>>>> upstream/master
//
// 最初の区切り文字を基準に分割するには、Cutを参照してください。
func Split(s, sep string) []string

// SplitAfterは、sをsepの後にスライスし、それらの部分文字列のスライスを返します。
//
// sがsepを含まず、sepが空でない場合、SplitAfterは長さ1のスライスを返します。その唯一の要素はsです。
//
// sepが空の場合、SplitAfterは各UTF-8シーケンスの後に分割します。sとsepの両方が空の場合、SplitAfterは空のスライスを返します。
//
<<<<<<< HEAD
// countが-1のSplitAfterNと同等です。
=======
// It is equivalent to [SplitAfterN] with a count of -1.
>>>>>>> upstream/master
func SplitAfter(s, sep string) []string

// Fieldsは、sをUnicode.IsSpaceによって定義される1つ以上の連続する空白文字の各インスタンスで分割し、sの部分文字列のスライスまたは空のスライスを返します。
func Fields(s string) []string

// FieldsFuncは、Unicodeコードポイントcがf(c)を満たす連続するランで文字列sを分割し、sのスライスの配列を返します。
// sのすべてのコードポイントがf(c)を満たすか、文字列が空の場合、空のスライスが返されます。
//
// FieldsFuncは、f(c)を呼び出す順序について保証せず、fが常に同じ値を返すことを前提としています。
func FieldsFunc(s string, f func(rune) bool) []string

// Joinは、最初の引数の要素を連結して単一の文字列を作成します。区切り文字列sepは、結果の文字列の要素間に配置されます。
func Join(elems []string, sep string) string

// HasPrefixは、文字列sがprefixで始まるかどうかを報告します。
func HasPrefix(s, prefix string) bool

// HasSuffixは、文字列sがsuffixで終わるかどうかを報告します。
func HasSuffix(s, suffix string) bool

// Mapは、マッピング関数に従ってすべての文字を変更した文字列sのコピーを返します。
// マッピング関数が負の値を返す場合、文字は置換されずに文字列から削除されます。
func Map(mapping func(rune) rune, s string) string

// Repeatは、文字列sのcount個のコピーからなる新しい文字列を返します。
//
// countが負の場合、または（len(s) * count）の結果がオーバーフローする場合、パニックが発生します。
func Repeat(s string, count int) string

// ToUpperは、すべてのUnicode文字を大文字にマップしたsを返します。
func ToUpper(s string) string

// ToLowerは、すべてのUnicode文字を小文字にマップしたsを返します。
func ToLower(s string) string

// ToTitleは、すべてのUnicode文字をUnicodeタイトルケースにマップしたsのコピーを返します。
func ToTitle(s string) string

// ToUpperSpecialは、Unicode文字をすべて、cで指定されたケースマッピングを使用して大文字にマップしたsのコピーを返します。
func ToUpperSpecial(c unicode.SpecialCase, s string) string

// ToLowerSpecialは、Unicode文字をすべて、cで指定されたケースマッピングを使用して小文字にマップしたsのコピーを返します。
func ToLowerSpecial(c unicode.SpecialCase, s string) string

// ToTitleSpecialは、Unicode文字をすべて、特別なケースルールに優先してUnicodeタイトルケースにマップしたsのコピーを返します。
func ToTitleSpecial(c unicode.SpecialCase, s string) string

// ToValidUTF8は、無効なUTF-8バイトシーケンスのランを置換文字列で置き換えたsのコピーを返します。置換文字列は空にすることができます。
func ToValidUTF8(s, replacement string) string

// Titleは、単語の先頭を表すすべてのUnicode文字をUnicodeタイトルケースにマップしたsのコピーを返します。
//
// Deprecated: Titleが単語の境界に使用するルールは、Unicode句読点を適切に処理しません。代わりに、golang.org/x/text/casesを使用してください。
func Title(s string) string

// TrimLeftFuncは、f(c)がtrueを返す最初のUnicodeコードポイントcを含まないように、文字列sの先頭からすべてのUnicodeコードポイントcを削除したスライスを返します。
func TrimLeftFunc(s string, f func(rune) bool) string

// TrimRightFuncは、f(c)がtrueを返す最後のUnicodeコードポイントcを含まないように、文字列sの末尾からすべてのUnicodeコードポイントcを削除したスライスを返します。
func TrimRightFunc(s string, f func(rune) bool) string

// TrimFuncは、f(c)がtrueを返す最初と最後のUnicodeコードポイントcを含まないように、文字列sの先頭と末尾からすべてのUnicodeコードポイントcを削除したスライスを返します。
func TrimFunc(s string, f func(rune) bool) string

// IndexFuncは、f(c)がtrueを返す最初のUnicodeコードポイントのインデックスを返します。見つからない場合は-1を返します。
func IndexFunc(s string, f func(rune) bool) int

// LastIndexFuncは、f(c)がtrueを返す最後のUnicodeコードポイントのインデックスを返します。見つからない場合は-1を返します。
func LastIndexFunc(s string, f func(rune) bool) int

// Trimは、cutsetに含まれるすべての先頭と末尾のUnicodeコードポイントを削除した文字列sのスライスを返します。
func Trim(s, cutset string) string

// TrimLeftは、cutsetに含まれるすべての先頭のUnicodeコードポイントを削除した文字列sのスライスを返します。
//
<<<<<<< HEAD
// 接頭辞を削除するには、代わりにTrimPrefixを使用してください。
=======
// To remove a prefix, use [TrimPrefix] instead.
>>>>>>> upstream/master
func TrimLeft(s, cutset string) string

// TrimRightは、cutsetに含まれるすべての末尾のUnicodeコードポイントを削除した文字列sのスライスを返します。
//
<<<<<<< HEAD
// 接尾辞を削除するには、代わりにTrimSuffixを使用してください。
=======
// To remove a suffix, use [TrimSuffix] instead.
>>>>>>> upstream/master
func TrimRight(s, cutset string) string

// TrimSpaceは、Unicodeで定義されるように、すべての先頭と末尾の空白を削除した文字列sのスライスを返します。
func TrimSpace(s string) string

// TrimPrefixは、指定された接頭辞文字列を除いたsを返します。
// sが接頭辞で始まらない場合、sは変更されずにそのまま返されます。
func TrimPrefix(s, prefix string) string

// TrimSuffixは、指定された接尾辞文字列を除いたsを返します。
// sが接尾辞で終わらない場合、sは変更されずにそのまま返されます。
func TrimSuffix(s, suffix string) string

// Replaceは、古いものの最初のn個の重複しないインスタンスが新しいものに置き換えられた文字列sのコピーを返します。
// oldが空の場合、文字列の先頭と各UTF-8シーケンスの後に一致し、kルーン文字列に対してk+1の置換が生成されます。
// n < 0の場合、置換の数に制限はありません。
func Replace(s, old, new string, n int) string

// ReplaceAllは、古いもののすべての重複しないインスタンスが新しいものに置き換えられた文字列sのコピーを返します。
// oldが空の場合、文字列の先頭と各UTF-8シーケンスの後に一致し、kルーン文字列に対してk+1の置換が生成されます。
func ReplaceAll(s, old, new string) string

// EqualFoldは、UTF-8文字列として解釈されたsとtが、単純なUnicodeの大文字小文字を区別しない比較において等しいかどうかを報告します。
// これは、大文字小文字を区別しない形式の大文字小文字を区別しない性質です。
func EqualFold(s, t string) bool

// Indexは、s内のsubstrの最初のインスタンスのインデックスを返します。substrがsに存在しない場合は-1を返します。
func Index(s, substr string) int

// 最初の sep のインスタンスを中心に s をスライスし、
// sep の前と後のテキストを返します。
// found は、sep が s に現れるかどうかを報告します。
// sep が s に現れない場合、cut は s、""、false を返します。
func Cut(s, sep string) (before, after string, found bool)

// CutPrefix は、指定された先頭接頭辞文字列を除いた s を返し、接頭辞が見つかったかどうかを報告します。
// s が prefix で始まらない場合、CutPrefix は s、false を返します。
// prefix が空の文字列の場合、CutPrefix は s、true を返します。
func CutPrefix(s, prefix string) (after string, found bool)

// CutSuffix は、指定された末尾接尾辞文字列を除いた s を返し、接尾辞が見つかったかどうかを報告します。
// s が suffix で終わらない場合、CutSuffix は s、false を返します。
// suffix が空の文字列の場合、CutSuffix は s、true を返します。
func CutSuffix(s, suffix string) (before string, found bool)
