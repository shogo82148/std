// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func Time(t *Time_t) (tt Time_t, err error)

func Listen(s int, n int) (err error)

func Shutdown(s, how int) (err error)

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)
