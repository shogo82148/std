// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Unix domain sockets stubs for Plan 9

package net

import (
	"github.com/shogo82148/std/time"
)

// UnixConn is an implementation of the Conn interface
// for connections to Unix domain sockets.
type UnixConn bool

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

// ReadFrom implements the PacketConn ReadFrom method.
func (c *UnixConn) ReadFrom(b []byte) (n int, addr Addr, err error)

// WriteTo implements the PacketConn WriteTo method.
func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err error)

// DialUnix connects to the remote address raddr on the network net,
// which must be "unix" or "unixgram".  If laddr is not nil, it is used
// as the local address for the connection.
func DialUnix(net string, laddr, raddr *UnixAddr) (c *UnixConn, err error)

// UnixListener is a Unix domain socket listener.
// Clients should typically use variables of type Listener
// instead of assuming Unix domain sockets.
type UnixListener bool

// ListenUnix announces on the Unix domain socket laddr and returns a Unix listener.
// Net must be "unix" (stream sockets).
func ListenUnix(net string, laddr *UnixAddr) (l *UnixListener, err error)

// Accept implements the Accept method in the Listener interface;
// it waits for the next call and returns a generic Conn.
func (l *UnixListener) Accept() (c Conn, err error)

// Close stops listening on the Unix address.
// Already accepted connections are not closed.
func (l *UnixListener) Close() error

// Addr returns the listener's network address.
func (l *UnixListener) Addr() Addr
