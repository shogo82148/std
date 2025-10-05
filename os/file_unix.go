// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package os

// DevNullはオペレーティングシステムの「nullデバイス」の名称です。
// Unix系のシステムでは"/dev/null"、Windowsでは"NUL"です。
const DevNull = "/dev/null"

// Truncateは指定されたファイルのサイズを変更します。
// ファイルがシンボリックリンクの場合は、リンク先のサイズを変更します。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func Truncate(name string, size int64) error

// Removeは指定されたファイルまたは（空の）ディレクトリを削除します。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func Remove(name string) error

// Link は newname を oldname ファイルのハードリンクとして作成します。
// エラーがある場合、*LinkError 型になります。
func Link(oldname, newname string) error

// Symlinkはnewnameをoldnameへのシンボリックリンクとして作成します。
// Windowsでは、存在しないoldnameへのシンボリックリンクはファイルのシンボリックリンクとして作成されます。
// もしoldnameが後でディレクトリとして作成された場合、シンボリックリンクは機能しません。
// エラーが発生した場合は、*LinkError型のエラーになります。
func Symlink(oldname, newname string) error
