// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comment

// A Doc is a parsed Go doc comment.
type Doc struct {
	Content []Block

	Links []*LinkDef
}

// A LinkDef is a single link definition.
type LinkDef struct {
	Text string
	URL  string
	Used bool
}

// A Block is block-level content in a doc comment,
// one of [*Code], [*Heading], [*List], or [*Paragraph].
type Block interface {
	block()
}

// A Heading is a doc comment heading.
type Heading struct {
	Text []Text
}

// A List is a numbered or bullet list.
// Lists are always non-empty: len(Items) > 0.
// In a numbered list, every Items[i].Number is a non-empty string.
// In a bullet list, every Items[i].Number is an empty string.
type List struct {
	Items []*ListItem

	ForceBlankBefore bool

	ForceBlankBetween bool
}

// BlankBefore reports whether a reformatting of the comment
// should include a blank line before the list.
// The default rule is the same as for [BlankBetween]:
// if the list item content contains any blank lines
// (meaning at least one item has multiple paragraphs)
// then the list itself must be preceded by a blank line.
// A preceding blank line can be forced by setting [List].ForceBlankBefore.
func (l *List) BlankBefore() bool

// BlankBetween reports whether a reformatting of the comment
// should include a blank line between each pair of list items.
// The default rule is that if the list item content contains any blank lines
// (meaning at least one item has multiple paragraphs)
// then list items must themselves be separated by blank lines.
// Blank line separators can be forced by setting [List].ForceBlankBetween.
func (l *List) BlankBetween() bool

// A ListItem is a single item in a numbered or bullet list.
type ListItem struct {
	Number string

	Content []Block
}

// A Paragraph is a paragraph of text.
type Paragraph struct {
	Text []Text
}

// A Code is a preformatted code block.
type Code struct {
	Text string
}

// A Text is text-level content in a doc comment,
// one of [Plain], [Italic], [*Link], or [*DocLink].
type Text interface {
	text()
}

// A Plain is a string rendered as plain text (not italicized).
type Plain string

// An Italic is a string rendered as italicized text.
type Italic string

// A Link is a link to a specific URL.
type Link struct {
	Auto bool
	Text []Text
	URL  string
}

// A DocLink is a link to documentation for a Go package or symbol.
type DocLink struct {
	Text []Text

	ImportPath string
	Recv       string
	Name       string
}

// A Parser is a doc comment parser.
// The fields in the struct can be filled in before calling Parse
// in order to customize the details of the parsing process.
type Parser struct {
	Words map[string]string

	LookupPackage func(name string) (importPath string, ok bool)

	LookupSym func(recv, name string) (ok bool)
}

// parseDoc is parsing state for a single doc comment.

// DefaultLookupPackage is the default package lookup
// function, used when [Parser].LookupPackage is nil.
// It recognizes names of the packages from the standard
// library with single-element import paths, such as math,
// which would otherwise be impossible to name.
//
// Note that the go/doc package provides a more sophisticated
// lookup based on the imports used in the current package.
func DefaultLookupPackage(name string) (importPath string, ok bool)

// Parse parses the doc comment text and returns the *Doc form.
// Comment markers (/* // and */) in the text must have already been removed.
func (p *Parser) Parse(text string) *Doc

// A span represents a single span of comment lines (lines[start:end])
// of an identified kind (code, heading, paragraph, and so on).

// A spanKind describes the kind of span.

const (
	_ spanKind = iota
)
