// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Fdは開いたファイルを参照するWindowsハンドルを返します。
// fが閉じられた場合、ファイルディスクリプタは無効になります。
// fがガベージコレクションされると、ファイナライザがファイルディスクリプタを閉じることがあります。
// ファイナライザが実行されるタイミングについては、runtime.SetFinalizerの詳細情報を参照してください。
// Unixシステムでは、これによりSetDeadlineメソッドが機能しなくなります。
func (file *File) Fd() uintptr

// NewFileは指定したファイルディスクリプタと名前の新しいFileを返します。
// fdが有効なファイルディスクリプタでない場合、返される値はnilになります。
func NewFile(fd uintptr, name string) *File

// DevNullはオペレーティングシステムの「nullデバイス」の名前です。
// Unix系のシステムでは、"/dev/null"です。Windowsでは"NUL"です。
const DevNull = "NUL"

// Truncateは指定されたファイルのサイズを変更します。
// もしファイルがシンボリックリンクである場合、リンクの対象のサイズを変更します。
func Truncate(name string, size int64) error

// Removeは指定されたファイルまたはディレクトリを削除します。
// エラーが発生した場合、*PathErrorの型で返されます。
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
<<<<<<< HEAD

// Readlinkは指定されたシンボリックリンクの宛先を返します。
// エラーが発生した場合、*PathErrorのタイプになります。
func Readlink(name string) (string, error)
=======
>>>>>>> upstream/master
