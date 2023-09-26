// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements a typechecker test harness. The packages specified
// in tests are typechecked. Error messages reported by the typechecker are
// compared against the error messages expected in the test files.
//
// Expected errors are indicated in the test files by putting a comment
// of the form /* ERROR "rx" */ immediately following an offending token.
// The harness will verify that an error matching the regular expression
// rx is reported at that source position. Consecutive comments may be
// used to indicate multiple errors for the same token position.
//
// For instance, the following test file indicates that a "not declared"
// error should be reported for the undeclared variable x:
//
//	package p
//	func f() {
//		_ = x /* ERROR "not declared" */ + 1
//	}

package types_test

import (
	. "go/types"
)

// Positioned errors are of the form filename:line:column: message .

// ERROR comments must start with text `ERROR "rx"` or `ERROR rx` where
// rx is a regular expression that matches the expected error message.
// Space around "rx" or rx is ignored. Use the form `ERROR HERE "rx"`
// for error messages that are located immediately after rather than
// at a token's position.
//

// goVersionRx matches a Go version string using '_', e.g. "go1_12".

// excludedForUnifiedBuild lists files that cannot be tested
// when using the unified build's export data.
// TODO(gri) enable as soon as the unified build supports this.
