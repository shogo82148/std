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

<<<<<<< HEAD
// NewEncoding returns a new padded Encoding defined by the given alphabet,
// which must be a 32-byte string that contains unique byte values and
// does not contain the padding character or CR / LF ('\r', '\n').
// The alphabet is treated as a sequence of byte values
// without any special treatment for multi-byte UTF-8.
// The resulting Encoding uses the default padding character ('='),
// which may be changed or disabled via [Encoding.WithPadding].
func NewEncoding(encoder string) *Encoding

// StdEncoding is the standard base32 encoding, as defined in RFC 4648.
var StdEncoding = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")

// HexEncoding is the “Extended Hex Alphabet” defined in RFC 4648.
// It is typically used in DNS.
var HexEncoding = NewEncoding("0123456789ABCDEFGHIJKLMNOPQRSTUV")

// WithPadding creates a new encoding identical to enc except
// with a specified padding character, or NoPadding to disable padding.
// The padding character must not be '\r' or '\n',
// must not be contained in the encoding's alphabet,
// must not be negative, and must be a rune equal or below '\xff'.
// Padding characters above '\x7f' are encoded as their exact byte value
// rather than using the UTF-8 representation of the codepoint.
func (enc Encoding) WithPadding(padding rune) *Encoding

// Encode encodes src using the encoding enc,
// writing [Encoding.EncodedLen](len(src)) bytes to dst.
//
// The encoding pads the output to a multiple of 8 bytes,
// so Encode is not appropriate for use on individual blocks
// of a large data stream. Use [NewEncoder] instead.
func (enc *Encoding) Encode(dst, src []byte)

// AppendEncode appends the base32 encoded src to dst
// and returns the extended buffer.
func (enc *Encoding) AppendEncode(dst, src []byte) []byte

// EncodeToString returns the base32 encoding of src.
=======
// NewEncodingは、与えられたアルファベットによって定義される新しいエンコーディングを返します。
// アルファベットは32バイトの文字列でなければなりません。
// アルファベットは、マルチバイトUTF-8に対する特別な処理なしに、バイト値のシーケンスとして扱われます。
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
// EncodedLen(len(src))バイトをdstに書き込みます。
//
// エンコーディングは、出力を8バイトの倍数にパディングするため、
// 大量のデータストリームの個々のブロックにEncodeを使用することは適切ではありません。
// 代わりにNewEncoder()を使用してください。
func (enc *Encoding) Encode(dst, src []byte)

// EncodeToStringは、srcのbase32エンコーディングを返します。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// Decode decodes src using the encoding enc. It writes at most
// [Encoding.DecodedLen](len(src)) bytes to dst and returns the number of bytes
// written. If src contains invalid base32 data, it will return the
// number of bytes successfully written and [CorruptInputError].
// Newline characters (\r and \n) are ignored.
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

// AppendDecode appends the base32 decoded src to dst
// and returns the extended buffer.
// If the input is malformed, it returns the partially decoded src and an error.
func (enc *Encoding) AppendDecode(dst, src []byte) ([]byte, error)

// DecodeString returns the bytes represented by the base32 string s.
=======
// Decodeは、エンコーディングencを使用してsrcをデコードし、
// DecodedLen(len(src))バイトをdstに書き込みます。
// srcに無効なbase32データが含まれている場合、
// 書き込まれたバイト数とCorruptInputErrorを返します。
// 改行文字（\rおよび\n）は無視されます。
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)

// DecodeStringは、base32文字列sによって表されるバイト列を返します。
>>>>>>> release-branch.go1.21
func (enc *Encoding) DecodeString(s string) ([]byte, error)

// NewDecoderは、新しいbase32ストリームデコーダーを構築します。
func NewDecoder(enc *Encoding, r io.Reader) io.Reader

// DecodedLenは、nバイトのbase32エンコードされたデータに対応するデコードされたデータの最大バイト数を返します。
func (enc *Encoding) DecodedLen(n int) int
