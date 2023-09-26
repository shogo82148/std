// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

const (
	VersionTLS10 = 0x0301
	VersionTLS11 = 0x0302
	VersionTLS12 = 0x0303
	VersionTLS13 = 0x0304

	// Deprecated: SSLv3 is cryptographically broken, and is no longer
	// supported by this package. See golang.org/issue/32716.
	VersionSSL30 = 0x0300
)

// VersionName returns the name for the provided TLS version number
// (e.g. "TLS 1.3"), or a fallback representation of the value if the
// version is not implemented by this package.
func VersionName(version uint16) string

// TLS record types.

// TLS handshake message types.

// TLS compression types.

// TLS extension numbers

// TLS signaling cipher suite values

// CurveID is the type of a TLS identifier for an elliptic curve. See
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8.
//
// In TLS 1.3, this type is called NamedGroup, but at this time this library
// only supports Elliptic Curve based groups. See RFC 8446, Section 4.2.7.
type CurveID uint16

const (
	CurveP256 CurveID = 23
	CurveP384 CurveID = 24
	CurveP521 CurveID = 25
	X25519    CurveID = 29
)

// TLS 1.3 Key Share. See RFC 8446, Section 4.2.8.

// TLS 1.3 PSK Key Exchange Modes. See RFC 8446, Section 4.2.9.

// TLS 1.3 PSK Identity. Can be a Session Ticket, or a reference to a saved
// session. See RFC 8446, Section 4.2.11.

// TLS Elliptic Curve Point Formats
// https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-9

// TLS CertificateStatusType (RFC 3546)

// Certificate types (for certificateRequestMsg)

// Signature algorithms (for internal signaling use). Starting at 225 to avoid overlap with
// TLS 1.2 codepoints (RFC 5246, Appendix A.4.1), with which these have nothing to do.

// directSigning is a standard Hash value that signals that no pre-hashing
// should be performed, and that the input should be signed directly. It is the
// hash function associated with the Ed25519 signature scheme.

// defaultSupportedSignatureAlgorithms contains the signature and hash algorithms that
// the code advertises as supported in a TLS 1.2+ ClientHello and in a TLS 1.2+
// CertificateRequest. The two fields are merged to match with TLS 1.3.
// Note that in TLS 1.2, the ECDSA algorithms are not constrained to P-256, etc.

// helloRetryRequestRandom is set as the Random value of a ServerHello
// to signal that the message is actually a HelloRetryRequest.

// testingOnlyForceDowngradeCanary is set in tests to force the server side to
// include downgrade canaries even if it's using its highers supported version.

// ConnectionState records basic TLS details about the connection.
type ConnectionState struct {
	Version uint16

	HandshakeComplete bool

	DidResume bool

	CipherSuite uint16

	NegotiatedProtocol string

	NegotiatedProtocolIsMutual bool

	ServerName string

	PeerCertificates []*x509.Certificate

	VerifiedChains [][]*x509.Certificate

	SignedCertificateTimestamps [][]byte

	OCSPResponse []byte

	TLSUnique []byte

	ekm func(label string, context []byte, length int) ([]byte, error)
}

// ExportKeyingMaterial returns length bytes of exported key material in a new
// slice as defined in RFC 5705. If context is nil, it is not used as part of
// the seed. If the connection was set to allow renegotiation via
// Config.Renegotiation, this function will return an error.
//
// There are conditions in which the returned values might not be unique to a
// connection. See the Security Considerations sections of RFC 5705 and RFC 7627,
// and https://mitls.org/pages/attacks/3SHAKE#channelbindings.
func (cs *ConnectionState) ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error)

// ClientAuthType declares the policy the server will follow for
// TLS Client Authentication.
type ClientAuthType int

const (
	// NoClientCert indicates that no client certificate should be requested
	// during the handshake, and if any certificates are sent they will not
	// be verified.
	NoClientCert ClientAuthType = iota
	// RequestClientCert indicates that a client certificate should be requested
	// during the handshake, but does not require that the client send any
	// certificates.
	RequestClientCert
	// RequireAnyClientCert indicates that a client certificate should be requested
	// during the handshake, and that at least one certificate is required to be
	// sent by the client, but that certificate is not required to be valid.
	RequireAnyClientCert
	// VerifyClientCertIfGiven indicates that a client certificate should be requested
	// during the handshake, but does not require that the client sends a
	// certificate. If the client does send a certificate it is required to be
	// valid.
	VerifyClientCertIfGiven
	// RequireAndVerifyClientCert indicates that a client certificate should be requested
	// during the handshake, and that at least one valid certificate is required
	// to be sent by the client.
	RequireAndVerifyClientCert
)

// ClientSessionCache is a cache of ClientSessionState objects that can be used
// by a client to resume a TLS session with a given server. ClientSessionCache
// implementations should expect to be called concurrently from different
// goroutines. Up to TLS 1.2, only ticket-based resumption is supported, not
// SessionID-based resumption. In TLS 1.3 they were merged into PSK modes, which
// are supported via this interface.
type ClientSessionCache interface {
	Get(sessionKey string) (session *ClientSessionState, ok bool)

	Put(sessionKey string, cs *ClientSessionState)
}

// SignatureScheme identifies a signature algorithm supported by TLS. See
// RFC 8446, Section 4.2.3.
type SignatureScheme uint16

