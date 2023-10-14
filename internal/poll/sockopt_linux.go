// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package poll

import "github.com/shogo82148/std/syscall"

// SetsockoptIPMreqn wraps the setsockopt network call with an IPMreqn argument.
func (fd *FD) SetsockoptIPMreqn(level, name int, mreq *syscall.IPMreqn) error
