// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// CertPool is a set of certificates.
type CertPool struct {
	byName map[string][]int

	lazyCerts []lazyCert

	haveSum map[sum224]bool

	systemPool bool
}

// lazyCert is minimal metadata about a Cert and a func to retrieve it
// in its normal expanded *Certificate form.

// NewCertPool returns a new, empty CertPool.
func NewCertPool() *CertPool

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
