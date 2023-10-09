// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/scanner"
	"github.com/shogo82148/std/go/token"
)

func ExampleScanner_Scan() {
	// src はトークン化したい入力です。
	src := []byte("cos(x) + 1i*sin(x) // Euler")

	// スキャナーを初期化する。
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positionsはfsetに対して相対的です。
	file := fset.AddFile("", fset.Base(), len(src)) // "file"という入力を登録する
	s.Init(file, src, nil /* エラーハンドラーなし */, scanner.ScanComments)

	// Scanの繰り返し呼び出しは、入力で見つかったトークンのシーケンスを返します。
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	// output:
	// 1:1	IDENT	"cos"
	// 1:4	(	""
	// 1:5	IDENT	"x"
	// 1:6	)	""
	// 1:8	+	""
	// 1:10	IMAG	"1i"
	// 1:12	*	""
	// 1:13	IDENT	"sin"
	// 1:16	(	""
	// 1:17	IDENT	"x"
	// 1:18	)	""
	// 1:20	COMMENT	"// Euler"
	// 1:28	;	"\n"
}
