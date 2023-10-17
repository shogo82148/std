// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crc32は32ビットの巡回冗長検査 (CRC-32) チェックサムを実装しています。
// 詳細はhttps://en.wikipedia.org/wiki/Cyclic_redundancy_checkを参照してください。
//
// 多項式は、LSBファースト形式、または反転表現としても表されます。
//
// 詳細はhttps://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Reversed_representations_and_reciprocal_polynomialsを参照してください。
package crc32

import (
	"github.com/shogo82148/std/hash"
)

// CRC-32のチェックサムのサイズ（バイト単位）。
const Size = 4

// 事前に定義された多項式。
const (

	// IEEEは、断然最も一般的なCRC-32多項式です。
	// ethernet (IEEE 802.3)、v.42、fddi、gzip、zip、pngなどで使用されています。
	IEEE = 0xedb88320

	// キャスタニョーリの多項式、iSCSIで使用されています。
	// IEEEよりも優れたエラー検出特性を持っています。
	// https://dx.doi.org/10.1109/26.231911
	Castagnoli = 0x82f63b78

	// クープマンの多項式。
	// IEEEよりもエラー検出性能が優れています。
	// https://dx.doi.org/10.1109/DSN.2002.1028931
	Koopman = 0xeb31d82e
)

// Tableは、効率的な処理のための多項式を表す256ワードのテーブルです。
type Table [256]uint32

// IEEETableは [IEEE] ポリノミアルのテーブルです。
var IEEETable = simpleMakeTable(IEEE)

// MakeTableは指定された多項式から構築された [Table] を返します。
// この [Table] の内容は変更してはいけません。
func MakeTable(poly uint32) *Table

// Newは [Table] によって表現される多項式を使用してCRC-32チェックサムを計算する新しい [hash.Hash32] を作成します。
// そのSumメソッドはビッグエンディアンのバイト順で値を配置します。
// 返されるHash32は、内部状態のマーシャリングとアンマーシャリングを実装するため、 [encoding.BinaryMarshaler] と [encoding.BinaryUnmarshaler] も実装しています。
func New(tab *Table) hash.Hash32

// NewIEEEは、 [IEEE] 多項式を使用してCRC-32チェックサムを計算する新しい [hash.Hash32] を作成します。そのSumメソッドは、値をビッグエンディアンのバイト順でレイアウトします。
// 返されるHash32は、 [encoding.BinaryMarshaler] および [encoding.BinaryUnmarshaler] も実装しており、ハッシュの内部状態をマーシャルおよびアンマーシャルすることができます。
func NewIEEE() hash.Hash32

// Updateはpのバイトをcrcに追加した結果を返します。
func Update(crc uint32, tab *Table, p []byte) uint32

// Checksumは [Table] で表されるポリノミアルを使用して、
// dataのCRC-32チェックサムを返します。
func Checksum(data []byte, tab *Table) uint32

// ChecksumIEEEは、 [IEEE] 多項式を使用してデータのCRC-32チェックサムを返します。
func ChecksumIEEE(data []byte) uint32
