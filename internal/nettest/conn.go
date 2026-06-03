// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nettest

import (
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/time"
)

// Conn is an in-memory test implementation of net.Conn.
type Conn struct {
	// Conns come in pairs.
	// Writes to one Conn are read by its peer, and vice-versa.
	//
	// A connHalf handles one direction of data flow.
	// A Conn consists of read and write halves.
	// A Conn's peer has the same halves, only swapped.
	//
	// A Conn reads from r and writes to w.
	r, w *connHalf

	// peer is the other endpoint.
	peer *Conn
}

// NewConnPair returns a pair of connected Conns.
func NewConnPair() (*Conn, *Conn)

// Peer returns the other end of the connection.
func (c *Conn) Peer() *Conn

// Read reads data from the connection.
func (c *Conn) Read(b []byte) (n int, err error)

// CanRead reports whether Read can proceed without blocking.
func (c *Conn) CanRead() bool

// Write writes data to the connection.
func (c *Conn) Write(b []byte) (n int, err error)

// IsClosed reports whether the connection has been closed.
// A connection is closed if [CloseRead] and [CloseWrite] are both called,
// or if [Close] is called.
//
// To identify when the other side of the Conn has been closed,
// use Conn.Peer().IsClosed().
func (c *Conn) IsClosed() bool

// CloseRead shuts down the reading side of the connection.
func (c *Conn) CloseRead() error

// CloseWrite shuts down the writing side of the connection.
func (c *Conn) CloseWrite() error

// Close closes the connection.
func (c *Conn) Close() error

// SetCloseError sets the error returned by Close.
// Close still closes the connection.
// A nil error restores the usual behavior.
func (c *Conn) SetCloseError(err error)

// LocalAddr returns the (fake) local network address.
func (c *Conn) LocalAddr() net.Addr

// SetLocalAddr sets the local address.
//
// To set the remote address, set the local address of Conn's peer.
func (c *Conn) SetLocalAddr(addr net.Addr)

// LocalAddr returns the (fake) remote network address.
func (c *Conn) RemoteAddr() net.Addr

// SetDeadline sets the read and write deadlines for the connection.
func (c *Conn) SetDeadline(t time.Time) error

// SetReadDeadline sets the read deadline for the connection.
func (c *Conn) SetReadDeadline(t time.Time) error

// SetWriteDeadline sets the write deadline for the connection.
func (c *Conn) SetWriteDeadline(t time.Time) error

// SetReadBufferSize sets the connection's read buffer.
// Writes to the other end of the connection will block so long as the buffer is full.
// Setting the size to 0 blocks all writes until the size is increased.
func (c *Conn) SetReadBufferSize(size int)

// SetReadError causes any currently blocked and future Read calls to return
// a net.OpError wrapping err. It does not affect the other half of the connection.
// Reads will return any buffered data before returning the error,
// including data written after the error is set and io.EOF after the other end is closed.
// A nil error restores the usual behavior.
func (c *Conn) SetReadError(err error)

// SetWriteError causes any currently blocked and future Write calls to return
// a net.OpError wrapping err. It does not affect the other half of the connection.
// Writes will not write data to the connection buffer while an error is set.
// A nil error restores the usual behavior.
func (c *Conn) SetWriteError(err error)
