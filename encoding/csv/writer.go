// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package csv

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// Writerは、CSVエンコーディングを使用してレコードを書き込みます。
//
<<<<<<< HEAD
// As returned by [NewWriter], a Writer writes records terminated by a
// newline and uses ',' as the field delimiter. The exported fields can be
// changed to customize the details before
// the first call to [Writer.Write] or [Writer.WriteAll].
//
// [Writer.Comma] is the field delimiter.
//
// If [Writer.UseCRLF] is true,
// the Writer ends each output line with \r\n instead of \n.
//
// The writes of individual records are buffered.
// After all data has been written, the client should call the
// [Writer.Flush] method to guarantee all data has been forwarded to
// the underlying [io.Writer].  Any errors that occurred should
// be checked by calling the [Writer.Error] method.
=======
// NewWriterによって返された場合、Writerは改行で終わるレコードを書き込み、
// フィールド区切り文字として「,」を使用します。
// 最初のWriteまたはWriteAll呼び出しの前に、エクスポートされたフィールドをカスタマイズすることができます。
//
// Commaはフィールドの区切り文字です。
//
// UseCRLFがtrueの場合、Writerは各出力行を\nではなく\r\nで終了します。
//
// 個々のレコードの書き込みはバッファリングされます。
// すべてのデータが書き込まれた後、クライアントはFlushメソッドを呼び出して、
// 基礎となるio.Writerにすべてのデータが転送されたことを保証する必要があります。
// 発生したエラーは、Errorメソッドを呼び出して確認する必要があります。
>>>>>>> release-branch.go1.21
type Writer struct {
	Comma   rune
	UseCRLF bool
	w       *bufio.Writer
}

// NewWriterは、wに書き込む新しいWriterを返します。
func NewWriter(w io.Writer) *Writer

<<<<<<< HEAD
// Write writes a single CSV record to w along with any necessary quoting.
// A record is a slice of strings with each string being one field.
// Writes are buffered, so [Writer.Flush] must eventually be called to ensure
// that the record is written to the underlying [io.Writer].
func (w *Writer) Write(record []string) error

// Flush writes any buffered data to the underlying [io.Writer].
// To check if an error occurred during Flush, call [Writer.Error].
func (w *Writer) Flush()

// Error reports any error that has occurred during
// a previous [Writer.Write] or [Writer.Flush].
func (w *Writer) Error() error

// WriteAll writes multiple CSV records to w using [Writer.Write] and
// then calls [Writer.Flush], returning any error from the Flush.
=======
// Writeは、必要に応じてクォーティングを行い、単一のCSVレコードをwに書き込みます。
// レコードは、各文字列が1つのフィールドである文字列のスライスです。
// 書き込みはバッファリングされるため、レコードが基礎となるio.Writerに書き込まれることを保証するには、
// 最終的にFlushを呼び出す必要があります。
func (w *Writer) Write(record []string) error

// Flushは、バッファリングされたデータを基礎となるio.Writerに書き込みます。
// Flush中にエラーが発生したかどうかを確認するには、Errorを呼び出します。
func (w *Writer) Flush()

// Errorは、以前のWriteまたはFlush中に発生したエラーを報告します。
func (w *Writer) Error() error

// WriteAllは、Writeを使用して複数のCSVレコードをwに書き込み、
// Flushを呼び出してからFlushからのエラーを返します。
>>>>>>> release-branch.go1.21
func (w *Writer) WriteAll(records [][]string) error
