// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/crypto/x509"
)

// A SessionState is a resumable session.
type SessionState struct {

	// Extra is ignored by crypto/tls, but is encoded by [SessionState.Bytes]
	// and parsed by [ParseSessionState].
	//
	// This allows [Config.UnwrapSession]/[Config.WrapSession] and
	// [ClientSessionCache] implementations to store and retrieve additional
	// data alongside this session.
	//
	// To allow different layers in a protocol stack to share this field,
	// applications must only append to it, not replace it, and must use entries
	// that can be recognized even if out of order (for example, by starting
	// with an id and version prefix).
	Extra [][]byte

	// EarlyData indicates whether the ticket can be used for 0-RTT in a QUIC
	// connection. The application may set this to false if it is true to
	// decline to offer 0-RTT even if supported.
	EarlyData bool

	version     uint16
	isClient    bool
	cipherSuite uint16
	// createdAt is the generation time of the secret on the sever (which for
	// TLS 1.0–1.2 might be earlier than the current session) and the time at
	// which the ticket was received on the client.
	createdAt         uint64
	secret            []byte
	extMasterSecret   bool
	peerCertificates  []*x509.Certificate
	activeCertHandles []*activeCert
	ocspResponse      []byte
	scts              [][]byte
	verifiedChains    [][]*x509.Certificate
	alpnProtocol      string

	// Client-side TLS 1.3-only fields.
	useBy  uint64
	ageAdd uint32
}

// Bytes encodes the session, including any private fields, so that it can be
// parsed by [ParseSessionState]. The encoding contains secret values critical
// to the security of future and possibly past sessions.
//
// The specific encoding should be considered opaque and may change incompatibly
// between Go versions.
func (s *SessionState) Bytes() ([]byte, error)

// ParseSessionState parses a [SessionState] encoded by [SessionState.Bytes].
func ParseSessionState(data []byte) (*SessionState, error)

// EncryptTicket encrypts a ticket with the Config's configured (or default)
// session ticket keys. It can be used as a [Config.WrapSession] implementation.
func (c *Config) EncryptTicket(cs ConnectionState, ss *SessionState) ([]byte, error)

// DecryptTicket decrypts a ticket encrypted by [Config.EncryptTicket]. It can
// be used as a [Config.UnwrapSession] implementation.
//
// If the ticket can't be decrypted or parsed, DecryptTicket returns (nil, nil).
func (c *Config) DecryptTicket(identity []byte, cs ConnectionState) (*SessionState, error)

// ClientSessionState contains the state needed by a client to
// resume a previous TLS session.
type ClientSessionState struct {
	ticket  []byte
	session *SessionState
}

// ResumptionState returns the session ticket sent by the server (also known as
// the session's identity) and the state necessary to resume this session.
//
// It can be called by [ClientSessionCache.Put] to serialize (with
// [SessionState.Bytes]) and store the session.
func (cs *ClientSessionState) ResumptionState() (ticket []byte, state *SessionState, err error)

// NewResumptionState returns a state value that can be returned by
// [ClientSessionCache.Get] to resume a previous session.
//
// state needs to be returned by [ParseSessionState], and the ticket and session
// state must have been returned by [ClientSessionState.ResumptionState].
func NewResumptionState(ticket []byte, state *SessionState) (*ClientSessionState, error)
