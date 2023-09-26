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

// ErrorString implements Error's String method by returning itself.
type ErrorString string

func (e ErrorString) Error() string

// NewError converts s to an ErrorString, which satisfies the Error interface.
func NewError(s string) error

func (e ErrorString) Is(target error) bool

func (e ErrorString) Temporary() bool

func (e ErrorString) Timeout() bool

// A Note is a string describing a process note.
// It implements the os.Signal interface.
type Note string

func (n Note) Signal()

func (n Note) String() string

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

// For testing: clients can set this flag to force
// creation of IPv6 sockets to return EAFNOSUPPORT.
var SocketDisableIPv6 bool

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err ErrorString)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err ErrorString)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr)

func Getpid() (pid int)

func Getppid() (ppid int)

func Read(fd int, p []byte) (n int, err error)

func Write(fd int, p []byte) (n int, err error)

// sys	fd2path(fd int, buf []byte) (err error)
func Fd2path(fd int) (path string, err error)

// sys	pipe(p *[2]int32) (err error)
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

// sys	await(s []byte) (n int, err error)
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

// sys	open(path string, mode int) (fd int, err error)
func Open(path string, mode int) (fd int, err error)

// sys	create(path string, mode int, perm uint32) (fd int, err error)
func Create(path string, mode int, perm uint32) (fd int, err error)

// sys	remove(path string) (err error)
func Remove(path string) error

// sys	stat(path string, edir []byte) (n int, err error)
func Stat(path string, edir []byte) (n int, err error)

// sys	bind(name string, old string, flag int) (err error)
func Bind(name string, old string, flag int) (err error)

// sys	mount(fd int, afd int, old string, flag int, aname string) (err error)
func Mount(fd int, afd int, old string, flag int, aname string) (err error)

// sys	wstat(path string, edir []byte) (err error)
func Wstat(path string, edir []byte) (err error)
