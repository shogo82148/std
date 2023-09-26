// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd || windows
// +build darwin freebsd linux netbsd openbsd windows

// UDP sockets

package net

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

var ErrWriteToConnected = errors.New("use of WriteTo with pre-connected UDP")

// UDPConn is the implementation of the Conn and PacketConn
// interfaces for UDP network connections.
type UDPConn struct {
	fd *netFD
}

// Read implements the Conn Read method.
func (c *UDPConn) Read(b []byte) (int, error)

// Write implements the Conn Write method.
func (c *UDPConn) Write(b []byte) (int, error)

// Close closes the UDP connection.
func (c *UDPConn) Close() error

// LocalAddr returns the local network address.
func (c *UDPConn) LocalAddr() Addr

// RemoteAddr returns the remote network address, a *UDPAddr.
func (c *UDPConn) RemoteAddr() Addr

// SetDeadline implements the Conn SetDeadline method.
func (c *UDPConn) SetDeadline(t time.Time) error

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *UDPConn) SetReadDeadline(t time.Time) error

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *UDPConn) SetWriteDeadline(t time.Time) error

// SetReadBuffer sets the size of the operating system's
// receive buffer associated with the connection.
func (c *UDPConn) SetReadBuffer(bytes int) error

// SetWriteBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
func (c *UDPConn) SetWriteBuffer(bytes int) error

// ReadFromUDP reads a UDP packet from c, copying the payload into b.
// It returns the number of bytes copied into b and the return address
// that was on the packet.
//
// ReadFromUDP can be made to time out and return an error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)

// WriteToUDP writes a UDP packet to addr via c, copying the payload from b.
//
// WriteToUDP can be made to time out and return
// an error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetWriteDeadline.
// On packet-oriented connections, write timeouts are rare.
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)

// WriteTo implements the PacketConn WriteTo method.
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)

// File returns a copy of the underlying os.File, set to blocking mode.
// It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
func (c *UDPConn) File() (f *os.File, err error)

// DialUDP connects to the remote address raddr on the network net,
// which must be "udp", "udp4", or "udp6".  If laddr is not nil, it is used
// as the local address for the connection.
func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error)

// ListenUDP listens for incoming UDP packets addressed to the
// local address laddr.  The returned connection c's ReadFrom
// and WriteTo methods can be used to receive and send UDP
// packets with per-packet addressing.
func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)

// ListenMulticastUDP listens for incoming multicast UDP packets
// addressed to the group address gaddr on ifi, which specifies
// the interface to join.  ListenMulticastUDP uses default
// multicast interface if ifi is nil.
func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
