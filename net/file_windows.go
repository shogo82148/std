// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/os"
)

func FileConn(f *os.File) (c Conn, err error)

func FileListener(f *os.File) (l Listener, err error)

func FilePacketConn(f *os.File) (c PacketConn, err error)
