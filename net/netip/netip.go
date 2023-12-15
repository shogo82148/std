// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package netip defines an IP address type that's a small value type.
// Building on that [Addr] type, the package also defines [AddrPort] (an
// IP address and a port) and [Prefix] (an IP address and a bit length
// prefix).
//
// Compared to the [net.IP] type, [Addr] type takes less memory, is immutable,
// and is comparable (supports == and being a map key).
package netip

import (
	"github.com/shogo82148/std/internal/intern"
)

// Addr represents an IPv4 or IPv6 address (with or without a scoped
// addressing zone), similar to [net.IP] or [net.IPAddr].
//
// Unlike [net.IP] or [net.IPAddr], Addr is a comparable value
// type (it supports == and can be a map key) and is immutable.
//
// The zero Addr is not a valid IP address.
// Addr{} is distinct from both 0.0.0.0 and ::.
type Addr struct {
	// addr is the hi and lo bits of an IPv6 address. If z==z4,
	// hi and lo contain the IPv4-mapped IPv6 address.
	//
	// hi and lo are constructed by interpreting a 16-byte IPv6
	// address as a big-endian 128-bit number. The most significant
	// bits of that number go into hi, the rest into lo.
	//
	// For example, 0011:2233:4455:6677:8899:aabb:ccdd:eeff is stored as:
	//  addr.hi = 0x0011223344556677
	//  addr.lo = 0x8899aabbccddeeff
	//
	// We store IPs like this, rather than as [16]byte, because it
	// turns most operations on IPs into arithmetic and bit-twiddling
	// operations on 64-bit registers, which is much faster than
	// bytewise processing.
	addr uint128

	// z is a combination of the address family and the IPv6 zone.
	//
	// nil means invalid IP address (for a zero Addr).
	// z4 means an IPv4 address.
	// z6noz means an IPv6 address without a zone.
	//
	// Otherwise it's the interned zone name string.
	z *intern.Value
}

// IPv6LinkLocalAllNodes returns the IPv6 link-local all nodes multicast
// address ff02::1.
func IPv6LinkLocalAllNodes() Addr

// IPv6LinkLocalAllRouters returns the IPv6 link-local all routers multicast
// address ff02::2.
func IPv6LinkLocalAllRouters() Addr

// IPv6Loopback returns the IPv6 loopback address ::1.
func IPv6Loopback() Addr

// IPv6Unspecified returns the IPv6 unspecified address "::".
func IPv6Unspecified() Addr

// IPv4Unspecified returns the IPv4 unspecified address "0.0.0.0".
func IPv4Unspecified() Addr

// AddrFrom4 returns the address of the IPv4 address given by the bytes in addr.
func AddrFrom4(addr [4]byte) Addr

// AddrFrom16 returns the IPv6 address given by the bytes in addr.
// An IPv4-mapped IPv6 address is left as an IPv6 address.
// (Use Unmap to convert them if needed.)
func AddrFrom16(addr [16]byte) Addr

// ParseAddr parses s as an IP address, returning the result. The string
// s can be in dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"),
// or IPv6 with a scoped addressing zone ("fe80::1cc0:3e8c:119f:c2e1%ens18").
func ParseAddr(s string) (Addr, error)

// MustParseAddr calls ParseAddr(s) and panics on error.
// It is intended for use in tests with hard-coded strings.
func MustParseAddr(s string) Addr

// AddrFromSlice parses the 4- or 16-byte byte slice as an IPv4 or IPv6 address.
// Note that a net.IP can be passed directly as the []byte argument.
// If slice's length is not 4 or 16, AddrFromSlice returns Addr{}, false.
func AddrFromSlice(slice []byte) (ip Addr, ok bool)

// IsValid reports whether the Addr is an initialized address (not the zero Addr).
//
// Note that "0.0.0.0" and "::" are both valid values.
func (ip Addr) IsValid() bool

// BitLen returns the number of bits in the IP address:
// 128 for IPv6, 32 for IPv4, and 0 for the zero Addr.
//
// Note that IPv4-mapped IPv6 addresses are considered IPv6 addresses
// and therefore have bit length 128.
func (ip Addr) BitLen() int

