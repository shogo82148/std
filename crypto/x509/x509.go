// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package x509 parses X.509-encoded keys and certificates.
package x509

import (
	"github.com/shogo82148/std/crypto/x509/pkix"
	"github.com/shogo82148/std/encoding/asn1"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
	"github.com/shogo82148/std/time"
)

// pkixPublicKey reflects a PKIX public key structure. See SubjectPublicKeyInfo
// in RFC 3280.

// ParsePKIXPublicKey parses a DER encoded public key. These values are
// typically found in PEM blocks with "BEGIN PUBLIC KEY".
func ParsePKIXPublicKey(derBytes []byte) (pub interface{}, err error)

// MarshalPKIXPublicKey serialises a public key to DER-encoded PKIX format.
func MarshalPKIXPublicKey(pub interface{}) ([]byte, error)

// RFC 5280,  4.2.1.1

type SignatureAlgorithm int

const (
	UnknownSignatureAlgorithm SignatureAlgorithm = iota
	MD2WithRSA
	MD5WithRSA
	SHA1WithRSA
	SHA256WithRSA
	SHA384WithRSA
	SHA512WithRSA
	DSAWithSHA1
	DSAWithSHA256
)

type PublicKeyAlgorithm int

const (
	UnknownPublicKeyAlgorithm PublicKeyAlgorithm = iota
	RSA
	DSA
)

// OIDs for signature algorithms
//
// pkcs-1 OBJECT IDENTIFIER ::= {
//    iso(1) member-body(2) us(840) rsadsi(113549) pkcs(1) 1 }
//
//
// RFC 3279 2.2.1 RSA Signature Algorithms
//
// md2WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 2 }
//
// md5WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 4 }
//
// sha-1WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 5 }
//
// dsaWithSha1 OBJECT IDENTIFIER ::= {
//    iso(1) member-body(2) us(840) x9-57(10040) x9cm(4) 3 }
//
//
// RFC 4055 5 PKCS #1 Version 1.5
//
// sha256WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 11 }
//
// sha384WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 12 }
//
// sha512WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 13 }
//
//
// RFC 5758 3.1 DSA Signature Algorithms
//
// dsaWithSha256 OBJECT IDENTIFIER ::= {
//    joint-iso-ccitt(2) country(16) us(840) organization(1) gov(101)
//    csor(3) algorithms(4) id-dsa-with-sha2(3) 2}
//

// RFC 3279, 2.3 Public Key Algorithms
//
// pkcs-1 OBJECT IDENTIFIER ::== { iso(1) member-body(2) us(840)
//    rsadsi(113549) pkcs(1) 1 }
//
// rsaEncryption OBJECT IDENTIFIER ::== { pkcs1-1 1 }
//
// id-dsa OBJECT IDENTIFIER ::== { iso(1) member-body(2) us(840)
//    x9-57(10040) x9cm(4) 1 }

// KeyUsage represents the set of actions that are valid for a given key. It's
// a bitmap of the KeyUsage* constants.
type KeyUsage int

const (
	KeyUsageDigitalSignature KeyUsage = 1 << iota
	KeyUsageContentCommitment
	KeyUsageKeyEncipherment
	KeyUsageDataEncipherment
	KeyUsageKeyAgreement
	KeyUsageCertSign
	KeyUsageCRLSign
	KeyUsageEncipherOnly
	KeyUsageDecipherOnly
)

// RFC 5280, 4.2.1.12  Extended Key Usage
//
// anyExtendedKeyUsage OBJECT IDENTIFIER ::= { id-ce-extKeyUsage 0 }
//
// id-kp OBJECT IDENTIFIER ::= { id-pkix 3 }
//
// id-kp-serverAuth             OBJECT IDENTIFIER ::= { id-kp 1 }
// id-kp-clientAuth             OBJECT IDENTIFIER ::= { id-kp 2 }
// id-kp-codeSigning            OBJECT IDENTIFIER ::= { id-kp 3 }
// id-kp-emailProtection        OBJECT IDENTIFIER ::= { id-kp 4 }
// id-kp-timeStamping           OBJECT IDENTIFIER ::= { id-kp 8 }
// id-kp-OCSPSigning            OBJECT IDENTIFIER ::= { id-kp 9 }

// ExtKeyUsage represents an extended set of actions that are valid for a given key.
// Each of the ExtKeyUsage* constants define a unique action.
type ExtKeyUsage int

const (
	ExtKeyUsageAny ExtKeyUsage = iota
	ExtKeyUsageServerAuth
	ExtKeyUsageClientAuth
	ExtKeyUsageCodeSigning
	ExtKeyUsageEmailProtection
	ExtKeyUsageTimeStamping
	ExtKeyUsageOCSPSigning
)

