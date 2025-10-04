// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package routebsd

// An InterfaceMessage represents an interface message.
type InterfaceMessage struct {
	Version int
	Type    int
	Flags   int
	Index   int
	Name    string
	Addrs   []Addr

	extOff int
	raw    []byte
}

// An InterfaceAddrMessage represents an interface address message.
type InterfaceAddrMessage struct {
	Version int
	Type    int
	Flags   int
	Index   int
	Addrs   []Addr

	raw []byte
}

// An InterfaceMulticastAddrMessage represents an interface multicast
// address message.
type InterfaceMulticastAddrMessage struct {
	Version int
	Type    int
	Flags   int
	Index   int
	Addrs   []Addr

	raw []byte
}
