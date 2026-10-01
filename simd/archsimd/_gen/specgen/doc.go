// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specgen

// ParagraphKind classifies a doc comment paragraph.
type ParagraphKind int

const (
	// SpecOwned paragraphs describe the behavioral semantics of an operation.
	SpecOwned ParagraphKind = iota

	// ImplementationNote paragraphs describe target-specific implementation
	// details (such as assembly instructions, required CPU features, emulation
	// notes, etc.).
	//
	// All implementation note paragraphs start with one of a fixed set of
	// prefixes (see [HasNotePrefix]).
	ImplementationNote

	// DirectiveComment paragraphs represent compiler or tool directives (e.g.
	// //go:noescape).
	DirectiveComment
)

func (k ParagraphKind) String() string

// Paragraph is a doc comment paragraph with its classification.
type Paragraph struct {
	Kind ParagraphKind
	Text string
}

// HasNotePrefix reports whether text starts with an implementation note prefix.
// The text should NOT include any leading "//".
func HasNotePrefix(text string) bool

// FormatComment formats a slice of paragraphs into standard Go comment text.
// The result has "//"-prefixed lines and ends with "\n" unless the whole result
// is empty.
//
// FormatComment checks that [SpecOwned] paragraphs must appear strictly before
// [ImplementationNote] paragraphs, which must appear strictly before
// [DirectiveComment]s. If the ordering rule is violated, it returns an error.
func FormatComment(paras []Paragraph) (string, error)
