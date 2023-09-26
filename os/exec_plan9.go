// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/syscall"
)

// The only signal values guaranteed to be present in the os package
// on all systems are Interrupt (send the process an interrupt) and
// Kill (force the process to exit). Interrupt is not implemented on
// Windows; using it with os.Process.Signal will return an error.
var (
	Interrupt Signal = syscall.Note("interrupt")
	Kill      Signal = syscall.Note("kill")
)

// ProcessState stores information about a process, as reported by Wait.
type ProcessState struct {
	pid    int
	status *syscall.Waitmsg
}

// Pid returns the process id of the exited process.
func (p *ProcessState) Pid() int

func (p *ProcessState) String() string

// ExitCode returns the exit code of the exited process, or -1
// if the process hasn't exited or was terminated by a signal.
func (p *ProcessState) ExitCode() int
