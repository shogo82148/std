// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package socktest

import "github.com/shogo82148/std/syscall"

// Socket wraps syscall.Socket.
func (sw *Switch) Socket(family, sotype, proto int) (s int, err error)

// Close wraps syscall.Close.
func (sw *Switch) Close(s int) (err error)

// Connect wraps syscall.Connect.
func (sw *Switch) Connect(s int, sa syscall.Sockaddr) (err error)

// Listen wraps syscall.Listen.
func (sw *Switch) Listen(s, backlog int) (err error)

// Accept wraps syscall.Accept.
func (sw *Switch) Accept(s int) (ns int, sa syscall.Sockaddr, err error)

// GetsockoptInt wraps syscall.GetsockoptInt.
func (sw *Switch) GetsockoptInt(s, level, opt int) (soerr int, err error)
