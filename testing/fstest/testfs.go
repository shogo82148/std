// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// fstestパッケージは、ファイルシステムの実装およびユーザーのテストをサポートする機能を実装します。
package fstest

import (
	"github.com/shogo82148/std/io/fs"
)

// TestFSはファイルシステムの実装をテストします。
// fsys内のファイルツリー全体を走査し、
// 各ファイルを開いて正しく動作するかを確認します。
// シンボリックリンクは辿りませんが、
// ファイルシステムが [fs.ReadLinkFS] を実装している場合は、そのLstat値を確認します。
// また、ファイルシステムが少なくとも期待されるファイルを含んでいるかもチェックします。
// 特別なケースとして、期待されるファイルが指定されていない場合、fsysは空でなければなりません。
// それ以外の場合、fsysは少なくとも指定されたファイルを含み、他のファイルを含んでもかまいません。
// fsysの内容はTestFSの実行中に同時に変更されてはなりません。
//
// TestFSが何かしらの不適切な振る舞いを見つけた場合、最初のエラーまたはエラーのリストを返します。
// 検査には [errors.Is] または [errors.As] を使用します。
//
// テスト内での典型的な使用法は以下の通りです：
//
//	if err := fstest.TestFS(myFS, "file/that/should/be/present"); err != nil {
//		t.Fatal(err)
//	}
func TestFS(fsys fs.FS, expected ...string) error
