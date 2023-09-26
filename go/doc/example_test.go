// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc_test

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/token"
)

// This example illustrates how to use NewFromFiles
// to compute package documentation with examples.
func ExampleNewFromFiles() {
	// src and test are two source files that make up
	// a package whose documentation will be computed.
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

	// Create the AST by parsing src and test.
	fset := token.NewFileSet()
	files := []*ast.File{
		mustParse(fset, "src.go", src),
		mustParse(fset, "src_test.go", test),
	}

	// Compute package documentation with examples.
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

// Parse literal source code as a *doc.Package.
