// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cookiejar

// tNow is the synthetic current time used as now during testing.

// testPSL implements PublicSuffixList with just two rules: "co.uk"
// and the default rule "*".
// The implementation has two intentional bugs:
//
//	PublicSuffix("www.buggy.psl") == "xy"
//	PublicSuffix("www2.buggy.psl") == "com"

// jarTest encapsulates the following actions on a jar:
//  1. Perform SetCookies with fromURL and the cookies from setCookies.
//     (Done at time tNow + 0 ms.)
//  2. Check that the entries in the jar matches content.
//     (Done at time tNow + 1001 ms.)
//  3. For each query in tests: Check that Cookies with toURL yields the
//     cookies in want.
//     (Query n done at tNow + (n+2)*1001 ms.)

// query contains one test of the cookies returned from Jar.Cookies.

// basicsTests contains fundamental tests. Each jarTest has to be performed on
// a fresh, empty Jar.

// updateAndDeleteTests contains jarTests which must be performed on the same
// Jar.

// chromiumBasicsTests contains fundamental tests. Each jarTest has to be
// performed on a fresh, empty Jar.

// chromiumDomainTests contains jarTests which must be executed all on the
// same Jar.

// chromiumDeletionTests must be performed all on the same Jar.

// domainHandlingTests tests and documents the rules for domain handling.
// Each test must be performed on an empty new Jar.
