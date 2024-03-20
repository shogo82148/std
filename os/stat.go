// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Statは指定されたファイルに関する [FileInfo] を返します。
// エラーが発生した場合、[*PathError] の型です。
func Stat(name string) (FileInfo, error)

// Lstatは、名前付きファイルに関する [FileInfo] を返します。
// ファイルがシンボリックリンクである場合、返されるFileInfoはシンボリックリンクを説明します。
// Lstatは、リンクをたどる試みをしません。
// エラーがある場合、[*PathError] 型になります。
//
// Windowsでは、ファイルが他の名前付きエンティティ（シンボリックリンクやマウントされたフォルダなど）の代替となるリパースポイントである場合、返されるFileInfoはリパースポイントを説明し、解決しようとしません。
func Lstat(name string) (FileInfo, error)
