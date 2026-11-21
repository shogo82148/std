// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upki

import (
	"github.com/shogo82148/std/crypto/x509"
	"github.com/shogo82148/std/os"
)

// Result is the aggregate result of checking a leaf certificate across all
// of its verified chains.
type Result uint8

const (
	// NotCovered means the cache could not determine the certificate's status.
	// Reasons include: the leaf has no embedded SCTs, no filter covers any of
	// the SCT (log ID, timestamp) pairs, the covering filters' own coverage
	// disagrees with the index, or covering filters were found but every one
	// reported the issuer as not enrolled.
	NotCovered Result = iota

	// Revoked means a covering filter listed the certificate as revoked.
	Revoked

	// NotRevoked means at least one covering filter answered conclusively that
	// the certificate is not in its revoked set, and no covering filter listed
	// it as revoked.
	NotRevoked
)

// Check checks the end-entity certificate in chains against the C2SP
// revocation cache in dir. The chains must already have been verified. Check
// takes ownership of indexFile and closes it before returning.
func Check(dir string, indexFile *os.File, chains [][]*x509.Certificate) (Result, error)

// ExtractSCTs decodes the RFC 6962 SignedCertificateTimestampList
// extension from cert and returns one [SCT] per entry.
//
// Returns (nil, nil) when the extension is absent.
//
// Entries using unknown SCT versions are ignored.
//
// It returns errMalformedSCTExtension if the extension is present but malformed.
//
// See RFC 6962 §3.3 for more information.
func ExtractSCTs(cert *x509.Certificate) ([]SCT, error)

// SCT is a parsed Signed Certificate Timestamp from a leaf certificate's
// embedded SCT extension.
//
// It carries the two fields CRLite needs to pick covering filters: the
// CT log's identifier and the log's reported issuance timestamp.
//
// Important: The SCT's signature is intentionally not retained.
// In the revocation checking context we assume the CT log inclusion
// proof was already validated.
type SCT struct {
	logID     logID
	timestamp timestamp
}
