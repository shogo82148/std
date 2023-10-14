// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// FormatFileInfoは、人間が読みやすい形式に整形されたinfoのバージョンを返します。
// FileInfoの実装は、Stringメソッドからこれを呼び出すことができます。
// "hello.go"という名前のファイル、100バイト、モード0o644、作成日時は1970年1月1日の正午の場合、出力は次のようになります。
//
// -rw-r--r-- 100 1970-01-01 12:00:00 hello.go
func FormatFileInfo(info FileInfo) string

// FormatDirEntry は dir の人間が読みやすい形式のフォーマット済みバージョンを返します。
// [DirEntry] の実装は、これを String メソッドから呼び出すことができます。
// 名前が subdir のディレクトリと名前が hello.go のファイルの出力は次の通りです：
//
// d subdir/
// - hello.go
func FormatDirEntry(dir DirEntry) string
