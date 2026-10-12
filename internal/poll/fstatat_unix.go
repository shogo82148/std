// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || wasip1

package poll

import (
	"github.com/shogo82148/std/syscall"
)

func (fd *FD) Fstatat(name string, s *syscall.Stat_t, flags int) error
