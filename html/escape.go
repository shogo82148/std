// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package html provides functions for escaping and unescaping HTML text.
package html

// These replacements permit compatibility with old numeric entities that
// assumed Windows-1252 encoding.
// http://www.whatwg.org/specs/web-apps/current-work/multipage/tokenization.html#consume-a-character-reference

// EscapeString escapes special characters like "<" to become "&lt;". It
// escapes only five such characters: amp, apos, lt, gt and quot.
// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
// always true.
func EscapeString(s string) string

// UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a
// larger range of entities than EscapeString escapes. For example, "&aacute;"
// unescapes to "á", as does "&#225;" and "&xE1;".
// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
// always true.
func UnescapeString(s string) string
