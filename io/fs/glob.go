// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// GlobFSはGlobメソッドを持つファイルシステムです。
type GlobFS interface {
	FS

	Glob(pattern string) ([]string, error)
}

<<<<<<< HEAD
// Glob returns the names of all files matching pattern or nil
// if there is no matching file. The syntax of patterns is the same
// as in [path.Match]. The pattern may describe hierarchical names such as
// usr/*/bin/ed.
//
// Glob ignores file system errors such as I/O errors reading directories.
// The only possible returned error is [path.ErrBadPattern], reporting that
// the pattern is malformed.
//
// If fs implements [GlobFS], Glob calls fs.Glob.
// Otherwise, Glob uses [ReadDir] to traverse the directory tree
// and look for matches for the pattern.
=======
// Globは、パターンに一致するすべてのファイルの名前を返します。一致するファイルがない場合はnilを返します。
// パターンの構文はpath.Matchと同じです。パターンはusr/*/bin/edのような階層的な名前を指定することができます。
//
// Globは、ディレクトリの読み取り時のI/Oエラーなどのファイルシステムのエラーを無視します。
// 返される唯一の可能なエラーは、path.ErrBadPatternで、パターンが不正であることを報告します。
//
// もしfsがGlobFSを実装している場合、Globはfs.Globを呼び出します。
// そうでない場合、Globはディレクトリツリーをトラバースし、パターンに一致するものを探します。
>>>>>>> release-branch.go1.21
func Glob(fsys FS, pattern string) (matches []string, err error)
