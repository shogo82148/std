// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
zlibパッケージは、RFC 1950で指定されているように、zlib形式の圧縮データの読み書きを実装します。

この実装は、読み取り中に解凍するフィルターと、書き込み中に圧縮するフィルターを提供します。
たとえば、圧縮されたデータをバッファーに書き込むには：

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()

そして、そのデータを読み戻すには：

	r, err := zlib.NewReader(&b)
	io.Copy(os.Stdout, r)
	r.Close()
*/
package zlib

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

var (
	// ErrChecksumは、無効なチェックサムを持つZLIBデータを読み取る場合に返されます。
	ErrChecksum = errors.New("zlib: invalid checksum")
	// ErrDictionaryは、無効な辞書を持つZLIBデータを読み取る場合に返されます。
	ErrDictionary = errors.New("zlib: invalid dictionary")
	// ErrHeaderは、無効なヘッダーを持つZLIBデータを読み取る場合に返されます。
	ErrHeader = errors.New("zlib: invalid header")
)

// Resetterは、[NewReader] または [NewReaderDict] によって返されたReadCloserをリセットし、
// 新しい基礎となるReaderに切り替えます。これにより、新しいReadCloserを割り当てる代わりに
// 再利用できます。
type Resetter interface {
	Reset(r io.Reader, dict []byte) error
}

// NewReaderは新しいReadCloserを作成します。
// 返されたReadCloserからの読み取りは、rからデータを読み取って解凍します。
// rが [io.ByteReader] を実装していない場合、解凍器はrから必要以上のデータを
// 読み取る可能性があります。
// 終了時にReadCloserのCloseを呼び出すのは呼び出し元の責任です。
//
// NewReaderによって返される [io.ReadCloser] は [Resetter] も実装しています。
func NewReader(r io.Reader) (io.ReadCloser, error)

// NewReaderDictは [NewReader] と同様ですが、プリセット辞書を使用します。
// 圧縮データが辞書を参照していない場合、NewReaderDictは辞書を無視します。
// 圧縮データが別の辞書を参照している場合、NewReaderDictは [ErrDictionary] を返します。
//
// NewReaderDictによって返されるReadCloserも [Resetter] を実装しています。
func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)
