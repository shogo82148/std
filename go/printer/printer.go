// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package printer implements printing of AST nodes.
package printer

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

// A Mode value is a set of flags (or 0). They control printing.
type Mode uint

const (
	RawFormat Mode = 1 << iota
	TabIndent
	UseSpaces
	SourcePos
)

// A Config node controls the output of Fprint.
type Config struct {
	Mode     Mode
	Tabwidth int
	Indent   int
}

// A CommentedNode bundles an AST node and corresponding comments.
// It may be provided as argument to any of the Fprint functions.
type CommentedNode struct {
	Node     any
	Comments []*ast.CommentGroup
}

// Fprint "pretty-prints" an AST node to output for a given configuration cfg.
// Position information is interpreted relative to the file set fset.
// The node type must be *ast.File, *CommentedNode, []ast.Decl, []ast.Stmt,
// or assignment-compatible to ast.Expr, ast.Decl, ast.Spec, or ast.Stmt.
func (cfg *Config) Fprint(output io.Writer, fset *token.FileSet, node any) error

// Fprint "pretty-prints" an AST node to output.
// It calls Config.Fprint with default settings.
// Note that gofmt uses tabs for indentation but spaces for alignment;
// use format.Node (package go/format) for output that matches gofmt.
func Fprint(output io.Writer, fset *token.FileSet, node any) error
