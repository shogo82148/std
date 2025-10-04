// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package routebsd

import (
	"github.com/shogo82148/std/net/netip"
)

// An Addr represents an address associated with packet routing.
type Addr interface {
	Family() int
}

// A LinkAddr represents a link-layer address.
type LinkAddr struct {
	Index int
	Name  string
	Addr  []byte
}

// Family implements the Family method of Addr interface.
func (a *LinkAddr) Family() int

// An InetAddr represent an internet address using IPv4 or IPv6.
type InetAddr struct {
	IP netip.Addr
}

func (a *InetAddr) Family() int
