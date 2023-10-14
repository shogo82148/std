// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes

import (
	"github.com/shogo82148/std/io"
)

<<<<<<< HEAD
// Readerはio.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner、io.RuneScannerのインターフェースを実装し、
// バイトスライスから読み取ります。
// Bufferとは異なり、Readerは読み込み専用であり、シークをサポートします。
// Readerのゼロ値は空のスライスのReaderのように動作します。
=======
// A Reader implements the io.Reader, io.ReaderAt, io.WriterTo, io.Seeker,
// io.ByteScanner, and io.RuneScanner interfaces by reading from
// a byte slice.
// Unlike a [Buffer], a Reader is read-only and supports seeking.
// The zero value for Reader operates like a Reader of an empty slice.
>>>>>>> upstream/master
type Reader struct {
	s        []byte
	i        int64
	prevRune int
}

// Lenはスライスの未読部分のバイト数を返します。
func (r *Reader) Len() int

<<<<<<< HEAD
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
=======
// Size returns the original length of the underlying byte slice.
// Size is the number of bytes available for reading via [Reader.ReadAt].
// The result is unaffected by any method calls except [Reader.Reset].
func (r *Reader) Size() int64

// Read implements the [io.Reader] interface.
func (r *Reader) Read(b []byte) (n int, err error)

// ReadAt implements the [io.ReaderAt] interface.
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// ReadByte implements the [io.ByteReader] interface.
func (r *Reader) ReadByte() (byte, error)

// UnreadByte complements [Reader.ReadByte] in implementing the [io.ByteScanner] interface.
func (r *Reader) UnreadByte() error

// ReadRune implements the [io.RuneReader] interface.
func (r *Reader) ReadRune() (ch rune, size int, err error)

// UnreadRune complements [Reader.ReadRune] in implementing the [io.RuneScanner] interface.
func (r *Reader) UnreadRune() error

// Seek implements the [io.Seeker] interface.
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteTo implements the [io.WriterTo] interface.
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// Reset resets the [Reader.Reader] to be reading from b.
func (r *Reader) Reset(b []byte)

// NewReader returns a new [Reader.Reader] reading from b.
>>>>>>> upstream/master
func NewReader(b []byte) *Reader
