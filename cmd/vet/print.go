// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the printf-checker.

package main

// isPrint records the print functions.
// If a key ends in 'f' then it is assumed to be a formatted print.

// formatState holds the parsed representation of a printf directive such as "%3.*[4]d".
// It is constructed by parsePrintfVerb.

// printfArgType encodes the types of expressions a printf verb accepts. It is a bitmask.

// Common flag sets for printf verbs.

// printVerbs identifies which flags are known to printf for each verb.

// printFormatRE is the regexp we match and report as a possible format string
// in the first argument to unformatted prints like fmt.Print.
// We exclude the space flag, so that printing a string like "x % y" is not reported as a format.
