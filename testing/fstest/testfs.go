// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fstest implements support for testing implementations and users of file systems.
package fstest

import (
	"github.com/shogo82148/std/io/fs"
)

// TestFSは、ファイルシステムの実装をテストします。
// fsys内の全てのファイルツリーを走査し、
// 各ファイルが正しく動作するかを開いて確認します。
// また、ファイルシステムが少なくとも期待されるファイルを含んでいることも確認します。
// 特別なケースとして、期待されるファイルが一つもリストされていない場合、fsysは空でなければなりません。
// それ以外の場合、fsysは少なくともリストされたファイルを含んでいなければなりません。他のファイルも含むことができます。
// fsysの内容は、TestFSと同時に変更してはなりません。
//
// TestFSが何かしらの不適切な動作を見つけた場合、それら全てを報告するエラーを返します。
// エラーテキストは複数行にわたり、検出された不適切な動作ごとに1行ずつ存在します。
//
// テスト内での典型的な使用法は以下の通りです：
//
//	if err := fstest.TestFS(myFS, "file/that/should/be/present"); err != nil {
//		t.Fatal(err)
//	}
func TestFS(fsys fs.FS, expected ...string) error
