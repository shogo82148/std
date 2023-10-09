// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/doc"
	"github.com/shogo82148/std/go/token"
)

// この例は、NewFromFilesを使用してパッケージのドキュメントと例を計算する方法を示しています。
func ExampleNewFromFiles() {

	// srcとtestは、ドキュメントが計算されるパッケージを構成する2つのソースファイルです。
	const src = `
// This is the package comment.
package p

import "fmt"

// This comment is associated with the Greet function.
func Greet(who string) {
	fmt.Printf("Hello, %s!\n", who)
}
`
	const test = `
package p_test

// This comment is associated with the ExampleGreet_world example.
func ExampleGreet_world() {
	Greet("world")
}
`

	// srcとtestを解析してASTを作成します。
	fset := token.NewFileSet()
	files := []*ast.File{
		mustParse(fset, "src.go", src),
		mustParse(fset, "src_test.go", test),
	}

	// 例を用いて計算パッケージのドキュメントを作成します。
	p, err := doc.NewFromFiles(fset, files, "example.com/p")
	if err != nil {
		panic(err)
	}

	fmt.Printf("package %s - %s", p.Name, p.Doc)
	fmt.Printf("func %s - %s", p.Funcs[0].Name, p.Funcs[0].Doc)
	fmt.Printf(" ⤷ example with suffix %q - %s", p.Funcs[0].Examples[0].Suffix, p.Funcs[0].Examples[0].Doc)

	// Output:
	// package p - This is the package comment.
	// func Greet - This comment is associated with the Greet function.
	//  ⤷ example with suffix "world" - This comment is associated with the ExampleGreet_world example.
}
