// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func Getpagesize() int

func TimespecToNsec(ts Timespec) int64

func NsecToTimespec(nsec int64) (ts Timespec)

func TimevalToNsec(tv Timeval) int64

func NsecToTimeval(nsec int64) (tv Timeval)

// sysnb	gettimeofday(tp *Timeval) (sec int32, usec int32, err error)
func Gettimeofday(tv *Timeval) (err error)

func SetKevent(k *Kevent_t, fd, mode, flags int)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)

func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)
