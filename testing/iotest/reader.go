// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// iotestパッケージは、主にテストに役立つReaderとWriterを実装します。
package iotest

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

// OneByteReaderは、各非空のReadをrから1バイト読み取ることで実装するReaderを返します。
func OneByteReader(r io.Reader) io.Reader

// HalfReaderは、要求されたバイト数の半分をrから読み取ることでReadを実装するReaderを返します。
func HalfReader(r io.Reader) io.Reader

// DataErrReaderは、Readerによってエラーが処理される方法を変更します。通常、
// Readerは最後のデータが読み取られた後の最初のRead呼び出しからエラー（通常はEOF）を返します。
// DataErrReaderはReaderをラップし、最終的なエラーが最終的なデータとともに返されるように、
// その動作を変更します。最終データの後の最初の呼び出しではなく。
func DataErrReader(r io.Reader) io.Reader

// ErrTimeoutは、偽のタイムアウトエラーです。
var ErrTimeout = errors.New("timeout")

<<<<<<< HEAD
// TimeoutReader returns [ErrTimeout] on the second read
// with no data. Subsequent calls to read succeed.
func TimeoutReader(r io.Reader) io.Reader

// ErrReader returns an [io.Reader] that returns 0, err from all Read calls.
func ErrReader(err error) io.Reader

// TestReader tests that reading from r returns the expected file content.
// It does reads of different sizes, until EOF.
// If r implements [io.ReaderAt] or [io.Seeker], TestReader also checks
// that those operations behave as they should.
=======
// TimeoutReaderは、データなしの2回目の読み取りでErrTimeoutを返します。
// その後の読み取りの呼び出しは成功します。
func TimeoutReader(r io.Reader) io.Reader

// ErrReaderは、全てのRead呼び出しから0, errを返すio.Readerを返します。
func ErrReader(err error) io.Reader

// TestReaderは、rからの読み取りが期待されるファイル内容を返すことをテストします。
// EOFまで、異なるサイズの読み取りを行います。
// もしrがio.ReaderAtまたはio.Seekerを実装しているなら、TestReaderはまた、
// それらの操作が適切に動作することも確認します。
>>>>>>> release-branch.go1.21
//
// TestReaderが何かしらの不適切な動作を見つけた場合、それら全てを報告するエラーを返します。
// エラーテキストは複数行にわたります。
func TestReader(r io.Reader, content []byte) error
