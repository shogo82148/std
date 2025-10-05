// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// lzwパッケージ は、T. A. Welchによって述べられた「高性能データ圧縮のためのテクニック」という文書で説明されている、Lempel-Ziv-Welch圧縮データフォーマットを実装します。
//
// 特に、これはGIFおよびPDFファイル形式で使用されるLZWを実装しており、可変幅コード（最大12ビット）および最初の2つの非文字コードはクリアコードとEOFコードを意味します。
//
// TIFFファイル形式は、LZWアルゴリズムの似ているが互換性のないバージョンを使用しています。
// 実装については、[golang.org/x/image/tiff/lzw] パッケージを参照してください。
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

// ReaderはLZW形式で圧縮されたデータを読み込むために使用できる [io.Reader] です。
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

// Readは[io.Reader]を実装し、基になるリーダーから非圧縮バイトを読み取ります。
func (r *Reader) Read(b []byte) (int, error)

// Closeは [Reader] を閉じ、将来の読み込み操作に対してエラーを返します。
// サブの [io.Reader] を閉じません。
func (r *Reader) Close() error

func (r *Reader) Reset(src io.Reader, order Order, litWidth int)

// NewReaderは新しい [io.ReadCloser] を作成します。
// 返された [io.ReadCloser] からの読み取りは、rからデータを読み取って解凍します。
// rが [io.ByteReader] も実装していない場合、
// デコンプレッサーはrから必要以上のデータを読み取る可能性があります。
// 読み取り完了時にReadCloserのCloseを呼び出すのは呼び出し元の責任です。
// リテラルコードに使用するビット数litWidthは、
// [2,8]の範囲でなければならず、通常は8です。これは圧縮時に
// 使用されたlitWidthと等しくなければなりません。
//
// 返された [io.ReadCloser] の基底型は、*[Reader] であることが保証されます。
func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
