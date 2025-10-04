// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen.go -output md5block.go
//go:generateコマンドを使用して、gen.goを実行し、md5block.goに出力します。

// md5 パッケージは、RFC 1321で定義されたMD5ハッシュアルゴリズムを実装します。
//
// MD5は暗号学的に破られており、安全なアプリケーションには使用されるべきではありません。
package md5

import (
	"github.com/shogo82148/std/hash"
)

// MD5チェックサムのバイト数。
const Size = 16

// MD5のブロックサイズ（バイト単位）。
const BlockSize = 64

<<<<<<< HEAD
// NewはMD5チェックサムを計算する新しいhash.Hashを返します。
// このハッシュは、内部状態をマーシャルおよびアンマーシャルするために
// [encoding.BinaryMarshaler]および[encoding.BinaryUnmarshaler]も実装しています。
=======
// New returns a new [hash.Hash] computing the MD5 checksum. The Hash
// also implements [encoding.BinaryMarshaler], [encoding.BinaryAppender] and
// [encoding.BinaryUnmarshaler] to marshal and unmarshal the internal
// state of the hash.
>>>>>>> upstream/release-branch.go1.25
func New() hash.Hash

// Sum はデータのMD5ハッシュ値を返します。
func Sum(data []byte) [Size]byte
