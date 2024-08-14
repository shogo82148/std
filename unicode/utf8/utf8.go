// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// utf8パッケージはUTF-8でエンコードされたテキストをサポートするための関数や定数を実装しています。ルーンとUTF-8バイトシーケンスの変換を行うための関数も含まれています。
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

// DecodeRuneはpの最初のUTF-8エンコーディングを展開し、ルーンと
// そのバイト幅を返します。もしpが空なら、([RuneError], 0)を返します。それ以外の場合、
// エンコーディングが無効なら、(RuneError, 1)を返します。これらは正しい、空でないUTF-8に対しては
// 不可能な結果です。
//
// エンコーディングが無効な場合は、それが不正なUTF-8である、範囲外のルーンをエンコードしている、
// または値のための最短可能なUTF-8エンコーディングでない場合です。
// それ以外の検証は行われません。
func DecodeRune(p []byte) (r rune, size int)

// DecodeRuneInStringは [DecodeRune] と同様ですが、入力は文字列です。もしsが
// 空なら、([RuneError], 0)を返します。それ以外の場合、
// エンコーディングが無効なら、(RuneError, 1)を返します。これらは正しい、空でない
// UTF-8に対しては不可能な結果です。
//
// エンコーディングが無効な場合は、それが不正なUTF-8である、範囲外のルーンをエンコードしている、
// または値のための最短可能なUTF-8エンコーディングでない場合です。
// それ以外の検証は行われません。
func DecodeRuneInString(s string) (r rune, size int)

// DecodeLastRuneはpの最後のUTF-8エンコーディングを展開し、ルーンと
// そのバイト幅を返します。もしpが空なら、([RuneError], 0)を返します。それ以外の場合、
// エンコーディングが無効なら、(RuneError, 1)を返します。これらは正しい、空でないUTF-8に対しては
// 不可能な結果です。
//
// エンコーディングが無効な場合は、それが不正なUTF-8である、範囲外のルーンをエンコードしている、
// または値のための最短可能なUTF-8エンコーディングでない場合です。
// それ以外の検証は行われません。
func DecodeLastRune(p []byte) (r rune, size int)

// DecodeLastRuneInStringは [DecodeLastRune] と同様ですが、入力は文字列です。もしsが
// 空なら、([RuneError], 0)を返します。それ以外の場合、
// エンコーディングが無効なら、(RuneError, 1)を返します。これらは正しい、空でない
// UTF-8に対しては不可能な結果です。
//
// エンコーディングが無効な場合は、それが不正なUTF-8である、範囲外のルーンをエンコードしている、
// または値のための最短可能なUTF-8エンコーディングでない場合です。
// それ以外の検証は行われません。
func DecodeLastRuneInString(s string) (r rune, size int)

// RuneLenはルーンをUTF-8でエンコードするために必要なバイト数を返します。
// ルーンがUTF-8でエンコードすることができない場合は、-1を返します。
func RuneLen(r rune) int

// EncodeRuneは、ルーンのUTF-8エンコーディングをpに書き込みます（これは十分に大きくなければなりません）。
// もしルーンが範囲外であれば、[RuneError] のエンコーディングを書き込みます。
// 書き込まれたバイト数を返します。
func EncodeRune(p []byte, r rune) int

// AppendRuneは、rのUTF-8エンコーディングをpの末尾に追加し、
// 拡張されたバッファを返します。もしルーンが範囲外であれば、
// [RuneError] のエンコーディングを追加します。
func AppendRune(p []byte, r rune) []byte

// RuneCount は p 内のルーンの数を返します。間違ったエンコーディングや短いエンコーディングは、1バイトの幅を持つ単一のルーンとして扱われます。
func RuneCount(p []byte) int

// RuneCountInString is like [RuneCount] but its input is a string.
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
