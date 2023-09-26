// Copyright 2009,2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// FreeBSD system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and wrap
// it in our own nicer implementation, either here or in
// syscall_bsd.go or syscall_unix.go.

package syscall

// See https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/versions.html.

// INO64_FIRST from /usr/src/lib/libc/sys/compat-ino64.h

type SockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [46]int8
	raw    RawSockaddrDatalink
}

func Pipe(p []int) error

func Pipe2(p []int, flags int) error

func GetsockoptIPMreqn(fd, level, opt int) (*IPMreqn, error)

func SetsockoptIPMreqn(fd, level, opt int, mreq *IPMreqn) (err error)

func Accept4(fd, flags int) (nfd int, sa Sockaddr, err error)

func Getfsstat(buf []Statfs_t, flags int) (n int, err error)

func Stat(path string, st *Stat_t) (err error)

func Lstat(path string, st *Stat_t) (err error)

func Fstat(fd int, st *Stat_t) (err error)

func Fstatat(fd int, path string, st *Stat_t, flags int) (err error)

func Statfs(path string, st *Statfs_t) (err error)

func Fstatfs(fd int, st *Statfs_t) (err error)

func Getdirentries(fd int, buf []byte, basep *uintptr) (n int, err error)

func Mknod(path string, mode uint32, dev uint64) (err error)
