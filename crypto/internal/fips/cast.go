// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips

// CAST runs the named Cryptographic Algorithm Self-Test (if compiled and
// operated in FIPS mode) and aborts the program (stopping the module
// input/output and entering the "error state") if the self-test fails.
//
// These are mandatory self-checks that must be performed by FIPS 140-3 modules
// before the algorithm is used. See Implementation Guidance 10.3.A.
//
// The name must not contain commas, colons, hashes, or equal signs.
//
// When calling this function, also add the calling package to cast_test.go.
func CAST(name string, f func() error)
