// Copyright 2009,2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// NetBSD system calls.
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
	Data   [24]int8
	raw    RawSockaddrDatalink
}

func Syscall9(trap, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)

// ParseDirent parses up to max directory entries in buf,
// appending the names to names. It returns the number
// bytes consumed from buf, the number of entries added
// to names, and the new names slice.
func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string)

// sysnb pipe2(p *[2]_C_int, flags _C_int) (err error)
func Pipe(p []int) (err error)

// sys getdents(fd int, buf []byte) (n int, err error)
func Getdirentries(fd int, buf []byte, basep *uintptr) (n int, err error)

// TODO
func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
