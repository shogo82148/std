// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Packは、伝統的なUnixのarツールの簡易版です。
Go言語で必要な操作のみを実装しています。

使用方法:

	go tool pack op file.a [name...]

Packは、アーカイブに対して操作を適用し、操作の引数として名前を使用します。

opは、次のいずれかの文字で指定される操作です:

	c	新しいアーカイブにファイル（ファイルシステムから）を追加する
	p	アーカイブからファイルを表示する
	r	アーカイブにファイル（ファイルシステムから）を追加する
	t	アーカイブからファイルを一覧表示する
	x	アーカイブからファイルを抽出する

cコマンドへのアーカイブ引数は、存在しないか有効なアーカイブファイルでなければならず、
新しいエントリを追加する前にクリアされます。ファイルが存在するがアーカイブではない場合はエラーです。

p、t、xコマンドでは、コマンドラインの名前がない場合、操作はアーカイブ内のすべてのファイルに適用されます。

Unixのarとは異なり、r操作は常にアーカイブに追記されます。
つまり、指定した名前のファイルがアーカイブに既に存在していても、追加されます。
このように、packのr操作はUnixのarのrq操作に近い動作です。

操作の末尾にv文字を追加する（pvまたはrvなど）と、冗長な操作が有効になります:
cおよびrコマンドでは、ファイルが追加されるたびに名前が表示されます。
pコマンドでは、各ファイルが名前で前置された行で表示されます。
tコマンドでは、一覧には追加のファイルメタデータが含まれます。
xコマンドでは、ファイルが抽出されるたびに名前が表示されます。
*/package main
