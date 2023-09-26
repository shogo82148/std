// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package syscall

import (
	"github.com/shogo82148/std/unsafe"
)

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// An Errno is an unsigned number describing an error condition.
// It implements the error interface. The zero Errno is by convention
// a non-error, so code to convert from Errno to error should use:
//
//	err = nil
//	if errno != 0 {
//		err = errno
//	}
type Errno uintptr

func (e Errno) Error() string

func (e Errno) Is(target error) bool

func (e Errno) Temporary() bool

func (e Errno) Timeout() bool

// Do the interface allocations only once for common
// Errno values.

// A Signal is a number describing a process signal.
// It implements the os.Signal interface.
type Signal int

func (s Signal) Signal()

func (s Signal) String() string

func Read(fd int, p []byte) (n int, err error)

func Write(fd int, p []byte) (n int, err error)

// For testing: clients can set this flag to force
// creation of IPv6 sockets to return EAFNOSUPPORT.
var SocketDisableIPv6 bool

type Sockaddr interface {
	sockaddr() (ptr unsafe.Pointer, len _Socklen, err error)
}

type SockaddrInet4 struct {
	Port int
	Addr [4]byte
	raw  RawSockaddrInet4
}

type SockaddrInet6 struct {
	Port   int
	ZoneId uint32
	Addr   [16]byte
	raw    RawSockaddrInet6
}

type SockaddrUnix struct {
	Name string
	raw  RawSockaddrUnix
}

func Bind(fd int, sa Sockaddr) (err error)

func Connect(fd int, sa Sockaddr) (err error)

func Getpeername(fd int) (sa Sockaddr, err error)

func GetsockoptInt(fd, level, opt int) (value int, err error)

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error)

func Sendto(fd int, p []byte, flags int, to Sockaddr) (err error)

func SetsockoptByte(fd, level, opt int, value byte) (err error)

func SetsockoptInt(fd, level, opt int, value int) (err error)

func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err error)

func SetsockoptIPMreq(fd, level, opt int, mreq *IPMreq) (err error)

func SetsockoptIPv6Mreq(fd, level, opt int, mreq *IPv6Mreq) (err error)

func SetsockoptICMPv6Filter(fd, level, opt int, filter *ICMPv6Filter) error

func SetsockoptLinger(fd, level, opt int, l *Linger) (err error)

func SetsockoptString(fd, level, opt int, s string) (err error)

func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (err error)

func Socket(domain, typ, proto int) (fd int, err error)

func Socketpair(domain, typ, proto int) (fd [2]int, err error)

func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
