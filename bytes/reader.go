// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/io"
)

// Readerはio.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner、io.RuneScannerのインターフェースを実装し、
// バイトスライスから読み取ります。
// Bufferとは異なり、Readerは読み込み専用であり、シークをサポートします。
// Readerのゼロ値は空のスライスのReaderのように動作します。
type Reader struct {
	s        []byte
	i        int64
	prevRune int
}

// Lenはスライスの未読部分のバイト数を返します。
func (r *Reader) Len() int

// Size は元のバイトスライスの長さを返します。
// Size は ReadAt を通じて読み取り可能なバイト数です。
// Reset を除くすべてのメソッド呼び出しによって結果は変わりません。
func (r *Reader) Size() int64

// Readはio.Readerインターフェースを実装します。
func (r *Reader) Read(b []byte) (n int, err error)

// ReadAtはio.ReaderAtインターフェースを実装します。
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// ReadByteはio.ByteReaderインターフェースを実装します。
func (r *Reader) ReadByte() (byte, error)

// UnreadByteはio.ByteScannerインターフェースを実装する際にReadByteを補完します。
func (r *Reader) UnreadByte() error

// ReadRuneはio.RuneReaderインターフェースを実装します。
func (r *Reader) ReadRune() (ch rune, size int, err error)

// UnreadRuneはio.RuneScannerインターフェースの実装においてReadRuneを補完します。
func (r *Reader) UnreadRune() error

// Seek は io.Seeker インターフェースを実装します。
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteToはio.WriterToインターフェースを実装します。
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// ResetはReaderをbからの読み取りにリセットします。
func (r *Reader) Reset(b []byte)

// NewReaderはbから読み込む新しいReaderを返す。
func NewReader(b []byte) *Reader
