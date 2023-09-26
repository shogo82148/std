// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// sigPerThreadSyscall is the same signal (SIGSETXID) used by glibc for
// per-thread syscalls on Linux. We use it for the same purpose in non-cgo
// binaries.

// Clone, the Linux rfork.

// startupRandomData holds random bytes initialized at startup. These come from
// the ELF AT_RANDOM auxiliary vector.

// secureMode holds the value of AT_SECURE passed in the auxiliary vector.

// perThreadSyscallArgs contains the system call number, arguments, and
// expected return values for a system call to be executed on all threads.

// perThreadSyscall is the system call to execute for the ongoing
// doAllThreadsSyscall.
//
// perThreadSyscall may only be written while mp.needPerThreadSyscall == 0 on
// all Ms.
