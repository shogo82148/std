// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd linux netbsd openbsd

package net

import (
	"github.com/shogo82148/std/os"
)

// FileConn returns a copy of the network connection corresponding to
// the open file f.  It is the caller's responsibility to close f when
// finished.  Closing c does not affect f, and closing f does not
// affect c.
func FileConn(f *os.File) (c Conn, err error)

// FileListener returns a copy of the network listener corresponding
// to the open file f.  It is the caller's responsibility to close l
// when finished.  Closing l does not affect f, and closing f does not
// affect l.
func FileListener(f *os.File) (l Listener, err error)

// FilePacketConn returns a copy of the packet network connection
// corresponding to the open file f.  It is the caller's
// responsibility to close f when finished.  Closing c does not affect
// f, and closing f does not affect c.
func FilePacketConn(f *os.File) (c PacketConn, err error)
