// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package printer_test

import (
	"bytes"
	"fmt"
	"go/printer"
	"strings"
)

func ExampleFprint() {
	// Parse source file and extract the AST without comments for
	// this function, with position information referring to the
	// file set fset.
	funcAST, fset := parseFunc("example_test.go", "ExampleFprint")

	// Print the function body into buffer buf.
	// The file set is provided to the printer so that it knows
	// about the original source formatting and can add additional
	// line breaks where they were present in the source.
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, funcAST.Body)

	// Remove braces {} enclosing the function body, unindent,
	// and trim leading and trailing white space.
	s := buf.String()
	s = s[1 : len(s)-1]
	s = strings.TrimSpace(strings.ReplaceAll(s, "\n\t", "\n"))

	// Print the cleaned-up body text to stdout.
	fmt.Println(s)

	// output:
	// funcAST, fset := parseFunc("example_test.go", "ExampleFprint")
	//
	// var buf bytes.Buffer
	// printer.Fprint(&buf, fset, funcAST.Body)
	//
	// s := buf.String()
	// s = s[1 : len(s)-1]
	// s = strings.TrimSpace(strings.ReplaceAll(s, "\n\t", "\n"))
	//
	// fmt.Println(s)
}
