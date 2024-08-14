// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// base64パッケージは、RFC 4648で指定されているように、base64エンコーディングを実装します。
package base64

import (
	"github.com/shogo82148/std/io"
)

// Encodingは、64文字のアルファベットによって定義される基数64のエンコーディング/デコーディングスキームです。
// 最も一般的なものは、RFC 4648で定義され、MIME（RFC 2045）およびPEM（RFC 1421）で使用される「base64」エンコーディングです。
// RFC 4648は、+と/の代わりに-と_が使用された標準エンコーディングを定義しています。
type Encoding struct {
	encode    [64]byte
	decodeMap [256]uint8
	padChar   rune
	strict    bool
}

const (
	StdPadding rune = '='
	NoPadding  rune = -1
)

// NewEncodingは、与えられたアルファベットによって定義されるパディングされた新しいエンコーディングを返します。
// アルファベットは、パディング文字またはCR / LF（'\r'、'\n'）を含まない64バイトの文字列でなければなりません。
// アルファベットは、マルチバイトUTF-8に対する特別な処理なしに、バイト値のシーケンスとして扱われます。
// 結果のエンコーディングは、デフォルトのパディング文字（'='）を使用します。
// パディング文字は [Encoding.WithPadding] を介して変更または無効化できます。
func NewEncoding(encoder string) *Encoding

// WithPaddingは、指定されたパディング文字または [NoPadding] を使用して、encと同一の新しいエンコーディングを作成します。
// パディング文字は'\r'または'\n'ではなく、エンコーディングのアルファベットに含まれていない必要があり、'\xff'以下のルーンである必要があります。
// '\x7f'より上のパディング文字は、コードポイントのUTF-8表現を使用する代わりに、その正確なバイト値としてエンコードされます。
func (enc Encoding) WithPadding(padding rune) *Encoding

// Strictは、RFC 4648 Section 3.5 で説明されているように、
// 末尾のパディングビットがゼロであることを要求する厳密なデコードが有効になっているencと同一の新しいエンコーディングを作成します。
//
// ただし、入力はまだ操作可能であり、改行文字（CRおよびLF）は引き続き無視されます。
func (enc Encoding) Strict() *Encoding

// StdEncodingは、RFC 4648で定義されている標準のbase64エンコーディングです。
var StdEncoding = NewEncoding(encodeStd)

// URLEncodingは、RFC 4648で定義されている代替のbase64エンコーディングです。
// 通常、URLやファイル名で使用されます。
var URLEncoding = NewEncoding(encodeURL)

// RawStdEncodingは、RFC 4648 Section 3.2 で定義されている標準の生の、パディングされていないbase64エンコーディングです。
// これは、[StdEncoding] と同じですが、パディング文字が省略されています。
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)

// RawURLEncodingは、RFC 4648で定義されているパディングされていない代替のbase64エンコーディングです。
// 通常、URLやファイル名で使用されます。
// これは、[URLEncoding] と同じですが、パディング文字が省略されています。
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)

// Encodeは、エンコーディングencを使用してsrcをエンコードし、
// [Encoding.EncodedLen](len(src))バイトをdstに書き込みます。
//
// エンコーディングは、出力を4バイトの倍数にパディングするため、
// 大量のデータストリームの個々のブロックにEncodeを使用することは適切ではありません。
// 代わりに [NewEncoder] を使用してください。
func (enc *Encoding) Encode(dst, src []byte)

// AppendEncodeは、base64でエンコードされたsrcをdstに追加し、
// 拡張されたバッファを返します。
func (enc *Encoding) AppendEncode(dst, src []byte) []byte

// EncodeToStringは、srcのbase64エンコーディングを返します。
func (enc *Encoding) EncodeToString(src []byte) string

// NewEncoderは、新しいbase64ストリームエンコーダーを返します。
// 返されたライターに書き込まれたデータはencを使用してエンコードされ、その後wに書き込まれます。
// Base64エンコーディングは4バイトブロックで動作します。
// 書き込みが完了したら、呼び出し元は、部分的に書き込まれたブロックをフラッシュするために返されたエンコーダーを閉じる必要があります。
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser

// EncodedLenは、長さnの入力バッファのbase64エンコーディングのバイト数を返します。
func (enc *Encoding) EncodedLen(n int) int

type CorruptInputError int64

func (e CorruptInputError) Error() string

// AppendDecodeは、base64でデコードされたsrcをdstに追加し、
// 拡張されたバッファを返します。
// 入力が不正な形式の場合、部分的にデコードされたsrcとエラーを返します。
func (enc *Encoding) AppendDecode(dst, src []byte) ([]byte, error)

// DecodeStringは、base64文字列sによって表されるバイト列を返します。
func (enc *Encoding) DecodeString(s string) ([]byte, error)

// Decodeは、エンコーディングencを使用してsrcをデコードし、
// [Encoding.DecodedLen](len(src))バイトをdstに書き込みます。
// srcに無効なbase64データが含まれている場合、
// 書き込まれたバイト数と [CorruptInputError] を返します。
// 改行文字（\rおよび\n）は無視されます。
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

// NewDecoderは、新しいbase64ストリームデコーダーを構築します。
func NewDecoder(enc *Encoding, r io.Reader) io.Reader

// DecodedLenは、nバイトのbase64エンコードされたデータに対応するデコードされたデータの最大バイト数を返します。
func (enc *Encoding) DecodedLen(n int) int
