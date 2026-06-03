// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nettest

import (
	"github.com/shogo82148/std/internal/gate"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/netip"
)

// Listener is an in-memory test implementation of net.Listener.
type Listener struct {
	gate      gate.Gate
	queue     queue[*Conn]
	closed    bool
	acceptErr error
	closeErr  error

	addr     net.Addr
	nextaddr netip.AddrPort
}

// NewListener returns a new Listener.
func NewListener() *Listener

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (li *Listener) Close() error

// Addr returns the listener's network address.
//
// The address is always a *net.TCPAddr.
func (li *Listener) Addr() net.Addr

// SetAddr sets the listener's network address.
func (li *Listener) SetAddr(addr net.Addr)

// NewConn returns a new connection to the listener.
//
// Accept will return the other side of the conn.
func (li *Listener) NewConn() *Conn

// NewConnConfig returns a new connection to the listener.
//
// The function f is called with the new client connection.
// After f returns, Accept will return the other side of the connection.
//
// For example, to create a connection from a specific IP address:
//
//	conn := li.NewConnConfig(func(conn *nettest.Conn) {
//		conn.SetLocalAddr(net.TCPAddrFromAddrPort(netip.MustParseAddrPort("10.0.0.1:1234")))
//	})
func (li *Listener) NewConnConfig(f func(*Conn)) *Conn

// Accept waits for and returns the next connection to the listener.
//
// The connections returned by Accept are always [*Conn]s.
func (li *Listener) Accept() (net.Conn, error)

// SetAcceptError causes any currently blocked and future Accept calls to return
// a net.OpError wrapping err.
// Accept will return any available connections before returning the error,
// including connections created after the error is set.
// A nil error restores the usual behavior.
func (li *Listener) SetAcceptError(err error)

// SetCloseError sets the error returned by Close.
// Close still closes the listener.
// A nil error restores the usual behavior.
func (li *Listener) SetCloseError(err error)
