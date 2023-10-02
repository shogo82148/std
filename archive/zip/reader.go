// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/sync"
)

var (
	ErrFormat       = errors.New("zip: not a valid zip file")
	ErrAlgorithm    = errors.New("zip: unsupported compression algorithm")
	ErrChecksum     = errors.New("zip: checksum error")
	ErrInsecurePath = errors.New("zip: insecure file path")
)

type fileListEntry struct{}

// Reader は、ZIP アーカイブからコンテンツを提供するための構造体です。
type Reader struct {
	r             io.ReaderAt
	File          []*File
	Comment       string
	decompressors map[uint16]Decompressor

	baseOffset int64

	fileListOnce sync.Once
	fileList     []fileListEntry
}

// ReadCloser は、不要になったときに閉じる必要がある Reader です。
type ReadCloser struct {
	f *os.File
	Reader
}

// File は、ZIP アーカイブ内の単一のファイルです。
// ファイル情報は、埋め込み FileHeader にあります。
// ファイルの内容は、Open を呼び出すことでアクセスできます。
type File struct {
	FileHeader
	zip          *Reader
	zipr         io.ReaderAt
	headerOffset int64
	zip64        bool
}

// OpenReader は、指定された名前の Zip ファイルを開き、ReadCloser を返します。
//
// アーカイブ内のファイルのいずれかが、[filepath.IsLocal] によって定義されるローカルでない名前
// またはバックスラッシュを含む名前を使用している場合、
// および GODEBUG 環境変数に `zipinsecurepath=0` が含まれている場合、
// OpenReader は ErrInsecurePath エラーを返すリーダーを返します。
// 将来の Go のバージョンでは、この動作がデフォルトで導入される可能性があります。
// ローカルでない名前を受け入れたいプログラムは、ErrInsecurePath エラーを無視して返されたリーダーを使用できます。
func OpenReader(name string) (*ReadCloser, error)

// NewReader は、指定されたサイズを持つと想定される r から読み取る新しい Reader を返します。
//
// アーカイブ内のファイルのいずれかが、[filepath.IsLocal] によって定義されるローカルでない名前
// またはバックスラッシュを含む名前を使用している場合、
// および GODEBUG 環境変数に `zipinsecurepath=0` が含まれている場合、
// NewReader は ErrInsecurePath エラーを返すリーダーを返します。
// 将来の Go のバージョンでは、この動作がデフォルトで導入される可能性があります。
// ローカルでない名前を受け入れたいプログラムは、ErrInsecurePath エラーを無視して返されたリーダーを使用できます。
func NewReader(r io.ReaderAt, size int64) (*Reader, error)

// RegisterDecompressor は、特定のメソッド ID にカスタムの解凍プログラムを登録または上書きします。
// メソッドの解凍プログラムが見つからない場合、Reader はパッケージレベルで解凍プログラムを検索します。
func (r *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)

// Close は、Zip ファイルを閉じ、I/O に使用できなくします。
func (rc *ReadCloser) Close() error

// DataOffset は、ファイルの圧縮された可能性のあるデータのオフセットを、zip ファイルの先頭からの相対位置で返します。
//
// ほとんどの呼び出し元は、データを透過的に解凍し、チェックサムを検証する Open を代わりに使用する必要があります。
func (f *File) DataOffset() (offset int64, err error)

// Open は、ファイルの内容にアクセスする ReadCloser を返します。
// 複数のファイルを同時に読み取ることができます。
func (f *File) Open() (io.ReadCloser, error)

// OpenRaw returns a Reader that provides access to the File's contents without
// decompression.
func (f *File) OpenRaw() (io.Reader, error)

// A fileListEntry is a File and its ename.
// If file == nil, the fileListEntry describes a directory without metadata.

// Open は、fs.FS.Open のセマンティクスを使用して、ZIP アーカイブ内の指定された名前のファイルを開きます。
// パスは常にスラッシュで区切られ、先頭に / または ../ 要素はありません。
func (r *Reader) Open(name string) (fs.File, error)
