// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/shogo82148/std/io"
)

// Readerは、文字列から読み取りを行うことで、io.Reader、io.ReaderAt、io.ByteReader、io.ByteScanner、
// io.RuneReader、io.RuneScanner、io.Seeker、およびio.WriterToインターフェースを実装します。
// Readerのゼロ値は、空の文字列のReaderのように動作します。
type Reader struct {
	s        string
	i        int64
	prevRune int
}

// Lenは、文字列の未読部分のバイト数を返します。
func (r *Reader) Len() int

// Sizeは、基礎となる文字列の元の長さを返します。
// Sizeは、ReadAtを介して読み取ることができるバイト数です。
// 返される値は常に同じであり、他のメソッドの呼び出しに影響を受けません。
func (r *Reader) Size() int64

// Readは、io.Readerインターフェースを実装します。
func (r *Reader) Read(b []byte) (n int, err error)

// ReadAtは、io.ReaderAtインターフェースを実装します。
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)

// ReadByteは、io.ByteReaderインターフェースを実装します。
func (r *Reader) ReadByte() (byte, error)

// UnreadByteは、io.ByteScannerインターフェースを実装します。
func (r *Reader) UnreadByte() error

// ReadRuneは、io.RuneReaderインターフェースを実装します。
func (r *Reader) ReadRune() (ch rune, size int, err error)

// UnreadRuneは、io.RuneScannerインターフェースを実装します。
func (r *Reader) UnreadRune() error

// Seekは、io.Seekerインターフェースを実装します。
func (r *Reader) Seek(offset int64, whence int) (int64, error)

// WriteToは、io.WriterToインターフェースを実装します。
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)

// Resetは、Readerをsから読み取るようにリセットします。
func (r *Reader) Reset(s string)

// NewReaderは、sから読み取る新しいReaderを返します。
// bytes.NewBufferStringに似ていますが、より効率的で書き込み不可能です。
func NewReader(s string) *Reader
