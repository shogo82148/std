// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"github.com/shogo82148/std/io"
)

type fileWriter struct{}

// Writerはtarアーカイブの順次書き込みを提供します。
// Write.WriteHeaderは提供されたHeaderで新しいファイルを開始し、
// その後、Writerはそのファイルのデータを提供するためのio.Writerとして扱うことができます。
type Writer struct {
	w    io.Writer
	pad  int64
	curr fileWriter
	hdr  Header
	blk  block

	err error
}

// NewWriterはwに書き込む新しいWriterを作成します。
func NewWriter(w io.Writer) *Writer

// Flushは現在のファイルのブロックパディングの書き込みを終了します。
// Flushを呼び出す前に、現在のファイルは完全に書き込まれている必要があります。
//
// これは、次のWriteHeaderまたはCloseの呼び出しで
// ファイルのパディングが暗黙的にフラッシュされるため、不要です。
func (tw *Writer) Flush() error

// WriteHeaderはhdrを書き込み、ファイルの内容を受け入れる準備をします。
// Header.Sizeは、次のファイルの書き込み可能なバイト数を決定します。
// 現在のファイルが完全に書き込まれていない場合、エラーが返されます。
// これにより、ヘッダを書き込す前に必要なパディングが暗黙的にフラッシュされます。
func (tw *Writer) WriteHeader(hdr *Header) error

// Writeはtarアーカイブ内の現在のファイルに書き込みます。
// WriteHeaderの後にHeader.Sizeバイト以上が書き込まれた場合、
// WriteはエラーErrWriteTooLongを返します。
//
// TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir、TypeFifoなどの特殊なタイプでWriteを呼び出すと、
// Header.Sizeが示す内容に関係なく、(0、ErrWriteTooLong)が返されます。
func (tw *Writer) Write(b []byte) (int, error)

// Closeはパディングをフラッシュし、フッターを書き込むことでtarアーカイブを閉じます。
// (WriteHeaderの前の)現在のファイルが完全に書き込まれていない場合、エラーが返されます。
func (tw *Writer) Close() error

// regFileWriter is a fileWriter for writing data to a regular file entry.

// sparseFileWriter is a fileWriter for writing data to a sparse file entry.

// zeroWriter may only be written with NULs, otherwise it returns errWriteHole.
