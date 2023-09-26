// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bzip2 implements bzip2 decompression.
package bzip2

import "github.com/shogo82148/std/io"

// A StructuralError is returned when the bzip2 data is found to be
// syntactically invalid.
type StructuralError string

func (s StructuralError) Error() string

// A reader decompresses bzip2 compressed data.

// NewReader returns an io.Reader which decompresses bzip2 data from r.
// If r does not also implement [io.ByteReader],
// the decompressor may read more data than necessary from r.
func NewReader(r io.Reader) io.Reader
