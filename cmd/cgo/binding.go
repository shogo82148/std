// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/go/types"
	"github.com/shogo82148/std/sync"
)

// BindingVar describes a C global variable bound with //cgo:var.
type BindingVar struct {
	Name    string
	GoName  string
	CName   string
	BindC   bool
	Type    string
	Library string
	LibC    bool
	Static  bool

	// Methods of //cgo:dynload types:
	RecvType string
	SymIndex int
}

// Bindings holds the parsed declarations and linker flags collected
// from the //cgo: pragmas in a package.
type Bindings struct {
	mu              sync.Mutex
	Package         string
	LdFlags         []string
	Vars            []BindingVar
	Funcs           []BindingFunc
	UsedBuiltins    map[string]bool
	TypeGoToC       map[string]string
	TypeGoToAnyC    map[string]string
	ConstCToGo      map[string]string
	GoOnlyCNames    map[string]bool
	DynloadTypes    []string
	DynloadCounts   map[string]int
	DynloadVars     []BindingVar
	TypeMap         map[string]BindingType
	UsedFuncPtrs    map[string]bool
	UsedExportCalls map[string]bool

	typeSpecs             map[string]*ast.TypeSpec
	resolveUsedTypeInFile func(expr ast.Expr, inPtr bool, imp fileImports)
	tpkg                  *types.Package
	typeInfo              *types.Info
	constPkg              *types.Package
	constFset             *token.FileSet
	constFilePos          map[string]token.Pos
	constSources          []bindingSource
}
