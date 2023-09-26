// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/syscall"
)

// Process stores the information about a process created by StartProcess.
type Process struct {
	Pid    int
	handle uintptr
	isdone uint32
}

// ProcAttr holds the attributes that will be applied to a new process
// started by StartProcess.
type ProcAttr struct {
	Dir string

	Env []string

	Files []*File

	Sys *syscall.SysProcAttr
}

// A Signal represents an operating system signal.
// The usual underlying implementation is operating system-dependent:
// on Unix it is syscall.Signal.
type Signal interface {
	String() string
	Signal()
}

// Getpid returns the process id of the caller.
func Getpid() int

// Getppid returns the process id of the caller's parent.
func Getppid() int
