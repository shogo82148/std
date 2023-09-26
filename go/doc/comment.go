// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Godoc comment extraction and comment -> HTML formatting.

package doc

import (
	"github.com/shogo82148/std/io"
)

// ToHTML converts comment text to formatted HTML.
// The comment was prepared by DocReader,
// so it is known not to have leading, trailing blank lines
// nor to have trailing spaces at the end of lines.
// The comment markers have already been removed.
//
// Turn each run of multiple \n into </p><p>.
// Turn each run of indented lines into a <pre> block without indent.
// Enclose headings with header tags.
//
// URLs in the comment text are converted into links; if the URL also appears
// in the words map, the link is taken from the map (if the corresponding map
// value is the empty string, the URL is not converted into a link).
//
// Go identifiers that appear in the words map are italicized; if the corresponding
// map value is not the empty string, it is considered a URL and the word is converted
// into a link.
func ToHTML(w io.Writer, text string, words map[string]string)

// ToText prepares comment text for presentation in textual output.
// It wraps paragraphs of text to width or fewer Unicode code points
// and then prefixes each line with the indent.  In preformatted sections
// (such as program text), it prefixes each non-blank line with preIndent.
func ToText(w io.Writer, text string, indent, preIndent string, width int)
