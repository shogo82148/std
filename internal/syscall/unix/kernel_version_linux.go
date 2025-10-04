// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

// KernelVersion returns major and minor kernel version numbers
// parsed from the syscall.Uname's Release field, or (0, 0) if
// the version can't be obtained or parsed.
func KernelVersion() (major, minor int)
