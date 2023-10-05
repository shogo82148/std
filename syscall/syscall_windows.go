// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Windows system calls.

package syscall

import (
	"github.com/shogo82148/std/unsafe"
)

type Handle uintptr

const InvalidHandle = ^Handle(0)

// StringToUTF16 returns the UTF-16 encoding of the UTF-8 string s,
// with a terminating NUL added. If s contains a NUL byte this
// function panics instead of returning an error.
//
// Deprecated: Use UTF16FromString instead.
func StringToUTF16(s string) []uint16

// UTF16FromString returns the UTF-16 encoding of the UTF-8 string
// s, with a terminating NUL added. If s contains a NUL byte at any
// location, it returns (nil, EINVAL). Unpaired surrogates
// are encoded using WTF-8.
func UTF16FromString(s string) ([]uint16, error)

// UTF16ToString returns the UTF-8 encoding of the UTF-16 sequence s,
// with a terminating NUL removed. Unpaired surrogates are decoded
// using WTF-8 instead of UTF-8 encoding.
func UTF16ToString(s []uint16) string

// StringToUTF16Ptr returns pointer to the UTF-16 encoding of
// the UTF-8 string s, with a terminating NUL added. If s
// contains a NUL byte this function panics instead of
// returning an error.
//
// Deprecated: Use UTF16PtrFromString instead.
func StringToUTF16Ptr(s string) *uint16

// UTF16PtrFromString returns pointer to the UTF-16 encoding of
// the UTF-8 string s, with a terminating NUL added. If s
// contains a NUL byte at any location, it returns (nil, EINVAL).
// Unpaired surrogates are encoded using WTF-8.
func UTF16PtrFromString(s string) (*uint16, error)

// Errno is the Windows error number.
//
// Errno values can be tested against error values using errors.Is.
// For example:
//
//	_, _, err := syscall.Syscall(...)
//	if errors.Is(err, fs.ErrNotExist) ...
type Errno uintptr

// FormatMessage is deprecated (msgsrc should be uintptr, not uint32, but can
// not be changed due to the Go 1 compatibility guarantee).
//
// Deprecated: Use FormatMessage from golang.org/x/sys/windows instead.
func FormatMessage(flags uint32, msgsrc uint32, msgid uint32, langid uint32, buf []uint16, args *byte) (n uint32, err error)

func (e Errno) Error() string

func (e Errno) Is(target error) bool

func (e Errno) Temporary() bool

func (e Errno) Timeout() bool

// NewCallback converts a Go function to a function pointer conforming to the stdcall calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
// Only a limited number of callbacks may be created in a single Go process, and any memory allocated
// for these callbacks is never released.
// Between NewCallback and NewCallbackCDecl, at least 1024 callbacks can always be created.
func NewCallback(fn any) uintptr

// NewCallbackCDecl converts a Go function to a function pointer conforming to the cdecl calling convention.
// This is useful when interoperating with Windows code requiring callbacks.
// The argument is expected to be a function with one uintptr-sized result. The function must not have arguments with size larger than the size of uintptr.
// Only a limited number of callbacks may be created in a single Go process, and any memory allocated
// for these callbacks is never released.
// Between NewCallback and NewCallbackCDecl, at least 1024 callbacks can always be created.
func NewCallbackCDecl(fn any) uintptr

func Open(path string, mode int, perm uint32) (fd Handle, err error)

func Read(fd Handle, p []byte) (n int, err error)

func Write(fd Handle, p []byte) (n int, err error)

func ReadFile(fd Handle, p []byte, done *uint32, overlapped *Overlapped) error

func WriteFile(fd Handle, p []byte, done *uint32, overlapped *Overlapped) error

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

func UtimesNano(path string, ts []Timespec) (err error)

func Fsync(fd Handle) (err error)

func Chmod(path string, mode uint32) (err error)

func LoadCancelIoEx() error

func LoadSetFileCompletionNotificationModes() error

// For testing: clients can set this flag to force
// creation of IPv6 sockets to return EAFNOSUPPORT.
var SocketDisableIPv6 bool

type RawSockaddrInet4 struct {
	Family uint16
	Port   uint16
	Addr   [4]byte
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte
	Scope_id uint32
}

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [100]int8
}

type Sockaddr interface {
	sockaddr() (ptr unsafe.Pointer, len int32, err error)
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

type RawSockaddrUnix struct {
	Family uint16
	Path   [UNIX_PATH_MAX]int8
}

type SockaddrUnix struct {
	Name string
	raw  RawSockaddrUnix
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

func LoadGetAddrInfo() error

func LoadConnectEx() error

func ConnectEx(fd Handle, sa Sockaddr, sendBuf *byte, sendDataLen uint32, bytesSent *uint32, overlapped *Overlapped) error

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

func TimespecToNsec(ts Timespec) int64

func NsecToTimespec(nsec int64) (ts Timespec)

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

func Getppid() (ppid int)

func Fchdir(fd Handle) (err error)

// TODO(brainman): fix all needed for os
func Link(oldpath, newpath string) (err error)
func Symlink(path, link string) (err error)

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

func LoadCreateSymbolicLink() error

// Readlink returns the destination of the named symbolic link.
func Readlink(path string, buf []byte) (n int, err error)

// Deprecated: CreateIoCompletionPort has the wrong function signature. Use x/sys/windows.CreateIoCompletionPort.
func CreateIoCompletionPort(filehandle Handle, cphandle Handle, key uint32, threadcnt uint32) (Handle, error)

// Deprecated: GetQueuedCompletionStatus has the wrong function signature. Use x/sys/windows.GetQueuedCompletionStatus.
func GetQueuedCompletionStatus(cphandle Handle, qty *uint32, key *uint32, overlapped **Overlapped, timeout uint32) error

// Deprecated: PostQueuedCompletionStatus has the wrong function signature. Use x/sys/windows.PostQueuedCompletionStatus.
func PostQueuedCompletionStatus(cphandle Handle, qty uint32, key uint32, overlapped *Overlapped) error

// RegEnumKeyEx enumerates the subkeys of an open registry key.
// Each call retrieves information about one subkey. name is
// a buffer that should be large enough to hold the name of the
// subkey plus a null terminating character. nameLen is its
// length. On return, nameLen will contain the actual length of the
// subkey.
//
// Should name not be large enough to hold the subkey, this function
// will return ERROR_MORE_DATA, and must be called again with an
// appropriately sized buffer.
//
// reserved must be nil. class and classLen behave like name and nameLen
// but for the class of the subkey, except that they are optional.
// lastWriteTime, if not nil, will be populated with the time the subkey
// was last written.
//
// The caller must enumerate all subkeys in order. That is
// RegEnumKeyEx must be called with index starting at 0, incrementing
// the index until the function returns ERROR_NO_MORE_ITEMS, or with
// the index of the last subkey (obtainable from RegQueryInfoKey),
// decrementing until index 0 is enumerated.
//
// Successive calls to this API must happen on the same OS thread,
// so call runtime.LockOSThread before calling this function.
func RegEnumKeyEx(key Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, classLen *uint32, lastWriteTime *Filetime) (regerrno error)

func GetStartupInfo(startupInfo *StartupInfo) error
