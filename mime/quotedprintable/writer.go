// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quotedprintable

import "github.com/shogo82148/std/io"

// A Writer is a quoted-printable writer that implements io.WriteCloser.
type Writer struct {
	Binary bool

	w    io.Writer
	i    int
	line [78]byte
	cr   bool
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer

// Write encodes p using quoted-printable encoding and writes it to the
// underlying io.Writer. It limits line length to 76 characters. The encoded
// bytes are not necessarily flushed until the Writer is closed.
func (w *Writer) Write(p []byte) (n int, err error)

// Close closes the Writer, flushing any unwritten data to the underlying
// io.Writer, but does not close the underlying io.Writer.
func (w *Writer) Close() error
