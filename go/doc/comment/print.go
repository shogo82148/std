// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

// A Printer is a doc comment printer.
// The fields in the struct can be filled in before calling
// any of the printing methods
// in order to customize the details of the printing process.
type Printer struct {
	// HeadingLevel is the nesting level used for
	// HTML and Markdown headings.
	// If HeadingLevel is zero, it defaults to level 3,
	// meaning to use <h3> and ###.
	HeadingLevel int

	// HeadingID is a function that computes the heading ID
	// (anchor tag) to use for the heading h when generating
	// HTML and Markdown. If HeadingID returns an empty string,
	// then the heading ID is omitted.
	// If HeadingID is nil, h.DefaultID is used.
	HeadingID func(h *Heading) string

	// DocLinkURL is a function that computes the URL for the given DocLink.
	// If DocLinkURL is nil, then link.DefaultURL(p.DocLinkBaseURL) is used.
	DocLinkURL func(link *DocLink) string

	// DocLinkBaseURL is used when DocLinkURL is nil,
	// passed to [DocLink.DefaultURL] to construct a DocLink's URL.
	// See that method's documentation for details.
	DocLinkBaseURL string

	// TextPrefix is a prefix to print at the start of every line
	// when generating text output using the Text method.
	TextPrefix string

	// TextCodePrefix is the prefix to print at the start of each
	// preformatted (code block) line when generating text output,
	// instead of (not in addition to) TextPrefix.
	// If TextCodePrefix is the empty string, it defaults to TextPrefix+"\t".
	TextCodePrefix string

	// TextWidth is the maximum width text line to generate,
	// measured in Unicode code points,
	// excluding TextPrefix and the newline character.
	// If TextWidth is zero, it defaults to 80 minus the number of code points in TextPrefix.
	// If TextWidth is negative, there is no limit.
	TextWidth int
}

// DefaultURL constructs and returns the documentation URL for l,
// using baseURL as a prefix for links to other packages.
//
// The possible forms returned by DefaultURL are:
//   - baseURL/ImportPath, for a link to another package
//   - baseURL/ImportPath#Name, for a link to a const, func, type, or var in another package
//   - baseURL/ImportPath#Recv.Name, for a link to a method in another package
//   - #Name, for a link to a const, func, type, or var in this package
//   - #Recv.Name, for a link to a method in this package
//
// If baseURL ends in a trailing slash, then DefaultURL inserts
// a slash between ImportPath and # in the anchored forms.
// For example, here are some baseURL values and URLs they can generate:
//
//	"/pkg/" → "/pkg/math/#Sqrt"
//	"/pkg"  → "/pkg/math#Sqrt"
//	"/"     → "/math/#Sqrt"
//	""      → "/math#Sqrt"
func (l *DocLink) DefaultURL(baseURL string) string

// DefaultID returns the default anchor ID for the heading h.
//
// The default anchor ID is constructed by converting every
// rune that is not alphanumeric ASCII to an underscore
// and then adding the prefix “hdr-”.
// For example, if the heading text is “Go Doc Comments”,
// the default ID is “hdr-Go_Doc_Comments”.
func (h *Heading) DefaultID() string

// Comment returns the standard Go formatting of the [Doc],
// without any comment markers.
func (p *Printer) Comment(d *Doc) []byte
