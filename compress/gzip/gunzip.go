// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// gzipパッケージは RFC 1952 で指定されている gzip 形式の圧縮ファイルの読み書きを実装しています。
package gzip

import (
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

var (
	// ErrChecksum は、無効なチェックサムを持つGZIPデータを読み取る際に返されます。
	ErrChecksum = errors.New("gzip: invalid checksum")
	// ErrHeader は無効なヘッダーを持つ GZIP データを読み取る際に返されます。
	ErrHeader = errors.New("gzip: invalid header")
)

// gzipファイルは、圧縮ファイルに関するメタデータを示すヘッダーを格納しています。
// そのヘッダーは、 [Writer] および [Reader] の構造体のフィールドとして公開されています。
//
// 文字列はUTF-8でエンコードする必要があり、UnicodeのコードポイントU+0001からU+00FFのみを含むことができます。
// これは、GZIPファイル形式の制約によるものです。
type Header struct {
	Comment string
	Extra   []byte
	ModTime time.Time
	Name    string
	OS      byte
}

// Readerは、gzip形式の圧縮ファイルから非圧縮データを取得するために読み取り可能な [io.Reader] です。
//
// 一般的に、gzipファイルはgzipファイルの連結であり、各ファイルには独自のヘッダがあります。
// Readerから読み取ると、各非圧縮データの連結が返されます。
// Readerのフィールドには最初のヘッダのみが記録されます。
//
// Gzipファイルには非圧縮データの長さとチェックサムが格納されています。
// [Reader.Read] は、非圧縮データの末尾に到達した場合、期待された長さやチェックサムがない場合に [ErrChecksum] を返します。
// クライアントは、 [Reader.Read] によって返されるデータを受け取るまで、仮のものとして扱うべきです。
// データの終端を示す [io.EOF] を受け取るまで。
type Reader struct {
	Header
	r            flate.Reader
	decompressor io.ReadCloser
	digest       uint32
	size         uint32
	buf          [512]byte
	err          error
	multistream  bool
}

// NewReader creates a new [Reader] reading the given reader.
// If r does not also implement [io.ByteReader],
// the decompressor may read more data than necessary from r.
//
// [Reader] を使用し終わった後は、呼び出し元の責任でCloseを呼び出す必要があります。
//
// [Reader] によって返される [Reader.Header] フィールドは有効です。
func NewReader(r io.Reader) (*Reader, error)

// Resetは [Reader] zの状態を破棄し、 [NewReader] からの元の状態の結果と同等にしますが、代わりにrから読み込みます。
// これにより、新しい [Reader] を割り当てる代わりに、 [Reader] を再利用することができます。
func (z *Reader) Reset(r io.Reader) error

// Multistreamは、リーダーがマルチストリームファイルに対応しているかどうかを制御します。
//
// 有効にすると（デフォルトでは有効）、 [Reader] は入力が各々個別にgzipされたデータストリームのシーケンスであることを期待し、
// 各データストリームにはヘッダとトレーラーがあり、EOFで終了します。
// このため、gzipで連結されたシーケンスの連結とgzip化は同等と見なされます。これはgzipリーダーの標準的な動作です。
//
// Multistream(false)を呼び出すと、この動作を無効にできます。
// 動作を無効にすることは、個々のgzipデータストリームを識別するファイル形式を読み込む場合や、
// gzipデータストリームと他のデータストリームを混在させるファイル形式を読み込む場合に便利です。
// このモードでは、 [Reader] がデータストリームの終端に達した場合、 [Reader.Read] は [io.EOF] を返します。
// 基底のリーダーは [io.ByteReader] を実装している必要があり、gzipストリームの直後に位置を残しておくようになります。
// 次のストリームを開始するには、z.Reset（r）を呼び出した後にz.Multistream(false)を呼び出します。
// 次のストリームが存在しない場合、z.Reset（r）は [io.EOF] を返します。
func (z *Reader) Multistream(ok bool)

// Readは、基になるReaderから圧縮されていないバイトを読み込むために [io.Reader] を実装しています。
func (z *Reader) Read(p []byte) (n int, err error)

// CloseはReaderを閉じます。ただし、基本となる [io.Reader] は閉じません。
// GZIPのチェックサムを検証するためには、 [io.EOF] まで完全に消費する必要があります。
func (z *Reader) Close() error
