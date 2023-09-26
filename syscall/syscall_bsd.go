// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || netbsd || openbsd
// +build darwin freebsd netbsd openbsd

// BSD system call wrappers shared by *BSD based systems
// including OS X (Darwin) and FreeBSD.  Like the other
// syscall_*.go files it is compiled as Go code but also
// used as input to mksyscall which parses the //sys
// lines and generates system call stubs.

package syscall

/*
 * Pseudo-system calls
 */
// The const provides a compile-time constant so clients
// can adjust to whether there is a working Getwd and avoid
// even linking this function into the binary.  See ../os/getwd.go.
const ImplementsGetwd = false

func Getwd() (string, error)

func Getgroups() (gids []int, err error)

func Setgroups(gids []int) (err error)

func ReadDirent(fd int, buf []byte) (n int, err error)

type WaitStatus uint32

func (w WaitStatus) Exited() bool

func (w WaitStatus) ExitStatus() int

func (w WaitStatus) Signaled() bool

func (w WaitStatus) Signal() Signal

func (w WaitStatus) CoreDump() bool

func (w WaitStatus) Stopped() bool

func (w WaitStatus) Continued() bool

func (w WaitStatus) StopSignal() Signal

func (w WaitStatus) TrapCause() int

func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, err error)

// For testing: clients can set this flag to force
// creation of IPv6 sockets to return EAFNOSUPPORT.
var SocketDisableIPv6 bool

type Sockaddr interface {
	sockaddr() (ptr uintptr, len _Socklen, err error)
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

func Accept(fd int) (nfd int, sa Sockaddr, err error)

func Getsockname(fd int) (sa Sockaddr, err error)

func Getpeername(fd int) (sa Sockaddr, err error)

func Bind(fd int, sa Sockaddr) (err error)

func Connect(fd int, sa Sockaddr) (err error)

func Socket(domain, typ, proto int) (fd int, err error)

func Socketpair(domain, typ, proto int) (fd [2]int, err error)

func GetsockoptByte(fd, level, opt int) (value byte, err error)

func GetsockoptInt(fd, level, opt int) (value int, err error)

func GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error)

func GetsockoptIPMreq(fd, level, opt int) (*IPMreq, error)

func GetsockoptIPv6Mreq(fd, level, opt int) (*IPv6Mreq, error)

func SetsockoptByte(fd, level, opt int, value byte) (err error)

func SetsockoptInt(fd, level, opt int, value int) (err error)

func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err error)

func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (err error)

func SetsockoptLinger(fd, level, opt int, l *Linger) (err error)

func SetsockoptIPMreq(fd, level, opt int, mreq *IPMreq) (err error)

func SetsockoptIPv6Mreq(fd, level, opt int, mreq *IPv6Mreq) (err error)

func SetsockoptString(fd, level, opt int, s string) (err error)

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error)

func Sendto(fd int, p []byte, flags int, to Sockaddr) (err error)

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error)

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error)

func Kevent(kq int, changes, events []Kevent_t, timeout *Timespec) (n int, err error)

func Sysctl(name string) (value string, err error)

func SysctlUint32(name string) (value uint32, err error)

// sys	utimes(path string, timeval *[2]Timeval) (err error)
func Utimes(path string, tv []Timeval) (err error)

// sys	futimes(fd int, timeval *[2]Timeval) (err error)
func Futimes(fd int, tv []Timeval) (err error)

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)

func Munmap(b []byte) (err error)
