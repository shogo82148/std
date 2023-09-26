// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/time"
)

// Dial connects to the address addr on the network net.
//
// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
// (IPv4-only), "ip6" (IPv6-only), "unix" and "unixpacket".
//
// For TCP and UDP networks, addresses have the form host:port.
// If host is a literal IPv6 address, it must be enclosed
// in square brackets.  The functions JoinHostPort and SplitHostPort
// manipulate addresses in this form.
//
// Examples:
//
//	Dial("tcp", "12.34.56.78:80")
//	Dial("tcp", "google.com:80")
//	Dial("tcp", "[de:ad:be:ef::ca:fe]:80")
//
// For IP networks, addr must be "ip", "ip4" or "ip6" followed
// by a colon and a protocol number or name.
//
// Examples:
//
//	Dial("ip4:1", "127.0.0.1")
//	Dial("ip6:ospf", "::1")
func Dial(net, addr string) (Conn, error)

// DialTimeout acts like Dial but takes a timeout.
// The timeout includes name resolution, if required.
func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)

// Listen announces on the local network address laddr.
// The network string net must be a stream-oriented network:
// "tcp", "tcp4", "tcp6", or "unix", or "unixpacket".
func Listen(net, laddr string) (Listener, error)

// ListenPacket announces on the local network address laddr.
// The network string net must be a packet-oriented network:
// "udp", "udp4", "udp6", "ip", "ip4", "ip6" or "unixgram".
func ListenPacket(net, addr string) (PacketConn, error)
