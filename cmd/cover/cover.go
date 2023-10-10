// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// Blockは解析で記録する基本ブロックに関する情報を表します。
// 注：基本ブロックの定義は制御構造に基づいています。&&や||は分割しません。
// 分割することはできますが、重要ではないと思われますので、手間をかける価値はありません。
type Block struct {
	startByte token.Pos
	endByte   token.Pos
	numStmt   int
}

// パッケージはパッケージ固有の状態を保持します。
type Package struct {
	mdb            *encodemeta.CoverageMetaDataBuilder
	counterLengths []int
}

// 関数は関数特有の状態を持ちます。
type Func struct {
	units      []coverage.CoverableUnit
	counterVar string
}

// Fileはパーサーで使用されるファイルの状態のラッパーです。
// 基本的なパースツリーウォーカーは、このタイプのメソッドです。
type File struct {
	fset    *token.FileSet
	name    string
	astFile *ast.File
	blocks  []Block
	content []byte
	edit    *edit.Buffer
	mdb     *encodemeta.CoverageMetaDataBuilder
	fn      Func
	pkg     *Package
}

// Visitはast.Visitorインターフェースを実装します。
func (f *File) Visit(node ast.Node) ast.Visitor
