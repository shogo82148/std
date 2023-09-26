// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iotest

import (
	"github.com/shogo82148/std/io"
)

// NewWriteLogger returns a writer that behaves like w except
// that it logs (using log.Printf) each write to standard error,
// printing the prefix and the hexadecimal data written.
func NewWriteLogger(prefix string, w io.Writer) io.Writer

// NewReadLogger returns a reader that behaves like r except
// that it logs (using log.Print) each read to standard error,
// printing the prefix and the hexadecimal data read.
func NewReadLogger(prefix string, r io.Reader) io.Reader
