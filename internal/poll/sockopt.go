// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || windows

package poll

import "github.com/shogo82148/std/syscall"

// SetsockoptInt wraps the setsockopt network call with an int argument.
func (fd *FD) SetsockoptInt(level, name, arg int) error

// SetsockoptInet4Addr wraps the setsockopt network call with an IPv4 address.
func (fd *FD) SetsockoptInet4Addr(level, name int, arg [4]byte) error

// SetsockoptLinger wraps the setsockopt network call with a Linger argument.
func (fd *FD) SetsockoptLinger(level, name int, l *syscall.Linger) error

// GetsockoptInt wraps the getsockopt network call with an int argument.
func (fd *FD) GetsockoptInt(level, name int) (int, error)
