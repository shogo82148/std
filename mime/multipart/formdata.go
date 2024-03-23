// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multipart

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/textproto"
)

// ErrMessageTooLargeは、メッセージフォームデータが処理するには大きすぎる場合、ReadFormによって返されます。
var ErrMessageTooLarge = errors.New("multipart: message too large")

// ReadFormは、パートのContent-Dispositionが"form-data"であるマルチパートメッセージ全体を解析します。
// メモリには最大でmaxMemoryバイト + 10MB（非ファイルパート用に予約）を格納します。
// メモリに格納できないファイルパートは、一時ファイルとしてディスクに格納されます。
// すべての非ファイルパートをメモリに格納できない場合、[ErrMessageTooLarge] を返します。
func (r *Reader) ReadForm(maxMemory int64) (*Form, error)

// Formは解析されたマルチパートフォームです。
// ファイルパートはメモリまたはディスクに保存され、
// [*FileHeader] のOpenメソッドを通じてアクセス可能です。
// Valueパートは文字列として保存されます。
// 両方ともフィールド名でキー化されます。
type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}

// RemoveAllは、[Form] に関連付けられた一時ファイルをすべて削除します。
func (f *Form) RemoveAll() error

// FileHeaderは、マルチパートリクエストのファイル部分を記述します。
type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64

	content   []byte
	tmpfile   string
	tmpoff    int64
	tmpshared bool
}

// Openは、[FileHeader] に関連付けられたFileを開き、それを返します。
func (fh *FileHeader) Open() (File, error)

// Fileは、マルチパートメッセージのファイル部分にアクセスするためのインターフェースです。
// その内容はメモリに保存されるか、またはディスクに保存されます。
// ディスクに保存されている場合、Fileの基礎となる具体的な型は*os.Fileになります。
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}
