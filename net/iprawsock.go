// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// (Raw) IP sockets

package net

// IPAddr represents the address of a IP end point.
type IPAddr struct {
	IP IP
}

// Network returns the address's network name, "ip".
func (a *IPAddr) Network() string

func (a *IPAddr) String() string

// ResolveIPAddr parses addr as a IP address and resolves domain
// names to numeric addresses on the network net, which must be
// "ip", "ip4" or "ip6".  A literal IPv6 host address must be
// enclosed in square brackets, as in "[::]".
func ResolveIPAddr(net, addr string) (*IPAddr, error)
