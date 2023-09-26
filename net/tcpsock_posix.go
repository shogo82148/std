// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || freebsd || linux || netbsd || openbsd || windows
// +build darwin freebsd linux netbsd openbsd windows

// TCP sockets

package net

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

// TCPConn is an implementation of the Conn interface
// for TCP network connections.
type TCPConn struct {
	fd *netFD
}

// Read implements the Conn Read method.
func (c *TCPConn) Read(b []byte) (n int, err error)

// ReadFrom implements the io.ReaderFrom ReadFrom method.
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)

// Write implements the Conn Write method.
func (c *TCPConn) Write(b []byte) (n int, err error)

// Close closes the TCP connection.
func (c *TCPConn) Close() error

// CloseRead shuts down the reading side of the TCP connection.
// Most callers should just use Close.
func (c *TCPConn) CloseRead() error

// CloseWrite shuts down the writing side of the TCP connection.
// Most callers should just use Close.
func (c *TCPConn) CloseWrite() error

// LocalAddr returns the local network address, a *TCPAddr.
func (c *TCPConn) LocalAddr() Addr

// RemoteAddr returns the remote network address, a *TCPAddr.
func (c *TCPConn) RemoteAddr() Addr

// SetDeadline implements the Conn SetDeadline method.
func (c *TCPConn) SetDeadline(t time.Time) error

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *TCPConn) SetReadDeadline(t time.Time) error

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *TCPConn) SetWriteDeadline(t time.Time) error

// SetReadBuffer sets the size of the operating system's
// receive buffer associated with the connection.
func (c *TCPConn) SetReadBuffer(bytes int) error

// SetWriteBuffer sets the size of the operating system's
// transmit buffer associated with the connection.
func (c *TCPConn) SetWriteBuffer(bytes int) error

// SetLinger sets the behavior of Close() on a connection
// which still has data waiting to be sent or to be acknowledged.
//
// If sec < 0 (the default), Close returns immediately and
// the operating system finishes sending the data in the background.
//
// If sec == 0, Close returns immediately and the operating system
// discards any unsent or unacknowledged data.
//
// If sec > 0, Close blocks for at most sec seconds waiting for
// data to be sent and acknowledged.
func (c *TCPConn) SetLinger(sec int) error

// SetKeepAlive sets whether the operating system should send
// keepalive messages on the connection.
func (c *TCPConn) SetKeepAlive(keepalive bool) error

// SetNoDelay controls whether the operating system should delay
// packet transmission in hopes of sending fewer packets
// (Nagle's algorithm).  The default is true (no delay), meaning
// that data is sent as soon as possible after a Write.
func (c *TCPConn) SetNoDelay(noDelay bool) error

// File returns a copy of the underlying os.File, set to blocking mode.
// It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
func (c *TCPConn) File() (f *os.File, err error)

// DialTCP connects to the remote address raddr on the network net,
// which must be "tcp", "tcp4", or "tcp6".  If laddr is not nil, it is used
// as the local address for the connection.
func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)

// TCPListener is a TCP network listener.
// Clients should typically use variables of type Listener
// instead of assuming TCP.
type TCPListener struct {
	fd *netFD
}

// ListenTCP announces on the TCP address laddr and returns a TCP listener.
// Net must be "tcp", "tcp4", or "tcp6".
// If laddr has a port of 0, it means to listen on some available port.
// The caller can use l.Addr() to retrieve the chosen address.
func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)

// AcceptTCP accepts the next incoming call and returns the new connection
// and the remote address.
func (l *TCPListener) AcceptTCP() (c *TCPConn, err error)

// Accept implements the Accept method in the Listener interface;
// it waits for the next call and returns a generic Conn.
func (l *TCPListener) Accept() (c Conn, err error)

// Close stops listening on the TCP address.
// Already Accepted connections are not closed.
func (l *TCPListener) Close() error

// Addr returns the listener's network address, a *TCPAddr.
func (l *TCPListener) Addr() Addr

// SetDeadline sets the deadline associated with the listener.
// A zero time value disables the deadline.
func (l *TCPListener) SetDeadline(t time.Time) error

// File returns a copy of the underlying os.File, set to blocking mode.
// It is the caller's responsibility to close f when finished.
// Closing c does not affect f, and closing f does not affect c.
func (l *TCPListener) File() (f *os.File, err error)
