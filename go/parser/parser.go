// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files.
//
// The [ParseFile] function reads file input from a string, []byte, or
// io.Reader, and produces an [ast.File] representing the complete
// abstract syntax tree of the file.
//
// The [ParseExprFrom] function reads a single source-level expression and
// produces an [ast.Expr], the syntax tree of the expression.
//
// The parser accepts a larger language than is syntactically permitted by
// the Go spec, for simplicity, and for improved robustness in the presence
// of syntax errors. For instance, in method declarations, the receiver is
// treated like an ordinary parameter list and thus may contain multiple
// entries where the spec permits exactly one. Consequently, the corresponding
// field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.
//
// Applications that need to parse one or more complete packages of Go
// source code may find it more convenient not to interact directly
// with the parser but instead to use the Load function in package
// [golang.org/x/tools/go/packages].
package parser
