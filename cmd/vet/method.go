// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the code to check canonical methods.

package main

type MethodSig struct {
	args    []string
	results []string
}

// canonicalMethods lists the input and output types for Go methods
// that are checked using dynamic interface checks.  Because the
// checks are dynamic, such methods would not cause a compile error
// if they have the wrong signature: instead the dynamic check would
// fail, sometimes mysteriously.  If a method is found with a name listed
// here but not the input/output types listed here, vet complains.
//
// A few of the canonical methods have very common names.
// For example, a type might implement a Scan method that
// has nothing to do with fmt.Scanner, but we still want to check
// the methods that are intended to implement fmt.Scanner.
// To do that, the arguments that have a = prefix are treated as
// signals that the canonical meaning is intended: if a Scan
// method doesn't have a fmt.ScanState as its first argument,
// we let it go.  But if it does have a fmt.ScanState, then the
// rest has to match.
