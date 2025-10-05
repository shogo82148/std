// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// DevNullはオペレーティングシステムの「nullデバイス」の名称です。
// Unix系のシステムでは"/dev/null"、Windowsでは"NUL"です。
const DevNull = "NUL"

// Truncateは指定されたファイルのサイズを変更します。
// もしファイルがシンボリックリンクである場合、リンクの対象のサイズを変更します。
func Truncate(name string, size int64) error

// Removeは指定されたファイルまたはディレクトリを削除します。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func Remove(name string) error

// Pipeは接続された一対のファイルを返します。rからの読み取りはwに書き込まれたバイトを返します。
// エラーがある場合は、ファイルとエラーを返します。返されたファイルのWindowsハンドルは、子プロセスに引き継がれるようにマークされています。
func Pipe() (r *File, w *File, err error)

// Linkはoldnameファイルへのハードリンクとしてnewnameを作成します。
// エラーが発生した場合、*LinkError型になります。
func Link(oldname, newname string) error

// Symlinkはnewnameをoldnameへのシンボリックリンクとして作成します。
// Windowsでは、存在しないoldnameへのシンボリックリンクはファイルシンボリックリンクとして作成されます。
// oldnameが後でディレクトリとして作成された場合、シンボリックリンクは機能しません。
// エラーが発生した場合、*LinkErrorの型になります。
func Symlink(oldname, newname string) error
