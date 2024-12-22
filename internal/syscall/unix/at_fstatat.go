// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || (linux && !(loong64 || mips64 || mips64le)) || netbsd || (openbsd && mips64)

package unix

import (
	"github.com/shogo82148/std/syscall"
)

func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error
