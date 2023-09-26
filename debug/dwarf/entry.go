// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DWARF debug information entry parser.
// An entry is a sequence of data items of a given format.
// The first word in the entry is an index into what DWARF
// calls the ``abbreviation table.''  An abbreviation is really
// just a type descriptor: it's an array of attribute tag/value format pairs.

package dwarf

// a single entry's description: a sequence of attributes

// a map from entry format ids to their descriptions

// An entry is a sequence of attribute/value pairs.
type Entry struct {
	Offset   Offset
	Tag      Tag
	Children bool
	Field    []Field
}

// A Field is a single attribute/value pair in an Entry.
type Field struct {
	Attr Attr
	Val  interface{}
}

// Val returns the value associated with attribute Attr in Entry,
// or nil if there is no such attribute.
//
// A common idiom is to merge the check for nil return with
// the check that the value has the expected dynamic type, as in:
//
//	v, ok := e.Val(AttrSibling).(int64);
func (e *Entry) Val(a Attr) interface{}

// An Offset represents the location of an Entry within the DWARF info.
// (See Reader.Seek.)
type Offset uint32

// A Reader allows reading Entry structures from a DWARF “info” section.
// The Entry structures are arranged in a tree.  The Reader's Next function
// return successive entries from a pre-order traversal of the tree.
// If an entry has children, its Children field will be true, and the children
// follow, terminated by an Entry with Tag 0.
type Reader struct {
	b            buf
	d            *Data
	err          error
	unit         int
	lastChildren bool
	lastSibling  Offset
}

// Reader returns a new Reader for Data.
// The reader is positioned at byte offset 0 in the DWARF “info” section.
func (d *Data) Reader() *Reader

// Seek positions the Reader at offset off in the encoded entry stream.
// Offset 0 can be used to denote the first entry.
func (r *Reader) Seek(off Offset)

// Next reads the next entry from the encoded entry stream.
// It returns nil, nil when it reaches the end of the section.
// It returns an error if the current offset is invalid or the data at the
// offset cannot be decoded as a valid Entry.
func (r *Reader) Next() (*Entry, error)

// SkipChildren skips over the child entries associated with
// the last Entry returned by Next.  If that Entry did not have
// children or Next has not been called, SkipChildren is a no-op.
func (r *Reader) SkipChildren()