const (
	// RSASSA-PKCS1-v1_5 algorithms.
	PKCS1WithSHA256 SignatureScheme = 0x0401
	PKCS1WithSHA384 SignatureScheme = 0x0501
	PKCS1WithSHA512 SignatureScheme = 0x0601

	// RSASSA-PSS algorithms with public key OID rsaEncryption.
	PSSWithSHA256 SignatureScheme = 0x0804
	PSSWithSHA384 SignatureScheme = 0x0805
	PSSWithSHA512 SignatureScheme = 0x0806

	// ECDSA algorithms. Only constrained to a specific curve in TLS 1.3.
	ECDSAWithP256AndSHA256 SignatureScheme = 0x0403
	ECDSAWithP384AndSHA384 SignatureScheme = 0x0503
	ECDSAWithP521AndSHA512 SignatureScheme = 0x0603

	// EdDSA algorithms.
	Ed25519 SignatureScheme = 0x0807

	// Legacy signature and hash algorithms for TLS 1.2.
	PKCS1WithSHA1 SignatureScheme = 0x0201
	ECDSAWithSHA1 SignatureScheme = 0x0203
)

// ClientHelloInfo contains information from a ClientHello message in order to
// guide application logic in the GetCertificate and GetConfigForClient callbacks.
type ClientHelloInfo struct {
	CipherSuites []uint16

	ServerName string

	SupportedCurves []CurveID

	SupportedPoints []uint8

	SignatureSchemes []SignatureScheme

	SupportedProtos []string

	SupportedVersions []uint16

	Conn net.Conn

	config *Config

	ctx context.Context
}

// Context returns the context of the handshake that is in progress.
// This context is a child of the context passed to HandshakeContext,
// if any, and is canceled when the handshake concludes.
func (c *ClientHelloInfo) Context() context.Context

// CertificateRequestInfo contains information from a server's
// CertificateRequest message, which is used to demand a certificate and proof
// of control from a client.
type CertificateRequestInfo struct {
	AcceptableCAs [][]byte

	SignatureSchemes []SignatureScheme

	Version uint16

	ctx context.Context
}

// Context returns the context of the handshake that is in progress.
// This context is a child of the context passed to HandshakeContext,
// if any, and is canceled when the handshake concludes.
func (c *CertificateRequestInfo) Context() context.Context

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
//
// Renegotiation is not defined in TLS 1.3.
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

	VerifyConnection func(ConnectionState) error

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

	UnwrapSession func(identity []byte, cs ConnectionState) (*SessionState, error)

	WrapSession func(ConnectionState, *SessionState) ([]byte, error)

	MinVersion uint16

	MaxVersion uint16

	CurvePreferences []CurveID

	DynamicRecordSizingDisabled bool

	Renegotiation RenegotiationSupport

	KeyLogWriter io.Writer

	mutex sync.RWMutex

	sessionTicketKeys []ticketKey

	autoSessionTicketKeys []ticketKey
}

// ticketKey is the internal representation of a session ticket key.

// maxSessionTicketLifetime is the maximum allowed lifetime of a TLS 1.3 session
// ticket, and the lifetime we set for all tickets we send.

// Clone returns a shallow clone of c or nil if c is nil. It is safe to clone a Config that is
// being used concurrently by a TLS client or server.
func (c *Config) Clone() *Config

// deprecatedSessionTicketKey is set as the prefix of SessionTicketKey if it was
// randomized for backwards compatibility but is not in use.

// SetSessionTicketKeys updates the session ticket keys for a server.
//
// The first key will be used when creating new tickets, while all keys can be
// used for decrypting tickets. It is safe to call this function while the
// server is running in order to rotate the session ticket keys. The function
// will panic if keys is empty.
//
// Calling this function will turn off automatic session ticket key rotation.
//
// If multiple servers are terminating connections for the same host they should
// all have the same session ticket keys. If the session ticket keys leaks,
// previously recorded and future TLS connections using those keys might be
// compromised.
func (c *Config) SetSessionTicketKeys(keys [][32]byte)

// roleClient and roleServer are meant to call supportedVersions and parents
// with more readability at the callsite.

// SupportsCertificate returns nil if the provided certificate is supported by
// the client that sent the ClientHello. Otherwise, it returns an error
// describing the reason for the incompatibility.
//
// If this ClientHelloInfo was passed to a GetConfigForClient or GetCertificate
// callback, this method will take into account the associated Config. Note that
// if GetConfigForClient returns a different Config, the change can't be
// accounted for by this method.
//
// This function will call x509.ParseCertificate unless c.Leaf is set, which can
// incur a significant performance cost.
func (chi *ClientHelloInfo) SupportsCertificate(c *Certificate) error

// SupportsCertificate returns nil if the provided certificate is supported by
// the server that sent the CertificateRequest. Otherwise, it returns an error
// describing the reason for the incompatibility.
func (cri *CertificateRequestInfo) SupportsCertificate(c *Certificate) error

// BuildNameToCertificate parses c.Certificates and builds c.NameToCertificate
// from the CommonName and SubjectAlternateName fields of each of the leaf
// certificates.
//
// Deprecated: NameToCertificate only allows associating a single certificate
// with a given name. Leave that field nil to let the library select the first
// compatible chain from Certificates.
func (c *Config) BuildNameToCertificate()

// writerMutex protects all KeyLogWriters globally. It is rarely enabled,
// and is only for debugging, so a global mutex saves space.

// A Certificate is a chain of one or more certificates, leaf first.
type Certificate struct {
	Certificate [][]byte

	PrivateKey crypto.PrivateKey

	SupportedSignatureAlgorithms []SignatureScheme

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

// CertificateVerificationError is returned when certificate verification fails during the handshake.
type CertificateVerificationError struct {
	UnverifiedCertificates []*x509.Certificate
	Err                    error
}

func (e *CertificateVerificationError) Error() string

func (e *CertificateVerificationError) Unwrap() error
