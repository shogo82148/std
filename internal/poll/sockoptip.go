// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

package poll

import "github.com/shogo82148/std/syscall"

// SetsockoptIPMreq wraps the setsockopt network call with an IPMreq argument.
func (fd *FD) SetsockoptIPMreq(level, name int, mreq *syscall.IPMreq) error

// SetsockoptIPv6Mreq wraps the setsockopt network call with an IPv6Mreq argument.
func (fd *FD) SetsockoptIPv6Mreq(level, name int, mreq *syscall.IPv6Mreq) error
