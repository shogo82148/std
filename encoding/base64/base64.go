// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package base64は、RFC 4648で指定されているように、base64エンコーディングを実装します。
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

<<<<<<< HEAD
// NewEncoding returns a new padded Encoding defined by the given alphabet,
// which must be a 64-byte string that contains unique byte values and
// does not contain the padding character or CR / LF ('\r', '\n').
// The alphabet is treated as a sequence of byte values
// without any special treatment for multi-byte UTF-8.
// The resulting Encoding uses the default padding character ('='),
// which may be changed or disabled via [Encoding.WithPadding].
func NewEncoding(encoder string) *Encoding

// WithPadding creates a new encoding identical to enc except
// with a specified padding character, or [NoPadding] to disable padding.
// The padding character must not be '\r' or '\n',
// must not be contained in the encoding's alphabet,
// must not be negative, and must be a rune equal or below '\xff'.
// Padding characters above '\x7f' are encoded as their exact byte value
// rather than using the UTF-8 representation of the codepoint.
=======
// NewEncodingは、与えられたアルファベットによって定義されるパディングされた新しいエンコーディングを返します。
// アルファベットは、パディング文字またはCR / LF（'\r'、'\n'）を含まない64バイトの文字列でなければなりません。
// アルファベットは、マルチバイトUTF-8に対する特別な処理なしに、バイト値のシーケンスとして扱われます。
// 結果のエンコーディングは、デフォルトのパディング文字（'='）を使用します。
// パディング文字はWithPaddingを介して変更または無効化できます。
func NewEncoding(encoder string) *Encoding

// WithPaddingは、指定されたパディング文字またはNoPaddingを使用して、encと同一の新しいエンコーディングを作成します。
// パディング文字は'\r'または'\n'ではなく、エンコーディングのアルファベットに含まれていない必要があり、'\xff'以下のルーンである必要があります。
// '\x7f'より上のパディング文字は、コードポイントのUTF-8表現を使用する代わりに、その正確なバイト値としてエンコードされます。
>>>>>>> release-branch.go1.21
func (enc Encoding) WithPadding(padding rune) *Encoding

// Strictは、RFC 4648セクション3.5で説明されているように、
// 末尾のパディングビットがゼロであることを要求する厳密なデコードが有効になっているencと同一の新しいエンコーディングを作成します。
//
// ただし、入力はまだ操作可能であり、改行文字（CRおよびLF）は引き続き無視されます。
func (enc Encoding) Strict() *Encoding

<<<<<<< HEAD
// StdEncoding is the standard base64 encoding, as defined in RFC 4648.
var StdEncoding = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

// URLEncoding is the alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
var URLEncoding = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

// RawStdEncoding is the standard raw, unpadded base64 encoding,
// as defined in RFC 4648 section 3.2.
// This is the same as [StdEncoding] but omits padding characters.
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)

// RawURLEncoding is the unpadded alternate base64 encoding defined in RFC 4648.
// It is typically used in URLs and file names.
// This is the same as [URLEncoding] but omits padding characters.
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)

// Encode encodes src using the encoding enc,
// writing [Encoding.EncodedLen](len(src)) bytes to dst.
//
// The encoding pads the output to a multiple of 4 bytes,
// so Encode is not appropriate for use on individual blocks
// of a large data stream. Use [NewEncoder] instead.
func (enc *Encoding) Encode(dst, src []byte)

// AppendEncode appends the base64 encoded src to dst
// and returns the extended buffer.
func (enc *Encoding) AppendEncode(dst, src []byte) []byte

// EncodeToString returns the base64 encoding of src.
=======
// StdEncodingは、RFC 4648で定義されている標準のbase64エンコーディングです。
var StdEncoding = NewEncoding(encodeStd)

// URLEncodingは、RFC 4648で定義されている代替のbase64エンコーディングです。
// 通常、URLやファイル名で使用されます。
var URLEncoding = NewEncoding(encodeURL)

// RawStdEncodingは、RFC 4648セクション3.2で定義されている標準の生の、パディングされていないbase64エンコーディングです。
// これは、StdEncodingと同じですが、パディング文字が省略されています。
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)

// RawURLEncodingは、RFC 4648で定義されているパディングされていない代替のbase64エンコーディングです。
// 通常、URLやファイル名で使用されます。
// これは、URLEncodingと同じですが、パディング文字が省略されています。
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)

// Encodeは、エンコーディングencを使用してsrcをエンコードし、
// EncodedLen(len(src))バイトをdstに書き込みます。
//
// エンコーディングは、出力を4バイトの倍数にパディングするため、
// 大量のデータストリームの個々のブロックにEncodeを使用することは適切ではありません。
// 代わりにNewEncoder()を使用してください。
func (enc *Encoding) Encode(dst, src []byte)

// EncodeToStringは、srcのbase64エンコーディングを返します。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// AppendDecode appends the base64 decoded src to dst
// and returns the extended buffer.
// If the input is malformed, it returns the partially decoded src and an error.
func (enc *Encoding) AppendDecode(dst, src []byte) ([]byte, error)

// DecodeString returns the bytes represented by the base64 string s.
func (enc *Encoding) DecodeString(s string) ([]byte, error)

// Decode decodes src using the encoding enc. It writes at most
// [Encoding.DecodedLen](len(src)) bytes to dst and returns the number of bytes
// written. If src contains invalid base64 data, it will return the
// number of bytes successfully written and [CorruptInputError].
// New line characters (\r and \n) are ignored.
=======
// DecodeStringは、base64文字列sによって表されるバイト列を返します。
func (enc *Encoding) DecodeString(s string) ([]byte, error)

// Decodeは、エンコーディングencを使用してsrcをデコードし、
// DecodedLen(len(src))バイトをdstに書き込みます。
// srcに無効なbase64データが含まれている場合、
// 書き込まれたバイト数とCorruptInputErrorを返します。
// 改行文字（\rおよび\n）は無視されます。
>>>>>>> release-branch.go1.21
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

// NewDecoderは、新しいbase64ストリームデコーダーを構築します。
func NewDecoder(enc *Encoding, r io.Reader) io.Reader

// DecodedLenは、nバイトのbase64エンコードされたデータに対応するデコードされたデータの最大バイト数を返します。
func (enc *Encoding) DecodedLen(n int) int
