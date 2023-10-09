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

// NewはSHA1チェックサムを計算する新しいhash.Hashを返します。Hashはまた、encoding.BinaryMarshalerとencoding.BinaryUnmarshalerを実装しており、ハッシュの内部状態をマーシャリングおよびアンマーシャリングすることができます。
func New() hash.Hash

// SumはデータのSHA-1チェックサムを返します。
func Sum(data []byte) [Size]byte
