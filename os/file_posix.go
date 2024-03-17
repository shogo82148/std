// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

package os

import (
	"github.com/shogo82148/std/time"
)

<<<<<<< HEAD
// Closeはファイルを閉じて、I/Oに使用できなくします。
// SetDeadlineをサポートするファイルでは、保留中のI/O操作はキャンセルされ、
// ErrClosedエラーとともに即座に返ります。
// Closeが既に呼び出されている場合はエラーが返されます。
func (f *File) Close() error

// Chownは指定されたファイルの数値UIDとGIDを変更します。
// ファイルがシンボリックリンクの場合、リンク先のUIDとGIDを変更します。
// -1のUIDまたはGIDはその値を変更しないことを意味します。
// エラーが発生した場合、型*PathErrorになります。
//
// WindowsまたはPlan 9の場合、Chownは常にsyscall.EWINDOWSまたは
// EPLAN9のエラーを*PathErrorでラップして返します。
func Chown(name string, uid, gid int) error

// Lchownは指定されたファイルの数値UIDとGIDを変更します。
// ファイルがシンボリックリンクの場合、リンク自体のUIDとGIDを変更します。
// エラーが発生した場合は、*PathError型のエラーが返されます。
//
// Windowsでは、常にsyscall.EWINDOWSエラーが返され、*PathErrorでラップされます。
func Lchown(name string, uid, gid int) error

// Chownは指定したファイルの数値uidとgidを変更します。
// エラーが発生した場合、それは*PathErrorの型です。
//
// Windowsでは、いつもsyscall.EWINDOWSのエラーを返し、*PathErrorにラップします。
func (f *File) Chown(uid, gid int) error

// Truncateはファイルのサイズを変更します。
// I/Oオフセットは変更しません。
// エラーが発生した場合、*PathError型になります。
=======
// Close closes the [File], rendering it unusable for I/O.
// On files that support [File.SetDeadline], any pending I/O operations will
// be canceled and return immediately with an [ErrClosed] error.
// Close will return an error if it has already been called.
func (f *File) Close() error

// Chown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link's target.
// A uid or gid of -1 means to not change that value.
// If there is an error, it will be of type [*PathError].
//
// On Windows or Plan 9, Chown always returns the [syscall.EWINDOWS] or
// EPLAN9 error, wrapped in *PathError.
func Chown(name string, uid, gid int) error

// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type [*PathError].
//
// On Windows, it always returns the [syscall.EWINDOWS] error, wrapped
// in *PathError.
func Lchown(name string, uid, gid int) error

// Chown changes the numeric uid and gid of the named file.
// If there is an error, it will be of type [*PathError].
//
// On Windows, it always returns the [syscall.EWINDOWS] error, wrapped
// in *PathError.
func (f *File) Chown(uid, gid int) error

// Truncate changes the size of the file.
// It does not change the I/O offset.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/master
func (f *File) Truncate(size int64) error

// Syncはファイルの現在の内容を安定したストレージへコミットします。
// 通常、これはファイルシステムのメモリ上のコピーをフラッシュし、
// 最近書き込まれたデータをディスクに書き込むことを意味しています。
func (f *File) Sync() error

<<<<<<< HEAD
// Chtimesは指定されたファイルのアクセス時間と修正時間を変更します。これはUnixのutime()やutimes()関数と同様です。
// ゼロのtime.Time値は、対応するファイルの時間を変更しません。
//
// 基礎となるファイルシステムは、値を切り捨てたり、より正確でない時間単位に丸めたりするかもしれません。
// エラーが発生した場合、*PathError型になります。
func Chtimes(name string, atime time.Time, mtime time.Time) error

// Chdirは現在の作業ディレクトリを指定されたディレクトリpathに変更します。
// pathはディレクトリでなければなりません。
// エラーが発生した場合、*PathErrorの型で返されます。
=======
// Chtimes changes the access and modification times of the named
// file, similar to the Unix utime() or utimes() functions.
// A zero [time.Time] value will leave the corresponding file time unchanged.
//
// The underlying filesystem may truncate or round the values to a
// less precise time unit.
// If there is an error, it will be of type [*PathError].
func Chtimes(name string, atime time.Time, mtime time.Time) error

// Chdir changes the current working directory to the file,
// which must be a directory.
// If there is an error, it will be of type [*PathError].
>>>>>>> upstream/master
func (f *File) Chdir() error
