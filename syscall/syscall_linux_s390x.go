// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// archHonorsR2 captures the fact that r2 is honored by the
// runtime.GOARCH.  Syscall conventions are generally r1, r2, err :=
// syscall(trap, ...).  Not all architectures define r2 in their
// ABI. See "man syscall".

func Time(t *Time_t) (tt Time_t, err error)

func Pipe(p []int) (err error)

func Pipe2(p []int, flags int) (err error)

func Listen(s int, n int) (err error)

func Shutdown(s, how int) (err error)

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)
