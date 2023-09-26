// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/syscall"
)

// Plan9Note implements the Signal interface on Plan 9.
type Plan9Note string

func (note Plan9Note) String() string

// ProcessState stores information about a process, as reported by Wait.
type ProcessState struct {
	pid    int
	status *syscall.Waitmsg
}

// Pid returns the process id of the exited process.
func (p *ProcessState) Pid() int

func (p *ProcessState) String() string
