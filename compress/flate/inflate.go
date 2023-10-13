// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package flateは、DEFLATE圧縮データ形式を実装しています。RFC 1951で説明されています。
// gzipとzlibパッケージは、DEFLATEベースのファイル形式へのアクセスを実装しています。
package flate

import (
	"github.com/shogo82148/std/io"
)

// CorruptInputError は指定されたオフセットで破損した入力の存在を報告します。
type CorruptInputError int64

func (e CorruptInputError) Error() string

// InternalErrorはflateコード自体のエラーを報告します。
type InternalError string

func (e InternalError) Error() string

// ReadErrorは、入力を読み取る中で遭遇したエラーを報告します。
//
// Deprecated: もはや返されません。
type ReadError struct {
	Offset int64
	Err    error
}

func (e *ReadError) Error() string

// WriteErrorは出力の書き込み中に遭遇したエラーを報告します。
//
// Deprecated: もう返されません。
type WriteError struct {
	Offset int64
	Err    error
}

func (e *WriteError) Error() string

<<<<<<< HEAD
// ResetterはNewReaderまたはNewReaderDictが返すReadCloserをリセットし、新しい基になるReaderに切り替えます。これにより、新しいものを割り当てる代わりにReadCloserを再利用することができます。
=======
// Resetter resets a ReadCloser returned by [NewReader] or [NewReaderDict]
// to switch to a new underlying [Reader]. This permits reusing a ReadCloser
// instead of allocating a new one.
>>>>>>> upstream/master
type Resetter interface {
	Reset(r io.Reader, dict []byte) error
}

<<<<<<< HEAD
// NewReader で必要とされる実際の読み取りインターフェース。
// 渡された io.Reader が ReadByte も持っていない場合、
// NewReader は自身のバッファリングを導入します。
=======
// The actual read interface needed by [NewReader].
// If the passed in io.Reader does not also have ReadByte,
// the [NewReader] will introduce its own buffering.
>>>>>>> upstream/master
type Reader interface {
	io.Reader
	io.ByteReader
}

// NewReader returns a new ReadCloser that can be used
// to read the uncompressed version of r.
// If r does not also implement [io.ByteReader],
// the decompressor may read more data than necessary from r.
// The reader returns [io.EOF] after the final block in the DEFLATE stream has
// been encountered. Any trailing data after the final block is ignored.
//
<<<<<<< HEAD
// NewReaderによって返されるReadCloserは、Resetterも実装しています。
func NewReader(r io.Reader) io.ReadCloser

// NewReaderDictはNewReaderと同じようにリーダーを初期化しますが、
// 事前に設定された辞書でリーダーを初期化します。
// 返されたリーダーは、与えられた辞書で圧縮解除されたデータストリームが開始されたかのように振る舞います。
// この辞書は既に読み取られています。通常、NewWriterDictで圧縮されたデータを読み込むためにNewReaderDictが使用されます。
//
// NewReaderによって返されたReadCloserはResetterも実装しています。
=======
// The ReadCloser returned by NewReader also implements [Resetter].
func NewReader(r io.Reader) io.ReadCloser

// NewReaderDict is like [NewReader] but initializes the reader
// with a preset dictionary. The returned [Reader] behaves as if
// the uncompressed data stream started with the given dictionary,
// which has already been read. NewReaderDict is typically used
// to read data compressed by NewWriterDict.
//
// The ReadCloser returned by NewReaderDict also implements [Resetter].
>>>>>>> upstream/master
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
