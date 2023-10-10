// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/parser"
	"github.com/shogo82148/std/go/token"
)

func Example_retrievePositionInfo() {
	fset := token.NewFileSet()

	const src = `package main

import "fmt"

import "go/token"

//line :1:5
type p = token.Pos

const bad = token.NoPos

//line fake.go:42:11
func ok(pos p) bool {
	return pos != bad
}

/*line :7:9*/func main() {
	fmt.Println(ok(bad) == bad.IsValid())
}
`

	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// f内の各宣言の場所と種類を表示する。
	for _, decl := range f.Decls {

		// ファイルセットを通じてファイル名、行番号、列番号を取得します。
		// 相対位置と絶対位置の両方を取得します。
		// 相対位置は、直前の行ディレクティブに対する相対的な位置です。
		// 絶対位置はソース内の正確な位置です。
		pos := decl.Pos()
		relPosition := fset.Position(pos)
		absPosition := fset.PositionFor(pos, false)

		// エラーが発生した場合は、FuncDeclまたはGenDeclのいずれかであるため、終了します。
		kind := "func"
		if gen, ok := decl.(*ast.GenDecl); ok {
			kind = gen.Tok.String()
		}

		// もし相対位置と絶対位置が異なる場合は、両方を表示する。
		fmtPosition := relPosition.String()
		if relPosition != absPosition {
			fmtPosition += "[" + absPosition.String() + "]"
		}

		fmt.Printf("%s: %s\n", fmtPosition, kind)
	}

	//Output:
	//
	// main.go:3:1: import
	// main.go:5:1: import
	// main.go:1:5[main.go:8:1]: type
	// main.go:3:1[main.go:10:1]: const
	// fake.go:42:11[main.go:13:1]: func
	// fake.go:7:9[main.go:17:14]: func
}
