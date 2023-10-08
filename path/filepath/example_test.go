// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/path/filepath"
)

func ExampleExt() {
	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))
	// Output:
	// No dots: ""
	// One dot: ".js"
	// Two dots: ".js"
}