// Zone returns ip's IPv6 scoped addressing zone, if any.
func (ip Addr) Zone() string

// Compare returns an integer comparing two IPs.
// The result will be 0 if ip == ip2, -1 if ip < ip2, and +1 if ip > ip2.
// The definition of "less than" is the same as the Less method.
func (ip Addr) Compare(ip2 Addr) int

// Less reports whether ip sorts before ip2.
// IP addresses sort first by length, then their address.
// IPv6 addresses with zones sort just after the same address without a zone.
func (ip Addr) Less(ip2 Addr) bool

// Is4 reports whether ip is an IPv4 address.
//
// It returns false for IPv4-mapped IPv6 addresses. See Addr.Unmap.
func (ip Addr) Is4() bool

// Is4In6 reports whether ip is an IPv4-mapped IPv6 address.
func (ip Addr) Is4In6() bool

// Is6 reports whether ip is an IPv6 address, including IPv4-mapped
// IPv6 addresses.
func (ip Addr) Is6() bool

// Unmap returns ip with any IPv4-mapped IPv6 address prefix removed.
//
// That is, if ip is an IPv6 address wrapping an IPv4 address, it
// returns the wrapped IPv4 address. Otherwise it returns ip unmodified.
func (ip Addr) Unmap() Addr

// WithZone returns an IP that's the same as ip but with the provided
// zone. If zone is empty, the zone is removed. If ip is an IPv4
// address, WithZone is a no-op and returns ip unchanged.
func (ip Addr) WithZone(zone string) Addr

// IsLinkLocalUnicast reports whether ip is a link-local unicast address.
func (ip Addr) IsLinkLocalUnicast() bool

// IsLoopback reports whether ip is a loopback address.
func (ip Addr) IsLoopback() bool

// IsMulticast reports whether ip is a multicast address.
func (ip Addr) IsMulticast() bool

// IsInterfaceLocalMulticast reports whether ip is an IPv6 interface-local
// multicast address.
func (ip Addr) IsInterfaceLocalMulticast() bool

// IsLinkLocalMulticast reports whether ip is a link-local multicast address.
func (ip Addr) IsLinkLocalMulticast() bool

// IsGlobalUnicast reports whether ip is a global unicast address.
//
// It returns true for IPv6 addresses which fall outside of the current
// IANA-allocated 2000::/3 global unicast space, with the exception of the
// link-local address space. It also returns true even if ip is in the IPv4
// private address space or IPv6 unique local address space.
// It returns false for the zero Addr.
//
// For reference, see RFC 1122, RFC 4291, and RFC 4632.
func (ip Addr) IsGlobalUnicast() bool

// IsPrivate reports whether ip is a private address, according to RFC 1918
// (IPv4 addresses) and RFC 4193 (IPv6 addresses). That is, it reports whether
// ip is in 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16, or fc00::/7. This is the
// same as net.IP.IsPrivate.
func (ip Addr) IsPrivate() bool

// IsUnspecified reports whether ip is an unspecified address, either the IPv4
// address "0.0.0.0" or the IPv6 address "::".
//
// Note that the zero Addr is not an unspecified address.
func (ip Addr) IsUnspecified() bool

// Prefix keeps only the top b bits of IP, producing a Prefix
// of the specified length.
// If ip is a zero Addr, Prefix always returns a zero Prefix and a nil error.
// Otherwise, if bits is less than zero or greater than ip.BitLen(),
// Prefix returns an error.
func (ip Addr) Prefix(b int) (Prefix, error)

// As16 returns the IP address in its 16-byte representation.
// IPv4 addresses are returned as IPv4-mapped IPv6 addresses.
// IPv6 addresses with zones are returned without their zone (use the
// Zone method to get it).
// The ip zero value returns all zeroes.
func (ip Addr) As16() (a16 [16]byte)

// As4 returns an IPv4 or IPv4-in-IPv6 address in its 4-byte representation.
// If ip is the zero Addr or an IPv6 address, As4 panics.
// Note that 0.0.0.0 is not the zero Addr.
func (ip Addr) As4() (a4 [4]byte)

