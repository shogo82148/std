// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

// UnixConn is an implementation of the Conn interface for connections
// to Unix domain sockets.
type UnixConn struct {
	conn
}

// ReadFromUnix reads a packet from c, copying the payload into b.  It
// returns the number of bytes copied into b and the source address of
// the packet.
//
// ReadFromUnix can be made to time out and return an error with
// Timeout() == true after a fixed time limit; see SetDeadline and
// SetReadDeadline.
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)

// ReadMsgUnix reads a packet from c, copying the payload into b and
// the associated out-of-band data into oob.  It returns the number of
// bytes copied into b, the number of bytes copied into oob, the flags
// that were set on the packet, and the source address of the packet.
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)

// WriteToUnix writes a packet to addr via c, copying the payload from b.
//
// WriteToUnix can be made to time out and return an error with
// Timeout() == true after a fixed time limit; see SetDeadline and
// SetWriteDeadline.  On packet-oriented connections, write timeouts
// are rare.
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)

// WriteTo implements the PacketConn WriteTo method.
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)

// WriteMsgUnix writes a packet to addr via c, copying the payload
// from b and the associated out-of-band data from oob.  It returns
// the number of payload and out-of-band bytes written.
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)

// CloseRead shuts down the reading side of the Unix domain connection.
// Most callers should just use Close.
func (c *UnixConn) CloseRead() error

// CloseWrite shuts down the writing side of the Unix domain connection.
// Most callers should just use Close.
func (c *UnixConn) CloseWrite() error

// DialUnix connects to the remote address raddr on the network net,
// which must be "unix", "unixgram" or "unixpacket".  If laddr is not
// nil, it is used as the local address for the connection.
func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error)

// UnixListener is a Unix domain socket listener.  Clients should
// typically use variables of type Listener instead of assuming Unix
// domain sockets.
type UnixListener struct {
	fd *netFD
}

// ListenUnix announces on the Unix domain socket laddr and returns a
// Unix listener.  The network net must be "unix" or "unixpacket".
func ListenUnix(net string, laddr *UnixAddr) (*UnixListener, error)

// AcceptUnix accepts the next incoming call and returns the new
// connection.
func (l *UnixListener) AcceptUnix() (*UnixConn, error)

// Accept implements the Accept method in the Listener interface; it
// waits for the next call and returns a generic Conn.
func (l *UnixListener) Accept() (Conn, error)

// Close stops listening on the Unix address.  Already accepted
// connections are not closed.
func (l *UnixListener) Close() error

// Addr returns the listener's network address.
// The Addr returned is shared by all invocations of Addr, so
// do not modify it.
func (l *UnixListener) Addr() Addr

// SetDeadline sets the deadline associated with the listener.
// A zero time value disables the deadline.
func (l *UnixListener) SetDeadline(t time.Time) error

// File returns a copy of the underlying os.File, set to blocking
// mode.  It is the caller's responsibility to close f when finished.
// Closing l does not affect f, and closing f does not affect l.
//
// The returned os.File's file descriptor is different from the
// connection's.  Attempting to change properties of the original
// using this duplicate may or may not have the desired effect.
func (l *UnixListener) File() (*os.File, error)

// ListenUnixgram listens for incoming Unix datagram packets addressed
// to the local address laddr.  The network net must be "unixgram".
// The returned connection's ReadFrom and WriteTo methods can be used
// to receive and send packets with per-packet addressing.
func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error)
