// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crc64は64ビットの巡回冗長検査（CRC-64）チェックサムを実装しています。
// 詳細はhttps://en.wikipedia.org/wiki/Cyclic_redundancy_checkを参照してください。
package crc64

import (
	"github.com/shogo82148/std/hash"
)

// CRC-64のチェックサムのバイト単位のサイズ。
const Size = 8

// 事前定義された多項式。
const (
	// ISO 3309で定義され、HDLCで使用されるISOポリノミアル。
	ISO = 0xD800000000000000

	// ECMA 182で定義されたECMA多項式。
	ECMA = 0xC96C5795D7870F42
)

// Table は効率的な処理のための多項式を表す256単語のテーブルです。
type Table [256]uint64

// MakeTable は指定された多項式から構築された [Table] を返します。
// このTableの内容は変更してはなりません。
func MakeTable(poly uint64) *Table

// Newは [Table] で表される多項式を使用してCRC-64チェックサムを計算する新しいhash.Hash64を作成します。Sumメソッドは値をビッグエンディアンのバイト順で並べます。返されるHash64は、内部状態をmarshalおよびunmarshalするための [encoding.BinaryMarshaler] および [encoding.BinaryUnmarshaler] も実装しています。
func New(tab *Table) hash.Hash64

// Updateはpのバイトをcrcに追加した結果を返します。
func Update(crc uint64, tab *Table, p []byte) uint64

// Checksum 関数は、 [Table] が表す多項式を使って、データの CRC-64 チェックサムを返します。
func Checksum(data []byte, tab *Table) uint64