// AsSlice returns an IPv4 or IPv6 address in its respective 4-byte or 16-byte representation.
func (ip Addr) AsSlice() []byte

// Next returns the address following ip.
// If there is none, it returns the zero Addr.
func (ip Addr) Next() Addr

// Prev returns the IP before ip.
// If there is none, it returns the IP zero value.
func (ip Addr) Prev() Addr

// String returns the string form of the IP address ip.
// It returns one of 5 forms:
//
//   - "invalid IP", if ip is the zero Addr
//   - IPv4 dotted decimal ("192.0.2.1")
//   - IPv6 ("2001:db8::1")
//   - "::ffff:1.2.3.4" (if Is4In6)
//   - IPv6 with zone ("fe80:db8::1%eth0")
//
// Note that unlike package net's IP.String method,
// IPv4-mapped IPv6 addresses format with a "::ffff:"
// prefix before the dotted quad.
func (ip Addr) String() string

// AppendTo appends a text encoding of ip,
// as generated by MarshalText,
// to b and returns the extended buffer.
func (ip Addr) AppendTo(b []byte) []byte

// StringExpanded is like String but IPv6 addresses are expanded with leading
// zeroes and no "::" compression. For example, "2001:db8::1" becomes
// "2001:0db8:0000:0000:0000:0000:0000:0001".
func (ip Addr) StringExpanded() string

// MarshalText implements the encoding.TextMarshaler interface,
// The encoding is the same as returned by String, with one exception:
// If ip is the zero Addr, the encoding is the empty string.
func (ip Addr) MarshalText() ([]byte, error)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The IP address is expected in a form accepted by ParseAddr.
//
// If text is empty, UnmarshalText sets *ip to the zero Addr and
// returns no error.
func (ip *Addr) UnmarshalText(text []byte) error

// MarshalBinary implements the encoding.BinaryMarshaler interface.
// It returns a zero-length slice for the zero Addr,
// the 4-byte form for an IPv4 address,
// and the 16-byte form with zone appended for an IPv6 address.
func (ip Addr) MarshalBinary() ([]byte, error)

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It expects data in the form generated by MarshalBinary.
func (ip *Addr) UnmarshalBinary(b []byte) error

// AddrPort is an IP and a port number.
type AddrPort struct {
	ip   Addr
	port uint16
}

// AddrPortFrom returns an AddrPort with the provided IP and port.
// It does not allocate.
func AddrPortFrom(ip Addr, port uint16) AddrPort

// Addr returns p's IP address.
func (p AddrPort) Addr() Addr

// Port returns p's port.
func (p AddrPort) Port() uint16

// ParseAddrPort parses s as an AddrPort.
//
// It doesn't do any name resolution: both the address and the port
// must be numeric.
func ParseAddrPort(s string) (AddrPort, error)

// MustParseAddrPort calls ParseAddrPort(s) and panics on error.
// It is intended for use in tests with hard-coded strings.
func MustParseAddrPort(s string) AddrPort

// IsValid reports whether p.Addr() is valid.
// All ports are valid, including zero.
func (p AddrPort) IsValid() bool

// Compare returns an integer comparing two AddrPorts.
// The result will be 0 if p == p2, -1 if p < p2, and +1 if p > p2.
// AddrPorts sort first by IP address, then port.
func (p AddrPort) Compare(p2 AddrPort) int

func (p AddrPort) String() string

// AppendTo appends a text encoding of p,
// as generated by MarshalText,
// to b and returns the extended buffer.
func (p AddrPort) AppendTo(b []byte) []byte

// MarshalText implements the encoding.TextMarshaler interface. The
// encoding is the same as returned by String, with one exception: if
// p.Addr() is the zero Addr, the encoding is the empty string.
func (p AddrPort) MarshalText() ([]byte, error)

// UnmarshalText implements the encoding.TextUnmarshaler
// interface. The AddrPort is expected in a form
// generated by MarshalText or accepted by ParseAddrPort.
func (p *AddrPort) UnmarshalText(text []byte) error

