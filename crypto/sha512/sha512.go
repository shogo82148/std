// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// sha512パッケージは、FIPS 180-4で定義されているSHA-384、SHA-512、SHA-512/224、およびSHA-512/256のハッシュアルゴリズムを実装しています。
//
// このパッケージによって返されるすべてのhash.Hash実装は、
// encoding.BinaryMarshalerとencoding.BinaryUnmarshalerも実装しており、
// ハッシュの内部状態をマーシャリングおよびアンマーシャリングすることができます。
package sha512

import (
	"github.com/shogo82148/std/hash"
)

const (
	// SizeはSHA-512のチェックサムのバイト単位のサイズです。
	Size = 64

	// Size224はSHA-512/224のチェックサムのサイズ（バイト単位）です。
	Size224 = 28

	// Size256はSHA-512/256チェックサムのバイト単位でのサイズです。
	Size256 = 32

	// Size384はSHA-384のチェックサムのバイト数です。
	Size384 = 48

	// BlockSizeは、SHA-512/224、SHA-512/256、SHA-384、およびSHA-512ハッシュ関数のブロックサイズ（バイト単位）です。
	BlockSize = 128
)

// NewはSHA-512チェックサムを計算する新しいhash.Hashを返します。
func New() hash.Hash

// New512_224はSHA-512/224チェックサムを計算する新しいhash.Hashを返します。
func New512_224() hash.Hash

// New512_256はSHA-512/256チェックサムを計算する新しいhash.Hashを返します。
func New512_256() hash.Hash

// New384はSHA-384チェックサムを計算する新しい [hash.Hash] を返します。このHashは
// [encoding.BinaryMarshaler]、[encoding.BinaryAppender]、および
// [encoding.BinaryUnmarshaler] も実装し、ハッシュの内部状態を
// マーシャル・アンマーシャルします。
func New384() hash.Hash

// Sum512は、データのSHA512ハッシュ値を返します。
func Sum512(data []byte) [Size]byte

// Sum384はデータのSHA384チェックサムを返します。
func Sum384(data []byte) [Size384]byte

// Sum512_224は、データのSum512/224チェックサムを返します。
func Sum512_224(data []byte) [Size224]byte

// Sum512_256はデータのSum512/256チェックサムを返します。
func Sum512_256(data []byte) [Size256]byte
