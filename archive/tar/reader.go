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

// NewReaderはrから読み取りを行う新しい [Reader] を作成します。
func NewReader(r io.Reader) *Reader

// Nextはtarアーカイブ内の次のエントリに進みます。
// Header.Sizeは次のファイルの読み取り可能なバイト数を決定します。
// 現在のファイルに残っているデータは自動的に破棄されます。
// アーカイブの末尾に達した場合、Nextはエラーio.EOFを返します。
//
// Nextが [filepath.IsLocal] によって定義されるローカルでない名前に遭遇し、
// GODEBUG環境変数に`tarinsecurepath=0`が含まれている場合、
// Nextは [ErrInsecurePath] エラーを伴うヘッダーを返します。
// 将来のGoのバージョンでは、この動作がデフォルトで導入される可能性があります。
// ローカルでない名前を受け入れたいプログラムは、 [ErrInsecurePath] エラーを無視して返されたヘッダーを使用できます。
func (tr *Reader) Next() (*Header, error)

// Read reads from the current file in the tar archive.
// It returns (0, io.EOF) when it reaches the end of that file,
// until [Next] is called to advance to the next file.
//
// 現在のファイルがスパースである場合、
// 穴としてマークされた領域はNULバイトとして読み戻されます。
//
// [TypeLink] 、 [TypeSymlink] 、 [TypeChar] 、 [TypeBlock] 、 [TypeDir] 、 [TypeFifo] などの特殊なタイプでReadを呼び出すと、
// [Header.Size] が示す内容に関係なく、(0, [io.EOF]) が返されます。
func (tr *Reader) Read(b []byte) (int, error)
