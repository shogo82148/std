// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package quotedprintable

import "github.com/shogo82148/std/io"

// Writerは、[io.WriteCloser] を実装するquoted-printableライターです。
type Writer struct {
	// バイナリモードでは、ライターの入力を純粋なバイナリとして扱い、
	// 行末のバイトをバイナリデータとして処理します。
	Binary bool

	w    io.Writer
	i    int
	line [78]byte
	cr   bool
}

// NewWriterは、wに書き込む新しい [Writer] を返します。
func NewWriter(w io.Writer) *Writer

// Writeは、pをquoted-printableエンコーディングでエンコードし、それを
// 基礎となる [io.Writer] に書き込みます。行の長さは76文字に制限されます。
// エンコードされたバイトは、[Writer] が閉じられるまで必ずしもフラッシュされません。
func (w *Writer) Write(p []byte) (n int, err error)

// Closeは [Writer] を閉じ、未書き込みのデータを基礎となる [io.Writer] にフラッシュしますが、
// 基礎となるio.Writerを閉じるわけではありません。
func (w *Writer) Close() error
