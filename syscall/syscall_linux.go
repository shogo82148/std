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

// sys	open(path string, mode int, perm uint32) (fd int, err error)
func Open(path string, mode int, perm uint32) (fd int, err error)

// sys	openat(dirfd int, path string, flags int, mode uint32) (fd int, err error)
func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error)

// sysnb	pipe(p *[2]_C_int) (err error)
func Pipe(p []int) (err error)

// sys	utimes(path string, times *[2]Timeval) (err error)
func Utimes(path string, tv []Timeval) (err error)

// sys	futimesat(dirfd int, path *byte, times *[2]Timeval) (err error)
func Futimesat(dirfd int, path string, tv []Timeval) (err error)

func Futimes(fd int, tv []Timeval) (err error)

const ImplementsGetwd = true

// sys	Getcwd(buf []byte) (n int, err error)
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

// sys	wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, err error)
func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, err error)

func Mkfifo(path string, mode uint32) (err error)

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

func Accept(fd int) (nfd int, sa Sockaddr, err error)

func Getsockname(fd int) (sa Sockaddr, err error)

func Getpeername(fd int) (sa Sockaddr, err error)

func Bind(fd int, sa Sockaddr) (err error)

func Connect(fd int, sa Sockaddr) (err error)

func Socket(domain, typ, proto int) (fd int, err error)

func Socketpair(domain, typ, proto int) (fd [2]int, err error)

func GetsockoptInt(fd, level, opt int) (value int, err error)

func GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error)

func GetsockoptIPMreq(fd, level, opt int) (*IPMreq, error)

func GetsockoptIPMreqn(fd, level, opt int) (*IPMreqn, error)

func GetsockoptIPv6Mreq(fd, level, opt int) (*IPv6Mreq, error)

func SetsockoptInt(fd, level, opt int, value int) (err error)

func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err error)

func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (err error)

func SetsockoptLinger(fd, level, opt int, l *Linger) (err error)

func SetsockoptIPMreq(fd, level, opt int, mreq *IPMreq) (err error)

func SetsockoptIPMreqn(fd, level, opt int, mreq *IPMreqn) (err error)

func SetsockoptIPv6Mreq(fd, level, opt int, mreq *IPv6Mreq) (err error)

func SetsockoptString(fd, level, opt int, s string) (err error)

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error)

func Sendto(fd int, p []byte, flags int, to Sockaddr) (err error)

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from Sockaddr, err error)

func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int) (err error)

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

func PtraceSingleStep(pid int) (err error)

func PtraceAttach(pid int) (err error)

func PtraceDetach(pid int) (err error)

// sys	reboot(magic1 uint, magic2 uint, cmd int, arg string) (err error)
func Reboot(cmd int) (err error)

func ReadDirent(fd int, buf []byte) (n int, err error)

func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string)

// sys	mount(source string, target string, fstype string, flags uintptr, data *byte) (err error)
func Mount(source string, target string, fstype string, flags uintptr, data string) (err error)

func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)

func Munmap(b []byte) (err error)
