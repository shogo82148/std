// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/shogo82148/std/syscall"
)

// Openat flags not supported by syscall.Open.
//
// These are invented values.
//
// When adding a new flag here, add an unexported version to
// the set of invented O_ values in syscall/types_windows.go
// to avoid overlap.
const (
	O_DIRECTORY    = 0x100000
	O_NOFOLLOW_ANY = 0x20000000
	O_OPEN_REPARSE = 0x40000000
)

func Openat(dirfd syscall.Handle, name string, flag int, perm uint32) (_ syscall.Handle, e1 error)

func Mkdirat(dirfd syscall.Handle, name string, mode uint32) error

func Deleteat(dirfd syscall.Handle, name string) error
