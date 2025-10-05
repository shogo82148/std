// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/time"
)

// DevNullはオペレーティングシステムの「nullデバイス」の名称です。
// Unix系のシステムでは、"/dev/null"となります。Windowsでは、"NUL"です。
const DevNull = "/dev/null"

// Closeは、Fileを閉じて、I/Oに使用できなくします。
// SetDeadlineをサポートするファイルでは、保留中のI/O操作がキャンセルされ、
// すぐにErrClosedエラーとともに戻ります。
// Closeは、すでに呼び出されている場合にエラーを返します。
func (f *File) Close() error

// Statはファイルを記述するFileInfo構造体を返します。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func (f *File) Stat() (FileInfo, error)

// Truncateはファイルのサイズを変更します。
// I/Oオフセットは変更しません。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func (f *File) Truncate(size int64) error

// Syncはファイルの現在の内容を安定したストレージにコミットします。
// 通常、これはファイルシステムのインメモリコピーのディスクへの最近書き込まれたデータのフラッシュを意味します。
func (f *File) Sync() error

// Truncateは指定されたファイルのサイズを変更します。
// ファイルがシンボリックリンクの場合は、リンク先のサイズを変更します。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func Truncate(name string, size int64) error

// Removeは指定されたファイルまたはディレクトリを削除します。
// エラーが発生した場合、エラーは[*PathError]型になります。
func Remove(name string) error

// Chtimesは、Unixのutime()やutimes()関数に似た、指定したファイルのアクセス時刻と修正時刻を変更します。
// time.Time値がゼロの場合、対応するファイルの時刻は変更されません。
//
// 根底のファイルシステムは、値をより精度の低い単位に切り捨てたり、四捨五入したりする場合があります。
// エラーがある場合、それは [*PathError] 型になります。
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

// Chownは指定されたファイルの数値uidとgidを変更します。
// ファイルがシンボリックリンクの場合は、リンク先のuidとgidを変更します。
// uidまたはgidに-1を指定すると、その値は変更されません。
// エラーが発生した場合、エラーは [*PathError] 型になります。
//
// WindowsまたはPlan 9では、Chownは常に [syscall.EWINDOWS] または
// [syscall.EPLAN9] エラーを [*PathError] でラップして返します。
func Chown(name string, uid, gid int) error

// Lchownは指定されたファイルの数値uidとgidを変更します。
// ファイルがシンボリックリンクの場合は、リンク自体のuidとgidを変更します。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func Lchown(name string, uid, gid int) error

// Chownは指定されたファイルの数値uidとgidを変更します。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func (f *File) Chown(uid, gid int) error

// Chdirはカレントワーキングディレクトリをそのファイルに変更します。
// ファイルはディレクトリである必要があります。
// エラーが発生した場合、エラーは [*PathError] 型になります。
func (f *File) Chdir() error
