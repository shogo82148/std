// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
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
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8
type CurveID uint16

const (
	CurveP256 CurveID = 23
	CurveP384 CurveID = 24
	CurveP521 CurveID = 25
	X25519    CurveID = 29
)

// TLS Elliptic Curve Point Formats
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-9

// TLS CertificateStatusType (RFC 3546)

// Certificate types (for certificateRequestMsg)

// Signature algorithms (for internal signaling use). Starting at 16 to avoid overlap with
// TLS 1.2 codepoints (RFC 5246, section A.4.1), with which these have nothing to do.

// supportedSignatureAlgorithms contains the signature and hash algorithms that
// the code advertises as supported in a TLS 1.2 ClientHello and in a TLS 1.2
// CertificateRequest. The two fields are merged to match with TLS 1.3.
// Note that in TLS 1.2, the ECDSA algorithms are not constrained to P-256, etc.

// ConnectionState records basic TLS details about the connection.
type ConnectionState struct {
	Version                     uint16
	HandshakeComplete           bool
	DidResume                   bool
	CipherSuite                 uint16
	NegotiatedProtocol          string
	NegotiatedProtocolIsMutual  bool
	ServerName                  string
	PeerCertificates            []*x509.Certificate
	VerifiedChains              [][]*x509.Certificate
	SignedCertificateTimestamps [][]byte
	OCSPResponse                []byte

	ekm func(label string, context []byte, length int) ([]byte, error)

	TLSUnique []byte
}

// ExportKeyingMaterial returns length bytes of exported key material in a new
// slice as defined in https://tools.ietf.org/html/rfc5705. If context is nil,
// it is not used as part of the seed. If the connection was set to allow
// renegotiation via Config.Renegotiation, this function will return an error.
func (cs *ConnectionState) ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error)

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
	verifiedChains     [][]*x509.Certificate
}

// ClientSessionCache is a cache of ClientSessionState objects that can be used
// by a client to resume a TLS session with a given server. ClientSessionCache
// implementations should expect to be called concurrently from different
// goroutines. Only ticket-based resumption is supported, not SessionID-based
// resumption.
type ClientSessionCache interface {
	Get(sessionKey string) (session *ClientSessionState, ok bool)

	Put(sessionKey string, cs *ClientSessionState)
}

// SignatureScheme identifies a signature algorithm supported by TLS. See
// https://tools.ietf.org/html/draft-ietf-tls-tls13-18#section-4.2.3.
type SignatureScheme uint16

const (
	PKCS1WithSHA1   SignatureScheme = 0x0201
	PKCS1WithSHA256 SignatureScheme = 0x0401
	PKCS1WithSHA384 SignatureScheme = 0x0501
	PKCS1WithSHA512 SignatureScheme = 0x0601

	PSSWithSHA256 SignatureScheme = 0x0804
	PSSWithSHA384 SignatureScheme = 0x0805
	PSSWithSHA512 SignatureScheme = 0x0806

	ECDSAWithP256AndSHA256 SignatureScheme = 0x0403
	ECDSAWithP384AndSHA384 SignatureScheme = 0x0503
	ECDSAWithP521AndSHA512 SignatureScheme = 0x0603

	// Legacy signature and hash algorithms for TLS 1.2.
	ECDSAWithSHA1 SignatureScheme = 0x0203
)

// ClientHelloInfo contains information from a ClientHello message in order to
// guide certificate selection in the GetCertificate callback.
type ClientHelloInfo struct {
	CipherSuites []uint16

	ServerName string

	SupportedCurves []CurveID

	SupportedPoints []uint8

	SignatureSchemes []SignatureScheme

	SupportedProtos []string

	SupportedVersions []uint16

	Conn net.Conn
}

// CertificateRequestInfo contains information from a server's
// CertificateRequest message, which is used to demand a certificate and proof
// of control from a client.
type CertificateRequestInfo struct {
	AcceptableCAs [][]byte

	SignatureSchemes []SignatureScheme
}

// RenegotiationSupport enumerates the different levels of support for TLS
// renegotiation. TLS renegotiation is the act of performing subsequent
// handshakes on a connection after the first. This significantly complicates
// the state machine and has been the source of numerous, subtle security
// issues. Initiating a renegotiation is not supported, but support for
// accepting renegotiation requests may be enabled.
//
// Even when enabled, the server may not change its identity between handshakes
// (i.e. the leaf certificate must be the same). Additionally, concurrent
// handshake and application data flow is not permitted so renegotiation can
// only be used with protocols that synchronise with the renegotiation, such as
// HTTPS.
type RenegotiationSupport int

const (
	// RenegotiateNever disables renegotiation.
	RenegotiateNever RenegotiationSupport = iota

	// RenegotiateOnceAsClient allows a remote server to request
	// renegotiation once per connection.
	RenegotiateOnceAsClient

	// RenegotiateFreelyAsClient allows a remote server to repeatedly
	// request renegotiation.
	RenegotiateFreelyAsClient
)

// A Config structure is used to configure a TLS client or server.
// After one has been passed to a TLS function it must not be
// modified. A Config may be reused; the tls package will also not
// modify it.
type Config struct {
	Rand io.Reader

	Time func() time.Time

	Certificates []Certificate

	NameToCertificate map[string]*Certificate

	GetCertificate func(*ClientHelloInfo) (*Certificate, error)

	GetClientCertificate func(*CertificateRequestInfo) (*Certificate, error)

	GetConfigForClient func(*ClientHelloInfo) (*Config, error)

	VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

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

	DynamicRecordSizingDisabled bool

	Renegotiation RenegotiationSupport

	KeyLogWriter io.Writer

	serverInitOnce sync.Once

	mutex sync.RWMutex

	sessionTicketKeys []ticketKey
}

// ticketKeyNameLen is the number of bytes of identifier that is prepended to
// an encrypted session ticket in order to identify the key used to encrypt it.

// ticketKey is the internal representation of a session ticket key.

// Clone returns a shallow clone of c. It is safe to clone a Config that is
// being used concurrently by a TLS client or server.
func (c *Config) Clone() *Config

// SetSessionTicketKeys updates the session ticket keys for a server. The first
// key will be used when creating new tickets, while all keys can be used for
// decrypting tickets. It is safe to call this function while the server is
// running in order to rotate the session ticket keys. The function will panic
// if keys is empty.
func (c *Config) SetSessionTicketKeys(keys [][32]byte)

// BuildNameToCertificate parses c.Certificates and builds c.NameToCertificate
// from the CommonName and SubjectAlternateName fields of each of the leaf
// certificates.
func (c *Config) BuildNameToCertificate()

// writerMutex protects all KeyLogWriters globally. It is rarely enabled,
// and is only for debugging, so a global mutex saves space.

// A Certificate is a chain of one or more certificates, leaf first.
type Certificate struct {
	Certificate [][]byte

	PrivateKey crypto.PrivateKey

	OCSPStaple []byte

	SignedCertificateTimestamps [][]byte

	Leaf *x509.Certificate
}

// lruSessionCache is a ClientSessionCache implementation that uses an LRU
// caching strategy.

// NewLRUClientSessionCache returns a ClientSessionCache with the given
// capacity that uses an LRU strategy. If capacity is < 1, a default capacity
// is used instead.
func NewLRUClientSessionCache(capacity int) ClientSessionCache

// TODO(jsing): Make these available to both crypto/x509 and crypto/tls.
