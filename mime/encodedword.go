// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mime

import (
	"github.com/shogo82148/std/io"
)

// WordEncoderは、RFC 2047のエンコードされたワードエンコーダです。
type WordEncoder byte

const (
	// BEncodingは、RFC 2045によって定義されたBase64エンコーディングスキームを表します。
	BEncoding = WordEncoder('b')
	// QEncodingは、RFC 2047によって定義されたQ-エンコーディングスキームを表します。
	QEncoding = WordEncoder('q')
)

// Encodeは、sのエンコードされた単語形式を返します。もしsが特殊文字を含まないASCIIであれば、
// それは変更されずに返されます。提供されたcharsetは、sのIANA
// 文字セット名です。それは大文字と小文字を区別しません。
func (e WordEncoder) Encode(charset, s string) string

// WordDecoderは、RFC 2047のエンコードされた単語を含むMIMEヘッダーをデコードします。
type WordDecoder struct {
	// CharsetReaderがnilでない場合、提供されたcharsetからUTF-8に変換する
	// charset変換リーダーを生成する関数を定義します。
	// Charsetは常に小文字です。utf-8、iso-8859-1、us-ascii charsetは
	// デフォルトで処理されます。
	// CharsetReaderの結果値のうちの1つは非nilでなければなりません。
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)
}

// Decodeは、RFC 2047のエンコードされた単語をデコードします。
func (d *WordDecoder) Decode(word string) (string, error)

// DecodeHeaderは、与えられた文字列のすべてのエンコードされた単語をデコードします。
// dの [WordDecoder.CharsetReader] がエラーを返す場合にのみエラーを返します。
func (d *WordDecoder) DecodeHeader(header string) (string, error)
