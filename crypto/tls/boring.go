// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

package tls

// default defaultFIPSCurvePreferences is the FIPS-allowed curves,
// in preference order (most preferable first).

// defaultCipherSuitesFIPS are the FIPS-allowed cipher suites.

// fipsSupportedSignatureAlgorithms currently are a subset of
// defaultSupportedSignatureAlgorithms without Ed25519 and SHA-1.
