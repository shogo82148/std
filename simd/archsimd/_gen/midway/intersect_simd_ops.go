// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/ast"
)

type MethodSet map[string]*ast.FuncDecl
type TypeMethods map[string]MethodSet

type Comments struct {
	Types     map[string]string            `yaml:"types"`
	Functions map[string]string            `yaml:"functions"`
	Methods   map[string]map[string]string `yaml:"methods"`
}

type ArchAndFiles struct {
	arch  string
	files []string
}

type TypeMethod struct {
	t, m string
}
