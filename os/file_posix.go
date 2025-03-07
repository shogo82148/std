// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

package os

import (
	"github.com/shogo82148/std/time"
)

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
// [syscall.EPLAN9] error, wrapped in [*PathError].
func Chown(name string, uid, gid int) error

// Lchown changes the numeric uid and gid of the named file.
// If the file is a symbolic link, it changes the uid and gid of the link itself.
// If there is an error, it will be of type [*PathError].
//
// On Windows, it always returns the [syscall.EWINDOWS] error, wrapped
// in [*PathError].
func Lchown(name string, uid, gid int) error

// Chown changes the numeric uid and gid of the named file.
// If there is an error, it will be of type [*PathError].
//
// On Windows, it always returns the [syscall.EWINDOWS] error, wrapped
// in [*PathError].
func (f *File) Chown(uid, gid int) error

// Truncate changes the size of the file.
// It does not change the I/O offset.
// If there is an error, it will be of type [*PathError].
func (f *File) Truncate(size int64) error

// Sync commits the current contents of the file to stable storage.
// Typically, this means flushing the file system's in-memory copy
// of recently written data to disk.
func (f *File) Sync() error

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
func (f *File) Chdir() error
