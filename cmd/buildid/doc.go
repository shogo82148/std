// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Buildidは、Goパッケージまたはバイナリに格納されているビルドIDを表示または更新します。

使用法：

	go tool buildid [-w] ファイル

デフォルトでは、buildidは指定したファイルに含まれるビルドIDを表示します。
-wオプションが指定された場合、buildidはファイルに含まれるビルドIDを正確に記録するためのコンテンツのハッシュで上書きします。

このツールは、goコマンドまたは他のビルドシステムによってのみ使用することを想定しています。
*/package main