// MarshalBinary implements the encoding.BinaryMarshaler interface.
// It returns Addr.MarshalBinary with an additional two bytes appended
// containing the port in little-endian.
func (p AddrPort) MarshalBinary() ([]byte, error)

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It expects data in the form generated by MarshalBinary.
func (p *AddrPort) UnmarshalBinary(b []byte) error

// Prefix is an IP address prefix (CIDR) representing an IP network.
//
// The first Bits() of Addr() are specified. The remaining bits match any address.
// The range of Bits() is [0,32] for IPv4 or [0,128] for IPv6.
type Prefix struct {
	ip Addr

	// bitsPlusOne stores the prefix bit length plus one.
	// A Prefix is valid if and only if bitsPlusOne is non-zero.
	bitsPlusOne uint8
}

// PrefixFrom returns a Prefix with the provided IP address and bit
// prefix length.
//
// It does not allocate. Unlike Addr.Prefix, PrefixFrom does not mask
// off the host bits of ip.
//
// If bits is less than zero or greater than ip.BitLen, Prefix.Bits
// will return an invalid value -1.
func PrefixFrom(ip Addr, bits int) Prefix

// Addr returns p's IP address.
func (p Prefix) Addr() Addr

// Bits returns p's prefix length.
//
// It reports -1 if invalid.
func (p Prefix) Bits() int

// IsValid reports whether p.Bits() has a valid range for p.Addr().
// If p.Addr() is the zero Addr, IsValid returns false.
// Note that if p is the zero Prefix, then p.IsValid() == false.
func (p Prefix) IsValid() bool

// IsSingleIP reports whether p contains exactly one IP.
func (p Prefix) IsSingleIP() bool

// ParsePrefix parses s as an IP address prefix.
// The string can be in the form "192.168.1.0/24" or "2001:db8::/32",
// the CIDR notation defined in RFC 4632 and RFC 4291.
// IPv6 zones are not permitted in prefixes, and an error will be returned if a
// zone is present.
//
// Note that masked address bits are not zeroed. Use Masked for that.
func ParsePrefix(s string) (Prefix, error)

// MustParsePrefix calls ParsePrefix(s) and panics on error.
// It is intended for use in tests with hard-coded strings.
func MustParsePrefix(s string) Prefix

// Masked returns p in its canonical form, with all but the high
// p.Bits() bits of p.Addr() masked off.
//
// If p is zero or otherwise invalid, Masked returns the zero Prefix.
func (p Prefix) Masked() Prefix

// Contains reports whether the network p includes ip.
//
// An IPv4 address will not match an IPv6 prefix.
// An IPv4-mapped IPv6 address will not match an IPv4 prefix.
// A zero-value IP will not match any prefix.
// If ip has an IPv6 zone, Contains returns false,
// because Prefixes strip zones.
func (p Prefix) Contains(ip Addr) bool

// Overlaps reports whether p and o contain any IP addresses in common.
//
// If p and o are of different address families or either have a zero
// IP, it reports false. Like the Contains method, a prefix with an
// IPv4-mapped IPv6 address is still treated as an IPv6 mask.
func (p Prefix) Overlaps(o Prefix) bool

// AppendTo appends a text encoding of p,
// as generated by MarshalText,
// to b and returns the extended buffer.
func (p Prefix) AppendTo(b []byte) []byte

// MarshalText implements the encoding.TextMarshaler interface,
// The encoding is the same as returned by String, with one exception:
// If p is the zero value, the encoding is the empty string.
func (p Prefix) MarshalText() ([]byte, error)

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The IP address is expected in a form accepted by ParsePrefix
// or generated by MarshalText.
func (p *Prefix) UnmarshalText(text []byte) error

// MarshalBinary implements the encoding.BinaryMarshaler interface.
// It returns Addr.MarshalBinary with an additional byte appended
// containing the prefix bits.
func (p Prefix) MarshalBinary() ([]byte, error)

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It expects data in the form generated by MarshalBinary.
func (p *Prefix) UnmarshalBinary(b []byte) error

// String returns the CIDR notation of p: "<ip>/<bits>".
func (p Prefix) String() string
