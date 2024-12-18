// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// sha256パッケージは、FIPS 180-4で定義されたSHA224およびSHA256ハッシュアルゴリズムを実装しています。
package sha256

import (
	"github.com/shogo82148/std/hash"
)

// SHA256のチェックサムのバイト数。
const Size = 32

// SHA224のチェックサムのサイズ（バイト単位）
const Size224 = 28

// SHA256とSHA224のブロックサイズ（バイト単位）です。
const BlockSize = 64

// NewはSHA256チェックサムを計算する新しいhash.Hashを返します。
// このハッシュは、内部状態をマーシャルおよびアンマーシャルするために
// [encoding.BinaryMarshaler] および [encoding.BinaryUnmarshaler] も実装しています。
func New() hash.Hash

// New224はSHA224チェックサムを計算する新しいhash.Hashを返します。
func New224() hash.Hash

// Sum256はdataのSHA256チェックサムを返します。
func Sum256(data []byte) [Size]byte

// Sum224はデータのSHA224ハッシュ値を返します。
func Sum224(data []byte) [Size224]byte
