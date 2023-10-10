// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは、各関数に対して（行、列）-（行-列）の範囲を計算するビジターを実装します。

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// FuncExtentはソース内の関数の範囲をファイルと位置で説明します。
type FuncExtent struct {
	name      string
	startLine int
	startCol  int
	endLine   int
	endCol    int
}

// FuncVisitorは、ファイルの関数の位置リストを構築するための訪問者を実装します。
type FuncVisitor struct {
	fset    *token.FileSet
	name    string
	astFile *ast.File
	funcs   []*FuncExtent
}

// Visit は ast.Visitor インターフェースを実装します。
func (v *FuncVisitor) Visit(node ast.Node) ast.Visitor

// Pkgは単一のパッケージを説明します。 'go list'のJSON出力と互換性があります。 'go help list'を参照してください。
type Pkg struct {
	ImportPath string
	Dir        string
	Error      *struct {
		Err string
	}
}
