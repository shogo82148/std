// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/parser"
	"github.com/shogo82148/std/go/token"
)

func ExampleParseFile() {
	fset := token.NewFileSet() // positionsはfsetに対して相対的な位置にあります。

	src := `package foo

import (
	"fmt"
	"time"
)

func bar() {
	fmt.Println(time.Now())
}`

	// インポートの処理をした後にsrcをパースしますが、それ以降の処理を停止します。
	f, err := parser.ParseFile(fset, "", src, parser.ImportsOnly)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ファイルのASTからインポートを出力する。
	for _, s := range f.Imports {
		fmt.Println(s.Path.Value)
	}

	// 出力：
	//
	// "fmt"
	// "time"
}
