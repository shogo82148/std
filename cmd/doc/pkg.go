// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/go/doc"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

type Package struct {
	writer      io.Writer
	name        string
	userPath    string
	pkg         *ast.Package
	file        *ast.File
	doc         *doc.Package
	build       *build.Package
	typedValue  map[*doc.Value]bool
	constructor map[*doc.Func]bool
	fs          *token.FileSet
	buf         pkgBuffer
}

func (p *Package) ToText(w io.Writer, text, prefix, codePrefix string)

// pkgBuffer is a wrapper for bytes.Buffer that prints a package clause the
// first time Write is called.

type PackageError string

func (p PackageError) Error() string

// pkg.Fatalf is like log.Fatalf, but panics so it can be recovered in the
// main do function, so it doesn't cause an exit. Allows testing to work
// without running a subprocess. The log prefix will be added when
// logged in main; it is not added here.
func (pkg *Package) Fatalf(format string, args ...any)

func (pkg *Package) Printf(format string, args ...any)
