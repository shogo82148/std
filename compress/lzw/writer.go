// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lzw

import (
	"github.com/shogo82148/std/io"
)

// A writer is a buffered, flushable writer.

// An errWriteCloser is an io.WriteCloser that always returns a given error.

// encoder is LZW compressor.

// errOutOfCodes is an internal error that means that the encoder has run out
// of unused codes and a clear code needs to be sent next.

// NewWriter creates a new io.WriteCloser.
// Writes to the returned io.WriteCloser are compressed and written to w.
// It is the caller's responsibility to call Close on the WriteCloser when
// finished writing.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8. Input bytes must be less than 1<<litWidth.
func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
