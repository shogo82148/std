// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

package syscall

func Open(path string, openmode int, perm uint32) (int, error)

func Close(fd int) error

func CloseOnExec(fd int)

func Mkdir(path string, perm uint32) error

func ReadDirent(fd int, buf []byte) (int, error)

func Stat(path string, st *Stat_t) error

func Lstat(path string, st *Stat_t) error

func Fstat(fd int, st *Stat_t) error

func Unlink(path string) error

func Rmdir(path string) error

func Chmod(path string, mode uint32) error

func Fchmod(fd int, mode uint32) error

func Chown(path string, uid, gid int) error

func Fchown(fd int, uid, gid int) error

func Lchown(path string, uid, gid int) error

func UtimesNano(path string, ts []Timespec) error

func Rename(from, to string) error

func Truncate(path string, length int64) error

func Ftruncate(fd int, length int64) error

func Getcwd(buf []byte) (n int, err error)

func Chdir(path string) (err error)

func Fchdir(fd int) error

func Readlink(path string, buf []byte) (n int, err error)

func Link(path, link string) error

func Symlink(path, link string) error

func Fsync(fd int) error

func Read(fd int, b []byte) (int, error)

func Write(fd int, b []byte) (int, error)

func Pread(fd int, b []byte, offset int64) (int, error)

func Pwrite(fd int, b []byte, offset int64) (int, error)

func Seek(fd int, offset int64, whence int) (int64, error)

func Dup(fd int) (int, error)

func Dup2(fd, newfd int) error

func Pipe(fd []int) error
