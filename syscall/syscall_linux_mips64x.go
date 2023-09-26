// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (mips64 || mips64le)
// +build linux
// +build mips64 mips64le

package syscall

func Getpagesize() int

func Time(t *Time_t) (tt Time_t, err error)

func TimespecToNsec(ts Timespec) int64

func NsecToTimespec(nsec int64) (ts Timespec)

func TimevalToNsec(tv Timeval) int64

func NsecToTimeval(nsec int64) (tv Timeval)

func Pipe(p []int) (err error)

func Pipe2(p []int, flags int) (err error)

func Ioperm(from int, num int, on int) (err error)

func Iopl(level int) (err error)

func Fstat(fd int, s *Stat_t) (err error)

func Lstat(path string, s *Stat_t) (err error)

func Stat(path string, s *Stat_t) (err error)

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)
