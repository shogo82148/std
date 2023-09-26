// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package x509 implements a subset of the X.509 standard.
//
// It allows parsing and generating certificates, certificate signing
// requests, certificate revocation lists, and encoded public and private keys.
// It provides a certificate verifier, complete with a chain builder.
//
// The package targets the X.509 technical profile defined by the IETF (RFC
// 2459/3280/5280), and as further restricted by the CA/Browser Forum Baseline
// Requirements. There is minimal support for features outside of these
// profiles, as the primary goal of the package is to provide compatibility
// with the publicly trusted TLS certificate ecosystem and its policies and
// constraints.
//
// On macOS and Windows, certificate verification is handled by system APIs, but
// the package aims to apply consistent validation rules across operating
// systems.
package x509

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/crypto/x509/pkix"
	"github.com/shogo82148/std/encoding/asn1"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math/big"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/url"
	"github.com/shogo82148/std/time"

	_ "github.com/shogo82148/std/crypto/sha1"
	_ "github.com/shogo82148/std/crypto/sha256"
	_ "github.com/shogo82148/std/crypto/sha512"

	cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

// pkixPublicKey reflects a PKIX public key structure. See SubjectPublicKeyInfo
// in RFC 3280.

// ParsePKIXPublicKey parses a public key in PKIX, ASN.1 DER form. The encoded
// public key is a SubjectPublicKeyInfo structure (see RFC 5280, Section 4.1).
//
// It returns a *rsa.PublicKey, *dsa.PublicKey, *ecdsa.PublicKey,
// ed25519.PublicKey (not a pointer), or *ecdh.PublicKey (for X25519).
// More types might be supported in the future.
//
// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".
func ParsePKIXPublicKey(derBytes []byte) (pub any, err error)

// MarshalPKIXPublicKey converts a public key to PKIX, ASN.1 DER form.
// The encoded public key is a SubjectPublicKeyInfo structure
// (see RFC 5280, Section 4.1).
//
// The following key types are currently supported: *rsa.PublicKey,
// *ecdsa.PublicKey, ed25519.PublicKey (not a pointer), and *ecdh.PublicKey.
// Unsupported key types result in an error.
//
// This kind of key is commonly encoded in PEM blocks of type "PUBLIC KEY".
func MarshalPKIXPublicKey(pub any) ([]byte, error)

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
	ECDSAWithSHA1
	ECDSAWithSHA256
	ECDSAWithSHA384
	ECDSAWithSHA512
	SHA256WithRSAPSS
	SHA384WithRSAPSS
	SHA512WithRSAPSS
	PureEd25519
)

func (algo SignatureAlgorithm) String() string

type PublicKeyAlgorithm int

const (
	UnknownPublicKeyAlgorithm PublicKeyAlgorithm = iota
	RSA
	DSA
	ECDSA
	Ed25519
)

func (algo PublicKeyAlgorithm) String() string

// OIDs for signature algorithms
//
//	pkcs-1 OBJECT IDENTIFIER ::= {
//		iso(1) member-body(2) us(840) rsadsi(113549) pkcs(1) 1 }
//
// RFC 3279 2.2.1 RSA Signature Algorithms
//
//	md2WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 2 }
//
//	md5WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 4 }
//
//	sha-1WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 5 }
//
//	dsaWithSha1 OBJECT IDENTIFIER ::= {
//		iso(1) member-body(2) us(840) x9-57(10040) x9cm(4) 3 }
//
// RFC 3279 2.2.3 ECDSA Signature Algorithm
//
//	ecdsa-with-SHA1 OBJECT IDENTIFIER ::= {
//		iso(1) member-body(2) us(840) ansi-x962(10045)
//		signatures(4) ecdsa-with-SHA1(1)}
//
// RFC 4055 5 PKCS #1 Version 1.5
//
//	sha256WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 11 }
//
//	sha384WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 12 }
//
//	sha512WithRSAEncryption OBJECT IDENTIFIER ::= { pkcs-1 13 }
//
// RFC 5758 3.1 DSA Signature Algorithms
//
//	dsaWithSha256 OBJECT IDENTIFIER ::= {
//		joint-iso-ccitt(2) country(16) us(840) organization(1) gov(101)
//		csor(3) algorithms(4) id-dsa-with-sha2(3) 2}
//
// RFC 5758 3.2 ECDSA Signature Algorithm
//
//	ecdsa-with-SHA256 OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//		us(840) ansi-X9-62(10045) signatures(4) ecdsa-with-SHA2(3) 2 }
//
//	ecdsa-with-SHA384 OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//		us(840) ansi-X9-62(10045) signatures(4) ecdsa-with-SHA2(3) 3 }
//
//	ecdsa-with-SHA512 OBJECT IDENTIFIER ::= { iso(1) member-body(2)
//		us(840) ansi-X9-62(10045) signatures(4) ecdsa-with-SHA2(3) 4 }
//
// RFC 8410 3 Curve25519 and Curve448 Algorithm Identifiers
//
//	id-Ed25519   OBJECT IDENTIFIER ::= { 1 3 101 112 }

