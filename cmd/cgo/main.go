// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Cgo; 概要についてはdoc.goを参照してください。

// TODO（rsc）：
// 正しい行番号のアノテーションを出力する。
// gcがアノテーションを理解するようにする。

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// パッケージは書く予定のパッケージに関する情報を集めます。
type Package struct {
	PackageName string
	PackagePath string
	PtrSize     int64
	IntSize     int64
	GccOptions  []string
	GccIsClang  bool
	CgoFlags    map[string][]string
	Written     map[string]bool
	Name        map[string]*Name
	ExpFunc     []*ExpFunc
	Decl        []ast.Decl
	GoFiles     []string
	GccFiles    []string
	Preamble    string
	typedefs    map[string]bool
	typedefList []typedefInfo
}

type File struct {
	AST      *ast.File
	Comments []*ast.CommentGroup
	Package  string
	Preamble string
	Ref      []*Ref
	Calls    []*Call
	ExpFunc  []*ExpFunc
	Name     map[string]*Name
	NamePos  map[*Name]token.Pos
	Edit     *edit.Buffer
}

// Callは、ASTのC.xxx関数の呼び出しを参照します。
type Call struct {
	Call     *ast.CallExpr
	Deferred bool
	Done     bool
}

// Ref型はAST内のC.xxxの形式の式を参照します。
type Ref struct {
	Name    *Name
	Expr    *ast.Expr
	Context astContext
	Done    bool
}

func (r *Ref) Pos() token.Pos

// NameはC.xxxに関する情報を収集します。
type Name struct {
	Go       string
	Mangle   string
	C        string
	Define   string
	Kind     string
	Type     *Type
	FuncType *FuncType
	AddError bool
	Const    string
}

// IsVarは、Kindが "var" または "fpvar" であるかどうかを報告します。
func (n *Name) IsVar() bool

// IsConstはKindが「iconst」、「fconst」または「sconst」であるかどうか報告します。
func (n *Name) IsConst() bool

// ExpFuncはCから呼び出すことができるエクスポートされた関数です。
// このような関数は、Goの入力ファイルに含まれるドキュメントコメントの行//export ExpNameによって識別されます。
type ExpFunc struct {
	Func    *ast.FuncDecl
	ExpName string
	Doc     string
}

// TypeReprは型の文字列表現を含む
type TypeRepr struct {
	Repr       string
	FormatArgs []interface{}
}

// TypeはCとGoの世界の両方でタイプに関する情報を収集します。
type Type struct {
	Size       int64
	Align      int64
	C          *TypeRepr
	Go         ast.Expr
	EnumValues map[string]int64
	Typedef    string
	BadPointer bool
}

// FuncTypeはCとGoの両方の世界における関数型に関する情報を収集します。
type FuncType struct {
	Params []*Type
	Result *Type
	Go     *ast.FuncType
}

// fについて記録すべきことを記録する
func (p *Package) Record(f *File)
