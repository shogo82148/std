// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/time"
)

// Fdは、開かれたファイルを参照する整数型のPlan 9ファイルディスクリプタを返します。
// fがクローズされると、ファイルディスクリプタは無効になります。
// fがガベージコレクションされると、ファイルディスクリプタをクローズするファイナライザが実行される可能性があります。
// この場合、ランタイムのSetFinalizerに関する詳細は、ファイナライザがいつ実行されるかについての情報を参照してください。Unixシステムでは、これによってSetDeadlineメソッドが動作しなくなります。
//
// 代替案として、f.SyscallConnメソッドを参照してください。
func (f *File) Fd() uintptr

// NewFile は指定されたファイルディスクリプタと名前で新しいファイルを返します。
// fd が有効なファイルディスクリプタでない場合、返される値は nil になります。
func NewFile(fd uintptr, name string) *File

// DevNullはオペレーティングシステムの「nullデバイス」の名称です。
// Unix系のシステムでは、"/dev/null"となります。Windowsでは、"NUL"です。
const DevNull = "/dev/null"

// Closeは、Fileを閉じて、I/Oに使用できなくします。
// SetDeadlineをサポートするファイルでは、保留中のI/O操作がキャンセルされ、
// すぐにErrClosedエラーとともに戻ります。
// Closeは、すでに呼び出されている場合にエラーを返します。
func (f *File) Close() error

// Statはファイルに関する情報を記述するFileInfo構造体を返します。
// エラーが発生した場合は、*PathErrorのタイプになります。
func (f *File) Stat() (FileInfo, error)

// Truncateはファイルのサイズを変更します。
// I/Oオフセットは変更されません。
// エラーが発生した場合、*PathError型となります。
func (f *File) Truncate(size int64) error

// Syncはファイルの現在の内容を安定したストレージにコミットします。
// 通常、これはファイルシステムのインメモリコピーのディスクへの最近書き込まれたデータのフラッシュを意味します。
func (f *File) Sync() error

// Truncateは指定されたファイルのサイズを変更します。
// もしファイルがシンボリックリンクなら、リンクの対象のサイズを変更します。
// エラーが発生した場合、*PathErrorの型となります。
func Truncate(name string, size int64) error

// Removeは指定されたファイルまたはディレクトリを削除します。
// エラーが発生した場合、それは*PathError型のエラーです。
func Remove(name string) error

// Chtimesは、Unixのutime()やutimes()関数に似た、指定したファイルのアクセス時刻と修正時刻を変更します。
// time.Time値がゼロの場合、対応するファイルの時刻は変更されません。
//
// 根底のファイルシステムは、値をより精度の低い単位に切り捨てたり、四捨五入したりする場合があります。
// エラーがある場合、それは*PathError型になります。
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
// Readlinkは指定されたシンボリックリンクの宛先を返します。
// エラーが発生した場合、*PathError型になります。
func Readlink(name string) (string, error)

// Chownは指定されたファイルの数値化されたuidとgidを変更します。
// ファイルがシンボリックリンクの場合、リンク先のuidとgidを変更します。
// uidまたはgidが-1の場合、その値は変更されません。
// エラーが発生した場合、*PathError型のエラーが返されます。
=======
// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// A uid or gid of -1 means to not change that value.
// If there is an error, it will be of type *PathError.
>>>>>>> upstream/master
//
// WindowsまたはPlan 9では、Chownは常にsyscall.EWINDOWSまたはEPLAN9エラーを返し、*PathErrorでラップされます。
func Chown(name string, uid, gid int) error

func Lchown(name string, uid, gid int) error

// Chownは指定されたファイルの数値UIDとGIDを変更します。
// エラーが発生した場合、それは*PathErrorのタイプになります。
func (f *File) Chown(uid, gid int) error

// Chdirは現在の作業ディレクトリを指定されたファイル（ディレクトリである必要があります）に変更します。
// エラーが発生した場合、それは*PathErrorの型です。
func (f *File) Chdir() error