// hashToPSSParameters contains the DER encoded RSA PSS parameters for the
// SHA256, SHA384, and SHA512 hashes as defined in RFC 3447, Appendix A.2.3.
// The parameters contain the following values:
//   - hashAlgorithm contains the associated hash identifier with NULL parameters
//   - maskGenAlgorithm always contains the default mgf1SHA1 identifier
//   - saltLength contains the length of the associated hash
//   - trailerField always contains the default trailerFieldBC value

// pssParameters reflects the parameters in an AlgorithmIdentifier that
// specifies RSA PSS. See RFC 3447, Appendix A.2.3.

// RFC 5480, 2.1.1.1. Named Curve
//
//	secp224r1 OBJECT IDENTIFIER ::= {
//	  iso(1) identified-organization(3) certicom(132) curve(0) 33 }
//
//	secp256r1 OBJECT IDENTIFIER ::= {
//	  iso(1) member-body(2) us(840) ansi-X9-62(10045) curves(3)
//	  prime(1) 7 }
//
//	secp384r1 OBJECT IDENTIFIER ::= {
//	  iso(1) identified-organization(3) certicom(132) curve(0) 34 }
//
//	secp521r1 OBJECT IDENTIFIER ::= {
//	  iso(1) identified-organization(3) certicom(132) curve(0) 35 }
//
// NB: secp256r1 is equivalent to prime256v1

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
//	anyExtendedKeyUsage OBJECT IDENTIFIER ::= { id-ce-extKeyUsage 0 }
//
//	id-kp OBJECT IDENTIFIER ::= { id-pkix 3 }
//
//	id-kp-serverAuth             OBJECT IDENTIFIER ::= { id-kp 1 }
//	id-kp-clientAuth             OBJECT IDENTIFIER ::= { id-kp 2 }
//	id-kp-codeSigning            OBJECT IDENTIFIER ::= { id-kp 3 }
//	id-kp-emailProtection        OBJECT IDENTIFIER ::= { id-kp 4 }
//	id-kp-timeStamping           OBJECT IDENTIFIER ::= { id-kp 8 }
//	id-kp-OCSPSigning            OBJECT IDENTIFIER ::= { id-kp 9 }

// ExtKeyUsage represents an extended set of actions that are valid for a given key.
// Each of the ExtKeyUsage* constants define a unique action.
type ExtKeyUsage int

const (
	ExtKeyUsageAny ExtKeyUsage = iota
	ExtKeyUsageServerAuth
	ExtKeyUsageClientAuth
	ExtKeyUsageCodeSigning
	ExtKeyUsageEmailProtection
	ExtKeyUsageIPSECEndSystem
	ExtKeyUsageIPSECTunnel
	ExtKeyUsageIPSECUser
	ExtKeyUsageTimeStamping
	ExtKeyUsageOCSPSigning
	ExtKeyUsageMicrosoftServerGatedCrypto
	ExtKeyUsageNetscapeServerGatedCrypto
	ExtKeyUsageMicrosoftCommercialCodeSigning
	ExtKeyUsageMicrosoftKernelCodeSigning
)

// extKeyUsageOIDs contains the mapping between an ExtKeyUsage and its OID.

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
	PublicKey          any

	Version             int
	SerialNumber        *big.Int
	Issuer              pkix.Name
	Subject             pkix.Name
	NotBefore, NotAfter time.Time
	KeyUsage            KeyUsage

	Extensions []pkix.Extension

	ExtraExtensions []pkix.Extension

	UnhandledCriticalExtensions []asn1.ObjectIdentifier

	ExtKeyUsage        []ExtKeyUsage
	UnknownExtKeyUsage []asn1.ObjectIdentifier

	BasicConstraintsValid bool
	IsCA                  bool

	MaxPathLen int

	MaxPathLenZero bool

	SubjectKeyId   []byte
	AuthorityKeyId []byte

	OCSPServer            []string
	IssuingCertificateURL []string

	DNSNames       []string
	EmailAddresses []string
	IPAddresses    []net.IP
	URIs           []*url.URL

	PermittedDNSDomainsCritical bool
	PermittedDNSDomains         []string
	ExcludedDNSDomains          []string
	PermittedIPRanges           []*net.IPNet
	ExcludedIPRanges            []*net.IPNet
	PermittedEmailAddresses     []string
	ExcludedEmailAddresses      []string
	PermittedURIDomains         []string
	ExcludedURIDomains          []string

	CRLDistributionPoints []string

	PolicyIdentifiers []asn1.ObjectIdentifier
}

