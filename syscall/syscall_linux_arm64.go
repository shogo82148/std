// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func Stat(path string, stat *Stat_t) (err error)

func Lchown(path string, uid int, gid int) (err error)

func Lstat(path string, stat *Stat_t) (err error)

func Getpagesize() int

func TimespecToNsec(ts Timespec) int64

func NsecToTimespec(nsec int64) (ts Timespec)

func TimevalToNsec(tv Timeval) int64

func NsecToTimeval(nsec int64) (tv Timeval)

func Pipe(p []int) (err error)

func Pipe2(p []int, flags int) (err error)

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)

func InotifyInit() (fd int, err error)

// TODO(dfc): constants that should be in zsysnum_linux_arm64.go, remove
// these when the deprecated syscalls that the syscall package relies on
// are removed.
const (
	SYS_GETPGRP      = 1060
	SYS_UTIMES       = 1037
	SYS_FUTIMESAT    = 1066
	SYS_PAUSE        = 1061
	SYS_USTAT        = 1070
	SYS_UTIME        = 1063
	SYS_LCHOWN       = 1032
	SYS_TIME         = 1062
	SYS_EPOLL_CREATE = 1042
	SYS_EPOLL_WAIT   = 1069
)
