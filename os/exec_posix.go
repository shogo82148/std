// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || windows
// +build darwin dragonfly freebsd linux netbsd openbsd windows

package os

import (
	"github.com/shogo82148/std/syscall"
)

// The only signal values guaranteed to be present on all systems
// are Interrupt (send the process an interrupt) and Kill (force
// the process to exit).
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)

// ProcessState stores information about a process, as reported by Wait.
type ProcessState struct {
	pid    int
	status syscall.WaitStatus
	rusage *syscall.Rusage
}

// Pid returns the process id of the exited process.
func (p *ProcessState) Pid() int

func (p *ProcessState) String() string