// ErrUnsupportedAlgorithm results from attempting to perform an operation that
// involves algorithms that are not currently implemented.
var ErrUnsupportedAlgorithm = errors.New("x509: cannot verify signature: algorithm unimplemented")

// An InsecureAlgorithmError indicates that the SignatureAlgorithm used to
// generate the signature is not secure, and the signature has been rejected.
//
// To temporarily restore support for SHA-1 signatures, include the value
// "x509sha1=1" in the GODEBUG environment variable. Note that this option will
// be removed in a future release.
type InsecureAlgorithmError SignatureAlgorithm

func (e InsecureAlgorithmError) Error() string

// ConstraintViolationError results when a requested usage is not permitted by
// a certificate. For example: checking a signature when the public key isn't a
// certificate signing key.
type ConstraintViolationError struct{}

func (ConstraintViolationError) Error() string

func (c *Certificate) Equal(other *Certificate) bool

// CheckSignatureFrom verifies that the signature on c is a valid signature from parent.
//
// This is a low-level API that performs very limited checks, and not a full
// path verifier. Most users should use [Certificate.Verify] instead.
func (c *Certificate) CheckSignatureFrom(parent *Certificate) error

// CheckSignature verifies that signature is a valid signature over signed from
// c's public key.
//
// This is a low-level API that performs no validity checks on the certificate.
//
// [MD5WithRSA] signatures are rejected, while [SHA1WithRSA] and [ECDSAWithSHA1]
// signatures are currently accepted.
func (c *Certificate) CheckSignature(algo SignatureAlgorithm, signed, signature []byte) error

// CheckCRLSignature checks that the signature in crl is from c.
//
// Deprecated: Use RevocationList.CheckSignatureFrom instead.
func (c *Certificate) CheckCRLSignature(crl *pkix.CertificateList) error

type UnhandledCriticalExtension struct{}

func (h UnhandledCriticalExtension) Error() string

// RFC 5280 4.2.1.4

// RFC 5280, 4.2.2.1

// RFC 5280, 4.2.1.14

// emptyASN1Subject is the ASN.1 DER encoding of an empty Subject, which is
// just an empty SEQUENCE.

// CreateCertificate creates a new X.509 v3 certificate based on a template.
// The following members of template are currently used:
//
//   - AuthorityKeyId
//   - BasicConstraintsValid
//   - CRLDistributionPoints
//   - DNSNames
//   - EmailAddresses
//   - ExcludedDNSDomains
//   - ExcludedEmailAddresses
//   - ExcludedIPRanges
//   - ExcludedURIDomains
//   - ExtKeyUsage
//   - ExtraExtensions
//   - IPAddresses
//   - IsCA
//   - IssuingCertificateURL
//   - KeyUsage
//   - MaxPathLen
//   - MaxPathLenZero
//   - NotAfter
//   - NotBefore
//   - OCSPServer
//   - PermittedDNSDomains
//   - PermittedDNSDomainsCritical
//   - PermittedEmailAddresses
//   - PermittedIPRanges
//   - PermittedURIDomains
//   - PolicyIdentifiers
//   - SerialNumber
//   - SignatureAlgorithm
//   - Subject
//   - SubjectKeyId
//   - URIs
//   - UnknownExtKeyUsage
//
// The certificate is signed by parent. If parent is equal to template then the
// certificate is self-signed. The parameter pub is the public key of the
// certificate to be generated and priv is the private key of the signer.
//
// The returned slice is the certificate in DER encoding.
//
// The currently supported key types are *rsa.PublicKey, *ecdsa.PublicKey and
// ed25519.PublicKey. pub must be a supported key type, and priv must be a
// crypto.Signer with a supported public key.
//
// The AuthorityKeyId will be taken from the SubjectKeyId of parent, if any,
// unless the resulting certificate is self-signed. Otherwise the value from
// template will be used.
//
// If SubjectKeyId from template is empty and the template is a CA, SubjectKeyId
// will be generated from the hash of the public key.
func CreateCertificate(rand io.Reader, template, parent *Certificate, pub, priv any) ([]byte, error)

// pemCRLPrefix is the magic string that indicates that we have a PEM encoded
// CRL.

// pemType is the type of a PEM encoded CRL.

