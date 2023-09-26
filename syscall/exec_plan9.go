// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fork, exec, wait, etc.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

var ForkLock sync.RWMutex

// StringSlicePtr is deprecated. Use SlicePtrFromStrings instead.
// If any string contains a NUL byte this function panics instead
// of returning an error.
func StringSlicePtr(ss []string) []*byte

// SlicePtrFromStrings converts a slice of strings to a slice of
// pointers to NUL-terminated byte slices. If any string contains
// a NUL byte, it returns (nil, EINVAL).
func SlicePtrFromStrings(ss []string) ([]*byte, error)

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

// WaitProcess waits until the pid of a
// running process is found in the queue of
// wait messages. It is used in conjunction
// with ForkExec/StartProcess to wait for a
// running process to exit.
func WaitProcess(pid int, w *Waitmsg) (err error)
