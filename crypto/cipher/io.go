// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

import "github.com/shogo82148/std/io"

// StreamReader wraps a Stream into an io.Reader. It calls XORKeyStream
// to process each slice of data which passes through.
type StreamReader struct {
	S Stream
	R io.Reader
}

func (r StreamReader) Read(dst []byte) (n int, err error)

// StreamWriter wraps a Stream into an io.Writer. It calls XORKeyStream
// to process each slice of data which passes through. If any Write call
// returns short then the StreamWriter is out of sync and must be discarded.
type StreamWriter struct {
	S   Stream
	W   io.Writer
	Err error
}

func (w StreamWriter) Write(src []byte) (n int, err error)

func (w StreamWriter) Close() error
