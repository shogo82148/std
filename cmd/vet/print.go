// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the printf-checker.

package main

// printfList records the formatted-print functions. The value is the location
// of the format parameter. Names are lower-cased so the lookup is
// case insensitive.

// printList records the unformatted-print functions. The value is the location
// of the first parameter to be printed.  Names are lower-cased so the lookup is
// case insensitive.

// formatState holds the parsed representation of a printf directive such as "%3.*[4]d".
// It is constructed by parsePrintfVerb.

// printfArgType encodes the types of expressions a printf verb accepts. It is a bitmask.

// Common flag sets for printf verbs.

// printVerbs identifies which flags are known to printf for each verb.
// TODO: A type that implements Formatter may do what it wants, and vet
// will complain incorrectly.
