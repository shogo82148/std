// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/shogo82148/std/syscall"
)

func Getgrouplist(name *byte, gid uint32, gids *uint32, n *int32) error

const (
	SC_GETGR_R_SIZE_MAX = 0x46
	SC_GETPW_R_SIZE_MAX = 0x47
)

type Passwd struct {
	Name   *byte
	Passwd *byte
	Uid    uint32
	Gid    uint32
	Change int64
	Class  *byte
	Gecos  *byte
	Dir    *byte
	Shell  *byte
	Expire int64
}

type Group struct {
	Name   *byte
	Passwd *byte
	Gid    uint32
	Mem    **byte
}

func Getpwnam(name *byte, pwd *Passwd, buf *byte, size uintptr, result **Passwd) syscall.Errno

func Getpwuid(uid uint32, pwd *Passwd, buf *byte, size uintptr, result **Passwd) syscall.Errno

func Getgrnam(name *byte, grp *Group, buf *byte, size uintptr, result **Group) syscall.Errno

func Getgrgid(gid uint32, grp *Group, buf *byte, size uintptr, result **Group) syscall.Errno

func Sysconf(key int32) int64
