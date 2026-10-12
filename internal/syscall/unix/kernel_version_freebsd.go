// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/syscall"
)

// KernelVersion returns major and minor kernel version numbers
// parsed from the syscall.Sysctl("kern.osrelease")'s value,
// or (0, 0) if the version can't be obtained or parsed.
func KernelVersion() (major, minor int)

// SupportCopyFileRange reports whether the kernel supports the copy_file_range(2).
// This function will examine both the kernel version and the availability of the system call.
var SupportCopyFileRange = sync.OnceValue(func() bool {

	if !KernelVersionGE(13, 0) {
		return false
	}
	_, err := CopyFileRange(0, nil, 0, nil, 0, 0)
	return err != syscall.ENOSYS
})
