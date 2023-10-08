// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zlib

import (
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
)

// これらの定数はflateパッケージからコピーされています。
// これにより、「compress/zlib」をインポートするコードが「compress/flate」もインポートする必要がなくなります。
const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)

// Writerは、書き込まれたデータを受け取り、そのデータの圧縮形式を下位のライターに書き込みます（NewWriterを参照）。
type Writer struct {
	w           io.Writer
	level       int
	dict        []byte
	compressor  *flate.Writer
	digest      hash.Hash32
	err         error
	scratch     [4]byte
	wroteHeader bool
}

// NewWriterは、新しいWriterを作成します。
// 返されたWriterに書き込まれたデータは圧縮され、wに書き込まれます。
//
// Writerを使用し終わったら、呼び出し元がCloseを呼び出す責任があります。
// 書き込みはバッファリングされ、Closeが呼び出されるまでフラッシュされない場合があります。
func NewWriter(w io.Writer) *Writer

// NewWriterLevelは、NewWriterと同様ですが、デフォルトの圧縮レベルを仮定する代わりに、
// 圧縮レベルを指定します。
//
// 圧縮レベルは、DefaultCompression、NoCompression、HuffmanOnly、BestSpeedからBestCompressionまでの
// 任意の整数値であることができます。レベルが有効である場合、返されるエラーはnilになります。
func NewWriterLevel(w io.Writer, level int) (*Writer, error)

// NewWriterLevelDictは、NewWriterLevelと同様ですが、圧縮に使用する辞書を指定します。
//
// 辞書はnilである場合があります。そうでない場合、その内容はWriterが閉じられるまで変更されないようにする必要があります。
func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)

// Resetは、Writer zの状態をクリアし、NewWriterLevelまたはNewWriterLevelDictからの初期状態と同等になるようにしますが、
// 代わりにwに書き込みます。
func (z *Writer) Reset(w io.Writer)

// Writeは、pの圧縮形式を基になるio.Writerに書き込みます。
// 圧縮されたバイトは、Writerが閉じられるか、明示的にフラッシュされるまで必ずしもフラッシュされません。
func (z *Writer) Write(p []byte) (n int, err error)

// Flushは、Writerをその基になるio.Writerにフラッシュします。
func (z *Writer) Flush() error

// Closeは、Writerを閉じ、書き込まれていないデータを基になるio.Writerにフラッシュしますが、
// 基になるio.Writerを閉じません。
func (z *Writer) Close() error
