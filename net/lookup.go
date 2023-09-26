// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/context"
)

// protocols contains minimal mappings between internet protocol
// names and numbers for platforms that don't have a complete list of
// protocol numbers.
//
// See https://www.iana.org/assignments/protocol-numbers
//
// On Unix, this map is augmented by readProtocols via lookupProtocol.

// services contains minimal mappings between services names and port
// numbers for platforms that don't have a complete list of port numbers.
//
// See https://www.iana.org/assignments/service-names-port-numbers
//
// On Unix, this map is augmented by readServices via goLookupPort.

// dnsWaitGroup can be used by tests to wait for all DNS goroutines to
// complete. This avoids races on the test hooks.

// maxPortBufSize is the longest reasonable name of a service
// (non-numeric port).
// Currently the longest known IANA-unregistered name is
// "mobility-header", so we use that length, plus some slop in case
// something longer is added in the future.

// DefaultResolver is the resolver used by the package-level Lookup
// functions and by Dialers without a specified Resolver.
var DefaultResolver = &Resolver{}

// A Resolver looks up names and numbers.
//
// A nil *Resolver is equivalent to a zero Resolver.
type Resolver struct {
	PreferGo bool

	StrictErrors bool

	Dial func(ctx context.Context, network, address string) (Conn, error)

	lookupGroup singleflight.Group
}

// LookupHost looks up the given host using the local resolver.
// It returns a slice of that host's addresses.
func LookupHost(host string) (addrs []string, err error)

// LookupHost looks up the given host using the local resolver.
// It returns a slice of that host's addresses.
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)

// LookupIP looks up host using the local resolver.
// It returns a slice of that host's IPv4 and IPv6 addresses.
func LookupIP(host string) ([]IP, error)

// LookupIPAddr looks up host using the local resolver.
// It returns a slice of that host's IPv4 and IPv6 addresses.
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)

// LookupIP looks up host for the given network using the local resolver.
// It returns a slice of that host's IP addresses of the type specified by
// network.
// network must be one of "ip", "ip4" or "ip6".
func (r *Resolver) LookupIP(ctx context.Context, network, host string) ([]IP, error)

// onlyValuesCtx is a context that uses an underlying context
// for value lookup if the underlying context hasn't yet expired.

var _ context.Context = (*onlyValuesCtx)(nil)

// LookupPort looks up the port for the given network and service.
func LookupPort(network, service string) (port int, err error)

// LookupPort looks up the port for the given network and service.
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)

// LookupCNAME returns the canonical name for the given host.
// Callers that do not care about the canonical name can call
// LookupHost or LookupIP directly; both take care of resolving
// the canonical name as part of the lookup.
//
// A canonical name is the final name after following zero
// or more CNAME records.
// LookupCNAME does not return an error if host does not
// contain DNS "CNAME" records, as long as host resolves to
// address records.
//
// The returned canonical name is validated to be a properly
// formatted presentation-format domain name.
func LookupCNAME(host string) (cname string, err error)

// LookupCNAME returns the canonical name for the given host.
// Callers that do not care about the canonical name can call
// LookupHost or LookupIP directly; both take care of resolving
// the canonical name as part of the lookup.
//
// A canonical name is the final name after following zero
// or more CNAME records.
// LookupCNAME does not return an error if host does not
// contain DNS "CNAME" records, as long as host resolves to
// address records.
//
// The returned canonical name is validated to be a properly
// formatted presentation-format domain name.
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (string, error)

// LookupSRV tries to resolve an SRV query of the given service,
// protocol, and domain name. The proto is "tcp" or "udp".
// The returned records are sorted by priority and randomized
// by weight within a priority.
//
// LookupSRV constructs the DNS name to look up following RFC 2782.
// That is, it looks up _service._proto.name. To accommodate services
// publishing SRV records under non-standard names, if both service
// and proto are empty strings, LookupSRV looks up name directly.
//
// The returned service names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the the remaining results, if any.
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)

// LookupSRV tries to resolve an SRV query of the given service,
// protocol, and domain name. The proto is "tcp" or "udp".
// The returned records are sorted by priority and randomized
// by weight within a priority.
//
// LookupSRV constructs the DNS name to look up following RFC 2782.
// That is, it looks up _service._proto.name. To accommodate services
// publishing SRV records under non-standard names, if both service
// and proto are empty strings, LookupSRV looks up name directly.
//
// The returned service names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the the remaining results, if any.
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error)

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
//
// The returned mail server names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the the remaining results, if any.
//
// LookupMX uses context.Background internally; to specify the context, use
// Resolver.LookupMX.
func LookupMX(name string) ([]*MX, error)

// LookupMX returns the DNS MX records for the given domain name sorted by preference.
//
// The returned mail server names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the the remaining results, if any.
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error)

// LookupNS returns the DNS NS records for the given domain name.
//
// The returned name server names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the the remaining results, if any.
//
// LookupNS uses context.Background internally; to specify the context, use
// Resolver.LookupNS.
func LookupNS(name string) ([]*NS, error)

// LookupNS returns the DNS NS records for the given domain name.
//
// The returned name server names are validated to be properly
// formatted presentation-format domain names. If the response contains
// invalid names, those records are filtered out and an error
// will be returned alongside the the remaining results, if any.
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error)

// LookupTXT returns the DNS TXT records for the given domain name.
func LookupTXT(name string) ([]string, error)

// LookupTXT returns the DNS TXT records for the given domain name.
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)

// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address.
//
// The returned names are validated to be properly formatted presentation-format
// domain names. If the response contains invalid names, those records are filtered
// out and an error will be returned alongside the the remaining results, if any.
//
// When using the host C library resolver, at most one result will be
// returned. To bypass the host resolver, use a custom Resolver.
func LookupAddr(addr string) (names []string, err error)

// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address.
//
// The returned names are validated to be properly formatted presentation-format
// domain names. If the response contains invalid names, those records are filtered
// out and an error will be returned alongside the the remaining results, if any.
func (r *Resolver) LookupAddr(ctx context.Context, addr string) ([]string, error)

// errMalformedDNSRecordsDetail is the DNSError detail which is returned when a Resolver.Lookup...
// method recieves DNS records which contain invalid DNS names. This may be returned alongside
// results which have had the malformed records filtered out.
