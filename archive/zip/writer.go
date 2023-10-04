// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

type (
	countWriter struct{}
	header      struct{}
	fileWriter  struct{}
)

// Writer は、ZIP ファイルのライターを実装します。
type Writer struct {
	cw          *countWriter
	dir         []*header
	last        *fileWriter
	closed      bool
	compressors map[uint16]Compressor
	comment     string

	// testHookCloseSizeOffset if non-nil is called with the size
	// of offset of the central directory at Close.
	testHookCloseSizeOffset func(size, offset uint64)
}

// NewWriter は、w に ZIP ファイルを書き込む新しい Writer を返します。
func NewWriter(w io.Writer) *Writer

// SetOffset は、zip データの開始オフセットを基になるライター内に設定します。
// これは、バイナリ実行可能ファイルなどに zip データが追加される場合に使用する必要があります。
// データが書き込まれる前に呼び出す必要があります。
func (w *Writer) SetOffset(n int64)

// Flush は、バッファリングされたデータを基になるライターにフラッシュします。
// Flush を呼び出す必要は通常ありません。Close を呼び出すだけで十分です。
func (w *Writer) Flush() error

// SetComment は、中央ディレクトリのコメントフィールドを設定します。
// Close を呼び出す前にのみ呼び出すことができます。
func (w *Writer) SetComment(comment string) error

// Close は、中央ディレクトリを書き込むことで zip ファイルの書き込みを終了します。
// 基になるライターを閉じません。
func (w *Writer) Close() error

// Create は、指定された名前を使用してファイルを zip ファイルに追加します。
// 返される Writer にファイルの内容を書き込む必要があります。
// ファイルの内容は、Deflate メソッドを使用して圧縮されます。
// 名前は相対パスである必要があります。
// ドライブレター（例：C：）または先頭のスラッシュで始まることはできず、
// スラッシュのみが許可されます。
// ファイルではなくディレクトリを作成するには、名前の末尾にスラッシュを追加します。
// 次の Create、CreateHeader、または Close を呼び出す前に、ファイルの内容を io.Writer に書き込む必要があります。
func (w *Writer) Create(name string) (io.Writer, error)

// CreateHeader は、ファイルメタデータに提供された FileHeader を使用して、zip アーカイブにファイルを追加します。
// Writer は fh を所有し、そのフィールドを変更する可能性があります。
// CreateHeader を呼び出した後、呼び出し元は fh を変更してはいけません。
//
// これは、ファイルの内容を書き込む必要がある Writer を返します。
// 次の Create、CreateHeader、CreateRaw、または Close を呼び出す前に、ファイルの内容を io.Writer に書き込む必要があります。
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)

// CreateRaw は、提供された FileHeader を使用して zip アーカイブにファイルを追加し、
// ファイルの内容を書き込むための Writer を返します。
// 次の Create、CreateHeader、CreateRaw、または Close を呼び出す前に、ファイルの内容を io.Writer に書き込む必要があります。
//
// CreateHeader とは異なり、Writer に渡されるバイトは圧縮されません。
func (w *Writer) CreateRaw(fh *FileHeader) (io.Writer, error)

// Copy は、ファイル f（Reader から取得された）を w にコピーします。
// これは、解凍、圧縮、および検証をバイパスして、生の形式で直接コピーします。
func (w *Writer) Copy(f *File) error

// RegisterCompressor は、特定のメソッド ID にカスタムの圧縮プログラムを登録または上書きします。
// メソッドの圧縮プログラムが見つからない場合、Writer はパッケージレベルで圧縮プログラムを検索します。
func (w *Writer) RegisterCompressor(method uint16, comp Compressor)

// AddFS adds the files from fs.FS to the archive.
// It walks the directory tree starting at the root of the filesystem
// adding each file to the zip using deflate while maintaining the directory structure.
func (w *Writer) AddFS(fsys fs.FS) error
