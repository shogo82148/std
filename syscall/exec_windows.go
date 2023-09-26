// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fork, exec, wait, etc.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

var ForkLock sync.RWMutex

// EscapeArg rewrites command line argument s as prescribed
// in https://msdn.microsoft.com/en-us/library/ms880421.
// This function returns "" (2 double quotes) if s is empty.
// Alternatively, these transformations are done:
//   - every back slash (\) is doubled, but only if immediately
//     followed by double quote (");
//   - every double quote (") is escaped by back slash (\);
//   - finally, s is wrapped with double quotes (arg -> "arg"),
//     but only if there is space or tab inside s.
func EscapeArg(s string) string

func CloseOnExec(fd Handle)

func SetNonblock(fd Handle, nonblocking bool) (err error)

// FullPath retrieves the full path of the specified file.
func FullPath(name string) (path string, err error)

type ProcAttr struct {
	Dir   string
	Env   []string
	Files []uintptr
	Sys   *SysProcAttr
}

type SysProcAttr struct {
	HideWindow        bool
	CmdLine           string
	CreationFlags     uint32
	Token             Token
	ProcessAttributes *SecurityAttributes
	ThreadAttributes  *SecurityAttributes
}

func StartProcess(argv0 string, argv []string, attr *ProcAttr) (pid int, handle uintptr, err error)

func Exec(argv0 string, argv []string, envv []string) (err error)
