// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/io"
)

// funcMap maps command names to functions that render their inputs safe.

// equivEscapers matches contextual escapers to equivalent template builtins.

// escaper collects type inferences about templates and changes needed to make
// templates injection safe.

// filterFailsafe is an innocuous word that is emitted in place of unsafe values
// by sanitizer functions. It is not a keyword in any programming language,
// contains no special characters, is not empty, and when it appears in output
// it is distinct enough that a developer can find the source of the problem
// via a search engine.

// redundantFuncs[a][b] implies that funcMap[b](funcMap[a](x)) == funcMap[a](x)
// for all x.

// delimEnds maps each delim to a string of characters that terminate it.

// HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.
func HTMLEscape(w io.Writer, b []byte)

// HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.
func HTMLEscapeString(s string) string

// HTMLEscaper returns the escaped HTML equivalent of the textual
// representation of its arguments.
func HTMLEscaper(args ...interface{}) string

// JSEscape writes to w the escaped JavaScript equivalent of the plain text data b.
func JSEscape(w io.Writer, b []byte)

// JSEscapeString returns the escaped JavaScript equivalent of the plain text data s.
func JSEscapeString(s string) string

// JSEscaper returns the escaped JavaScript equivalent of the textual
// representation of its arguments.
func JSEscaper(args ...interface{}) string

// URLQueryEscaper returns the escaped value of the textual representation of
// its arguments in a form suitable for embedding in a URL query.
func URLQueryEscaper(args ...interface{}) string
