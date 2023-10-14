// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements a regression test harness for syntax errors.
// The files in the testdata directory are parsed and the reported
// errors are compared against the errors declared in those files.
//
// Errors are declared in place in the form of "error comments",
// just before (or on the same line as) the offending token.
//
// Error comments must be of the form // ERROR rx or /* ERROR rx */
// where rx is a regular expression that matches the reported error
// message. The rx text comprises the comment text after "ERROR ",
// with any white space around it stripped.
//
// If the line comment form is used, the reported error's line must
// match the line of the error comment.
//
// If the regular comment form is used, the reported error's position
// must match the position of the token immediately following the
// error comment. Thus, /* ERROR ... */ comments should appear
// immediately before the position where the error is reported.
//
// Currently, the test harness only supports one error comment per
// token. If multiple error comments appear before a token, only
// the last one is considered.

package syntax
