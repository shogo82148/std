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

<<<<<<< HEAD
// NewWriterは、ランダムな境界を持つ新しいマルチパートWriterを返し、wに書き込みます。
func NewWriter(w io.Writer) *Writer

// BoundaryはWriterの境界を返します。
func (w *Writer) Boundary() string

// SetBoundaryは、Writerのデフォルトのランダムに生成された
// 境界セパレータを明示的な値で上書きします。
=======
// NewWriter returns a new multipart [Writer] with a random boundary,
// writing to w.
func NewWriter(w io.Writer) *Writer

// Boundary returns the [Writer]'s boundary.
func (w *Writer) Boundary() string

// SetBoundary overrides the [Writer]'s default randomly-generated
// boundary separator with an explicit value.
>>>>>>> upstream/master
//
// SetBoundaryはパートが作成される前に呼び出す必要があり、特定のASCII文字のみを
// 含むことができ、非空であり、最大で70バイトの長さでなければなりません。
func (w *Writer) SetBoundary(boundary string) error

<<<<<<< HEAD
// FormDataContentTypeは、このWriterのBoundaryを持つHTTP
// multipart/form-dataのContent-Typeを返します。
func (w *Writer) FormDataContentType() string

// CreatePartは、提供されたヘッダーを持つ新しいマルチパートセクションを作成します。
// パートのボディは、返されたWriterに書き込むべきです。
// CreatePartを呼び出した後、以前のパートにはもう書き込むことができません。
func (w *Writer) CreatePart(header textproto.MIMEHeader) (io.Writer, error)

// CreateFormFileは、CreatePartの便利なラッパーです。これは、
// 提供されたフィールド名とファイル名で新しいform-dataヘッダーを作成します。
func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)

// CreateFormFieldは、与えられたフィールド名を使用してヘッダーを作成し、
// CreatePartを呼び出します。
func (w *Writer) CreateFormField(fieldname string) (io.Writer, error)

// WriteFieldはCreateFormFieldを呼び出し、その後で与えられた値を書き込みます。
=======
// FormDataContentType returns the Content-Type for an HTTP
// multipart/form-data with this [Writer]'s Boundary.
func (w *Writer) FormDataContentType() string

// CreatePart creates a new multipart section with the provided
// header. The body of the part should be written to the returned
// [Writer]. After calling CreatePart, any previous part may no longer
// be written to.
func (w *Writer) CreatePart(header textproto.MIMEHeader) (io.Writer, error)

// CreateFormFile is a convenience wrapper around [Writer.CreatePart]. It creates
// a new form-data header with the provided field name and file name.
func (w *Writer) CreateFormFile(fieldname, filename string) (io.Writer, error)

// CreateFormField calls [Writer.CreatePart] with a header using the
// given field name.
func (w *Writer) CreateFormField(fieldname string) (io.Writer, error)

// WriteField calls [Writer.CreateFormField] and then writes the given value.
>>>>>>> upstream/master
func (w *Writer) WriteField(fieldname, value string) error

// Closeはマルチパートメッセージを終了し、終了境界線を出力に書き込みます。
func (w *Writer) Close() error
