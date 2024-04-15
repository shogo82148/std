// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/shogo82148/std/runtime"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/syscall"
)

// KernelVersion returns major and minor kernel version numbers
// parsed from the syscall.Uname's Version field, or (0, 0) if the
// version can't be obtained or parsed.
func KernelVersion() (major int, minor int)

// SupportSockNonblockCloexec tests if SOCK_NONBLOCK and SOCK_CLOEXEC are supported
// for socket() system call, returns true if affirmative.
var SupportSockNonblockCloexec = sync.OnceValue(func() bool {

	s, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM|syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC, 0)
	if err == nil {
		syscall.Close(s)
		return true
	}
	if err != syscall.EPROTONOSUPPORT && err != syscall.EINVAL {

		major, minor := KernelVersion()
		if runtime.GOOS == "illumos" {
			return major > 5 || (major == 5 && minor >= 11)
		}
		return major > 11 || (major == 11 && minor >= 4)
	}
	return false
})

// SupportAccept4 tests whether accept4 system call is available.
var SupportAccept4 = sync.OnceValue(func() bool {
	for {

		_, _, err := syscall.Accept4(0, syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC)
		if err == syscall.EINTR {
			continue
		}
		return err != syscall.ENOSYS
	}
})
