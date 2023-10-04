// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// CertPool is a set of certificates.
type CertPool struct {
	byName map[string][]int

	// lazyCerts contains funcs that return a certificate,
	// lazily parsing/decompressing it as needed.
	lazyCerts []lazyCert

	// haveSum maps from sum224(cert.Raw) to true. It's used only
	// for AddCert duplicate detection, to avoid CertPool.contains
	// calls in the AddCert path (because the contains method can
	// call getCert and otherwise negate savings from lazy getCert
	// funcs).
	haveSum map[sum224]bool

	// systemPool indicates whether this is a special pool derived from the
	// system roots. If it includes additional roots, it requires doing two
	// verifications, one using the roots provided by the caller, and one using
	// the system platform verifier.
	systemPool bool
}

// NewCertPool returns a new, empty CertPool.
func NewCertPool() *CertPool

// Clone returns a copy of s.
func (s *CertPool) Clone() *CertPool

// SystemCertPool returns a copy of the system cert pool.
//
// On Unix systems other than macOS the environment variables SSL_CERT_FILE and
// SSL_CERT_DIR can be used to override the system default locations for the SSL
// certificate file and SSL certificate files directory, respectively. The
// latter can be a colon-separated list.
//
// Any mutations to the returned pool are not written to disk and do not affect
// any other pool returned by SystemCertPool.
//
// New changes in the system cert pool might not be reflected in subsequent calls.
func SystemCertPool() (*CertPool, error)

// AddCert adds a certificate to a pool.
func (s *CertPool) AddCert(cert *Certificate)

// AppendCertsFromPEM attempts to parse a series of PEM encoded certificates.
// It appends any certificates found to s and reports whether any certificates
// were successfully parsed.
//
// On many Linux systems, /etc/ssl/cert.pem will contain the system wide set
// of root CAs in a format suitable for this function.
func (s *CertPool) AppendCertsFromPEM(pemCerts []byte) (ok bool)

// Subjects returns a list of the DER-encoded subjects of
// all of the certificates in the pool.
//
// Deprecated: if s was returned by SystemCertPool, Subjects
// will not include the system roots.
func (s *CertPool) Subjects() [][]byte

// Equal reports whether s and other are equal.
func (s *CertPool) Equal(other *CertPool) bool
