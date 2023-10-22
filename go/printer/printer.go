// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package printerは、ASTノードの印刷を実装しています。
package printer

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

// モード値はフラグの集合（または0）です。これらは印刷を制御します。
type Mode uint

const (
	RawFormat Mode = 1 << iota
	TabIndent
	UseSpaces
	SourcePos
)

// ConfigノードはFprintの出力を制御します。
type Config struct {
	Mode     Mode
	Tabwidth int
	Indent   int
}

// CommentedNodeは、ASTノードと対応するコメントをまとめたものです。
// これは、[Fprint] 関数の引数として提供することができます。
type CommentedNode struct {
	Node     any
	Comments []*ast.CommentGroup
}

// Fprintは与えられた設定cfgに対して、ASTノードを出力に「きれいに表示」します。
// 位置情報はファイルセットfsetを基準に解釈されます。
// ノードの型は *[ast.File]、*[CommentedNode]、[][ast.Decl]、[][ast.Stmt]、または [ast.Expr]、[ast.Decl]、[ast.Spec]、[ast.Stmt] に互換性のあるものである必要があります。
func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node any) error

// FprintはASTノードを出力に「整形表示」します。
// それはデフォルトの設定でConfig.Fprintを呼び出します。
// gofmt はインデントにタブを使用し、整列にはスペースを使用することに注意してください。
// gofmtと一致する出力にはformat.Node（パッケージgo/format）を使用してください。
func Fprint(output io.Writer, fset *token.FileSet, node any) error
