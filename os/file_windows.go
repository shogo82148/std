// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

<<<<<<< HEAD
// Fdは、開いているファイルを参照するWindowsハンドルを返します。
// もしfが閉じている場合、ファイルディスクリプタは無効になります。
// もしfがガベージコレクションされる場合、ファイナライザーがファイルディスクリプタを閉じる可能性があり、
// それにより無効になります。ファイナライザーがいつ実行されるかの詳細については、
// [runtime.SetFinalizer] を参照してください。Unixシステムでは、これにより [File.SetDeadline]
// メソッドが動作しなくなります。
func (file *File) Fd() uintptr

// NewFileは指定したファイルディスクリプタと名前の新しいFileを返します。
// fdが有効なファイルディスクリプタでない場合、返される値はnilになります。
func NewFile(fd uintptr, name string) *File

// DevNullはオペレーティングシステムの「nullデバイス」の名前です。
// Unix系のシステムでは、"/dev/null"です。Windowsでは"NUL"です。
=======
// DevNull is the name of the operating system's “null device.”
// On Unix-like systems, it is "/dev/null"; on Windows, "NUL".
>>>>>>> upstream/release-branch.go1.25
const DevNull = "NUL"

// Truncateは指定されたファイルのサイズを変更します。
// もしファイルがシンボリックリンクである場合、リンクの対象のサイズを変更します。
func Truncate(name string, size int64) error

<<<<<<< HEAD
// Removeは指定されたファイルまたはディレクトリを削除します。
// エラーが発生した場合、*PathErrorの型で返されます。
=======
// Remove removes the named file or directory.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
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
