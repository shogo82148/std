// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// IPAddr represents the address of an IP end point.
type IPAddr struct {
	IP   IP
	Zone string
}

// Network returns the address's network name, "ip".
func (a *IPAddr) Network() string

func (a *IPAddr) String() string

// ResolveIPAddr parses addr as an IP address of the form "host" or
// "ipv6-host%zone" and resolves the domain name on the network net,
// which must be "ip", "ip4" or "ip6".
func ResolveIPAddr(net, addr string) (*IPAddr, error)
