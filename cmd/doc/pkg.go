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

func (pkg *Package) ToText(w io.Writer, text, prefix, codePrefix string)

type PackageError string

func (p PackageError) Error() string

// pkg.Fatalfはlog.Fatalfと似ていますが、エラーが発生するため、メインのdo関数で回復が可能であり、プログラムが終了しないようになっています。サブプロセスを実行しないでテストを実行できるようにします。ログの接頭辞はメインで追加されますが、ここでは追加されません。
func (pkg *Package) Fatalf(format string, args ...any)

func (pkg *Package) Printf(format string, args ...any)
