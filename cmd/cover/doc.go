// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Coverは、'go test -coverprofile=cover.out'で生成されるカバレッジプロファイルを解析するためのプログラムです。

Cover is also used by 'go test -cover' to rewrite the source code with
annotations to track which parts of each function are executed (this
is referred to "instrumentation"). Cover can operate in "legacy mode"
on a single Go source file at a time, or when invoked by the Go tool
it will process all the source files in a single package at a time
(package-scope instrumentation is enabled via "-pkgcfg" option).

インストゥルメンテーションされたコードを生成する際、カバレッジツールはソースコードを調査しておおよその基本ブロック情報を計算します。したがって、バイナリ書き換えカバレッジツールよりも移植性は高くなりますが、少し機能は制限されます。たとえば、&&および||式の内部にはプローブを挿入せず、単一の文に複数の関数リテラルがある場合には僅かに混乱する可能性があります。

cgoを使用するパッケージのカバレッジを計算する場合、カバレッジツールは入力ではなく、cgoの前処理の出力に適用する必要があります。なぜなら、カバレッジツールはcgoにとって重要なコメントを削除するからです。

使用方法については、次を参照してください：

	go help testflag
	go tool cover -help
*/package main
