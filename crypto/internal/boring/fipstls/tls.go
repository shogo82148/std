// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

// Package fipstls allows control over whether crypto/tls requires FIPS-approved settings.
// This package only exists with GOEXPERIMENT=boringcrypto, but the effects are independent
// of the use of BoringCrypto.
package fipstls

// Force forces crypto/tls to restrict TLS configurations to FIPS-approved settings.
// By design, this call is impossible to undo (except in tests).
//
// Note that this call has an effect even in programs using
// standard crypto (that is, even when Enabled = false).
func Force()

// Abandon allows non-FIPS-approved settings.
// If called from a non-test binary, it panics.
func Abandon()

// Required reports whether FIPS-approved settings are required.
func Required() bool
