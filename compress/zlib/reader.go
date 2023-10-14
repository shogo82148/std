// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージzlibは、RFC 1950で指定されているように、zlib形式の圧縮データの読み書きを実装します。

この実装は、読み取り中に解凍するフィルターと、書き込み中に圧縮するフィルターを提供します。
たとえば、圧縮されたデータをバッファに書き込むには：

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

// Resetter resets a ReadCloser returned by NewReader or NewReaderDict
// to switch to a new underlying Reader. This permits reusing a ReadCloser
// instead of allocating a new one.
type Resetter interface {
	// Reset discards any buffered data and resets the Resetter as if it was
	// newly initialized with the given reader.
	Reset(r io.Reader, dict []byte) error
}

// NewReader creates a new ReadCloser.
// Reads from the returned ReadCloser read and decompress data from r.
// If r does not implement io.ByteReader, the decompressor may read more
// data than necessary from r.
// It is the caller's responsibility to call Close on the ReadCloser when done.
//
// The ReadCloser returned by NewReader also implements Resetter.
func NewReader(r io.Reader) (io.ReadCloser, error)

// NewReaderDict is like NewReader but uses a preset dictionary.
// NewReaderDict ignores the dictionary if the compressed data does not refer to it.
// If the compressed data refers to a different dictionary, NewReaderDict returns ErrDictionary.
//
// The ReadCloser returned by NewReaderDict also implements Resetter.
func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)
