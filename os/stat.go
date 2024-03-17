// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

<<<<<<< HEAD
// Statは指定されたファイルに関するFileInfoを返します。
// エラーが発生した場合、*PathErrorの型です。
func Stat(name string) (FileInfo, error)

// Lstatは、名前付きファイルに関するFileInfoを返します。
// ファイルがシンボリックリンクである場合、返されるFileInfoはシンボリックリンクを説明します。
// Lstatは、リンクをたどる試みをしません。
// エラーがある場合、*PathError型になります。
=======
// Stat returns a [FileInfo] describing the named file.
// If there is an error, it will be of type [*PathError].
func Stat(name string) (FileInfo, error)

// Lstat returns a [FileInfo] describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link. Lstat makes no attempt to follow the link.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/master
//
// Windowsでは、ファイルが他の名前付きエンティティ（シンボリックリンクやマウントされたフォルダなど）の代替となるリパースポイントである場合、返されるFileInfoはリパースポイントを説明し、解決しようとしません。
func Lstat(name string) (FileInfo, error)
