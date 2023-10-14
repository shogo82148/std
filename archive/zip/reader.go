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

	// Some JAR files are zip files with a prefix that is a bash script.
	// The baseOffset field is the start of the zip file proper.
	baseOffset int64

	// fileList is a list of files sorted by ename,
	// for use by the Open method.
	fileListOnce sync.Once
	fileList     []fileListEntry
}

<<<<<<< HEAD
// ReadCloser は、不要になったときに閉じる必要がある Reader です。
=======
// A ReadCloser is a [Reader] that must be closed when no longer needed.
>>>>>>> upstream/master
type ReadCloser struct {
	f *os.File
	Reader
}

<<<<<<< HEAD
// File は、ZIP アーカイブ内の単一のファイルです。
// ファイル情報は、埋め込み FileHeader にあります。
// ファイルの内容は、Open を呼び出すことでアクセスできます。
=======
// A File is a single file in a ZIP archive.
// The file information is in the embedded [FileHeader].
// The file content can be accessed by calling [File.Open].
>>>>>>> upstream/master
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

<<<<<<< HEAD
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
=======
// NewReader returns a new [Reader] reading from r, which is assumed to
// have the given size in bytes.
//
// If any file inside the archive uses a non-local name
// (as defined by [filepath.IsLocal]) or a name containing backslashes
// and the GODEBUG environment variable contains `zipinsecurepath=0`,
// NewReader returns the reader with an [ErrInsecurePath] error.
// A future version of Go may introduce this behavior by default.
// Programs that want to accept non-local names can ignore
// the [ErrInsecurePath] error and use the returned reader.
func NewReader(r io.ReaderAt, size int64) (*Reader, error)

// RegisterDecompressor registers or overrides a custom decompressor for a
// specific method ID. If a decompressor for a given method is not found,
// [Reader] will default to looking up the decompressor at the package level.
>>>>>>> upstream/master
func (r *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)

// Close は、Zip ファイルを閉じ、I/O に使用できなくします。
func (rc *ReadCloser) Close() error

// DataOffset は、ファイルの圧縮された可能性のあるデータのオフセットを、zip ファイルの先頭からの相対位置で返します。
//
<<<<<<< HEAD
// ほとんどの呼び出し元は、データを透過的に解凍し、チェックサムを検証する Open を代わりに使用する必要があります。
func (f *File) DataOffset() (offset int64, err error)

// Open は、ファイルの内容にアクセスする ReadCloser を返します。
// 複数のファイルを同時に読み取ることができます。
=======
// Most callers should instead use [File.Open], which transparently
// decompresses data and verifies checksums.
func (f *File) DataOffset() (offset int64, err error)

// Open returns a [ReadCloser] that provides access to the [File]'s contents.
// Multiple files may be read concurrently.
>>>>>>> upstream/master
func (f *File) Open() (io.ReadCloser, error)

// OpenRaw returns a [Reader] that provides access to the [File]'s contents without
// decompression.
func (f *File) OpenRaw() (io.Reader, error)

// Openは、fs.FS.Openのセマンティクスを使用して、ZIPアーカイブ内の指定されたファイルを開きます。
// パスは常にスラッシュで区切られ、先頭に/または../要素はありません。
func (r *Reader) Open(name string) (fs.File, error)
