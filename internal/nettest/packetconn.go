// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nettest

import (
	"github.com/shogo82148/std/internal/gate"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

// A PacketNet is a group of communicating [PacketConn]s.
type PacketNet struct {
	mu    sync.Mutex
	conns map[netAddr]*PacketConn
}

// NewPacketNet returns a new PacketNet.
func NewPacketNet() *PacketNet

// NewConn returns a new [PacketConn] listening on the given address.
// It returns an error if there is an existing listener on this address.
func (n *PacketNet) NewConn(a net.Addr) (*PacketConn, error)

type PacketConn struct {
	gate         gate.Gate
	queue        queue[*packet]
	closed       bool
	readErr      error
	writeErr     error
	closeErr     error
	readDeadline connDeadline

	net  *PacketNet
	addr net.Addr
}

// ReadFrom reads a packet from the connection, copying the payload into b.
func (p *PacketConn) ReadFrom(b []byte) (n int, addr net.Addr, err error)

// WriteTo writes a packet with payload b to addr.
// addr must be a [*net.UDPAddr].
//
// WriteTo appends the packet to the recipient's receive buffer.
// If no recipient is listening on addr or if the recipient's
// receive buffer is full, the packet is silently discarded.
func (p *PacketConn) WriteTo(b []byte, addr net.Addr) (n int, err error)

// Close closes the connection.
func (p *PacketConn) Close() error

// LocalAddr returns the (fake) local network address.
func (p *PacketConn) LocalAddr() net.Addr

// SetReadDeadline sets the read deadline for the connection.
// PacketConns have no write deadline.
func (p *PacketConn) SetDeadline(t time.Time) error

// SetReadDeadline sets the read deadline for the connection.
func (p *PacketConn) SetReadDeadline(t time.Time) error

// SetWriteDeadline has no effect.
// Writes to PacketConns never block.
func (p *PacketConn) SetWriteDeadline(t time.Time) error

// SetReadError causes any currently blocked and future ReadFrom calls to return
// a net.OpError wrapping err. It does not affect the other half of the connection.
// Reads will return any buffered data before returning the error,
// including data written after the error is set.
// A nil error restores the usual behavior.
func (c *PacketConn) SetReadError(err error)

// SetWriteError causes any currently blocked and future WriteTo calls to return
// a net.OpError wrapping err. It does not affect the other half of the connection.
// Writes will not write data while an error is set.
// A nil error restores the usual behavior.
func (c *PacketConn) SetWriteError(err error)

// SetCloseError sets the error returned by Close.
// Close still closes the connection.
// A nil error restores the usual behavior.
func (c *PacketConn) SetCloseError(err error)

// CanRead reports whether [ReadFrom] can return at least one byte or an error.
// If [ReadFrom] would block, CanRead returns false.
func (p *PacketConn) CanRead() bool

// IsClosed reports whether the connection has been closed.
func (p *PacketConn) IsClosed() bool
