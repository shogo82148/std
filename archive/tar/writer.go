// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

<<<<<<< HEAD
type fileWriter struct{}

// Writerはtarアーカイブの順次書き込みを提供します。
// Write.WriteHeaderは提供されたHeaderで新しいファイルを開始し、
// その後、Writerはそのファイルのデータを提供するためのio.Writerとして扱うことができます。
=======
// Writer provides sequential writing of a tar archive.
// [Writer.WriteHeader] begins a new file with the provided [Header],
// and then Writer can be treated as an io.Writer to supply that file's data.
>>>>>>> upstream/master
type Writer struct {
	w    io.Writer
	pad  int64
	curr fileWriter
	hdr  Header
	blk  block

	// err is a persistent error.
	// It is only the responsibility of every exported method of Writer to
	// ensure that this error is sticky.
	err error
}

// NewWriterはwに書き込む新しいWriterを作成します。
func NewWriter(w io.Writer) *Writer

// Flushは現在のファイルのブロックパディングの書き込みを終了します。
// Flushを呼び出す前に、現在のファイルは完全に書き込まれている必要があります。
//
<<<<<<< HEAD
// これは、次のWriteHeaderまたはCloseの呼び出しで
// ファイルのパディングが暗黙的にフラッシュされるため、不要です。
=======
// This is unnecessary as the next call to [Writer.WriteHeader] or [Writer.Close]
// will implicitly flush out the file's padding.
>>>>>>> upstream/master
func (tw *Writer) Flush() error

// WriteHeaderはhdrを書き込み、ファイルの内容を受け入れる準備をします。
// Header.Sizeは、次のファイルの書き込み可能なバイト数を決定します。
// 現在のファイルが完全に書き込まれていない場合、エラーが返されます。
// これにより、ヘッダを書き込す前に必要なパディングが暗黙的にフラッシュされます。
func (tw *Writer) WriteHeader(hdr *Header) error

// AddFSは、fs.FSからファイルをアーカイブに追加します。
// ファイルシステムのルートから開始してディレクトリツリーを走査し、
// 各ファイルをtarアーカイブに追加しながらディレクトリ構造を維持します。
func (tw *Writer) AddFS(fsys fs.FS) error

<<<<<<< HEAD
// Writeは、tarアーカイブの現在のファイルに書き込みます。
// WriteHeaderの後にHeader.Sizeバイト以上が書き込まれた場合、WriteはErrWriteTooLongエラーを返します。
//
// TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir、TypeFifoなどの特殊なタイプでWriteを呼び出すと、
// Header.Sizeが示す内容に関係なく、(0、ErrWriteTooLong)が返されます。
func (tw *Writer) Write(b []byte) (int, error)

// Closeはパディングをフラッシュし、フッターを書き込むことでtarアーカイブを閉じます。
// (WriteHeaderの前の)現在のファイルが完全に書き込まれていない場合、エラーが返されます。
=======
// Write writes to the current file in the tar archive.
// Write returns the error [ErrWriteTooLong] if more than
// Header.Size bytes are written after [Writer.WriteHeader].
//
// Calling Write on special types like [TypeLink], [TypeSymlink], [TypeChar],
// [TypeBlock], [TypeDir], and [TypeFifo] returns (0, [ErrWriteTooLong]) regardless
// of what the [Header.Size] claims.
func (tw *Writer) Write(b []byte) (int, error)

// Close closes the tar archive by flushing the padding, and writing the footer.
// If the current file (from a prior call to [Writer.WriteHeader]) is not fully written,
// then this returns an error.
>>>>>>> upstream/master
func (tw *Writer) Close() error
