// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux
// +build linux

package syscall

// SysProcIDMap holds Container ID to Host ID mappings used for User Namespaces in Linux.
// See user_namespaces(7).
type SysProcIDMap struct {
	ContainerID int
	HostID      int
	Size        int
}

type SysProcAttr struct {
	Chroot       string
	Credential   *Credential
	Ptrace       bool
	Setsid       bool
	Setpgid      bool
	Setctty      bool
	Noctty       bool
	Ctty         int
	Foreground   bool
	Pgid         int
	Pdeathsig    Signal
	Cloneflags   uintptr
	Unshareflags uintptr
	UidMappings  []SysProcIDMap
	GidMappings  []SysProcIDMap

	GidMappingsEnableSetgroups bool
	AmbientCaps                []uintptr
}
