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
	O_WRITE_ATTRS  = 0x80000000
)

func Openat(dirfd syscall.Handle, name string, flag uint64, perm uint32) (_ syscall.Handle, e1 error)

func Mkdirat(dirfd syscall.Handle, name string, mode uint32) error

func Deleteat(dirfd syscall.Handle, name string, options uint32) error

func Renameat(olddirfd syscall.Handle, oldpath string, newdirfd syscall.Handle, newpath string) error

func Linkat(olddirfd syscall.Handle, oldpath string, newdirfd syscall.Handle, newpath string) error

// SymlinkatFlags configure Symlinkat.
//
// Symbolic links have two properties: They may be directory or file links,
// and they may be absolute or relative.
//
// The Windows API defines flags describing these properties
// (SYMBOLIC_LINK_FLAG_DIRECTORY and SYMLINK_FLAG_RELATIVE),
// but the flags are passed to different system calls and
// do not have distinct values, so we define our own enumeration
// that permits expressing both.
type SymlinkatFlags uint

const (
	SYMLINKAT_DIRECTORY = SymlinkatFlags(1 << iota)
	SYMLINKAT_RELATIVE
)

func Symlinkat(oldname string, newdirfd syscall.Handle, newname string, flags SymlinkatFlags) error
