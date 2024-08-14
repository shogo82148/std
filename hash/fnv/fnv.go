// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// fnvパッケージは、 Glenn Fowler、Landon Curt Noll、およびPhong Voによって作成された
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

// New32は新しい32ビットFNV-1 [hash.Hash] を返します。
// そのSumメソッドは値をビッグエンディアンのバイト順で配置します。
func New32() hash.Hash32

// New32aは32ビットのFNV-1a [hash.Hash] を新たに作成します。
// そのSumメソッドは値をビッグエンディアンバイト順で表示します。
func New32a() hash.Hash32

// New64は新しい64ビットのFNV-1 [hash.Hash] を返します。
// そのSumメソッドは値をビッグエンディアンのバイトオーダーで配置します。
func New64() hash.Hash64

// New64a は新しい64ビットのFNV-1a [hash.Hash] を返します。
// そのSumメソッドは値をビッグエンディアンのバイト順序で配置します。
func New64a() hash.Hash64

// New128は新しい128ビットのFNV-1 [hash.Hash] を返します。
// Sumメソッドは値をビッグエンディアンのバイトオーダーで配置します。
func New128() hash.Hash

// New128aは新しい128ビットのFNV-1a [hash.Hash] を返します。
// そのSumメソッドは値をビッグエンディアンのバイト順で配置します。
func New128a() hash.Hash
