// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzip

import (
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/io"
)

// これらの定数はflateパッケージからコピーされています。そのため、「compress/gzip」をインポートするコードは、「compress/flate」もインポートする必要はありません。
const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)

// Writerはio.WriteCloserです。
// Writerへの書き込みは圧縮され、wに書き込まれます。
type Writer struct {
	Header
	w           io.Writer
	level       int
	wroteHeader bool
	compressor  *flate.Writer
	digest      uint32
	size        uint32
	closed      bool
	buf         [10]byte
	err         error
}

<<<<<<< HEAD
// NewWriterは新しいWriterを返します。
// 返されたWriterに書き込まれたデータは圧縮され、wに書き込まれます。
//
// Writerが終了したら、呼び出し元はCloseを呼ぶ責任があります。
// 書き込みはバッファリングされ、Closeが呼ばれるまでフラッシュされない場合があります。
=======
// NewWriter returns a new [Writer].
// Writes to the returned writer are compressed and written to w.
//
// It is the caller's responsibility to call Close on the [Writer] when done.
// Writes may be buffered and not flushed until Close.
>>>>>>> upstream/master
//
// Writer.Headerのフィールドを設定したい呼び出し元は、
// Write、Flush、またはCloseの最初の呼び出しの前に設定する必要があります。
func NewWriter(w io.Writer) *Writer

<<<<<<< HEAD
// NewWriterLevel関数は、デフォルトの圧縮レベルを仮定する代わりに、圧縮レベルを指定して
// NewWriter関数と同様の処理を行います。
//
// 圧縮レベルは、DefaultCompression、NoCompression、HuffmanOnly、またはBestSpeedからBestCompressionまでの
// いずれかの整数値を指定できます。レベルが有効である場合、返されるエラーはnilになります。
func NewWriterLevel(w io.Writer, level int) (*Writer, error)

// ResetはWriter zの状態を破棄し、NewWriterまたはNewWriterLevelの元の状態と同等にし、
// ただし、wに書き込むことができます。これにより、新しいWriterを割り当てる代わりに
// Writerを再利用することができます。
func (z *Writer) Reset(w io.Writer)

// Writeはpを圧縮された形式で基になるio.Writerに書き込みます。
// 圧縮されたバイトは、Writerが閉じられるまで必ずフラッシュされるわけではありません。
=======
// NewWriterLevel is like [NewWriter] but specifies the compression level instead
// of assuming [DefaultCompression].
//
// The compression level can be [DefaultCompression], [NoCompression], [HuffmanOnly]
// or any integer value between [BestSpeed] and [BestCompression] inclusive.
// The error returned will be nil if the level is valid.
func NewWriterLevel(w io.Writer, level int) (*Writer, error)

// Reset discards the [Writer] z's state and makes it equivalent to the
// result of its original state from [NewWriter] or [NewWriterLevel], but
// writing to w instead. This permits reusing a [Writer] rather than
// allocating a new one.
func (z *Writer) Reset(w io.Writer)

// Write writes a compressed form of p to the underlying [io.Writer]. The
// compressed bytes are not necessarily flushed until the [Writer] is closed.
>>>>>>> upstream/master
func (z *Writer) Write(p []byte) (int, error)

// Flushは、保留中の圧縮データを下位のライターに書き込むために使用されます。
//
// これは、主に圧縮されたネットワークプロトコルで有用であり、リモートのリーダーがパケットを再構築するために十分なデータを持っていることを保証します。データが書き込まれるまで、Flushは戻りません。下位のライターがエラーを返した場合、Flushはそのエラーを返します。
//
// zlibライブラリの用語では、FlushはZ_SYNC_FLUSHと同等です。
func (z *Writer) Flush() error

<<<<<<< HEAD
// Closeは、書き込まれていないデータを書き込み元のio.Writerにフラッシュし、GZIPのフッターを書き込んでWriterを閉じます。
// これは、書き込み元のio.Writerを閉じません。
=======
// Close closes the [Writer] by flushing any unwritten data to the underlying
// [io.Writer] and writing the GZIP footer.
// It does not close the underlying [io.Writer].
>>>>>>> upstream/master
func (z *Writer) Close() error
