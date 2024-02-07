// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/time"
)

// ResponseControllerは、HTTPハンドラーがレスポンスを制御するために使用されます。
//
<<<<<<< HEAD
// Handler.ServeHTTPメソッドが返された後にResponseControllerを使用することはできません。
=======
// A ResponseController may not be used after the [Handler.ServeHTTP] method has returned.
>>>>>>> upstream/release-branch.go1.22
type ResponseController struct {
	rw ResponseWriter
}

<<<<<<< HEAD
// NewResponseControllerは、リクエスト用のResponseControllerを作成します。
//
// ResponseWriterは、Handler.ServeHTTPメソッドに渡された元の値である必要があります。
// または、元のResponseWriterを返すUnwrapメソッドを持っている必要があります。
=======
// NewResponseController creates a [ResponseController] for a request.
//
// The ResponseWriter should be the original value passed to the [Handler.ServeHTTP] method,
// or have an Unwrap method returning the original ResponseWriter.
>>>>>>> upstream/release-branch.go1.22
//
// ResponseWriterが次のいずれかのメソッドを実装している場合、ResponseControllerは
// 適切に呼び出します。
//
//	Flush()
//	FlushError() error // エラーを返す代替Flush
//	Hijack() (net.Conn、*bufio.ReadWriter、error)
//	SetReadDeadline(deadline time.Time) error
//	SetWriteDeadline(deadline time.Time) error
//	EnableFullDuplex() error
//
<<<<<<< HEAD
// ResponseWriterがメソッドをサポートしていない場合、ResponseControllerは
// ErrNotSupportedに一致するエラーを返します。
=======
// If the ResponseWriter does not support a method, ResponseController returns
// an error matching [ErrNotSupported].
>>>>>>> upstream/release-branch.go1.22
func NewResponseController(rw ResponseWriter) *ResponseController

// Flushは、バッファリングされたデータをクライアントにフラッシュします。
func (c *ResponseController) Flush() error

// Hijackは、呼び出し元が接続を引き継ぐことを可能にします。
// 詳細については、Hijackerインターフェースを参照してください。
func (c *ResponseController) Hijack() (net.Conn, *bufio.ReadWriter, error)

// SetReadDeadlineは、ボディを含むリクエスト全体の読み取りの期限を設定します。
// 期限が超過した後にリクエストボディから読み取りを行うと、エラーが返されます。
// ゼロ値は期限がないことを意味します。
//
// 期限が超過した後に読み取り期限を設定しても、期限は延長されません。
func (c *ResponseController) SetReadDeadline(deadline time.Time) error

// SetWriteDeadlineは、レスポンスの書き込みの期限を設定します。
// 期限が超過した後にレスポンスボディに書き込みを行うと、ブロックされず、
// データがバッファリングされている場合は成功する可能性があります。
// ゼロ値は期限がないことを意味します。
//
// 期限が超過した後に書き込み期限を設定しても、期限は延長されません。
func (c *ResponseController) SetWriteDeadline(deadline time.Time) error

<<<<<<< HEAD
// EnableFullDuplexは、リクエストハンドラがRequest.Bodyからの読み取りを交互に行い、
// ResponseWriterへの書き込みと交互に行うことを示します。
=======
// EnableFullDuplex indicates that the request handler will interleave reads from [Request.Body]
// with writes to the [ResponseWriter].
>>>>>>> upstream/release-branch.go1.22
//
// HTTP/1リクエストの場合、Go HTTPサーバーはデフォルトで、レスポンスの書き込みを開始する前に
// リクエストボディの未読部分を消費し、ハンドラがリクエストから読み取りとレスポンスの書き込みを
// 同時に行うことを防止します。EnableFullDuplexを呼び出すと、この動作が無効になり、
// ハンドラがリクエストからの読み取りを続けながらレスポンスを同時に書き込むことができるようになります。
//
// HTTP/2リクエストの場合、Go HTTPサーバーは常に並行して読み取りとレスポンスを許可します。
func (c *ResponseController) EnableFullDuplex() error
