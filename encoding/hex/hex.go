// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// hexパッケージは16進数のエンコードとデコードを実装します。
package hex

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// EncodedLenはn個の元のバイトのエンコーディングの長さを返します。
// 具体的には、n * 2を返します。
func EncodedLen(n int) int

// Encodeは、srcを [EncodedLen](len(src))バイトのdstにエンコードします。
// 便宜上、dstに書き込まれたバイト数を返しますが、この値は常に [EncodedLen](len(src))です。
// Encodeは16進数エンコーディングを実装します。
func Encode(dst, src []byte) int

// AppendEncodeは、16進数でエンコードされたsrcをdstに追加し、
// 拡張されたバッファを返します。
func AppendEncode(dst, src []byte) []byte

// ErrLengthは、[Decode] または [DecodeString] を使用して奇数長の入力をデコードしようとする試みを報告します。
// ストリームベースのDecoderは、ErrLengthの代わりに [io.ErrUnexpectedEOF] を返します。
var ErrLength = errors.New("encoding/hex: odd length hex string")

// InvalidByteError の値は、16 進数文字列に無効なバイトが含まれている場合のエラーを記述します。
type InvalidByteError byte

func (e InvalidByteError) Error() string

// DecodedLenはxソースバイトのデコード結果の長さを返します。
// 具体的には、x / 2 を返します。
func DecodedLen(x int) int

// Decodeは、srcを [DecodedLen](len(src))バイトにデコードし、
// dstに書き込まれた実際のバイト数を返します。
//
// Decodeは、srcが16進文字のみを含み、かつsrcの長さが偶数であることを期待しています。
// もし入力が不正な場合、Decodeはエラーが発生する前にデコードされたバイト数を返します。
func Decode(dst, src []byte) (int, error)

// AppendDecodeは、16進数でデコードされたsrcをdstに追加し、
// 拡張されたバッファを返します。
// 入力が不正な形式の場合、部分的にデコードされたsrcとエラーを返します。
func AppendDecode(dst, src []byte) ([]byte, error)

// EncodeToStringは、srcの16進数エンコーディングを返します。
func EncodeToString(src []byte) string

// DecodeStringは16進数の文字列sによって表されるバイトを返します。
//
// DecodeStringは、srcが16進数の文字のみを含み、かつ偶数の長さであることを期待しています。
// 入力が不正な場合、DecodeStringはエラーが発生する前にデコードされたバイトを返します。
func DecodeString(s string) ([]byte, error)

// Dumpは指定されたデータの16進ダンプを含む文字列を返します。16進ダンプの形式は、コマンドラインの`hexdump -C`の出力と一致します。
func Dump(data []byte) string

// NewEncoderは、wに小文字の16進数文字を書き込む [io.Writer] を返します。
func NewEncoder(w io.Writer) io.Writer

// NewDecoderは、rから16進数文字をデコードする [io.Reader] を返します。
// NewDecoderは、rが偶数個の16進数文字のみを含むことを期待します。
func NewDecoder(r io.Reader) io.Reader

// Dumperは、書き込まれたすべてのデータの16進ダンプをwに書き込む [io.WriteCloser] を返します。
// ダンプの形式は、コマンドライン上の`hexdump -C`の出力と一致します。
func Dumper(w io.Writer) io.WriteCloser
