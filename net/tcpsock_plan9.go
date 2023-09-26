// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TCP for Plan 9

package net

import (
	"github.com/shogo82148/std/time"
)

// TCPConn is an implementation of the Conn interface
// for TCP network connections.
type TCPConn struct {
	plan9Conn
}

// SetDeadline implements the Conn SetDeadline method.
func (c *TCPConn) SetDeadline(t time.Time) error

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *TCPConn) SetReadDeadline(t time.Time) error

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *TCPConn) SetWriteDeadline(t time.Time) error

// CloseRead shuts down the reading side of the TCP connection.
// Most callers should just use Close.
func (c *TCPConn) CloseRead() error

// CloseWrite shuts down the writing side of the TCP connection.
// Most callers should just use Close.
func (c *TCPConn) CloseWrite() error

// DialTCP connects to the remote address raddr on the network net,
// which must be "tcp", "tcp4", or "tcp6".  If laddr is not nil, it is used
// as the local address for the connection.
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err error)

// TCPListener is a TCP network listener.
// Clients should typically use variables of type Listener
// instead of assuming TCP.
type TCPListener struct {
	plan9Listener
}

// ListenTCP announces on the TCP address laddr and returns a TCP listener.
// Net must be "tcp", "tcp4", or "tcp6".
// If laddr has a port of 0, it means to listen on some available port.
// The caller can use l.Addr() to retrieve the chosen address.
func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err error)
