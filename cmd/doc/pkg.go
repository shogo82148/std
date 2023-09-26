// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/go/doc"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
)

type Package struct {
	writer   io.Writer
	name     string
	userPath string
	pkg      *ast.Package
	file     *ast.File
	doc      *doc.Package
	build    *build.Package
	fs       *token.FileSet
	buf      bytes.Buffer
}

type PackageError string

func (p PackageError) Error() string

// pkg.Fatalf is like log.Fatalf, but panics so it can be recovered in the
// main do function, so it doesn't cause an exit. Allows testing to work
// without running a subprocess. The log prefix will be added when
// logged in main; it is not added here.
func (pkg *Package) Fatalf(format string, args ...interface{})

func (pkg *Package) Printf(format string, args ...interface{})
