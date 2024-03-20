// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

package os

import (
	"github.com/shogo82148/std/time"
)

// Closeは [File] を閉じて、I/Oに使用できなくします。
// [File.SetDeadline] をサポートするファイルでは、保留中のI/O操作はキャンセルされ、
// [ErrClosed] エラーとともに即座に返ります。
// Closeが既に呼び出されている場合はエラーが返されます。
func (f *File) Close() error

// Chownは指定されたファイルの数値UIDとGIDを変更します。
// ファイルがシンボリックリンクの場合、リンク先のUIDとGIDを変更します。
// -1のUIDまたはGIDはその値を変更しないことを意味します。
// エラーが発生した場合、型 [*PathError] になります。
//
// WindowsまたはPlan 9の場合、Chownは常に [syscall.EWINDOWS] または
// EPLAN9のエラーを*PathErrorでラップして返します。
func Chown(name string, uid, gid int) error

// Lchownは指定されたファイルの数値UIDとGIDを変更します。
// ファイルがシンボリックリンクの場合、リンク自体のUIDとGIDを変更します。
// エラーが発生した場合は、[*PathError] 型のエラーが返されます。
//
// Windowsでは、常に [syscall.EWINDOWS] エラーが返され、*PathErrorでラップされます。
func Lchown(name string, uid, gid int) error

// Chownは指定したファイルの数値uidとgidを変更します。
// エラーが発生した場合、それは [*PathError] の型です。
//
// Windowsでは、いつも [syscall.EWINDOWS] のエラーを返し、*PathErrorにラップします。
func (f *File) Chown(uid, gid int) error

// Truncateはファイルのサイズを変更します。
// I/Oオフセットは変更しません。
// エラーが発生した場合、[*PathError] 型になります。
func (f *File) Truncate(size int64) error

// Syncはファイルの現在の内容を安定したストレージへコミットします。
// 通常、これはファイルシステムのメモリ上のコピーをフラッシュし、
// 最近書き込まれたデータをディスクに書き込むことを意味しています。
func (f *File) Sync() error

// Chtimesは指定されたファイルのアクセス時間と修正時間を変更します。これはUnixのutime()やutimes()関数と同様です。
// ゼロの [time.Time] 値は、対応するファイルの時間を変更しません。
//
// 基礎となるファイルシステムは、値を切り捨てたり、より正確でない時間単位に丸めたりするかもしれません。
// エラーが発生した場合、[*PathError] 型になります。
func Chtimes(name string, atime time.Time, mtime time.Time) error

// Chdirは現在の作業ディレクトリを指定されたディレクトリpathに変更します。
// pathはディレクトリでなければなりません。
// エラーが発生した場合、[*PathError] の型で返されます。
func (f *File) Chdir() error
