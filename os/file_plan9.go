// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/time"
)

<<<<<<< HEAD
// Fdは、開かれたファイルを参照する整数型のPlan 9ファイルディスクリプタを返します。
// fがクローズされると、ファイルディスクリプタは無効になります。
// fがガベージコレクションされると、ファイルディスクリプタをクローズするファイナライザが実行される可能性があります。
// この場合、[runtime.SetFinalizer] に関する詳細は、ファイナライザがいつ実行されるかについての情報を参照してください。Unixシステムでは、これによって [File.SetDeadline] メソッドが動作しなくなります。
//
// 代替案として、f.SyscallConnメソッドを参照してください。
func (f *File) Fd() uintptr

// NewFile は指定されたファイルディスクリプタと名前で新しいファイルを返します。
// fd が有効なファイルディスクリプタでない場合、返される値は nil になります。
func NewFile(fd uintptr, name string) *File

// DevNullはオペレーティングシステムの「nullデバイス」の名称です。
// Unix系のシステムでは、"/dev/null"となります。Windowsでは、"NUL"です。
=======
// DevNull is the name of the operating system's “null device.”
// On Unix-like systems, it is "/dev/null"; on Windows, "NUL".
>>>>>>> upstream/release-branch.go1.25
const DevNull = "/dev/null"

// Closeは、Fileを閉じて、I/Oに使用できなくします。
// SetDeadlineをサポートするファイルでは、保留中のI/O操作がキャンセルされ、
// すぐにErrClosedエラーとともに戻ります。
// Closeは、すでに呼び出されている場合にエラーを返します。
func (f *File) Close() error

<<<<<<< HEAD
// Statはファイルに関する情報を記述するFileInfo構造体を返します。
// エラーが発生した場合は、*PathErrorのタイプになります。
func (f *File) Stat() (FileInfo, error)

// Truncateはファイルのサイズを変更します。
// I/Oオフセットは変更されません。
// エラーが発生した場合、*PathError型となります。
=======
// Stat returns the FileInfo structure describing file.
// If there is an error, it will be of type [*PathError].
func (f *File) Stat() (FileInfo, error)

// Truncate changes the size of the file.
// It does not change the I/O offset.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
func (f *File) Truncate(size int64) error

// Syncはファイルの現在の内容を安定したストレージにコミットします。
// 通常、これはファイルシステムのインメモリコピーのディスクへの最近書き込まれたデータのフラッシュを意味します。
func (f *File) Sync() error

<<<<<<< HEAD
// Truncateは指定されたファイルのサイズを変更します。
// もしファイルがシンボリックリンクなら、リンクの対象のサイズを変更します。
// エラーが発生した場合、*PathErrorの型となります。
func Truncate(name string, size int64) error

// Removeは指定されたファイルまたはディレクトリを削除します。
// エラーが発生した場合、それは*PathError型のエラーです。
=======
// Truncate changes the size of the named file.
// If the file is a symbolic link, it changes the size of the link's target.
// If there is an error, it will be of type [*PathError].
func Truncate(name string, size int64) error

// Remove removes the named file or directory.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
func Remove(name string) error

// Chtimesは、Unixのutime()やutimes()関数に似た、指定したファイルのアクセス時刻と修正時刻を変更します。
// time.Time値がゼロの場合、対応するファイルの時刻は変更されません。
//
<<<<<<< HEAD
// 根底のファイルシステムは、値をより精度の低い単位に切り捨てたり、四捨五入したりする場合があります。
// エラーがある場合、それは*PathError型になります。
=======
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
func Chtimes(name string, atime time.Time, mtime time.Time) error

// Pipeは、接続されたファイルのペアを返します；rから読み取られたバイトはwに書き込まれます。ファイルとエラー（あれば）を返します。
func Pipe() (r *File, w *File, err error)

// Linkはoldnameファイルへのハードリンクとしてnewnameを作成します。
// エラーが発生した場合、*LinkError型になります。
func Link(oldname, newname string) error

// Symlinkはnewnameをoldnameへのシンボリックリンクとして作成します。
// Windowsでは、存在しないoldnameへのシンボリックリンクはファイルのシンボリックリンクとして作成されます。
// その後oldnameがディレクトリとして作成された場合、シンボリックリンクは機能しません。
// エラーが発生した場合、*LinkErrorの型になります。
func Symlink(oldname, newname string) error

<<<<<<< HEAD
// Chownは指定されたファイルの数値化されたuidとgidを変更します。
// ファイルがシンボリックリンクの場合、リンク先のuidとgidを変更します。
// uidまたはgidが-1の場合、その値は変更されません。
// エラーが発生した場合、*PathError型のエラーが返されます。
//
// WindowsまたはPlan 9では、Chownは常にsyscall.EWINDOWSまたはEPLAN9エラーを返し、*PathErrorでラップされます。
func Chown(name string, uid, gid int) error

func Lchown(name string, uid, gid int) error

// Chownは指定されたファイルの数値UIDとGIDを変更します。
// エラーが発生した場合、それは*PathErrorのタイプになります。
func (f *File) Chown(uid, gid int) error

// Chdirは現在の作業ディレクトリを指定されたファイル（ディレクトリである必要があります）に変更します。
// エラーが発生した場合、それは*PathErrorの型です。
=======
// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// A uid or gid of -1 means to not change that value.
// If there is an error, it will be of type [*PathError].
//
// On Windows or Plan 9, Chown always returns the [syscall.EWINDOWS] or
// [syscall.EPLAN9] error, wrapped in [*PathError].
func Chown(name string, uid, gid int) error

// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type [*PathError].
func Lchown(name string, uid, gid int) error

// Chown changes the numeric uid and gid of the named file.
// If there is an error, it will be of type [*PathError].
func (f *File) Chown(uid, gid int) error

// Chdir changes the current working directory to the file,
// which must be a directory.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/release-branch.go1.25
func (f *File) Chdir() error
