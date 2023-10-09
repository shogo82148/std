// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 入力ASTを解析し、Prog構造体を準備します。

package main

// ParseGoは、指定されたファイルから読み込まれたGoソースコードから取得した情報でfを埋めます。
// これには、import "C"コメントに付属しているCの前文、C.xxxへの参照のリスト、
// エクスポートされた関数のリスト、および実際のASTが含まれ、書き直されて出力されます。
func (f *File) ParseGo(abspath string, src []byte)
