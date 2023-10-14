// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package syscall provides the syscall primitives required for the runtime.
package syscall

// Syscall6 calls system call number 'num' with arguments a1-6.
func Syscall6(num, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, errno uintptr)

func EpollCreate1(flags int32) (fd int32, errno uintptr)

func EpollWait(epfd int32, events []EpollEvent, maxev, waitms int32) (n int32, errno uintptr)

func EpollCtl(epfd, op, fd int32, event *EpollEvent) (errno uintptr)
