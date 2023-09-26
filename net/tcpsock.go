// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// TCPAddr represents the address of a TCP end point.
type TCPAddr struct {
	IP   IP
	Port int
	Zone string
}

// Network returns the address's network name, "tcp".
func (a *TCPAddr) Network() string

func (a *TCPAddr) String() string

// ResolveTCPAddr returns an address of TCP end point.
//
// The network must be a TCP network name.
//
// If the host in the address parameter is not a literal IP address or
// the port is not a literal port number, ResolveTCPAddr resolves the
// address to an address of TCP end point.
// Otherwise, it parses the address as a pair of literal IP address
// and port number.
// The address parameter can use a host name, but this is not
// recommended, because it will return at most one of the host name's
// IP addresses.
//
// See func Dial for a description of the network and address
// parameters.
func ResolveTCPAddr(network, address string) (*TCPAddr, error)

// TCPConn is an implementation of the Conn interface for TCP network
// connections.
type TCPConn struct {
	conn
}

// SyscallConn returns a raw network connection.
// This implements the syscall.Conn interface.
func (c *TCPConn) SyscallConn() (syscall.RawConn, error)

// ReadFrom implements the io.ReaderFrom ReadFrom method.
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)

// CloseRead shuts down the reading side of the TCP connection.
// Most callers should just use Close.
func (c *TCPConn) CloseRead() error

// CloseWrite shuts down the writing side of the TCP connection.
// Most callers should just use Close.
func (c *TCPConn) CloseWrite() error

// SetLinger sets the behavior of Close on a connection which still
// has data waiting to be sent or to be acknowledged.
//
// If sec < 0 (the default), the operating system finishes sending the
// data in the background.
//
// If sec == 0, the operating system discards any unsent or
// unacknowledged data.
//
// If sec > 0, the data is sent in the background as with sec < 0. On
// some operating systems after sec seconds have elapsed any remaining
// unsent data may be discarded.
func (c *TCPConn) SetLinger(sec int) error

// SetKeepAlive sets whether the operating system should send
// keepalive messages on the connection.
func (c *TCPConn) SetKeepAlive(keepalive bool) error

// SetKeepAlivePeriod sets period between keep alives.
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error

// SetNoDelay controls whether the operating system should delay
// packet transmission in hopes of sending fewer packets (Nagle's
// algorithm).  The default is true (no delay), meaning that data is
// sent as soon as possible after a Write.
func (c *TCPConn) SetNoDelay(noDelay bool) error

// DialTCP acts like Dial for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)

// TCPListener is a TCP network listener. Clients should typically
// use variables of type Listener instead of assuming TCP.
type TCPListener struct {
	fd *netFD
}

// SyscallConn returns a raw network connection.
// This implements the syscall.Conn interface.
//
// The returned RawConn only supports calling Control. Read and
// Write return an error.
func (l *TCPListener) SyscallConn() (syscall.RawConn, error)

// AcceptTCP accepts the next incoming call and returns the new
// connection.
func (l *TCPListener) AcceptTCP() (*TCPConn, error)

// Accept implements the Accept method in the Listener interface; it
// waits for the next call and returns a generic Conn.
func (l *TCPListener) Accept() (Conn, error)

// Close stops listening on the TCP address.
// Already Accepted connections are not closed.
func (l *TCPListener) Close() error

// Addr returns the listener's network address, a *TCPAddr.
// The Addr returned is shared by all invocations of Addr, so
// do not modify it.
func (l *TCPListener) Addr() Addr

// SetDeadline sets the deadline associated with the listener.
// A zero time value disables the deadline.
func (l *TCPListener) SetDeadline(t time.Time) error

// File returns a copy of the underlying os.File, set to blocking
// mode. It is the caller's responsibility to close f when finished.
// Closing l does not affect f, and closing f does not affect l.
//
// The returned os.File's file descriptor is different from the
// connection's. Attempting to change properties of the original
// using this duplicate may or may not have the desired effect.
func (l *TCPListener) File() (f *os.File, err error)

// ListenTCP acts like Listen for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If the IP field of laddr is nil or an unspecified IP address,
// ListenTCP listens on all available unicast and anycast IP addresses
// of the local system.
// If the Port field of laddr is 0, a port number is automatically
// chosen.
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
