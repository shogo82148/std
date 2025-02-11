// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package unix

import (
	"github.com/shogo82148/std/syscall"
)

// The values of these constants are not part of the WASI API.
const (
	// UTIME_OMIT is the sentinel value to indicate that a time value should not
	// be changed. It is useful for example to indicate for example with UtimesNano
	// to avoid changing AccessTime or ModifiedTime.
	// Its value must match syscall/fs_wasip1.go
	UTIME_OMIT = -0x2

	AT_REMOVEDIR        = 0x200
	AT_SYMLINK_NOFOLLOW = 0x100
)

func Unlinkat(dirfd int, path string, flags int) error

func Openat(dirfd int, path string, flags int, perm uint32) (int, error)

func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error

func Readlinkat(dirfd int, path string, buf []byte) (int, error)

func Mkdirat(dirfd int, path string, mode uint32) error

func Fchmodat(dirfd int, path string, mode uint32, flags int) error
