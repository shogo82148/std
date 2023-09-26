// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

// Fork, exec, wait, etc.

package syscall

import (
	errorspkg "errors"
	"github.com/shogo82148/std/sync"
)

var ForkLock sync.RWMutex

// StringSlicePtr converts a slice of strings to a slice of pointers
// to NUL-terminated byte arrays. If any string contains a NUL byte
// this function panics instead of returning an error.
//
// Deprecated: Use SlicePtrFromStrings instead.
func StringSlicePtr(ss []string) []*byte

// SlicePtrFromStrings converts a slice of strings to a slice of
// pointers to NUL-terminated byte arrays. If any string contains
// a NUL byte, it returns (nil, EINVAL).
func SlicePtrFromStrings(ss []string) ([]*byte, error)

func CloseOnExec(fd int)

func SetNonblock(fd int, nonblocking bool) (err error)

// Credential holds user and group identities to be assumed
// by a child process started by StartProcess.
type Credential struct {
	Uid         uint32
	Gid         uint32
	Groups      []uint32
	NoSetGroups bool
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

// execveLibc is non-nil on OS using libc syscall, set to execve in exec_libc.go; this
// avoids a build dependency for other platforms.

// Exec invokes the execve(2) system call.
func Exec(argv0 string, argv []string, envv []string) (err error)
