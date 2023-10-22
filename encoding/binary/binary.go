// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// binaryパッケージは、数値とバイトシーケンスの間の単純な変換、
// およびvarintのエンコードとデコードを実装します。
//
// 数値は、固定サイズの値を読み書きすることによって変換されます。
// 固定サイズの値は、固定サイズの算術型（bool、int8、uint8、int16、float32、complex64など）
// または固定サイズの値のみを含む配列または構造体です。
//
// varint関数は、可変長エンコーディングを使用して単一の整数値をエンコードおよびデコードします。
// より小さい値は、より少ないバイトを必要とします。
// 仕様については、以下を参照してください。
// https://developers.google.com/protocol-buffers/docs/encoding。
//
<<<<<<< HEAD
// This package favors simplicity over efficiency. Clients that require
// high-performance serialization, especially for large data structures,
// should look at more advanced solutions such as the [encoding/gob]
// package or [google.golang.org/protobuf] for protocol buffers.
=======
// このパッケージは、効率よりもシンプルさを重視しています。
// 特に大規模なデータ構造に対して高性能なシリアル化が必要なクライアントは、
// encoding/gobパッケージやプロトコルバッファなどのより高度なソリューションを検討する必要があります。
>>>>>>> release-branch.go1.21
package binary

import (
	"github.com/shogo82148/std/io"
)

<<<<<<< HEAD
// A ByteOrder specifies how to convert byte slices into
// 16-, 32-, or 64-bit unsigned integers.
//
// It is implemented by [LittleEndian], [BigEndian], and [NativeEndian].
=======
// ByteOrderは、バイトスライスを16、32、または64ビットの符号なし整数に変換する方法を指定します。
>>>>>>> release-branch.go1.21
type ByteOrder interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
	String() string
}

<<<<<<< HEAD
// AppendByteOrder specifies how to append 16-, 32-, or 64-bit unsigned integers
// into a byte slice.
//
// It is implemented by [LittleEndian], [BigEndian], and [NativeEndian].
=======
// AppendByteOrderは、16、32、または64ビットの符号なし整数をバイトスライスに追加する方法を指定します。
>>>>>>> release-branch.go1.21
type AppendByteOrder interface {
	AppendUint16([]byte, uint16) []byte
	AppendUint32([]byte, uint32) []byte
	AppendUint64([]byte, uint64) []byte
	String() string
}

<<<<<<< HEAD
// LittleEndian is the little-endian implementation of [ByteOrder] and [AppendByteOrder].
var LittleEndian littleEndian

// BigEndian is the big-endian implementation of [ByteOrder] and [AppendByteOrder].
=======
// LittleEndianは、ByteOrderおよびAppendByteOrderのリトルエンディアン実装です。
var LittleEndian littleEndian

// BigEndianは、ByteOrderおよびAppendByteOrderのビッグエンディアン実装です。
>>>>>>> release-branch.go1.21
var BigEndian bigEndian

// Readは、rからdataに対して構造化されたバイナリデータを読み取ります。
// dataは、固定サイズの値または固定サイズの値のスライスへのポインタである必要があります。
// rから読み取られたバイトは、指定されたバイトオーダーを使用してデコードされ、
// dataの連続するフィールドに書き込まれます。
// ブール値をデコードする場合、ゼロバイトはfalseとしてデコードされ、
// それ以外の非ゼロバイトはtrueとしてデコードされます。
// 構造体に読み込む場合、ブランク（_）フィールド名を持つフィールドのデータはスキップされます。
// つまり、パディングにブランクフィールド名を使用できます。
// 構造体に読み込む場合、すべての非ブランクフィールドはエクスポートされている必要があります。
// そうでない場合、Readはパニックを引き起こす可能性があります。
//
<<<<<<< HEAD
// The error is [io.EOF] only if no bytes were read.
// If an [io.EOF] happens after reading some but not all the bytes,
// Read returns [io.ErrUnexpectedEOF].
=======
// エラーがEOFであるのは、バイトが読み込まれなかった場合のみです。
// 一部のバイトが読み込まれた後にEOFが発生した場合、
// ReadはErrUnexpectedEOFを返します。
>>>>>>> release-branch.go1.21
func Read(r io.Reader, order ByteOrder, data any) error

// Writeは、データのバイナリ表現をwに書き込みます。
// データは、固定サイズの値または固定サイズの値のスライス、またはそのようなデータへのポインタである必要があります。
// ブール値は1がtrue、0がfalseとして1バイトでエンコードされます。
// wに書き込まれたバイトは、指定されたバイトオーダーを使用してエンコードされ、
// データの連続するフィールドから読み取られます。
// 構造体を書き込む場合、ブランク（_）フィールド名を持つフィールドのデータはゼロ値で書き込まれます。
func Write(w io.Writer, order ByteOrder, data any) error

<<<<<<< HEAD
// Size returns how many bytes [Write] would generate to encode the value v, which
// must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.
// If v is neither of these, Size returns -1.
=======
// Sizeは、値vをエンコードするためにWriteが生成するバイト数を返します。
// vは、固定サイズの値または固定サイズの値のスライス、またはそのようなデータへのポインタである必要があります。
// vがこれらのいずれでもない場合、Sizeは-1を返します。
>>>>>>> release-branch.go1.21
func Size(v any) int
