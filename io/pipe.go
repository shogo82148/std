// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Pipe adapter to connect code expecting an io.Reader
// with code expecting an io.Writer.

package io

import (
	"github.com/shogo82148/std/errors"
)

// ErrClosedPipe is the error used for read or write operations on a closed pipe.
var ErrClosedPipe = errors.New("io: read/write on closed pipe")

// A pipe is the shared pipe structure underlying PipeReader and PipeWriter.

// A PipeReader is the read half of a pipe.
type PipeReader struct {
	p *pipe
}

// Read implements the standard Read interface:
// it reads data from the pipe, blocking until a writer
// arrives or the write end is closed.
// If the write end is closed with an error, that error is
// returned as err; otherwise err is EOF.
func (r *PipeReader) Read(data []byte) (n int, err error)

// Close closes the reader; subsequent writes to the
// write half of the pipe will return the error ErrClosedPipe.
func (r *PipeReader) Close() error

// CloseWithError closes the reader; subsequent writes
// to the write half of the pipe will return the error err.
func (r *PipeReader) CloseWithError(err error) error

// A PipeWriter is the write half of a pipe.
type PipeWriter struct {
	p *pipe
}

// Write implements the standard Write interface:
// it writes data to the pipe, blocking until readers
// have consumed all the data or the read end is closed.
// If the read end is closed with an error, that err is
// returned as err; otherwise err is ErrClosedPipe.
func (w *PipeWriter) Write(data []byte) (n int, err error)

// Close closes the writer; subsequent reads from the
// read half of the pipe will return no bytes and EOF.
func (w *PipeWriter) Close() error

// CloseWithError closes the writer; subsequent reads from the
// read half of the pipe will return no bytes and the error err.
func (w *PipeWriter) CloseWithError(err error) error

// Pipe creates a synchronous in-memory pipe.
// It can be used to connect code expecting an io.Reader
// with code expecting an io.Writer.
// Reads on one end are matched with writes on the other,
// copying data directly between the two; there is no internal buffering.
// It is safe to call Read and Write in parallel with each other or with
// Close. Close will complete once pending I/O is done. Parallel calls to
// Read, and parallel calls to Write, are also safe:
// the individual calls will be gated sequentially.
func Pipe() (*PipeReader, *PipeWriter)
