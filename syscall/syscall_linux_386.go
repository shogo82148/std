// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func Getpagesize() int

func TimespecToNsec(ts Timespec) int64

func NsecToTimespec(nsec int64) (ts Timespec)

func TimevalToNsec(tv Timeval) int64

func NsecToTimeval(nsec int64) (tv Timeval)

// Underlying system call writes to newoffset via pointer.
// Implemented in assembly to avoid allocation.
func Seek(fd int, offset int64, whence int) (newoffset int64, err error)

func Listen(s int, n int) (err error)

func Shutdown(s, how int) (err error)

func Fstatfs(fd int, buf *Statfs_t) (err error)

func Statfs(path string, buf *Statfs_t) (err error)

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)
