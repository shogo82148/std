// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// bytesパッケージはバイトスライスの操作のための関数を実装します。
// これは [strings] パッケージの機能に類似しています。
package bytes

import (
	"github.com/shogo82148/std/unicode"
)

// Equalは、aとbが同じ長さで同じバイトを含むかどうかを報告します。
// nilの引数は空のスライスと等価です。
func Equal(a, b []byte) bool

// Compare関数は2つのバイトスライスを辞書的に比較して整数を返します。
// a == bの場合は0、a < bの場合は-1、a > bの場合は+1となります。
// nilの引数は空スライスと同等です。
func Compare(a, b []byte) int

// Count は s において非重複の sep の出現回数を数えます。
// もし sep が空のスライスなら、Count は s 中の UTF-8 エンコードされたコードポイントの数に 1 を加えた値を返します。
func Count(s, sep []byte) int

// subsliceがb内に含まれているかどうかを報告します。
func Contains(b, subslice []byte) bool

// ContainsAnyは、chars内のUTF-8エンコードされたコードポイントのいずれかがb内に含まれているかどうかを報告します。
func ContainsAny(b []byte, chars string) bool

// ContainsRuneは、UTF-8でエンコードされたバイトスライスbにルーンが含まれているかどうかを報告します。
func ContainsRune(b []byte, r rune) bool

// ContainsFuncは、UTF-8エンコードされたコードポイントの中で、bのどれかがf(r)を満たすかどうかを報告します。
func ContainsFunc(b []byte, f func(rune) bool) bool

// IndexByteはb内の最初のcのインスタンスのインデックスを返します。もしcがbに存在しない場合は、-1を返します。
func IndexByte(b []byte, c byte) int

// LastIndex関数は、s内のsepの最後のインスタンスのインデックスを返します。sepがs内に存在しない場合は、-1を返します。
func LastIndex(s, sep []byte) int

// LastIndexByteは、cがs内で最後に出現するインデックスを返します。cがs内に存在しない場合は-1を返します。
func LastIndexByte(s []byte, c byte) int

// IndexRuneはsをUTF-8でエンコードされたコードポイントのシーケンスとして解釈します。
// sの中で指定されたルーンの最初の出現のバイトインデックスを返します。
// sにルーンが含まれていない場合は-1を返します。
// rが [utf8.RuneError] である場合、無効なUTF-8バイトシーケンスの最初のインスタンスを返します。
func IndexRune(s []byte, r rune) int

// IndexAnyはsをUTF-8エンコードされたUnicodeのコードポイントのシーケンスとして解釈します。
// sの中でcharsのいずれかのUnicodeコードポイントの最初の出現のバイトインデックスを返します。
// charsが空であるか、共通のコードポイントがない場合は-1を返します。
func IndexAny(s []byte, chars string) int

// LastIndexAnyは、sをUTF-8でエンコードされたUnicodeコードポイントの
// シーケンスとして解釈します。charsに含まれる任意のUnicodeコードポイントの
// 最後の出現のバイトインデックスを返します。charsが空である場合や、
// 共通のコードポイントが存在しない場合は、-1を返します。
func LastIndexAny(s []byte, chars string) int

// SplitNは、sをsepで区切り、そのセパレーターの間のサブスライスのスライスを返します。
// sepが空の場合、SplitNは各UTF-8シーケンスの後に分割します。
// countは返すサブスライスの数を決定します：
//
//   - n> 0：最大でn個のサブスライス；最後のサブスライスは分割パートが含まれます。
//   - n == 0：結果はnilです（ゼロのサブスライス）
//   - n < 0：すべてのサブスライス
//
// 最初のセパレーターの周りで分割するには、[Cut] を参照してください。
func SplitN(s, sep []byte, n int) [][]byte

// SplitAfterNはsをsepの各インスタンスの後ろでサブスライスに分割し、それらのサブスライスのスライスを返します。
// sepが空である場合、SplitAfterNはUTF-8シーケンスの後ろで分割します。
// countは返すサブスライスの数を決定します：
//
//   - n > 0：最大nのサブスライス；最後のサブスライスは分割されていない残りになります。
//   - n == 0：結果はnilです（サブスライスはゼロ個）
//   - n < 0：すべてのサブスライス
func SplitAfterN(s, sep []byte, n int) [][]byte

// Split関数は、sをsepで区切ったすべてのサブスライスから成るスライスを返します。
// sepが空の場合、Split関数はUTF-8シーケンスごとに区切ります。
// これは、SplitN関数にカウント-1を指定した場合と同等です。
//
// 最初の区切り文字で区切る場合は、[Cut] 関数を参照してください。
func Split(s, sep []byte) [][]byte

