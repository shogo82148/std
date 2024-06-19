// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// ParseCertificate parses a single certificate from the given ASN.1 DER data.
//
// Before Go 1.23, ParseCertificate accepted certificates with negative serial
// numbers. This behavior can be restored by including "x509negativeserial=1" in
// the GODEBUG environment variable.
func ParseCertificate(der []byte) (*Certificate, error)

// ParseCertificates parses one or more certificates from the given ASN.1 DER
// data. The certificates must be concatenated with no intermediate padding.
func ParseCertificates(der []byte) ([]*Certificate, error)

// ParseRevocationList parses a X509 v2 [Certificate] Revocation List from the given
// ASN.1 DER data.
func ParseRevocationList(der []byte) (*RevocationList, error)
