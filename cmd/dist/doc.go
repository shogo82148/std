// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DistはGoの配布をブートストラップ、ビルド、テストするのに役立ちます。
//
// 使用方法：
//
//	go tool dist [コマンド]
//
// コマンドは以下の通りです：
//
//	banner         インストールバナーを表示する
//	bootstrap      すべてを再ビルドする
//	clean          すべてのビルドファイルを削除する
//	env [-p]       環境を表示する（-p：$PATHを含む）
//	install [dir]  個々のディレクトリをインストールする
//	list [-json]   サポートされているすべてのプラットフォームをリストする
//	test [-h]      Goテストを実行する
//	version        Goのバージョンを表示する
package main
