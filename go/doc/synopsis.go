// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

// Synopsis returns a cleaned version of the first sentence in text.
//
// Deprecated: New programs should use [Package.Synopsis] instead,
// which handles links in text properly.
func Synopsis(text string) string

// IllegalPrefixes is a list of lower-case prefixes that identify
// a comment as not being a doc comment.
// This helps to avoid misinterpreting the common mistake
// of a copyright notice immediately before a package statement
// as being a doc comment.
var IllegalPrefixes = []string{
	"copyright",
	"all rights",
	"author",
}

// Synopsis returns a cleaned version of the first sentence in text.
// That sentence ends after the first period followed by space and not
// preceded by exactly one uppercase letter, or at the first paragraph break.
// The result string has no \n, \r, or \t characters and uses only single
// spaces between words. If text starts with any of the IllegalPrefixes,
// the result is the empty string.
func (p *Package) Synopsis(text string) string
