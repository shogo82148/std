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

// NewはMD5チェックサムを計算する新しい [hash.Hash] を返します。
// このハッシュは、内部状態をマーシャルおよびアンマーシャルするために
// [encoding.BinaryMarshaler]、[encoding.AppendBinary]、および
// [encoding.BinaryUnmarshaler] も実装しています。
func New() hash.Hash

// Sum はデータのMD5ハッシュ値を返します。
func Sum(data []byte) [Size]byte
