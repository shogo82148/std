// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socktest

import (
	"github.com/shogo82148/std/syscall"
)

// WSASocket wraps [syscall.WSASocket].
func (sw *Switch) WSASocket(family, sotype, proto int32, protinfo *syscall.WSAProtocolInfo, group uint32, flags uint32) (s syscall.Handle, err error)

// Closesocket wraps [syscall.Closesocket].
func (sw *Switch) Closesocket(s syscall.Handle) (err error)

// Connect wraps [syscall.Connect].
func (sw *Switch) Connect(s syscall.Handle, sa syscall.Sockaddr) (err error)

// ConnectEx wraps [syscall.ConnectEx].
func (sw *Switch) ConnectEx(s syscall.Handle, sa syscall.Sockaddr, b *byte, n uint32, nwr *uint32, o *syscall.Overlapped) (err error)

// Listen wraps [syscall.Listen].
func (sw *Switch) Listen(s syscall.Handle, backlog int) (err error)

// AcceptEx wraps [syscall.AcceptEx].
func (sw *Switch) AcceptEx(ls syscall.Handle, as syscall.Handle, b *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, rcvd *uint32, overlapped *syscall.Overlapped) error
