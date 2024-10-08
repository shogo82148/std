// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "github.com/shogo82148/std/syscall"

func PidFDSendSignal(pidfd uintptr, s syscall.Signal) error

func PidFDOpen(pid, flags int) (uintptr, error)
