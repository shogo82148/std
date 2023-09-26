// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (ppc64 || ppc64le)
// +build linux
// +build ppc64 ppc64le

package syscall

// archHonorsR2 captures the fact that r2 is honored by the
// runtime.GOARCH.  Syscall conventions are generally r1, r2, err :=
// syscall(trap, ...).  Not all architectures define r2 in their
// ABI. See "man syscall".

func Pipe(p []int) (err error)

func Pipe2(p []int, flags int) (err error)

func (r *PtraceRegs) PC() uint64

func (r *PtraceRegs) SetPC(pc uint64)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)

func SyncFileRange(fd int, off int64, n int64, flags int) error
