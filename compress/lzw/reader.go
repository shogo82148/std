// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージ lzw は、T. A. Welchによって述べられた「高性能データ圧縮のためのテクニック」という文書で説明されている、Lempel-Ziv-Welch圧縮データフォーマットを実装します。
//
// 特に、これはGIFおよびPDFファイル形式で使用されるLZWを実装しており、可変幅コード（最大12ビット）および最初の2つの非文字コードはクリアコードとEOFコードを意味します。
//
// TIFFファイル形式は、LZWアルゴリズムの似ているが互換性のないバージョンを使用しています。実装については、golang.org/x/image/tiff/lzwパッケージを参照してください。
package lzw

import (
	"github.com/shogo82148/std/io"
)

// OrderはLZWデータストリーム内のビットの順序を指定します。
type Order int

const (
	// LSBは、GIFファイルフォーマットで使用されるLeast Significant Bits（最下位ビット優先）の意味です。
	LSB Order = iota

	// MSBは、TIFFおよびPDFファイル形式で使用される、最上位ビットを優先することを意味します。
	MSB
)

// ReaderはLZW形式で圧縮されたデータを読み込むために使用できるio.Readerです。
type Reader struct {
	r        io.ByteReader
	bits     uint32
	nBits    uint
	width    uint
	read     func(*Reader) (uint16, error)
	litWidth int
	err      error

	// 最初の1<<litWidthのコードはリテラルコードです。
	// 次の2つのコードはクリアとEOFを意味します。
	// 他の有効なコードは [lo, hi] の範囲にあります。ここで、lo := clear + 2 であり、
	// 各コードが現れるたびに上限が増加します。
	//
	// overflowはhiがコード幅を超えるコードです。常に1 << widthと等しくなります。
	//
	// lastは最後に見たコード、またはdecoderInvalidCodeです。
	//
	// 不変事項はhi < overflowです。
	clear, eof, hi, overflow, last uint16

	// [lo, hi]の各コードcは2バイト以上に展開されます。ただし、c != hiの場合：
	//   suffix[c]はこれらのバイトの最後のバイトです。
	//   prefix[c]は最後のバイト以外のコードです。
	//   このコードは、リテラルコードまたは[lo, c)内の別のコードである場合があります。
	// c == hiの場合は特別なケースです。
	suffix [1 << maxWidth]uint8
	prefix [1 << maxWidth]uint16

	// outputは一時的な出力バッファです。
	// 文字コードはバッファの先頭から蓄積されます。
	// リテラルコードは、バッファの末尾から右から左にデコードされ、
	// バッファの先頭にコピーされる接尾辞のシーケンスにデコードされます。
	// バッファに >= 1<<maxWidth バイトが含まれている場合、フラッシュされます。
	// したがって、常にコード全体をデコードするためのスペースがあります。
	output [2 * 1 << maxWidth]byte
	o      int
	toRead []byte
}

// Readはio.Readerを実装し、基になるReaderから非圧縮バイトを読み取ります。
func (r *Reader) Read(b []byte) (int, error)

// CloseはReaderを閉じ、将来の読み込み操作に対してエラーを返します。
// サブのio.Readerを閉じません。
func (r *Reader) Close() error

// ResetはReaderの状態をクリアし、新しいReaderとして再利用するために使用されます。
func (r *Reader) Reset(src io.Reader, order Order, litWidth int)

// NewReader creates a new io.ReadCloser.
// Reads from the returned io.ReadCloser read and decompress data from r.
// If r does not also implement [io.ByteReader],
// the decompressor may read more data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser when
// finished reading.
// The number of bits to use for literal codes, litWidth, must be in the
// range [2,8] and is typically 8. It must equal the litWidth
// used during compression.
//
// 返されたio.ReadCloserの基底型は、*Readerであることが保証されます。
func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
