// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quotedprintable

import "github.com/shogo82148/std/io"

<<<<<<< HEAD
// Writerは、io.WriteCloserを実装するquoted-printableライターです。
=======
// A Writer is a quoted-printable writer that implements [io.WriteCloser].
>>>>>>> upstream/master
type Writer struct {
	// バイナリモードでは、ライターの入力を純粋なバイナリとして扱い、
	// 行末のバイトをバイナリデータとして処理します。
	Binary bool

	w    io.Writer
	i    int
	line [78]byte
	cr   bool
}

<<<<<<< HEAD
// NewWriterは、wに書き込む新しいWriterを返します。
func NewWriter(w io.Writer) *Writer

// Writeは、pをquoted-printableエンコーディングでエンコードし、それを
// 基礎となるio.Writerに書き込みます。行の長さは76文字に制限されます。
// エンコードされたバイトは、Writerが閉じられるまで必ずしもフラッシュされません。
func (w *Writer) Write(p []byte) (n int, err error)

// CloseはWriterを閉じ、未書き込みのデータを基礎となるio.Writerにフラッシュしますが、
// 基礎となるio.Writerを閉じるわけではありません。
=======
// NewWriter returns a new [Writer] that writes to w.
func NewWriter(w io.Writer) *Writer

// Write encodes p using quoted-printable encoding and writes it to the
// underlying [io.Writer]. It limits line length to 76 characters. The encoded
// bytes are not necessarily flushed until the [Writer] is closed.
func (w *Writer) Write(p []byte) (n int, err error)

// Close closes the [Writer], flushing any unwritten data to the underlying
// [io.Writer], but does not close the underlying io.Writer.
>>>>>>> upstream/master
func (w *Writer) Close() error