// SplitAfterは、sをsepの各インスタンスの後にスライスし、それらのサブスライスのスライスを返します。
// sepが空の場合、UTF-8のシーケンスの後に分割します。
// これは、countが-1のSplitAfterNと同等です。
func SplitAfter(s, sep []byte) [][]byte

// Fieldsは、sをUTF-8エンコードされたコードポイントのシーケンスとして解釈します。
// [unicode.IsSpace] で定義される1つ以上の連続する空白文字の各インスタンスの周りでスライスsを分割し、
// sの部分スライスのスライスを返します。sが空白のみを含む場合は空のスライスを返します。
func Fields(s []byte) [][]byte

// FieldsFuncは、sをUTF-8でエンコードされたコードポイントのシーケンスとして解釈します。
// それは、f(c)を満たすコードポイントcの連続を各ランでsを分割し、sのサブスライスのスライスを返します。
// sのすべてのコードポイントがf(c)を満たすか、またはlen(s) == 0の場合、空のスライスが返されます。
//
// FieldsFuncは、f(c)をどの順序で呼び出すかについて保証はなく、fは常に同じ値を返すと仮定しています。
func FieldsFunc(s []byte, f func(rune) bool) [][]byte

// Join関数は、sの要素を連結して新しいバイトスライスを作成します。結果のスライスの要素間にはセパレーターsepが配置されます。
func Join(s [][]byte, sep []byte) []byte

// HasPrefixは、バイトスライスsがprefixで始まるかどうかを報告します。
func HasPrefix(s, prefix []byte) bool

// HasSuffixは、バイトスライスsがsuffixで終わるかどうかを報告します。
func HasSuffix(s, suffix []byte) bool

// Map関数は、与えられたマッピング関数に基づいて、バイトスライスsのすべての文字が変更されたコピーを返します。
// マッピング関数が負の値を返すと、文字は置換せずにバイトスライスから削除されます。
// sと出力の文字はUTF-8エンコードされたコードポイントとして解釈されます。
func Map(mapping func(r rune) rune, s []byte) []byte

// Repeat は、bのcount回のコピーからなる新しいバイトスライスを返します。
//
// countが負数であるか、(len(b) * count)の結果がオーバーフローする場合、パニックが発生します。
func Repeat(b []byte, count int) []byte

// ToUpperは、すべてのUnicode文字を大文字に変換したバイトスライスsのコピーを返します。
func ToUpper(s []byte) []byte

// ToLowerは、すべてのUnicodeの文字を小文字にマッピングしたバイトスライスsのコピーを返します。
func ToLower(s []byte) []byte

// ToTitleはsをUTF-8でエンコードされたバイト列として扱い、すべてのUnicodeの文字をタイトルケースにマップしたコピーを返します。
func ToTitle(s []byte) []byte

// ToUpperSpecialはsをUTF-8エンコードされたバイトとして扱い、すべてのUnicodeの文字をその大文字に変換したコピーを返します。特殊な大文字変換ルールを優先します。
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte

// ToLowerSpecialはUTF-8エンコードされたバイト列sを扱い、ユニコードの文字をすべて小文字に変換し、特殊なケースのルールを優先します。
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte

// ToTitleSpecialはUTF-8でエンコードされたバイト列としてsを扱い、すべてのUnicode文字をタイトルケースにマッピングしたコピーを返します。特殊なケースのルールに優先します。
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte

// ToValidUTF8は、sをUTF-8でエンコードされたバイトとして処理し、各バイトの連続が不正なUTF-8を表す場合に、置換バイト（空の場合もあります）で置き換えられたコピーを返します。
func ToValidUTF8(s, replacement []byte) []byte

// TitleはUTF-8でエンコードされたバイト列sをUnicodeの文字として扱い、単語の先頭にあるすべての文字をタイトルケースにマッピングしたコピーを返します。
//
// 廃止予定: Titleが単語の境界を処理する際、Unicodeの句読点を適切に扱えません。golang.org/x/text/casesを代わりに使用してください。
func Title(s []byte) []byte

// TrimLeftFuncはUTF-8でエンコードされたバイト列sを処理し、f(c)を満たすすべての先頭のUTF-8エンコードされたコードポイントcを除いたsのサブスライスを返します。
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte

// TrimRightFuncは、末尾にあるすべてのトレイリングUTF-8エンコードされたコードポイントcをスライスして、f（c）を満たすものを取り除いたsの部分スライスを返します。
func TrimRightFunc(s []byte, f func(r rune) bool) []byte

// TrimFunc は、前方および後方のすべての先頭と末尾をスライスして、f(c) で指定された条件を満たすすべての UTF-8 エンコードされたコードポイント c を除去して、s のサブスライスを返します。
func TrimFunc(s []byte, f func(r rune) bool) []byte

