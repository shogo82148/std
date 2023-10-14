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

<<<<<<< HEAD
// MakeTable は指定された多項式から構築されたTableを返します。
// このTableの内容は変更してはなりません。
func MakeTable(poly uint64) *Table

// NewはTableで表される多項式を使用してCRC-64チェックサムを計算する新しいhash.Hash64を作成します。Sumメソッドは値をビッグエンディアンのバイト順で並べます。返されるHash64は、内部状態をmarshalおよびunmarshalするためのencoding.BinaryMarshalerおよびencoding.BinaryUnmarshalerも実装しています。
=======
// MakeTable returns a [Table] constructed from the specified polynomial.
// The contents of this [Table] must not be modified.
func MakeTable(poly uint64) *Table

// New creates a new hash.Hash64 computing the CRC-64 checksum using the
// polynomial represented by the [Table]. Its Sum method will lay the
// value out in big-endian byte order. The returned Hash64 also
// implements [encoding.BinaryMarshaler] and [encoding.BinaryUnmarshaler] to
// marshal and unmarshal the internal state of the hash.
>>>>>>> upstream/master
func New(tab *Table) hash.Hash64

// Updateはpのバイトをcrcに追加した結果を返します。
func Update(crc uint64, tab *Table, p []byte) uint64

<<<<<<< HEAD
// Checksum 関数は、Table が表す多項式を使って、データの CRC-64 チェックサムを返します。
=======
// Checksum returns the CRC-64 checksum of data
// using the polynomial represented by the [Table].
>>>>>>> upstream/master
func Checksum(data []byte, tab *Table) uint64
