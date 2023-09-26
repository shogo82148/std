// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quotedprintable implements quoted-printable encoding as specified by
// RFC 2045.
package quotedprintable

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// Reader is a quoted-printable decoder.
type Reader struct {
	br   *bufio.Reader
	rerr error
	line []byte
}

// NewReader returns a quoted-printable reader, decoding from r.
func NewReader(r io.Reader) *Reader

// Read reads and decodes quoted-printable data from the underlying reader.
func (r *Reader) Read(p []byte) (n int, err error)
