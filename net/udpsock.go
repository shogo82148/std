// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/net/netip"
	"github.com/shogo82148/std/syscall"
)

// UDPAddr represents the address of a UDP end point.
type UDPAddr struct {
	IP   IP
	Port int
	Zone string
}

// AddrPort returns the UDPAddr a as a netip.AddrPort.
//
// If a.Port does not fit in a uint16, it's silently truncated.
//
// If a is nil, a zero value is returned.
func (a *UDPAddr) AddrPort() netip.AddrPort

// Network returns the address's network name, "udp".
func (a *UDPAddr) Network() string

func (a *UDPAddr) String() string

// ResolveUDPAddr returns an address of UDP end point.
//
// The network must be a UDP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveUDPAddr resolves the
// address to an address of UDP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveUDPAddr(network, address string) (*UDPAddr, error)

// UDPAddrFromAddrPort returns addr as a UDPAddr. If addr.IsValid() is false,
// then the returned UDPAddr will contain a nil IP field, indicating an
// address family-agnostic unspecified address.
func UDPAddrFromAddrPort(addr netip.AddrPort) *UDPAddr

// An addrPortUDPAddr is a netip.AddrPort-based UDP address that satisfies the Addr interface.

// UDPConn is the implementation of the Conn and PacketConn interfaces
// for UDP network connections.
type UDPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
// This implements the syscall.Conn interface.
func (c *UDPConn) SyscallConn() (syscall.RawConn, error)

// ReadFromUDP acts like ReadFrom but returns a UDPAddr.
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)

// ReadFromUDPAddrPort acts like ReadFrom but returns a netip.AddrPort.
func (c *UDPConn) ReadFromUDPAddrPort(b []byte) (n int, addr netip.AddrPort, err error)

// ReadMsgUDP reads a message from c, copying the payload into b and
// the associated out-of-band data into oob. It returns the number of
// bytes copied into b, the number of bytes copied into oob, the flags
// that were set on the message and the source address of the message.
//
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
// used to manipulate IP-level socket options in oob.
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)

// ReadMsgUDPAddrPort is like ReadMsgUDP but returns an netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error)

// WriteToUDP acts like WriteTo but takes a UDPAddr.
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)

// WriteToUDPAddrPort acts like WriteTo but takes a netip.AddrPort.
func (c *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error)

// WriteTo implements the PacketConn WriteTo method.
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)

// WriteMsgUDP writes a message to addr via c if c isn't connected, or
// to c's remote address if c is connected (in which case addr must be
// nil). The payload is copied from b and the associated out-of-band
// data is copied from oob. It returns the number of payload and
// out-of-band bytes written.
//
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
// used to manipulate IP-level socket options in oob.
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)

// WriteMsgUDPAddrPort is like WriteMsgUDP but takes a netip.AddrPort instead of a UDPAddr.
func (c *UDPConn) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error)

// DialUDP acts like Dial for UDP networks.
//
// The network must be a UDP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)

// ListenUDP acts like ListenPacket for UDP networks.
//
// The network must be a UDP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenUDP listens on all available IP addresses of the local system
// except multicast IP addresses.
// If the Port field of laddr is 0, a port number is automatically
// chosen.
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)

// ListenMulticastUDP acts like ListenPacket for UDP networks but
// takes a group address on a specific network interface.
//
// The network must be a UDP network name; see func Dial for details.
//
// ListenMulticastUDP listens on all available IP addresses of the
// local system including the group, multicast IP address.
// If ifi is nil, ListenMulticastUDP uses the system-assigned
// multicast interface, although this is not recommended because the
// assignment depends on platforms and sometimes it might require
// routing configuration.
// If the Port field of gaddr is 0, a port number is automatically
// chosen.
//
// ListenMulticastUDP is just for convenience of simple, small
// applications. There are golang.org/x/net/ipv4 and
// golang.org/x/net/ipv6 packages for general purpose uses.
//
// Note that ListenMulticastUDP will set the IP_MULTICAST_LOOP socket option
// to 0 under IPPROTO_IP, to disable loopback of multicast packets.
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
