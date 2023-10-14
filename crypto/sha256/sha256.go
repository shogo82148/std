// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージsha256は、FIPS 180-4で定義されたSHA224およびSHA256ハッシュアルゴリズムを実装しています。
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

<<<<<<< HEAD
// NewはSHA256ハッシュチェックサムを計算する新しいhash.Hashを返します。Hashは
// encoding.BinaryMarshalerおよびencoding.BinaryUnmarshalerも実装しており、内部の
// ハッシュの状態をマーシャリングおよびアンマーシャリングすることができます。
=======
// New returns a new hash.Hash computing the SHA256 checksum. The Hash
// also implements [encoding.BinaryMarshaler] and
// [encoding.BinaryUnmarshaler] to marshal and unmarshal the internal
// state of the hash.
>>>>>>> upstream/master
func New() hash.Hash

// New224はSHA224チェックサムを計算する新しいhash.Hashを返します。
func New224() hash.Hash

// Sum256はdataのSHA256チェックサムを返します。
func Sum256(data []byte) [Size]byte

// Sum224はデータのSHA224ハッシュ値を返します。
func Sum224(data []byte) [Size224]byte
