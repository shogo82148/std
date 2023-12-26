// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージquotedprintableは、RFC 2045で指定されているquoted-printableエンコーディングを実装します。
package quotedprintable

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// Readerは、quoted-printableデコーダーです。
type Reader struct {
	br   *bufio.Reader
	rerr error
	line []byte
}

// NewReaderは、rからデコードするquoted-printableリーダーを返します。
func NewReader(r io.Reader) *Reader

// Readは、基礎となるリーダーからquoted-printableデータを読み取り、デコードします。
func (r *Reader) Read(p []byte) (n int, err error)
