// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package base32は、RFC 4648で指定されているように、base32エンコーディングを実装します。
package base32

import (
	"github.com/shogo82148/std/io"
)

// Encodingは、32文字のアルファベットによって定義される基数32のエンコーディング/デコーディングスキームです。
// 最も一般的なものは、SASL GSSAPIで導入され、RFC 4648で標準化された「base32」エンコーディングです。
// 代替の「base32hex」エンコーディングは、DNSSECで使用されます。
type Encoding struct {
	encode    [32]byte
	decodeMap [256]uint8
	padChar   rune
}

const (
	StdPadding rune = '='
	NoPadding  rune = -1
)

// NewEncodingは、与えられたアルファベットで定義されたパディングされたエンコーディングを返します。
// アルファベットは、パディング文字またはCR / LF（'\r'、'\n'）を含まず、一意のバイト値を含む32バイトの文字列である必要があります。
// アルファベットは、マルチバイトUTF-8の特別な処理なしにバイト値のシーケンスとして扱われます。
// 結果のエンコーディングは、デフォルトのパディング文字（'='）を使用します。
// パディング文字を変更または無効にするには、[Encoding.WithPadding] を使用できます。
func NewEncoding(encoder string) *Encoding

// StdEncodingは、RFC 4648で定義されている標準のbase32エンコーディングです。
var StdEncoding = NewEncoding(encodeStd)

// HexEncodingは、RFC 4648で定義されている「Extended Hex Alphabet」です。
// 通常、DNSで使用されます。
var HexEncoding = NewEncoding(encodeHex)

// WithPaddingは、指定されたパディング文字またはNoPaddingを使用して、encと同一の新しいエンコーディングを作成します。
// パディング文字は'\r'または'\n'ではなく、エンコーディングのアルファベットに含まれていない必要があり、'\xff'以下のルーンである必要があります。
// '\x7f'より上のパディング文字は、コードポイントのUTF-8表現を使用する代わりに、その正確なバイト値としてエンコードされます。
func (enc Encoding) WithPadding(padding rune) *Encoding

// Encodeは、エンコーディングencを使用してsrcをエンコードし、
// [Encoding.EncodedLen](len(src))バイトをdstに書き込みます。
//
// エンコーディングは、出力を8バイトの倍数にパディングするため、
// 大量のデータストリームの個々のブロックにEncodeを使用することは適切ではありません。
// 代わりに [NewEncoder] を使用してください。
func (enc *Encoding) Encode(dst, src []byte)

// EncodeToStringは、srcのbase32エンコーディングを返します。
func (enc *Encoding) EncodeToString(src []byte) string

// NewEncoderは、新しいbase32ストリームエンコーダーを返します。
// 返されたライターに書き込まれたデータはencを使用してエンコードされ、その後wに書き込まれます。
// Base32エンコーディングは5バイトブロックで動作します。
// 書き込みが完了したら、呼び出し元は、部分的に書き込まれたブロックをフラッシュするために返されたエンコーダーを閉じる必要があります。
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

// EncodedLenは、長さnの入力バッファのbase32エンコーディングのバイト数を返します。
func (enc *Encoding) EncodedLen(n int) int

type CorruptInputError int64

func (e CorruptInputError) Error() string

// Decodeは、エンコーディングencを使用してsrcをデコードし、
// [Encoding.DecodedLen](len(src))バイトをdstに書き込みます。
// srcに無効なbase32データが含まれている場合、
// 書き込まれたバイト数と [CorruptInputError] を返します。
// 改行文字（\rおよび\n）は無視されます。
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

// DecodeStringは、base32文字列sによって表されるバイト列を返します。
func (enc *Encoding) DecodeString(s string) ([]byte, error)

// NewDecoderは、新しいbase32ストリームデコーダーを構築します。
func NewDecoder(enc *Encoding, r io.Reader) io.Reader

// DecodedLenは、nバイトのbase32エンコードされたデータに対応するデコードされたデータの最大バイト数を返します。
func (enc *Encoding) DecodedLen(n int) int
