// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// LookupHost looks up the given host using the local resolver.
// It returns an array of that host's addresses.
func LookupHost(host string) (addrs []string, err error)

// LookupIP looks up host using the local resolver.
// It returns an array of that host's IPv4 and IPv6 addresses.
func LookupIP(host string) (addrs []IP, err error)

// LookupPort looks up the port for the given network and service.
func LookupPort(network, service string) (port int, err error)

// LookupCNAME returns the canonical DNS host for the given name.
// Callers that do not care about the canonical name can call
// LookupHost or LookupIP directly; both take care of resolving
// the canonical name as part of the lookup.
func LookupCNAME(name string) (cname string, err error)

// LookupSRV tries to resolve an SRV query of the given service,
// protocol, and domain name.  The proto is "tcp" or "udp".
// The returned records are sorted by priority and randomized
// by weight within a priority.
//
// LookupSRV constructs the DNS name to look up following RFC 2782.
// That is, it looks up _service._proto.name.  To accommodate services
// publishing SRV records under non-standard names, if both service
// and proto are empty strings, LookupSRV looks up name directly.
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
func LookupMX(name string) (mx []*MX, err error)

// LookupTXT returns the DNS TXT records for the given domain name.
func LookupTXT(name string) (txt []string, err error)

// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address.
func LookupAddr(addr string) (name []string, err error)
