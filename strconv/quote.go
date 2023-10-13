// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run makeisprint.go -output isprint.go

package strconv

// Quoteは、sを表すダブルクォートで囲まれたGo文字列リテラルを返します。
// 返される文字列は、制御文字およびIsPrintによって定義された印刷不可能な文字のために、
// Goエスケープシーケンス（\t、\n、\xFF、\u0100）を使用します。
func Quote(s string) string

// AppendQuoteは、sを表すダブルクォートで囲まれたGo文字列リテラル（Quoteによって生成されたもの）をdstに追加し、拡張されたバッファを返します。
func AppendQuote(dst []byte, s string) []byte

// QuoteToASCIIは、sを表すダブルクォートで囲まれたGo文字列リテラルを返します。
// 返される文字列は、制御文字およびIsPrintによって定義された印刷不可能な文字のために、
// Goエスケープシーケンス（\t、\n、\xFF、\u0100）を使用します。
func QuoteToASCII(s string) string

// AppendQuoteToASCIIは、sを表すダブルクォートで囲まれたGo文字列リテラル（QuoteToASCIIによって生成されたもの）をdstに追加し、拡張されたバッファを返します。
func AppendQuoteToASCII(dst []byte, s string) []byte

// QuoteToGraphicは、sを表すダブルクォートで囲まれたGo文字列リテラルを返します。
// 返される文字列は、Unicodeのグラフィック文字（IsGraphicによって定義される）を変更せず、
// 非グラフィック文字に対してGoエスケープシーケンス（\t、\n、\xFF、\u0100）を使用します。
func QuoteToGraphic(s string) string

// AppendQuoteToGraphicは、sを表すダブルクォートで囲まれたGo文字列リテラル（QuoteToGraphicによって生成されたもの）をdstに追加し、拡張されたバッファを返します。
func AppendQuoteToGraphic(dst []byte, s string) []byte

// QuoteRuneは、runeを表すシングルクォートで囲まれたGo文字リテラルを返します。
// 返される文字列は、制御文字およびIsPrintによって定義された印刷不可能な文字のために、
// Goエスケープシーケンス（\t、\n、\xFF、\u0100）を使用します。
// rが有効なUnicodeコードポイントでない場合、Unicode置換文字U+FFFDとして解釈されます。
func QuoteRune(r rune) string

// AppendQuoteRuneは、runeを表すシングルクォートで囲まれたGo文字リテラル（QuoteRuneによって生成されたもの）をdstに追加し、拡張されたバッファを返します。
func AppendQuoteRune(dst []byte, r rune) []byte

// QuoteRuneToASCIIは、runeを表すシングルクォートで囲まれたGo文字リテラルを返します。
// 返される文字列は、制御文字およびIsPrintによって定義された印刷不可能な文字のために、
// Goエスケープシーケンス（\t、\n、\xFF、\u0100）を使用します。
// rが有効なUnicodeコードポイントでない場合、Unicode置換文字U+FFFDとして解釈されます。
func QuoteRuneToASCII(r rune) string

// AppendQuoteRuneToASCIIは、runeを表すシングルクォートで囲まれたGo文字リテラル（QuoteRuneToASCIIによって生成されたもの）をdstに追加し、拡張されたバッファを返します。
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte

// QuoteRuneToGraphicは、runeを表すシングルクォートで囲まれたGo文字リテラルを返します。
// 返される文字列は、Unicodeのグラフィック文字（IsGraphicによって定義される）を変更せず、
// 非グラフィック文字に対してGoエスケープシーケンス（\t、\n、\xFF、\u0100）を使用します。
// rが有効なUnicodeコードポイントでない場合、Unicode置換文字U+FFFDとして解釈されます。
func QuoteRuneToGraphic(r rune) string

// AppendQuoteRuneToGraphicは、runeを表すシングルクォートで囲まれたGo文字リテラル（QuoteRuneToGraphicによって生成されたもの）をdstに追加し、拡張されたバッファを返します。
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte

// CanBackquoteは、タブ以外の制御文字を含まない単一行のバッククォート文字列として、文字列sを変更せずに表現できるかどうかを報告します。
func CanBackquote(s string) bool

// UnquoteCharは、文字列sによって表されるエスケープされた文字列または文字リテラルの最初の文字またはバイトをデコードします。
// 4つの値を返します。
//
// 1. value：デコードされたUnicodeコードポイントまたはバイト値
// 2. multibyte：デコードされた文字がマルチバイトUTF-8表現を必要とするかどうかを示すブール値
// 3. tail：文字の後の残りの文字列
// 4. エラー：文字が構文的に有効である場合はnilになります。
//
// 2番目の引数であるquoteは、解析されるリテラルの種類を指定し、許可されるエスケープされた引用符文字を決定します。
// シングルクォートに設定すると、\'のシーケンスを許可し、エスケープされていない'を許可しません。
// ダブルクォートに設定すると、\"を許可し、エスケープされていない"を許可しません。
// ゼロに設定すると、どちらのエスケープも許可せず、両方の引用符文字がエスケープされていない状態で現れることを許可します。
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)

// QuotedPrefixは、sのプレフィックスにある引用符で囲まれた文字列（Unquoteで理解される形式）を返します。
// sが有効な引用符で囲まれた文字列で始まらない場合、QuotedPrefixはエラーを返します。
func QuotedPrefix(s string) (string, error)

// Unquoteは、sを単一引用符、二重引用符、またはバッククォートで囲まれたGo文字列リテラルとして解釈し、sが引用する文字列値を返します。
// （sが単一引用符で囲まれている場合、それはGo文字リテラルであることに注意してください。Unquoteは対応する1文字の文字列を返します。）
func Unquote(s string) (string, error)

// IsPrintは、文字がGoによって印刷可能と定義されているかどうかを報告します。
// これは、unicode.IsPrintと同じ定義で、文字、数字、句読点、記号、ASCIIスペースを含みます。
func IsPrint(r rune) bool

// IsGraphicは、Unicodeによってグラフィックとして定義されたルーンかどうかを報告します。
// このような文字には、カテゴリL、M、N、P、S、およびZsからの文字、数字、句読点、記号、スペースが含まれます。
func IsGraphic(r rune) bool
