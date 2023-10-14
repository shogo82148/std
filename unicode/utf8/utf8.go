// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージutf8はUTF-8でエンコードされたテキストをサポートするための関数や定数を実装しています。ルーンとUTF-8バイトシーケンスの変換を行うための関数も含まれています。
// 詳細はhttps://en.wikipedia.org/wiki/UTF-8を参照してください。
package utf8

// エンコーディングに基本的な数値。
const (
	RuneError = '\uFFFD'
	RuneSelf  = 0x80
	MaxRune   = '\U0010FFFF'
	UTFMax    = 4
)

// FullRuneは、pのバイトがルーンの完全なUTF-8エンコードで始まるかどうかを報告します。
// 無効なエンコーディングは完全なルーンと見なされるため、幅-1のエラールーンとして変換されます。
func FullRune(p []byte) bool

// FullRuneInString は FullRune と似ていますが、入力は文字列です。
func FullRuneInString(s string) bool

// DecodeRune関数は、pの最初のUTF-8エンコーディングを解読し、ルーンとそのバイト幅を返します。もしpが空の場合は（RuneError、0）を返します。それ以外の場合、もしエンコーディングが無効な場合は（RuneError、1）を返します。これらの結果は、正しい非空のUTF-8に対して不可能なものです。
// エンコーディングが無効な場合は、正しくないUTF-8をエンコードしているか、範囲外のルーンをエンコードしているか、値の最短のUTF-8エンコーディングではない場合です。他の検証は行われません。
func DecodeRune(p []byte) (r rune, size int)

// DecodeRuneInStringはDecodeRuneと同様ですが、入力が文字列です。もしsが空ならば、(RuneError, 0)を返します。さもなければ、もしエンコーディングが無効な場合は、(RuneError, 1)を返します。これらは、正しい非空のUTF-8では不可能な結果です。
// エンコーディングが無効なのは、正しくないUTF-8であるか、範囲外のルーンをエンコードしているか、または値の最短のUTF-8エンコーディングではない場合です。他の検証は行われません。
func DecodeRuneInString(s string) (r rune, size int)

// DecodeLastRune関数は、pに格納されたUTF-8エンコーディングの最後の文字を解読し、そのルーンとバイトの幅を返します。もしpが空であれば (RuneError, 0)を返します。また、もしエンコーディングが無効であれば (RuneError, 1) を返します。これらの結果は、正しい、空でないUTF-8に対してはあり得ない結果です。
// エンコーディングが無効なのは、正しいUTF-8ではない場合、範囲外のルーンをエンコードしている場合、あるいは値に対して最も短い可能なUTF-8エンコーディングではない場合です。他のバリデーションは行われません。
func DecodeLastRune(p []byte) (r rune, size int)

// DecodeLastRuneInStringはDecodeLastRuneと同様ですが、入力は文字列です。もしsが空の場合は(RuneError, 0)を返します。そうでない場合、エンコードが無効な場合は(RuneError, 1)を返します。これらの結果は、正しい、空でないUTF-8に対しては起こり得ない結果です。
// エンコードが無効な場合とは、不正なUTF-8をエンコードしている場合、範囲外のルーンをエンコードしている場合、または最も短い可能なUTF-8エンコードではない場合を指します。他のバリデーションは行われません。
func DecodeLastRuneInString(s string) (r rune, size int)

// RuneLenはルーンをエンコードするために必要なバイト数を返します。
// ルーンがUTF-8でエンコードすることができない場合は、-1を返します。
func RuneLen(r rune) int

// EncodeRuneは、p（十分に大きくなければなりません）にルーンのUTF-8エンコードを書き込みます。
// ルーンが範囲外の場合は、RuneErrorのエンコードを書き込みます。
// 書き込まれたバイト数を返します。
func EncodeRune(p []byte, r rune) int

// AppendRuneは、rのUTF-8エンコーディングをpの末尾に追加し、
// 拡張されたバッファを返します。もしrが範囲外である場合、
// RuneErrorのエンコーディングを追加します。
func AppendRune(p []byte, r rune) []byte

// RuneCount は p 内のルーンの数を返します。間違ったエンコーディングや短いエンコーディングは、1バイトの幅を持つ単一のルーンとして扱われます。
func RuneCount(p []byte) int

// RuneCountInStringはRuneCountと同じですが、入力は文字列です。
func RuneCountInString(s string) (n int)

// RuneStartは、バイトが符号化された（おそらく無効な）ルーンの最初のバイトであるかどうかを報告します。2番目以降のバイトは常に上位2ビットが10に設定されます。
func RuneStart(b byte) bool

// Validは、pが完全に有効なUTF-8エンコードされたルーンで構成されているかどうかを示します。
func Valid(p []byte) bool

// ValidStringは、sが完全に有効なUTF-8エンコードされたルーンで構成されているかどうかを報告します。
func ValidString(s string) bool

// ValidRuneは、rがUTF-8として正当にエンコードされるかどうかを報告します。
// 範囲外のコードポイントやサロゲートペアの半分は不正です。
func ValidRune(r rune) bool
