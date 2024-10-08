// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// utf16パッケージはUTF-16シーケンスのエンコードとデコードを実装します。
package utf16

// IsSurrogateは指定されたUnicodeコードポイントが
// 代理ペアに現れることができるかどうかを報告します。
func IsSurrogate(r rune) bool

// DecodeRune はサロゲートペアのUTF-16デコードを返します。
// サロゲートペアが正しいUTF-16のサロゲートペアでない場合、
// DecodeRune はUnicodeの代替コードポイントU+FFFDを返します。
func DecodeRune(r1, r2 rune) rune

// EncodeRuneは与えられたルーンに対して、UTF-16サロゲートペアのr1、r2を返します。
// もしルーンが有効なUnicodeコードポイントでない場合やエンコーディングが必要ではない場合、
// EncodeRuneはU+FFFD、U+FFFDを返します。
func EncodeRune(r rune) (r1, r2 rune)

// RuneLenは、ルーンのUTF-16エンコーディングに含まれる16ビットワードの数を返します。
// ルーンがUTF-16でエンコードするための有効な値でない場合、-1を返します。
func RuneLen(r rune) int

// EncodeはUnicodeコードポイントの列sのUTF-16エンコーディングを返します。
func Encode(s []rune) []uint16

// AppendRuneはUnicodeのコードポイントrのUTF-16エンコーディングを
// pの末尾に追加し、拡張されたバッファを返します。コードポイントが有効な
// Unicodeのコードポイントでない場合、U+FFFDのエンコーディングを追加します。
func AppendRune(a []uint16, r rune) []uint16

// DecodeはUTF-16エンコーディングsで表されるUnicodeのコードポイントのシーケンスを返します。
func Decode(s []uint16) []rune
