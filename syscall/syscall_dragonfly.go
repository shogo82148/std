// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DragonFly BSD system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and wrap
// it in our own nicer implementation, either here or in
// syscall_bsd.go or syscall_unix.go.

package syscall

// See version list in https://github.com/DragonFlyBSD/DragonFlyBSD/blob/master/sys/sys/param.h

// First __DragonFly_version after September 2019 ABI changes
// http://lists.dragonflybsd.org/pipermail/users/2019-September/358280.html

type SockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [12]int8
	Rcf    uint16
	Route  [16]uint16
	raw    RawSockaddrDatalink
}

func Pipe(p []int) (err error)

// sys	extpread(fd int, p []byte, flags int, offset int64) (n int, err error)
func Pread(fd int, p []byte, offset int64) (n int, err error)

// sys	extpwrite(fd int, p []byte, flags int, offset int64) (n int, err error)
func Pwrite(fd int, p []byte, offset int64) (n int, err error)

func Accept4(fd, flags int) (nfd int, sa Sockaddr, err error)

func Getfsstat(buf []Statfs_t, flags int) (n int, err error)
