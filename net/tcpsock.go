// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	IP   IP
	Port int
	Zone string
}

// Network returns the address's network name, "tcp".
func (a *TCPAddr) Network() string

func (a *TCPAddr) String() string

// ResolveTCPAddr parses addr as a TCP address of the form "host:port"
// or "[ipv6-host%zone]:port" and resolves a pair of domain name and
// port name on the network net, which must be "tcp", "tcp4" or
// "tcp6".  A literal address or host name for IPv6 must be enclosed
// in square brackets, as in "[::1]:80", "[ipv6-host]:http" or
// "[ipv6-host%zone]:80".
func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
