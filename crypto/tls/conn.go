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
	conn        net.Conn
	isClient    bool
	handshakeFn func() error

	handshakeStatus uint32

	handshakeMutex sync.Mutex
	handshakeErr   error
	vers           uint16
	haveVers       bool
	config         *Config

	handshakes       int
	didResume        bool
	cipherSuite      uint16
	ocspResponse     []byte
	scts             [][]byte
	peerCertificates []*x509.Certificate

	verifiedChains [][]*x509.Certificate

	serverName string

	secureRenegotiation bool

	ekm func(label string, context []byte, length int) ([]byte, error)

	resumptionSecret []byte

	ticketKeys []ticketKey

	clientFinishedIsFirst bool

	closeNotifyErr error

	closeNotifySent bool

	clientFinished [12]byte
	serverFinished [12]byte

	clientProtocol string

	in, out   halfConn
	rawInput  bytes.Buffer
	input     bytes.Reader
	hand      bytes.Buffer
	buffering bool
	sendBuf   []byte

	bytesSent   int64
	packetsSent int64

	retryCount int

	activeCall int32

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

// SetWriteDeadline sets the write deadline on the underlying connection.
// A zero value for t means Write will not time out.
// After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.
func (c *Conn) SetWriteDeadline(t time.Time) error

// A halfConn represents one direction of the record layer
// connection, either sending or receiving.

// cbcMode is an interface for block ciphers using cipher block chaining.

// RecordHeaderError is returned when a TLS record header is invalid.
type RecordHeaderError struct {
	Msg string

	RecordHeader [5]byte

	Conn net.Conn
}

func (e RecordHeaderError) Error() string

// atLeastReader reads from R, stopping with EOF once at least N bytes have been
// read. It is different from an io.LimitedReader in that it doesn't cut short
// the last Read call, and in that it considers an early EOF an error.

// outBufPool pools the record-sized scratch buffers used by writeRecordLocked.

// Write writes data to the connection.
//
// As Write calls Handshake, in order to prevent indefinite blocking a deadline
// must be set for both Read and Write before Write is called when the handshake
// has not yet completed. See SetDeadline, SetReadDeadline, and
// SetWriteDeadline.
func (c *Conn) Write(b []byte) (int, error)

// Read reads data from the connection.
//
// As Read calls Handshake, in order to prevent indefinite blocking a deadline
// must be set for both Read and Write before Read is called when the handshake
// has not yet completed. See SetDeadline, SetReadDeadline, and
// SetWriteDeadline.
func (c *Conn) Read(b []byte) (int, error)

// Close closes the connection.
func (c *Conn) Close() error

// CloseWrite shuts down the writing side of the connection. It should only be
// called once the handshake has completed and does not call CloseWrite on the
// underlying connection. Most callers should just use Close.
func (c *Conn) CloseWrite() error

// Handshake runs the client or server handshake
// protocol if it has not yet been run.
//
// Most uses of this package need not call Handshake explicitly: the
// first Read or Write will call it automatically.
//
// For control over canceling or setting a timeout on a handshake, use
// the Dialer's DialContext method.
func (c *Conn) Handshake() error

// ConnectionState returns basic TLS details about the connection.
func (c *Conn) ConnectionState() ConnectionState

// OCSPResponse returns the stapled OCSP response from the TLS server, if
// any. (Only valid for client connections.)
func (c *Conn) OCSPResponse() []byte

// VerifyHostname checks that the peer certificate chain is valid for
// connecting to host. If so, it returns nil; if not, it returns an error
// describing the problem.
func (c *Conn) VerifyHostname(host string) error
