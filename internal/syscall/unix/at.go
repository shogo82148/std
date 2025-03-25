// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || linux || netbsd || (openbsd && mips64)

package unix

func Unlinkat(dirfd int, path string, flags int) error

func Openat(dirfd int, path string, flags int, perm uint32) (int, error)

func Readlinkat(dirfd int, path string, buf []byte) (int, error)

func Mkdirat(dirfd int, path string, mode uint32) error

func Fchmodat(dirfd int, path string, mode uint32, flags int) error

func Fchownat(dirfd int, path string, uid, gid int, flags int) error

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) error

func Linkat(olddirfd int, oldpath string, newdirfd int, newpath string, flag int) error
