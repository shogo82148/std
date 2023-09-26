// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/context"
)

// QUICEncryptionLevel represents a QUIC encryption level used to transmit
// handshake messages.
type QUICEncryptionLevel int

const (
	QUICEncryptionLevelInitial = QUICEncryptionLevel(iota)
	QUICEncryptionLevelEarly
	QUICEncryptionLevelHandshake
	QUICEncryptionLevelApplication
)

func (l QUICEncryptionLevel) String() string

// A QUICConn represents a connection which uses a QUIC implementation as the underlying
// transport as described in RFC 9001.
//
// Methods of QUICConn are not safe for concurrent use.
type QUICConn struct {
	conn *Conn

	sessionTicketSent bool
}

// A QUICConfig configures a QUICConn.
type QUICConfig struct {
	TLSConfig *Config
}

// A QUICEventKind is a type of operation on a QUIC connection.
type QUICEventKind int

const (
	// QUICNoEvent indicates that there are no events available.
	QUICNoEvent QUICEventKind = iota

	// QUICSetReadSecret and QUICSetWriteSecret provide the read and write
	// secrets for a given encryption level.
	// QUICEvent.Level, QUICEvent.Data, and QUICEvent.Suite are set.
	//
	// Secrets for the Initial encryption level are derived from the initial
	// destination connection ID, and are not provided by the QUICConn.
	QUICSetReadSecret
	QUICSetWriteSecret

	// QUICWriteData provides data to send to the peer in CRYPTO frames.
	// QUICEvent.Data is set.
	QUICWriteData

	// QUICTransportParameters provides the peer's QUIC transport parameters.
	// QUICEvent.Data is set.
	QUICTransportParameters

	// QUICTransportParametersRequired indicates that the caller must provide
	// QUIC transport parameters to send to the peer. The caller should set
	// the transport parameters with QUICConn.SetTransportParameters and call
	// QUICConn.NextEvent again.
	//
	// If transport parameters are set before calling QUICConn.Start, the
	// connection will never generate a QUICTransportParametersRequired event.
	QUICTransportParametersRequired

	// QUICRejectedEarlyData indicates that the server rejected 0-RTT data even
	// if we offered it. It's returned before QUICEncryptionLevelApplication
	// keys are returned.
	QUICRejectedEarlyData

	// QUICHandshakeDone indicates that the TLS handshake has completed.
	QUICHandshakeDone
)

// A QUICEvent is an event occurring on a QUIC connection.
//
// The type of event is specified by the Kind field.
// The contents of the other fields are kind-specific.
type QUICEvent struct {
	Kind QUICEventKind

	Level QUICEncryptionLevel

	Data []byte

	Suite uint16
}

// QUICClient returns a new TLS client side connection using QUICTransport as the
// underlying transport. The config cannot be nil.
//
// The config's MinVersion must be at least TLS 1.3.
func QUICClient(config *QUICConfig) *QUICConn

// QUICServer returns a new TLS server side connection using QUICTransport as the
// underlying transport. The config cannot be nil.
//
// The config's MinVersion must be at least TLS 1.3.
func QUICServer(config *QUICConfig) *QUICConn

// Start starts the client or server handshake protocol.
// It may produce connection events, which may be read with NextEvent.
//
// Start must be called at most once.
func (q *QUICConn) Start(ctx context.Context) error

// NextEvent returns the next event occurring on the connection.
// It returns an event with a Kind of QUICNoEvent when no events are available.
func (q *QUICConn) NextEvent() QUICEvent

// Close closes the connection and stops any in-progress handshake.
func (q *QUICConn) Close() error

// HandleData handles handshake bytes received from the peer.
// It may produce connection events, which may be read with NextEvent.
func (q *QUICConn) HandleData(level QUICEncryptionLevel, data []byte) error

type QUICSessionTicketOptions struct {
	EarlyData bool
}

// SendSessionTicket sends a session ticket to the client.
// It produces connection events, which may be read with NextEvent.
// Currently, it can only be called once.
func (q *QUICConn) SendSessionTicket(opts QUICSessionTicketOptions) error

// ConnectionState returns basic TLS details about the connection.
func (q *QUICConn) ConnectionState() ConnectionState

// SetTransportParameters sets the transport parameters to send to the peer.
//
// Server connections may delay setting the transport parameters until after
// receiving the client's transport parameters. See QUICTransportParametersRequired.
func (q *QUICConn) SetTransportParameters(params []byte)