// A Certificate represents an X.509 certificate.
type Certificate struct {
	Raw                     []byte
	RawTBSCertificate       []byte
	RawSubjectPublicKeyInfo []byte
	RawSubject              []byte
	RawIssuer               []byte

	Signature          []byte
	SignatureAlgorithm SignatureAlgorithm

	PublicKeyAlgorithm PublicKeyAlgorithm
	PublicKey          interface{}

	Version             int
	SerialNumber        *big.Int
	Issuer              pkix.Name
	Subject             pkix.Name
	NotBefore, NotAfter time.Time
	KeyUsage            KeyUsage

	ExtKeyUsage        []ExtKeyUsage
	UnknownExtKeyUsage []asn1.ObjectIdentifier

	BasicConstraintsValid bool
	IsCA                  bool
	MaxPathLen            int

	SubjectKeyId   []byte
	AuthorityKeyId []byte

	DNSNames       []string
	EmailAddresses []string

	PermittedDNSDomainsCritical bool
	PermittedDNSDomains         []string

	PolicyIdentifiers []asn1.ObjectIdentifier
}

// ErrUnsupportedAlgorithm results from attempting to perform an operation that
// involves algorithms that are not currently implemented.
var ErrUnsupportedAlgorithm = errors.New("crypto/x509: cannot verify signature: algorithm unimplemented")

// ConstraintViolationError results when a requested usage is not permitted by
// a certificate. For example: checking a signature when the public key isn't a
// certificate signing key.
type ConstraintViolationError struct{}

func (ConstraintViolationError) Error() string

func (c *Certificate) Equal(other *Certificate) bool

// Entrust have a broken root certificate (CN=Entrust.net Certification
// Authority (2048)) which isn't marked as a CA certificate and is thus invalid
// according to PKIX.
// We recognise this certificate by its SubjectPublicKeyInfo and exempt it
// from the Basic Constraints requirement.
// See http://www.entrust.net/knowledge-base/technote.cfm?tn=7869
//
// TODO(agl): remove this hack once their reissued root is sufficiently
// widespread.

// CheckSignatureFrom verifies that the signature on c is a valid signature
// from parent.
func (c *Certificate) CheckSignatureFrom(parent *Certificate) (err error)

// CheckSignature verifies that signature is a valid signature over signed from
// c's public key.
func (c *Certificate) CheckSignature(algo SignatureAlgorithm, signed, signature []byte) (err error)

// CheckCRLSignature checks that the signature in crl is from c.
func (c *Certificate) CheckCRLSignature(crl *pkix.CertificateList) (err error)

type UnhandledCriticalExtension struct{}

func (h UnhandledCriticalExtension) Error() string

// RFC 5280 4.2.1.4

// RFC 5280, 4.2.1.10

// ParseCertificate parses a single certificate from the given ASN.1 DER data.
func ParseCertificate(asn1Data []byte) (*Certificate, error)

// ParseCertificates parses one or more certificates from the given ASN.1 DER
// data. The certificates must be concatenated with no intermediate padding.
func ParseCertificates(asn1Data []byte) ([]*Certificate, error)

// CreateCertificate creates a new certificate based on a template. The
// following members of template are used: SerialNumber, Subject, NotBefore,
// NotAfter, KeyUsage, BasicConstraintsValid, IsCA, MaxPathLen, SubjectKeyId,
// DNSNames, PermittedDNSDomainsCritical, PermittedDNSDomains.
//
// The certificate is signed by parent. If parent is equal to template then the
// certificate is self-signed. The parameter pub is the public key of the
// signee and priv is the private key of the signer.
//
// The returned slice is the certificate in DER encoding.
//
// The only supported key type is RSA (*rsa.PublicKey for pub, *rsa.PrivateKey
// for priv).
func CreateCertificate(rand io.Reader, template, parent *Certificate, pub interface{}, priv interface{}) (cert []byte, err error)

// pemCRLPrefix is the magic string that indicates that we have a PEM encoded
// CRL.

// pemType is the type of a PEM encoded CRL.

// ParseCRL parses a CRL from the given bytes. It's often the case that PEM
// encoded CRLs will appear where they should be DER encoded, so this function
// will transparently handle PEM encoding as long as there isn't any leading
// garbage.
func ParseCRL(crlBytes []byte) (certList *pkix.CertificateList, err error)

// ParseDERCRL parses a DER encoded CRL from the given bytes.
func ParseDERCRL(derBytes []byte) (certList *pkix.CertificateList, err error)

// CreateCRL returns a DER encoded CRL, signed by this Certificate, that
// contains the given list of revoked certificates.
//
// The only supported key type is RSA (*rsa.PrivateKey for priv).
func (c *Certificate) CreateCRL(rand io.Reader, priv interface{}, revokedCerts []pkix.RevokedCertificate, now, expiry time.Time) (crlBytes []byte, err error)
