// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Plan 9 system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and
// wrap it in our own nicer implementation.

package syscall

const ImplementsGetwd = true

// ErrnoはGOOS plan9では使用されません。
//
// これは、Go 1のAPI契約（api/go1.txtの"pkg syscall, type Errno uintptr"）で
// 存在すると規定されており、エコシステム内のコードはしばしばその存在を
// 前提としているために存在します。これにより、ビルドタグ付きのファイルを使わずに、
// 代わりにruntime.GOOSのチェックを使用して移植可能なコードを書くことができます。
type Errno uintptr

func (e Errno) Error() string
func (e Errno) Temporary() bool
func (e Errno) Timeout() bool

// ErrorStringは、自分自身を返すことでErrorのStringメソッドを実装します。
//
// ErrorStringの値は、[errors.Is] を使用してエラー値と比較できます。
// 例えば:
//
//	_, _, err := syscall.Syscall(...)
//	if errors.Is(err, fs.ErrNotExist) ...
type ErrorString string

func (e ErrorString) Error() string

// NewErrorは、sをErrorインターフェースを満たすErrorStringに変換します。
func NewError(s string) error

func (e ErrorString) Is(target error) bool

func (e ErrorString) Temporary() bool

func (e ErrorString) Timeout() bool

// Noteは、プロセスのノートを記述する文字列です。
// これはos.Signalインターフェースを実装します。
type Note string

func (n Note) Signal()

func (n Note) String() string

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

// テスト用: クライアントはこのフラグを設定することで、
// IPv6ソケットの作成で強制的に [EAFNOSUPPORT] を返すようにできます。
var SocketDisableIPv6 bool

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err ErrorString)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err ErrorString)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)

func Getpid() (pid int)

func Getppid() (ppid int)

func Read(fd int, p []byte) (n int, err error)

func Write(fd int, p []byte) (n int, err error)

func Fd2path(fd int) (path string, err error)

func Pipe(p []int) (err error)

func Seek(fd int, offset int64, whence int) (newoffset int64, err error)

func Mkdir(path string, mode uint32) (err error)

type Waitmsg struct {
	Pid  int
	Time [3]uint32
	Msg  string
}

func (w Waitmsg) Exited() bool
func (w Waitmsg) Signaled() bool

func (w Waitmsg) ExitStatus() int

func Await(w *Waitmsg) (err error)

func Unmount(name, old string) (err error)

func Fchdir(fd int) (err error)

type Timespec struct {
	Sec  int32
	Nsec int32
}

type Timeval struct {
	Sec  int32
	Usec int32
}

func NsecToTimeval(nsec int64) (tv Timeval)

func Gettimeofday(tv *Timeval) error

func Getegid() (egid int)
func Geteuid() (euid int)
func Getgid() (gid int)
func Getuid() (uid int)

func Getgroups() (gids []int, err error)

func Open(path string, mode int) (fd int, err error)

func Create(path string, mode int, perm uint32) (fd int, err error)

func Remove(path string) error

func Stat(path string, edir []byte) (n int, err error)

func Bind(name string, old string, flag int) (err error)

func Mount(fd int, afd int, old string, flag int, aname string) (err error)

func Wstat(path string, edir []byte) (err error)
