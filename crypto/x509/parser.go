// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package x509

import (
	cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

// ParseCertificate parses a single certificate from the given ASN.1 DER data.
func ParseCertificate(der []byte) (*Certificate, error)

// ParseCertificates parses one or more certificates from the given ASN.1 DER
// data. The certificates must be concatenated with no intermediate padding.
func ParseCertificates(der []byte) ([]*Certificate, error)

// The X.509 standards confusingly 1-indexed the version names, but 0-indexed
// the actual encoded version, so the version for X.509v2 is 1.

// ParseRevocationList parses a X509 v2 Certificate Revocation List from the given
// ASN.1 DER data.
func ParseRevocationList(der []byte) (*RevocationList, error)
