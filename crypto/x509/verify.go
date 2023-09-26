// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/time"
)

type InvalidReason int

const (
	// NotAuthorizedToSign results when a certificate is signed by another
	// which isn't marked as a CA certificate.
	NotAuthorizedToSign InvalidReason = iota
	// Expired results when a certificate has expired, based on the time
	// given in the VerifyOptions.
	Expired
	// CANotAuthorizedForThisName results when an intermediate or root
	// certificate has a name constraint which doesn't permit a DNS or
	// other name (including IP address) in the leaf certificate.
	CANotAuthorizedForThisName
	// TooManyIntermediates results when a path length constraint is
	// violated.
	TooManyIntermediates
	// IncompatibleUsage results when the certificate's key usage indicates
	// that it may only be used for a different purpose.
	IncompatibleUsage
	// NameMismatch results when the subject name of a parent certificate
	// does not match the issuer name in the child.
	NameMismatch
	// NameConstraintsWithoutSANs results when a leaf certificate doesn't
	// contain a Subject Alternative Name extension, but a CA certificate
	// contains name constraints.
	NameConstraintsWithoutSANs
	// UnconstrainedName results when a CA certificate contains permitted
	// name constraints, but leaf certificate contains a name of an
	// unsupported or unconstrained type.
	UnconstrainedName
	// TooManyConstraints results when the number of comparision operations
	// needed to check a certificate exceeds the limit set by
	// VerifyOptions.MaxConstraintComparisions. This limit exists to
	// prevent pathological certificates can consuming excessive amounts of
	// CPU time to verify.
	TooManyConstraints
	// CANotAuthorizedForExtKeyUsage results when an intermediate or root
	// certificate does not permit a requested extended key usage.
	CANotAuthorizedForExtKeyUsage
)

// CertificateInvalidError results when an odd error occurs. Users of this
// library probably want to handle all these errors uniformly.
type CertificateInvalidError struct {
	Cert   *Certificate
	Reason InvalidReason
	Detail string
}

func (e CertificateInvalidError) Error() string

// HostnameError results when the set of authorized names doesn't match the
// requested name.
type HostnameError struct {
	Certificate *Certificate
	Host        string
}

func (h HostnameError) Error() string

// UnknownAuthorityError results when the certificate issuer is unknown
type UnknownAuthorityError struct {
	Cert *Certificate

	hintErr error

	hintCert *Certificate
}

func (e UnknownAuthorityError) Error() string

// SystemRootsError results when we fail to load the system root certificates.
type SystemRootsError struct {
	Err error
}

func (se SystemRootsError) Error() string

// errNotParsed is returned when a certificate without ASN.1 contents is
// verified. Platform-specific verification needs the ASN.1 contents.

// VerifyOptions contains parameters for Certificate.Verify. It's a structure
// because other PKIX verification APIs have ended up needing many options.
type VerifyOptions struct {
	DNSName       string
	Intermediates *CertPool
	Roots         *CertPool
	CurrentTime   time.Time

	KeyUsages []ExtKeyUsage

	MaxConstraintComparisions int
}

// rfc2821Mailbox represents a “mailbox” (which is an email address to most
// people) by breaking it into the “local” (i.e. before the '@') and “domain”
// parts.

// Verify attempts to verify c by building one or more chains from c to a
// certificate in opts.Roots, using certificates in opts.Intermediates if
// needed. If successful, it returns one or more chains where the first
// element of the chain is c and the last element is from opts.Roots.
//
// If opts.Roots is nil and system roots are unavailable the returned error
// will be of type SystemRootsError.
//
// Name constraints in the intermediates will be applied to all names claimed
// in the chain, not just opts.DNSName. Thus it is invalid for a leaf to claim
// example.com if an intermediate doesn't permit it, even if example.com is not
// the name being validated. Note that DirectoryName constraints are not
// supported.
//
// Extended Key Usage values are enforced down a chain, so an intermediate or
// root that enumerates EKUs prevents a leaf from asserting an EKU not in that
// list.
//
// WARNING: this function doesn't do any revocation checking.
func (c *Certificate) Verify(opts VerifyOptions) (chains [][]*Certificate, err error)

// maxChainSignatureChecks is the maximum number of CheckSignatureFrom calls
// that an invocation of buildChains will (tranistively) make. Most chains are
// less than 15 certificates long, so this leaves space for multiple chains and
// for failed checks due to different intermediates having the same Subject.

// VerifyHostname returns nil if c is a valid certificate for the named host.
// Otherwise it returns an error describing the mismatch.
func (c *Certificate) VerifyHostname(h string) error
