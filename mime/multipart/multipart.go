// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

/*
Package multipart implements MIME multipart parsing, as defined in RFC
2046.

The implementation is sufficient for HTTP (RFC 2388) and the multipart
bodies generated by popular browsers.
*/
package multipart

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/textproto"
)

// TODO(bradfitz): inline these once the compiler can inline them in
// read-only situation (such as bytes.HasSuffix)

// A Part represents a single part in a multipart body.
type Part struct {
	Header textproto.MIMEHeader

	buffer *bytes.Buffer
	mr     *Reader

	disposition       string
	dispositionParams map[string]string
}

// FormName returns the name parameter if p has a Content-Disposition
// of type "form-data".  Otherwise it returns the empty string.
func (p *Part) FormName() string

// FileName returns the filename parameter of the Part's
// Content-Disposition header.
func (p *Part) FileName() string

// NewReader creates a new multipart Reader reading from r using the
// given MIME boundary.
func NewReader(reader io.Reader, boundary string) *Reader

// Read reads the body of a part, after its headers and before the
// next part (if any) begins.
func (p *Part) Read(d []byte) (n int, err error)

func (p *Part) Close() error

// Reader is an iterator over parts in a MIME multipart body.
// Reader's underlying parser consumes its input as needed.  Seeking
// isn't supported.
type Reader struct {
	bufReader *bufio.Reader

	currentPart *Part
	partsRead   int

	nl, nlDashBoundary, dashBoundaryDash, dashBoundary []byte
}

// NextPart returns the next part in the multipart or an error.
// When there are no more parts, the error io.EOF is returned.
func (r *Reader) NextPart() (*Part, error)
