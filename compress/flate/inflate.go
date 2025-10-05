// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package flateは、RFC 1951で説明されているDEFLATE圧縮データ形式を実装します。
// [compress/gzip] および [compress/zlib] パッケージは、
// DEFLATEベースのファイル形式へのアクセスを実装します。
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

// Resetterは [NewReader] または [NewReaderDict] が返すReadCloserをリセットし、新しい基になる [Reader] に切り替えます。これにより、新しいものを割り当てる代わりにReadCloserを再利用することができます。
type Resetter interface {
	Reset(r io.Reader, dict []byte) error
}

// [NewReader] で必要な実際の読み取りインターフェース。
// 渡された [io.Reader] がReadByteも持たない場合、
// [NewReader] は独自のバッファリングを導入します。
type Reader interface {
	io.Reader
	io.ByteReader
}

// NewReaderは、rの非圧縮版を読み取るために使用できる新しいReadCloserを返します。
// rが [io.ByteReader] も実装していない場合、
// デコンプレッサーはrから必要以上のデータを読み取る可能性があります。
// リーダーは、DEFLATEストリーム内の最終ブロックに遭遇した後に
// [io.EOF] を返します。最終ブロック後の末尾データは無視されます。
//
// NewReaderによって返される [io.ReadCloser] は、 [Resetter] も実装しています。
func NewReader(r io.Reader) io.ReadCloser

// NewReaderDictは [NewReader] と同様ですが、リーダーを
// プリセット辞書で初期化します。返されたリーダーは、
// 非圧縮データストリームが指定された辞書で始まったかのように動作し、
// その辞書は既に読み取られています。NewReaderDictは通常、
// [NewWriterDict] で圧縮されたデータを読み取るために使用されます。
//
// NewReaderによって返されたReadCloserは [Resetter] も実装しています。
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
