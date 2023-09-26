// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the printf-checker.

package main

// isFormattedPrint records the formatted-print functions. Names are
// lower-cased so the lookup is case insensitive.

// isPrint records the unformatted-print functions. Names are lower-cased
// so the lookup is case insensitive.

// formatState holds the parsed representation of a printf directive such as "%3.*[4]d".
// It is constructed by parsePrintfVerb.

// printfArgType encodes the types of expressions a printf verb accepts. It is a bitmask.

// Common flag sets for printf verbs.

// printVerbs identifies which flags are known to printf for each verb.
