// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import (
	"github.com/shogo82148/std/syscall"
)

// SendFile wraps the TransmitFile call.
func SendFile(fd *FD, src syscall.Handle, n int64) (written int64, err error)
