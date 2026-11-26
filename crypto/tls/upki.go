// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/errors"
)

// UPKIRevocationConfig controls local, offline UPKI revocation checking of
// server certificates.
//
// It doesn't configure downloading or updating revocation data.
type UPKIRevocationConfig struct {
	// Disabled disables UPKI revocation checking, including automatic
	// discovery of a platform-default cache.
	Disabled bool

	// Directory is a C2SP upki-revocation cache directory. It contains
	// index.bin directly.
	//
	// If Directory is empty, platform-default directories are searched for
	// index.bin in user-before-system order. Discovery skips missing paths,
	// non-directory path components, and paths the process lacks permission
	// to open.
	Directory string

	// MissingData controls the result when a lookup needs an index but its
	// path is missing, a parent is not a directory, or access is denied.
	// By default, the connection proceeds if automatic discovery finds no
	// available index. [UPKIRevocationPolicyReject] requires an available index.
	// Errors reading or parsing an opened index or its filters are fatal.
	MissingData UPKIRevocationPolicy

	// NotCovered controls the result when the revocation data cannot
	// determine a certificate's status, including when covering filters have
	// no information about the certificate's issuer. By default the connection
	// proceeds. [UPKIRevocationPolicyReject] produces an error if there isn't
	// sufficient coverage to make a revocation determination.
	NotCovered UPKIRevocationPolicy

	// MissingSCTs controls the result when a certificate contains no embedded
	// Signed Certificate Timestamps. A malformed embedded SCT extension is an
	// error if a cache is explicitly configured or found during discovery.
	// By default a certificate without SCTs proceeds without error.
	// [UPKIRevocationPolicyReject] produces an error if the SCTs required for
	// a revocation determination aren't present.
	MissingSCTs UPKIRevocationPolicy
}

// ErrUPKIRevoked is returned as a wrapped [CertificateVerificationError]
// when a certificate is determined to be revoked.
var ErrUPKIRevoked = errors.New("certificate is revoked according to UPKI")

// UPKIRevocationPolicy controls the result of an inconclusive UPKI revocation
// check.
type UPKIRevocationPolicy uint8

const (
	// UPKIRevocationPolicyDefault uses the default for the corresponding
	// UPKIRevocationConfig field.
	UPKIRevocationPolicyDefault UPKIRevocationPolicy = iota

	// UPKIRevocationPolicyAllow allows the connection to proceed when the
	// revocation check is inconclusive for the corresponding reason.
	UPKIRevocationPolicyAllow

	// UPKIRevocationPolicyReject rejects the connection when the revocation
	// check is inconclusive for the corresponding reason.
	UPKIRevocationPolicyReject
)
