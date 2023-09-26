// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

// DNS client: see RFC 1035.
// Has to be linked into package net for Dial.

// TODO(rsc):
//	Could potentially handle many outstanding lookups faster.
//	Could have a small cache.
//	Random UDP source port (net.Dial should do that for us).
//	Random request IDs.

package net

// A dnsConn represents a DNS transport endpoint.

// dnsPacketConn implements the dnsConn interface for RFC 1035's
// "UDP usage" transport mechanism. Conn is a packet-oriented connection,
// such as a *UDPConn.

// dnsStreamConn implements the dnsConn interface for RFC 1035's
// "TCP usage" transport mechanism. Conn is a stream-oriented connection,
// such as a *TCPConn.

// A resolverConfig represents a DNS stub resolver configuration.

// hostLookupOrder specifies the order of LookupHost lookup strategies.
// It is basically a simplified representation of nsswitch.conf.
// "files" means /etc/hosts.
