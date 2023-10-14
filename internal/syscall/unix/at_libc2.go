// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || (openbsd && !mips64)

package unix

import (
	"github.com/shogo82148/std/syscall"
)

func Unlinkat(dirfd int, path string, flags int) error

func Openat(dirfd int, path string, flags int, perm uint32) (int, error)

func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error
