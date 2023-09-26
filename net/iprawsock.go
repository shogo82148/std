// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/syscall"
)

// IPAddr represents the address of an IP end point.
type IPAddr struct {
	IP   IP
	Zone string
}

// Network returns the address's network name, "ip".
func (a *IPAddr) Network() string

func (a *IPAddr) String() string

// ResolveIPAddr returns an address of IP end point.
//
// The network must be an IP network name.
//
// If the host in the address parameter is not a literal IP address,
// ResolveIPAddr resolves the address to an address of IP end point.
// Otherwise, it parses the address as a literal IP address.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveIPAddr(network, address string) (*IPAddr, error)

// IPConn is the implementation of the Conn and PacketConn interfaces
// for IP network connections.
type IPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
// This implements the syscall.Conn interface.
func (c *IPConn) SyscallConn() (syscall.RawConn, error)

// ReadFromIP acts like ReadFrom but returns an IPAddr.
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)

// ReadFrom implements the PacketConn ReadFrom method.
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)

// ReadMsgIP reads a message from c, copying the payload into b and
// the associated out-of-band data into oob. It returns the number of
// bytes copied into b, the number of bytes copied into oob, the flags
// that were set on the message and the source address of the message.
//
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
// used to manipulate IP-level socket options in oob.
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)

// WriteToIP acts like WriteTo but takes an IPAddr.
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)

// WriteTo implements the PacketConn WriteTo method.
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)

// WriteMsgIP writes a message to addr via c, copying the payload from
// b and the associated out-of-band data from oob. It returns the
// number of payload and out-of-band bytes written.
//
// The packages golang.org/x/net/ipv4 and golang.org/x/net/ipv6 can be
// used to manipulate IP-level socket options in oob.
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)

// DialIP acts like Dial for IP networks.
//
// The network must be an IP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)

// ListenIP acts like ListenPacket for IP networks.
//
// The network must be an IP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenIP listens on all available IP addresses of the local system
// except multicast IP addresses.
func ListenIP(network string, laddr *IPAddr) (*IPConn, error)
