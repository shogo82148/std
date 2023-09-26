// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// A Reader implements convenience methods for reading requests
// or responses from a text protocol network connection.
type Reader struct {
	R   *bufio.Reader
	dot *dotReader
	buf []byte
}

// NewReader returns a new Reader reading from r.
func NewReader(r *bufio.Reader) *Reader

// ReadLine reads a single line from r,
// eliding the final \n or \r\n from the returned string.
func (r *Reader) ReadLine() (string, error)

// ReadLineBytes is like ReadLine but returns a []byte instead of a string.
func (r *Reader) ReadLineBytes() ([]byte, error)

// ReadContinuedLine reads a possibly continued line from r,
// eliding the final trailing ASCII white space.
// Lines after the first are considered continuations if they
// begin with a space or tab character.  In the returned data,
// continuation lines are separated from the previous line
// only by a single space: the newline and leading white space
// are removed.
//
// For example, consider this input:
//
//	Line 1
//	  continued...
//	Line 2
//
// The first call to ReadContinuedLine will return "Line 1 continued..."
// and the second will return "Line 2".
//
// A line consisting of only white space is never continued.
func (r *Reader) ReadContinuedLine() (string, error)

// ReadContinuedLineBytes is like ReadContinuedLine but
// returns a []byte instead of a string.
func (r *Reader) ReadContinuedLineBytes() ([]byte, error)

// ReadCodeLine reads a response code line of the form
//
//	code message
//
// where code is a three-digit status code and the message
// extends to the rest of the line.  An example of such a line is:
//
//	220 plan9.bell-labs.com ESMTP
//
// If the prefix of the status does not match the digits in expectCode,
// ReadCodeLine returns with err set to &Error{code, message}.
// For example, if expectCode is 31, an error will be returned if
// the status is not in the range [310,319].
//
// If the response is multi-line, ReadCodeLine returns an error.
//
// An expectCode <= 0 disables the check of the status code.
func (r *Reader) ReadCodeLine(expectCode int) (code int, message string, err error)

// ReadResponse reads a multi-line response of the form:
//
//	code-message line 1
//	code-message line 2
//	...
//	code message line n
//
// where code is a three-digit status code. The first line starts with the
// code and a hyphen. The response is terminated by a line that starts
// with the same code followed by a space. Each line in message is
// separated by a newline (\n).
//
// See page 36 of RFC 959 (http://www.ietf.org/rfc/rfc959.txt) for
// details.
//
// If the prefix of the status does not match the digits in expectCode,
// ReadResponse returns with err set to &Error{code, message}.
// For example, if expectCode is 31, an error will be returned if
// the status is not in the range [310,319].
//
// An expectCode <= 0 disables the check of the status code.
func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error)

// DotReader returns a new Reader that satisfies Reads using the
// decoded text of a dot-encoded block read from r.
// The returned Reader is only valid until the next call
// to a method on r.
//
// Dot encoding is a common framing used for data blocks
// in text protocols such as SMTP.  The data consists of a sequence
// of lines, each of which ends in "\r\n".  The sequence itself
// ends at a line containing just a dot: ".\r\n".  Lines beginning
// with a dot are escaped with an additional dot to avoid
// looking like the end of the sequence.
//
// The decoded form returned by the Reader's Read method
// rewrites the "\r\n" line endings into the simpler "\n",
// removes leading dot escapes if present, and stops with error io.EOF
// after consuming (and discarding) the end-of-sequence line.
func (r *Reader) DotReader() io.Reader

// ReadDotBytes reads a dot-encoding and returns the decoded data.
//
// See the documentation for the DotReader method for details about dot-encoding.
func (r *Reader) ReadDotBytes() ([]byte, error)

// ReadDotLines reads a dot-encoding and returns a slice
// containing the decoded lines, with the final \r\n or \n elided from each.
//
// See the documentation for the DotReader method for details about dot-encoding.
func (r *Reader) ReadDotLines() ([]string, error)

// ReadMIMEHeader reads a MIME-style header from r.
// The header is a sequence of possibly continued Key: Value lines
// ending in a blank line.
// The returned map m maps CanonicalMIMEHeaderKey(key) to a
// sequence of values in the same order encountered in the input.
//
// For example, consider this input:
//
//	My-Key: Value 1
//	Long-Key: Even
//	       Longer Value
//	My-Key: Value 2
//
// Given that input, ReadMIMEHeader returns the map:
//
//	map[string][]string{
//		"My-Key": {"Value 1", "Value 2"},
//		"Long-Key": {"Even Longer Value"},
//	}
func (r *Reader) ReadMIMEHeader() (MIMEHeader, error)

// CanonicalMIMEHeaderKey returns the canonical format of the
// MIME header key s.  The canonicalization converts the first
// letter and any letter following a hyphen to upper case;
// the rest are converted to lowercase.  For example, the
// canonical key for "accept-encoding" is "Accept-Encoding".
// MIME header keys are assumed to be ASCII only.
func CanonicalMIMEHeaderKey(s string) string
