// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージ strconv は基本データ型の文字列表現への変換を実装します。
//
// # 数値の変換
//
// 最も一般的な数値の変換は Atoi (文字列から整数へ) と Itoa (整数から文字列へ) です。
//
//	i, err := strconv.Atoi("-42")
//	s := strconv.Itoa(-42)
//
// これらは10進数とGoのint型を仮定しています。
//
// [ParseBool]、[ParseFloat]、[ParseInt]、および [ParseUint] は文字列を値に変換します：
//
//	b, err := strconv.ParseBool("true")
//	f, err := strconv.ParseFloat("3.1415", 64)
//	i, err := strconv.ParseInt("-42", 10, 64)
//	u, err := strconv.ParseUint("42", 10, 64)
//
// パース関数は最も広い型（float64、int64、およびuint64）を返しますが、サイズ引数が
// より狭い幅を指定している場合、結果はその狭い型にデータの損失なく変換できます：
//
//	s := "2147483647" // 最大のint32
//	i64, err := strconv.ParseInt(s, 10, 32)
//	...
//	i := int32(i64)
//
// [FormatBool]、[FormatFloat]、[FormatInt]、および [FormatUint] は値を文字列に変換します：
//
//	s := strconv.FormatBool(true)
//	s := strconv.FormatFloat(3.1415, 'E', -1, 64)
//	s := strconv.FormatInt(-42, 16)
//	s := strconv.FormatUint(42, 16)
//
// [AppendBool]、[AppendFloat]、[AppendInt]、および [AppendUint] は類似していますが、
// フォーマットされた値を宛先スライスに追加します。
//
// # 文字列の変換
//
// [Quote] および [QuoteToASCII] は文字列をクォートされたGo文字リテラルに変換します。
// 後者は非ASCII Unicodeを \u でエスケープして、結果がASCII文字列であることを保証します：
//
//	q := strconv.Quote("Hello, 世界")
//	q := strconv.QuoteToASCII("Hello, 世界")
//
// [QuoteRune] および [QuoteRuneToASCII] は類似していますが、runeを受け入れて、
// クォートされたGo runeリテラルを返します。
//
// [Unquote] および [UnquoteChar] はGo文字列およびruneリテラルのクォートを解除します。
package strconv
