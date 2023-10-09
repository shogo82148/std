// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

import "github.com/shogo82148/std/io"

// StreamReaderはStreamをio.Readerにラップします。それは各データスライスを通過する際にXORKeyStreamを呼び出して処理します。
type StreamReader struct {
	S Stream
	R io.Reader
}

func (r StreamReader) Read(dst []byte) (n int, err error)

// StreamWriterはStreamをio.Writerにラップします。それはXORKeyStreamを呼び出して
// 通過するデータの各スライスを処理します。もしWrite呼び出しがshortを返す場合、
// StreamWriterは同期が取れておらず、破棄する必要があります。
// StreamWriterには内部のバッファリングはなく、データを書き込むためにCloseを呼び出す必要はありません。
type StreamWriter struct {
	S   Stream
	W   io.Writer
	Err error
}

func (w StreamWriter) Write(src []byte) (n int, err error)

// Closeは基礎となるWriterを閉じ、そのCloseの返り値を返します。Writerがio.Closerでもある場合は、それを返します。そうでなければnilを返します。
func (w StreamWriter) Close() error
