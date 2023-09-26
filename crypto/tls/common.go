// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// TLS record types.

// TLS handshake message types.

// TLS compression types.

// TLS extension numbers

// TLS Elliptic Curves
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8

// TLS Elliptic Curve Point Formats
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-9

// TLS CertificateStatusType (RFC 3546)

// Certificate types (for certificateRequestMsg)

// ConnectionState records basic TLS details about the connection.
type ConnectionState struct {
	HandshakeComplete          bool
	CipherSuite                uint16
	NegotiatedProtocol         string
	NegotiatedProtocolIsMutual bool

	ServerName string

	PeerCertificates []*x509.Certificate

	VerifiedChains [][]*x509.Certificate
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

// A Config structure is used to configure a TLS client or server. After one
// has been passed to a TLS function it must not be modified.
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
