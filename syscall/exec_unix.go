// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd
// +build darwin freebsd linux netbsd openbsd

// Fork, exec, wait, etc.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

var ForkLock sync.RWMutex

// Convert array of string to array of NUL-terminated byte pointer.
// If any string contains a NUL byte this function panics instead
// of returning an error.
func StringSlicePtr(ss []string) []*byte

func CloseOnExec(fd int)

func SetNonblock(fd int, nonblocking bool) (err error)

// Credential holds user and group identities to be assumed
// by a child process started by StartProcess.
type Credential struct {
	Uid    uint32
	Gid    uint32
	Groups []uint32
}

// ProcAttr holds attributes that will be applied to a new process started
// by StartProcess.
type ProcAttr struct {
	Dir   string
	Env   []string
	Files []uintptr
	Sys   *SysProcAttr
}

// Combination of fork and exec, careful to be thread safe.
func ForkExec(argv0 string, argv []string, attr *ProcAttr) (pid int, err error)

// StartProcess wraps ForkExec for package os.
func StartProcess(argv0 string, argv []string, attr *ProcAttr) (pid int, handle uintptr, err error)

// Ordinary exec.
func Exec(argv0 string, argv []string, envv []string) (err error)
