// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

package syscall

var (
	Stdin  = 0
	Stdout = 1
	Stderr = 2
)

func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// An Errno is an unsigned number describing an error condition.
// It implements the error interface.  The zero Errno is by convention
// a non-error, so code to convert from Errno to error should use:
//
//	err = nil
//	if errno != 0 {
//		err = errno
//	}
type Errno uintptr

func (e Errno) Error() string

func (e Errno) Temporary() bool

func (e Errno) Timeout() bool

// A Signal is a number describing a process signal.
// It implements the os.Signal interface.
type Signal int

func (s Signal) Signal()

func (s Signal) String() string

func Read(fd int, p []byte) (n int, err error)

func Write(fd int, p []byte) (n int, err error)

func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