// TrimPrefixは、指定されたプレフィックス文字列を先頭から除去したsを返します。
// sがプレフィックスで始まらない場合、sは変更されずに返されます。
func TrimPrefix(s, prefix []byte) []byte

// TrimSuffixは、指定された末尾の接尾辞文字列を除いたsを返します。
// もしsが接尾辞で終わっていない場合はsは変更されずにそのまま返されます。
func TrimSuffix(s, suffix []byte) []byte

// IndexFuncは、sをUTF-8エンコードされたコードポイントのシーケンスとして解釈します。
// sの中でf(c)を満たす最初のUnicodeコードポイントのバイトインデックスを返します。
// 該当するコードポイントがない場合は-1を返します。
func IndexFunc(s []byte, f func(r rune) bool) int

// LastIndexFuncはsをUTF-8でエンコードされたコードポイントのシーケンスとして解釈します。
// f(c)を満たす最後のUnicodeコードポイントのバイトインデックスを、f(c)を満たすものがない場合は-1を返します。
func LastIndexFunc(s []byte, f func(r rune) bool) int

// Trimは、cutsetに含まれるすべての先頭と末尾のUTF-8エンコードされたコードポイントをスライスして、sのサブスライスを返します。
func Trim(s []byte, cutset string) []byte

// TrimLeftは、cutsetに含まれるすべての先行するUTF-8エンコードされたコードポイントを除去することによって、sの一部のサブスライスを返します。
func TrimLeft(s []byte, cutset string) []byte

// TrimRightは、cutsetに含まれるすべての末尾のUTF-8エンコードされたコードポイントを取り除いて、sの一部をサブスライスとして返します。
func TrimRight(s []byte, cutset string) []byte

// TrimSpaceは、Unicodeで定義されたように、sの先頭と末尾のすべての空白を取り除いた部分列を返します。
func TrimSpace(s []byte) []byte

// RunesはUTF-8でエンコードされたコードポイントのシーケンスとしてsを解釈します。
// sと同等のルーン（Unicodeのコードポイント）のスライスを返します。
func Runes(s []byte) []rune

// Replaceは、スライスsの最初のn個の重ならない
// oldのインスタンスをnewで置き換えたスライスのコピーを返します。
// oldが空の場合、スライスの先頭とUTF-8シーケンスの後に一致し、
// kランスライスに対してk+1回の置換がされます。
// nが負の場合、置換の数に制限はありません。
func Replace(s, old, new []byte, n int) []byte

// ReplaceAllは、スライスsのすべての非重複インスタンスをoldからnewに置き換えたスライスのコピーを返します。
// oldが空の場合、スライスの先頭とUTF-8シーケンスの後に一致し、kランスライスに対してk+1回の置換が行われます。
func ReplaceAll(s, old, new []byte) []byte

// EqualFoldは、UTF-8文字列として解釈されたsとtが、単純なUnicodeの大文字小文字を区別しない比較で等しいかどうかを報告します。これは、大文字小文字を区別しない形式の一般的なUnicodeです。
func EqualFold(s, t []byte) bool

// Indexは、sの最初のsepのインスタンスのインデックスを返します。sepがsに存在しない場合は、-1を返します。
func Index(s, sep []byte) int

// 最初の「sep」のインスタンス周りのスライス「s」を切り取り、
// sepの前と後のテキストを返します。
// 見つかった結果は「sep」が「s」に現れるかどうかを報告します。
// もし「sep」が「s」に現れない場合、cutは「s」とnil、falseを返します。
//
// Cutは元のスライス「s」のスライスを返します、コピーではありません。
func Cut(s, sep []byte) (before, after []byte, found bool)

// Cloneはb[:len(b)]のコピーを返します。
// 結果には余分な未使用の容量があるかもしれません。
// Clone(nil)はnilを返します。
func Clone(b []byte) []byte

// CutPrefixは与えられた先頭接頭辞のバイトスライスを取り除き、
// 接頭辞が見つかったかどうかを報告します。
// もしsが接頭辞で始まっていない場合、CutPrefixはs、falseを返します。
// もし接頭辞が空のバイトスライスの場合、CutPrefixはs、trueを返します。
//
// CutPrefixは元のスライスsの断片を返します、コピーではありません。
func CutPrefix(s, prefix []byte) (after []byte, found bool)

// CutSuffixは与えられた終了サフィックスのバイトスライスを除いたsを返し、そのサフィックスが見つかったかどうかを報告します。
// もしsがサフィックスで終わらない場合、CutSuffixはs、falseを返します。
// もしサフィックスが空のバイトスライスである場合、CutSuffixはs、trueを返します。
//
// CutSuffixは元のスライスsのスライスを返しますが、コピーではありません。
func CutSuffix(s, suffix []byte) (before []byte, found bool)
