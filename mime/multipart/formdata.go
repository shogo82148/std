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

<<<<<<< HEAD
// ReadFormは、パートのContent-Dispositionが"form-data"であるマルチパートメッセージ全体を解析します。
// メモリには最大でmaxMemoryバイト + 10MB（非ファイルパート用に予約）を格納します。
// メモリに格納できないファイルパートは、一時ファイルとしてディスクに格納されます。
// すべての非ファイルパートをメモリに格納できない場合、ErrMessageTooLargeを返します。
func (r *Reader) ReadForm(maxMemory int64) (*Form, error)

// Formは解析されたマルチパートフォームです。
// ファイルパートはメモリまたはディスクに保存され、
// *FileHeaderのOpenメソッドを通じてアクセス可能です。
// Valueパートは文字列として保存されます。
// 両方ともフィールド名でキー化されます。
=======
// ReadForm parses an entire multipart message whose parts have
// a Content-Disposition of "form-data".
// It stores up to maxMemory bytes + 10MB (reserved for non-file parts)
// in memory. File parts which can't be stored in memory will be stored on
// disk in temporary files.
// It returns [ErrMessageTooLarge] if all non-file parts can't be stored in
// memory.
func (r *Reader) ReadForm(maxMemory int64) (*Form, error)

// Form is a parsed multipart form.
// Its File parts are stored either in memory or on disk,
// and are accessible via the [*FileHeader]'s Open method.
// Its Value parts are stored as strings.
// Both are keyed by field name.
>>>>>>> upstream/master
type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}

<<<<<<< HEAD
// RemoveAllは、Formに関連付けられた一時ファイルをすべて削除します。
=======
// RemoveAll removes any temporary files associated with a [Form].
>>>>>>> upstream/master
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

<<<<<<< HEAD
// Openは、FileHeaderに関連付けられたFileを開き、それを返します。
=======
// Open opens and returns the [FileHeader]'s associated File.
>>>>>>> upstream/master
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
