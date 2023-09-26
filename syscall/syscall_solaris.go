// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Solaris system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and wrap
// it in our own nicer implementation, either here or in
// syscall_solaris.go or syscall_unix.go.

package syscall

type SockaddrDatalink struct {
	Family uint16
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [244]int8
	raw    RawSockaddrDatalink
}

func Pipe(p []int) (err error)

func Getsockname(fd int) (sa Sockaddr, err error)

const ImplementsGetwd = true

func Getwd() (wd string, err error)

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

func Gethostname() (name string, err error)

func UtimesNano(path string, ts []Timespec) error

// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command.
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error

func Accept(fd int) (nfd int, sa Sockaddr, err error)

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error)

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error)

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error)

func Getexecname() (path string, err error)

func Utimes(path string, tv []Timeval) error
