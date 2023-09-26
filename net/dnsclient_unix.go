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

// A dnsDialer provides dialing suitable for DNS queries.

// A dnsConn represents a DNS transport endpoint.

// A resolverConfig represents a DNS stub resolver configuration.

// hostLookupOrder specifies the order of LookupHost lookup strategies.
// It is basically a simplified representation of nsswitch.conf.
// "files" means /etc/hosts.
