// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows system calls.

package syscall

type Handle uintptr

const InvalidHandle = ^Handle(0)

// StringToUTF16 returns the UTF-16 encoding of the UTF-8 string s.
// If s contains a NUL byte this function panics instead of
// returning an error.
func StringToUTF16(s string) []uint16

// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL removed.
func UTF16ToString(s []uint16) string

// StringToUTF16Ptr returns pointer to the UTF-16 encoding of the UTF-8 string s, with a terminating NUL added.
// If s contains a NUL byte this function panics instead of
// returning an error.
func StringToUTF16Ptr(s string) *uint16

func Getpagesize() int

// Errno is the Windows error number.
type Errno uintptr

func (e Errno) Error() string

func (e Errno) Temporary() bool

func (e Errno) Timeout() bool

// Converts a Go function to a function pointer conforming
// to the stdcall calling convention.  This is useful when
// interoperating with Windows code requiring callbacks.
// Implemented in ../runtime/windows/syscall.goc
func NewCallback(fn interface{}) uintptr

func Exit(code int)

func Open(path string, mode int, perm uint32) (fd Handle, err error)

func Read(fd Handle, p []byte) (n int, err error)

func Write(fd Handle, p []byte) (n int, err error)

func Seek(fd Handle, offset int64, whence int) (newoffset int64, err error)

func Close(fd Handle) (err error)

var (
	Stdin  = getStdHandle(STD_INPUT_HANDLE)
	Stdout = getStdHandle(STD_OUTPUT_HANDLE)
	Stderr = getStdHandle(STD_ERROR_HANDLE)
)

const ImplementsGetwd = true

func Getwd() (wd string, err error)

func Chdir(path string) (err error)

func Mkdir(path string, mode uint32) (err error)

func Rmdir(path string) (err error)

func Unlink(path string) (err error)

func Rename(oldpath, newpath string) (err error)

func ComputerName() (name string, err error)

func Ftruncate(fd Handle, length int64) (err error)

func Gettimeofday(tv *Timeval) (err error)

func Pipe(p []Handle) (err error)

func Utimes(path string, tv []Timeval) (err error)

func Fsync(fd Handle) (err error)

func Chmod(path string, mode uint32) (err error)

// For testing: clients can set this flag to force
// creation of IPv6 sockets to return EAFNOSUPPORT.
var SocketDisableIPv6 bool

type RawSockaddrInet4 struct {
	Family uint16
	Port   uint16
	Addr   [4]byte
	Zero   [8]uint8
}

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [96]int8
}

type Sockaddr interface {
	sockaddr() (ptr uintptr, len int32, err error)
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
}

type SockaddrUnix struct {
	Name string
}

func (rsa *RawSockaddrAny) Sockaddr() (Sockaddr, error)

func Socket(domain, typ, proto int) (fd Handle, err error)

func SetsockoptInt(fd Handle, level, opt int, value int) (err error)

func Bind(fd Handle, sa Sockaddr) (err error)

func Connect(fd Handle, sa Sockaddr) (err error)

func Getsockname(fd Handle) (sa Sockaddr, err error)

func Getpeername(fd Handle) (sa Sockaddr, err error)

func Listen(s Handle, n int) (err error)

func Shutdown(fd Handle, how int) (err error)

func WSASendto(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to Sockaddr, overlapped *Overlapped, croutine *byte) (err error)

// Invented structures to support what package os expects.
type Rusage struct {
	CreationTime Filetime
	ExitTime     Filetime
	KernelTime   Filetime
	UserTime     Filetime
}

type WaitStatus struct {
	ExitCode uint32
}

func (w WaitStatus) Exited() bool

func (w WaitStatus) ExitStatus() int

func (w WaitStatus) Signal() Signal

func (w WaitStatus) CoreDump() bool

func (w WaitStatus) Stopped() bool

func (w WaitStatus) Continued() bool

func (w WaitStatus) StopSignal() Signal

func (w WaitStatus) Signaled() bool

func (w WaitStatus) TrapCause() int

// Timespec is an invented structure on Windows, but here for
// consistency with the syscall package for other operating systems.
type Timespec struct {
	Sec  int64
	Nsec int64
}

func Accept(fd Handle) (nfd Handle, sa Sockaddr, err error)
func Recvfrom(fd Handle, p []byte, flags int) (n int, from Sockaddr, err error)

func Sendto(fd Handle, p []byte, flags int, to Sockaddr) (err error)
func SetsockoptTimeval(fd Handle, level, opt int, tv *Timeval) (err error)

type Linger struct {
	Onoff  int32
	Linger int32
}

type IPMreq struct {
	Multiaddr [4]byte
	Interface [4]byte
}

type IPv6Mreq struct {
	Multiaddr [16]byte
	Interface uint32
}

func GetsockoptInt(fd Handle, level, opt int) (int, error)
func SetsockoptLinger(fd Handle, level, opt int, l *Linger) (err error)
func SetsockoptInet4Addr(fd Handle, level, opt int, value [4]byte) (err error)

func SetsockoptIPMreq(fd Handle, level, opt int, mreq *IPMreq) (err error)

func SetsockoptIPv6Mreq(fd Handle, level, opt int, mreq *IPv6Mreq) (err error)

func Getpid() (pid int)

func FindFirstFile(name *uint16, data *Win32finddata) (handle Handle, err error)

func FindNextFile(handle Handle, data *Win32finddata) (err error)

// TODO(brainman): fix all needed for os
func Getppid() (ppid int)

func Fchdir(fd Handle) (err error)
func Link(oldpath, newpath string) (err error)
func Symlink(path, link string) (err error)
func Readlink(path string, buf []byte) (n int, err error)

func Fchmod(fd Handle, mode uint32) (err error)
func Chown(path string, uid int, gid int) (err error)
func Lchown(path string, uid int, gid int) (err error)
func Fchown(fd Handle, uid int, gid int) (err error)

func Getuid() (uid int)
func Geteuid() (euid int)
func Getgid() (gid int)
func Getegid() (egid int)
func Getgroups() (gids []int, err error)

type Signal int

func (s Signal) Signal()

func (s Signal) String() string
