// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Pipe adapter to connect code expecting an io.Reader
// with code expecting an io.Writer.

package io

import (
	"github.com/shogo82148/std/errors"
)

// ErrClosedPipeは、クローズされたパイプに対する読み取りまたは書き込み操作で使用されるエラーです。
var ErrClosedPipe = errors.New("io: read/write on closed pipe")

// PipeReaderは、パイプの読み取り側です。
type PipeReader struct{ pipe }

// Readは、標準のReadインターフェースを実装します。
// パイプからデータを読み取り、ライターが到着するか、書き込み側が閉じられるまでブロックします。
// 書き込み側がエラーで閉じられた場合、そのエラーがerrとして返されます。
// それ以外の場合、errはEOFです。
func (r *PipeReader) Read(data []byte) (n int, err error)

<<<<<<< HEAD
// Closeは、リーダーを閉じます。パイプの書き込み半分への後続の書き込みは、
// エラーErrClosedPipeを返します。
=======
// Close closes the reader; subsequent writes to the
// write half of the pipe will return the error [ErrClosedPipe].
>>>>>>> upstream/master
func (r *PipeReader) Close() error

// CloseWithErrorは、リーダーを閉じます。パイプの書き込み半分への後続の書き込みは、エラーerrを返します。
//
// CloseWithErrorは、以前にエラーが存在する場合、前のエラーを上書きせず、常にnilを返します。
func (r *PipeReader) CloseWithError(err error) error

// PipeWriterは、パイプの書き込み側です。
type PipeWriter struct{ r PipeReader }

<<<<<<< HEAD
// Writeは、標準のWriteインターフェースを実装します。
// データをパイプに書き込み、1つ以上のリーダーがすべてのデータを消費するか、
// 読み取り側が閉じられるまでブロックします。
// 読み取り側がエラーで閉じられた場合、そのエラーがerrとして返されます。
// それ以外の場合、errはErrClosedPipeです。
=======
// Write implements the standard Write interface:
// it writes data to the pipe, blocking until one or more readers
// have consumed all the data or the read end is closed.
// If the read end is closed with an error, that err is
// returned as err; otherwise err is [ErrClosedPipe].
>>>>>>> upstream/master
func (w *PipeWriter) Write(data []byte) (n int, err error)

// Close closes the writer; subsequent reads from the
// read half of the pipe will return no bytes and EOF.
func (w *PipeWriter) Close() error

// CloseWithError closes the writer; subsequent reads from the
// read half of the pipe will return no bytes and the error err,
// or EOF if err is nil.
//
// CloseWithError never overwrites the previous error if it exists
// and always returns nil.
func (w *PipeWriter) CloseWithError(err error) error

// Pipe creates a synchronous in-memory pipe.
// It can be used to connect code expecting an [io.Reader]
// with code expecting an [io.Writer].
//
// Reads and Writes on the pipe are matched one to one
// except when multiple Reads are needed to consume a single Write.
// That is, each Write to the [PipeWriter] blocks until it has satisfied
// one or more Reads from the [PipeReader] that fully consume
// the written data.
// The data is copied directly from the Write to the corresponding
// Read (or Reads); there is no internal buffering.
//
// It is safe to call Read and Write in parallel with each other or with Close.
// Parallel calls to Read and parallel calls to Write are also safe:
// the individual calls will be gated sequentially.
func Pipe() (*PipeReader, *PipeWriter)
