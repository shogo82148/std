// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// これらの例は、flagパッケージのより複雑な使用方法を示しています。
package flag_test

func Example() {

	// ここには興味深いコードの変数が宣言されますが、
	// それらのフラグをフラグパッケージが見るためには、通常はmainの開始時に実行する必要があります（initではなく）：
	// 	flag.Parse()
	// このコードはパッケージのテストスイートの一部である「Example」と呼ばれる関数であり、
	// すでにフラグが解析されていますので、ここでは呼び出しません。
	// ただし、pkg.go.devで表示される場合、この関数は「main」という名前に変更され、スタンドアロンの例として実行できます。
}
