// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package csv reads and writes comma-separated values (CSV) files.
// There are many kinds of CSV files; this package supports the format
// described in RFC 4180.
//
// A csv file contains zero or more records of one or more fields per record.
// Each record is separated by the newline character. The final record may
// optionally be followed by a newline character.
//
//	field1,field2,field3
//
// White space is considered part of a field.
//
// Carriage returns before newline characters are silently removed.
//
// Blank lines are ignored. A line with only whitespace characters (excluding
// the ending newline character) is not considered a blank line.
//
// Fields which start and stop with the quote character " are called
// quoted-fields. The beginning and ending quote are not part of the
// field.
//
// The source:
//
//	normal string,"quoted-field"
//
// results in the fields
//
//	{`normal string`, `quoted-field`}
//
// Within a quoted-field a quote character followed by a second quote
// character is considered a single quote.
//
//	"the ""word"" is true","a ""quoted-field"""
//
// results in
//
//	{`the "word" is true`, `a "quoted-field"`}
//
// Newlines and commas may be included in a quoted-field
//
//	"Multi-line
//	field","comma is ,"
//
// results in
//
//	{`Multi-line
//	field`, `comma is ,`}
package csv

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// A ParseError is returned for parsing errors.
// Line numbers are 1-indexed and columns are 0-indexed.
type ParseError struct {
	StartLine int
	Line      int
	Column    int
	Err       error
}

func (e *ParseError) Error() string

func (e *ParseError) Unwrap() error

// These are the errors that can be returned in ParseError.Err.
var (
	ErrBareQuote  = errors.New("bare \" in non-quoted-field")
	ErrQuote      = errors.New("extraneous or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")

	// Deprecated: ErrTrailingComma is no longer used.
	ErrTrailingComma = errors.New("extra delimiter at end of line")
)

// A Reader reads records from a CSV-encoded file.
//
// As returned by NewReader, a Reader expects input conforming to RFC 4180.
// The exported fields can be changed to customize the details before the
// first call to Read or ReadAll.
//
// The Reader converts all \r\n sequences in its input to plain \n,
// including in multiline field values, so that the returned data does
// not depend on which line-ending convention an input file uses.
type Reader struct {
	Comma rune

	Comment rune

	FieldsPerRecord int

	LazyQuotes bool

	TrimLeadingSpace bool

	ReuseRecord bool

	TrailingComma bool

	r *bufio.Reader

	numLine int

	offset int64

	rawBuffer []byte

	recordBuffer []byte

	fieldIndexes []int

	fieldPositions []position

	lastRecord []string
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader

// Read reads one record (a slice of fields) from r.
// If the record has an unexpected number of fields,
// Read returns the record along with the error ErrFieldCount.
// If the record contains a field that cannot be parsed,
// Read returns a partial record along with the parse error.
// The partial record contains all fields read before the error.
// If there is no data left to be read, Read returns nil, io.EOF.
// If ReuseRecord is true, the returned slice may be shared
// between multiple calls to Read.
func (r *Reader) Read() (record []string, err error)

// FieldPos returns the line and column corresponding to
// the start of the field with the given index in the slice most recently
// returned by Read. Numbering of lines and columns starts at 1;
// columns are counted in bytes, not runes.
//
// If this is called with an out-of-bounds index, it panics.
func (r *Reader) FieldPos(field int) (line, column int)

// InputOffset returns the input stream byte offset of the current reader
// position. The offset gives the location of the end of the most recently
// read row and the beginning of the next row.
func (r *Reader) InputOffset() int64

// pos holds the position of a field in the current line.

// ReadAll reads all the remaining records from r.
// Each record is a slice of fields.
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reported.
func (r *Reader) ReadAll() (records [][]string, err error)
