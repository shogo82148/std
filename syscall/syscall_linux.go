// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and
// wrap it in our own nicer implementation.

package syscall

func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

//go:uintptrkeepalive
//go:nosplit
//go:norace
//go:linkname RawSyscall
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

//go:uintptrkeepalive
//go:nosplit
//go:linkname Syscall
func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

//go:uintptrkeepalive
//go:nosplit
//go:linkname Syscall6
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

func Access(path string, mode uint32) (err error)

func Chmod(path string, mode uint32) (err error)

func Chown(path string, uid int, gid int) (err error)

func Creat(path string, mode uint32) (fd int, err error)

func Faccessat(dirfd int, path string, mode uint32, flags int) (err error)

func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error)

func Link(oldpath string, newpath string) (err error)

func Mkdir(path string, mode uint32) (err error)

func Mknod(path string, mode uint32, dev int) (err error)

func Open(path string, mode int, perm uint32) (fd int, err error)

func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error)

func Pipe(p []int) error

func Pipe2(p []int, flags int) error

func Readlink(path string, buf []byte) (n int, err error)

func Rename(oldpath string, newpath string) (err error)

func Rmdir(path string) error

func Symlink(oldpath string, newpath string) (err error)

func Unlink(path string) error

func Unlinkat(dirfd int, path string) error

func Utimes(path string, tv []Timeval) (err error)

func UtimesNano(path string, ts []Timespec) (err error)

func Futimesat(dirfd int, path string, tv []Timeval) (err error)

func Futimes(fd int, tv []Timeval) (err error)

const ImplementsGetwd = true

func Getwd() (wd string, err error)

func Getgroups() (gids []int, err error)

func Setgroups(gids []int) (err error)

type WaitStatus uint32

func (w WaitStatus) Exited() bool

func (w WaitStatus) Signaled() bool

func (w WaitStatus) Stopped() bool

func (w WaitStatus) Continued() bool

func (w WaitStatus) CoreDump() bool

func (w WaitStatus) ExitStatus() int

func (w WaitStatus) Signal() Signal

func (w WaitStatus) StopSignal() Signal

func (w WaitStatus) TrapCause() int

func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, err error)

func Mkfifo(path string, mode uint32) (err error)

type SockaddrLinklayer struct {
	Protocol uint16
	Ifindex  int
	Hatype   uint16
	Pkttype  uint8
	Halen    uint8
	Addr     [8]byte
	raw      RawSockaddrLinklayer
}

type SockaddrNetlink struct {
	Family uint16
	Pad    uint16
	Pid    uint32
	Groups uint32
	raw    RawSockaddrNetlink
}

func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error)

func Getsockname(fd int) (sa Sockaddr, err error)

func GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error)

func GetsockoptIPMreq(fd, level, opt int) (*IPMreq, error)

func GetsockoptIPMreqn(fd, level, opt int) (*IPMreqn, error)

func GetsockoptIPv6Mreq(fd, level, opt int) (*IPv6Mreq, error)

func GetsockoptIPv6MTUInfo(fd, level, opt int) (*IPv6MTUInfo, error)

func GetsockoptICMPv6Filter(fd, level, opt int) (*ICMPv6Filter, error)

func GetsockoptUcred(fd, level, opt int) (*Ucred, error)

func SetsockoptIPMreqn(fd, level, opt int, mreq *IPMreqn) (err error)

// BindToDevice binds the socket associated with fd to device.
func BindToDevice(fd int, device string) (err error)

func PtracePeekText(pid int, addr uintptr, out []byte) (count int, err error)

func PtracePeekData(pid int, addr uintptr, out []byte) (count int, err error)

func PtracePokeText(pid int, addr uintptr, data []byte) (count int, err error)

func PtracePokeData(pid int, addr uintptr, data []byte) (count int, err error)

func PtraceGetRegs(pid int, regsout *PtraceRegs) (err error)

func PtraceSetRegs(pid int, regs *PtraceRegs) (err error)

func PtraceSetOptions(pid int, options int) (err error)

func PtraceGetEventMsg(pid int) (msg uint, err error)

func PtraceCont(pid int, signal int) (err error)

func PtraceSyscall(pid int, signal int) (err error)

func PtraceSingleStep(pid int) (err error)

func PtraceAttach(pid int) (err error)

func PtraceDetach(pid int) (err error)

func Reboot(cmd int) (err error)

func ReadDirent(fd int, buf []byte) (n int, err error)

func Mount(source string, target string, fstype string, flags uintptr, data string) (err error)

func Getpgrp() (pid int)

// AllThreadsSyscall performs a syscall on each OS thread of the Go
// runtime. It first invokes the syscall on one thread. Should that
// invocation fail, it returns immediately with the error status.
// Otherwise, it invokes the syscall on all of the remaining threads
// in parallel. It will terminate the program if it observes any
// invoked syscall's return value differs from that of the first
// invocation.
//
// AllThreadsSyscall is intended for emulating simultaneous
// process-wide state changes that require consistently modifying
// per-thread state of the Go runtime.
//
// AllThreadsSyscall is unaware of any threads that are launched
// explicitly by cgo linked code, so the function always returns
// ENOTSUP in binaries that use cgo.
//
//go:uintptrescapes
func AllThreadsSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

// AllThreadsSyscall6 is like AllThreadsSyscall, but extended to six
// arguments.
//
//go:uintptrescapes
func AllThreadsSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

func Setegid(egid int) (err error)

func Seteuid(euid int) (err error)

func Setgid(gid int) (err error)

func Setregid(rgid, egid int) (err error)

func Setresgid(rgid, egid, sgid int) (err error)

func Setresuid(ruid, euid, suid int) (err error)

func Setreuid(ruid, euid int) (err error)

func Setuid(uid int) (err error)

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)

func Munmap(b []byte) (err error)
