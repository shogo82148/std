// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TLS low level connection and record layer

package tls

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

// A Conn represents a secured connection.
// It implements the net.Conn interface.
type Conn struct {
	conn     net.Conn
	isClient bool

	handshakeMutex    sync.Mutex
	vers              uint16
	haveVers          bool
	config            *Config
	handshakeComplete bool
	didResume         bool
	cipherSuite       uint16
	ocspResponse      []byte
	peerCertificates  []*x509.Certificate

	verifiedChains [][]*x509.Certificate

	serverName string

	clientProtocol         string
	clientProtocolFallback bool

	connErr

	in, out  halfConn
	rawInput *block
	input    *block
	hand     bytes.Buffer

	tmp [16]byte
}

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr

// SetDeadline sets the read and write deadlines associated with the connection.
// A zero value for t means Read and Write will not time out.
// After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.
func (c *Conn) SetDeadline(t time.Time) error

// SetReadDeadline sets the read deadline on the underlying connection.
// A zero value for t means Read will not time out.
func (c *Conn) SetReadDeadline(t time.Time) error

// SetWriteDeadline sets the write deadline on the underlying conneciton.
// A zero value for t means Write will not time out.
// After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.
func (c *Conn) SetWriteDeadline(t time.Time) error

// A halfConn represents one direction of the record layer
// connection, either sending or receiving.

// A block is a simple data buffer.

// Write writes data to the connection.
func (c *Conn) Write(b []byte) (int, error)

// Read can be made to time out and return a net.Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (c *Conn) Read(b []byte) (n int, err error)

// Close closes the connection.
func (c *Conn) Close() error

// Handshake runs the client or server handshake
// protocol if it has not yet been run.
// Most uses of this package need not call Handshake
// explicitly: the first Read or Write will call it automatically.
func (c *Conn) Handshake() error

// ConnectionState returns basic TLS details about the connection.
func (c *Conn) ConnectionState() ConnectionState

// OCSPResponse returns the stapled OCSP response from the TLS server, if
// any. (Only valid for client connections.)
func (c *Conn) OCSPResponse() []byte

// VerifyHostname checks that the peer certificate chain is valid for
// connecting to host.  If so, it returns nil; if not, it returns an error
// describing the problem.
func (c *Conn) VerifyHostname(host string) error
