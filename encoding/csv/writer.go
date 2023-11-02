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
// [NewWriter] によって返された場合、Writerは改行で終わるレコードを書き込み、
// フィールド区切り文字として「,」を使用します。
// 最初の [Writer.Write] または [Writer.WriteAll] 呼び出しの前に、エクスポートされたフィールドをカスタマイズすることができます。
//
// [Writer.Comma] はフィールドの区切り文字です。
//
// [Writer.UseCRLF] がtrueの場合、Writerは各出力行を\nではなく\r\nで終了します。
//
// 個々のレコードの書き込みはバッファリングされます。
// すべてのデータが書き込まれた後、クライアントは [Writer.Flush] メソッドを呼び出して、
// 基礎となる [io.Writer] にすべてのデータが転送されたことを保証する必要があります。
// 発生したエラーは、[Writer.Error] メソッドを呼び出して確認する必要があります。
type Writer struct {
	Comma   rune
	UseCRLF bool
	w       *bufio.Writer
}

// NewWriterは、wに書き込む新しいWriterを返します。
func NewWriter(w io.Writer) *Writer

// Writeは、必要に応じてクォーティングを行い、単一のCSVレコードをwに書き込みます。
// レコードは、各文字列が1つのフィールドである文字列のスライスです。
// 書き込みはバッファリングされるため、レコードが基礎となる [io.Writer] に書き込まれることを保証するには、
// 最終的に [Writer.Flush] を呼び出す必要があります。
func (w *Writer) Write(record []string) error

// Flushは、バッファリングされたデータを基礎となる [io.Writer] に書き込みます。
// Flush中にエラーが発生したかどうかを確認するには、[Writer.Error] を呼び出します。
func (w *Writer) Flush()

// Errorは、以前の [Writer.Write] または [Writer.Flush] 中に発生したエラーを報告します。
func (w *Writer) Error() error

// WriteAllは、[Writer.Write] を使用して複数のCSVレコードをwに書き込み、
// [Writer.Flush] を呼び出してからFlushからのエラーを返します。
func (w *Writer) WriteAll(records [][]string) error
