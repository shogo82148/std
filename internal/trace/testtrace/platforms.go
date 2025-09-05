// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testtrace

import (
	"github.com/shogo82148/std/testing"
)

// MustHaveSyscallEvents skips the current test if the current
// platform does not support true system call events.
func MustHaveSyscallEvents(t *testing.T)

// HasSyscallEvents returns true if the current platform
// has real syscall events available.
func HasSyscallEvents() bool
