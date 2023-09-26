// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Aix system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and
// wrap it in our own nicer implementation.

package syscall

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// Constant expected by package but not supported
const (
	_ = iota
	TIOCSCTTY
	SYS_EXECVE
	SYS_FCNTL
)

const (
	F_DUPFD_CLOEXEC = 0
	// AF_LOCAL doesn't exist on AIX
	AF_LOCAL = AF_UNIX
)

func (ts *StTimespec_t) Unix() (sec int64, nsec int64)

func (ts *StTimespec_t) Nano() int64

func Access(path string, mode uint32) (err error)

func Pipe(p []int) (err error)

func Readlink(path string, buf []byte) (n int, err error)

func Utimes(path string, tv []Timeval) error

func UtimesNano(path string, ts []Timespec) error

func Unlinkat(dirfd int, path string) (err error)

const ImplementsGetwd = true

func Getwd() (ret string, err error)

func Getcwd(buf []byte) (n int, err error)

func Getgroups() (gids []int, err error)

func Setgroups(gids []int) (err error)

func Gettimeofday(tv *Timeval) (err error)

func ReadDirent(fd int, buf []byte) (n int, err error)

func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, err error)

func Fsync(fd int) error

func Getsockname(fd int) (sa Sockaddr, err error)

func Accept(fd int) (nfd int, sa Sockaddr, err error)

type SockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [120]uint8
	raw    RawSockaddrDatalink
}

type WaitStatus uint32

func (w WaitStatus) Stopped() bool
func (w WaitStatus) StopSignal() Signal

func (w WaitStatus) Exited() bool
func (w WaitStatus) ExitStatus() int

func (w WaitStatus) Signaled() bool
func (w WaitStatus) Signal() Signal

func (w WaitStatus) Continued() bool

func (w WaitStatus) CoreDump() bool

func (w WaitStatus) TrapCause() int

func PtracePeekText(pid int, addr uintptr, out []byte) (count int, err error)

func PtracePeekData(pid int, addr uintptr, out []byte) (count int, err error)

func PtracePokeText(pid int, addr uintptr, data []byte) (count int, err error)

func PtracePokeData(pid int, addr uintptr, data []byte) (count int, err error)

func PtraceCont(pid int, signal int) (err error)

func PtraceSingleStep(pid int) (err error)

func PtraceAttach(pid int) (err error)

func PtraceDetach(pid int) (err error)

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)

func Munmap(b []byte) (err error)
