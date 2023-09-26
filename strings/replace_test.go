// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings_test

import (
	"log"
	. "strings"
)

var _ = log.Printf

type ReplacerTest struct {
	r   *Replacer
	in  string
	out string
}

var ReplacerTests = []ReplacerTest{

	{htmlEscaper, "No changes", "No changes"},
	{htmlEscaper, "I <3 escaping & stuff", "I &lt;3 escaping &amp; stuff"},
	{htmlEscaper, "&&&", "&amp;&amp;&amp;"},

	{replacer, "fooaaabar", "foo3[aaa]b1[a]r"},
	{replacer, "long, longerst, longer", "short, most long, medium"},
	{replacer, "XiX", "YiY"},

	{capitalLetters, "brad", "BrAd"},
	{capitalLetters, Repeat("a", (32<<10)+123), Repeat("A", (32<<10)+123)},

	{blankToXReplacer, "oo", "XOXOX"},
}

// pickAlgorithmTest is a test that verifies that given input for a
// Replacer that we pick the correct algorithm.
