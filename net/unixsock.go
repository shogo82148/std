// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// UnixAddr represents the address of a Unix domain socket end point.
type UnixAddr struct {
	Name string
	Net  string
}

// Network returns the address's network name, "unix", "unixgram" or
// "unixpacket".
func (a *UnixAddr) Network() string

func (a *UnixAddr) String() string

// ResolveUnixAddr returns an address of Unix domain socket end point.
//
// The network must be a Unix network name.
//
// See func [Dial] for a description of the network and address
// parameters.
func ResolveUnixAddr(network, address string) (*UnixAddr, error)

// UnixConn is an implementation of the [Conn] interface for connections
// to Unix domain sockets.
type UnixConn struct {
	conn
}

// SyscallConn returns a raw network connection.
// This implements the [syscall.Conn] interface.
func (c *UnixConn) SyscallConn() (syscall.RawConn, error)

// CloseRead shuts down the reading side of the Unix domain connection.
// Most callers should just use [UnixConn.Close].
func (c *UnixConn) CloseRead() error

// CloseWrite shuts down the writing side of the Unix domain connection.
// Most callers should just use [UnixConn.Close].
func (c *UnixConn) CloseWrite() error

// ReadFromUnix acts like [UnixConn.ReadFrom] but returns a [UnixAddr].
func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)

// ReadFrom implements the [PacketConn].ReadFrom method.
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)

// ReadMsgUnix reads a message from c, copying the payload into b and
// the associated out-of-band data into oob. It returns the number of
// bytes copied into b, the number of bytes copied into oob, the flags
// that were set on the message and the source address of the message.
//
// Note that if len(b) == 0 and len(oob) > 0, this function will still
// read (and discard) 1 byte from the connection.
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)

// WriteToUnix acts like [UnixConn.WriteTo] but takes a [UnixAddr].
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)

// WriteTo implements the [PacketConn].WriteTo method.
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)

// WriteMsgUnix writes a message to addr via c, copying the payload
// from b and the associated out-of-band data from oob. It returns the
// number of payload and out-of-band bytes written.
//
// Note that if len(b) == 0 and len(oob) > 0, this function will still
// write 1 byte to the connection.
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)

// DialUnix acts like [Dial] for Unix networks.
//
// The network must be a Unix network name; see func [Dial] for details.
//
// If laddr is non-nil, it is used as the local address for the
// connection.
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)

// UnixListener is a Unix domain socket listener. Clients should
// typically use variables of type [Listener] instead of assuming Unix
// domain sockets.
type UnixListener struct {
	fd         *netFD
	path       string
	unlink     bool
	unlinkOnce sync.Once
}

// SyscallConn returns a raw network connection.
// This implements the [syscall.Conn] interface.
//
// The returned [syscall.RawConn] only supports calling Control. Read and
// Write return an error.
func (l *UnixListener) SyscallConn() (syscall.RawConn, error)

// AcceptUnix accepts the next incoming call and returns the new
// connection.
func (l *UnixListener) AcceptUnix() (*UnixConn, error)

// Accept implements the Accept method in the [Listener] interface.
// Returned connections will be of type [*UnixConn].
func (l *UnixListener) Accept() (Conn, error)

// Close stops listening on the Unix address. Already accepted
// connections are not closed.
func (l *UnixListener) Close() error

// Addr returns the listener's network address.
// The [Addr] returned is shared by all invocations of Addr, so
// do not modify it.
func (l *UnixListener) Addr() Addr

// SetDeadline sets the deadline associated with the listener.
// A zero time value disables the deadline.
func (l *UnixListener) SetDeadline(t time.Time) error

// File returns a copy of the underlying [os.File].
// It is the caller's responsibility to close f when finished.
// Closing l does not affect f, and closing f does not affect l.
//
// The returned [os.File]'s file descriptor is different from the
// connection's. Attempting to change properties of the original
// using this duplicate may or may not have the desired effect.
func (l *UnixListener) File() (f *os.File, err error)

// ListenUnix acts like [Listen] for Unix networks.
//
// The network must be "unix" or "unixpacket".
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)

// ListenUnixgram acts like [ListenPacket] for Unix networks.
//
// The network must be "unixgram".
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
