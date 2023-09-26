// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains main runtime AIX syscalls.
// Pollset syscalls are in netpoll_aix.go.
// The implementation is based on Solaris and Windows.
// Each syscall is made by calling its libc symbol using asmcgocall and asmsyscall6
// assembly functions.

package runtime

// asmsyscall6 calls the libc symbol using a C convention.
// It's defined in sys_aix_ppc64.go.
