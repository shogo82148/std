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
	// certificate has a name constraint which doesn't include the name
	// being checked.
	CANotAuthorizedForThisName
	// TooManyIntermediates results when a path length constraint is
	// violated.
	TooManyIntermediates
)

// CertificateInvalidError results when an odd error occurs. Users of this
// library probably want to handle all these errors uniformly.
type CertificateInvalidError struct {
	Cert   *Certificate
	Reason InvalidReason
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
	cert *Certificate
}

func (e UnknownAuthorityError) Error() string

// VerifyOptions contains parameters for Certificate.Verify. It's a structure
// because other PKIX verification APIs have ended up needing many options.
type VerifyOptions struct {
	DNSName       string
	Intermediates *CertPool
	Roots         *CertPool
	CurrentTime   time.Time
}

// Verify attempts to verify c by building one or more chains from c to a
// certificate in opts.Roots, using certificates in opts.Intermediates if
// needed. If successful, it returns one or more chains where the first
// element of the chain is c and the last element is from opts.Roots.
//
// WARNING: this doesn't do any revocation checking.
func (c *Certificate) Verify(opts VerifyOptions) (chains [][]*Certificate, err error)

// VerifyHostname returns nil if c is a valid certificate for the named host.
// Otherwise it returns an error describing the mismatch.
func (c *Certificate) VerifyHostname(h string) error
