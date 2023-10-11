// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Statは指定されたファイルに関するFileInfoを返します。
// エラーが発生した場合、*PathErrorの型です。
func Stat(name string) (FileInfo, error)

// Lstatは指定したファイルに関する情報を返す。
// ファイルがシンボリックリンクの場合、返されるFileInfoは
// シンボリックリンクに関する情報を記述する。
// Lstatはリンクを辿る試みを行わない。
// エラーが発生した場合、そのエラーは*PathError型になる。
func Lstat(name string) (FileInfo, error)
