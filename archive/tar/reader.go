// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tar

import (
	"github.com/shogo82148/std/io"
)

type fileReader struct{}
type block struct{}

// Readerはtarアーカイブの内容に順次アクセスするためのものです。
// Reader.Next はアーカイブ内の次のファイル（最初のファイルを含む）に進み、
// その後、Readerはファイルのデータにアクセスするためのio.Readerとして扱うことができます。
type Reader struct {
	r    io.Reader
	pad  int64
	curr fileReader
	blk  block

	// err is a persistent error.
	// It is only the responsibility of every exported method of Reader to
	// ensure that this error is sticky.
	err error
}

<<<<<<< HEAD
// NewReaderはrから読み取りを行う新しいReaderを作成します。
=======
// NewReader creates a new [Reader] reading from r.
>>>>>>> upstream/master
func NewReader(r io.Reader) *Reader

// Nextはtarアーカイブ内の次のエントリに進みます。
// Header.Sizeは次のファイルの読み取り可能なバイト数を決定します。
// 現在のファイルに残っているデータは自動的に破棄されます。
// アーカイブの末尾に達した場合、Nextはエラーio.EOFを返します。
//
<<<<<<< HEAD
// Nextが [filepath.IsLocal] によって定義されるローカルでない名前に遭遇し、
// GODEBUG環境変数に`tarinsecurepath=0`が含まれている場合、
// NextはErrInsecurePathエラーを持つヘッダを返します。
// Goの将来のバージョンでは、この動作がデフォルトで導入される可能性があります。
// ローカルでない名前を受け入れたいプログラムは、
// ErrInsecurePathエラーを無視して返されたヘッダを使用できます。
func (tr *Reader) Next() (*Header, error)

// Readはtarアーカイブ内の現在のファイルから読み取りを行います。
// 次のファイルに進むためにNextが呼び出されるまで、
// ファイルの終わりに達すると(0、io.EOF)を返します。
=======
// If Next encounters a non-local name (as defined by [filepath.IsLocal])
// and the GODEBUG environment variable contains `tarinsecurepath=0`,
// Next returns the header with an [ErrInsecurePath] error.
// A future version of Go may introduce this behavior by default.
// Programs that want to accept non-local names can ignore
// the [ErrInsecurePath] error and use the returned header.
func (tr *Reader) Next() (*Header, error)

// Read reads from the current file in the tar archive.
// It returns (0, io.EOF) when it reaches the end of that file,
// until [Next] is called to advance to the next file.
>>>>>>> upstream/master
//
// 現在のファイルがスパースである場合、
// 穴としてマークされた領域はNULバイトとして読み戻されます。
//
<<<<<<< HEAD
// TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir、TypeFifoなどの特殊なタイプでReadを呼び出すと、
// Header.Sizeが示す内容に関係なく、(0、io.EOF)が返されます。
=======
// Calling Read on special types like [TypeLink], [TypeSymlink], [TypeChar],
// [TypeBlock], [TypeDir], and [TypeFifo] returns (0, [io.EOF]) regardless of what
// the [Header.Size] claims.
>>>>>>> upstream/master
func (tr *Reader) Read(b []byte) (int, error)
