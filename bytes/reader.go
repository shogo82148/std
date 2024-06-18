// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/io"
)

// Readerはio.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner、io.RuneScannerのインターフェースを実装し、
// バイトスライスから読み取ります。
// [Buffer] とは異なり、Readerは読み込み専用であり、シークをサポートします。
// Readerのゼロ値は空のスライスのReaderのように動作します。
type Reader struct {
	s        []byte
	i        int64
	prevRune int
}

// Lenはスライスの未読部分のバイト数を返します。
func (r *Reader) Len() int

// Size は元のバイトスライスの長さを返します。
// Size は [Reader.ReadAt] を通じて読み取り可能なバイト数です。
// [Reader.Reset] を除くすべてのメソッド呼び出しによって結果は変わりません。
func (r *Reader) Size() int64

// Readは [io.Reader] インターフェースを実装します。
func (r *Reader) Read(b []byte) (n int, err error)

// ReadAtは [io.ReaderAt] インターフェースを実装します。
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// ReadByteは [io.ByteReader] インターフェースを実装します。
func (r *Reader) ReadByte() (byte, error)

// UnreadByteは [io.ByteScanner] インターフェースを実装する際に [Reader.ReadByte] を補完します。
func (r *Reader) UnreadByte() error

// ReadRuneは [io.RuneReader] インターフェースを実装します。
func (r *Reader) ReadRune() (ch rune, size int, err error)

// UnreadRuneは [io.RuneScanner] インターフェースの実装において [Reader.ReadRune] を補完します。
func (r *Reader) UnreadRune() error

// Seek は [io.Seeker] インターフェースを実装します。
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteToは [io.WriterTo] インターフェースを実装します。
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// Resetは [Reader] をbからの読み取りにリセットします。
func (r *Reader) Reset(b []byte)

// NewReaderはbから読み込む新しい [Reader] を返す。
func NewReader(b []byte) *Reader
