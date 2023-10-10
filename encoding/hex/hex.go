// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hex は16進数のエンコードとデコードを実装します。
package hex

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// EncodedLenはn個の元のバイトのエンコーディングの長さを返します。
// 具体的には、n * 2を返します。
func EncodedLen(n int) int

// Encodeは、srcをdstのEncodedLen(len(src))バイトにエンコードします。
// 便宜上、dstに書き込まれたバイト数を返しますが、この値は常にEncodedLen(len(src))です。
// Encodeは16進数エンコーディングを実装しています。
func Encode(dst, src []byte) int

// ErrLengthは、DecodeまたはDecodeStringを使用して奇数の長さの入力をデコードしようとしたことを報告します。
// ストリームベースのデコーダーは、ErrLengthの代わりにio.ErrUnexpectedEOFを返します。
var ErrLength = errors.New("encoding/hex: odd length hex string")

// InvalidByteError の値は、16 進数文字列に無効なバイトが含まれている場合のエラーを記述します。
type InvalidByteError byte

func (e InvalidByteError) Error() string

// DecodedLenはxソースバイトのデコード結果の長さを返します。
// 具体的には、x / 2 を返します。
func DecodedLen(x int) int

// DecodeはsrcをDecodedLen（len（src））バイトにデコードし、
// dstに書き込まれた実際のバイト数を返します。
//
// Decodeは、srcが16進文字のみを含み、かつsrcの長さが偶数であることを期待しています。
// もし入力が不正な場合、Decodeはエラーが発生する前にデコードされたバイト数を返します。
func Decode(dst, src []byte) (int, error)

// EncodeToStringはsrcの16進数エンコーディングを返します。
func EncodeToString(src []byte) string

// DecodeStringは16進数の文字列sによって表されるバイトを返します。
//
// DecodeStringは、srcが16進数の文字のみを含み、かつ偶数の長さであることを期待しています。
// 入力が不正な場合、DecodeStringはエラーが発生する前にデコードされたバイトを返します。
func DecodeString(s string) ([]byte, error)

// Dumpは指定されたデータの16進ダンプを含む文字列を返します。16進ダンプの形式は、コマンドラインの`hexdump -C`の出力と一致します。
func Dump(data []byte) string

// NewEncoderは、小文字の16進数文字をwに書き込むio.Writerを返します。
func NewEncoder(w io.Writer) io.Writer

// NewDecoderは、rから16進数の文字列をデコードするio.Readerを返します。
// NewDecoderは、rに偶数個の16進数の文字列のみが含まれていることを期待しています。
func NewDecoder(r io.Reader) io.Reader

// Dumperは書き込まれた全データの16進ダンプをwに書き込むWriteCloserを返します。ダンプのフォーマットはコマンドライン上での `hexdump -C` の出力に一致します。
func Dumper(w io.Writer) io.WriteCloser
