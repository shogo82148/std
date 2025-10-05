// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zlib

import (
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
)

// これらの定数は [flate] パッケージからコピーされており、
// [compress/zlib] をインポートするコードが [compress/flate] もインポートする必要がないようになっています。
const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)

// Writerは書き込まれたデータを受け取り、そのデータの圧縮された形式を
// 基になるライターに書き込みます（[NewWriter] を参照）。
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

// NewWriterは新しい [Writer] を作成します。
// 返されたWriterへの書き込みは圧縮されてwに書き込まれます。
//
// Writerを使用し終わったら、呼び出し元がCloseを呼び出す責任があります。
// 書き込みはバッファリングされ、Closeが呼び出されるまでフラッシュされない場合があります。
func NewWriter(w io.Writer) *Writer

// NewWriterLevelは [NewWriter] と同様ですが、[DefaultCompression] を仮定する代わりに
// 圧縮レベルを指定します。
//
// 圧縮レベルは、[DefaultCompression]、[NoCompression]、[HuffmanOnly]
// または [BestSpeed] から [BestCompression] までの整数値のいずれかを指定できます。
// レベルが有効である場合、返されるエラーはnilになります。
func NewWriterLevel(w io.Writer, level int) (*Writer, error)

// NewWriterLevelDictは [NewWriterLevel] と同様ですが、圧縮に使用する辞書を
// 指定します。
//
// 辞書はnilである場合があります。そうでない場合、その内容はWriterが閉じられるまで変更されないようにする必要があります。
func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)

// Resetは[Writer] zの状態をクリアし、[NewWriterLevel] または [NewWriterLevelDict] からの
// 初期状態と同等にしますが、代わりにwに書き込みます。
func (z *Writer) Reset(w io.Writer)

// Writeはpの圧縮された形式を基になる [io.Writer] に書き込みます。
// 圧縮されたバイトは、[Writer] が閉じられるか
// 明示的にフラッシュされるまで必ずしもフラッシュされません。
func (z *Writer) Write(p []byte) (n int, err error)

// Flushは、Writerをその基になる [io.Writer] にフラッシュします。
func (z *Writer) Flush() error

// Closeは、Writerを閉じ、書き込まれていないデータを基になる [io.Writer] にフラッシュしますが、
// 基になる [io.Writer] を閉じません。
func (z *Writer) Close() error