// ParseCRL parses a CRL from the given bytes. It's often the case that PEM
// encoded CRLs will appear where they should be DER encoded, so this function
// will transparently handle PEM encoding as long as there isn't any leading
// garbage.
//
// Deprecated: Use ParseRevocationList instead.
func ParseCRL(crlBytes []byte) (*pkix.CertificateList, error)

// ParseDERCRL parses a DER encoded CRL from the given bytes.
//
// Deprecated: Use ParseRevocationList instead.
func ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error)

// CreateCRL returns a DER encoded CRL, signed by this Certificate, that
// contains the given list of revoked certificates.
//
// Deprecated: this method does not generate an RFC 5280 conformant X.509 v2 CRL.
// To generate a standards compliant CRL, use CreateRevocationList instead.
func (c *Certificate) CreateCRL(rand io.Reader, priv any, revokedCerts []pkix.RevokedCertificate, now, expiry time.Time) (crlBytes []byte, err error)

// CertificateRequest represents a PKCS #10, certificate signature request.
type CertificateRequest struct {
	Raw                      []byte
	RawTBSCertificateRequest []byte
	RawSubjectPublicKeyInfo  []byte
	RawSubject               []byte

	Version            int
	Signature          []byte
	SignatureAlgorithm SignatureAlgorithm

	PublicKeyAlgorithm PublicKeyAlgorithm
	PublicKey          any

	Subject pkix.Name

	Attributes []pkix.AttributeTypeAndValueSET

	Extensions []pkix.Extension

	ExtraExtensions []pkix.Extension

	DNSNames       []string
	EmailAddresses []string
	IPAddresses    []net.IP
	URIs           []*url.URL
}

// oidExtensionRequest is a PKCS #9 OBJECT IDENTIFIER that indicates requested
// extensions in a CSR.

// CreateCertificateRequest creates a new certificate request based on a
// template. The following members of template are used:
//
//   - SignatureAlgorithm
//   - Subject
//   - DNSNames
//   - EmailAddresses
//   - IPAddresses
//   - URIs
//   - ExtraExtensions
//   - Attributes (deprecated)
//
// priv is the private key to sign the CSR with, and the corresponding public
// key will be included in the CSR. It must implement crypto.Signer and its
// Public() method must return a *rsa.PublicKey or a *ecdsa.PublicKey or a
// ed25519.PublicKey. (A *rsa.PrivateKey, *ecdsa.PrivateKey or
// ed25519.PrivateKey satisfies this.)
//
// The returned slice is the certificate request in DER encoding.
func CreateCertificateRequest(rand io.Reader, template *CertificateRequest, priv any) (csr []byte, err error)

// ParseCertificateRequest parses a single certificate request from the
// given ASN.1 DER data.
func ParseCertificateRequest(asn1Data []byte) (*CertificateRequest, error)

// CheckSignature reports whether the signature on c is valid.
func (c *CertificateRequest) CheckSignature() error

// RevocationList contains the fields used to create an X.509 v2 Certificate
// Revocation list with CreateRevocationList.
type RevocationList struct {
	Raw []byte

	RawTBSRevocationList []byte

	RawIssuer []byte

	Issuer pkix.Name

	AuthorityKeyId []byte

	Signature []byte

	SignatureAlgorithm SignatureAlgorithm

	RevokedCertificates []pkix.RevokedCertificate

	Number *big.Int

	ThisUpdate time.Time

	NextUpdate time.Time

	Extensions []pkix.Extension

	ExtraExtensions []pkix.Extension
}

// These structures reflect the ASN.1 structure of X.509 CRLs better than
// the existing crypto/x509/pkix variants do. These mirror the existing
// certificate structs in this file.
//
// Notably, we include issuer as an asn1.RawValue, mirroring the behavior of
// tbsCertificate and allowing raw (unparsed) subjects to be passed cleanly.

// CreateRevocationList creates a new X.509 v2 Certificate Revocation List,
// according to RFC 5280, based on template.
//
// The CRL is signed by priv which should be the private key associated with
// the public key in the issuer certificate.
//
// The issuer may not be nil, and the crlSign bit must be set in KeyUsage in
// order to use it as a CRL issuer.
//
// The issuer distinguished name CRL field and authority key identifier
// extension are populated using the issuer certificate. issuer must have
// SubjectKeyId set.
func CreateRevocationList(rand io.Reader, template *RevocationList, issuer *Certificate, priv crypto.Signer) ([]byte, error)

// CheckSignatureFrom verifies that the signature on rl is a valid signature
// from issuer.
func (rl *RevocationList) CheckSignatureFrom(parent *Certificate) error
