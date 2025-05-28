// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TLS low level connection and record layer

package tls

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// A Conn represents a secured connection.
// It implements the net.Conn interface.
type Conn struct {
	// constant
	conn        net.Conn
	isClient    bool
	handshakeFn func(context.Context) error
	quic        *quicState

	// isHandshakeComplete is true if the connection is currently transferring
	// application data (i.e. is not currently processing a handshake).
	// isHandshakeComplete is true implies handshakeErr == nil.
	isHandshakeComplete atomic.Bool
	// constant after handshake; protected by handshakeMutex
	handshakeMutex sync.Mutex
	handshakeErr   error
	vers           uint16
	haveVers       bool
	config         *Config
	// handshakes counts the number of handshakes performed on the
	// connection so far. If renegotiation is disabled then this is either
	// zero or one.
	handshakes       int
	extMasterSecret  bool
	didResume        bool
	didHRR           bool
	cipherSuite      uint16
	curveID          CurveID
	peerSigAlg       SignatureScheme
	ocspResponse     []byte
	scts             [][]byte
	peerCertificates []*x509.Certificate
	// verifiedChains contains the certificate chains that we built, as
	// opposed to the ones presented by the server.
	verifiedChains [][]*x509.Certificate
	// serverName contains the server name indicated by the client, if any.
	serverName string
	// secureRenegotiation is true if the server echoed the secure
	// renegotiation extension. (This is meaningless as a server because
	// renegotiation is not supported in that case.)
	secureRenegotiation bool
	// ekm is a closure for exporting keying material.
	ekm func(label string, context []byte, length int) ([]byte, error)
	// resumptionSecret is the resumption_master_secret for handling
	// or sending NewSessionTicket messages.
	resumptionSecret []byte
	echAccepted      bool

	// ticketKeys is the set of active session ticket keys for this
	// connection. The first one is used to encrypt new tickets and
	// all are tried to decrypt tickets.
	ticketKeys []ticketKey

	// clientFinishedIsFirst is true if the client sent the first Finished
	// message during the most recent handshake. This is recorded because
	// the first transmitted Finished message is the tls-unique
	// channel-binding value.
	clientFinishedIsFirst bool

	// closeNotifyErr is any error from sending the alertCloseNotify record.
	closeNotifyErr error
	// closeNotifySent is true if the Conn attempted to send an
	// alertCloseNotify record.
	closeNotifySent bool

	// clientFinished and serverFinished contain the Finished message sent
	// by the client or server in the most recent handshake. This is
	// retained to support the renegotiation extension and tls-unique
	// channel-binding.
	clientFinished [12]byte
	serverFinished [12]byte

	// clientProtocol is the negotiated ALPN protocol.
	clientProtocol string

	// input/output
	in, out   halfConn
	rawInput  bytes.Buffer
	input     bytes.Reader
	hand      bytes.Buffer
	buffering bool
	sendBuf   []byte

	// bytesSent counts the bytes of application data sent.
	// packetsSent counts packets.
	bytesSent   int64
	packetsSent int64

	// retryCount counts the number of consecutive non-advancing records
	// received by Conn.readRecord. That is, records that neither advance the
	// handshake, nor deliver application data. Protected by in.Mutex.
	retryCount int

	// activeCall indicates whether Close has been call in the low bit.
	// the rest of the bits are the number of goroutines in Conn.Write.
	activeCall atomic.Int32

	tmp [16]byte
}

// LocalAddr returns the local network address.
func (c *Conn) LocalAddr() net.Addr

// RemoteAddr returns the remote network address.
func (c *Conn) RemoteAddr() net.Addr

// SetDeadline sets the read and write deadlines associated with the connection.
// A zero value for t means [Conn.Read] and [Conn.Write] will not time out.
// After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.
func (c *Conn) SetDeadline(t time.Time) error

// SetReadDeadline sets the read deadline on the underlying connection.
// A zero value for t means [Conn.Read] will not time out.
func (c *Conn) SetReadDeadline(t time.Time) error

// SetWriteDeadline sets the write deadline on the underlying connection.
// A zero value for t means [Conn.Write] will not time out.
// After a [Conn.Write] has timed out, the TLS state is corrupt and all future writes will return the same error.
func (c *Conn) SetWriteDeadline(t time.Time) error

// NetConn returns the underlying connection that is wrapped by c.
// Note that writing to or reading from this connection directly will corrupt the
// TLS session.
func (c *Conn) NetConn() net.Conn

// RecordHeaderError is returned when a TLS record header is invalid.
type RecordHeaderError struct {
	// Msg contains a human readable string that describes the error.
	Msg string
	// RecordHeader contains the five bytes of TLS record header that
	// triggered the error.
	RecordHeader [5]byte
	// Conn provides the underlying net.Conn in the case that a client
	// sent an initial handshake that didn't look like TLS.
	// It is nil if there's already been a handshake or a TLS alert has
	// been written to the connection.
	Conn net.Conn
}

func (e RecordHeaderError) Error() string

// Write writes data to the connection.
//
// As Write calls [Conn.Handshake], in order to prevent indefinite blocking a deadline
// must be set for both [Conn.Read] and Write before Write is called when the handshake
// has not yet completed. See [Conn.SetDeadline], [Conn.SetReadDeadline], and
// [Conn.SetWriteDeadline].
func (c *Conn) Write(b []byte) (int, error)

// Read reads data from the connection.
//
// As Read calls [Conn.Handshake], in order to prevent indefinite blocking a deadline
// must be set for both Read and [Conn.Write] before Read is called when the handshake
// has not yet completed. See [Conn.SetDeadline], [Conn.SetReadDeadline], and
// [Conn.SetWriteDeadline].
func (c *Conn) Read(b []byte) (int, error)

// Close closes the connection.
func (c *Conn) Close() error

// CloseWrite shuts down the writing side of the connection. It should only be
// called once the handshake has completed and does not call CloseWrite on the
// underlying connection. Most callers should just use [Conn.Close].
func (c *Conn) CloseWrite() error

// Handshake runs the client or server handshake
// protocol if it has not yet been run.
//
// Most uses of this package need not call Handshake explicitly: the
// first [Conn.Read] or [Conn.Write] will call it automatically.
//
// For control over canceling or setting a timeout on a handshake, use
// [Conn.HandshakeContext] or the [Dialer]'s DialContext method instead.
//
// In order to avoid denial of service attacks, the maximum RSA key size allowed
// in certificates sent by either the TLS server or client is limited to 8192
// bits. This limit can be overridden by setting tlsmaxrsasize in the GODEBUG
// environment variable (e.g. GODEBUG=tlsmaxrsasize=4096).
func (c *Conn) Handshake() error

// HandshakeContext runs the client or server handshake
// protocol if it has not yet been run.
//
// The provided Context must be non-nil. If the context is canceled before
// the handshake is complete, the handshake is interrupted and an error is returned.
// Once the handshake has completed, cancellation of the context will not affect the
// connection.
//
// Most uses of this package need not call HandshakeContext explicitly: the
// first [Conn.Read] or [Conn.Write] will call it automatically.
func (c *Conn) HandshakeContext(ctx context.Context) error

// ConnectionState returns basic TLS details about the connection.
func (c *Conn) ConnectionState() ConnectionState

// OCSPResponse returns the stapled OCSP response from the TLS server, if
// any. (Only valid for client connections.)
func (c *Conn) OCSPResponse() []byte

// VerifyHostname checks that the peer certificate chain is valid for
// connecting to host. If so, it returns nil; if not, it returns an error
// describing the problem.
func (c *Conn) VerifyHostname(host string) error
