// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzip

import (
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/io"
)

// これらの定数は [flate] パッケージからコピーされており、
// [compress/gzip] をインポートするコードが [compress/flate] もインポートする必要がないようになっています。
const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)

// Writerは [io.WriteCloser] です。
// Writerへの書き込みは圧縮されてwに書き込まれます。
type Writer struct {
	Header
	w           io.Writer
	level       int
	wroteHeader bool
	closed      bool
	buf         [10]byte
	compressor  *flate.Writer
	digest      uint32
	size        uint32
	err         error
}

// NewWriterは新しい [Writer] を返します。
// 返されたWriterに書き込まれたデータは圧縮され、wに書き込まれます。
//
// [Writer] が終了したら、呼び出し元はCloseを呼ぶ責任があります。
// 書き込みはバッファリングされ、Closeが呼ばれるまでフラッシュされない場合があります。
//
// Writer.[Header] のフィールドを設定したい呼び出し元は、
// Write、Flush、またはCloseの最初の呼び出しの前に設定する必要があります。
func NewWriter(w io.Writer) *Writer

// NewWriterLevel関数は、デフォルトの圧縮レベルを仮定する代わりに、圧縮レベルを指定して
// [NewWriter] 関数と同様の処理を行います。
//
// 圧縮レベルは、 [DefaultCompression] 、 [NoCompression] 、 [HuffmanOnly] 、または [BestSpeed] から [BestCompression] までの
// いずれかの整数値を指定できます。レベルが有効である場合、返されるエラーはnilになります。
func NewWriterLevel(w io.Writer, level int) (*Writer, error)

// Resetは [Writer] zの状態を破棄し、 [NewWriter] または [NewWriterLevel] の元の状態と同等にし、
// ただし、wに書き込むことができます。これにより、新しい [Writer] を割り当てる代わりに
// [Writer] を再利用することができます。
func (z *Writer) Reset(w io.Writer)

// Writeはpを圧縮された形式で基になる [io.Writer] に書き込みます。
// 圧縮されたバイトは、 [Writer] が閉じられるまで必ずフラッシュされるわけではありません。
func (z *Writer) Write(p []byte) (int, error)

// Flushは、保留中の圧縮データを下位のライターに書き込むために使用されます。
//
// これは、主に圧縮されたネットワークプロトコルで有用であり、リモートのリーダーがパケットを再構築するために十分なデータを持っていることを保証します。データが書き込まれるまで、Flushは戻りません。下位のライターがエラーを返した場合、Flushはそのエラーを返します。
//
// zlibライブラリの用語では、FlushはZ_SYNC_FLUSHと同等です。
func (z *Writer) Flush() error

// Closeは、書き込まれていないデータを書き込み元の [io.Writer] にフラッシュし、GZIPのフッターを書き込んで [Writer] を閉じます。
// これは、書き込み元の [io.Writer] を閉じません。
func (z *Writer) Close() error
