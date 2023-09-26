// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || solaris
// +build aix solaris

// This file handles forkAndExecInChild function for OS using libc syscall like AIX or Solaris.

package syscall

type SysProcAttr struct {
	Chroot     string
	Credential *Credential
	Setsid     bool

	Setpgid bool

	Setctty bool
	Noctty  bool
	Ctty    int

	Foreground bool
	Pgid       int
}
