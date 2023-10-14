// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージfnvは、 Glenn Fowler、Landon Curt Noll、およびPhong Voによって作成された
// FNV-1およびFNV-1aという非暗号化ハッシュ関数を実装しています。
// 詳細は、以下を参照してください。
// https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function.
//
// このパッケージによって返されるすべてのhash.Hashの実装も、
// encoding.BinaryMarshalerとencoding.BinaryUnmarshalerを実装しています。
// これにより、ハッシュの内部状態をマーシャリングおよびアンマーシャリングすることができます。
package fnv

import (
	"github.com/shogo82148/std/hash"
)

<<<<<<< HEAD
// New32は新しい32ビットFNV-1ハッシュ.Hashを返します。
// そのSumメソッドは値をビッグエンディアンのバイト順で配置します。
func New32() hash.Hash32

// New32aは32ビットのFNV-1aハッシュ.Hashを新たに作成します。
// そのSumメソッドは値をビッグエンディアンバイト順で表示します。
func New32a() hash.Hash32

// New64は新しい64ビットのFNV-1ハッシュのhash.Hashを返します。
// そのSumメソッドは値をビッグエンディアンのバイトオーダーで配置します。
func New64() hash.Hash64

// New64a は新しい64ビットのFNV-1aハッシュを返します。
// そのSumメソッドは値をビッグエンディアンのバイト順序で配置します。
func New64a() hash.Hash64

// New128は新しい128ビットのFNV-1ハッシュ.Hashを返します。
// Sumメソッドは値をビッグエンディアンのバイトオーダーで配置します。
func New128() hash.Hash

// New128aは新しい128ビットのFNV-1aハッシュを返します。
// そのSumメソッドは値をビッグエンディアンのバイト順で配置します。
=======
// New32 returns a new 32-bit FNV-1 [hash.Hash].
// Its Sum method will lay the value out in big-endian byte order.
func New32() hash.Hash32

// New32a returns a new 32-bit FNV-1a [hash.Hash].
// Its Sum method will lay the value out in big-endian byte order.
func New32a() hash.Hash32

// New64 returns a new 64-bit FNV-1 [hash.Hash].
// Its Sum method will lay the value out in big-endian byte order.
func New64() hash.Hash64

// New64a returns a new 64-bit FNV-1a [hash.Hash].
// Its Sum method will lay the value out in big-endian byte order.
func New64a() hash.Hash64

// New128 returns a new 128-bit FNV-1 [hash.Hash].
// Its Sum method will lay the value out in big-endian byte order.
func New128() hash.Hash

// New128a returns a new 128-bit FNV-1a [hash.Hash].
// Its Sum method will lay the value out in big-endian byte order.
>>>>>>> upstream/master
func New128a() hash.Hash
