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

	err error
}

// NewReaderはrから読み取りを行う新しいReaderを作成します。
func NewReader(r io.Reader) *Reader

// Nextはtarアーカイブ内の次のエントリに進みます。
// Header.Sizeは次のファイルの読み取り可能なバイト数を決定します。
// 現在のファイルに残っているデータは自動的に破棄されます。
// アーカイブの末尾に達した場合、Nextはエラーio.EOFを返します。
//
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
//
// 現在のファイルがスパースである場合、
// 穴としてマークされた領域はNULバイトとして読み戻されます。
//
// TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir、TypeFifoなどの特殊なタイプでReadを呼び出すと、
// Header.Sizeが示す内容に関係なく、(0、io.EOF)が返されます。
func (tr *Reader) Read(b []byte) (int, error)

// regFileReader is a fileReader for reading data from a regular file entry.

// sparseFileReader is a fileReader for reading data from a sparse file entry.
