// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// suffixarrayパッケージは、インメモリの接尾辞配列を使用して対数時間での部分文字列検索を実装します。
//
// 使用例:
//
//	// データに対してインデックスを作成する
//	index := suffixarray.New(data)
//
//	// バイトスライス s を検索する
//	offsets1 := index.Lookup(s, -1) // data 内の s の出現インデックスのリスト
//	offsets2 := index.Lookup(s, 3)  // data 内で最大3つのインデックスでの s の出現リスト

package suffixarray

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/regexp"
)

// Indexは高速な部分文字列検索のための接尾辞配列を実装しています。
type Index struct {
	data []byte
	sa   ints
}

// Newはデータのために新しい [Index] を作成します。
// [Index] の作成時間はO(N)であり、Nはdataの長さです。
func New(data []byte) *Index

// Readはrからインデックスをxに読み込む。xはnilであってはならない。
func (x *Index) Read(r io.Reader) error

// Writeはインデックスxをwに書き込みます。
func (x *Index) Write(w io.Writer) error

// Bytes はインデックスが作成されたデータを返します。
// このデータは変更してはいけません。
func (x *Index) Bytes() []byte

// Lookupは、バイト文字列sがインデックス付きデータに出現する最大n箇所のインデックスの非ソートリストを返します。n < 0の場合、すべての出現箇所が返されます。
// sが空である、sが見つからない、またはn == 0の場合、結果はnilです。
// ルックアップ時間はO(log(N)*len(s) + len(result))であり、Nはインデックス付きデータのサイズです。
func (x *Index) Lookup(s []byte, n int) (result []int)

// FindAllIndexは正規表現rの非重複マッチのソートされたリストを返します。
// マッチは、x.Bytes()の一致するスライスを指定するインデックスのペアです。
// nが0未満の場合、すべてのマッチが連続して返されます。
// そうでない場合、最大でnマッチが返される可能性がありますが、連続しているとは限りません。
// マッチがない場合、またはn == 0の場合は、結果はnilです。
func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)
