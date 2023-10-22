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

<<<<<<< HEAD
// CommentedNodeは、ASTノードと対応するコメントをまとめたものです。
// これは、Fprint関数の引数として提供することができます。
=======
// A CommentedNode bundles an AST node and corresponding comments.
// It may be provided as argument to any of the [Fprint] functions.
>>>>>>> upstream/master
type CommentedNode struct {
	Node     any
	Comments []*ast.CommentGroup
}

<<<<<<< HEAD
// Fprintは与えられた設定cfgに対して、ASTノードを出力に「きれいに表示」します。
// 位置情報はファイルセットfsetを基準に解釈されます。
// ノードの型は *ast.File、*CommentedNode、[]ast.Decl、[]ast.Stmt、または ast.Expr、ast.Decl、ast.Spec、ast.Stmtに互換性のあるものである必要があります。
func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node any) error

// FprintはASTノードを出力に「整形表示」します。
// それはデフォルトの設定でConfig.Fprintを呼び出します。
// gofmt はインデントにタブを使用し、整列にはスペースを使用することに注意してください。
// gofmtと一致する出力にはformat.Node（パッケージgo/format）を使用してください。
=======
// Fprint "pretty-prints" an AST node to output for a given configuration cfg.
// Position information is interpreted relative to the file set fset.
// The node type must be *[ast.File], *[CommentedNode], [][ast.Decl], [][ast.Stmt],
// or assignment-compatible to [ast.Expr], [ast.Decl], [ast.Spec], or [ast.Stmt].
func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node any) error

// Fprint "pretty-prints" an AST node to output.
// It calls [Config.Fprint] with default settings.
// Note that gofmt uses tabs for indentation but spaces for alignment;
// use format.Node (package go/format) for output that matches gofmt.
>>>>>>> upstream/master
func Fprint(output io.Writer, fset *token.FileSet, node any) error
