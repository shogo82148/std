// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd || windows
// +build darwin freebsd linux netbsd openbsd windows

// Unix domain sockets

package net

import (
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

// UnixConn is an implementation of the Conn interface
// for connections to Unix domain sockets.
type UnixConn struct {
	fd *netFD
}

// Read implements the Conn Read method.
func (c *UnixConn) Read(b []byte) (n int, err error)

// Write implements the Conn Write method.
func (c *UnixConn) Write(b []byte) (n int, err error)

// Close closes the Unix domain connection.
func (c *UnixConn) Close() error

// LocalAddr returns the local network address, a *UnixAddr.
// Unlike in other protocols, LocalAddr is usually nil for dialed connections.
func (c *UnixConn) LocalAddr() Addr

// RemoteAddr returns the remote network address, a *UnixAddr.
// Unlike in other protocols, RemoteAddr is usually nil for connections
// accepted by a listener.
func (c *UnixConn) RemoteAddr() Addr

// SetDeadline implements the Conn SetDeadline method.
func (c *UnixConn) SetDeadline(t time.Time) error

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *UnixConn) SetReadDeadline(t time.Time) error

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *UnixConn) SetWriteDeadline(t time.Time) error

// SetReadBuffer sets the size of the operating system's
// receive buffer associated with the connection.
func (c *UnixConn) SetReadBuffer(bytes int) error

// SetWriteBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
func (c *UnixConn) SetWriteBuffer(bytes int) error

// ReadFromUnix reads a packet from c, copying the payload into b.
// It returns the number of bytes copied into b and the source address
// of the packet.
//
// ReadFromUnix can be made to time out and return
// an error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetReadDeadline.
func (c *UnixConn) ReadFromUnix(b []byte) (n int, addr *UnixAddr, err error)

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UnixConn) ReadFrom(b []byte) (n int, addr Addr, err error)

// WriteToUnix writes a packet to addr via c, copying the payload from b.
//
// WriteToUnix can be made to time out and return
// an error with Timeout() == true after a fixed time limit;
// see SetDeadline and SetWriteDeadline.
// On packet-oriented connections, write timeouts are rare.
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (n int, err error)

// WriteTo implements the PacketConn WriteTo method.
func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err error)

// ReadMsgUnix reads a packet from c, copying the payload into b
// and the associated out-of-band data into oob.
// It returns the number of bytes copied into b, the number of
// bytes copied into oob, the flags that were set on the packet,
// and the source address of the packet.
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)

// WriteMsgUnix writes a packet to addr via c, copying the payload from b
// and the associated out-of-band data from oob.  It returns the number
// of payload and out-of-band bytes written.
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)

// File returns a copy of the underlying os.File, set to blocking mode.
// It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
func (c *UnixConn) File() (f *os.File, err error)

// DialUnix connects to the remote address raddr on the network net,
// which must be "unix" or "unixgram".  If laddr is not nil, it is used
// as the local address for the connection.
func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error)

// UnixListener is a Unix domain socket listener.
// Clients should typically use variables of type Listener
// instead of assuming Unix domain sockets.
type UnixListener struct {
	fd   *netFD
	path string
}

// ListenUnix announces on the Unix domain socket laddr and returns a Unix listener.
// Net must be "unix" (stream sockets).
func ListenUnix(net string, laddr *UnixAddr) (*UnixListener, error)

// AcceptUnix accepts the next incoming call and returns the new connection
// and the remote address.
func (l *UnixListener) AcceptUnix() (*UnixConn, error)

// Accept implements the Accept method in the Listener interface;
// it waits for the next call and returns a generic Conn.
func (l *UnixListener) Accept() (c Conn, err error)

// Close stops listening on the Unix address.
// Already accepted connections are not closed.
func (l *UnixListener) Close() error

// Addr returns the listener's network address.
func (l *UnixListener) Addr() Addr

// SetDeadline sets the deadline associated with the listener.
// A zero time value disables the deadline.
func (l *UnixListener) SetDeadline(t time.Time) (err error)

// File returns a copy of the underlying os.File, set to blocking mode.
// It is the caller's responsibility to close f when finished.
// Closing l does not affect f, and closing f does not affect l.
func (l *UnixListener) File() (f *os.File, err error)

// ListenUnixgram listens for incoming Unix datagram packets addressed to the
// local address laddr.  The returned connection c's ReadFrom
// and WriteTo methods can be used to receive and send UDP
// packets with per-packet addressing.  The network net must be "unixgram".
func ListenUnixgram(net string, laddr *UnixAddr) (*UDPConn, error)
