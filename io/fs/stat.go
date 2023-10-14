// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// StatFSは、Statメソッドを持つファイルシステムです。
type StatFS interface {
	FS

	// Stat returns a FileInfo describing the file.
	// If there is an error, it should be of type *PathError.
	Stat(name string) (FileInfo, error)
}

// Statはファイルシステムから指定されたファイルに関する [FileInfo] を返します。
//
// もしfsが [StatFS] を実装している場合、Statはfs.Statを呼び出します。
// そうでない場合、Statは [File] を開いて統計情報を取得します。
func Stat(fsys FS, name string) (FileInfo, error)
