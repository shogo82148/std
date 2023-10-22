// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binary

import (
	"github.com/shogo82148/std/io"
)

// MaxVarintLenNは、Nビット整数の可変長エンコードの最大長です。
const (
	MaxVarintLen16 = 3
	MaxVarintLen32 = 5
	MaxVarintLen64 = 10
)

// AppendUvarintは、PutUvarintによって生成されたxのvarintエンコード形式をbufに追加し、拡張されたバッファを返します。
func AppendUvarint(buf []byte, x uint64) []byte

// PutUvarintは、uint64をbufにエンコードし、書き込まれたバイト数を返します。
// バッファが小さすぎる場合、PutUvarintはパニックを引き起こします。
func PutUvarint(buf []byte, x uint64) int

// Uvarintは、bufからuint64をデコードし、その値と読み取られたバイト数（> 0）を返します。
// エラーが発生した場合、値は0で、バイト数nは<= 0です。
//
//	n == 0: バッファが小さすぎます
//	n < 0: 64ビットより大きい値（オーバーフロー）で、-nは読み取られたバイト数です
func Uvarint(buf []byte) (uint64, int)

// AppendVarintは、PutVarintによって生成されたxのvarintエンコード形式をbufに追加し、拡張されたバッファを返します。
func AppendVarint(buf []byte, x int64) []byte

// PutVarintは、int64をbufにエンコードし、書き込まれたバイト数を返します。
// バッファが小さすぎる場合、PutVarintはパニックを引き起こします。
func PutVarint(buf []byte, x int64) int

// Varintは、bufからint64をデコードし、その値と読み取られたバイト数（> 0）を返します。
// エラーが発生した場合、値は0で、バイト数nは<= 0です。
//
//	n == 0: バッファが小さすぎます
//	n < 0: 64ビットより大きい値（オーバーフロー）で、-nは読み取られたバイト数です
func Varint(buf []byte) (int64, int)

// ReadUvarintは、rから符号なし整数を読み取り、uint64として返します。
// エラーがEOFであるのは、バイトが読み込まれなかった場合のみです。
// 一部のバイトが読み込まれた後にEOFが発生した場合、
// ReadUvarintはio.ErrUnexpectedEOFを返します。
func ReadUvarint(r io.ByteReader) (uint64, error)

// ReadVarintは、rから符号付き整数を読み取り、int64として返します。
// エラーがEOFであるのは、バイトが読み込まれなかった場合のみです。
// 一部のバイトが読み込まれた後にEOFが発生した場合、
// ReadVarintはio.ErrUnexpectedEOFを返します。
func ReadVarint(r io.ByteReader) (int64, error)
