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
