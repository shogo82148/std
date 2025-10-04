// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/shogo82148/std/syscall"
)

// IsNonblock returns whether the file descriptor fd is opened
// in non-blocking mode, that is, the [syscall.FILE_FLAG_OVERLAPPED] flag
// was set when the file was opened.
func IsNonblock(fd syscall.Handle) (nonblocking bool, err error)
