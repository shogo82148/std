// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iotest

import (
	"github.com/shogo82148/std/io"
)

// NewWriteLoggerは、wのように振る舞うライターを返しますが、
// 書き込みごとに（[log.Printf] を使用して）標準エラーにログを出力し、
// プレフィックスと書き込まれた16進数データを印刷します。
func NewWriteLogger(prefix string, w io.Writer) io.Writer

// NewReadLoggerは、rと同様に動作するリーダーを返しますが、
// 読み取りごとに（[log.Printf] を使用して）標準エラーにログを出力し、
// プレフィックスと読み取った16進数データを印刷します。
func NewReadLogger(prefix string, r io.Reader) io.Reader
