// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multipart

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/textproto"
)

// A Writer generates multipart messages.
type Writer struct {
	w        io.Writer
	boundary string
	lastpart *part
}

// NewWriter returns a new multipart [Writer] with a random boundary,
// writing to w.
func NewWriter(w io.Writer) *Writer

// Boundary returns the [Writer]'s boundary.
func (w *Writer) Boundary() string

// SetBoundary overrides the [Writer]'s default randomly-generated
// boundary separator with an explicit value.
//
// SetBoundary must be called before any parts are created, may only
// contain certain ASCII characters, and must be non-empty and
// at most 70 bytes long.
func (w *Writer) SetBoundary(boundary string) error

// FormDataContentType returns the Content-Type for an HTTP
// multipart/form-data with this [Writer]'s Boundary.
func (w *Writer) FormDataContentType() string

// CreatePart creates a new multipart section with the provided
// header. The body of the part should be written to the returned
// [Writer]. After calling CreatePart, any previous part may no longer
// be written to.
func (w *Writer) CreatePart(header textproto.MIMEHeader) (io.Writer, error)

// CreateFormFile is a convenience wrapper around [Writer.CreatePart]. It creates
// a new form-data header with the provided field name and file name.
func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)

// CreateFormField calls [Writer.CreatePart] with a header using the
// given field name.
func (w *Writer) CreateFormField(fieldname string) (io.Writer, error)

// WriteField calls [Writer.CreateFormField] and then writes the given value.
func (w *Writer) WriteField(fieldname, value string) error

// Close finishes the multipart message and writes the trailing
// boundary end line to the output.
func (w *Writer) Close() error
