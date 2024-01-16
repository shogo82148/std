// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// Writerは、テキストプロトコルネットワーク接続にリクエストまたはレスポンスを書き込むための便利なメソッドを実装します。
type Writer struct {
	W   *bufio.Writer
	dot *dotWriter
}

// NewWriterはwに書き込む新しい [Writer] を返します。
func NewWriter(w *bufio.Writer) *Writer

// PrintfLineはフォーマットされた出力を\r\nに続けて書き込みます。
func (w *Writer) PrintfLine(format string, args ...any) error

// DotWriterは、wにドットエンコードを書き込むために使用できるライターを返します。
// 必要な場合に先行するドットを挿入し、改行文字 \n を \r\n に変換し、
// DotWriterが閉じられるときに最後の .\r\n 行を追加します。
// 次にwのメソッドを呼び出す前に、呼び出し元はDotWriterを閉じる必要があります。
//
// dot-encodingの詳細については、[Reader.DotReader] メソッドのドキュメントを参照してください。
func (w *Writer) DotWriter() io.WriteCloser
