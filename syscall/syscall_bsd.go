// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd
// +build darwin dragonfly freebsd netbsd openbsd

// BSD system call wrappers shared by *BSD based systems
// including OS X (Darwin) and FreeBSD.  Like the other
// syscall_*.go files it is compiled as Go code but also
// used as input to mksyscall which parses the //sys
// lines and generates system call stubs.

package syscall

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

func Accept(fd int) (nfd int, sa Sockaddr, err error)

func Getsockname(fd int) (sa Sockaddr, err error)

func GetsockoptByte(fd, level, opt int) (value byte, err error)

func GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error)

func GetsockoptIPMreq(fd, level, opt int) (*IPMreq, error)

func GetsockoptIPv6Mreq(fd, level, opt int) (*IPv6Mreq, error)

func GetsockoptIPv6MTUInfo(fd, level, opt int) (*IPv6MTUInfo, error)

func GetsockoptICMPv6Filter(fd, level, opt int) (*ICMPv6Filter, error)

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error)

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error)

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error)

func Kevent(kq int, changes, events []Kevent_t, timeout *Timespec) (n int, err error)

func Sysctl(name string) (value string, err error)

func SysctlUint32(name string) (value uint32, err error)

// sys	utimes(path string, timeval *[2]Timeval) (err error)
func Utimes(path string, tv []Timeval) (err error)

func UtimesNano(path string, ts []Timespec) error

// sys	futimes(fd int, timeval *[2]Timeval) (err error)
func Futimes(fd int, tv []Timeval) (err error)

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)

func Munmap(b []byte) (err error)
