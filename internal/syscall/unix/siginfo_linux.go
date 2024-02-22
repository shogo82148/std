// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/shogo82148/std/syscall"
)

// SiginfoChild is a struct filled in by Linux waitid syscall.
// In C, siginfo_t contains a union with multiple members;
// this struct corresponds to one used when Signo is SIGCHLD.
//
// NOTE fields are exported to be used by TestSiginfoChildLayout.
type SiginfoChild struct {
	Signo int32
	siErrnoCode
	_ [is64bit]int32

	Pid    int32
	Uid    uint32
	Status int32

	// Pad to 128 bytes.
	_ [128 - (6+is64bit)*4]byte
}

// WaitStatus converts SiginfoChild, as filled in by the waitid syscall,
// to syscall.WaitStatus.
func (s *SiginfoChild) WaitStatus() (ws syscall.WaitStatus)
