// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/shogo82148/std/syscall"
)

const (
	P_PID   = 1
	P_PIDFD = 3
)

func Waitid(idType int, id int, info *SiginfoChild, options int, rusage *syscall.Rusage) error
