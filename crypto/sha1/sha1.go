// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージsha1は、RFC 3174で定義されているSHA-1ハッシュアルゴリズムを実装しています。
//
// SHA-1は暗号学的に破られており、セキュアなアプリケーションには使用すべきではありません。
package sha1

import (
	"github.com/shogo82148/std/hash"
)

// SHA-1チェックサムのサイズ（バイト単位）。
const Size = 20

// SHA-1のブロックサイズ（バイト単位）です。
const BlockSize = 64

// New512_224はSHA1チェックサムを計算する新しい [hash.Hash] を返します。
// このハッシュは、内部状態をマーシャルおよびアンマーシャルするために
// [encoding.BinaryMarshaler]、[encoding.BinaryAppender]、および
// [encoding.BinaryUnmarshaler] も実装しています。
func New() hash.Hash

// SumはデータのSHA-1チェックサムを返します。
func Sum(data []byte) [Size]byte
