// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// テストカバレッジのサポート。

package testing

// CoverBlock は単一の基本ブロックのカバレッジデータを記録します。
// フィールドは1からインデックス付けされています。エディタのように、
// ファイルの開始行は1です。例えば、列はバイト単位で測定されます。
// 注: この struct はテストインフラストラクチャ内部で使用されるものであり、変更される可能性があります。
// まだ Go 1 互換性ガイドラインにカバーされていません。
type CoverBlock struct {
	Line0 uint32
	Col0  uint16
	Line1 uint32
	Col1  uint16
	Stmts uint16
}

// Coverはテストカバレッジのチェックに関する情報を記録します。
// 注意: この構造体はテストインフラストラクチャに内部的に使用され、変更される可能性があります。
// Go 1の互換性ガイドラインではまだ対象外です。
type Cover struct {
	Mode            string
	Counters        map[string][]uint32
	Blocks          map[string][]CoverBlock
	CoveredPackages string
}

// RegisterCoverはテストのためのカバレッジデータアキュムレータを記録します。
// 注意: この関数はテストインフラストラクチャ内部のものであり、変更される可能性があります。
// まだGo 1互換性ガイドラインには含まれていません。
func RegisterCover(c Cover)
