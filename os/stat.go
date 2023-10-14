// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Statは指定されたファイルに関するFileInfoを返します。
// エラーが発生した場合、*PathErrorの型です。
func Stat(name string) (FileInfo, error)

<<<<<<< HEAD
// Lstat returns a FileInfo describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link. Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
//
// On Windows, if the file is a reparse point that is a surrogate for another
// named entity (such as a symbolic link or mounted folder), the returned
// FileInfo describes the reparse point, and makes no attempt to resolve it.
=======
// Lstatは指定したファイルに関する情報を返す。
// ファイルがシンボリックリンクの場合、返されるFileInfoは
// シンボリックリンクに関する情報を記述する。
// Lstatはリンクを辿る試みを行わない。
// エラーが発生した場合、そのエラーは*PathError型になる。
>>>>>>> release-branch.go1.21
func Lstat(name string) (FileInfo, error)
