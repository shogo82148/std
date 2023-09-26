// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func EpollCreate(size int) (fd int, err error)

func Stat(path string, stat *Stat_t) (err error)

func Lchown(path string, uid int, gid int) (err error)

func Lstat(path string, stat *Stat_t) (err error)

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error)

func Time(t *Time_t) (Time_t, error)

func Utime(path string, buf *Utimbuf) error

// Getrlimit prefers the prlimit64 system call. See issue 38604.
func Getrlimit(resource int, rlim *Rlimit) error

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)

func InotifyInit() (fd int, err error)

func Pause() error
