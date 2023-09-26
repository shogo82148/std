// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (mips || mipsle)
// +build linux
// +build mips mipsle

package syscall

// archHonorsR2 captures the fact that r2 is honored by the
// runtime.GOARCH.  Syscall conventions are generally r1, r2, err :=
// syscall(trap, ...).  Not all architectures define r2 in their
// ABI. See "man syscall".

func Syscall9(trap, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)

func Fstatfs(fd int, buf *Statfs_t) (err error)

func Statfs(path string, buf *Statfs_t) (err error)

func Seek(fd int, offset int64, whence int) (off int64, err error)

func Pipe2(p []int, flags int) (err error)

func Pipe(p []int) (err error)

func Getrlimit(resource int, rlim *Rlimit) (err error)

func Setrlimit(resource int, rlim *Rlimit) (err error)

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)
