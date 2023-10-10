// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package format_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/format"
	"github.com/shogo82148/std/go/parser"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/log"
)

func ExampleNode() {
	const expr = "(6+2*3)/4"

	// parser.ParseExprは引数を解析し、対応するast.Nodeを返します。
	node, err := parser.ParseExpr(expr)
	if err != nil {
		log.Fatal(err)
	}

	// ノード用のFileSetを作成します。ノードは実際のソースファイルから
	// 来ないため、fsetは空になります。
	fset := token.NewFileSet()

	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(buf.String())

	// Output: (6 + 2*3) / 4
}
