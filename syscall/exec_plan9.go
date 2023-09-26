// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fork, exec, wait, etc.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

var ForkLock sync.RWMutex

// Convert array of string to array
// of NUL-terminated byte pointer.
func StringSlicePtr(ss []string) []*byte

type ProcAttr struct {
	Dir   string
	Env   []string
	Files []uintptr
	Sys   *SysProcAttr
}

type SysProcAttr struct {
	Rfork int
}

// Combination of fork and exec, careful to be thread safe.
func ForkExec(argv0 string, argv []string, attr *ProcAttr) (pid int, err error)

// StartProcess wraps ForkExec for package os.
func StartProcess(argv0 string, argv []string, attr *ProcAttr) (pid int, handle uintptr, err error)

// Ordinary exec.
func Exec(argv0 string, argv []string, envv []string) (err error)
