// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt_test

import (
	. "fmt"
)

type ScanTest struct {
	text string
	in   interface{}
	out  interface{}
}

type ScanfTest struct {
	format string
	text   string
	in     interface{}
	out    interface{}
}

type ScanfMultiTest struct {
	format string
	text   string
	in     []interface{}
	out    []interface{}
	err    string
}

type FloatTest struct {
	text string
	in   float64
	out  float64
}

// Xs accepts any non-empty run of the verb character
type Xs string

// IntString accepts an integer followed immediately by a string.
// It tests the embedding of a scan within a scan.
type IntString struct {
	i int
	s string
}

// myStringReader implements Read but not ReadRune, allowing us to test our readRune wrapper
// type that creates something that can read runes given only Read().

// eofCounter is a special Reader that counts reads at end of file.

type TwoLines string

// simpleReader is a strings.Reader that implements only Read, not ReadRune.
// Good for testing readahead.

// runeScanner implements the Scanner interface for TestScanStateCount.

// RecursiveInt accepts a string matching %d.%d.%d....
// and parses it into a linked list.
// It allows us to benchmark recursive descent style scanners.
type RecursiveInt struct {
	i    int
	next *RecursiveInt
}

// 800 is small enough to not overflow the stack when using gccgo on a
// platform that does not support split stack.
