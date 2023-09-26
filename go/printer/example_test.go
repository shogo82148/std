// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package printer_test

func ExampleFprint() {
	printSelf()

	// Output:
	// funcAST, fset := parseFunc("example_test.go", "printSelf")
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
