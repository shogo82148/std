// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips

// CAST runs the named Cryptographic Algorithm Self-Test (if operated in FIPS
// mode) and aborts the program (stopping the module input/output and entering
// the "error state") if the self-test fails.
//
// CASTs are mandatory self-checks that must be performed by FIPS 140-3 modules
// before the algorithm is used. See Implementation Guidance 10.3.A.
//
// The name must not contain commas, colons, hashes, or equal signs.
//
// If a package p calls CAST from its init function, an import of p should also
// be added to crypto/internal/fipstest. If a package p calls CAST on the first
// use of the algorithm, an invocation of that algorithm should be added to
// fipstest.TestConditionals.
func CAST(name string, f func() error)

// PCT runs the named Pairwise Consistency Test (if operated in FIPS mode) and
// returns any errors. If an error is returned, the key must not be used.
//
// PCTs are mandatory for every key pair that is generated/imported, including
// ephemeral keys (which effectively doubles the cost of key establishment). See
// Implementation Guidance 10.3.A Additional Comment 1.
//
// The name must not contain commas, colons, hashes, or equal signs.
//
// If a package p calls PCT during key generation, an invocation of that
// function should be added to fipstest.TestConditionals.
func PCT(name string, f func() error) error
