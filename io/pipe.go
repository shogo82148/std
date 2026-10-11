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

// Closeは、リーダーを閉じます。パイプの書き込み半分への後続の書き込みは、
// エラー [ErrClosedPipe] を返します。
func (r *PipeReader) Close() error

// CloseWithErrorは、リーダーを閉じます。パイプの書き込み半分への後続の書き込みは、エラーerrを返します。
//
// CloseWithErrorは、以前にエラーが存在する場合、前のエラーを上書きせず、常にnilを返します。
func (r *PipeReader) CloseWithError(err error) error

// PipeWriterは、パイプの書き込み側です。
type PipeWriter struct{ r PipeReader }

// Writeは、標準のWriteインターフェースを実装します。
// データをパイプに書き込み、1つ以上のリーダーがすべてのデータを消費するか、
// 読み取り側が閉じられるまでブロックします。
// 読み取り側がエラーで閉じられた場合、そのエラーがerrとして返されます。
// それ以外の場合、errは [ErrClosedPipe] です。
func (w *PipeWriter) Write(data []byte) (n int, err error)

// Closeはライターを閉じます。その後のパイプの読み取り側からの読み取りは、
// バイトを返さずEOFを返します。
func (w *PipeWriter) Close() error

// CloseWithErrorはライターを閉じます。その後のパイプの読み取り側からの読み取りは、
// バイトを返さずエラーerrを返します。errがnilの場合はEOFを返します。
//
// CloseWithErrorは、以前のエラーが存在する場合はそれを上書きせず、
// 常にnilを返します。
func (w *PipeWriter) CloseWithError(err error) error

// Pipeは同期的なインメモリパイプを作成します。
// これは、[io.Reader] を期待するコードと [io.Writer] を期待するコードを
// 接続するために使用できます。
//
// パイプ上のReadとWriteは、1つのWriteを消費するために複数のReadが必要な場合を除き、
// 1対1で対応します。
// つまり、[PipeWriter] への各Writeは、書き込まれたデータを完全に消費する
// [PipeReader] からの1つ以上のReadを満たすまでブロックします。
// データはWriteから対応するRead（または複数のRead）に直接コピーされます。
// 内部バッファリングはありません。
//
// ReadとWriteを互いに、またはCloseと並行して呼び出しても安全です。
// Readの並行呼び出しやWriteの並行呼び出しも安全です。
// 個々の呼び出しは順番に処理されます。
func Pipe() (*PipeReader, *PipeWriter)
