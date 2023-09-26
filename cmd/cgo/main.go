// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Cgo; see gmp.go for an overview.

// TODO(rsc):
//	Emit correct line number annotations.
//	Make gc understand the annotations.

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// A Package collects information about the package we're going to write.
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
}

// A File collects information about a single Go input file.
type File struct {
	AST      *ast.File
	Comments []*ast.CommentGroup
	Package  string
	Preamble string
	Ref      []*Ref
	Calls    []*Call
	ExpFunc  []*ExpFunc
	Name     map[string]*Name
}

// A Call refers to a call of a C.xxx function in the AST.
type Call struct {
	Call     *ast.CallExpr
	Deferred bool
}

// A Ref refers to an expression of the form C.xxx in the AST.
type Ref struct {
	Name    *Name
	Expr    *ast.Expr
	Context string
}

func (r *Ref) Pos() token.Pos

// A Name collects information about C.xxx.
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

// IsVar reports whether Kind is either "var" or "fpvar"
func (n *Name) IsVar() bool

// A ExpFunc is an exported function, callable from C.
// Such functions are identified in the Go input file
// by doc comments containing the line //export ExpName
type ExpFunc struct {
	Func    *ast.FuncDecl
	ExpName string
	Doc     string
}

// A TypeRepr contains the string representation of a type.
type TypeRepr struct {
	Repr       string
	FormatArgs []interface{}
}

// A Type collects information about a type in both the C and Go worlds.
type Type struct {
	Size       int64
	Align      int64
	C          *TypeRepr
	Go         ast.Expr
	EnumValues map[string]int64
	Typedef    string
}

// A FuncType collects information about a function type in both the C and Go worlds.
type FuncType struct {
	Params []*Type
	Result *Type
	Go     *ast.FuncType
}

// This flag is for bootstrapping a new Go implementation,
// to generate Go types that match the data layout and
// constant values used in the host's C libraries and system calls.

// Record what needs to be recorded about f.
func (p *Package) Record(f *File)
