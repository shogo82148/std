// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

type SysProcAttr struct {
	Chroot     string
	Credential *Credential
	Ptrace     bool
	Setsid     bool

	Setpgid bool

	Setctty bool
	Noctty  bool
	Ctty    int

	Foreground bool
	Pgid       int
}
