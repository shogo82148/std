// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// UDP sockets

package net

// UDPAddr represents the address of a UDP end point.
type UDPAddr struct {
	IP   IP
	Port int
}

// Network returns the address's network name, "udp".
func (a *UDPAddr) Network() string

func (a *UDPAddr) String() string

// ResolveUDPAddr parses addr as a UDP address of the form
// host:port and resolves domain names or port names to
// numeric addresses on the network net, which must be "udp",
// "udp4" or "udp6".  A literal IPv6 host address must be
// enclosed in square brackets, as in "[::]:80".
func ResolveUDPAddr(net, addr string) (*UDPAddr, error)
