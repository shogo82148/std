// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fipstest collects external tests that would ordinarily live in
// crypto/internal/fips140/... packages. That tree gets snapshot at each
// validation, while we want tests to evolve and still apply to all versions of
// the module. Also, we can't fix failing tests in a module snapshot, so we need
// to either minimize, skip, or remove them. Finally, the module needs to avoid
// importing internal packages like testenv and cryptotest to avoid locking in
// their APIs.
//
// Also, this package includes the ACVP and functional testing harnesses.
package fipstest
