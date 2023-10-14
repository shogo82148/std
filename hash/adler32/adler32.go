// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package adler32はAdler-32チェックサムを実装しています。
//
// これはRFC 1950で定義されています：
//
//	Adler-32はバイトごとに累積される2つの合計で構成されています: s1は
//	すべてのバイトの合計で、s2はすべてのs1の値の合計です。両方の合計は
//	65521でモジュロ演算が行われます。s1は1で初期化され、s2はゼロです。Adler-32
//	チェックサムはs2*65536 + s1として、最も重要なバイトを最初に（ネットワークの）順序で格納されます。
package adler32

import (
	"github.com/shogo82148/std/hash"
)

// Adler-32チェックサムのバイト単位のサイズ。
const Size = 4

// NewはAdler-32チェックサムを計算する新しいhash.Hash32を返します。
// Sumメソッドは値をビッグエンディアンのバイト順で配置します。
// 返されるHash32は、encoding.BinaryMarshalerとencoding.BinaryUnmarshalerも実装しており、
// ハッシュの内部状態をマーシャリングおよびアンマーシャリングすることができます。
func New() hash.Hash32

// ChecksumはdataのAdler-32チェックサムを返します。
func Checksum(data []byte) uint32
