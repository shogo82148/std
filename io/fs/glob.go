// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// GlobFSはGlobメソッドを持つファイルシステムです。
type GlobFS interface {
	FS

	Glob(pattern string) ([]string, error)
}

// Globは、パターンに一致するすべてのファイルの名前を返します。一致するファイルがない場合はnilを返します。
// パターンの構文は [path.Match] と同じです。パターンはusr/*/bin/edのような階層的な名前を指定することができます。
//
// Globは、ディレクトリの読み取り時のI/Oエラーなどのファイルシステムのエラーを無視します。
// 返される唯一の可能なエラーは、 [path.ErrBadPattern] で、パターンが不正であることを報告します。
//
// もしfsが [GlobFS] を実装している場合、Globはfs.Globを呼び出します。
// そうでない場合、Globは [ReadDir] を使用してディレクトリツリーをトラバースし、パターンに一致するものを探します。
func Glob(fsys FS, pattern string) (matches []string, err error)
