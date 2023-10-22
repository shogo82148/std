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
type Writer struct {
	Comma   rune
	UseCRLF bool
	w       *bufio.Writer
}

// NewWriterは、wに書き込む新しいWriterを返します。
func NewWriter(w io.Writer) *Writer

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
func (w *Writer) WriteAll(records [][]string) error
