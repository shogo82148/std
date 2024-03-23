// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multipart

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/textproto"
)

// Writerはマルチパートメッセージを生成します。
type Writer struct {
	w        io.Writer
	boundary string
	lastpart *part
}

// NewWriterは、ランダムな境界を持つ新しいマルチパート [Writer] を返し、wに書き込みます。
func NewWriter(w io.Writer) *Writer

// Boundaryは [Writer] の境界を返します。
func (w *Writer) Boundary() string

// SetBoundaryは、[Writer] のデフォルトのランダムに生成された
// 境界セパレータを明示的な値で上書きします。
//
// SetBoundaryはパートが作成される前に呼び出す必要があり、特定のASCII文字のみを
// 含むことができ、非空であり、最大で70バイトの長さでなければなりません。
func (w *Writer) SetBoundary(boundary string) error

// FormDataContentTypeは、この [Writer] のBoundaryを持つHTTP
// multipart/form-dataのContent-Typeを返します。
func (w *Writer) FormDataContentType() string

// CreatePartは、提供されたヘッダーを持つ新しいマルチパートセクションを作成します。
// パートのボディは、返された [Writer] に書き込むべきです。
// CreatePartを呼び出した後、以前のパートにはもう書き込むことができません。
func (w *Writer) CreatePart(header textproto.MIMEHeader) (io.Writer, error)

// CreateFormFileは、[Writer.CreatePart] の便利なラッパーです。これは、
// 提供されたフィールド名とファイル名で新しいform-dataヘッダーを作成します。
func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)

// CreateFormFieldは、与えられたフィールド名を使用してヘッダーを作成し、
// [Writer.CreatePart] を呼び出します。
func (w *Writer) CreateFormField(fieldname string) (io.Writer, error)

// WriteFieldは [Writer.CreateFormField] を呼び出し、その後で与えられた値を書き込みます。
func (w *Writer) WriteField(fieldname, value string) error

// Closeはマルチパートメッセージを終了し、終了境界線を出力に書き込みます。
func (w *Writer) Close() error
