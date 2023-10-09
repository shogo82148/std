// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/format"
	"github.com/shogo82148/std/go/parser"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/strings"
)

// この例は、GoプログラムのASTを検査する方法を示しています。
func ExampleInspect() {
	// srcはASTを検査したい入力です。
	src := `
package p
const c = 1.0
var X = f(3.14)*2 + c
`

	// srcを解析してASTを作成する。
	fset := token.NewFileSet() // ポジションはfsetに対して相対的です。
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	// AST を調査し、すべての識別子とリテラルを表示します。
	ast.Inspect(f, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		}
		if s != "" {
			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		}
		return true
	})

	// Output:
	// src.go:2:9:	p
	// src.go:3:7:	c
	// src.go:3:11:	1.0
	// src.go:4:5:	X
	// src.go:4:9:	f
	// src.go:4:11:	3.14
	// src.go:4:17:	2
	// src.go:4:21:	c
}

// この例では、デバッグ用に出力されるASTの形状を示しています。
func ExamplePrint() {
	// srcはASTを出力したい入力です。
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// src を解析してASTを作成します。
	fset := token.NewFileSet() // ポジションはfsetに対して相対的です。
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// ASTを出力する。
	ast.Print(fset, f)

	// Output:
	//      0  *ast.File {
	//      1  .  Package: 2:1
	//      2  .  Name: *ast.Ident {
	//      3  .  .  NamePos: 2:9
	//      4  .  .  Name: "main"
	//      5  .  }
	//      6  .  Decls: []ast.Decl (len = 1) {
	//      7  .  .  0: *ast.FuncDecl {
	//      8  .  .  .  Name: *ast.Ident {
	//      9  .  .  .  .  NamePos: 3:6
	//     10  .  .  .  .  Name: "main"
	//     11  .  .  .  .  Obj: *ast.Object {
	//     12  .  .  .  .  .  Kind: func
	//     13  .  .  .  .  .  Name: "main"
	//     14  .  .  .  .  .  Decl: *(obj @ 7)
	//     15  .  .  .  .  }
	//     16  .  .  .  }
	//     17  .  .  .  Type: *ast.FuncType {
	//     18  .  .  .  .  Func: 3:1
	//     19  .  .  .  .  Params: *ast.FieldList {
	//     20  .  .  .  .  .  Opening: 3:10
	//     21  .  .  .  .  .  Closing: 3:11
	//     22  .  .  .  .  }
	//     23  .  .  .  }
	//     24  .  .  .  Body: *ast.BlockStmt {
	//     25  .  .  .  .  Lbrace: 3:13
	//     26  .  .  .  .  List: []ast.Stmt (len = 1) {
	//     27  .  .  .  .  .  0: *ast.ExprStmt {
	//     28  .  .  .  .  .  .  X: *ast.CallExpr {
	//     29  .  .  .  .  .  .  .  Fun: *ast.Ident {
	//     30  .  .  .  .  .  .  .  .  NamePos: 4:2
	//     31  .  .  .  .  .  .  .  .  Name: "println"
	//     32  .  .  .  .  .  .  .  }
	//     33  .  .  .  .  .  .  .  Lparen: 4:9
	//     34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
	//     35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
	//     36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10
	//     37  .  .  .  .  .  .  .  .  .  Kind: STRING
	//     38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
	//     39  .  .  .  .  .  .  .  .  }
	//     40  .  .  .  .  .  .  .  }
	//     41  .  .  .  .  .  .  .  Ellipsis: -
	//     42  .  .  .  .  .  .  .  Rparen: 4:25
	//     43  .  .  .  .  .  .  }
	//     44  .  .  .  .  .  }
	//     45  .  .  .  .  }
	//     46  .  .  .  .  Rbrace: 5:1
	//     47  .  .  .  }
	//     48  .  .  }
	//     49  .  }
	//     50  .  FileStart: 1:1
	//     51  .  FileEnd: 5:3
	//     52  .  Scope: *ast.Scope {
	//     53  .  .  Objects: map[string]*ast.Object (len = 1) {
	//     54  .  .  .  "main": *(obj @ 11)
	//     55  .  .  }
	//     56  .  }
	//     57  .  Unresolved: []*ast.Ident (len = 1) {
	//     58  .  .  0: *(obj @ 29)
	//     59  .  }
	//     60  .  GoVersion: ""
	//     61  }
}

// この例は、ast.CommentMapを使用して、Goプログラムの変数宣言を削除しながら正しいコメントの関連を保持する方法を示しています。
func ExampleCommentMap() {

	// src は、私たちが操作するためのASTを作成する入力です。
	src := `
// This is the package comment.
package main

// This comment is associated with the hello constant.
const hello = "Hello, World!" // line comment 1

// This comment is associated with the foo variable.
var foo = hello // line comment 2

// This comment is associated with the main function.
func main() {
	fmt.Println(hello) // line comment 3
}
`

	// src をパースしてASTを作成する。
	fset := token.NewFileSet() // positionsはfsetに対して相対的です。
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// ast.File のコメントから ast.CommentMap を作成します。
	// これにより、コメントと AST ノードの関連付けが保持されます。
	cmap := ast.NewCommentMap(fset, f, f.Comments)

	// 最初の変数宣言を宣言リストから削除します。
	for i, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.VAR {
			copy(f.Decls[i:], f.Decls[i+1:])
			f.Decls = f.Decls[:len(f.Decls)-1]
			break
		}
	}

	// コメントマップを使用して、もはや必要でないコメント（変数宣言に関連するコメント）をフィルタリングし、新しいコメントリストを作成します。
	f.Comments = cmap.Filter(f).Comments()

	// 変更されたASTを出力します。
	var buf strings.Builder
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	fmt.Printf("%s", buf.String())

	// Output:
	// // This is the package comment.
	// package main
	//
	// // This comment is associated with the hello constant.
	// const hello = "Hello, World!" // line comment 1
	//
	// // This comment is associated with the main function.
	// func main() {
	// 	fmt.Println(hello) // line comment 3
	// }
}
