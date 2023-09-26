// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

const (
	VersionSSL30 = 0x0300
	VersionTLS10 = 0x0301
	VersionTLS11 = 0x0302
	VersionTLS12 = 0x0303
)

// TLS record types.

// TLS handshake message types.

// TLS compression types.

// TLS extension numbers

// TLS signaling cipher suite values

// CurveID is the type of a TLS identifier for an elliptic curve. See
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8
type CurveID uint16

const (
	CurveP256 CurveID = 23
	CurveP384 CurveID = 24
	CurveP521 CurveID = 25
)

// TLS Elliptic Curve Point Formats
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-9

// TLS CertificateStatusType (RFC 3546)

// Certificate types (for certificateRequestMsg)

// Hash functions for TLS 1.2 (See RFC 5246, section A.4.1)

// Signature algorithms for TLS 1.2 (See RFC 5246, section A.4.1)

// signatureAndHash mirrors the TLS 1.2, SignatureAndHashAlgorithm struct. See
// RFC 5246, section A.4.1.

// supportedSKXSignatureAlgorithms contains the signature and hash algorithms
// that the code advertises as supported in a TLS 1.2 ClientHello.

// supportedClientCertSignatureAlgorithms contains the signature and hash
// algorithms that the code advertises as supported in a TLS 1.2
// CertificateRequest.

// ConnectionState records basic TLS details about the connection.
type ConnectionState struct {
	Version                    uint16
	HandshakeComplete          bool
	DidResume                  bool
	CipherSuite                uint16
	NegotiatedProtocol         string
	NegotiatedProtocolIsMutual bool
	ServerName                 string
	PeerCertificates           []*x509.Certificate
	VerifiedChains             [][]*x509.Certificate
}

// ClientAuthType declares the policy the server will follow for
// TLS Client Authentication.
type ClientAuthType int

const (
	NoClientCert ClientAuthType = iota
	RequestClientCert
	RequireAnyClientCert
	VerifyClientCertIfGiven
	RequireAndVerifyClientCert
)

// ClientSessionState contains the state needed by clients to resume TLS
// sessions.
type ClientSessionState struct {
	sessionTicket      []uint8
	vers               uint16
	cipherSuite        uint16
	masterSecret       []byte
	serverCertificates []*x509.Certificate
}

// ClientSessionCache is a cache of ClientSessionState objects that can be used
// by a client to resume a TLS session with a given server. ClientSessionCache
// implementations should expect to be called concurrently from different
// goroutines.
type ClientSessionCache interface {
	Get(sessionKey string) (session *ClientSessionState, ok bool)

	Put(sessionKey string, cs *ClientSessionState)
}

// A Config structure is used to configure a TLS client or server.
// After one has been passed to a TLS function it must not be
// modified. A Config may be reused; the tls package will also not
// modify it.
type Config struct {
	Rand io.Reader

	Time func() time.Time

	Certificates []Certificate

	NameToCertificate map[string]*Certificate

	RootCAs *x509.CertPool

	NextProtos []string

	ServerName string

	ClientAuth ClientAuthType

	ClientCAs *x509.CertPool

	InsecureSkipVerify bool

	CipherSuites []uint16

	PreferServerCipherSuites bool

	SessionTicketsDisabled bool

	SessionTicketKey [32]byte

	ClientSessionCache ClientSessionCache

	MinVersion uint16

	MaxVersion uint16

	CurvePreferences []CurveID

	serverInitOnce sync.Once
}

// BuildNameToCertificate parses c.Certificates and builds c.NameToCertificate
// from the CommonName and SubjectAlternateName fields of each of the leaf
// certificates.
func (c *Config) BuildNameToCertificate()

// A Certificate is a chain of one or more certificates, leaf first.
type Certificate struct {
	Certificate [][]byte
	PrivateKey  crypto.PrivateKey

	OCSPStaple []byte

	Leaf *x509.Certificate
}

// A TLS record.

// lruSessionCache is a ClientSessionCache implementation that uses an LRU
// caching strategy.

// NewLRUClientSessionCache returns a ClientSessionCache with the given
// capacity that uses an LRU strategy. If capacity is < 1, a default capacity
// is used instead.
func NewLRUClientSessionCache(capacity int) ClientSessionCache

// TODO(jsing): Make these available to both crypto/x509 and crypto/tls.
