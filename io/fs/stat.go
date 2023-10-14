// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// StatFSは、Statメソッドを持つファイルシステムです。
type StatFS interface {
	FS

	Stat(name string) (FileInfo, error)
}

<<<<<<< HEAD
// Stat returns a [FileInfo] describing the named file from the file system.
//
// If fs implements [StatFS], Stat calls fs.Stat.
// Otherwise, Stat opens the [File] to stat it.
=======
// Statはファイルシステムから指定されたファイルに関するFileInfoを返します。
//
// もしfsがStatFSを実装している場合、Statはfs.Statを呼び出します。
// そうでない場合、Statはファイルを開いて統計情報を取得します。
>>>>>>> release-branch.go1.21
func Stat(fsys FS, name string) (FileInfo, error)
