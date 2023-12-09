// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iotest

import (
	"github.com/shogo82148/std/io"
)

<<<<<<< HEAD
// NewWriteLogger returns a writer that behaves like w except
// that it logs (using [log.Printf]) each write to standard error,
// printing the prefix and the hexadecimal data written.
func NewWriteLogger(prefix string, w io.Writer) io.Writer

// NewReadLogger returns a reader that behaves like r except
// that it logs (using [log.Printf]) each read to standard error,
// printing the prefix and the hexadecimal data read.
=======
// NewWriteLoggerは、wのように振る舞うライターを返しますが、
// 書き込みごとに（log.Printfを使用して）標準エラーにログを出力し、
// プレフィックスと書き込まれた16進数データを印刷します。
func NewWriteLogger(prefix string, w io.Writer) io.Writer

// NewReadLoggerは、rと同様に動作するリーダーを返しますが、
// 読み取りごとに（log.Printfを使用して）標準エラーにログを出力し、
// プレフィックスと読み取った16進数データを印刷します。
>>>>>>> release-branch.go1.21
func NewReadLogger(prefix string, r io.Reader) io.Reader
