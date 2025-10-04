// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// sha1パッケージは、RFC 3174で定義されているSHA-1ハッシュアルゴリズムを実装しています。
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

<<<<<<< HEAD
// NewはSHA1チェックサムを計算する新しいhash.Hashを返します。
// このハッシュは、内部状態をマーシャルおよびアンマーシャルするために
// [encoding.BinaryMarshaler]および[encoding.BinaryUnmarshaler]も実装しています。
=======
// New returns a new [hash.Hash] computing the SHA1 checksum. The Hash
// also implements [encoding.BinaryMarshaler], [encoding.BinaryAppender] and
// [encoding.BinaryUnmarshaler] to marshal and unmarshal the internal
// state of the hash.
>>>>>>> upstream/release-branch.go1.25
func New() hash.Hash

// SumはデータのSHA-1チェックサムを返します。
func Sum(data []byte) [Size]byte
